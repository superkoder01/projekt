package impl

import (
	"bytes"
	"crypto/hmac"
	"crypto/sha256"
	b64 "encoding/base64"
	"encoding/hex"
	"encoding/json"
	"io/ioutil"
	"mime/multipart"
	"net"
	"net/http"
	"regexp"
	"strconv"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mongo/model/billing_management/contract"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/conversion-service.git/pkg/conversion"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/signer-service.git/internal/configuration"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/signer-service.git/internal/error"
	http_utils "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/signer-service.git/internal/utils/http"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/signer-service.git/internal/utils/http/mSzafir"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/signer-service.git/service/billing_management"
	"go.uber.org/zap"
)

const (
	HTTP         = "http://"
	StatusOk     = "OK"
	StatusSigned = "SIGNED"
	FilePrefix   = "/signed_contract_"
	ServiceName  = "Signer-Service"
)

type signerService struct {
	logger        *zap.SugaredLogger
	conversion    conversion.Client
	mSzafirConfig *configuration.MSzafirConfig
}

func NewReportService(logger *zap.SugaredLogger, conversion conversion.Client, mSzafirConfig *configuration.MSzafirConfig) *signerService {
	return &signerService{logger: logger, conversion: conversion, mSzafirConfig: mSzafirConfig}
}

func getLastSignedContractName(contractModel contract.Contract, signedContractPath string) string {

	//// TODO: Handle Payload.SigningDetails of authorisation json type.

	var timestamp time.Time
	signingId := ""
	for _, value := range contractModel.Payload.SigningDetails {
		if strings.ToUpper(value.SigningStatus) == StatusSigned {
			if value.Timestamp.After(timestamp) {
				timestamp = value.Timestamp
				signingId = value.SigningId
			}
		}
	}

	if signingId != "" {
		return signedContractPath + FilePrefix + signingId + ".pdf"
	}
	return ""
}

//// TODO: Change function signature to accept not only contract or create new, additional function

func (s *signerService) InitSign(ctx *gin.Context, body contract.Contract) (string, error) {
	signedContract := getLastSignedContractName(body, s.mSzafirConfig.SignedContractsPath)

	var pdf []byte
	var err error

	if signedContract != "" {
		s.logger.Infof("Read signed contract before: %s", signedContract)
		pdf, err = ioutil.ReadFile(signedContract)
	} else {
		s.logger.Infof("Converting new contract id: %s", body.Id)
		pdf, _, err = s.conversion.Convert(&body)
	}

	if err != nil {
		s.logger.Errorf("Cannot read file ! Service error %v", err)
		if conversion.IsConnectionError(err) {
			return "", e.ApiErrEndpointForbidden
		}
		return "", err
	} else {
		s.logger.Infof("Initializing signing procedure")
		hauth, err := EncodeToBase64(PrepareHauth(s.mSzafirConfig))
		if err != nil {
			s.logger.Errorf("Cannot Encode To Base64. %s", err.Error())
			return "", e.ApiErrEndpointForbidden
		}
		s.logger.Debugf("init hauth: " + hauth)

		url := s.mSzafirConfig.Prefix + "init?hauth=" + hauth
		request, err := NewFileUploadRequest(HTTP+net.JoinHostPort(s.mSzafirConfig.Host, s.mSzafirConfig.Port)+url, "files", pdf)

		if err != nil {
			s.logger.Errorf("Cannot create init request. %s", err.Error())
			return "", e.ApiErrEndpointForbidden
		}
		client := &http.Client{}
		resp, err := client.Do(request)
		if err != nil {
			s.logger.Errorf("Cannot send init request. %s", err.Error())
			return "", e.ApiErrEndpointForbidden
		} else {
			response, err := ioutil.ReadAll(resp.Body)
			if err != nil {
				s.logger.Errorf("Cannot read init response. %s", err.Error())
				return "", e.ApiErrEndpointForbidden
			}

			var details mSzafir.InitResponse
			if err = json.Unmarshal(response, &details); err != nil {
				s.logger.Errorf("Cannot convert init response. %s", err.Error())
				return "", e.ApiErrEndpointForbidden
			}

			s.logger.Debugf("SigningId: " + details.SigningId)
			s.logger.Debugf("Status: " + details.Status)
			fileIds := strings.Join(details.FileIds, "")
			s.logger.Debugf("FileIds: " + fileIds)

			loginHauth, err := EncodeToBase64(PrepareLoginHauth(details.SigningId, details.FileIds, s.mSzafirConfig.Password))
			if err != nil {
				s.logger.Errorf("Cannot Encode To Base64. %s", err.Error())
				return "", e.ApiErrEndpointForbidden
			}
			s.logger.Debugf("Login hauth: " + loginHauth)
			ctx.Header("signingId", details.SigningId)
			ctx.Header("fileId", fileIds)
			return loginHauth, nil
		}
	}

	return "", nil
}

func (s *signerService) SigningCompletedNotification(ctx *gin.Context) error {
	s.logger.Debugf("SigningCompletedNotification")

	err := ctx.Request.ParseForm()
	if err != nil {
		return e.ApiErrEndpointForbidden
	}
	signingId := ctx.Request.Form.Get("signingId")
	s.logger.Debugf("Got SigningId: " + signingId)

	s.logger.Debugf("Getting status")
	resp := http_utils.Get(ctx, s.mSzafirConfig.Host, s.mSzafirConfig.Port, s.mSzafirConfig.Prefix+"status/"+signingId)

	response, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		s.logger.Errorf("Cannot read status response. %s", err.Error())
		return e.ApiErrEndpointForbidden
	}

	var details mSzafir.StatusResponse
	if err = json.Unmarshal(response, &details); err != nil {
		s.logger.Errorf("Cannot convert status response. %s", err.Error())
		return e.ApiErrEndpointForbidden
	}

	s.logger.Infof("Got status response")
	s.logger.Infof("Completed: %t", details.Completed)
	s.logger.Infof("Status: %s", details.Status)
	s.logger.Infof("StatusCode: %d", details.StatusCode)
	s.logger.Infof("FileUrls: %s", details.FileUrls)

	if details.Status == StatusOk {
		regexp := regexp.MustCompile(`://.*/mini-portal`)
		details.FileUrls[0] = regexp.ReplaceAllString(details.FileUrls[0], "://"+s.mSzafirConfig.Host+"/mini-portal")
		s.logger.Infof("Replaced fileUrls to: %s", details.FileUrls)
		billing_management.UpdateContractSigningStatus(ctx, signingId, &details)
	} else {
		s.logger.Errorf("Signing Status Error")
	}

	return nil
}

func PrepareLoginHauth(signingId string, fileIds []string, password string) *mSzafir.LoginHauth {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)
	hmacData := strconv.FormatInt(timestamp, 10) + signingId + strings.Join(fileIds[:], "")

	key := []byte(password)
	sig := hmac.New(sha256.New, key)
	sig.Write([]byte(hmacData))
	Hmac := hex.EncodeToString(sig.Sum(nil))

	return &mSzafir.LoginHauth{
		Timestamp: timestamp,
		SigningId: signingId,
		FileIds:   fileIds,
		Hmac:      Hmac,
	}

}

func PrepareHauth(config *configuration.MSzafirConfig) *mSzafir.Hauth {
	timestamp := time.Now().UnixNano() / int64(time.Millisecond)

	hmacData := strconv.FormatInt(timestamp, 10) + config.Mode + config.UrlSigningCompleted + config.UrlSigningCompletedNotification

	key := []byte(config.Password)
	sig := hmac.New(sha256.New, key)
	sig.Write([]byte(hmacData))

	Hmac := hex.EncodeToString(sig.Sum(nil))

	return &mSzafir.Hauth{
		Timestamp:                       timestamp,
		Mode:                            config.Mode,
		UrlSigningCompleted:             config.UrlSigningCompleted,
		UrlSigningCompletedNotification: config.UrlSigningCompletedNotification,
		Hmac:                            Hmac,
	}
}

func EncodeToBase64(v interface{}) (string, error) {
	var buf bytes.Buffer
	encoder := b64.NewEncoder(b64.StdEncoding, &buf)
	err := json.NewEncoder(encoder).Encode(v)
	if err != nil {
		return "", err
	}
	encoder.Close()
	return buf.String(), nil
}

func NewFileUploadRequest(uri string, paramName string, pdf []byte) (*http.Request, error) {

	body := new(bytes.Buffer)
	writer := multipart.NewWriter(body)
	part, err := writer.CreateFormFile(paramName, "files")
	if err != nil {
		return nil, err
	}
	part.Write(pdf)

	err = writer.Close()
	if err != nil {
		return nil, err
	}

	test, _ := http.NewRequest("POST", uri, body)
	test.Header.Set("Content-Type", writer.FormDataContentType())
	return test, nil
}
