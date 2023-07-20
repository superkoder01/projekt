import BaseApi from "@/api/base.api";
import {ApiTypeEnum} from "@/services/logger/api-type.enum";
import {LoginData} from "@/models/login-data";
import {AuthResponse} from "@/models/auth-response";
import {Customer} from "@/models/customer";
import {LocalSpinner} from "@/services/model/localSpinner";
import {RequestResponseApi} from "@/models/request-response-api";
import {DataHolder} from "@/models/data-holder";
import {PagingModel} from "@/services/model/paging.model";
import {User} from "@/models/user";

export class UsersApi extends BaseApi {

  protected getApiType(): ApiTypeEnum {
    return ApiTypeEnum.USERS_API;
  }

  private readonly API_URL = '/api/management';
  private readonly API_SUPER_ADMINISTRATORS_URL = this.API_URL + '/users/superAdmins'
  private readonly API_CURRENT_LOGGED_USER_DETAILS_URL = this.API_URL + '/users/details'
  private readonly API_USER_BY_ID = this.API_URL + '/users/:userId'

  getSuperAdmins(lockScreen: boolean, localSpinner: LocalSpinner | null, pagination?: PagingModel) :Promise<RequestResponseApi<DataHolder<User>>>{
    return this.getDataFromUrl<DataHolder<User>>(this.API_SUPER_ADMINISTRATORS_URL, lockScreen, localSpinner, false, pagination);
  }

  getCurrentLoggedUser(lockScreen: boolean, localSpinner: LocalSpinner | null) : Promise<RequestResponseApi<User>> {
    return this.getDataFromUrl<User>(this.API_CURRENT_LOGGED_USER_DETAILS_URL, lockScreen, localSpinner, false);
  }
  getUserById(userId: number, lockScreen: boolean, localSpinner: LocalSpinner | null): Promise<RequestResponseApi<User>> {
    const url = this.API_USER_BY_ID.replace(':userId', userId.toString());
    return this.getDataFromUrl<User>(url, lockScreen, localSpinner);
  }
}
