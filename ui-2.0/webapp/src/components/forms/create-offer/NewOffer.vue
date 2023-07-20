<template>
  <div>
    <div class="card">
      <Steps :model="items" :readonly="true"/>
    </div>
    <router-view v-slot="{Component}" v-model:newOffer="newOffer" :useSelectedOffer="useSelectedOffer" @prevPage="prevPage($event)"
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
import {
  EstimatedAnnualElectricityConsumption,
  OfferForm,

} from "@/components/forms/create-offer/OfferForm";
import {useContextStore} from "@/store/context.store";

import {
  Address,
  Conditions,
  Contact,
  Customer,
  DateFormat,
  DeliveryPeriod,
  Duration, Offer,
  OfferActivePeriod,
  OfferDetails,
  OfferDraft, OfferDraftConditions,
  OfferDraftDetails,
  OfferDraftPayload,
  OfferPayload,
  OfferPriceList,
  PhoneNumber,
  Repurchase,
  Seller,
  TariffGroupType
} from "@/models/billing/billing";
import {formatSendDate} from "@/utils/date-formatter";
import factoryApi from "@/api/factory.api";
import {computed} from '@vue/reactivity';
import {onBeforeMount} from "@vue/runtime-core";

const router = useRouter();
const toast = useToast();
const i18n = useI18n();

const props = defineProps({
  useSelectedOffer: {
    type: Boolean
  }
});
const currentPath = router.currentRoute.value.path;

const newOffer = ref<OfferForm>({} as OfferForm);

const items = computed(() => [
  {
    label: i18n.t('FORMS.HEADERS.GENERAL_INFO'),
    to: currentPath + ""
  },
  {
    label: i18n.t('FORMS.HEADERS.TECHNICAL_DATA'),
    to: currentPath + "/offer_step2",
  },
]);

const nextPage = (event: any) => {
  router.push(items.value[event.pageIndex + 1].to);
};
const prevPage = (event: any) => {
  router.push(items.value[event.pageIndex - 1].to);
};

onBeforeMount(() => {
  if(props.useSelectedOffer) {
    const selectedOffer = useContextStore().selectedOffer;
    factoryApi.draftOffersApi().getOfferDraftById(selectedOffer.payload.offerDetails.offerDraftId , true, null).then(response => {
      if (response.error === null && response.data) {
        newOffer.value.offerDraft = response.data;
        console.log(newOffer.value.offerDraft)
      }

    });
    // newOffer.value.offerDraft = {} as OfferDraft;
    // newOffer.value.offerDraft.payload = {} as OfferDraftPayload;
    // newOffer.value.offerDraft.payload.offerDetails = {} as OfferDraftDetails;
    //
    // newOffer.value.offerDraft.payload.offerDetails.title = selectedOffer.payload.offerDetails.title;
    // newOffer.value.offerDraft.payload.offerDetails.tariffGroup = selectedOffer.payload.offerDetails.tariffGroup as TariffGroupType;
    // newOffer.value.offerDraft.payload.offerDetails.agreementType = selectedOffer.payload.offerDetails.agreementType;
    // newOffer.value.offerDraft.payload.offerDetails.type = selectedOffer.payload.offerDetails.type;
    //
    // newOffer.value.offerDraft.payload.conditions = {} as OfferDraftConditions;
    // newOffer.value.offerDraft.payload.conditions.duration = selectedOffer.payload.conditions.duration as DeliveryPeriod;
    // newOffer.value.offerDraft.payload.conditions.billingPeriod = selectedOffer.payload.conditions.billingPeriod as Duration;
    // newOffer.value.offerDraft.payload.conditions.invoiceDueDate = selectedOffer.payload.conditions.invoiceDueDate;
    // newOffer.value.offerDraft.payload.priceList = selectedOffer.payload.priceList[0] as OfferPriceList;
    // newOffer.value.offerDraft.payload.repurchase = selectedOffer.payload.repurchase as Repurchase;

    newOffer.value.numberOfPPE = Number(selectedOffer.payload.conditions.numberOfPPE);
    newOffer.value.estimatedAnnualElectricityProduction = selectedOffer.payload.conditions.estimatedAnnualElectricityConsumption as EstimatedAnnualElectricityConsumption;
    newOffer.value.estimatedAnnualElectricityConsumption = selectedOffer.payload.conditions.estimatedAnnualElectricityConsumption as EstimatedAnnualElectricityConsumption;
  }
})

const complete = () => {
  if(props.useSelectedOffer){
    const selectedOffer = useContextStore().selectedOffer;
    let updatedOffer: Offer = {} as Offer;
    updatedOffer.payload = selectedOffer.payload;
    updatedOffer.header = selectedOffer.header;

    updatedOffer.payload.conditions.numberOfPPE = newOffer.value.numberOfPPE.toString();
    updatedOffer.payload.conditions.estimatedAnnualElectricityProduction = newOffer.value.estimatedAnnualElectricityProduction;
    updatedOffer.payload.conditions.estimatedAnnualElectricityConsumption = newOffer.value.estimatedAnnualElectricityConsumption;
    updatedOffer.payload.offerDetails.title = newOffer.value.offerDraft.payload.offerDetails.title;
    updatedOffer.payload.offerDetails.agreementType = newOffer.value.offerDraft.payload.offerDetails.agreementType;
    updatedOffer.payload.offerDetails.type = newOffer.value.offerDraft.payload.offerDetails.type;
    updatedOffer.payload.conditions.duration = newOffer.value.offerDraft.payload.conditions.duration;
    updatedOffer.payload.conditions.billingPeriod = newOffer.value.offerDraft.payload.conditions.billingPeriod;
    updatedOffer.payload.conditions.invoiceDueDate = newOffer.value.offerDraft.payload.conditions.invoiceDueDate;
    updatedOffer.payload.priceList[0] = newOffer.value.offerDraft.payload.priceList;
    updatedOffer.payload.repurchase = newOffer.value.offerDraft.payload.repurchase;
    if(selectedOffer.id)
      factoryApi.offersApi().updateOfferById(selectedOffer.id, updatedOffer, onSuccessEdit, onFailEdit);
  } else {
    let requestBody: any = {
      "Header": {
        "Version": "1.0",
        "Provider": useContextStore().loggedProvider?.name,
        "Content": {
          "Type": "offer",
          "Catg": "prosumer"
        }
      }
    };
    const customer = useContextStore().selectedCustomer;
    const provider = useContextStore().loggedProvider;
    console.log(customer);
    let offer: OfferPayload = {} as OfferPayload;
    offer.offerDetails = {} as OfferDetails;
    offer.offerDetails.title = newOffer.value.offerDraft.payload.offerDetails.title;
    offer.offerDetails.number = " ";
    offer.offerDetails.status = "DRAFT";
    offer.offerDetails.offerDraftId = newOffer.value.offerDraft.id;
    offer.offerDetails.tariffGroup = newOffer.value.offerDraft.payload.offerDetails.tariffGroup;
    offer.offerDetails.creationDate = formatSendDate(new Date().toString(), DateFormat.SEND_DATE_FORMAT);
    offer.offerDetails.agreementType = newOffer.value.offerDraft.payload.offerDetails.agreementType;
    offer.offerDetails.customerId = customer.id.toString();
    offer.offerDetails.type = newOffer.value.offerDraft.payload.offerDetails.type;

    offer.sellerDtls = {} as Seller;
    offer.sellerDtls.legalName = provider.name;
    offer.sellerDtls.displayName = provider.name;
    offer.sellerDtls.krs = provider.krs;
    offer.sellerDtls.nip = provider.nip;
    offer.sellerDtls.address = {} as Address;
    if (provider.apartmentNumber) {
      offer.sellerDtls.address.street = provider.street + " " + provider.buildingNumber + "/" + provider.apartmentNumber;
    } else {
      offer.sellerDtls.address.street = provider.street + " " + provider.buildingNumber;
    }
    offer.sellerDtls.address.postCode = provider.postalCode;
    offer.sellerDtls.address.city = provider.city;
    offer.sellerDtls.contact = {} as Contact;
    offer.sellerDtls.contact.address = {} as Address;
    if (provider.apartmentNumber) {
      offer.sellerDtls.contact.address.street = provider.street + " " + provider.buildingNumber + "/" + provider.apartmentNumber;
    } else {
      offer.sellerDtls.contact.address.street = provider.street + " " + provider.buildingNumber;
    }

    offer.sellerDtls.contact.address.postCode = provider.postalCode;
    offer.sellerDtls.contact.address.city = provider.city;
    offer.sellerDtls.contact.phoneNumbers = [{}] as PhoneNumber[];
    offer.sellerDtls.contact.phoneNumbers[0].number = provider.phoneNumber;
    offer.sellerDtls.contact.phoneNumbers[0].type = "fix";
    offer.sellerDtls.contact.email = provider.email;
    offer.sellerDtls.contact.www = provider.www;

    offer.customerDtls = {} as Customer;
    offer.customerDtls.customerId = customer.id.toString();
    offer.customerDtls.firstName = customer.firstName;
    offer.customerDtls.lastName = customer.lastName;
    offer.customerDtls.displayName = customer.firstName + " " + customer.lastName;
    offer.customerDtls.address = {} as Address;
    offer.customerDtls.pesel = customer.pesel;
    if (customer.apartmentNumber) {
      offer.customerDtls.address.street = customer.street + " " + customer.buildingNumber + "/" + customer.apartmentNumber;
    } else {
      offer.customerDtls.address.street = customer.street + " " + customer.buildingNumber;
    }
    offer.customerDtls.address.postCode = customer.postalCode;
    offer.customerDtls.address.city = customer.city;
    offer.customerDtls.contact = {} as Contact;
    offer.customerDtls.contact.address = {} as Address;
    if (customer.apartmentNumber) {
      offer.customerDtls.contact.address.street = customer.street + " " + customer.buildingNumber + "/" + customer.apartmentNumber;
    } else {
      offer.customerDtls.contact.address.street = customer.street + " " + customer.buildingNumber;
    }
    offer.customerDtls.contact.address.postCode = customer.postalCode;
    offer.customerDtls.contact.address.city = customer.city;
    offer.customerDtls.contact.phoneNumbers = [{}] as PhoneNumber[];
    offer.customerDtls.contact.phoneNumbers[0].number = customer.phone;
    offer.customerDtls.contact.phoneNumbers[0].type = "mobile";
    offer.customerDtls.contact.email = customer.email;
    offer.customerDtls.contact.www = "nie wiem co tu wpisac"; //TODO:

    offer.conditions = {} as Conditions;
    offer.conditions.duration = newOffer.value.offerDraft.payload.conditions.duration;
    offer.conditions.billingPeriod = newOffer.value.offerDraft.payload.conditions.billingPeriod;
    offer.conditions.invoiceDueDate = newOffer.value.offerDraft.payload.conditions.invoiceDueDate;
    offer.conditions.estimatedAnnualElectricityConsumption = newOffer.value.estimatedAnnualElectricityConsumption;
    offer.conditions.estimatedAnnualElectricityProduction = newOffer.value.estimatedAnnualElectricityProduction;
    offer.conditions.numberOfPPE = newOffer.value.numberOfPPE.toString();
    offer.conditions.offerActivePeriod = {} as OfferActivePeriod;
    offer.conditions.offerActivePeriod.startDate = formatSendDate(newOffer.value.startDate.toString(), DateFormat.SEND_DATE_FORMAT);
    offer.conditions.offerActivePeriod.endDate = formatSendDate(newOffer.value.endDate.toString(), DateFormat.SEND_DATE_FORMAT);

    offer.priceList = [{}] as OfferPriceList[];
    offer.priceList[0] = newOffer.value.offerDraft.payload.priceList;

    offer.repurchase = {} as Repurchase;
    offer.repurchase = newOffer.value.offerDraft.payload.repurchase;


    requestBody["Payload"] = offer;
    console.log(customer);
    console.log(requestBody);
    factoryApi.offersApi().saveNewOffer(requestBody, onSuccess, onFail);
  }
};

const onSuccess = () => {
  toast.success(i18n.t('GLOBALS.TOASTS.ADDED_PRICING'));
  router.push('/customers/customer_tabs/offers');
};

const onFail = () => {
  toast.error(i18n.t('GLOBALS.TOASTS.ERROR_POST'));
};

const onSuccessEdit = () => {
  toast.success("Pomyślnie zaktualizowano ofertę");
  router.push('/customers/customer_tabs/offers');
};

const onFailEdit = () => {
  toast.error("Edycja oferty nie powiodla się");
};
</script>

<style scoped lang="scss">

</style>
