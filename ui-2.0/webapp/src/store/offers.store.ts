import { defineStore } from 'pinia';
import {DataHolder} from "@/models/data-holder";
import {LocalSpinner} from "@/services/model/localSpinner";
import {PagingModel} from "@/services/model/paging.model";
import factoryApi from "@/api/factory.api";
import { ResponseError } from "@/models/request-response-api";
import {Offer, Pricing} from "@/models/billing/billing";

export const useOffersStore = defineStore({
  id: 'offersStore',
  state: () => {
    return {
      offers : Object(new DataHolder<Offer>())
    };
  },
  actions: {
    setOffers (offers:  DataHolder<Offer>) {
      this.offers = offers;
    },
    fetchOffers(lockScreen: boolean, localSpinner: LocalSpinner | null, pagination: PagingModel) {
      factoryApi.offersApi().getOffers(lockScreen, localSpinner, pagination).then(response => {
        if (response.error === null) {
          console.log(response.data)
          this.offers = response.data;
        } else {
          //TODO: error handling
        }
      });
    },
    fetchOffersForSelectedCustomer(customerId: number, lockScreen: boolean, localSpinner: LocalSpinner | null, pagination: PagingModel) {
      factoryApi.offersApi().getOffersByCustomerId(customerId,lockScreen, localSpinner, pagination).then(response => {
        if (response.error === null) {
          this.offers = response.data;
        } else {
          //TODO: error handling
        }
      });
    },

    updateOfferById(offerId: string,offer: Offer, successCallback: () => void , failCallback: (error : ResponseError) => void) {
      factoryApi.offersApi().updateOfferById(offerId,offer, successCallback, failCallback);
    }
   },
  getters: {
    getListDataHolder ():DataHolder<Offer> {
      return this.offers;
    }
  }
});
