package utils

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mongo/model/invoice"
	role "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mongo/model/utils/enum"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/error"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/service/management_and_login"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/commons.git/billing"
	apiUtils "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/management-and-login.git/pkg/api_utils"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"strconv"
	"time"
)

type SortObject struct {
	Name  string
	Order int
}

type countResult struct {
	Id    string
	Total int64
}

const (
	// Collection names
	CONTRACTS          = "contracts"
	OFFER              = "offers"
	OFFER_DRAFTS       = "offer_drafts"
	INVOICES           = "invoices"
	PRICINGS           = "pricings"
	TARIFF_GROUP_LABEL = "labels"

	// Property paths
	PROVIDER_NAME_PATH         = "header.provider"
	CONTRACTS_CUSTOMER_ID_PATH = "payload.contractDetails.customerId"
	OFFER_CUSTOMER_ID_PATH     = "payload.offerDetails.customerId"
	INVOICE_CUSTOMER_ID_PATH   = "payload.invoiceDetails.customerId"

	// Invoice types
	INVOICE            = "invoice"
	REPURCHASE         = "repurchase"
	REPURCHASE_DETAILS = "repurchase-details"
)

func SetFindOptions(query *Query) *options.FindOptions {
	options := new(options.FindOptions)
	if query.Limit != 0 {
		options.SetSkip(int64(query.Offset))
		options.SetLimit(int64(query.Limit))
	}
	if query.Sort != nil {
		options.SetSort(&SortObject{query.Sort.Value, query.Sort.Order})
	}
	return options
}

func SetAggregateOptions(query *Query, p *mongo.Pipeline) {
	options := SetFindOptions(query)
	if options.Sort != nil {
		if obj, ok := options.Sort.(*SortObject); ok {
			sortStage := bson.D{{"$sort", bson.D{{obj.Name, obj.Order}}}}
			*p = append(*p, sortStage)
		}
	}
	if options.Skip != nil {
		skipStage := bson.D{{"$skip", options.Skip}}
		*p = append(*p, skipStage)
	}
	if options.Limit != nil {
		limitStage := bson.D{{"$limit", options.Limit}}
		*p = append(*p, limitStage)
	}
	return
}

func SetCustomerIdPipeline(ctx *gin.Context, query *Query, p *mongo.Pipeline, collection string) {
	if len(query.FilterFields) != len(query.FilterValues) {
		e.HandleError(e.FilteredFieldsAndFilteredValuesLengthsDoNotMatch, ctx)
	}
	name := management_and_login.GetProvider(ctx).Name
	andFilters := []bson.D{{{PROVIDER_NAME_PATH, name}}}

	tokenRole, err := apiUtils.GetTokenRole(ctx)
	if err != nil {
		e.HandleError(err, ctx)
	}

	switch tokenRole {
	case role.SUPER_AGENT.Name(), role.AGENT.Name(), role.TRADER.Name():
		sId := ctx.Param("id")
		if sId == "" {
			var filters []bson.D
			accounts := management_and_login.GetWorkersCustomerAccounts(ctx)
			for _, acc := range accounts.Elements {
				filters = append(filters, bson.D{{getCustomerIdPath(collection), strconv.Itoa(acc.ID)}})
			}
			andFilters = append(andFilters, bson.D{{"$or", filters}})
		} else {
			andFilters = append(andFilters, bson.D{{getCustomerIdPath(collection), sId}})
		}
	case role.ADMINISTRATOR_BASIC.Name(), role.ADMINISTRATOR_FULL.Name(), role.SUPER_ADMIN.Name():
		sId := ctx.Param("id")
		if sId != "" {
			andFilters = append(andFilters, bson.D{{getCustomerIdPath(collection), sId}})
		} else {
			break
		}
	case role.PROSUMER.Name():
		id, err := apiUtils.GetTokenCustomerAccountID(ctx)
		if err != nil {
			e.HandleError(err, ctx)
		}
		if id != 0 {
			andFilters = append(andFilters, bson.D{{getCustomerIdPath(collection), strconv.Itoa(id)}})
		}
	}

	for i := 0; i < len(query.FilterFields); i++ {
		andFilters = append(andFilters, bson.D{{query.FilterFields[i], query.FilterValues[i]}})
	}

	*p = append(*p, bson.D{{"$match", bson.D{
		{"$and", andFilters},
	}}})
	return
}

func SetProviderPipeline(ctx *gin.Context, query *Query, p *mongo.Pipeline) {
	name := management_and_login.GetProvider(ctx).Name
	andFilters := []bson.D{{{PROVIDER_NAME_PATH, name}}}
	if len(query.FilterFields) != len(query.FilterValues) {
		e.HandleError(e.FilteredFieldsAndFilteredValuesLengthsDoNotMatch, ctx)
	}
	for i := 0; i < len(query.FilterFields); i++ {
		andFilters = append(andFilters, bson.D{{query.FilterFields[i], query.FilterValues[i]}})
	}
	*p = append(*p, bson.D{{"$match", bson.D{
		{"$and", andFilters},
	}}})
	return
}

func SetProviderFilters(ctx *gin.Context) interface{} {
	name := management_and_login.GetProvider(ctx).Name
	andFilters := []bson.D{{{PROVIDER_NAME_PATH, name}}}
	res := bson.D{
		{"$and", andFilters},
	}
	return res
}

func SetCount(p *mongo.Pipeline) {
	*p = append(*p, bson.D{{"$group", bson.D{{"_id", "$Id"}, {"total", bson.D{{"$sum", 1}}}}}})
	//*p = append(*p, bson.D{{"$group", bson.M{"_id": "$_id.Id", "count": bson.M{"$sum": 1}}}})
	//*p = append(*p, bson.D{{"$count", "amount"}})

	return
}

func getCustomerIdPath(collection string) string {
	switch collection {
	case CONTRACTS:
		return CONTRACTS_CUSTOMER_ID_PATH
	case OFFER:
		return OFFER_CUSTOMER_ID_PATH
	case INVOICES:
		return INVOICE_CUSTOMER_ID_PATH
	default:
		return ""
	}
}

func GenerateDocumentNumber(provider string, collection string, count int64) string {
	dt := time.Now()
	return fmt.Sprintf("%s/%s/%s/%d", generateDocumentName(provider, collection, count+1), AppendLeadingZero(dt.Day()), AppendLeadingZero(int(dt.Month())), dt.Year())
}

func GenerateCreationDate() string {
	dt := time.Now()
	return fmt.Sprintf("%d/%s/%s %s:%s:%s", dt.Year(), AppendLeadingZero(int(dt.Month())), AppendLeadingZero(dt.Day()), AppendLeadingZero(dt.Hour()), AppendLeadingZero(dt.Minute()), AppendLeadingZero(dt.Second()))
}

func generateDocumentName(provider string, collection string, count int64) string {
	res := ""
	// TODO: uzyj nazwy skróconej providera
	if provider == "Keno Energia Sp. z o.o." {
		res += "KE."
	}
	res += appendCollectionInfo(collection) + "." + strconv.FormatInt(count, 10)
	return res
}

func appendCollectionInfo(collection string) string {
	switch collection {
	case CONTRACTS:
		return "C"
	case OFFER:
		return "O"
	default:
		return ""
	}
}

func SetInvoicePayloadModel(elem *invoice.Invoice) error {
	switch elem.Header.Content.Type {
	case INVOICE:
		payload := &billing.ProsumentPayload{}
		res, err := unmarshalBSON(payload, &elem.Payload)
		if err != nil {
			log.Fatal(err)
			return err
		}
		elem.Payload = res
	case REPURCHASE:
		payload := &billing.ProsumentRepurchasePayload{}
		res, err := unmarshalBSON(payload, &elem.Payload)
		if err != nil {
			log.Fatal(err)
			return err
		}
		elem.Payload = res
	case REPURCHASE_DETAILS:
		payload := &billing.ProsumentRepurchaseDetailsPayload{}
		res, err := unmarshalBSON(payload, &elem.Payload)
		if err != nil {
			log.Fatal(err)
			return err
		}
		elem.Payload = res
	}
	return nil
}

func unmarshalBSON(payload interface{}, data *interface{}) (interface{}, error) {
	byteArr, err := bson.Marshal(data)
	if err != nil {
		return nil, err
	}
	err = bson.Unmarshal(byteArr, payload)
	if err != nil {
		return nil, err
	}
	return payload, nil
}

func GetCountElementsResponse(ctx *gin.Context, count *mongo.Cursor) (int64, error) {
	res := &countResult{}
	for count.Next(ctx) {
		err := count.Decode(&res)
		if err != nil {
			return 0, err
		} else {
			return res.Total, nil
		}
	}
	return 0, nil
}
