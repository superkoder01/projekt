/*
BMSFES. 
Copyright (C) 2022-2035 

This file is part of the BMSFES solution. 
BMSFES is free software: you can redistribute it and/or modify 
it under the terms of the GNU Affero General Public License as 
published by the Free Software Foundation, either version 3 of the 
License, or (at your option) any later version.
 
BMSFES solution is distributed in the hope that it will be useful, 
but WITHOUT ANY WARRANTY; without even the implied warranty of 
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the 
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License 
along with this program.  If not, see <http://www.gnu.org/licenses/>.
*/
package handler

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/op/go-logging"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mongo/model/offer"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/error"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/service/mongo"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/notification-service-rabbitmq.git/config"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/notification-service-rabbitmq.git/domain/model"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/notification-service-rabbitmq.git/messagebroker/rabbitmq"
)

var (
	logger = logging.MustGetLogger("pricing_handler")
	EMAIL  = "email"
)

type EmailHandler struct {
	config          *config.RabbitMQProducerConfig
	offerService    mongo.Service
	contractService mongo.Service
}

func NewEmailHandler(config *config.RabbitMQProducerConfig, contractService mongo.Service, offerService mongo.Service) *EmailHandler {
	return &EmailHandler{config: config, offerService: offerService, contractService: contractService}
}

func (eh *EmailHandler) SendEmailOffer(ctx *gin.Context) {
	id := ctx.Param("id")
	offerModel, err := eh.offerService.FindOne(id)

	if err != nil {
		e.HandleError(err, ctx)
	}
	offerBody, ok := offerModel.(*offer.Offer)
	if !ok {
		e.HandleError(e.ApiErrInvalidDataModel, ctx)
	}

	producer := rabbitmq.NewRabbitMQProducer(*eh.config, logger)
	err = producer.InitializeConnection()
	if err != nil {
		e.HandleError(err, ctx)
	}
	offerData := &offer.Offer{
		Header:  offerBody.Header,
		Payload: offerBody.Payload,
	}
	recipient := offerBody.Payload.CustomerDetails.Contact.Email
	offerBytes, err := json.Marshal(&offerData)
	offerString := string(offerBytes)
	buf := bytes.NewBuffer([]byte(offerString))
	decoder := json.NewDecoder(buf)
	err = decoder.Decode(&offerData)
	if err != nil {
		e.HandleError(err, ctx)
	}
	email := CreateOfferEmailMessage(recipient, offerString)
	bbb, err := Serialize_M(email)
	if err != nil {
		e.HandleError(err, ctx)
	}
	err = producer.PublishMessage(model.Message{
		Payload: bbb,
	}, EMAIL)
	//TODO: update
	//if err != nil {
	//	return
	//} else {
	//	var newStatus = &utils.UpdateStatus{Status: "sent"}
	//	_, err = eh.offerService.Update(id, newStatus)
	//	if err != nil {
	//		e.HandleError(err, ctx)
	//	}
	//}
	ctx.JSON(200, nil)
}

func DeserializeEmail(offerString string) *offer.Offer {
	var msg offer.Offer
	buf := bytes.NewBuffer([]byte(offerString))
	decoder := json.NewDecoder(buf)
	err := decoder.Decode(&msg)
	fmt.Println(err)
	return &msg
}

func CreateOfferEmailMessage(recipient string, offerString string) *model.Email {
	var i interface{}
	i = DeserializeEmail(offerString)
	return &model.Email{
		Destination: []string{recipient},
		Title:       "Oferta sprzedaży-zakupu energii elektrycznej",
		Body:        "Dzień dobry, w załączniku \"Oferta sprzedaży-zakupu energii elektrycznej\".",
		InvoiceData: &i,
	}
}
func Serialize_M(msg interface{}) ([]byte, error) {
	var b bytes.Buffer
	encoder := json.NewEncoder(&b)
	err := encoder.Encode(msg)
	return b.Bytes(), err
}
