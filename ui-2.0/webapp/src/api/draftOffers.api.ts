import BaseApi from "@/api/base.api";
import {ApiTypeEnum} from "@/services/logger/api-type.enum";
import {LocalSpinner} from "@/services/model/localSpinner";
import {PagingModel} from "@/services/model/paging.model";
import {DataHolder} from "@/models/data-holder";
import {RequestResponseApi} from "@/models/request-response-api";
import {Provider} from "@/models/provider";
import {Offer, OfferDraft} from "@/models/billing/billing";


export class DraftOffersApi extends BaseApi {

  getApiType(): ApiTypeEnum {
    return ApiTypeEnum.DRAFT_OFFERS_API;
  }

  private readonly DRAFT_OFFERS_URL = '/api/core/draft_offer';

  public async getDraftOffers(lockScreen: boolean, localSpinner: LocalSpinner | null, pagination?: PagingModel): Promise<RequestResponseApi<DataHolder<OfferDraft>>> {
    return this.getDataFromUrl<DataHolder<OfferDraft>>(this.DRAFT_OFFERS_URL, lockScreen, localSpinner, false, pagination);
  }

  public async getOfferDraftById(offerDraftId:string, lockScreen: boolean, localSpinner: LocalSpinner | null) : Promise<RequestResponseApi<OfferDraft>> {
    return this.getDataFromUrl<OfferDraft>(this.DRAFT_OFFERS_URL+"/"+offerDraftId, lockScreen, localSpinner, false);
  }

  public async saveNewOfferDraft(offerDraft: OfferDraft, successCallback: () => void, failCallback: () => void) {
    this.axiosCall<OfferDraft>({
      method: 'POST',
      url: this.DRAFT_OFFERS_URL,
      data: offerDraft,
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
