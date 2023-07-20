package email

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/op/go-logging"
	"github.com/streadway/amqp"
	conf "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/internal/configuration"
)

var (
	logger = logging.MustGetLogger("email")
)

type email struct {
	Destination          []string `json:"destination"`
	CopyDestination      []string `json:"cc-destination,omitempty"`
	BlindCopyDestination []string `json:"bcc-destination,omitempty"`
	Title                string   `json:"title"`
	Body                 string   `json:"body,omitempty"`
}

type emailInput struct {
	Destination    string
	Login          string
	ActivationCode string
}

func NewEmailInput(destination string, login string, activationCode string) *emailInput {
	return &emailInput{Destination: destination, Login: login, ActivationCode: activationCode}
}

func (e *email) Send() error {
	email := conf.GetEmailNotificationConfig()
	source := fmt.Sprintf("amqp://%s:%s@%s:%s", email.User, email.Password, email.Host, email.Port)
	logger.Debugf("Sending email notification to: %s", e.Destination)
	connection, err := amqp.Dial(source)
	if err != nil {
		return err
	}

	channel, err := connection.Channel()
	if err != nil {
		return err
	}

	b, err := json.Marshal(e)
	if err != nil {
		return err
	}

	err = channel.Publish(
		"",
		email.Queue,
		false,
		false,
		amqp.Publishing{
			DeliveryMode: amqp.Persistent,
			ContentType:  "application/json",
			Body:         b,
			MessageId:    uuid.New().String(),
		})
	if err != nil {
		return err
	}

	logger.Debugf("Email notification sent to: %s", e.Destination)

	return nil
}

func (e *emailInput) BuildActivationMessage() *email {
	eConf := conf.GetActivationMessageConfig()
	return &email{
		Destination: []string{e.Destination},
		Title:       fmt.Sprintf(eConf.Title, e.Login),
		Body:        fmt.Sprintf(eConf.Body, e.ActivationCode),
	}
}

func (e *emailInput) BuildPasswordResetMessage() *email {
	eConf := conf.GetPasswordResetMessageConfig()
	return &email{
		Destination: []string{e.Destination},
		Title:       eConf.Title,
		Body:        fmt.Sprintf(eConf.Body, e.ActivationCode),
	}
}
