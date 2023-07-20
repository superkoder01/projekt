<template>
  <div>
    <div class="card">
      <Steps :model="items" :readonly="true"/>
    </div>
    <router-view v-slot="{Component}" v-model:newOfferDraft="newOfferDraftForm" @prevPage="prevPage($event)"
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
import {FormOfferDraft} from "@/components/forms/create-draft-offer/OfferDraft";
import {useContextStore} from "@/store/context.store";
import {
  DateFormat,
  OfferDraft,
  OfferDraftConditions,
  OfferDraftDetails,
  OfferDraftPayload,
  OfferPriceList,
  TariffGroupType
} from "@/models/billing/billing";
import {formatSendDate} from "@/utils/date-formatter";
import factoryApi from "@/api/factory.api";
import moment from "moment";
import { computed } from '@vue/reactivity';

const router = useRouter();
const toast = useToast();
const i18n = useI18n();

const currentPath = router.currentRoute.value.path;

const newOfferDraftForm = ref<FormOfferDraft>({} as FormOfferDraft);

const items = computed(() => [
  {
    label: i18n.t('FORMS.HEADERS.ABOUT_OFFER'),
    to: currentPath + ""
  },
  {
    label: i18n.t('FORMS.HEADERS.PRICING'),
    to: currentPath + "/step2",
  },
  {
    label: i18n.t('FORMS.HEADERS.CONDITIONS'),
    to: currentPath + "/step3",
  }
]);


const nextPage = (event: any) => {
  router.push(items.value[event.pageIndex + 1].to);
};
const prevPage = (event: any) => {
  router.push(items.value[event.pageIndex - 1].to);
};

const startDate = ref<string>((moment(new Date())).format('YYYY-MM-DD'));
const endDate = ref<string>((moment(new Date())).format('YYYY-MM-DD'));

const complete = () => {
  const temp = newOfferDraftForm.value;
  console.log("temp:" + JSON.stringify(temp));

  let offerDraftPayload: OfferDraftPayload = {} as OfferDraftPayload;
  offerDraftPayload.offerDetails = {} as OfferDraftDetails;
  offerDraftPayload.offerDetails.title = temp.title;
  offerDraftPayload.offerDetails.agreementType = temp.agreementType;
  offerDraftPayload.offerDetails.tariffGroup = newOfferDraftForm.value.tariffGroup;
  offerDraftPayload.offerDetails.type = temp.type;

  offerDraftPayload.conditions = {} as OfferDraftConditions;
  offerDraftPayload.conditions.billingPeriod = temp.billingPeriod;
  offerDraftPayload.conditions.duration = temp.duration;
  offerDraftPayload.conditions.invoiceDueDate = temp.invoiceDueDate.toString();
  offerDraftPayload.conditions.startDate =  formatSendDate(temp.startDate,DateFormat.SEND_DATE_FORMAT) ;
  offerDraftPayload.conditions.endDate = formatSendDate(temp.endDate, DateFormat.SEND_DATE_FORMAT);

  if (temp.price !== undefined) {
    offerDraftPayload.priceList = {} as OfferPriceList;
    offerDraftPayload.priceList.name = temp.price.name;
    offerDraftPayload.priceList.type = temp.price.type;
    //offerDraftPayload.PriceList.StartDate = formatDate(temp.price.StartDate,DateFormat.SEND_DATE_FORMAT); //TODO
    offerDraftPayload.priceList.osd = "ENID";
    offerDraftPayload.priceList.tariffGroup = TariffGroupType.G11;

    if (temp.price.commercialFee !== undefined) {
      offerDraftPayload.priceList.commercialFee = temp.price.commercialFee;
      // offerDraftPayload.priceList.commercialFee = new Array<OfferCommercialFee>();
      // temp.price.commercialFee.forEach((fee: CommercialFee) => {
      //   offerDraftPayload.priceList.commercialFee.push({
      //     price: fee.price,
      //     from: fee.from,
      //     to: fee.to,
      //     unit: fee.unit
      //   });
      // });
    }
    offerDraftPayload.priceList.zones = temp.price.zones;
  }

  if (temp.repurchase !== undefined) {
    offerDraftPayload.repurchase = {id: temp.repurchase.id, name: temp.repurchase.name, type:temp.repurchase.type, price:temp.repurchase.price};
  }
  let requestBody: OfferDraft = {
    "header": {
      "version": "1.0",
      "provider": useContextStore().loggedProvider?.name,
      "content": {
        "type": "draft_offer",
        "category": "prosumer"
      }
    },
    "payload" : offerDraftPayload
  };
  console.log(JSON.stringify(requestBody));
  factoryApi.draftOffersApi().saveNewOfferDraft(requestBody, onSuccess, onFail)
};

const onSuccess = () => {
  toast.success(i18n.t('GLOBALS.TOASTS.ADDED_DRAFT_OFFER'));
  router.push('/product_catalog/offer_draft');
};

const onFail = () => {
  toast.error(i18n.t('GLOBALS.TOASTS.ERROR_POST'));
};
</script>

<style scoped lang="scss">

</style>
