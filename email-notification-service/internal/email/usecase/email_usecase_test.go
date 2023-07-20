package usecase

import (
	"NotificationEmailService/config"
	"NotificationEmailService/internal/domain/model"
	"NotificationEmailService/internal/logger"
	"NotificationEmailService/internal/test"
	"NotificationEmailService/pkg/email"
	"errors"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/conversion-service.git/pkg/conversion"
	rabbitMqError "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/notification-service-rabbitmq.git/domain/error"
	rabbitMqDomain "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/notification-service-rabbitmq.git/domain/model"
	"testing"
)

func TestEmailUseCase_ValidateEmailAddress(t *testing.T) {
	validEmail := []string{"piotr.piasecki@ovoo.pl", "pp@pp.pp.pl"}
	emailDesignationAddresses := []string{"piotr.piasecki@ovoo.pl", "piotrek", "piotrek@", "", "1213@", "pp@pp.pp.pl"}

	receivers := removeInvalidEmailReceivers(emailDesignationAddresses, logger.NewApiLogger(&config.Logger{}), "identifier")
	require.Len(t, receivers, len(validEmail))
	require.Equal(t, receivers, validEmail)
}

func TestSendMailEmptyMessage(t *testing.T) {
	e, err := sendByBytes([]byte{}, nil, nil)
	assert.Equal(t, "support@ovoo.pl", e.Destination[0])
	assert.Equal(t, "Email notification service error", e.Title)
	assert.Equal(t, rabbitMqError.IncorrectMessage{Service: "Email-Service", Reason: "Empty message", Identifier: "messageId"}, err)
}

func TestSendMailEmptyBody(t *testing.T) {
	e, err := sendByEmail(t, &rabbitMqDomain.Email{}, nil, nil)
	assert.Equal(t, "support@ovoo.pl", e.Destination[0])
	assert.Equal(t, "Email notification service error", e.Title)
	assert.Equal(t, rabbitMqError.IncorrectMessage{Service: "Email-Service", Reason: "Empty body", Identifier: "[workerName:messageId]"}, err)
}

func TestSendMailEmptyReceivers(t *testing.T) {
	e, err := sendByEmail(t, &rabbitMqDomain.Email{Body: "body"}, nil, nil)
	assert.Equal(t, "support@ovoo.pl", e.Destination[0])
	assert.Equal(t, "Email notification service error", e.Title)
	assert.Equal(t, rabbitMqError.IncorrectMessage{Service: "Email-Service", Reason: "Empty receivers", Identifier: "[workerName:messageId]"}, err)
}

func TestSendMailConversionCommunicationError(t *testing.T) {
	e, err := sendByEmail(t, &rabbitMqDomain.Email{
		Body:        "body",
		Destination: []string{"test@ovoo.pl"},
		Attachments: test.CreateAttachments("a", "a")}, conversion.NewConnectionErrorByDetails("test-error"), nil)
	assert.Nil(t, e)
	assert.Equal(t, rabbitMqError.NotReadyMessage{Service: "Email-Service", Reason: "Cannot generate file. err: test-error", Identifier: "[workerName:messageId]"}, err)
}

func TestSendMailConversionInternalError(t *testing.T) {
	e, err := sendByEmail(t, &rabbitMqDomain.Email{
		Body:        "body",
		Destination: []string{"test@ovoo.pl"},
		Attachments: test.CreateAttachments("a", "a")}, conversion.NewInternalErrorByDetails("test-error"), nil)
	assert.Equal(t, "support@ovoo.pl", e.Destination[0])
	assert.Equal(t, "Email notification service error", e.Title)
	assert.Equal(t, conversion.NewInternalErrorByDetails("test-error"), err)
}

func TestSendMailSendNotReadyMessageError(t *testing.T) {
	sendError := rabbitMqError.NotReadyMessage{Service: "Email-Service", Reason: "Cannot send email. err: test-error", Identifier: "[workerName:messageId]"}
	e, err := sendByEmail(t, &rabbitMqDomain.Email{Body: "body", Destination: []string{"test@ovoo.pl"}}, nil, sendError)
	assert.Nil(t, e)
	assert.Equal(t, sendError, err)
}

func TestSendMailSendError(t *testing.T) {
	sendError := errors.New("test-error")
	e, err := sendByEmail(t, &rabbitMqDomain.Email{Body: "body", Destination: []string{"test@ovoo.pl"}}, nil, sendError)
	assert.Equal(t, "support@ovoo.pl", e.Destination[0])
	assert.Equal(t, "Email notification service error", e.Title)
	assert.Equal(t, sendError, err)
}

func TestSendMail(t *testing.T) {
	e, err := sendByEmail(t, &rabbitMqDomain.Email{Body: "body", Destination: []string{"test@ovoo.pl"}, Title: "title"}, nil, nil)
	assert.Equal(t, "test@ovoo.pl", e.Destination[0])
	assert.Equal(t, "title", e.Title)
	assert.Equal(t, "body", e.Body)
	assert.Nil(t, err)
}

func sendByEmail(t *testing.T, e *rabbitMqDomain.Email, conversionError error, smtpError error) (*rabbitMqDomain.Email, error) {
	bytes, err := email.ToBytes(e)
	assert.Nil(t, err)
	return sendByBytes(bytes, conversionError, smtpError)
}

func sendByBytes(message []byte, conversionError error, smtpError error) (*rabbitMqDomain.Email, error) {
	smtpClient := mockSmtp{error: smtpError}
	conversionClient := mockConversion{error: conversionError}
	err := NewEmailUseCase(&smtpClient, logger.NewApiLogger(&config.Logger{}), &conversionClient, &config.ServiceConfig{
		ErrorEmailTo: []string{"support@ovoo.pl"},
	}).SendEmail(message, "messageId", "workerName")
	return smtpClient.email, err
}

type mockSmtp struct {
	email *rabbitMqDomain.Email
	error error
}

func (s *mockSmtp) Send(email *model.EmailWrapper) error {
	if s.error == nil {
		s.email = email.Message
		return nil
	} else {
		err := s.error
		s.error = nil
		return err
	}
}

type mockConversion struct {
	error error
}

func (c *mockConversion) Convert(any) ([]byte, string, error) {
	return nil, "", c.error
}
