import {useSplashStore} from '@/store/splash.store';
import {useToast} from 'vue-toastification';
import axios, {AxiosError, AxiosInstance, AxiosRequestConfig} from 'axios';
import {LogLevel} from '@/services/logger/log-level';
import {PagingModel} from "@/services/model/paging.model";
import {LocalSpinner} from "@/services/model/localSpinner";
import {useUserStore} from "@/store/user.store";
import {AuthResponse} from "@/models/auth-response";
import {Mutex} from "async-mutex";
import {LoggedApi} from "@/api/logged.api";
import {RequestResponseApi, ResponseError} from "@/models/request-response-api";
import {useRouter} from "vue-router";

const toast = useToast();
const refreshUrl = '/api/management/refresh';

export default abstract class BaseApi extends LoggedApi {
  protected axiosInstance: AxiosInstance;

  static mutex = new Mutex();

  constructor () {
    super();
    this.axiosInstance = axios.create({
    });
  }

  private refreshCounter = 0;

  protected async axiosCall<T> (config: AxiosRequestConfig, lockScreen: boolean, localSpinner: LocalSpinner | null, skipErrorToast = false, useMutex = true): Promise<RequestResponseApi<T>> {
    let release;
    if(useMutex){
      this.logToConsole(LogLevel.DEBUG,"Waiting for mutex");
      release = await BaseApi.mutex.acquire();
      this.logToConsole(LogLevel.DEBUG,"Mutex ACQUIRED");
    }
    try {
      const userStore = useUserStore();
      if(userStore.isLoggedIn) {
        config.headers = { Authorization: 'Bearer ' + userStore.accessToken};
      } else {
        //TODO
      }
      this.before(lockScreen, localSpinner);

      this.logToConsole(LogLevel.DEBUG, 'Axios Request: ', JSON.stringify(config));
      const data = await this.axiosInstance.request<T>(config);
      this.logToConsole(LogLevel.DEBUG, 'Axios Response', JSON.stringify(data));
      return new RequestResponseApi<T>(null, data.data);
    } catch (err) {
      this.logToConsole(LogLevel.ERROR, 'Axios Response', JSON.stringify(err));
      //Check if 401 and refresh token once
      const error = err as Error | AxiosError;
      const returnedError = Object() as ResponseError;
      if(axios.isAxiosError(error)){
        if(error.response?.status == 401 && this.refreshCounter < 1){
          this.refreshCounter ++;
          const temp = await this.refreshToken().then(isRefreshOk => {
            if(isRefreshOk){
              return this.axiosCall<T>(config, lockScreen, localSpinner, skipErrorToast, false);
            } else {
              //refresh error (refresh token expired)
              this.logToConsole(LogLevel.DEBUG, "refresh token error - logging out");
              useUserStore().loggedOut();
              useRouter().push('Home');
              skipErrorToast = true;
            }
          }).finally(()=> this.refreshCounter--);
          if(temp != undefined) {
            return temp;
          }
        } else {
          if(error.response?.data instanceof Blob) {
            returnedError.message = await (error.response?.data as Blob).text();
          }else{
            returnedError.message = error.response?.data as string;
          }
          returnedError.status = error.response?.status;
          returnedError.statusText = error.response?.statusText;
        }
      } else {
        returnedError.message = error.message;
      }
      if (!skipErrorToast) {
        toast.error('Error requesting service:' + this.getApiType() + error.message);
      }
      return new RequestResponseApi<T>(returnedError);
    } finally {
      this.after(lockScreen, localSpinner);
      if(useMutex && release != undefined){
        release();
        this.logToConsole(LogLevel.DEBUG,"Mutex RELEASED");
      }
    }
  }

  private async refreshToken():Promise<boolean>{
    this.logToConsole(LogLevel.DEBUG, 'Refreshing Token');
    const userStore = useUserStore();
    const config: AxiosRequestConfig = {method:"POST", url:refreshUrl};
    config.headers = { Authorization: 'Bearer ' + userStore.refreshToken};
    try {
      const value = await this.axiosInstance.request<AuthResponse>(config);
      if (value.data.access_token != undefined) {
        useUserStore().setAccessToken(value.data.access_token);
      }
      if (value.data.refresh_token != undefined) {
        useUserStore().setRefreshToken(value.data.refresh_token);
      }
      useUserStore().setIsLoggedIn();
      this.logToConsole(LogLevel.DEBUG, 'Refresh Token OK');
      return true;
    } catch (err) {
      this.logToConsole(LogLevel.DEBUG, 'Refresh Token ERROR: Axios Response', JSON.stringify(err));
      return false;
    }
  }

  protected getDataFromUrl<T>(url: string,  lockScreen: boolean, localSpinner: (LocalSpinner | null) = null, skipErrorToast = false , pagination?: PagingModel) :Promise<RequestResponseApi<T>>{
    return this.axiosCall<T>({
      method: 'GET',
      url: url,
      params: pagination?.toAxiosParams()
    }, true, null, skipErrorToast);
  }

  // protected createUrlWithVariable( url:string, variable:Record<string, string>): string{
  //
  //   return url.replace(':'+variable[])
  // }

  private before (lockScreen: boolean, localSpinner: LocalSpinner | null) {
    if(lockScreen) {
      useSplashStore().increment();
    }
    localSpinner?.turnOnFunction();
  }

  private after (lockScreen: boolean, localSpinner: LocalSpinner | null) {
    if(lockScreen) {
      useSplashStore().decrement();
    }
    localSpinner?.turnOffFunction();
  }
}
