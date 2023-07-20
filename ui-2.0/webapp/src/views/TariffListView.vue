<template>
  <DataTableWrapper :service="tariffDataTableService" :global-filter-fields="['']">
    <template v-slot:empty>No pricing found.</template>
    <template v-slot:columns>
      <Column field="id" :header="$t('id')" sortable ></Column>
      <Column field="name" :header="$t('name')" sortable ></Column>
      <Column field="tariffGroupLabelName" :header="$t('tariffGroupLabelName')" sortable ></Column>
      <Column field="startDate" :header="$t('startDate')" type="date" sortable ></Column>
      <Column field="endDate" :header="$t('endDate')" type="date" sortable ></Column>
      <Column :expander="true" headerStyle="width: 3rem" />
    </template>
    <template v-slot:expanded-columns="expandedData">
      <span v-for="(fee, index) of expandedData.expandedData.data.fees" :key="fee.value">
        <div>Opłata {{ index }}: {{fee.price}}</div>
      </span>
    </template>
    <template v-slot:paginatorstart ></template>
  </DataTableWrapper>
  <Button label="UTWÓRZ GRUPĘ TARYFOWĄ" @click="router.push('/product_catalog/tariff_pricing/new_tariff_group')" icon="pi pi-angle-right" icon-pos="right"></Button>
</template>

<script setup lang="ts">

import { TariffGroup } from "@/components/forms/tariff-group/TariffGroup";
import DataTableWrapper from "@/components/layout/DataTableWrapper.vue";
import { DataHolder } from "@/models/data-holder";
import { DataTableService } from "@/services/dataTableService";
import { DefaultSortingModel } from "@/services/model/defaultSorting.model";
import { LocalSpinner } from "@/services/model/localSpinner";
import { PagingModel } from "@/services/model/paging.model";
import { useTariffStore } from "@/store/tariff.store";
import { useRouter } from 'vue-router';

const router = useRouter();

const tariffDataTableService: DataTableService<TariffGroup> = {
  getDefaultSorting():DefaultSortingModel{
    return new DefaultSortingModel( "",  "");
  },
  fetchListData(pagination: PagingModel, lockScreen: boolean, localSpinner: LocalSpinner | null): void {
    //TODO: implement fetching tarifflist
    useTariffStore().getListDataHolder;
  },
  getListDataHolder(): DataHolder<TariffGroup>{
    return useTariffStore().getListDataHolder;
  }
};

const props = defineProps({
  expandedData: {
    type: Object
  }
});

</script>

<style scoped lang="scss">

</style>
