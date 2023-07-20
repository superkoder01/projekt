import { User } from '@/models/user';
import { defineStore } from 'pinia';
import {DataHolder} from "@/models/data-holder";
import {LocalSpinner} from "@/services/model/localSpinner";
import {PagingModel} from "@/services/model/paging.model";
import factoryApi from "@/api/factory.api";

export const useUsersStore = defineStore({
  id: 'usersStore',
  state: () => {
    return {
      users : Object(new DataHolder<User>())
    };
  },
  actions: {
    fetchSuperAdmins(lockScreen: boolean, localSpinner: LocalSpinner | null, pagination: PagingModel){
      factoryApi.usersApi().getSuperAdmins(lockScreen, localSpinner, pagination).then(
        (response) => {
          if (response.error === null) {
            this.users = response.data;
          } else {
            //TODO: error handling
          }
        }
      );
    },
    // fetchSuperAdminById(userId: number, lockScreen: boolean, localSpinner: LocalSpinner | null) {
    //   factoryApi.usersApi().getUserById(userId, lockScreen, localSpinner).then(
    //     (response) => {
    //       if(response.error === null) {
    //         this.selectedUserData = response.data as User;
    //       } else {
    //         //error handling
    //       }
    //     }
    //   );
    // },
  },
  getters: {
    getListDataHolder ():DataHolder<User> {
      return this.users;
    }
  }
});
