<template>
  <div class="card m-3">
    <div class="card-body">
      <Form @submit="complete" :validation-schema="schema" v-slot="{errors}">
        <div class="p-fluid formgrid grid ">
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="estimatedAnnualElectricityConsumption" name="estimatedAnnualElectricityConsumption" placeholder=" " type="number" class="form-control"
                   :class="{'is-invalid': errors.estimatedAnnualElectricityConsumption}"/>
            <span>{{$t("FORMS.PLACEHOLDERS.ESTIMATED_ENERGY_USAGE")}}</span>
            <div class="invalid-feedback">{{ errors.estimatedAnnualElectricityConsumption ? $t(errors.estimatedAnnualElectricityConsumption) : '' }}</div>
          </div>
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="estimatedAnnualElectricityProduction" name="estimatedAnnualElectricityProduction" placeholder=" " type="number" class="form-control"
                   :class="{'is-invalid': errors.estimatedAnnualElectricityProduction}"/>
            <span>{{$t("FORMS.PLACEHOLDERS.ESTIMATED_ENERGY_PRODUCTION")}}</span>
            <div class="invalid-feedback">{{ errors.estimatedAnnualElectricityProduction ? $t(errors.estimatedAnnualElectricityProduction) : '' }}</div>
          </div>
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="numberOfPPE" name="numberOfPPE" placeholder=" " type="number" class="form-control" min="0"
                   :class="{'is-invalid': errors.numberOfPPE}"/>
            <span>{{$t("FORMS.PLACEHOLDERS.NUMBER_OF_PPE")}}</span>
            <div class="invalid-feedback">{{ errors.numberOfPPE ? $t(errors.numberOfPPE): '' }}</div>
          </div>

        </div>
        <div class="grid grid-nogutter justify-content-between">
          <Button :label="$t('GLOBALS.BUTTONS.BACK')" @click="prevPage"  icon="pi pi-angle-left"></Button>
          <Button :label="useSelectedOffer ? $t('GLOBALS.BUTTONS.SAVE') : $t('GLOBALS.BUTTONS.CREATE')" type="submit" icon="pi pi-angle-right" icon-pos="right"></Button>
        </div>
      </Form>
    </div>
  </div>
</template>

<script setup lang="ts">
import {defineEmits, PropType, ref} from 'vue';
import {Field, Form} from 'vee-validate';
import * as Yup from 'yup';
import {onBeforeMount} from "@vue/runtime-core";
import { OfferForm} from "@/components/forms/create-offer/OfferForm";

const emit = defineEmits(['complete', 'prevPage', 'update:newOffer']);

const props = defineProps({

  newOffer: {
    type: Object as PropType<OfferForm>,
    required: true
  },
  useSelectedOffer: {
    type: Boolean
  }
});

onBeforeMount(() => {
  if(props.useSelectedOffer) {
    numberOfPPE.value = props.newOffer.numberOfPPE;
    estimatedAnnualElectricityConsumption.value = props.newOffer.estimatedAnnualElectricityConsumption.amount;
    estimatedAnnualElectricityProduction.value = props.newOffer.estimatedAnnualElectricityProduction.amount;
  }
});



const numberOfPPE = ref(1);
const estimatedAnnualElectricityConsumption = ref();
const estimatedAnnualElectricityProduction = ref();

const schema = Yup.object().shape({
  numberOfPPE: Yup.number()
                  .typeError('FORMS.VALIDATION_ERRORS.NUMBERS.NaN')
                  .positive('FORMS.VALIDATION_ERRORS.NUMBERS.NEGATIVE')
                  .integer('FORMS.VALIDATION_ERRORS.NUMBERS.NOT_AN_INT')
                  .required('FORMS.VALIDATION_ERRORS.REQUIRED'),
  estimatedAnnualElectricityConsumption: Yup.string()
    .required('FORMS.VALIDATION_ERRORS.REQUIRED'),

});


const prevPage = () => {
  emit('prevPage', {pageIndex: 1});
};

const complete = () => {

  let updatedNewOffer: OfferForm = props.newOffer;
  updatedNewOffer.numberOfPPE = numberOfPPE.value;

  const startDate = new Date();
  const endDate = new Date(startDate.getFullYear(), startDate.getMonth(), startDate.getDate()+30);
  updatedNewOffer.startDate = startDate;
  updatedNewOffer.endDate = endDate;
  updatedNewOffer.estimatedAnnualElectricityConsumption = {amount: estimatedAnnualElectricityConsumption.value, unit: "MWh"};
  updatedNewOffer.estimatedAnnualElectricityProduction = {amount: estimatedAnnualElectricityProduction.value, unit: "MWh"};

  emit('complete');
  emit('update:newOffer', updatedNewOffer);
};

</script>
<style scoped lang="scss">

</style>
