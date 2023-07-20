import BaseApi from "@/api/base.api";
import {ApiTypeEnum} from "@/services/logger/api-type.enum";
import {LocalSpinner} from "@/services/model/localSpinner";
import {RequestResponseApi} from "@/models/request-response-api";
import {DataHolder} from "@/models/data-holder";
import {PagingModel} from "@/services/model/paging.model";
import {Invoice} from "@/models/invoice/invoice";
import {Pricing} from "@/models/billing/billing";

export class InvoicesApi extends BaseApi {

  protected getApiType(): ApiTypeEnum {
    return ApiTypeEnum.PROVIDER_API;
  }

  private readonly API_URL = '/api/core';
  private readonly API_INVOICES_URL = this.API_URL + '/invoice';
  private readonly API_INVOICE_BY_ID_URL = this.API_URL + '/invoice/:id';
  private readonly API_INVOICES_BY_CUSTOMER_ID_URL = this.API_URL + '/invoice/customer/:id';

  getInvoices(lockScreen: boolean, localSpinner: LocalSpinner | null, pagination?: PagingModel): Promise<RequestResponseApi<DataHolder<Invoice>>> {
    return this.getDataFromUrl<DataHolder<Invoice>>(this.API_INVOICES_URL, lockScreen, localSpinner, false, pagination);
  }

  getInvoiceById(invoiceId: number, lockScreen: boolean, localSpinner: LocalSpinner | null){
    const url = this.API_INVOICE_BY_ID_URL.replace(':id', invoiceId.toString());
    return this.getDataFromUrl<Invoice>(url, lockScreen, localSpinner, false);
  }

  getInvoicesByCustomerId(customerId: number, lockScreen: boolean, localSpinner: LocalSpinner | null, pagination?: PagingModel): Promise<RequestResponseApi<DataHolder<Invoice>>> {
    const url = this.API_INVOICES_BY_CUSTOMER_ID_URL.replace(':id', customerId.toString());
    return this.getDataFromUrl<DataHolder<Invoice>>(url, lockScreen, localSpinner, false, pagination);
  }
}
