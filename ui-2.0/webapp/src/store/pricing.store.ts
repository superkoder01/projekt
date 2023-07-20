import {defineStore} from "pinia";
import {DataHolder} from "@/models/data-holder";
import {LocalSpinner} from "@/services/model/localSpinner";
import {PagingModel} from "@/services/model/paging.model";
import factoryApi from "@/api/factory.api";
import {Pricing, PricingHolder} from "@/models/billing/billing";
import {ServiceTypeEnum} from "@/models/billing/enum/service-type.enum";

export const usePricingStore = defineStore({
  id: 'pricingStore',
  state: () => {
    return {
      pricing: Object(new DataHolder<Pricing>())
    };
  },
  actions: {
    setPricing(pricing: DataHolder<Pricing>) {
      this.pricing = pricing;
    },
    fetchPricing(lockScreen: boolean, localSpinner: LocalSpinner | null, pagination?: PagingModel) {
      factoryApi.pricingApi().getPricing(lockScreen, localSpinner, pagination).then(
        (response) => {
          console.log(response.data);
          if (response.error === null) {
            if (response.data !== undefined)
              this.pricing = response.data;
          } else {
            //TODO: error handling
          }
        }
      );
    },
  },
  getters: {
    getListDataHolder(): DataHolder<Pricing> {
      return this.pricing;
    },
    getPricingSALE(): Array<Pricing> | null {
      if (this.pricing.elements!= undefined) {
        return this.pricing.elements
          .filter((obj: PricingHolder) => {
            return obj.payload.type === ServiceTypeEnum.SALE;
          })
          .map((obj: PricingHolder)=>{
            if(obj.payload.id.length === 0){
              obj.payload.id = obj.id;
            }
            return obj.payload;
        });
      }
      return null;
    },
    getPricingREPURCHASE(): Array<Pricing> | null {
      if (this.pricing.elements!= undefined) {
        return this.pricing.elements
          .filter((obj: PricingHolder) => {
            return obj.payload.type === ServiceTypeEnum.REPURCHASE;
          })
          .map((obj: PricingHolder)=>{
            if(obj.payload.id.length === 0){
              obj.payload.id = obj.id;
            }
            return obj.payload;
          });
      }
      return null;
    },
    getPricingREPURCHASE_RDN(): Array<Pricing> | null {
      if (this.pricing.elements!= undefined) {
        return this.pricing.elements
          .filter((obj: PricingHolder) => {
            return obj.payload.type === ServiceTypeEnum.REPURCHASE_RDN;
          })
          .map((obj: PricingHolder)=>{
            if(obj.payload.id.length === 0){
              obj.payload.id = obj.id;
            }
            return obj.payload;
          });
      }
      return null;
    }

  }
});
