import { defineStore } from 'pinia';
import {DataHolder} from "@/models/data-holder";
import factoryApi from "@/api/factory.api";
import {LocalSpinner} from "@/services/model/localSpinner";
import {PagingModel} from "@/services/model/paging.model";
import {ContractsHolder} from "@/models/contracts-holder";
import { Contract, Offer } from "@/models/billing/billing";
import { ResponseError } from "@/models/request-response-api";

export const useContractsStore = defineStore({
  id: 'contractsStore',
  state: () => {
    return {
      contracts: Object(new DataHolder<Contract>()),
      contract: Object(ContractsHolder)
    };
  },
  actions: {
    fetchContracts (lockScreen: boolean, localSpinner: LocalSpinner | null, pagination?: PagingModel) {
      factoryApi.contractsApi().getContracts(lockScreen, localSpinner, pagination).then(
        (response) => {
          console.log(response.data)
          if (response.error === null) {
            if (response.data !== undefined)
              this.contracts = response.data;
          } else {
            //TODO: error handling
          }
        }
      );
    },
    fetchContractsForCustomer (customerId: number, lockScreen: boolean, localSpinner: LocalSpinner | null, pagination?: PagingModel) {
      factoryApi.contractsApi().getContractsByCustomerId(customerId, lockScreen, localSpinner, pagination).then(
        (response) => {
          if (response.error === null) {
            if (response.data !== undefined)
              this.contracts = response.data;
          } else {
            //TODO: error handling
          }
        }
      );
    },
    updateContractById(contractId: string, contract: Contract, successCallback: () => void , failCallback: (error : ResponseError) => void) {
      factoryApi.contractsApi().updateContractById(contractId, contract, successCallback, failCallback);
    }
  },
  getters: {
    getListDataHolder ():DataHolder<Contract> {
      return this.contracts;
    }
  }
});
