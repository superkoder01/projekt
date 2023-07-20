import BaseApi from "@/api/base.api";
import {ApiTypeEnum} from "@/services/logger/api-type.enum";
import {Emplayee} from "@/models/emplayee";
import {LocalSpinner} from "@/services/model/localSpinner";
import {RequestResponseApi, ResponseError} from "@/models/request-response-api";
import {AxiosError} from "axios";
import { DataHolder } from "@/models/data-holder";
import { PagingModel } from "@/services/model/paging.model";

export class WorkersApi extends BaseApi {

  protected getApiType(): ApiTypeEnum {
    return ApiTypeEnum.WORKER_API;
  }

  private readonly API_URL = '/api/management';
  private readonly API_WORKERS_URL = this.API_URL + '/workers/';
  private readonly API_WORKERS_QUERY = this.API_WORKERS_URL + 'query/'
  private readonly API_WORKER_CUSTOMER_URL = this.API_URL + '/workers/details';
  private readonly API_WORKER_BY_ID = this.API_URL + '/workers/';

  getCurrentWorkerData(lockScreen: boolean, localSpinner: LocalSpinner | null) :Promise<RequestResponseApi<Emplayee>>{
    return this.getDataFromUrl<Emplayee>(this.API_WORKER_CUSTOMER_URL, lockScreen, localSpinner);
  }

  getWorkerDataById(workerId: number, lockScreen: boolean, localSpinner: LocalSpinner | null,) {
    return this.getDataFromUrl<Emplayee>(this.API_WORKER_BY_ID + workerId, lockScreen, localSpinner);
  }
  getWorkers(lockScreen: boolean, localSpinner: LocalSpinner | null, pagination?: PagingModel) :Promise<RequestResponseApi<DataHolder<Emplayee>>> {
    return this.getDataFromUrl<DataHolder<Emplayee>>(this.API_WORKERS_QUERY,lockScreen, localSpinner, false, pagination);
  }

  saveNewWorker(worker: Emplayee, successCallback: () => void, failCallback: (error: ResponseError) => void) {
    this.axiosCall({
      method: 'POST',
      url: this.API_URL + '/workerUsers',
      data: worker,
    }, true, null, true).then(value => {
      if (value.error === null) {
        if (successCallback !== undefined) {
          successCallback();
        }
      } else if(failCallback !== undefined){
        failCallback(value.error);
      }
    });
  }
  updateWorkerById(workerId: number, worker: Emplayee,successCallback: () => void, failCallback: () => void) {
    this.axiosCall({
      method: 'PUT',
      url: this.API_WORKERS_URL + workerId,
      data: worker
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
  deleteWorker(workerId: number, successCallback: () => void, failCallback: () => void) {
    this.axiosCall({
      method: 'DELETE',
      url: this.API_WORKERS_URL + workerId,
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
