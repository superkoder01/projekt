package gateway

import (
	"NotificationSmsService/internal/domain/model"
	"NotificationSmsService/pkg/config"
	"NotificationSmsService/pkg/logger"
	"errors"
	customError "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/notification-service-rabbitmq.git/domain/error"
	"net/http"
	"strings"
)

type SmsGateway struct {
	config       config.GatewayConfig
	logger       logger.Logger
	apiUrl       string
	senderId     string
	methodType   string
	url          string
	apiVersion   string
	apiServiceId string
	apiKey       string
}

func NewGateway(config config.GatewayConfig, logger logger.Logger) *SmsGateway {
	return &SmsGateway{
		config:       config,
		logger:       logger,
		apiUrl:       config.ApiUrl,
		senderId:     config.SenderId,
		methodType:   config.MethodType,
		url:          config.Host,
		apiVersion:   config.ApiVersion,
		apiServiceId: config.ApiServiceId,
		apiKey:       config.ApiKey,
	}
}

func (gateway *SmsGateway) Send(wrapper *model.SmsWrapper) error {
	identifier := wrapper.GetMessageIdentifiers()

	convertedSms, err := model.SerializeSmsToXml(wrapper, gateway.apiUrl, gateway.senderId)
	if err != nil {
		gateway.logger.Errorf("%v Cannot convert sms: %v. %v", identifier, wrapper.Message, err)
		return customError.IncorrectMessage{Service: "SMS-Service", Reason: "Invalid message format ! Cannot convert to xml!", Identifier: identifier}
	}

	request, err := gateway.createHttpRequest(string(convertedSms[:]))
	if err != nil {
		gateway.logger.Errorf("Cannot prepare HTTP request with SMS!", err)
		return err
	}
	client := http.Client{}

	gateway.logger.Infof("%v Sending POST request: %v", identifier, request)
	gateway.logger.Debugf("%v Request body: %v", identifier, string(convertedSms[:]))
	do, err := client.Do(request)
	gateway.logger.Infof("%v Request response statusCode: %v", identifier, do.Status)

	if err != nil {
		gateway.logger.Errorf("%v Unexpected request error !", identifier)
		return err
	} else {
		if do.StatusCode == 200 {
			return nil
		} else if do.StatusCode == 400 {
			gateway.logger.Errorf("%v Fatal Error invalid gateway configuration !! ", identifier)
			return customError.MessageServiceError{Service: "SMS-Service", Reason: "Gateway refuse to process request", Identifier: identifier}
		} else if do.StatusCode == 500 {
			gateway.logger.Errorf("Gateway error ! Retrying scheduled")
			return customError.NotReadyMessage{Service: "SMS-Service", Reason: "Server exception", Identifier: identifier}
		} else {
			gateway.logger.Errorf("%v Unknown sms-gateway response, statusCode: %v", identifier, do.StatusCode)
			return errors.New("invalid gateway response")
		}
	}
}

func (gateway *SmsGateway) createHttpRequest(sms string) (*http.Request, error) {
	request, err := http.NewRequest(gateway.methodType, gateway.url, strings.NewReader(sms))
	if err != nil {
		return &http.Request{}, err
	}
	request.Header = http.Header{
		"Content-Type":   []string{"application/xml"},
		"API_VERSION":    []string{gateway.apiVersion},
		"API_SERVICE_ID": []string{gateway.apiServiceId},
		"API_KEY":        []string{gateway.apiKey},
	}
	return request, nil
}
