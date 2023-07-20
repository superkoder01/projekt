package impl

import (
	"bytes"
	"encoding/json"
	"github.com/op/go-logging"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/api/model/blockchain_user"
	conf "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/internal/configuration"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/internal/email"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/internal/error"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao/entity"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/billing_dao/mysql"
	"io/ioutil"
	"net"
	"net/http"
	"strings"
	"time"
)

const (
	HTTP = "http://"
	JSON = "application/json"
)

var (
	logger = logging.MustGetLogger("service")
)

var (
	userFields = []string{"login", "email"}
)

type BlockchainAccount struct {
	UserName   string `json:"userName"`
	PublicKey  string `json:"publicKey"`
	AccAddress string `json:"accAddress"`
}

func attachActivationCode(en *entity.User, randomCode string) error {
	logger.Debugf("Attaching activation code to user entity")
	en.SetActivationCode(randomCode)

	return nil
}

func sendEmailNotificationOnUserCreation(en entity.Entity) error {
	logger.Debugf("Sending email notification on user creation")

	user, ok := en.(*entity.User)
	if !ok {
		return e.ErrInternalServerError
	}

	if strings.TrimSpace(user.Email) == "" {
		return e.ApiErrInvalidEmail
	}

	eIn := email.NewEmailInput(user.Email, user.Login, user.ActivationCode)
	return eIn.BuildActivationMessage().Send()
}

func sendEmailNotificationOnPasswordForgot(user *entity.User) error {
	logger.Debugf("Sending email notification on password forgot")

	if strings.TrimSpace(user.Email) == "" {
		return e.ApiErrInvalidEmail
	}

	eIn := email.NewEmailInput(user.Email, user.Login, user.ActivationCode)
	return eIn.BuildPasswordResetMessage().Send()
}

func createUserEntity(tx mysql.Session, en entity.Entity, activationCode string) error {
	user, ok := en.(*entity.User)
	if !ok {
		return e.ApiErrInvalidDataModel
	}

	var err error
	if err = attachActivationCode(user, activationCode); err != nil {
		return err
	}
	user.AddedDate = time.Now()

	// TODO uncomment when blockchain will be ready
	//var blockchainAccount *BlockchainAccount
	//if blockchainAccount, err = createBlockchainAccount(user.Email); err != nil {
	//	return err
	//}
	//user.SetBlockchainPubKey(blockchainAccount.PublicKey)
	//user.SetBlockchainAccAddress(blockchainAccount.AccAddress)

	if err = tx.Create(en).Error(); err != nil {
		return err
	}

	if err = sendEmailNotificationOnUserCreation(user); err != nil {
		return err
	}

	return nil
}

func createBlockchainAccount(email string) (*BlockchainAccount, error) {
	logger.Debugf("Creating blockchain account for: %s", email)

	userBytes, err := json.Marshal(&blockchain_user.BlockchainUser{
		Name: email,
	})
	if err != nil {
		return nil, err
	}

	// Get blockchain config
	bc := conf.GetBlockchainConfig()
	// Post create blockchain account request
	resp, err := http.Post(HTTP+net.JoinHostPort(bc.AdapterHost, bc.AdapterPort)+bc.Endpoint, JSON, bytes.NewBuffer(userBytes))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return nil, e.BlockchainErrAccountCreation
	}

	// Read body value
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, e.BlockchainErrAccountCreation
	}

	// Unmarshal []bytes into BlockchainAccount struct
	ba := BlockchainAccount{}
	if err = json.Unmarshal(body, &ba); err != nil {
		return nil, e.BlockchainErrAccountInvalid
	}

	return &ba, nil
}

// generateRegistrationCode returns twelve digits random number
func generateRegistrationCode() int64 {
	return time.Now().UnixNano() / (1 << 22)
}

func isOnList(record string, list []string) bool {
	listMap := make(map[string]struct{})
	for _, v := range list {
		listMap[v] = struct{}{}
	}
	_, ok := listMap[record]
	return ok
}
