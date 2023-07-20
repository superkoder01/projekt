<template>
  <div>
    <div class="card no-border sticky">
      <Steps :model="items" :readonly="true" />
    </div>

    <router-view v-slot="{Component}" v-model:newCustomer="newCustomer" :isBusinessClient="isBusinessClient" :useSelectedCustomer="useSelectedCustomer" @prevPage="prevPage($event)" @nextPage="nextPage($event)" @complete="complete">
      <keep-alive>
        <component :is="Component" />
      </keep-alive>
    </router-view>
  </div>

</template>

<script setup lang="ts">

import Steps from 'primevue/steps';
import {PropType, ref} from 'vue';
import { useRouter } from 'vue-router';
import { Customer as CustomerModel } from '@/components/forms/create-customer/Customer';
import {Customer} from "@/models/customer";
import { useToast } from 'vue-toastification';
import { useUserStore } from '@/store/user.store';
import { useI18n } from 'vue-i18n';
import { useCustomerStore } from '@/store/customers.store';
import { RoleEnum } from '@/services/permissions/role-enum';
import { computed } from '@vue/reactivity';
import {onBeforeMount} from "@vue/runtime-core";
import {useContextStore} from "@/store/context.store";

const router = useRouter();
const toast = useToast();
const userStore = useUserStore();
const i18n = useI18n();

const currentPath = router.currentRoute.value.path;

const props = defineProps({
  useSelectedCustomer: {
    type: Boolean
  },
  isBusinessClient: {
    type: Boolean
  }
});


let items = ref();
function setItems(){
  if (props.isBusinessClient){
    items = computed(()=> [
      {
        label: i18n.t('FORMS.HEADERS.BUSINESS_DATA'),
        to: currentPath + ""
      },
      {
        label: i18n.t('FORMS.HEADERS.PERSONAL_DATA'),
        to: currentPath + "/personal_step_2"
      },
      {
        label: i18n.t('FORMS.HEADERS.ADDRESS'),
        to: currentPath + "/address_step_2",
      }])
  }
  else {
    items = computed(()=> [
      {
        label: i18n.t('FORMS.HEADERS.PERSONAL_DATA'),
        to: currentPath + ""
      },
      {
        label: i18n.t('FORMS.HEADERS.ADDRESS'),
        to: currentPath + "/address_step_2",
      }]);
  }
}

setItems();

onBeforeMount(() => {
  console.log(currentPath)
  if (props.useSelectedCustomer) {
    const selected = useContextStore().selectedCustomer;
    console.log(selected)
    newCustomer.value.city = selected.city;
    newCustomer.value.providerId = selected.providerId;
    newCustomer.value.firstName = selected.firstName;
    newCustomer.value.lastName = selected.lastName;
    newCustomer.value.phone = selected.phone;
    newCustomer.value.customerTypeName = selected.customerTypeName;
    newCustomer.value.pesel = selected.pesel;
    newCustomer.value.email = selected.email;
    newCustomer.value.street = selected.street;
    newCustomer.value.buildingNumber = selected.buildingNumber;
    newCustomer.value.apartmentNumber = selected.apartmentNumber;
    newCustomer.value.postalCode = selected.postalCode;
    newCustomer.value.province = selected.province;
    newCustomer.value.country = selected.country;

  }
});
//const newCustomer = ref<Customer>({} as Customer);
const newCustomer = ref<CustomerModel>({} as CustomerModel);
const nextPage = (event: any) => {
  router.push(items.value[event.pageIndex + 1].to);
};
const prevPage = (event: any) => {
  router.push(items.value[event.pageIndex - 1].to);
};

const complete = () => {
  if(props.useSelectedCustomer) {
    const selected = useContextStore().selectedCustomer;
    let customerToUpdate = {} as Customer;
    customerToUpdate.id = selected.id;
    customerToUpdate.providerId = selected.providerId;
    customerToUpdate.workerId = selected.workerId;
    customerToUpdate.customerTypeName = newCustomer.value.customerTypeName;
    customerToUpdate.status = selected.status;
    customerToUpdate.firstName = newCustomer.value.firstName;
    customerToUpdate.lastName = newCustomer.value.lastName;
    customerToUpdate.country = newCustomer.value.country;
    customerToUpdate.province = newCustomer.value.province;
    customerToUpdate.city = newCustomer.value.city;
    customerToUpdate.postalCode = newCustomer.value.postalCode;
    customerToUpdate.buildingNumber = newCustomer.value.buildingNumber;
    customerToUpdate.apartmentNumber = newCustomer.value.apartmentNumber;
    customerToUpdate.street = newCustomer.value.street;
    customerToUpdate.email = newCustomer.value.email;
    customerToUpdate.phone = newCustomer.value.phone;
    customerToUpdate.pesel = newCustomer.value.pesel;
    useCustomerStore().updateCustomerById(customerToUpdate.id, customerToUpdate, onSuccessUpdate, onFailUpdate);
  }else {
    newCustomer.value.businessType = props.isBusinessClient ? 'B2B' : 'B2C';
    newCustomer.value.providerId = parseInt(userStore.providerId);
    newCustomer.value.roleId = RoleEnum.PROSUMER; // RoleEnum.PROSUMER;
    newCustomer.value.password = ''; //must be set to empty string
    console.log(newCustomer.value);

    useCustomerStore().createCustomer(newCustomer.value, onSuccess, onFail);
  }



};

const onSuccess = () => {
  router.push('/customers/customer_tabs');
  toast.success(i18n.t('GLOBALS.TOASTS.ADDED_CUSTOMER'));
};

const onFail = () => {
  toast.error(i18n.t('GLOBALS.TOASTS.ERROR_POST'));
};

const onSuccessUpdate = () => {
  router.push('/customers/customer_tabs');
  toast.success("Zaktualizowano klienta");
};

const onFailUpdate = () => {
  toast.error("Nie udało się zaktualizować klienta");
};

</script>

<style scoped lang="scss">

</style>
