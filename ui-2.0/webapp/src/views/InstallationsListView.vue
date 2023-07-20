<template>

    <DataTableWrapper :service="installationsDataTableService" :global-filter-fields="['city']">

      <template v-slot:empty>{{ $t('INSTALLATIONS_VIEW.OTHER.NOT_FOUND')}}</template>
      <template v-slot:columns>
        <Column field="meterNumber" :header="$t('INSTALLATIONS_VIEW.TABLE_HEADERS.METER_NUMBER')" sortable></Column>
        <Column field="address" :header="$t('INSTALLATIONS_VIEW.TABLE_HEADERS.ADDRESS')" sortable></Column>
        <Column field="city" :header="$t('INSTALLATIONS_VIEW.TABLE_HEADERS.CITY')" sortable></Column>
        <Column field="sapCode" :header="$t('INSTALLATIONS_VIEW.TABLE_HEADERS.PPE_NUMBER')" sortable></Column>
<!--        <Column field="" :header="$t('INSTALLATIONS_VIEW.TABLE_HEADERS.INSTALLATION_POWER')" sortable></Column>-->
<!--        <Column field="" :header="$t('INSTALLATIONS_VIEW.TABLE_HEADERS.OSD')" sortable></Column>-->
      </template>

      <template v-slot:paginatorstart ></template>
    </DataTableWrapper>
    <Button v-if="props.useSelectedCustomer" @click="router.push('/customers/customer_tabs/installations/new_installation')"> {{ $t('GLOBALS.BUTTONS.ADD_INSTALLATION')}}</Button>
</template>

<script setup lang="ts">
import DataTableWrapper from '@/components/layout/DataTableWrapper.vue';
import { DataHolder } from '@/models/data-holder';
import { DataTableService } from '@/services/dataTableService';
import { LocalSpinner } from '@/services/model/localSpinner';
import { PagingModel } from '@/services/model/paging.model';
import { useInstallationsStore } from '@/store/installations.store';
import { useRouter } from 'vue-router';
import { useContextStore } from "@/store/context.store";
import {DefaultSortingModel} from "@/services/model/defaultSorting.model";
import {Installation} from "@/components/forms/create-installation/Installation";
const router = useRouter();
const installationsStore = useInstallationsStore();

const props = defineProps({
  useSelectedCustomer:{
    type: Boolean,
    default:false
  }
});

const installationsDataTableService : DataTableService<Installation> = {
  getDefaultSorting():DefaultSortingModel{
    return new DefaultSortingModel( "",  "");
  },
  fetchListData(pagination: PagingModel, lockScreen: boolean, localSpinner: LocalSpinner | null): void {
    if(props.useSelectedCustomer) {
      const selectedCustomerId = useContextStore().selectedCustomer.id;
      installationsStore.fetchSelectedCustomerAccessPoints(selectedCustomerId, true, null);
    } else{
      installationsStore.fetchInstallations(lockScreen, localSpinner, pagination);
    }
  },
  getListDataHolder(): DataHolder<Installation> {
    return installationsStore.getListDataHolder;
  }
};

</script>

<style scoped lang="scss">

</style>
