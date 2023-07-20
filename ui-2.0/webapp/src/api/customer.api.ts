import { DataHolder } from '@/models/data-holder';
import {ApiTypeEnum} from "@/services/logger/api-type.enum";
import {Customer} from "@/models/customer";
import BaseApi from "@/api/base.api";
import {LocalSpinner} from "@/services/model/localSpinner";
import {RequestResponseApi} from "@/models/request-response-api";
import { Customer as CustomerModel} from '@/components/forms/create-customer/Customer';
import { PagingModel } from '@/services/model/paging.model';
import { CheckReponse } from '@/models/check-response';

export class CustomerApi extends BaseApi {

  protected getApiType(): ApiTypeEnum {
    return ApiTypeEnum.CUSTOMER_API;
  }

  private readonly API_URL = '/api/management';
  private readonly API_CUSTOMERS = this.API_URL + '/customerAccounts/query'
  private readonly API_CURRENT_CUSTOMER_URL = this.API_URL + '/customerAccounts/details';
  private readonly API_CUSTOMER_BY_ID = this.API_URL + '/customerAccounts/';
  private readonly API_CUSTOMERS_EMAIL_CHECK = this.API_URL + '/customerAccounts/check'

  async checkEmailAvailability(email:string, lockScreen: boolean, localSpinner: LocalSpinner | null):Promise<boolean>{
    const pagination = new PagingModel(['email'], 0).updateFilter(email);
    const dataFromUrl = this.getDataFromUrl<CheckReponse>(this.API_CUSTOMERS_EMAIL_CHECK, lockScreen, localSpinner, false, pagination);
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

  getCurrentCustomerData(lockScreen: boolean, localSpinner: LocalSpinner | null) :Promise<RequestResponseApi<Customer>>{
    return this.getDataFromUrl<Customer>(this.API_CURRENT_CUSTOMER_URL, lockScreen, localSpinner);
  }
  getCustomersData(lockScreen: boolean, localSpinner: LocalSpinner | null, pagination?: PagingModel) :Promise<RequestResponseApi<DataHolder<Customer>>>{
    return this.getDataFromUrl<DataHolder<Customer>>(this.API_CUSTOMERS, lockScreen, localSpinner, false, pagination);
  }

  getCustomerDataById(customerId: number, lockScreen: boolean, localSpinner: LocalSpinner | null,){
    return this.getDataFromUrl<Customer>(this.API_CUSTOMER_BY_ID + customerId, lockScreen, localSpinner);
  }
  saveNewCustomer(customer: CustomerModel, successCallback: () => void, failCallback: () => void) {
    this.axiosCall({
      method: 'POST',
      url: this.API_URL + '/customerUsers',
      data: customer,
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
  updateCustomerById(customerId: number,customer: Customer, successCallback: () => void, failCallback: () => void) {
    this.axiosCall({
      method: 'PUT',
      url: this.API_CUSTOMER_BY_ID + customerId,
      data: customer,
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
  deleteCustomerById(customerId: number, successCallback: () => void, failCallback: () => void) {
    this.axiosCall({
      method: 'DELETE',
      url: this.API_CUSTOMER_BY_ID + customerId
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
