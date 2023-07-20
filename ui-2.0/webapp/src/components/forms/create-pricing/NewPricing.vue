<template>
  <div>
    <div class="card">
      <Steps :model="items" :readonly="true"/>
    </div>
    <router-view v-slot="{Component}" v-model:newPricing="newPricing" @prevPage="prevPage($event)"
                 @nextPage="nextPage($event)" @complete="complete">
      <keep-alive>
        <component :is="Component"/>
      </keep-alive>
    </router-view>
  </div>
</template>

<script setup lang="ts">

import Steps from 'primevue/steps';
import {ref} from 'vue';
import {useRouter} from 'vue-router';
import {useToast} from 'vue-toastification';
import {useI18n} from 'vue-i18n';
import {FormPricing} from "@/components/forms/create-pricing/Pricing";
import {useContextStore} from "@/store/context.store";
import {CommercialFeeUnitType, DateFormat, Pricing, Repurchase} from "@/models/billing/billing";
import factoryApi from "@/api/factory.api";
import {ServiceTypeEnum} from "@/models/billing/enum/service-type.enum";
import { computed } from '@vue/reactivity';

const router = useRouter();
const toast = useToast();
const i18n = useI18n();

const currentPath = router.currentRoute.value.path;

const newPricing = ref<FormPricing>({} as FormPricing);

const items = computed(() => [
  {
    label: i18n.t('FORMS.HEADERS.GENERAL_INFO'),
    to: currentPath + ""
  },
  {
    label: i18n.t('FORMS.HEADERS.FEES'),
    to: currentPath + "/pricing_2",
  },
]);

const nextPage = (event: any) => {
  router.push(items.value[event.pageIndex + 1].to);
};
const prevPage = (event: any) => {
  router.push(items.value[event.pageIndex - 1].to);
};

const complete = () => {

  let requestBody: any = {
    "Header": {
      "Version": "1.0",
      "Provider": useContextStore().loggedProvider?.name,
      "Content": {
        "Type": "pricing",
        "Catg": ""
      }
    }
  };

  if (newPricing.value.type === ServiceTypeEnum.SALE) {
    let newSalePricing: Pricing = {} as Pricing;
    newSalePricing.name = newPricing.value.name;
    newSalePricing.type = newPricing.value.type;
    //newSalePricing.StartDate = formatSendDate(newPricing.value.startDate.toString(), DateFormat.SEND_DATE_FORMAT);
    //newSalePricing.EndDate =  formatSendDate(newPricing.value.endDate.toString(), DateFormat.SEND_DATE_FORMAT);
    newSalePricing.osd = "ENID";//TODO: skąd brać dane do tego?
    newSalePricing.tariffGroup = newPricing.value.tariffGroup;
    newSalePricing.zones = [{
      id: "1",
      name: "całodobowa",
      unit: "MWh",
      cost: newPricing.value.zones,
      currency: "pln"
    }];
    newSalePricing.commercialFee = newPricing.value.commercialFee;
    requestBody["Payload"] = newSalePricing;
    factoryApi.pricingApi().saveNewPricing(requestBody, onSuccess, onFail);
  } else if (newPricing.value.type === ServiceTypeEnum.REPURCHASE) {
    let newFixedPricing: Repurchase = {} as Repurchase;
    newFixedPricing.name = newPricing.value.name;
    newFixedPricing.type = newPricing.value.type;
    newFixedPricing.id = "";
    newFixedPricing.price = {unit: CommercialFeeUnitType.KWP, cost: newPricing.value.fixedPrice, currency: "pln"};
    requestBody["Payload"] = newFixedPricing;
    factoryApi.pricingApi().saveNewPricing(requestBody, onSuccess, onFail);
  } else if (newPricing.value.type === ServiceTypeEnum.REPURCHASE_RDN) {
    let newRdnPricing: Pricing = {} as Pricing;
    newRdnPricing.name = newPricing.value.name;
    newRdnPricing.type = newPricing.value.type;
    newRdnPricing.id = "";
    newRdnPricing.price = {};
    requestBody["Payload"] = newRdnPricing;
    factoryApi.pricingApi().saveNewPricing(requestBody, onSuccess, onFail);
  }
  console.log(requestBody);
};

const onSuccess = () => {
  toast.success(i18n.t('GLOBALS.TOASTS.ADDED_PRICING'));
  router.push('/product_catalog/pricing');
};

const onFail = () => {
  toast.error(i18n.t('GLOBALS.TOASTS.ERROR_POST'));
};
</script>

<style scoped lang="scss">

</style>
