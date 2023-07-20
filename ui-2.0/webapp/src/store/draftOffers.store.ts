import {defineStore} from "pinia";
import {DataHolder} from "@/models/data-holder";
import {LocalSpinner} from "@/services/model/localSpinner";
import {PagingModel} from "@/services/model/paging.model";
import factoryApi from "@/api/factory.api";
import {OfferDraft} from "@/models/billing/billing";

export const useDraftOffersStore = defineStore({
  id: 'draftOffersStore',
  state: () => {
    return {
      draftOffers : Object(new DataHolder<OfferDraft>())
    };
  },
  actions: {
    fetchDraftOffers (lockScreen: boolean, localSpinner: LocalSpinner | null, pagination?: PagingModel) {
      factoryApi.draftOffersApi().getDraftOffers(lockScreen, localSpinner, pagination).then(
        (response) => {
          if (response.error === null) {
            if (response.data !== undefined)
              console.log(response.data);
              this.draftOffers = response.data;
          } else {
            //TODO: error handling
          }
        }
      );
    },
    createNewOfferDraft(newOfferDraft: OfferDraft, successCallback: () => void , failCallback: () => void){
       factoryApi.draftOffersApi().saveNewOfferDraft(newOfferDraft, successCallback, failCallback);
    }
  },
  getters: {
    getListDataHolder ():DataHolder<OfferDraft> {
      return this.draftOffers;
    },
    getOfferDraft(): Array<OfferDraft> {
      return this.draftOffers.elements;
    }
  }
});
