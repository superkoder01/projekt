<template>
  <div>

    <div class="card no-border sticky">
      <Steps :model="items" :readonly="true" />
    </div>

    <router-view v-slot="{Component}" v-model:newTariffGroup="newTariffGroup" @prevPage="prevPage($event)" @nextPage="nextPage($event)" @complete="complete">
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
import {TariffGroup} from "@/components/forms/tariff-group/TariffGroup";
import { useToast } from 'vue-toastification';
import { useI18n } from 'vue-i18n';
import { computed } from '@vue/reactivity';
import { useTariffGroupStore } from '@/store/tariffGroup.store';

const router = useRouter();
const toast = useToast();
const i18n = useI18n();

const currentPath = router.currentRoute.value.path;
const tariffGroupStore = useTariffGroupStore();

const items = computed(()=>[
  {
    label: i18n.t('FORMS.HEADERS.GENERAL_INFO'),
    to: currentPath + ""
  },
  {
    label: i18n.t('FORMS.HEADERS.FEES'),
    to: currentPath + "/fees_2",
  }
]);
const newTariffGroup = ref<TariffGroup>({distributionNetworkOperatorID:-1, name:'', tariffGroupLabelName:'', startDate: new Date(), endDate: new Date(), fees:[]});

const nextPage = (event: any) => {
  router.push(items.value[event.pageIndex + 1].to);
};
const prevPage = (event: any) => {
  router.push(items.value[event.pageIndex - 1].to);
};

const complete = () => {
  tariffGroupStore.createNewTariffGroup(newTariffGroup.value, onSuccess, onFail);
};

const onSuccess = () => {
  toast.success(i18n.t('GLOBALS.TOASTS.ADDED_TARRIF_GROUP'));
  router.push('/product_catalog/tariff_pricing');
};

const onFail = () => {
  toast.error(i18n.t('GLOBALS.TOASTS.ERROR_POST'));
};
</script>

<style scoped lang="scss">

</style>
