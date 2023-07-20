import { TariffGroup } from '@/components/forms/tariff-group/TariffGroup';
import  BaseApi  from '@/api/base.api';
import { ApiTypeEnum } from '@/services/logger/api-type.enum';
import { DataHolder } from '@/models/data-holder';
import { RequestResponseApi, ResponseError } from '@/models/request-response-api';
import { LocalSpinner } from '@/services/model/localSpinner';
import { PagingModel } from '@/services/model/paging.model';

export class TariffGroupApi extends BaseApi {

  protected getApiType(): ApiTypeEnum {
    return ApiTypeEnum.TARRIF_GROUP_API;
  }

  private readonly API_BASE = 'api/core/';
  private readonly API_DISTRIBUTION_NETWORK_OPERATOR = this.API_BASE + 'distribution_network_operator';
  private readonly API_PARAMETER_NAME = this.API_BASE + 'parameter_name';
  private readonly API_GROUP_OSD = this.API_BASE + '/tarrif/group/osd'

  getDistributionNetworkOperator(lockScreen: boolean, localSpinner: LocalSpinner | null, pagination?: PagingModel) :Promise<RequestResponseApi<DataHolder<TariffGroup>>> {
    return this.getDataFromUrl<DataHolder<TariffGroup>>(this.API_DISTRIBUTION_NETWORK_OPERATOR, lockScreen, localSpinner, false, pagination);
  }
  getFees(lockScreen: boolean, localSpinner: LocalSpinner | null, pagination?: PagingModel) :Promise<RequestResponseApi<DataHolder<any>>> {
    return this.getDataFromUrl<DataHolder<any>>(this.API_PARAMETER_NAME, lockScreen, localSpinner, false, pagination);
  }
  saveNewTarrifGroup(tarrifGroup: TariffGroup, successCallback: () => void, failCallback: (error : ResponseError) => void) {
    this.axiosCall<TariffGroup>({
      method: 'POST',
      url: this.API_GROUP_OSD,
      data: tarrifGroup,
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
