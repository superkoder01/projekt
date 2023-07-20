package impl

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mariadb/model"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mariadb/model/fee"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mariadb/model/tariff_group"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/api/mariadb/model/tariff_group_parameter"
	e "gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/error"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/internal/utils"
	"gitlab.sce-ovoo.pl/ovoo/products/chain4energy/billing/billing-management.git/service"
	"io/ioutil"
	"strconv"
)

type tariffGroupOsdHandler struct {
	tariffGroupService          service.Service
	tariffGroupParameterService service.Service
}

func NewTariffGroupOsdHandler(tg service.Service, tgp service.Service) *tariffGroupOsdHandler {
	return &tariffGroupOsdHandler{tariffGroupService: tg, tariffGroupParameterService: tgp}
}

func (c *tariffGroupOsdHandler) List(ctx *gin.Context) {
	models, err := c.tariffGroupService.List()
	mdls := []model.Model{}

	if err != nil {
		e.HandleError(err, ctx)
	}

	for _, model := range models {
		tariffGroup, ok := model.(*tariff_group.TariffGroup)
		if !ok {
			e.HandleError(e.ErrInternalServerError, ctx)
		}
		mdls = append(mdls, c.getFees(tariffGroup, ctx))
	}
	if err != nil {
		e.HandleError(err, ctx)
	}

	ctx.JSON(200, utils.WrapQueryResult(int64(len(mdls)), mdls))
}

func (c *tariffGroupOsdHandler) Create(ctx *gin.Context) {
	newTariffGroupOsd, err := ioutil.ReadAll(ctx.Request.Body)
	if err != nil {
		e.HandleError(err, ctx)
	}
	var newTariffGroup tariff_group.TariffGroup
	err = json.Unmarshal(newTariffGroupOsd, &newTariffGroup)
	if err != nil {
		e.HandleError(err, ctx)
	}
	//TODO: dorobić transakcje
	var createTariffGroup model.Model
	createTariffGroup, err = c.tariffGroupService.Create(&newTariffGroup)
	if err != nil {
		e.HandleError(err, ctx)
	}

	tariffGroup, ok := createTariffGroup.(*tariff_group.TariffGroup)
	if !ok {
		e.HandleError(e.ErrInternalServerError, ctx)
	}

	var newFees fee.Fees
	err = json.Unmarshal(newTariffGroupOsd, &newFees)
	if err != nil {
		e.HandleError(err, ctx)
	}

	for _, fee := range newFees.Fees {
		newTariffGroupParameter := &tariff_group_parameter.TariffGroupParameter{
			TariffGroupID:   tariffGroup.ID,
			ParameterNameID: fee.NameID,
			Price:           fee.Price,
		}

		_, err = c.tariffGroupParameterService.Create(newTariffGroupParameter)
		if err != nil {
			e.HandleError(err, ctx)
		}
	}

	ctx.JSON(200, createTariffGroup)
}

func (c *tariffGroupOsdHandler) GetWithFilter(ctx *gin.Context) {
}

func (c *tariffGroupOsdHandler) GetByID(ctx *gin.Context) {
	sId := ctx.Param("id")
	id, err := strconv.Atoi(sId)
	if err != nil {
		e.HandleError(e.ErrInternalServerError, ctx)
	}

	model, err := c.tariffGroupService.GetByID(id)
	if err != nil {
		e.HandleError(err, ctx)
	}

	tariffGroup, ok := model.(*tariff_group.TariffGroup)
	if !ok {
		e.HandleError(e.ErrInternalServerError, ctx)
	}

	ctx.JSON(200, c.getFees(tariffGroup, ctx))

}

func (c *tariffGroupOsdHandler) UpdateByID(ctx *gin.Context) {
}

func (c *tariffGroupOsdHandler) DeleteByID(ctx *gin.Context) {
}

func (c *tariffGroupOsdHandler) Get(ctx *gin.Context) {
}

func (c *tariffGroupOsdHandler) Update(ctx *gin.Context) {
}

func (c *tariffGroupOsdHandler) Delete(ctx *gin.Context) {
}

func (c *tariffGroupOsdHandler) getFees(tg *tariff_group.TariffGroup, ctx *gin.Context) *tariff_group.GetTariffGroup {
	tgParameters, err := c.tariffGroupParameterService.GetWithFilter(tariff_group_parameter.TariffGroupParameter{TariffGroupID: tg.ID})
	if err != nil {
		e.HandleError(err, ctx)
	}
	var fees []fee.Fee
	for _, parameter := range tgParameters {
		tariffGroupParameter, ok := parameter.(*tariff_group_parameter.TariffGroupParameter)
		if !ok {
			e.HandleError(e.ErrInternalServerError, ctx)
		}
		mdFee := &fee.Fee{NameID: tariffGroupParameter.ParameterNameID, Price: tariffGroupParameter.Price}
		fees = append(fees, *mdFee)
	}
	return &tariff_group.GetTariffGroup{TariffGroup: *tg, Fees: fees}
}
