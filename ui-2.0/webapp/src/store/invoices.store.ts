import { defineStore } from 'pinia';
import {DataHolder} from "@/models/data-holder";
import factoryApi from "@/api/factory.api";
import {LocalSpinner} from "@/services/model/localSpinner";
import {PagingModel} from "@/services/model/paging.model";
import {Invoice} from "@/models/invoice/invoice";

export const useInvoicesStore = defineStore({
  id: 'invoicesStore',
  state: () => {
    return {
      invoices : Object(new DataHolder<Invoice>())
    };
  },
  actions: {
    fetchInvoices (lockScreen: boolean, localSpinner: LocalSpinner | null, pagination?: PagingModel) {
      factoryApi.invoicesApi().getInvoices(lockScreen, localSpinner, pagination).then(
        (response) => {
          if (response.error === null) {
            if (response.data !== undefined)
              this.invoices = response.data;
          } else {
            //TODO: error handling
          }
        }
      );
    },
    fetchInvoicesForCustomer (customerId: number, lockScreen: boolean, localSpinner: LocalSpinner | null, pagination?: PagingModel) {
      factoryApi.invoicesApi().getInvoicesByCustomerId(customerId, lockScreen, localSpinner, pagination).then(
        (response) => {
          if (response.error === null) {
            if (response.data !== undefined)
              this.invoices = response.data;
          } else {
            //TODO: error handling
          }
        }
      );
    }
  },
  getters: {
    getListDataHolder ():DataHolder<Invoice> {
      return this.invoices;
    }
  }
});
