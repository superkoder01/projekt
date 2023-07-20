<template>
  <div class="card m-3">
    <h5 class="card-header">{{ $t('FORMS.HEADERS.ABOUT_OFFER') }} </h5>
    <div class="card-body">
      <Form @submit="nextPage" :validation-schema="schema" v-slot="{errors}" :initial-values="{startDate:startDate}">
        <div class="p-fluid formgrid grid ">
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="title" name="title" placeholder=" " type="text" class="form-control" :class="{'is-invalid': errors.title}"></Field>
            <span>{{ $t('FORMS.PLACEHOLDERS.NAME')}}</span>
            <div class="invalid-feedback">{{errors.title ? $t(errors.title) : '' }}</div>
          </div>
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="selectedAgreementType" required name="agreementType" as="select" class="form-control" :class="{'is-invalid': errors.agreementType}" >
              <option v-for="agreementType in agreementTypeList" :key="agreementType" :value="agreementType">{{agreementType}}</option>
            </Field>
            <span>{{ $t('FORMS.PLACEHOLDERS.AGREEMENT_TYPE') }}</span>
            <div class="invalid-feedback">{{errors.agreementType ? $t(errors.agreementType): '' }}</div>
          </div>
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="selectedType" required name="type" as="select" placeholder=" "  class="form-control" :class="{'is-invalid': errors.type}" :disabled="selectedAgreementType==='B2B'">
              <option v-for="type in typeList" :key="type" :value="type" >{{type}}</option>
            </Field>
            <span>{{ $t('FORMS.PLACEHOLDERS.TYPE')}}</span>
            <div class="invalid-feedback">{{errors.type ? $t(errors.type): '' }}</div>
          </div>
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="selectedTariff" required name="tariff" as="select" class="form-control"
                   :class="{'is-invalid': errors.tariff}">
              <option v-for="tariff in tariffList" :key="tariff" :value="tariff.name">{{ tariff.name }}</option>
            </Field>
            <span>{{ $t('FORMS.PLACEHOLDERS.TARIFF_GROUP') }}</span>
            <div class="invalid-feedback">{{errors.tariff ? $t(errors.tariff): '' }}</div>
          </div>


          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="startDate" name="startDate" type="date" class="form-control"
                   :class="{'is-invalid': errors.startDate}"></Field>
            <span>{{ $t('FORMS.PLACEHOLDERS.SALE_PERIOD_FROM') }}</span>
            <div class="invalid-feedback">{{ errors.startDate }}</div>
          </div>
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="endDate" name="endDate" type="date" class="form-control"
                   :class="{'is-invalid': errors.endDate}"></Field>
            <span>{{ $t('FORMS.PLACEHOLDERS.SALE_PERIOD_TO') }}</span>
            <div class="invalid-feedback">{{ errors.endDate }}</div>
          </div>

        </div>
        <div class="grid grid-nogutter justify-content-between">
          <Button :label="$t('GLOBALS.BUTTONS.BACK')" @click="goBack" type="submit" icon="pi pi-angle-left" icon-pos="left"></Button>
          <Button :label="$t('GLOBALS.BUTTONS.NEXT')" type="submit" icon="pi pi-angle-right" icon-pos="right"></Button>
        </div>
      </Form>
    </div>
  </div>
</template>

<script setup lang="ts">
import { ref, defineEmits, PropType, onMounted, watch } from "vue";
import {TariffGroup} from "@/components/forms/tariff-group/TariffGroup";
import { Form, Field } from 'vee-validate';
import * as Yup from 'yup';
import {DataHolder} from "@/models/data-holder";
import {onBeforeMount} from "@vue/runtime-core";
import {FormPricing} from "@/components/forms/create-pricing/Pricing";
import {FormOfferDraft} from "@/components/forms/create-draft-offer/OfferDraft";
import router from "@/router";
import { TariffGroupType } from "@/models/billing/billing";
import { ServiceTypeEnum } from "@/models/billing/enum/service-type.enum";
import moment from "moment";

const emit = defineEmits(['nextPage', 'prevPage', 'update:updatedOfferDraft']);

const props = defineProps({

  newOfferDraft: {
    type:  Object as PropType<FormOfferDraft>,
    required: true
  }
});

onBeforeMount(()=> {
  console.log();
});

const title = ref();
const selectedType = ref();
const typeList = ref(['Kompleksowa', 'Rozdzielona']);
const selectedAgreementType = ref();
const agreementTypeList = ref(['B2C', 'B2B']);
const tariffList = ref([{id: 1, name: TariffGroupType.G11}]);
const selectedTariff = ref();
const startDate = ref<string>((moment(new Date())).format('YYYY-MM-DD'));
const endDate = ref<string>((moment(new Date())).format('YYYY-MM-DD'));

const schema = Yup.object().shape({
  title: Yup.string().required('FORMS.VALIDATION_ERRORS.REQUIRED'),
  tariff: Yup.string().required('FORMS.VALIDATION_ERRORS.REQUIRED'),
  type: Yup.string().required('FORMS.VALIDATION_ERRORS.REQUIRED'),
  agreementType: Yup.string().required('FORMS.VALIDATION_ERRORS.REQUIRED'),

});

watch(() => selectedAgreementType.value, () => {
  selectedType.value = selectedAgreementType.value === 'B2B' ? 'Rozdzielona' : 'Kompleksowa';
});

const nextPage = () => {

  let updatedOfferDraft: FormOfferDraft = props.newOfferDraft;
  updatedOfferDraft.title = title.value;
  updatedOfferDraft.type = selectedType.value;
  updatedOfferDraft.agreementType = selectedAgreementType.value;
  updatedOfferDraft.tariffGroup = selectedTariff.value;
  updatedOfferDraft.startDate = startDate.value;
  updatedOfferDraft.endDate = endDate.value;


  emit('nextPage', {pageIndex: 0});
  emit('update:updatedOfferDraft', updatedOfferDraft);
};

const goBack = () => {
  router.push("/product_catalog/offer_draft");
}

</script>
<style scoped lang="scss">

</style>
