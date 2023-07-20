package ports

//
//import (
//	mariadbdao "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/internal/repositories/billingrepo/mariadb/dao"
//	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/c4e-iges.git/pkg/billing/deprecated/mediation"
//	"time"
//)
//
//type IBillingRepo interface {
//	SaveInvoice(invoice mariadbdao.Invoice) error
//	GetProductionHistory(accessPointId uint) []mariadbdao.EnergyHistory
//	SaveCurrentEnergy(history mariadbdao.EnergyHistory)
//	GetAccessPoint(meterNumber string) mariadbdao.ServiceAccessPoint
//	SaveEnergyExcess(production []mariadbdao.EnergyProduction)
//	GetVatRate(startDate, endDate time.Time) float32
//	GetExcessEnergyProduction(startDate time.Time, accessPointId uint) mediation.ResaleRecord
//}
//
//type BillingRepoFactory interface {
//	MakeRepo() IBillingRepo
//}
