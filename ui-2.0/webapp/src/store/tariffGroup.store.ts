import factoryApi from '@/api/factory.api';
import { TariffGroup } from '@/components/forms/tariff-group/TariffGroup';
import { DataHolder } from '@/models/data-holder';
import { ResponseError } from '@/models/request-response-api';
import { LocalSpinner } from '@/services/model/localSpinner';
import { PagingModel } from '@/services/model/paging.model';
import { defineStore } from "pinia";

export const useTariffGroupStore = defineStore({
  id: 'tarrifGroupStore',
  state: () => {
    return{
      distributionNetworkOperator: Object(new DataHolder<TariffGroup>()),
      fees: Object(new DataHolder<any>())
    };      
  },
  actions: {
    setDistributionNetworkOperator(operator: DataHolder<TariffGroup>){
      this.distributionNetworkOperator = operator;
    },
    setFees(fees: any) {
      this.fees = fees;
    },
    fetchDistributionNetworkOperator(lockScreen: boolean, localSpinner: LocalSpinner | null) {
      factoryApi.tarrifGroupApi().getDistributionNetworkOperator(lockScreen, localSpinner).then(
        (response)=>{
          if(response.error === null) {
            this.distributionNetworkOperator = response.data;
          }
          else {
            // TODO: error handling
          }
        }
      );
    },
    fetchFees(lockScreen: boolean, localSpinner: LocalSpinner | null) {
      factoryApi.tarrifGroupApi().getFees(lockScreen, localSpinner).then(
        (response)=>{
          if(response.error === null) {
            this.fees = response.data;
          }
          else {
            // TODO: error handling
          }
        }
      );
    },
    createNewTariffGroup(tariffGroup: TariffGroup, successCallback: () => void , failCallback: (error : ResponseError) => void){
      factoryApi.tarrifGroupApi().saveNewTarrifGroup(tariffGroup, successCallback, failCallback);
    }
  },
  getters: {
    getDistributionNetworkOperator (): DataHolder<TariffGroup> {
      return this.distributionNetworkOperator;
    },
    getFees(): DataHolder<any> {
      return this.fees;
    }
  }
});