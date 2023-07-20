<template>

  <DataTable :value="service.getListDataHolder().elements"
             class="row-spacing"
             selectionMode="single"
             :showGridlines="false"
             :lazy="true"
             :removable-sort="true"
             :totalRecords="service.getListDataHolder().amount"
             :paginator="true" :rows="10" v-model:first="firstRecord"
             template="FirstPageLink PrevPageLink CurrentPageReport NextPageLink LastPageLink"
             @page="pageEvent($event)"
             @sort="sortEvent($event)"
             sortMode="single"
             v-model:expandedRows="expandedRows"
             :rowsPerPageOptions="[10,20,50]"
  >

    <template #header v-if="slotPassed('header')">
      <slot name="header"></slot>
    </template>
    <template #groupheader v-if="slotPassed('groupheader')">
      <slot name="groupheader"></slot>
    </template>

    <template #empty v-if="slotPassed('empty')">
      <slot name="empty"></slot>
    </template>
    <template #loading v-if="slotPassed('loading')">
      <slot name="loading"></slot>
    </template>

    <slot name="columns"></slot>

    <template v-if="slotPassed('expanded-columns')" #expansion="slotProps">
      <div class="expansion-container">
        <slot :expandedData="slotProps" name="expanded-columns"></slot>
      </div>
    </template>

    <template #paginatorstart v-if="slotPassed('paginatorstart') || slotPassed('paginatorstartCustom')">
      <slot v-if="slotPassed('paginatorstartCustom')" name="paginatorstartCustom"  />
      <Button v-if="slotPassed('paginatorstart')" type="button" icon="pi pi-refresh" class="p-button-text" @click="reload()"/>
    </template>

    <template #paginatorend v-if="slotPassed('paginatorend')" >
      <slot name="paginatorend"></slot>
    </template>
    <template #footer v-if="slotPassed('footer')">
      <slot name="footer"></slot>
    </template>
    <template #groupfooter v-if="slotPassed('groupfooter')">
      <slot name="groupfooter"></slot>
    </template>

  </DataTable>
</template>

<script setup lang="ts">

import {PropType, ref, useSlots, watch} from "vue";
import {DataTablePageEvent, DataTableSortEvent} from "primevue/datatable";
import {DataTableService} from "@/services/dataTableService";
import {PagingModel} from "@/services/model/paging.model";
import {useGlobalFilterStore} from "@/store/global-filter.store";
import {LocalSpinner} from "@/services/model/localSpinner";
import {onBeforeMount} from "@vue/runtime-core";
import {DefaultSortingModel} from "@/services/model/defaultSorting.model";

const globalFilter = useGlobalFilterStore();
const expandedRows = ref([]);
const firstRecord = ref(0);

//TODO: dlaczego ??? to musi być tu tworzone a nie w mount
let paging = new PagingModel(["login"]);
const slots = useSlots();

const props = defineProps({
  service: {
    type: Object as PropType<DataTableService>,
    required: true
  },
  globalFilterFields: {
    type: Array
  }
});


function slotPassed(slotName: string) {
  return slots[slotName];
}

// function test() {
//   console.log(props.columns)
// }

onBeforeMount(() => {
  paging = new PagingModel(props.globalFilterFields as string[]);
  if(props.service.getDefaultSorting() !== undefined){
    paging.setDefaultSorting(props.service.getDefaultSorting() as DefaultSortingModel);
  }
  props.service.fetchListData(paging, true, null);
});


function pageEvent(event: DataTablePageEvent) {
  console.log('onPage: ' + JSON.stringify(event, null, 2));
  props.service.fetchListData(paging.fromDataTablePageEvent(event), true, null);
}

function sortEvent(event: DataTableSortEvent) {
  console.log('onPage: ' + JSON.stringify(event, null, 2));
  props.service.fetchListData(paging.fromDataTableSortEvent(event), true, null);
}

const localSpinner: LocalSpinner = {
  turnOnFunction: () => {
    globalFilter.setIsLoading(true);
  },
  turnOffFunction: () => {
    globalFilter.setIsLoading(false);
  }
};

let timerId = 0;

function onFilterInput() {
  console.log('onFilterInput: ' + globalFilter.getFilter);
  clearTimeout(timerId);
  timerId = setTimeout(() => {
    props.service.fetchListData(paging.updateFilter(globalFilter.getFilter), false, localSpinner);
  }, 500);
}

watch(() => globalFilter.getFilter, () => onFilterInput());


function reload() {
  console.log('reload: ');
  paging.clearSorting();
  props.service.fetchListData(paging, true, null);
}

</script>

<style scoped lang="scss">

</style>

<style lang="scss">
.p-datatable-row-expansion {
  background-color: var(--main-color) !important;
  color: var(--header-text-color) !important;
  overflow: hidden !important;
  border-radius: 10px;
}

.expansion-container {
  display: flex;
  flex-wrap: wrap;

  div{
    width: 100%;
    display: grid;
    span{
      width:100%;
      place-items: center;
      div{
        place-items: start;
      }
    }
  }
  .detail-row{
    padding: 10px 0;
    display: inline-flex;

    .detail-header{
      color: var(--secondary-color);
    }
    .detail-item{
      color: white;
    }
  }
  span {
    flex-grow: 1;
    width: 45%;
    padding: 5px;
    margin-left: 5px;
    margin-right: 5px;
    box-sizing: border-box;

    div {
      color: var(--secondary-color);
    }
  }

}

:root {
  --gray: #D3D3D3;
}

.p-datatable .p-datatable-tbody tr > td {
  border: solid 1px var(--gray);
  border-left: 0;
  border-right: 0;
}

.p-datatable .p-datatable-tbody > tr > td:first-child {
  border-top-left-radius: 10px;
  border-bottom-left-radius: 10px;
  border: solid 1px var(--gray);
  border-right: 0;
}

.p-datatable .p-datatable-tbody > tr > td:last-child {
  border-top-right-radius: 10px;
  border-bottom-right-radius: 10px;
  border: solid 1px var(--gray);
  border-left: 0;
}

.row-spacing table {
  border-collapse: separate;
  border-spacing: 0px 10px;
}

.status {
  padding: 5px 8px;
  border-radius: 7px;
}

.p-datatable .p-datatable-thead > tr > th {
  border-width: 0;
  color: gray;
  padding-bottom: 0px;
}

.p-datatable .p-sortable-column:not(.p-highlight):hover {
  background-color: transparent;
}

div.p-datatable-wrapper > table {
  border-collapse: separate;
}
</style>
