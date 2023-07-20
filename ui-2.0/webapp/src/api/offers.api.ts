import BaseApi from "@/api/base.api";
import {ApiTypeEnum} from "@/services/logger/api-type.enum";
import {LocalSpinner} from "@/services/model/localSpinner";
import {PagingModel} from "@/services/model/paging.model";
import {DataHolder} from "@/models/data-holder";
import { RequestResponseApi, ResponseError } from "@/models/request-response-api";
import {Offer, Pricing} from "@/models/billing/billing";
import { Provider } from "@/models/provider";


export class OffersApi extends BaseApi {

  getApiType(): ApiTypeEnum {
    return ApiTypeEnum.OFFERS_API;
  }

  private readonly OFFERS_URL = '/api/core/offer';
  private readonly CUSTOMER_OFFERS_URL = this.OFFERS_URL+"/customer/";


  public async getOffers(lockScreen: boolean, localSpinner: LocalSpinner | null, pagination?: PagingModel): Promise<RequestResponseApi<DataHolder<Pricing>>> {
    return this.getDataFromUrl<DataHolder<Pricing>>(this.OFFERS_URL, lockScreen, localSpinner, false, pagination);
  }

  public async getOffersByCustomerId(customerId: number, lockScreen: boolean, localSpinner: LocalSpinner | null, pagination?: PagingModel) : Promise<RequestResponseApi<DataHolder<Pricing>>> {
    return this.getDataFromUrl<DataHolder<Pricing>>(this.CUSTOMER_OFFERS_URL+customerId, lockScreen, localSpinner, false, pagination);
  }

  public async getOfferById(offerId:string, lockScreen: boolean, localSpinner: LocalSpinner | null) : Promise<RequestResponseApi<Offer>> {
    return this.getDataFromUrl<Offer>(this.OFFERS_URL+"/"+offerId, lockScreen, localSpinner, false);
  }

  saveNewOffer(newOffer: Offer, successCallback: () => void, failCallback: () => void) {
    this.axiosCall({
      method: 'POST',
      url: this.OFFERS_URL,
      data: newOffer
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
  updateOfferById(offerId:string, offer: Offer, successCallback: () => void, failCallback: (error : ResponseError) => void) {
    const url = this.OFFERS_URL+"/"+offerId;
    this.axiosCall<Offer>({
      method: 'PUT',
      url: url,
      data: offer,
    }, true, null, true).then(value => {
      if (value.error === null) {
        if (successCallback !== undefined) {
          successCallback();
        }
      }
      else if(failCallback !== undefined){
        failCallback(value.error);
      }
    });
  }
}
