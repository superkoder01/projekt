<template>
  <div>
<!--    <div class="about">-->
<!--      <h1>This is an providers page</h1>-->
<!--    </div>-->

    <DataTableWrapper :service="providersDataTableService" :global-filter-fields="['name']">
      <template v-slot:empty>No customers found.</template>
      <template v-slot:columns>
        <Column field="isActive" :header="$t('PROVIDERS_VIEW.TABLE_HEADERS.STATUS')" sortable >
          <template #body="{data}">
            <span class="status" v-if="data.status"><Icon name="CheckCircle" /> {{$t('TABLES.SUPER_ADMINS.STATE.ACTIVE')}}</span>
            <span class="status" v-if="!data.status"><Icon name="Circle" /> {{$t('TABLES.SUPER_ADMINS.STATE.INACTIVE')}}</span>
            <!-- <span class="status" :style="{'background-color': getStateColor(data.isActive)}" > {{ data.isActive ? $t('TABLES.SUPER_ADMINS.STATE.ACTIVE') : $t('TABLES.SUPER_ADMINS.STATE.INACTIVE')  }} </span> -->
          </template>
        </Column>
        <Column field="name" :header="$t('PROVIDERS_VIEW.TABLE_HEADERS.PROVIDER')" sortable></Column>
        <Column field="country" :header="$t('PROVIDERS_VIEW.TABLE_HEADERS.COUNTRY')"  sortable></Column>
        <Column field="type" :header="$t('PROVIDERS_VIEW.TABLE_HEADERS.TYPE')"  sortable></Column>
        <Column field="www" :header="$t('PROVIDERS_VIEW.TABLE_HEADERS.WEBSITE')"  sortable></Column>
        <Column style="text-align: right">
          <template #body="slotProps">
            <span class="actions">
<!--              <Button @click="changeProviderStatus(slotProps.data)">{{$t("GLOBALS.BUTTONS."+buttonStatus(slotProps.data.status))}}</Button>-->
              <Button class="delete" disabled><Icon name="Delete" />{{$t('TABLES.ACTIONS.delete')}}</Button>
              <Button type="button" class="preview" @click="showDetails(slotProps.data)"><Icon name="FileText" />{{$t('TABLES.ACTIONS.preview')}}</Button>
            </span>
          </template>
        </Column>
      </template>
      <template v-slot:paginatorstart></template>
    </DataTableWrapper>
    <Button label="DODAJ PARTNERA" @click="router.push('/providers/new_provider')" icon="pi pi-angle-right" icon-pos="right"></Button>
  </div>
</template>

<script setup lang="ts">

import Column from "primevue/column";
import DataTableWrapper from '@/components/layout/DataTableWrapper.vue';
import {Provider} from "@/models/provider";
import {useRouter} from "vue-router";
import {useContextStore} from "@/store/context.store";
import {DataTableService} from "@/services/dataTableService";
import {PagingModel} from "@/services/model/paging.model";
import {LocalSpinner} from "@/services/model/localSpinner";
import {useProvidersStore} from "@/store/providers.store";
import {DataHolder} from "@/models/data-holder";
import {DefaultSortingModel} from "@/services/model/defaultSorting.model";
import { useToast } from "vue-toastification";
import { useI18n } from "vue-i18n";

const router = useRouter();
const toast = useToast();
const i18n = useI18n();

const providersDataTableService: DataTableService<Provider> = {
  getDefaultSorting():DefaultSortingModel{
    return new DefaultSortingModel( "",  "");
  },
  fetchListData(pagination: PagingModel, lockScreen: boolean, localSpinner: LocalSpinner | null): void {
    useProvidersStore().fetchProviders(lockScreen, localSpinner, pagination);
  },
  getListDataHolder(): DataHolder<Provider> {
    return useProvidersStore().getListDataHolder;
  }
};

function showDetails(provider: Provider) {
  console.log("Selected provider: " + JSON.stringify(provider));
  useContextStore().setSelectedProvider(provider);
  router.push('/providers/provider_tabs');
}

const changeProviderStatus = (selectedProvider: Provider) => {
  let providerToUpdate = {...selectedProvider};
  providerToUpdate.status = !providerToUpdate.status;
  //selectedProvider.status = !selectedProvider.status;
  useProvidersStore().updateProvider(providerToUpdate, onSuccess, onFail);
};
const buttonStatus = (status: boolean) => {
  return status ? "DEACTIVATE" : "ACTIVATE";
};

const onSuccess = () => {
  toast.success("Zmieniono status partnera");
};

const onFail = () => {
  toast.error("Zmiana statusu partnera nie powiodła się");
};
</script>

<style scoped>

</style>
<style>

</style>
