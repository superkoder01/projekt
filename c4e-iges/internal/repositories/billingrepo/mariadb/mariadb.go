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
package mariadb

//
//import (
//	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/config"
//	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/logger"
//	mariadbdao "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/repositories/billingrepo/mariadb/dao"
//	mariadbsession "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/repositories/billingrepo/mariadb/session"
//	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/pkg/billing/deprecated/mediation"
//	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/pkg/utils/datetime"
//	"gorm.io/gorm"
//	"time"
//)
//
//// Represents
//type mariadb struct {
//	client *gorm.DB
//	logger logger.Logger
//}
//
//func NewBillingRepo(cfg *config.AppConfig, logger logger.Logger) *mariadb {
//	session, err := mariadbsession.New(cfg)
//	if err != nil {
//		panic("failed to create connection to mariadb billingrepo")
//	}
//	return &mariadb{
//		client: session.GetConnection(),
//		logger: logger,
//	}
//}
//
//func (r *mariadb) SaveInvoice(invoice mariadbdao.Invoice) error {
//	r.client.Create(invoice)
//	return nil
//}
//
//func (r *mariadb) SaveEnergyExcess(production []mariadbdao.EnergyProduction) {
//	r.client.Create(production)
//}
//
//func (r *mariadb) GetVatRate(startDate, endDate time.Time) float32 {
//	var rate mariadbdao.VatRate
//	r.client.Select("RATE").Where("START_DATE < ? AND END_DATE > ? AND VAT_GROUP = ?", startDate, endDate, mariadbdao.Gxx).First(&rate)
//
//	return float32(rate.Value) / 100
//}
//
//func (r *mariadb) GetExcessEnergyProduction(startDate time.Time, accessPointId uint) mediation.ResaleRecord {
//	var production mariadbdao.EnergyProduction
//	r.logger.Debugf("energyExcess for date %v", startDate.Format(datetime.DBDateShortForm))
//	/*end date should be equal or -1 day*/
//	r.client.Where("SERVICE_ACCESS_POINT_ID = ? AND END_DATE = ?", accessPointId, startDate.Format(datetime.DBDateShortForm)).Find(&production)
//	return mediation.ResaleRecord{
//		From:              production.StartDate,
//		To:                production.EndDate,
//		Excess:            production.EnergyAmount,
//		UnitPrice:         production.NetPrice,
//		PreviousMeterRead: -1, //empty values
//		CurrentMeterRead:  -1,
//	}
//}
//
//func (r *mariadb) GetProductionHistory(accessPointId uint) []mariadbdao.EnergyHistory {
//	var history []mariadbdao.EnergyHistory
//	r.client.Where("SERVICE_ACCESS_POINT_ID = ?", accessPointId).Order("PERIOD ASC").Find(&history)
//
//	return history
//}
//
//func (r *mariadb) SaveCurrentEnergy(history mariadbdao.EnergyHistory) {
//	r.logger.Debugf("Save CurrentEnergy: %v", history)
//	r.client.Create(history)
//}
//
//func (r *mariadb) GetAccessPoint(meterNumber string) mariadbdao.ServiceAccessPoint {
//	var accessPoint mariadbdao.ServiceAccessPoint
//	r.client.Where("METER_NUMBER = ?", meterNumber).First(&accessPoint)
//	return accessPoint
//}
