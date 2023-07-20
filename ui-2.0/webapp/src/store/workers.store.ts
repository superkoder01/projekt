import { defineStore } from 'pinia';
import {DataHolder} from "@/models/data-holder";
import factoryApi from "@/api/factory.api";
import { Emplayee } from "@/models/emplayee";
import {LocalSpinner} from "@/services/model/localSpinner";
import { PagingModel } from '@/services/model/paging.model';

export const useWorkersStore = defineStore({
  id: 'workersStore',
  state: () => {
    return {
      workers : Object(new DataHolder<Emplayee>()),
      currentWorkerData: {} as Emplayee,
      selectedWorkerData: {} as Emplayee
    };
  },
  actions: {
    setWorkers (workers: DataHolder<Emplayee>) {
      this.workers = workers;
    },
    deleteWorker (workerId: number, successCallback: () => void , failCallback: () => void) {
      factoryApi.workersApi().deleteWorker(workerId, successCallback, failCallback);
    },
    createWorker (worker: Emplayee, successCallback: () => void , failCallback: () => void) {
      factoryApi.workersApi().saveNewWorker(worker, successCallback, failCallback);
    },
    updateWorkerById (workerId: number, worker: Emplayee,successCallback: () => void , failCallback: () => void) {
      factoryApi.workersApi().updateWorkerById(workerId, worker, successCallback, failCallback);
    },
    fetchWorkersData(lockScreen: boolean, localSpinner: LocalSpinner | null, pagination: PagingModel){
      factoryApi.workersApi().getWorkers(lockScreen, localSpinner, pagination).then(
        (response) => {
          if (response.error === null){
            this.workers = response.data;
          }
          else {
            //TODO: error handling
          }
        });
    },
    fetchCurrentWorkerData(lockScreen: boolean, localSpinner: LocalSpinner | null){
      factoryApi.workersApi().getCurrentWorkerData(lockScreen, localSpinner).then(

        (response)=>{
          console.log("Then " + JSON.stringify(response));
          if(response.error === null) {
            console.log("Then");
            if(response.data !== undefined) {
              this.currentWorkerData = response.data;
              console.log("Then");
            }
          } else {
            //TODO: error handling
          }
        }
      );
    },
    fetchSelectedWorkerData(workerId: number, lockScreen: boolean, localSpinner: LocalSpinner | null) {
      factoryApi.workersApi().getWorkerDataById(workerId, lockScreen, localSpinner).then(
        (response) => {
          if (response.error === null) {
            if (response.data !== undefined)
              this.selectedWorkerData = response.data;
          } else {
            //TODO: error handling
          }
        }
      );
    },
  },
  getters: {
    getListDataHolder (): DataHolder<Emplayee> {
      return this.workers;
    }

  }
});
