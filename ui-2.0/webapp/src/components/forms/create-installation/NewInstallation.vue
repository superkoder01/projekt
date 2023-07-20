<template>
  <div>
    <div class="card">
      <Steps :model="items" :readonly="true"/>
    </div>
    <router-view v-slot="{Component}" v-model:newInstallation="newInstallation" @prevPage="prevPage($event)"
                 @nextPage="nextPage($event)" @complete="complete">
      <keep-alive>
        <component :is="Component"/>
      </keep-alive>
    </router-view>
  </div>
</template>

<script setup lang="ts">

import Steps from 'primevue/steps';
import { ref} from 'vue';
import { useRouter } from 'vue-router';
import { useToast } from 'vue-toastification';
import { useUserStore } from '@/store/user.store';
import { useI18n } from 'vue-i18n';
import { Installation } from '@/components/forms/create-installation/Installation';
import { useInstallationsStore } from '@/store/installations.store';
import { useContextStore } from "@/store/context.store";
import { computed } from '@vue/reactivity';

const router = useRouter();
const toast = useToast();
const userStore = useUserStore();
const contextStore = useContextStore();
const i18n = useI18n();

const currentPath = router.currentRoute.value.path;

const items = computed(() => [
  {
    label: i18n.t('FORMS.HEADERS.INSTALLATION_DATA'),
    to: currentPath + ""
  },
  {
    label: i18n.t('FORMS.HEADERS.ADDRESS'),
    to: currentPath + "/address_2",
  }
]);

const newInstallation =ref<Installation>({} as Installation);

const nextPage = (event: any) => {
  router.push(items.value[event.pageIndex + 1].to);
};
const prevPage = (event: any) => {
  router.push(items.value[event.pageIndex - 1].to);
};

const complete = () => {
    newInstallation.value.providerId = parseInt(userStore.providerId);
    newInstallation.value.accountId = contextStore.selectedCustomer.id;
    useInstallationsStore().createNewInstallation(newInstallation.value, onSuccess, onFail);
};

const onSuccess = () => {
  router.push('/customers/customer_tabs/installations');
  toast.success(i18n.t('GLOBALS.TOASTS.CREATED_INSTALLATION'));
};

const onFail = () => {
  toast.error(i18n.t('GLOBALS.TOASTS.ERROR_POST'));
};

</script>

<style scoped lang="scss">

</style>
