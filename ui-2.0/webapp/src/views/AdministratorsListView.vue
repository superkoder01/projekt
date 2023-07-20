<template>
<div>
  <div>
    <DataTableWrapper :service="providerAdminsDataTableService" :global-filter-fields="['']">
      <template v-slot:empty>{{ $t('EMPLOYEES_VIEW.OTHER.NO_ADMINS_FOUND') }} </template>
      <template v-slot:columns>
        <Column field="isActive" :header="$t('TABLES.SUPER_ADMINS.STATE.TITLE')" sortable >
          <template #body="{data}">
            <span class="status" v-if="data.isActive"><Icon name="CheckCircle" /> {{$t('TABLES.SUPER_ADMINS.STATE.ACTIVE')}}</span>
            <span class="status" v-if="!data.isActive"><Icon name="Circle" /> {{$t('TABLES.SUPER_ADMINS.STATE.INACTIVE')}}</span>
            <!-- <span class="status" :style="{'background-color': getStateColor(data.isActive)}" > {{ data.isActive ? $t('TABLES.SUPER_ADMINS.STATE.ACTIVE') : $t('TABLES.SUPER_ADMINS.STATE.INACTIVE')  }} </span> -->
          </template>
        </Column>
        <Column field="login" :header="$t('TABLES.PROVIDER_ADMINS.LOGIN')" sortable ></Column>
        <Column field="firstName" :header="$t('TABLES.PROVIDER_ADMINS.FIRST_NAME')" sortable ></Column>
        <Column field="lastName" :header="$t('TABLES.PROVIDER_ADMINS.LAST_NAME')"  sortable ></Column>
        <Column field="phone" :header="$t('TABLES.PROVIDER_ADMINS.PHONE')" sortable ></Column>
        <Column style="text-align: right">
          <template #body="slotProps">
            <span class="actions">
              <Button class="delete" disabled><Icon name="Delete" />{{$t('TABLES.ACTIONS.delete')}}</Button>
              <Button type="button" class="preview" @click="showUserDetails(slotProps.data)"><Icon name="FileText" />{{$t('TABLES.ACTIONS.preview')}}</Button>
            </span>
          </template>
        </Column>
      </template>

      <template v-slot:paginatorstart ></template>
    </DataTableWrapper>
        <Button :label="$t('GLOBALS.BUTTONS.ADD_ADMIN')" @click="addProviderAdmin()" icon="pi pi-angle-right"
              icon-pos="right"></Button>
  </div>
</div>
</template>

<script setup lang="ts">
import DataTableWrapper from '@/components/layout/DataTableWrapper.vue';
import {DataTableService} from "@/services/dataTableService";
import {PagingModel} from "@/services/model/paging.model";
import {LocalSpinner} from "@/services/model/localSpinner";
import {DataHolder} from "@/models/data-holder";
import {useProvidersStore} from "@/store/providers.store";
import {useContextStore} from "@/store/context.store";
import { useRouter } from 'vue-router';

import { RoleEnum } from '@/services/permissions/role-enum';
import { computed } from 'vue';
import {DefaultSortingModel} from "@/services/model/defaultSorting.model";
import { Emplayee } from "@/models/emplayee";
import {Provider} from "@/models/provider";

const router = useRouter();

const providerAdminsDataTableService:DataTableService<Provider> = {
  getDefaultSorting():DefaultSortingModel{
    return new DefaultSortingModel( "desc",  "addedDate");
  },
  fetchListData(pagination: PagingModel, lockScreen: boolean, localSpinner: LocalSpinner | null): void {
    const selectedProviderId = useContextStore().selectedProvider.id;
    useProvidersStore().fetchProviderAdministrators(selectedProviderId, lockScreen, localSpinner, pagination);
  },
  getListDataHolder(): DataHolder<Provider> {
    return useProvidersStore().getProviderAdmins;
  }
};

const addProviderAdmin = () => {
    router.push({name: 'admin_personal_step', params: {
    passedRoleID: RoleEnum.ADMINISTRATOR_FULL,
    previousPath: router.currentRoute.value.path,
    selectedProviderId: useContextStore().selectedProvider.id
  }});
};

function showUserDetails(user: Emplayee) {
  console.log(useProvidersStore().getProviderAdmins);
  console.log(user); //TODO - przejscie do szczegolow pracownika
  useContextStore().selectedEmployee = Object.assign({} as Emplayee, user);
  // useContextStore().setBreadcrumbLastElement('/providers/provider_tabs/provider_admins', user.firstName + " " + user.lastName);
  router.push({name: 'provider_admins', params: {userType: 'admin'}});
}

</script>

<style scoped>

</style>
