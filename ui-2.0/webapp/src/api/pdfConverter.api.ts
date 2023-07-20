import {ApiTypeEnum} from "@/services/logger/api-type.enum";
import BaseApi from "@/api/base.api";
import {Invoice} from "@/models/invoice/invoice";
import {LocalSpinner} from "@/services/model/localSpinner";
import {Contract, Offer} from "@/models/billing/billing";
import {ResponseError} from "@/models/request-response-api";

export class PdfConverterApi extends BaseApi {

  private readonly PDF_CONVERTER_API_URL = '/api/convert-pdf';

  protected getApiType(): ApiTypeEnum {
    return ApiTypeEnum.PDF_CONVERTER_API;
  }

  downloadDocumentPdf(data: Invoice|Contract|Offer, lockScreen: boolean, localSpinner: LocalSpinner | null, successCallback: (data:any) => void, failCallback: (error : ResponseError) => void, isOpenInNewWindow = true){
    this.axiosCall({
      method: 'POST',
      url: this.PDF_CONVERTER_API_URL,
      data: data,
      responseType: 'blob',
    }, lockScreen, localSpinner, true).then(value => {
      if (value.error === null) {
        if(isOpenInNewWindow){
          this.openInNewWindow(value.data);
        }
        if (successCallback !== undefined) {
          successCallback(value.data);
        }
      }
      else if(failCallback !== undefined){
        failCallback(value.error);
      }
    });
  }



  private openInNewWindow(data: any){
    const reader = new FileReader();
    reader.readAsDataURL(data as Blob);
    reader.onload = (e) => {
      const a = document.createElement('a');
      a.download = 'invoice_.pdf';
      a.href = e.target?.result as string;
      document.body.appendChild(a);
      a.click()
      document.body.removeChild(a);
    }
  }
}

