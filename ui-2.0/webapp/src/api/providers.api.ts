import { FunctionalUser } from './../components/forms/create-func-user/FunctionalUser';
import BaseApi from "@/api/base.api";
import {ApiTypeEnum} from "@/services/logger/api-type.enum";
import {LocalSpinner} from "@/services/model/localSpinner";
import {RequestResponseApi, ResponseError} from "@/models/request-response-api";
import {DataHolder} from "@/models/data-holder";
import {PagingModel} from "@/services/model/paging.model";
import {Provider} from "@/models/provider";
import {AxiosError} from "axios";
import { CheckReponse } from '@/models/check-response';

export class ProvidersApi extends BaseApi {

  protected getApiType(): ApiTypeEnum {
    return ApiTypeEnum.PROVIDER_API;
  }

  private readonly API_URL = '/api/management';
  private readonly API_PROVIDERS_URL = this.API_URL + '/providers/query';
  private readonly API_SELECTED_PROVIDER_URL = this.API_URL + '/providers/';
  private readonly API_NEW_PROVIDER_URL = this.API_URL + '/providers/';
  private readonly API_PROVIDERS_ADMINS_URL = this.API_URL + '/providers/:id/administrators';
  private readonly API_PROVIDERS_ADD_ADMIN = this.API_URL + '/workerUsers';
  private readonly API_UPDATE_PROVIDER = this.API_URL + '/providers/:providerId';
  private readonly API_PROVIDER_FIELD_CHECK = this.API_URL + '/providers/check';

  getProviders(lockScreen: boolean, localSpinner: LocalSpinner | null, pagination?: PagingModel): Promise<RequestResponseApi<DataHolder<Provider>>> {
    return this.getDataFromUrl<DataHolder<Provider>>(this.API_PROVIDERS_URL, lockScreen, localSpinner, false, pagination);
  }

  getProviderById(providerId: number, lockScreen: boolean, localSpinner: LocalSpinner | null, pagination?: PagingModel): Promise<RequestResponseApi<Provider>> {
    return this.getDataFromUrl<Provider>(this.API_SELECTED_PROVIDER_URL + providerId, lockScreen, localSpinner, false, pagination);
  }

  getProviderInfo(): Promise<RequestResponseApi<Provider>> {
    return this.getDataFromUrl<Provider>('/api/management/providers/details', true, null, false);
  }

  getProviderAdmins(providerId: number, lockScreen: boolean, localSpinner: LocalSpinner | null, pagination?: PagingModel): Promise<RequestResponseApi<DataHolder<Provider>>> {
    const url = this.API_PROVIDERS_ADMINS_URL.replace(':id', providerId.toString());
    return this.getDataFromUrl<DataHolder<Provider>>(url, lockScreen, localSpinner, false, pagination);
  }
  saveNewProvider(provider: Provider, successCallback: () => void, failCallback: (error : ResponseError) => void) {
    this.axiosCall<Provider>({
      method: 'POST',
      url: this.API_NEW_PROVIDER_URL,
      data: provider,
    }, true, null, true).then(value => {
      if (value.error === null) {
        if (successCallback !== undefined) {
          successCallback();
        }
      }
      else if(failCallback !== undefined){
        failCallback(value.error);
      }
    });
  }
  saveNewProviderWithAdmin(provider: Provider, admin: FunctionalUser, successCallback: () => void, failCallback: (error : ResponseError) => void){
    this.axiosCall<Provider>({
      method: 'POST',
      url: this.API_NEW_PROVIDER_URL,
      data: provider,
    }, true, null, true).then(value => {
      if (value.error === null) {
        if (successCallback !== undefined) {
          admin.providerId = value.data?.id as number;
          this.axiosCall({
            method: 'POST',
            url: this.API_PROVIDERS_ADD_ADMIN,
            data: admin,
          }, true, null, true);
          successCallback();
        }
      }
      else if(failCallback !== undefined){
        failCallback(value.error);
      }
    });
  }
  updateProvider(provider: Provider, successCallback: () => void, failCallback: (error : ResponseError) => void) {
    const url = this.API_UPDATE_PROVIDER.replace(':providerId', provider.id.toString());
    this.axiosCall<Provider>({
      method: 'PUT',
      url: url,
      data: provider,
    }, true, null, true).then(value => {
      if (value.error === null) {
        if (successCallback !== undefined) {
          successCallback();
        }
      }
      else if(failCallback !== undefined){
        failCallback(value.error);
      }
    });
  }
  async checkFieldAvailability(field: string, fieldName: string, lockScreen: boolean, localSpinner: LocalSpinner | null):Promise<boolean>{
    const pagination = new PagingModel([fieldName], 0).updateFilter(field);
    const dataFromUrl = this.getDataFromUrl<CheckReponse>(this.API_PROVIDER_FIELD_CHECK, lockScreen, localSpinner, false, pagination);
    return dataFromUrl.then(value => {
      if (value.error === null) {
        if(value.data!=undefined) {
          return !value.data.checkResult;
        }
      }
      return true;
    });
    return true;
  }
}
