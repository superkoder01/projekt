<template>
  <div class="card m-3">
    <div class="card-body">
      <Form @submit="complete" :validation-schema="schema" v-slot="{errors}">
        <div class="p-fluid formgrid grid ">
          <template v-if="newPricing.type===ServiceTypeEnum.REPURCHASE_RDN">
            <div  class="rdn-info-container">
                <div> {{ $t('PRICING_VIEW.OTHER.RDN_INFO') }}</div>
                <span class="link" @click="goToRDNPage"> https://tge.pl/energia-elektryczna-rdn </span>
            </div>
          </template>


          <template v-if="newPricing.type===ServiceTypeEnum.REPURCHASE">
            <div class="field col-12 md:col-6 mb-4">
              <Field v-model="fixedPrice" name="fixedPrice" placeholder=" "  type="text" class="form-control" :class="{'is-invalid': errors.fixedPrice}" />
              <span>{{$t('FORMS.PLACEHOLDERS.PRICE')}}</span>
              <div class="invalid-feedback">{{errors.fixedPrice ? $t(errors.fixedPrice) : ''}}</div>
            </div>
          </template>

          <template v-if="newPricing.type===ServiceTypeEnum.SALE">
            <div class="field col-12 md:col-6 mb-4">
              <Field v-model="zonePrice" name="zones" placeholder=" " type="text" class="form-control" :class="{'is-invalid': errors.zones}" />
              <span>{{$t('FORMS.PLACEHOLDERS.ZONE_PRICE')}}</span>
              <div class="invalid-feedback">{{errors.zones ? $t(errors.zones) : ''}}</div>
            </div>
            <div>{{$t('FORMS.PLACEHOLDERS.MONTHLY_FEE')}}</div>
            <div v-for="(fee, index) of feeList" :key="fee" class="field col-12 md:col-6 mb-4">
              <Field v-model="fee.price.cost" :rules="isRequired" :name="`fee${index}`" placeholder=" " type="text" class="form-control" :class="{'is-invalid': errors[`fee${index}`]}"></Field>
              <span>{{ fee.from }} - {{ fee.to }} kWp</span>
              <div class="invalid-feedback">{{errors[`fee${index}`]}}</div>
            </div>
          </template>
        </div>
        <div class="grid grid-nogutter justify-content-between">
          <Button :label="$t('GLOBALS.BUTTONS.BACK')" @click="prevPage"  icon="pi pi-angle-left"></Button>
          <Button :label="$t('GLOBALS.BUTTONS.CREATE')" type="submit" icon="pi pi-angle-right"></Button>
        </div>
      </Form>
    </div>
  </div>
</template>

<script setup lang="ts">
import {defineEmits, PropType, reactive, ref} from 'vue';
import {Field, Form} from 'vee-validate';
import * as Yup from 'yup';
import {FormPricing} from "@/components/forms/create-pricing/Pricing";
import {CalendarUnitType, CommercialFee, CommercialFeeUnitType} from "@/models/billing/billing";
import {ServiceTypeEnum} from "@/models/billing/enum/service-type.enum";

const emit = defineEmits(['complete', 'prevPage','update:newPricing']);

const props = defineProps({

  newPricing: {
    type: Object as PropType<FormPricing>,
    required: true
  }
});
const isRequired = Yup.number().required('This field is required');
// eslint-disable-next-line vue/no-setup-props-destructure
const showZones =  props.newPricing.type;
const schema = Yup.object().shape({
  showZones: Yup.string(),
  zones: Yup.string().when("showZones", {
    is: 'fixed',
    then: Yup.string().required('FORMS.VALIDATION_ERRORS.REQUIRED')
  })

});

const feeList = reactive<CommercialFee[]>([
  {from:'0', to:'5', unit: CommercialFeeUnitType.KWP, price: {calendarUnit: CalendarUnitType.MONTH,currency: "pln"}},
  {from:'5.1', to: '10', unit: CommercialFeeUnitType.KWP, price: {calendarUnit: CalendarUnitType.MONTH,currency: "pln"}},
  {from:'10.1', to: '20', unit: CommercialFeeUnitType.KWP, price: {calendarUnit: CalendarUnitType.MONTH,currency: "pln"}},
  {from:'20.1', unit: CommercialFeeUnitType.KWP, price: {calendarUnit: CalendarUnitType.MONTH,currency: "pln"}}
]);

const zonePrice = ref();
const fixedPrice = ref();

const prevPage = () => {
  emit('prevPage', {pageIndex: 1});
};

const complete = () => {

  let updatedNewPricing: FormPricing = props.newPricing;

  if(props.newPricing.type === ServiceTypeEnum.REPURCHASE) {
    updatedNewPricing.fixedPrice = fixedPrice.value;
  }
  else if(props.newPricing.type === ServiceTypeEnum.SALE) {
    updatedNewPricing.zones = zonePrice.value;
    updatedNewPricing.commercialFee = feeList;
  }

  emit('complete');
  emit('update:newPricing', updatedNewPricing);
};
function goToRDNPage() {
  window.open('https://tge.pl/energia-elektryczna-rdn', '_blank')?.focus();
}
</script>

<style scoped>
.rdn-info-container{
  margin-top: -10px;
  margin-bottom: 1rem;
}
.link{
      color: blue;
}
.link:hover {
      cursor: pointer;
    }
</style>
