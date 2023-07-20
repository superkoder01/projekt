package usecase

import (
	"NotificationEmailService/config"
	"NotificationEmailService/internal/domain/model"
	"NotificationEmailService/internal/email/smtp"
	"encoding/json"
	"fmt"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/conversion-service.git/pkg/conversion"
	customError "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/notification-service-rabbitmq.git/domain/error"
	rabbitMqDomain "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/notification-service-rabbitmq.git/domain/model"
	"go.uber.org/zap"
)

type EmailUseCase struct {
	smtp       smtp.Smtp
	logger     *zap.SugaredLogger
	conversion conversion.Client
	config     *config.ServiceConfig
}

func NewEmailUseCase(smtp smtp.Smtp, logger *zap.SugaredLogger, conversion conversion.Client, config *config.ServiceConfig) *EmailUseCase {
	return &EmailUseCase{
		smtp:       smtp,
		logger:     logger,
		conversion: conversion,
		config:     config,
	}
}

func (emailUseCase EmailUseCase) SendEmail(message []byte, messageId, workerName string) error {
	if len(message) < 1 {
		emailUseCase.logger.Errorf("%v Empty message", messageId)
		return returnError(
			emailUseCase.sendErrorEmailByBody(fmt.Sprintf("Error: Empty message")),
			customError.IncorrectMessage{Service: "Email-Service", Reason: "Empty message", Identifier: messageId})
	}

	emailWrapper, err := model.CreateWrappedEmail(message, messageId, workerName)
	identifier := emailWrapper.GetMessageIdentifiers()

	if err != nil {
		emailUseCase.logger.Errorf("%v Cannot decode JSON", identifier)
		return returnError(
			emailUseCase.sendErrorEmailByBody(fmt.Sprintf("Error: Cannot decode JSON\nError details: %v", err)),
			customError.IncorrectMessage{Service: "Email-Service", Reason: fmt.Sprintf("Cannot decode JSON. err: %v", err), Identifier: identifier})
	}

	unwrappedMail := emailWrapper.Message
	if len(unwrappedMail.Body) < 1 {
		emailUseCase.logger.Errorf("%v Empty body", identifier)
		return returnError(
			emailUseCase.sendErrorEmailByEmailAndErrorType(unwrappedMail, "Empty body"),
			customError.IncorrectMessage{Service: "Email-Service", Reason: "Empty body", Identifier: identifier})
	}

	unwrappedMail.Destination = removeInvalidEmailReceivers(unwrappedMail.Destination, emailUseCase.logger, identifier)
	unwrappedMail.CCDestination = removeInvalidEmailReceivers(unwrappedMail.CCDestination, emailUseCase.logger, identifier)
	unwrappedMail.BCCDestination = removeInvalidEmailReceivers(unwrappedMail.BCCDestination, emailUseCase.logger, identifier)

	if len(unwrappedMail.Destination)+len(unwrappedMail.CCDestination)+len(unwrappedMail.BCCDestination) < 1 {
		emailUseCase.logger.Errorf("%v Empty receivers", identifier)
		return returnError(
			emailUseCase.sendErrorEmailByEmailAndErrorType(unwrappedMail, "Empty receivers"),
			customError.IncorrectMessage{Service: "Email-Service", Reason: "Empty receivers", Identifier: identifier})
	}

	/*check json*/
	if len(unwrappedMail.Attachments) > 0 {
		emailUseCase.logger.Infof("Message with attachments ! Processing...")

		for fileName, attachment := range unwrappedMail.Attachments {
			attachmentData, _ := json.Marshal(attachment)
			emailUseCase.logger.Debugf("attachment %v, data: %v", fileName, string(attachmentData[:]))

			file, _, err := emailUseCase.conversion.Convert(attachmentData)
			if err != nil {
				emailUseCase.logger.Errorf("Cannot generate file ! Service error %v", err)
				if conversion.IsConnectionError(err) {
					return customError.NotReadyMessage{Service: "Email-Service", Reason: fmt.Sprintf("Cannot generate file. err: %v", err), Identifier: identifier}
				} else {
					return returnError(
						emailUseCase.sendErrorEmailByEmailAndErrorTypeAndError(unwrappedMail, "Cannot generate file", err),
						err)
				}
			} else {
				emailWrapper.FileAttachment[fileName] = file
			}
		}
	} else {
		emailUseCase.logger.Infof("Message with no attachments")
	}

	if err := emailUseCase.smtp.Send(emailWrapper); err != nil {
		emailUseCase.logger.Errorf("%v Cannot send email. err: %v", identifier, err)
		if isNotReadyMessageError(err) {
			return err
		} else {
			return returnError(
				emailUseCase.sendErrorEmailByEmailAndErrorTypeAndError(unwrappedMail, "Cannot send email", err),
				err)
		}
	}

	return nil
}

func isNotReadyMessageError(error error) bool {
	switch error.(type) {
	case customError.NotReadyMessage:
		return true
	default:
		return false
	}
}

func returnError(firstError error, secondError error) error {
	if firstError != nil {
		return firstError
	} else {
		return secondError
	}
}

func (e EmailUseCase) sendErrorEmailByEmailAndErrorTypeAndError(email *rabbitMqDomain.Email, errorType string, error error) error {
	return e.sendErrorEmailByBody(fmt.Sprintf("To: %v\nCC: %v\nBCC: %v\nTitle: %v\n:Body: %v\nAttachments: %v\nError: %v\nError details: %v\nEnvironment: %v",
		email.Destination,
		email.CCDestination,
		email.BCCDestination,
		email.Title,
		email.Body,
		convertAttachments(email),
		errorType,
		error,
		e.config.ErrorEmailEnvironment))
}

func (e EmailUseCase) sendErrorEmailByEmailAndErrorType(email *rabbitMqDomain.Email, errorType string) error {
	return e.sendErrorEmailByBody(fmt.Sprintf("To: %v\nCC: %v\nBCC: %v\nTitle: %v\n:Body: %v\nAttachments: %v\nError: %v",
		email.Destination,
		email.CCDestination,
		email.BCCDestination,
		email.Title,
		email.Body,
		convertAttachments(email),
		errorType))
}

func convertAttachments(email *rabbitMqDomain.Email) map[string]string {
	attachments := make(map[string]string)
	for fileName, attachment := range email.Attachments {
		data, err := json.Marshal(attachment)
		if err == nil {
			attachments[fileName] = string(data)
		} else {
			attachments[fileName] = fmt.Sprintf("Error: %v", err)
		}
	}
	return attachments
}

func (e EmailUseCase) sendErrorEmailByBody(body string) error {
	if len(e.config.ErrorEmailTo) == 0 {
		e.logger.Warnf("Cannot send error mail. ErrorEmailTo is not set")
		return nil
	}

	return e.smtp.Send(&model.EmailWrapper{Message: &rabbitMqDomain.Email{
		Destination: e.config.ErrorEmailTo,
		Title:       "Email notification service error",
		Body:        body,
	}})
}

func removeInvalidEmailReceivers(slice []string, logger *zap.SugaredLogger, identifier string) []string {
	for i, rcount, size := 0, 0, len(slice); i < size; i++ {
		j := i - rcount

		if validateEmail(slice[j]) == false {
			logger.Warnf("%v Invalid email address ! Skipping receiver: %v", identifier, slice[j])
			slice = append(slice[:j], slice[j+1:]...)
			rcount++
		}
	}
	return slice
}

type InvoiceData struct {
	Header  Header
	Payload Payload
}

type Header struct {
	Version  string
	Provider string
	Content  Content
}

type Content struct {
	Type string
	Catg string
}

type Payload struct {
	InvDtls InvDtls
}

type InvDtls struct {
	_Id      string
	Num      string
	IssueDt  string
	SenderDt string
	Type     string
}
