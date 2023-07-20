<template>
  <div>
      <DataTableWrapper :service="customers" :global-filter-fields="['firstName', 'lastName', 'email', 'phone']">
        <template v-slot:empty> {{ $t('CUSTOMERS_VIEW.NOT_FOUND') }}</template>
        <template v-slot:columns>
          <Column field="id" :header="$t('CUSTOMERS_VIEW.TABLE_HEADERS.ID')" sortable></Column>
          <Column field="firstName" :header="$t('CUSTOMERS_VIEW.TABLE_HEADERS.FIRST_NAME')" sortable></Column>
          <Column field="lastName" :header="$t('CUSTOMERS_VIEW.TABLE_HEADERS.LAST_NAME')" sortable></Column>
          <Column field="email" :header="$t('CUSTOMERS_VIEW.TABLE_HEADERS.EMAIL')" sortable></Column>
          <Column field="phone" :header="$t('CUSTOMERS_VIEW.TABLE_HEADERS.PHONE')" sortable></Column>
          <Column field="province" :header="$t('CUSTOMERS_VIEW.TABLE_HEADERS.PROVINCE')" sortable></Column>
          <Column field="customerTypeName" :header="$t('CUSTOMERS_VIEW.TABLE_HEADERS.TYPE')" sortable>
            <template #body="{data}">
              <span>
                {{$t('GLOBALS.ROLES.' + data.customerTypeName)}}
              </span>
            </template>
          </Column>
          <Column style="text-align: right">
          <template #body="slotProps">
            <span class="actions">
              <Button class="delete" @click="deleteCustomer(slotProps.data)"><Icon name="Delete"/>{{ $t('TABLES.ACTIONS.delete') }}</Button>
              <Button type="button" class="preview" @click="setSelectedCustomer(slotProps.data)"><Icon name="FileText"/>{{ $t('TABLES.ACTIONS.preview') }}</Button>
            </span>
          </template>
        </Column>
        </template>
        <template v-slot:paginatorstart ></template>
      </DataTableWrapper>
      <Button :label="$t('GLOBALS.BUTTONS.ADD_CUSTOMER')" @click="router.push('/customers/new_customer')" icon="pi pi-angle-right" icon-pos="right"></Button>
      <Button :label="$t('GLOBALS.BUTTONS.ADD_BUSINESS_CUSTOMER')" @click="router.push('/customers/new_business_customer')" icon="pi pi-angle-right" icon-pos="right"></Button>
  </div>
</template>

<script setup lang="ts">
import Column from "primevue/column";
import DataTableWrapper from '@/components/layout/DataTableWrapper.vue';
import {useRouter} from 'vue-router';
import {Customer} from "@/models/customer";
import {useContextStore} from "@/store/context.store";
import {DataTableService} from "@/services/dataTableService";
import {DefaultSortingModel} from "@/services/model/defaultSorting.model";
import {PagingModel} from "@/services/model/paging.model";
import {LocalSpinner} from "@/services/model/localSpinner";
import {DataHolder} from "@/models/data-holder";
import {useCustomerStore} from "@/store/customers.store";
import {useToast} from "vue-toastification";

const toast = useToast();
const router = useRouter();

const customers: DataTableService<Customer> = {
  getDefaultSorting(): DefaultSortingModel {
    return new DefaultSortingModel("", "");
  },
  fetchListData(pagination: PagingModel, lockScreen: boolean, localSpinner: LocalSpinner | null): void {
    useCustomerStore().fetchCustomersData(lockScreen, localSpinner, pagination);
  },
  getListDataHolder(): DataHolder<Customer> {
    return useCustomerStore().listDataHolder;
  }
};

function setSelectedCustomer(customer: Customer) {
  useContextStore().selectedCustomer = Object.assign({} as Customer, customer);
  router.push('/customers/customer_tabs');
}

const deleteCustomer = (data: Customer) => {
  console.log(data);
  useCustomerStore().deleteCustomer(data.id, onSuccess, onFail)
};

const onSuccess = () => {
  toast.success("Usunięto klienta");
};

const onFail = () => {
  toast.error("Nie udało się usunąć klienta");
};
</script>

<style scoped lang="scss">

</style>
