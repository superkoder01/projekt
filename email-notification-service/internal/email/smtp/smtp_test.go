package smtp

import (
	"NotificationEmailService/config"
	"NotificationEmailService/internal/logger"
	"go.uber.org/zap"
)

var cfg config.SmtpConfig
var appLogger *zap.SugaredLogger

func init() {
	cfg = config.SmtpConfig{
		User: "test@test.pl",
	}
	appLogger = logger.NewApiLogger(&config.Logger{})
}

//func TestEmailConversionWithWalidEmail(t *testing.T) {
//	mailer := New(cfg, appLogger)
//
//	validMail := &email.Email{
//		From: "<test@test.pl>",
//		Text: []byte("sample text"),
//		To:   []string{"test@tester.pl"},
//	}
//
//	result, err := mailer.convertEmail(
//		&model.Email{
//			Body:        "sample text",
//			Destination: []string{"test@tester.pl"},
//		}, []byte(""), "identifier",
//	)
//	require.Equal(t, validMail, result)
//	require.NoError(t, err)
//}
//
//func TestEmptyEmailConversion(t *testing.T) {
//	mailer := New(cfg, appLogger)
//
//	_, err := mailer.convertEmail(&model.Email{}, []byte(""), "identifier")
//	require.Error(t, err, "invalid email content")
//}
//
//func TestEmptyEmailConversion_2(t *testing.T) {
//	mailer := New(cfg, appLogger)
//
//	_, err := mailer.convertEmail(&model.Email{Destination: []string{"valid@email.pl"}}, []byte(""), "identifier")
//	require.Error(t, err, "invalid email content")
//}
//
//func TestEmailConversionWithWalidEmailWithFile(t *testing.T) {
//	mailer := New(cfg, appLogger)
//
//	validMail := &email.Email{
//		From: "<test@test.pl>",
//		Text: []byte("sample text"),
//		To:   []string{"test@tester.pl"},
//	}
//	_, err := validMail.AttachFile("../../../config/config-local-email.yaml")
//
//	result, err := mailer.convertEmail(
//		&model.Email{
//			Body:        "sample text",
//			Destination: []string{"test@tester.pl"},
//			//FileName:    "../../../config/config-local-email.yaml",
//		}, []byte(""), "identifier",
//	)
//	require.Equal(t, validMail, result)
//	require.NoError(t, err)
//}
//
//func TestEmailConversionWithWalidEmailWithNotExistingFile(t *testing.T) {
//	mailer := New(cfg, appLogger)
//
//	_, err := mailer.convertEmail(
//		&model.Email{
//			Body:        "sample text",
//			Destination: []string{"test@tester.pl"},
//			//FileName:    "someRandomFileName123",
//		}, []byte(""), "identifier",
//	)
//	require.Error(t, err)
//}
