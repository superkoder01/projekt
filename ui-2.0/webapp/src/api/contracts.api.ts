import BaseApi from "@/api/base.api";
import {ApiTypeEnum} from "@/services/logger/api-type.enum";
import {LocalSpinner} from "@/services/model/localSpinner";
import { RequestResponseApi, ResponseError } from "@/models/request-response-api";
import {DataHolder} from "@/models/data-holder";
import {PagingModel} from "@/services/model/paging.model";
import {Invoice} from "@/models/invoice";
import { Contract, Offer, OfferPayload } from "@/models/billing/billing";

export class ContractsApi extends BaseApi {

  protected getApiType(): ApiTypeEnum {
    return ApiTypeEnum.PROVIDER_API;
  }

  private readonly API_URL = '/api/core';
  private readonly API_CONTRACTS_URL = this.API_URL + '/contract';
  private readonly API_CONTRACTS_BY_CUSTOMER_ID_URL = this.API_URL + '/contract/customer/:id';

  getContracts(lockScreen: boolean, localSpinner: LocalSpinner | null, pagination?: PagingModel): Promise<RequestResponseApi<DataHolder<Contract>>> {
    return this.getDataFromUrl<DataHolder<Contract>>(this.API_CONTRACTS_URL, lockScreen, localSpinner, false, pagination);
  }

  getContractsByCustomerId(customerId: number, lockScreen: boolean, localSpinner: LocalSpinner | null, pagination?: PagingModel): Promise<RequestResponseApi<DataHolder<Contract>>> {
    const url = this.API_CONTRACTS_BY_CUSTOMER_ID_URL.replace(':id', customerId.toString());
    return this.getDataFromUrl<DataHolder<Contract>>(url, lockScreen, localSpinner, false, pagination);
  }

  saveNewContract(newContract: Contract, successCallback: () => void, failCallback: () => void) {
    this.axiosCall({
      method: 'POST',
      url: this.API_CONTRACTS_URL,
      data: newContract
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
  updateContractById(contractId:string, contract: Contract, successCallback: () => void, failCallback: (error : ResponseError) => void) {
    const url = this.API_CONTRACTS_URL+"/"+contractId;
    this.axiosCall<Offer>({
      method: 'PUT',
      url: url,
      data: contract,
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
}
