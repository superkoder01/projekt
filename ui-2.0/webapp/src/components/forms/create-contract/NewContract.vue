<template>
  <div>
    <div class="card">
      <Steps :model="items" :readonly="true"/>
    </div>
    <router-view v-slot="{Component}" v-model:newContract="newContract" :useSelectedContract="useSelectedContract" @prevPage="prevPage($event)"
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
import { ContractForm, FormServiceAccessPoint, ServiceAccessPoint } from "@/components/forms/create-contract/Contract";
import { onBeforeMount } from "@vue/runtime-core";
import { useContextStore } from "@/store/context.store";
import {
  Contract,
  ContractConditions, ContractDetails,
  ContractPayload,
  DateFormat, Duration, Offer,
} from "@/models/billing/billing";
import { formatSendDate } from "@/utils/date-formatter";
import factoryApi from "@/api/factory.api";
import { computed } from '@vue/reactivity';


const router = useRouter();
const toast = useToast();
const i18n = useI18n();

const props = defineProps({
  useSelectedContract: {
    type: Boolean
  }
});

const currentPath = router.currentRoute.value.path;

const newContract = ref<ContractForm>({} as ContractForm);

onBeforeMount(() => {
  if(!props.useSelectedContract) {
    newContract.value.offer = useContextStore().selectedOffer;
  }
  if(props.useSelectedContract) {
    const selectedContract = useContextStore().getSelectedContract;
    newContract.value.tpaParameter = selectedContract.payload.contractDetails.tpaParameter;
    newContract.value.clientType = selectedContract.payload.contractDetails.clientType;
    newContract.value.serviceAccessPoints = selectedContract.payload.serviceAccessPoints;
    newContract.value.startDate = selectedContract.payload.conditions.startDate;
    newContract.value.bankAccountNumber = selectedContract.payload.sellerDtls.bankAccountNumber;
    factoryApi.offersApi().getOfferById(selectedContract.payload.contractDetails.offerId, true, null).then(response => {
      if (response.error === null && response.data) {
        newContract.value.offer = response.data;
      }
    });

  }
});

const items = computed(() => [
  {
    label: i18n.t('FORMS.HEADERS.GENERAL_INFO'),
    to: currentPath + ""
  },
  {
    label: i18n.t('FORMS.HEADERS.OSD_DATA'),
    to: currentPath + "/contract_step2",
  },
  {
    label: i18n.t('FORMS.HEADERS.SAP_DATA'),
    to: currentPath + "/contract_step3",
  },
  {
    label: i18n.t('FORMS.HEADERS.EXTRA'),
    to: currentPath + "/contract_step4",
  },
]);

const nextPage = (event: any) => {
  router.push(items.value[event.pageIndex + 1].to);
};
const prevPage = (event: any) => {
  router.push(items.value[event.pageIndex - 1].to);
};

const complete = () => {
  const provider = useContextStore().loggedProvider;

  let requestBody: any = {
    "Header": {
      "Version": "1.0",
      "Provider": useContextStore().loggedProvider?.name,
      "Content": {
        "Type": "contract",
        "Catg": "prosumer"
      }
    }
  };
  let contract: ContractPayload = {} as ContractPayload;

  if(props.useSelectedContract) {
    const selectedContract = useContextStore().getSelectedContract;
    let updatedContract: Contract = {} as Contract;
    updatedContract.payload = selectedContract.payload;
    updatedContract.header = selectedContract.header;

    updatedContract.payload.contractDetails.tpaParameter = newContract.value.tpaParameter;
    updatedContract.payload.contractDetails.clientType = newContract.value.clientType;
    updatedContract.payload.serviceAccessPoints = newContract.value.serviceAccessPoints;
    updatedContract.payload.sellerDtls.bankAccountNumber = newContract.value.bankAccountNumber;
    updatedContract.payload.conditions.startDate = formatSendDate(newContract.value.startDate, DateFormat.SEND_DATE_FORMAT);
    const startDate = new Date(newContract.value.startDate);
    updatedContract.payload.conditions.endDate = formatSendDate(new Date(startDate.getFullYear(), startDate.getMonth()+Number(newContract.value.offer.payload.conditions.duration.number), startDate.getDate()).toString(), DateFormat.SEND_DATE_FORMAT);
    console.log(selectedContract);
    if(selectedContract.id)
      factoryApi.contractsApi().updateContractById(selectedContract.id, updatedContract, onSuccess, onFail);
  } else {
    contract.contractDetails = {} as ContractDetails;
    contract.contractDetails.title = "";
    contract.contractDetails.type = newContract.value.offer.payload.offerDetails.type;
    contract.contractDetails.number = "";
    contract.contractDetails.offerId = newContract.value.offer.id;
    contract.contractDetails.creationDate = formatSendDate(new Date().toString(), DateFormat.SEND_DATE_FORMAT);
    contract.contractDetails.state = "DRAFT";
    contract.contractDetails.customerId = newContract.value.offer.payload.offerDetails.customerId;
    contract.contractDetails.tariffGroup = newContract.value.serviceAccessPoints[0].tariffGroup;
    contract.contractDetails.agreementType = newContract.value.offer.payload.offerDetails.agreementType;
    contract.contractDetails.tpaParameter = newContract.value.tpaParameter;
    contract.contractDetails.clientType = newContract.value.clientType;


    contract.sellerDtls = newContract.value.offer.payload.sellerDtls;
    contract.sellerDtls.bankAccountNumber = newContract.value.bankAccountNumber;

    contract.customerDtls = newContract.value.offer.payload.customerDtls;

    contract.conditions = {} as ContractConditions;
    contract.conditions.signatureDate = "";
    contract.conditions.startDate = formatSendDate(newContract.value.startDate, DateFormat.SEND_DATE_FORMAT);
    const startDate = new Date(newContract.value.startDate);
    contract.conditions.endDate = formatSendDate(new Date(startDate.getFullYear(), startDate.getMonth()+Number(newContract.value.offer.payload.conditions.duration.number), startDate.getDate()).toString(), DateFormat.SEND_DATE_FORMAT);
    contract.conditions.duration = {} as Duration; // Nie wiem czy mam to wypełniać czy backend to wypelnia
    contract.conditions.billingPeriod = newContract.value.offer.payload.conditions.billingPeriod;
    contract.conditions.invoiceDueDate = newContract.value.offer.payload.conditions.invoiceDueDate;
    contract.conditions.estimatedAnnualElectricityConsumption = newContract.value.offer.payload.conditions.estimatedAnnualElectricityConsumption;

    contract.serviceAccessPoints = [{}] as FormServiceAccessPoint[];
    contract.serviceAccessPoints = newContract.value.serviceAccessPoints;

    contract.priceList = newContract.value.offer.payload.priceList;

    const commercialFee: any = newContract.value.offer.payload.priceList[0].commercialFee.find(obj => {
      if(obj.to === undefined) return true;
      return newContract.value.serviceAccessPoints[0].contractedPower.amount >= Number(obj.from) && newContract.value.serviceAccessPoints[0].contractedPower.amount <= Number(obj.to);
    });
    console.log(commercialFee);
    contract.priceList[0].commercialFee = commercialFee.price;
    contract.repurchase = newContract.value.offer.payload.repurchase;
    requestBody["Payload"] = contract;
    factoryApi.contractsApi().saveNewContract(requestBody, onSuccess, onFail);
  }



};

const onSuccess = () => {
  toast.success(i18n.t('GLOBALS.TOASTS.ADDED_PRICING'));
  router.push('/customers/customer_tabs/contracts');
};

const onFail = () => {
  toast.error(i18n.t('GLOBALS.TOASTS.ERROR_POST'));
};
</script>

<style scoped lang="scss">

</style>
