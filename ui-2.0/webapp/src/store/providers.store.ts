import { FunctionalUser } from './../components/forms/create-func-user/FunctionalUser';
import { Provider } from '@/models/provider';
import { defineStore } from 'pinia';
import {DataHolder} from "@/models/data-holder";
import factoryApi from "@/api/factory.api";
import {LocalSpinner} from "@/services/model/localSpinner";
import {PagingModel} from "@/services/model/paging.model";
import {ResponseError} from "@/models/request-response-api";

export const useProvidersStore = defineStore({
  id: 'providersStore',
  state: () => {
    return {
      providers : Object(new DataHolder<Provider>()),
      selectedProviderData: Object(Provider),
      providerAdmins : Object(new DataHolder<Provider>()),
    };
  },
  actions: {
    setProviders (providers: DataHolder<Provider>) {
      this.providers = providers;
      console.log(providers);
    },
    setSelectedProviderData(provider: Provider) {
      this.selectedProviderData = provider;
    },
    setSelectedProviderAdmins(providerAdmins: DataHolder<Provider>) {
      this.providerAdmins = providerAdmins;
    },
    fetchProviders(lockScreen: boolean, localSpinner: LocalSpinner | null, pagination: PagingModel){
      factoryApi.providersApi().getProviders(lockScreen, localSpinner, pagination).then(
        (response)=>{
          if (response.error === null) {
            this.providers = response.data;
          } else {
            //TODO: error handling
          }
        }
      );
    },
    fetchProviderById(providerId:number, lockScreen: boolean, localSpinner: LocalSpinner | null){
      factoryApi.providersApi().getProviderById(providerId, lockScreen, localSpinner).then(
        (response)=>{
          if (response.error === null) {
            this.selectedProviderData = response.data;
          } else {
            //TODO: error handling
          }
        }
      );
    },
    fetchProviderAdministrators(providerId:number,lockScreen: boolean, localSpinner: LocalSpinner | null, pagination: PagingModel){
      factoryApi.providersApi().getProviderAdmins(providerId, lockScreen, localSpinner, pagination).then(
        (response)=>{
          if (response.error === null) {
            this.providerAdmins = response.data;
          } else {
            //TODO: error handling
          }
        }
      );
    },
    createNewProvider(provider: Provider, successCallback: () => void , failCallback: (error : ResponseError) => void){
      factoryApi.providersApi().saveNewProvider(provider, successCallback, failCallback);
    },
    createNewProviderWithAdmin(provider: Provider, admin: FunctionalUser, successCallback: () => void , failCallback: (error : ResponseError) => void){
      factoryApi.providersApi().saveNewProviderWithAdmin(provider, admin, successCallback, failCallback);
    },
    updateProvider(provider: Provider, successCallback: () => void , failCallback: (error : ResponseError) => void) {
      factoryApi.providersApi().updateProvider(provider, successCallback, failCallback);
    }
  },

  getters: {
    getListDataHolder (): DataHolder<Provider> {
      return this.providers;
    },
    getProviderAdmins(): DataHolder<Provider> {
      return this.providerAdmins;
    }
  }
});
