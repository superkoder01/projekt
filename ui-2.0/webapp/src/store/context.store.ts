import {defineStore} from "pinia";
import {Customer} from "@/models/customer";
import {Provider} from "@/models/provider";
import {Emplayee} from "@/models/emplayee";
import {BreadcrumbElement} from "@/store/BreadcrumpElement";
import {Contract, Offer} from "@/models/billing/billing";
import factoryApi from "@/api/factory.api";
import {User} from "@/models/user";
import {useUserStore} from "@/store/user.store";
import {RoleEnum} from "@/services/permissions/role-enum";

export const useContextStore = defineStore({
  id: 'contextStore',

  state: () => {
    return {
      selectedCustomer: {} as Customer,
      _selectedProvider: {} as Provider,
      selectedOffer: {} as Offer,
      selectedContract: {} as Contract,
      selectedEmployee: {} as Emplayee,
      // selectedSuperAdmin: {} as User,
      _breadcrumbLastElement: {} as BreadcrumbElement,
      loggedProvider: Object(Provider),
      _currentLoggedUser: Object(User),
      _currentLoggedUserData: Object(Customer) || Object(Emplayee),
      _allowExtendedValidation: true,
    };
  },
  actions: {
    setBreadcrumbLastElement(path:string, value:string) :void {
      this._breadcrumbLastElement = new BreadcrumbElement(path, value);
    },
    setSelectedProvider(provider:Provider){
      this._selectedProvider = provider;
    },
    // setLoggedProvider(provider: Provider) {
    //   this.loggedProvider = provider;
    // },

    fetchProviderInfo() {
      factoryApi.providersApi().getProviderInfo().then(response => {
        if (response.error === null) {
          console.log(response.data);
          this.loggedProvider = response.data;
        } else {
          //TODO: error handling
        }
      });
    },
    fetchCurrentLoggedUserWithData() {
      factoryApi.usersApi().getCurrentLoggedUser(true, null).then(response => {
        if (response.error === null) {
          console.log(response.data);
          this._currentLoggedUser = response.data;
          if(this._currentLoggedUser.workerId) {
            this._currentLoggedUserData = factoryApi.workersApi().getWorkerDataById(this._currentLoggedUser.workerId, true, null)
            .then(response => {
              if (response.error === null) {
                if (response.data !== undefined)
                  this._currentLoggedUserData = response.data;
              } else {
                //TODO: error handling
              }
            });
          } else {
            if(useUserStore().currentRole === RoleEnum.PROSUMER){
              this._currentLoggedUserData = factoryApi.customerApi().getCurrentCustomerData(true, null).then(response => {
                if (response.error === null) {
                  if (response.data !== undefined)
                    this._currentLoggedUserData = response.data;
                } else {
                  //TODO: error handling
                }
              });
            } else{
              console.debug("Error: ")
            }
          }
        } else {
          //TODO: error handling
        }
      });
    }

  },
  getters: {
    breadcrumbLastElement () : BreadcrumbElement {
      return this._breadcrumbLastElement;
    },
    currentLoggedUser () : User  {
      return this._currentLoggedUser;
    },
    currentLoggedUserData() : Customer | Emplayee {
      return this._currentLoggedUserData;
    },
    currentLoggedProvider(): Provider {
      return this.loggedProvider;
    },
    // getSelectedSuperAdmin(): User{
    //   return this.selectedSuperAdmin;
    // },
    getSelectedEmployee(): Emplayee{
      return this.selectedEmployee;
    },
    selectedProvider(): Provider{
      return this._selectedProvider;
    },
    allowExtendedValidation(): boolean {
      return this._allowExtendedValidation;
    },
    getSelectedContract(): Contract {
      return this.selectedContract;
    }
  },
  persist: {
    enabled: true
  },
});
