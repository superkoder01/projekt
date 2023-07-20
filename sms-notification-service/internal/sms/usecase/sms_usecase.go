package usecase

import (
	"NotificationSmsService/internal/domain/model"
	"NotificationSmsService/internal/sms/gateway"
	"NotificationSmsService/pkg/logger"
	customError "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/notification-service-rabbitmq.git/domain/error"
	"strings"
)

type SmsUseCase struct {
	gateway gateway.SmsGateway
	logger  logger.Logger
}

func NewSmsUseCase(gateway gateway.SmsGateway, logger logger.Logger) *SmsUseCase {
	return &SmsUseCase{
		gateway: gateway,
		logger:  logger,
	}
}

func (smsUseCase *SmsUseCase) SendSms(message []byte, messageId, workerName string) error {
	if len(message) < 1 {
		smsUseCase.logger.Errorf("%v Message is empty !", messageId)
		return customError.IncorrectMessage{Service: "SMS-Service", Reason: "Message is empty ! Cannot decode !", Identifier: messageId}
	}

	smsWrapper, err := model.CreateWrappedSms(message, messageId, workerName)
	identifier := smsWrapper.GetMessageIdentifiers()
	if err != nil {
		smsUseCase.logger.Errorf("%v Cannot decode json message !", identifier)
		return customError.IncorrectMessage{Service: "SMS-Service", Reason: "Invalid message format ! Cannot decode !", Identifier: identifier}
	}

	unwrappedSms := smsWrapper.Message
	if len(unwrappedSms.Msisdn) < 1 {
		return customError.IncorrectMessage{Service: "SMS-Service", Reason: "No receiver for message", Identifier: identifier}
	}

	numberNormalization(unwrappedSms.Msisdn)

	if err := smsUseCase.gateway.Send(smsWrapper); err != nil {
		smsUseCase.logger.Errorf("%v Failed to send sms ! %v", identifier, err)
		return err
	}

	return nil
}

func numberNormalization(msisdn []string) {
	for _, number := range msisdn {
		if !strings.HasPrefix(number, "0048") {
			number = "0048" + number
		}
	}
}
