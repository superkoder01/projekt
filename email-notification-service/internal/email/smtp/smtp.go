package smtp

import (
	"NotificationEmailService/config"
	"NotificationEmailService/config/smtp_security"
	"NotificationEmailService/internal/domain/model"
	"bytes"
	"crypto/tls"
	"fmt"
	"github.com/jordan-wright/email"
	customError "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/notification-service-rabbitmq.git/domain/error"
	rabbitMqModel "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/notification-service-rabbitmq.git/domain/model"
	"go.uber.org/zap"
	"mime"
	netSmtp "net/smtp"
)

type Smtp interface {
	Send(email *model.EmailWrapper) error
}

type smtp struct {
	config config.SmtpConfig
	logger *zap.SugaredLogger
}

func New(config config.SmtpConfig, logger *zap.SugaredLogger) *smtp {
	return &smtp{
		config: config,
		logger: logger,
	}
}

func (s *smtp) authorize() netSmtp.Auth {
	s.logger.Debugf("UserName: %v, Password: %v", s.config.User, s.config.Password)
	return netSmtp.PlainAuth("", s.config.User, s.config.Password, s.config.Host)
}

func (s *smtp) convertEmail(message *rabbitMqModel.Email, attachments map[string][]byte, identifier string) (*email.Email, error) {
	e := &email.Email{
		From:    s.config.SenderName + "<" + s.config.User + ">",
		To:      message.Destination,
		Cc:      message.CCDestination,
		Bcc:     message.BCCDestination,
		Subject: message.Title,
		Text:    []byte(message.Body),
	}

	fileExtension := ".pdf"
	for fileName, attachment := range attachments {
		_, err := e.Attach(bytes.NewReader(attachment), fileName+fileExtension, mime.TypeByExtension("application/pdf"))
		if err != nil {
			s.logger.Warnf("%v Cannot attache file %v", identifier, err)
			return &email.Email{}, err
		}
	}

	return e, nil
}

func (s *smtp) Send(wrapper *model.EmailWrapper) error {
	identifier := wrapper.GetMessageIdentifiers()
	e, err := s.convertEmail(wrapper.Message, wrapper.FileAttachment, identifier)
	if err != nil {
		s.logger.Errorf("%v Cannot send mail: %v", identifier, err)
		return customError.IncorrectMessage{Service: "Email-Service", Reason: "Invalid message format ! Cannot convert!"}
	}

	s.logger.Infof("%v Sending email message: %v", identifier, e)
	s.logger.Debugf("%v Email text: %v", identifier, string(e.Text[:]))

	sendErr := s.send(s.config.Host+":"+s.config.Port, e)
	if sendErr != nil {
		s.logger.Errorf("%v Cannot send mail: %v", identifier, err)
		return customError.NotReadyMessage{Service: "Email-Service", Reason: fmt.Sprintf("Cannot send mail. err: %v", err), Identifier: identifier}
	}

	return nil
}

func (s *smtp) send(addr string, data *email.Email) error {
	switch s.config.Security {
	case smtp_security.None:
		return data.Send(addr, s.authorize())
	case smtp_security.TLS:
		return data.SendWithTLS(addr, s.authorize(), &tls.Config{
			InsecureSkipVerify: true,
			ServerName:         s.config.Host,
		})
	case smtp_security.STARTTLS:
		return data.SendWithStartTLS(addr, s.authorize(), &tls.Config{
			InsecureSkipVerify: true,
			ServerName:         s.config.Host,
		})
	default:
		return data.Send(addr, s.authorize())
	}
}
