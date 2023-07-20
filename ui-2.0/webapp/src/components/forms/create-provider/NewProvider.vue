<template>
  <div>
    <div class="card no-border sticky">
      <Steps :model="items" :readonly="true" />
    </div>

    <router-view v-slot="{Component}" :useSelectedProvider="useSelectedProvider" v-model:newProvider="newProvider" v-model:newProviderAdmin="newProviderAdmin" @prevPage="prevPage($event)" @nextPage="nextPage($event)" @complete="complete($event)">
      <keep-alive>
        <component :is="Component" />
      </keep-alive>
    </router-view>
  </div>

</template>

<script setup lang="ts">
import Steps from 'primevue/steps';
import { ref } from 'vue';
import { useRouter } from 'vue-router';
import { useToast } from 'vue-toastification';
import { useI18n } from 'vue-i18n';
import { Provider } from '@/models/provider';
import { FunctionalUser } from '../create-func-user/FunctionalUser';
import { useProvidersStore } from '@/store/providers.store';
import {AxiosError} from "axios";
import {ResponseError} from "@/models/request-response-api";
import { computed } from '@vue/reactivity';
import {onBeforeMount} from "@vue/runtime-core";
import {useContextStore} from "@/store/context.store";

const router = useRouter();
const toast = useToast();
const i18n = useI18n();

const props = defineProps({
  useSelectedProvider: {
    type: Boolean
  },
});

onBeforeMount(() => {
  if(props.useSelectedProvider) {
    const selectedProvider = useContextStore().selectedProvider
    newProvider.value.id = selectedProvider.id;
    newProvider.value.name = selectedProvider.name;
    newProvider.value.type = selectedProvider.type;
    newProvider.value.status = selectedProvider.status;
    newProvider.value.nip = selectedProvider.nip;
    newProvider.value.regon = selectedProvider.regon;
    newProvider.value.krs = selectedProvider.krs;
    newProvider.value.email = selectedProvider.email;
    newProvider.value.phoneNumber = selectedProvider.phoneNumber;
    newProvider.value.blockchainAccAddress = selectedProvider.blockchainAccAddress;
    newProvider.value.street = selectedProvider.street;
    newProvider.value.buildingNumber = selectedProvider.buildingNumber;
    newProvider.value.apartmentNumber = selectedProvider.apartmentNumber;
    newProvider.value.postalCode = selectedProvider.postalCode;
    newProvider.value.province = selectedProvider.province;
    newProvider.value.city = selectedProvider.city;
    newProvider.value.country = selectedProvider.country;
    newProvider.value.licenseExpirationDate = selectedProvider.licenseExpirationDate;
    newProvider.value.www = selectedProvider.www;
  }
});

const currentPath = router.currentRoute.value.path;


const items = computed(() => [
  {
    label: i18n.t('FORMS.HEADERS.PROVIDER_DATA'),
    to: currentPath + ""
  },
  {
    label: i18n.t('FORMS.HEADERS.CONTACT'),
    to: currentPath + "/address_2",
  },
    {
    label: i18n.t('FORMS.HEADERS.ADMIN_DATA'),
    to: currentPath + "/admin_3",
  }
]);
const newProvider = ref<Provider>({} as Provider);
const newProviderAdmin = ref<FunctionalUser>({} as FunctionalUser);


const nextPage = (event: any) => {
  router.push(items.value[event.pageIndex + 1].to);
};
const prevPage = (event: any) => {
  router.push(items.value[event.pageIndex - 1].to);
};

const complete = (event: any) => {
  if(props.useSelectedProvider) {
    useProvidersStore().updateProvider(newProvider.value, onSuccessEdit, onFailEdit);
  } else {
    if(!event.isAdminAdded){
      useProvidersStore().createNewProvider(newProvider.value, onSuccess, onFail);
    }
    else{
      useProvidersStore().createNewProviderWithAdmin(newProvider.value, newProviderAdmin.value, onSuccess, onFail);
    }
  }

};

const onSuccess = () => {
  router.push('/providers');
  toast.success(i18n.t('GLOBALS.TOASTS.ADDED_PROVIDER'));
};

const onFail = (error: ResponseError ) => {
  toast.error(i18n.t('GLOBALS.TOASTS.ERROR_POST') + "\n" + error.status + " " + error.message);
};

const onSuccessEdit = () => {
  router.push('/providers');
  toast.success("Pomyślnie zaktualizowano partnera");
};

const onFailEdit = (error: ResponseError ) => {
  toast.error("Aktaulizacja nie powiodła się" + "\n" + error.status + " " + error.message);
};
</script>




<style scoped lang="scss">

</style>
