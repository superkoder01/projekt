import { FunctionalUser } from './../components/forms/create-func-user/FunctionalUser';
import BaseApi from "@/api/base.api";
import {ApiTypeEnum} from "@/services/logger/api-type.enum";
import {LocalSpinner} from "@/services/model/localSpinner";
import {RequestResponseApi} from "@/models/request-response-api";
import {DataHolder} from "@/models/data-holder";
import {PagingModel} from "@/services/model/paging.model";
import {Provider} from "@/models/provider";
import {Pricing, Repurchase} from "@/models/billing/billing";

export class PricingApi extends BaseApi {

  protected getApiType(): ApiTypeEnum {
    return ApiTypeEnum.PRICING_API;
  }

  private readonly API_URL = '/api/core';
  private readonly API_PRICING_URL = this.API_URL + '/pricing';

  getPricing(lockScreen: boolean, localSpinner: LocalSpinner | null, pagination?: PagingModel): Promise<RequestResponseApi<DataHolder<Pricing>>> {
    return this.getDataFromUrl<DataHolder<Pricing>>(this.API_PRICING_URL, lockScreen, localSpinner, false, pagination);
  }

  saveNewPricing(newPricing: Pricing | Repurchase, successCallback: () => void, failCallback: () => void) {
    this.axiosCall({
      method: 'POST',
      url: this.API_PRICING_URL,
      data: newPricing
    }, true, null, true).then(value => {
      if (value.error === null) {
        if (successCallback !== undefined) {
          successCallback();
        }
      }
      else if(failCallback !== undefined){
        failCallback();
      }
    });
  }
}
