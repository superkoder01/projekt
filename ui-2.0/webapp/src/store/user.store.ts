import {defineStore} from 'pinia';
import {RoleEnum} from '@/services/permissions/role-enum';
import {TokenObject} from "@/models/token-object";
import {TokenData} from "@/models/token-data";
import {LoginData} from "@/models/login-data";
import factoryApi from "@/api/factory.api";
import {User} from "@/models/user";
import {NewPassword} from '@/models/new-password';
import {useContextStore} from "@/store/context.store";
import {LocalSpinner} from '@/services/model/localSpinner';

interface UserState {
  accessTokenObject: TokenObject | null,
  refreshTokenObject: TokenObject | null,
  _isLoggedIn: boolean,
  _currentLoggedUser: User | null,
  userById: User | null
}

export const useUserStore = defineStore({
  id: 'userStore',
  state: (): UserState => {
    return {
      accessTokenObject: Object(TokenObject),
      refreshTokenObject: Object(TokenObject),
      _isLoggedIn: false,
      _currentLoggedUser: Object(User),
      userById: Object(User)
    };
  },
  actions: {
    setAccessToken(accessToken: TokenObject) {
      this.accessTokenObject = accessToken;
    },
    setRefreshToken(refreshToken: TokenObject) {
      this.refreshTokenObject = refreshToken;
    },
    setIsLoggedIn() {
      this._isLoggedIn = true;
    },
    loggedOut() {
      this._isLoggedIn = false;
      this.accessTokenObject = null;
      this.refreshTokenObject = null;
      this._currentLoggedUser = null;
    },
    login(loginData: LoginData, successCallback: () => void, failCallback: () => void) {
      factoryApi.userApi().login(loginData).then(response => {
        if (response.error === null) {
          if (response.data?.access_token != undefined) {
            this.setAccessToken(response.data?.access_token);
            useContextStore().fetchProviderInfo();
            useContextStore().fetchCurrentLoggedUserWithData();
          }
          if (response.data?.refresh_token != undefined) {
            this.setRefreshToken(response.data?.refresh_token);
          }
          this.setIsLoggedIn();
          if (successCallback !== undefined) {
            successCallback();
          }
        } else {
          // this.logToConsole(LogLevel.ERROR, 'Response value error: ' + value.error + ' data:' + value.data?.data);
          if (failCallback !== undefined) {
            failCallback();
          }
        }
      });
    },
    activateUser(passwordData: NewPassword, userActivateCode: string, successCallback: () => void, failCallback: () => void) {
      factoryApi.userApi().activateUser(passwordData, userActivateCode, successCallback, failCallback);
    },
    fetchUserById(userId: number, lockScreen: boolean, localSpinner: LocalSpinner | null) {
      factoryApi.usersApi().getUserById(userId, lockScreen, localSpinner).then(
        (response) => {
          if (response.error === null) {
            if (response.data !== undefined)
              this.userById = response.data;
          } else {
            //TODO: error handling
          }
        }
      );
    },
  },
  getters: {
    accessToken(): string {
      if (this.accessTokenObject !== null) {
        return this.accessTokenObject.token;
      }
      return "";
    },
    refreshToken(): string {
      if (this.refreshTokenObject !== null) {
        return this.refreshTokenObject.token;
      }
      return "";
    },
    isLoggedIn(): boolean {
      return this._isLoggedIn;
    },
    // currentLoggedUser(): User {
    //   return this._currentLoggedUser;
    // },
    currentRole(): RoleEnum {
      if (this.accessTokenObject !== null) {
        const role = decodeToken(this.accessTokenObject.token).role;
        return RoleEnum[role as keyof typeof RoleEnum];
      }
      return RoleEnum.UNDEFINED;
    },
    providerId(): string {
      if(this.accessTokenObject!== null) {
        return decodeToken(this.accessTokenObject.token).providerId;
      }
      return '';
    },
    userId(): string {
      if(this.accessTokenObject!== null) {
        return decodeToken(this.accessTokenObject.token).userId;
      }
      return '';
    },
    customerAccountId(): string {
      if(this.accessTokenObject!== null) {
        return decodeToken(this.accessTokenObject.token).customerAccountId;
      }
      return '';
    },
    workerId(): string {
      if(this.accessTokenObject!== null) {
        return decodeToken(this.accessTokenObject.token).workerId;
      }
      return '';
    },
  },
  persist: {
    enabled: true
  }
});

function decodeToken(accessToken: string): TokenData {
  const base64Url = accessToken.split('.')[1];
  const base64 = base64Url.replace(/-/g, '+').replace(/_/g, '/');
  const jsonPayload = decodeURIComponent(atob(base64).split('').map(function (c) {
    return '%' + ('00' + c.charCodeAt(0).toString(16)).slice(-2);
  }).join(''));
  return JSON.parse(jsonPayload);
}
