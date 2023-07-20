<template>
  <div class="card m-3">
    <h5 class="card-header">{{ $t('FORMS.HEADERS.CONDITIONS') }} </h5>
    <div class="card-body">
      <Form @submit="complete" :validation-schema="schema" v-slot="{errors}">
        <div class="p-fluid formgrid grid ">

          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="deliveryPeriodNumber" name="deliveryPeriodNumber" placeholder=" " type="number" class="form-control" :class="{'is-invalid': errors.deliveryPeriodNumber}"></Field>
            <span>{{ $t('FORMS.PLACEHOLDERS.MONTH_CONTRACT_DURATION') }}</span>
            <div class="invalid-feedback">{{errors.deliveryPeriodNumber ? $t(errors.deliveryPeriodNumber) : '' }}</div>
          </div>

          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="billingPeriodNumber" required name="billingPeriodNumber" as="select" class="form-control"
                   :class="{'is-invalid': errors.billingPeriodNumber}">
              <option v-for="period in billingPeriodNumberList" :key="period" :value="period">{{ $t("FORMS.PLACEHOLDERS."+period+"_MONTHS") }} </option>
            </Field>
            <span>{{ $t('FORMS.PLACEHOLDERS.MONTH_BILLING_PERIOD_DURATION') }}</span>
            <div class="invalid-feedback">{{errors.billingPeriodNumber ? $t(errors.billingPeriodNumber) : '' }}</div>

          </div>

          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="invoiceDueDate" required name="invoiceDueDate" as="select" class="form-control"
                   :class="{'is-invalid': errors.invoiceDueDate}">
              <option v-for="period in invoiceDueDateList" :key="period" :value="period">{{ $t("FORMS.PLACEHOLDERS."+period+"_DAYS") }} </option>
            </Field>
            <span>{{$t("FORMS.PLACEHOLDERS.INVOICE_DUE_DATE")}}</span>
            <div class="invalid-feedback">{{ errors.invoiceDueDate ? $t(errors.invoiceDueDate) : '' }}</div>

          </div>



        </div>
        <div class="grid grid-nogutter justify-content-between">
          <Button :label="$t('GLOBALS.BUTTONS.BACK')" @click="goBack" type="submit" icon="pi pi-angle-left" icon-pos="left"></Button>
          <Button :label="$t('GLOBALS.BUTTONS.CREATE')" type="submit" icon="pi pi-angle-right"></Button>
        </div>
      </Form>
    </div>
  </div>
</template>

<script setup lang="ts">
import {ref, defineEmits, PropType} from 'vue';
import { Form, Field } from 'vee-validate';
import * as Yup from 'yup';
import {onBeforeMount} from "@vue/runtime-core";
import {FormOfferDraft} from "@/components/forms/create-draft-offer/OfferDraft";
import { CalendarUnitType, DateFormat, toCalendarUnitTypeEnumKey } from "@/models/billing/billing";
import moment from "moment";
import { ServiceTypeEnum } from "@/models/billing/enum/service-type.enum";
import { formatSendDate } from "@/utils/date-formatter";

const emit = defineEmits(['complete', 'prevPage', 'update:updatedOfferDraft']);

const props = defineProps({

  newOfferDraft: {
    type: Object as PropType<FormOfferDraft>,
    required: true
  }
});

onBeforeMount(()=> {
  console.log();
});

const deliveryPeriodNumber = ref();
const billingPeriodNumber = ref();
const invoiceDueDate = ref();
const billingPeriodNumberList = ref(['1', '3', '6', '12']);
const invoiceDueDateList = ref(['7', '14', '30']);

const schema = Yup.object().shape({
  deliveryPeriodNumber: Yup.string().required('FORMS.VALIDATION_ERRORS.REQUIRED'),
  billingPeriodNumber: Yup.string().required('FORMS.VALIDATION_ERRORS.REQUIRED'),
  invoiceDueDate: Yup.string().required('FORMS.VALIDATION_ERRORS.REQUIRED'),
  startDate: Yup.date().when('serviceType', {
    is: (serviceType: string) => serviceType === ServiceTypeEnum.SALE,
    then: Yup.date().required('FORMS.VALIDATION_ERRORS.REQUIRED')
  }),
});

const complete = () => {
  console.log("ASD");
  let updatedOfferDraft: FormOfferDraft = props.newOfferDraft;
  updatedOfferDraft.duration = {number: deliveryPeriodNumber.value, calendarUnit:CalendarUnitType.MONTH};
  updatedOfferDraft.billingPeriod = {number: billingPeriodNumber.value, calendarUnit:CalendarUnitType.MONTH};
  updatedOfferDraft.invoiceDueDate = invoiceDueDate.value;

  emit('update:updatedOfferDraft', updatedOfferDraft);
  emit('complete');

};

const goBack = () => {
  emit('prevPage', {pageIndex: 2});
};

</script>
<style scoped lang="scss">

</style>
