import { Installation } from './../components/forms/create-installation/Installation';
import factoryApi from '@/api/factory.api';
import {defineStore} from "pinia";
import {DataHolder} from "@/models/data-holder";
import { LocalSpinner } from "@/services/model/localSpinner";
import { PagingModel } from "@/services/model/paging.model";

export const useInstallationsStore = defineStore({
  id: 'installationsStore',
  state: () => {
    return {
      installations : Object(new DataHolder<Installation>())
    };
  },
  actions: {
    fetchInstallations(lockScreen: boolean, localSpinner: LocalSpinner | null, pagination?: PagingModel){
      factoryApi.installationsApi().getInstallations(lockScreen, localSpinner, pagination).then(
        (response) => {
          if (response.error === null) {
            if (response.data !== undefined)
              this.installations = response.data;
          } else {
            //TODO: error handling
          }
      });
    },
    fetchInstallationsById(installationId: number, lockScreen: boolean, localSpinner: LocalSpinner | null, pagination?: PagingModel){
      factoryApi.installationsApi().getInstallationById(installationId, lockScreen, localSpinner, pagination).then(
        (response) => {
          if (response.error === null) {
            if (response.data !== undefined)
              this.installations = response.data;
          } else {
            //TODO: error handling
          }
      });
    },
    fetchSelectedCustomerAccessPoints(customerId: number, lockScreen: boolean, localSpinner: LocalSpinner | null, pagination?: PagingModel) {
      factoryApi.installationsApi().getCustomerAccessPoints(customerId, lockScreen, localSpinner, pagination).then(
        (response) => {

          if (response.error === null) {
            if (response.data !== undefined)
              console.log(response.data);

            this.installations = response.data;
          } else {
            //TODO: error handling
          }
        }
      );
    },
    createNewInstallation(installation: Installation, successCallback: () => void , failCallback: () => void){
      factoryApi.installationsApi().saveNewInstallation(installation, successCallback, failCallback);
    },
  },
  getters: {
    getListDataHolder ():DataHolder<Installation> {
      return this.installations;
    }
  }
});
