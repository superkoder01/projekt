<template>

  <div class="card m-3">
    <div class="card-body" v-if="newContract.offer">

      <div>
        <h4> Umowa {{ newContract.offer.payload.offerDetails.type }} dla oferty
          {{ newContract.offer.payload.offerDetails.title }} o numerze
          {{ newContract.offer.payload.offerDetails.number }} </h4>
      </div>
      <Accordion class="acc-margin">
        <AccordionTab :header="`${$t('FORMS.HEADERS.CUSTOMER_INFO')}: ` + newContract.offer.payload.customerDtls.firstName + ' ' + newContract.offer.payload.customerDtls.lastName">
          <h2 style="margin-top: 20px">{{ $t("FORMS.HEADERS.CUSTOMER_INFO") }}</h2>
          <div class="detailed-data" style="margin-top: 20px">
            <span><span><Icon name="User"/></span>
            <span><div class="label-small">{{ $t("FORMS.PLACEHOLDERS.FULL_NAME") }}:</div>{{ newContract.offer.payload.customerDtls.firstName }} {{ newContract.offer.payload.customerDtls.lastName }}</span></span>

            <span><span><Icon name="Contact"/></span>
            <span><div class="label-small">{{ $t("FORMS.PLACEHOLDERS.PESEL") }}:</div>{{ newContract.offer.payload.customerDtls.pesel }}</span></span>

            <span><span><Icon name="None"/></span>
            <span><div class="label-small">{{ $t("FORMS.PLACEHOLDERS.AGREEMENT_TYPE") }}:</div>{{ newContract.offer.payload.offerDetails.agreementType }}</span></span>
          </div>

          <h2 style="margin-top: 20px">{{ $t("FORMS.HEADERS.ADDRESS") }}</h2>
          <div class="detailed-data" style="margin-top: 20px">
            <span><span><Icon name="Home"/></span>
            <span><div class="label-small">{{ $t("FORMS.PLACEHOLDERS.CITY") }}:</div>{{ newContract.offer.payload.customerDtls.address.city }}</span></span>

            <span><span><Icon name="None"/></span>
            <span><div class="label-small">{{ $t("FORMS.PLACEHOLDERS.POSTAL_CODE") }}</div>{{ newContract.offer.payload.customerDtls.address.postCode }}</span></span>

            <span><span><Icon name="None"/></span>
            <span><div class="label-small">{{ $t("FORMS.PLACEHOLDERS.STREET") }}</div>{{ newContract.offer.payload.customerDtls.address.street }}</span></span>
          </div>
          <h2 style="margin-top: 20px">{{ $t("FORMS.HEADERS.ADDRESS_FOR_CORRESPONDENCE") }}</h2>
          <div class="detailed-data" style="margin-top: 20px">
            <span><span><Icon name="Home"/></span>
            <span><div class="label-small">{{ $t("FORMS.PLACEHOLDERS.CITY") }}:</div>{{ newContract.offer.payload.customerDtls.contact.address.city }}</span></span>

            <span><span><Icon name="None"/></span>
            <span><div class="label-small">{{ $t("FORMS.PLACEHOLDERS.POSTAL_CODE") }}</div>{{ newContract.offer.payload.customerDtls.contact.address.postCode }}</span></span>

            <span><span><Icon name="None"/></span>
            <span><div class="label-small">{{ $t("FORMS.PLACEHOLDERS.STREET") }}</div>{{ newContract.offer.payload.customerDtls.contact.address.street }}</span></span>
          </div>
        </AccordionTab>
      </Accordion>

      <Form @submit="nextPage" :validation-schema="schema" v-slot="{errors}">
        <div class="p-fluid formgrid grid ">

          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="selectedContractType" required name="selectedContractType" as="select" class="form-control" :class="{'is-invalid': errors.selectedContractType}">
              <option v-for="contractType in contractTypeList" :key="contractType" :value="contractType">{{ contractType }}</option>
            </Field>
            <span>{{ $t("FORMS.PLACEHOLDERS.CONTRACT_TYPE") }}</span>
            <div class="invalid-feedback">{{ errors.selectedContractType ? $t(errors.selectedContractType) : '' }}</div>
            <TooltipComponent :tooltip-text="$t('TOOLTIP.CONTRACT_TYPE')"/>
          </div>

          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="selectedTPA" required name="selectedTPA" as="select" class="form-control" :class="{'is-invalid': errors.selectedTPA}">
              <option v-for="tpa in tpaList" :key="tpa" :value="tpa">{{ tpa }}</option>
            </Field>
            <span>{{ $t("FORMS.PLACEHOLDERS.TPA_PARAMETER") }}</span>
            <div class="invalid-feedback">{{ errors.selectedTPA ? $t(errors.selectedTPA) : '' }}</div>
          </div>

          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="currentSeller" name="currentSeller" placeholder=" " type="text" class="form-control" :class="{'is-invalid': errors.currentSeller}"/>
            <span>{{ $t("FORMS.PLACEHOLDERS.CURRENT_SELLER") }}</span>
            <div class="invalid-feedback">{{ errors.currentSeller ? $t(errors.currentSeller) : '' }}</div>
          </div>

        </div>
        <div class="grid grid-nogutter justify-content-between">
          <Button :label="$t('GLOBALS.BUTTONS.NEXT')" type="submit" icon="pi pi-angle-right" icon-pos="right" v-tooltip="'tesggdfgsdgsdfgdsfg'"></Button>
        </div>
      </Form>
    </div>
  </div>

</template>

<script setup lang="ts">
import {defineEmits, PropType, ref, watch} from 'vue';
import {Field, Form} from 'vee-validate';
import * as Yup from 'yup';
import {onBeforeMount} from "@vue/runtime-core";
import {useDraftOffersStore} from "@/store/draftOffers.store";
import {ContractForm, CurrentSeller, FormServiceAccessPoint} from "@/components/forms/create-contract/Contract";
import TooltipComponent from '@/components/features/TooltipComponent.vue';
import {useContextStore} from "@/store/context.store";

const emit = defineEmits(['nextPage', 'prevPage', 'update:newContract']);

const props = defineProps({

  newContract: {
    type: Object as PropType<ContractForm>,
    required: true
  },
  useSelectedContract: {
    type: Boolean
  }
});
const offerDraftStore = useDraftOffersStore();
const contractTypeList = ref(['Odbiorca', 'Wytwórca', 'Prosument', 'Odbiorca z mikroinstalacją']);
const selectedContractType = ref();
const tpaList = ref(['nowe przyłącze', 'pierwsza zmiana sprzedawcy', 'kolejna zmiana sprzedawcy']);
const currentSeller = ref();
const selectedTPA = ref();


onBeforeMount(() => {
  offerDraftStore.fetchDraftOffers(true, null);

  if(props.useSelectedContract) {
    selectedContractType.value = props.newContract.clientType;
    selectedTPA.value = props.newContract.tpaParameter;
  }
});

watch(() => props.newContract.offer, () => {
  currentSeller.value = props.newContract.serviceAccessPoints[0].currentSeller.name;
});

const schema = Yup.object().shape({
  selectedContractType: Yup.string()
    .required("FORMS.VALIDATION_ERRORS.CONTRACT_TYPE"),
  selectedTPA: Yup.string()
    .required("FORMS.VALIDATION_ERRORS.TPA_PARAMETER"),
  currentSeller: Yup.string()
    .required("FORMS.VALIDATION_ERRORS.CURRENT_SELLER")
});


const nextPage = () => {
  let updatedNewContract: ContractForm = props.newContract;
  if(props.useSelectedContract) {
    updatedNewContract.tpaParameter = selectedTPA.value;
    updatedNewContract.clientType = selectedContractType.value;
    updatedNewContract.serviceAccessPoints[0].currentSeller.name = currentSeller.value;

  } else {
    updatedNewContract.serviceAccessPoints = [{}] as FormServiceAccessPoint[];
    updatedNewContract.serviceAccessPoints[0] = {} as FormServiceAccessPoint;
    updatedNewContract.serviceAccessPoints[0].currentSeller = {} as CurrentSeller;
    updatedNewContract.serviceAccessPoints[0].currentSeller.name = currentSeller.value;
    updatedNewContract.tpaParameter = selectedTPA.value;
    updatedNewContract.clientType = selectedContractType.value;

  }
  emit('nextPage', {pageIndex: 0});
  emit('update:newContract', updatedNewContract);
};

</script>
<style scoped lang="scss">
.acc-margin {
  margin-bottom: 20px;
}

</style>
