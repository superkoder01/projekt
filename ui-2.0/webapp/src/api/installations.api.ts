import BaseApi from '@/api/base.api';
import { Installation } from '@/components/forms/create-installation/Installation';
import { DataHolder } from '@/models/data-holder';
import { RequestResponseApi } from '@/models/request-response-api';
import { ApiTypeEnum } from '@/services/logger/api-type.enum';
import { LocalSpinner } from '@/services/model/localSpinner';
import { PagingModel } from '@/services/model/paging.model';


export class InstallationsApi extends BaseApi {

    protected getApiType(): ApiTypeEnum {
        return ApiTypeEnum.INSTALLATION_API;
    }

    private readonly API_BASE = '/api/management';
    private readonly API_SERVICE_POINTS_URL = this.API_BASE + '/serviceAccessPoints';
    private readonly API_SERVICE_POINTS_QUERY_URL = this.API_BASE + '/serviceAccessPoints/query';
    private readonly API_SERVICE_POINTS_BY_URL = this.API_BASE + 'serviceAccessPoints/:serviceAccessPointId';
    private readonly API_ACCESS_POINTS =this.API_BASE+ '/customerAccounts/:customerAccountId/serviceAccessPoints';

    getInstallations(lockScreen: boolean, localSpinner: LocalSpinner | null, pagination?: PagingModel): Promise<RequestResponseApi<DataHolder<Installation>>> {
        return this.getDataFromUrl<DataHolder<Installation>>(this.API_SERVICE_POINTS_QUERY_URL, lockScreen, localSpinner, false, pagination);
    }
    getInstallationById(installationId: number, lockScreen: boolean, localSpinner: LocalSpinner | null, pagination?: PagingModel): Promise<RequestResponseApi<DataHolder<Installation>>> {
        const url = this.API_SERVICE_POINTS_BY_URL.replace(':serviceAccessPointId', installationId.toString());
        return this.getDataFromUrl<DataHolder<Installation>>(url, lockScreen, localSpinner, false, pagination);
    }
    getCustomerAccessPoints(customerId: number, lockScreen: boolean, localSpinner: LocalSpinner | null,pagination?: PagingModel){
      const url = this.API_ACCESS_POINTS.replace(':customerAccountId', customerId.toString());
      return this.getDataFromUrl<DataHolder<Installation>>(url, lockScreen, localSpinner, false, pagination);
    }
    saveNewInstallation(installation: Installation, successCallback: () => void, failCallback: () => void) {
        this.axiosCall<Installation>({
          method: 'POST',
          url: this.API_SERVICE_POINTS_URL,
          data: installation,
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
