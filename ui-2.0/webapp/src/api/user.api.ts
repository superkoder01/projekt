import BaseApi from "@/api/base.api";
import {ApiTypeEnum} from "@/services/logger/api-type.enum";
import {LoginData} from "@/models/login-data";
import {AuthResponse} from "@/models/auth-response";
import { NewPassword } from "@/models/new-password";

export class UserApi extends BaseApi {

  protected getApiType(): ApiTypeEnum {
    return ApiTypeEnum.USER_API;
  }

  private readonly API_URL_BASE = '/api/management/'
  private readonly API_URL_AUTHENTICATE = this.API_URL_BASE + 'authenticate';
  private readonly API_URL_PASSWORD = this.API_URL_BASE + 'activate/:activationCode'


  login(loginData: LoginData) {
    return this.axiosCall<AuthResponse>({
      method: 'POST',
      url: this.API_URL_AUTHENTICATE,
      data: loginData},
      true, null, true);
  }
  activateUser(passwordData: NewPassword, userActivateCode: string, successCallback: () => void, failCallback: () => void) {
    const url = this.API_URL_PASSWORD.replace(':activationCode', userActivateCode.toString());
    this.axiosCall({
      method: 'PUT',
      url: url,
      data: passwordData,
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
