<template>
  <span>
    <ConfirmPopup></ConfirmPopup>
    <h2>{{ $t("SECTION_TITLES.SUPER_ADMINS") }}</h2>
    <DataTableWrapper :autoLayout="true" :service="superAdmins" :global-filter-fields="['firstName','lastName', 'email', 'phone']">
      <template v-slot:empty>{{ $t('TABLES.SUPER_ADMINS.NOT_FOUND')}}</template>
      <template v-slot:columns>
        <Column field="firstName" :header="$t('TABLES.SUPER_ADMINS.FIRST_NAME')" sortable ></Column>
        <Column field="lastName" :header="$t('TABLES.SUPER_ADMINS.LAST_NAME')" sortable ></Column>
         <Column field="email" :header="$t('TABLES.SUPER_ADMINS.EMAIL')" sortable ></Column>
        <Column field="phone" :header="$t('TABLES.SUPER_ADMINS.PHONE')" sortable ></Column>
        <Column field="isActive" :header="$t('TABLES.SUPER_ADMINS.STATE.TITLE')" sortable >
          <template #body="{data}">
            <span class="status" v-if="data.isActive"><Icon name="CheckCircle" /> {{$t('TABLES.SUPER_ADMINS.STATE.ACTIVE')}}</span>
            <span class="status" v-if="!data.isActive"><Icon name="Circle" /> {{$t('TABLES.SUPER_ADMINS.STATE.INACTIVE')}}</span>
            <!-- <span class="status" :style="{'background-color': getStateColor(data.isActive)}" > {{ data.isActive ? $t('TABLES.SUPER_ADMINS.STATE.ACTIVE') : $t('TABLES.SUPER_ADMINS.STATE.INACTIVE')  }} </span> -->
          </template>
        </Column>
        <Column style="text-align: right">
          <template #body="slotProps">
            <span class="actions">
              <Button class="delete" @click="deletePopup($event, slotProps.data)" :disabled="isDisabled(slotProps.data.workerId)"><Icon name="Delete" />{{$t('TABLES.ACTIONS.delete')}}</Button>
              <Button class="preview" @click="showEmployeeDetails(slotProps.data)"><Icon name="FileText" />{{$t('TABLES.ACTIONS.preview')}}</Button>
            </span>
            <!-- <Button type="button" style="background-color: red;" icon="pi pi-trash" @click="deleteSuperAdmin(data)" :disabled="isDisabled(data.workerId)"></Button> -->
          </template>
        </Column>
      </template>
      <template v-slot:paginatorstart ></template>
    </DataTableWrapper>
        <Button :label="$t('GLOBALS.BUTTONS.ADD_SUPER_ADMIN')" @click="addSuperAdmin()" icon="pi pi-angle-right" icon-pos="right"></Button>
  </span>
</template>

<script setup lang="ts">
import DataTableWrapper from '../components/layout/DataTableWrapper.vue';
import { RoleEnum } from '../services/permissions/role-enum';
import { useRouter } from 'vue-router';
import {DataTableService} from "../services/dataTableService";
import {PagingModel} from "../services/model/paging.model";
import {LocalSpinner} from "../services/model/localSpinner";
import {useUsersStore} from "../store/users.store";
import {DataHolder} from "../models/data-holder";
import { FunctionalUser } from '../components/forms/create-func-user/FunctionalUser';
import { useWorkersStore } from '../store/workers.store';
import { useToast } from 'vue-toastification';
import { useUserStore } from '@/store/user.store';
import { useI18n } from 'vue-i18n';
import { User } from '../models/user';
import { useConfirm } from "primevue/useconfirm";
import {useContextStore} from "@/store/context.store";
import {DefaultSortingModel} from "@/services/model/defaultSorting.model";
import {Emplayee} from "@/models/emplayee";

const i18n = useI18n();
const toast = useToast();
const confirm = useConfirm();

const superAdmins: DataTableService<User> = {
  getDefaultSorting():DefaultSortingModel{
    return new DefaultSortingModel( "",  "");
  },
  fetchListData(pagination: PagingModel, lockScreen: boolean, localSpinner: LocalSpinner | null): void {
    useUsersStore().fetchSuperAdmins(lockScreen, localSpinner, pagination);
  },
  getListDataHolder(): DataHolder<User> {
    return useUsersStore().getListDataHolder;
  }
};

const router = useRouter();

// function getStateColor(status: boolean): string {
//   if(status === true) return '#ccffcc';
//   return '#ff9999';
// }

function isDisabled(workerId: number): boolean{
  if(useUserStore().workerId != undefined
    && workerId != undefined
    && useUserStore().workerId == workerId.toString()){
    return true;
  }
  return false;
}

function addSuperAdmin(){
  router.push({name: 'super_admin_personal_step', params: {
    passedRoleID: RoleEnum.SUPER_ADMIN,
    previousPath: router.currentRoute.value.path,
    selectedProviderId: useUserStore().providerId,
  }});
}

function showEmployeeDetails(employee: Emplayee) {
  console.log(employee); //TODO - przejscie do szczegolow pracownika
  useContextStore().selectedEmployee = employee;
  router.push('/super_admins/details');
}

const deletePopup = (event, superAdmin: FunctionalUser) => {
  confirm.require({
    target: event.currentTarget,
    message: i18n.t('GLOBALS.POPUPS.SUPERADMIN.delete_message'),
    icon: 'pi pi-info-circle',
    acceptClass: 'p-button-danger',
    acceptLabel: i18n.t('GLOBALS.POPUPS.GENERAL.yes'),
    rejectLabel: i18n.t('GLOBALS.POPUPS.GENERAL.no'),
    accept: () => {
      deleteSuperAdmin(superAdmin);
    },
    reject: () => {
      console.log('SUPERUSER DELETE TERMINATED');
    }
  });
};

function deleteSuperAdmin(admin: FunctionalUser){
  console.log("deleteSuperAdmin " + JSON.stringify(admin));
  useWorkersStore().deleteWorker(admin.workerId, onSuccess, onFail);
}
const onSuccess = () => {
  location.reload();
  toast.success(i18n.t('GLOBALS.TOASTS.ADDED_SUPER_ADMIN'));
};

const onFail = () => {
  toast.error(i18n.t('GLOBALS.TOASTS.ERROR_DELETE'));
};
</script>

<style scoped lang="scss">

</style>
