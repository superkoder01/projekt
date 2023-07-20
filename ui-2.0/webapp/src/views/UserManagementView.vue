<template>
<span>
  <span v-if="mainView">
    <DataTableWrapper :service="workers" :global-filter-fields="['firstName','lastName','email','phone']">
      <template v-slot:empty>{{ $t('EMPLOYEES_VIEW.OTHER.NO_EMPLOYEES_FOUND') }} </template>
      <template v-slot:columns>
        <Column field="isActive" :header="$t('TABLES.SUPER_ADMINS.STATE.TITLE')" sortable >
            <template #body="{data}">
            <div style="width: max-content"> 
              <span class="status" v-if="data.status"><Icon name="CheckCircle" /> {{$t('TABLES.SUPER_ADMINS.STATE.ACTIVE')}}</span>
              <span class="status" v-if="!data.status"><Icon name="Circle" /> {{$t('TABLES.SUPER_ADMINS.STATE.INACTIVE')}}</span>
            </div>
            </template>
          </Column>
        <Column field="firstName" :header="$t('EMPLOYEES_VIEW.TABLE_HEADERS.FIRST_NAME')" sortable ></Column>
        <Column field="lastName" :header="$t('EMPLOYEES_VIEW.TABLE_HEADERS.LAST_NAME')"  sortable ></Column>
        <Column field="role" :header="$t('EMPLOYEES_VIEW.TABLE_HEADERS.ROLE')"  sortable >
          <template #body="{data}">
            <span>
              {{$t('GLOBALS.ROLES.' + data.role)}}
            </span>
          </template>
        </Column>
        <Column field="email" :header="$t('EMPLOYEES_VIEW.TABLE_HEADERS.EMAIL')"  sortable ></Column>
        <Column field="phone" :header="$t('EMPLOYEES_VIEW.TABLE_HEADERS.PHONE')"  sortable ></Column>
        <Column style="text-align: right">
            <template #body="slotProps">
              <span class="actions">
                <Button class="delete" disabled><Icon name="Delete" />{{$t('TABLES.ACTIONS.delete')}}</Button>
                <Button type="button" class="preview" @click="showEmployeeDetails(slotProps.data)"><Icon name="FileText" />{{$t('TABLES.ACTIONS.preview')}}</Button>
              </span>
            </template>
          </Column>
      </template>

      <template v-slot:paginatorstart ></template>
    </DataTableWrapper>
    <Button v-if="permissionService.canAccessFeature(FeatureEnum.MANAGE_FULL_ADMIN)" :label="$t('GLOBALS.BUTTONS.ADD_ADMIN')" @click="setProperCreateUserView(RoleEnum.ADMINISTRATOR_FULL, 'admin_personal_step')" icon="pi pi-angle-right"
              icon-pos="right"></Button>
    <Button v-if="permissionService.canAccessFeature(FeatureEnum.MANAGE_TRADER)" :label="$t('GLOBALS.BUTTONS.ADD_TRADER')" @click="setProperCreateUserView(RoleEnum.TRADER, 'trader_personal_step')" icon="pi pi-angle-right"
            icon-pos="right"></Button>
    <Button v-if="permissionService.canAccessFeature(FeatureEnum.MANAGE_SUPER_AGENT)" :label="$t('GLOBALS.BUTTONS.ADD_SUPER_AGENT')" @click="setProperCreateUserView(RoleEnum.SUPER_AGENT, 'super_agent_personal_step')" icon="pi pi-angle-right"
                icon-pos="right"></Button>
    <Button v-if="permissionService.canAccessFeature(FeatureEnum.MANAGE_AGENT)" :label="$t('GLOBALS.BUTTONS.ADD_AGENT')" @click="setProperCreateUserView(RoleEnum.AGENT, 'agent_personal_step')" icon="pi pi-angle-right"
            icon-pos="right"></Button>
  </span>
  <router-view v-if="!mainView"></router-view>
</span>

</template>

<script setup lang="ts">
import DataTableWrapper from '@/components/layout/DataTableWrapper.vue';
import { PermissionsService } from '@/services/permissions/permissions.service';
import { useRouter } from 'vue-router';
import { FeatureEnum } from "@/services/permissions/feature-enum";
import { RoleEnum } from '@/services/permissions/role-enum';
import { useUserStore } from '@/store/user.store';
import { Emplayee } from '../models/emplayee';
import {useContextStore} from "@/store/context.store";
import { computed } from 'vue';
import { DataTableService } from '@/services/dataTableService';
import { DefaultSortingModel } from '@/services/model/defaultSorting.model';
import { DataHolder } from '@/models/data-holder';
import { useWorkersStore } from '@/store/workers.store';
import { LocalSpinner } from '@/services/model/localSpinner';
import { PagingModel } from '@/services/model/paging.model';


const router = useRouter();
const currentPath = router.currentRoute.value.path;
const permissionService = new PermissionsService();
const userStore = useUserStore();
const workersStore = useWorkersStore();

const workers: DataTableService<Emplayee> = {
  getDefaultSorting():DefaultSortingModel{
    return new DefaultSortingModel( "",  "");
  },
  fetchListData(pagination: PagingModel, lockScreen: boolean, localSpinner: LocalSpinner | null):void {
    useWorkersStore().fetchWorkersData(lockScreen, localSpinner, pagination);
  },
  getListDataHolder(): DataHolder<Emplayee> {
    return workersStore.getListDataHolder;
  }
};

function setProperCreateUserView(userEnum: RoleEnum, name: string) {
    if(userEnum === RoleEnum.ADMINISTRATOR_FULL) {
      router.push({name: name, params: {passedRoleID: userEnum, selectedProviderId: userStore.providerId}});
    }
    router.push({name: name});
}

function showEmployeeDetails(employee: Emplayee) {
  console.log(employee); //TODO - przejscie do szczegolow pracownika
  useContextStore().selectedEmployee = Object.assign({} as Emplayee, employee);
  useContextStore().setBreadcrumbLastElement('/employees/details', employee.firstName + " " + employee.lastName);
  router.push({name: 'employee_details_view'});
}

const mainView = computed(() => {
  console.log(router.currentRoute.value);
  if(router.currentRoute.value.matched.length > 1) {
    return false;
  } else {
    return true;
  }
});

</script>

<style scoped lang="scss">
button{
  margin-left: 3rem;
  margin-bottom: 0.3rem;
}
</style>
