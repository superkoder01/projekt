<template>
  <div class="card no-border m-3">
  <h5 class="card-header">Adres </h5>
    <div class="card-body">
      <Form @submit="complete" :validation-schema="schema" v-slot="{errors}">
        <div class="p-fluid formgrid grid ">
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="street" name="street" placeholder=" " type="text" class="form-control" :class="{'is-invalid': errors.street}"></Field>
            <span>{{$t('FORMS.PLACEHOLDERS.STREET')}}</span>
            <div class="invalid-feedback">{{errors.street ? $t(errors.street) : ''}}</div>
          </div>
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="buildingNumber" name="buildingNumber" placeholder=" " type="text" class="form-control" :class="{'is-invalid': errors.buildingNumber}"></Field>
            <span>{{$t('FORMS.PLACEHOLDERS.BUILDING_NUMBER')}}</span>
            <div class="invalid-feedback">{{errors.buildingNumber ? $t(errors.buildingNumber) : ''}}</div>
          </div>
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="apartmentNumber" name="apartmentNumber" placeholder=" " type="text" class="form-control" :class="{'is-invalid': errors.apartmentNumber}"></Field>
            <span>{{$t('FORMS.PLACEHOLDERS.APARTMENT_NUMBER')}}</span>
            <div class="invalid-feedback">{{errors.apartmentNumber ? $t(errors.apartmentNumber) : ''}}</div>
          </div>
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="city" name="city" placeholder=" " type="text" class="form-control" :class="{'is-invalid': errors.city}"></Field>
            <span>{{$t('FORMS.PLACEHOLDERS.CITY')}}</span>
            <div class="invalid-feedback">{{errors.city ? $t(errors.city) : ''}}</div>
          </div>
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="postalCode" name="postalCode" placeholder=" " type="text" class="form-control" :class="{'is-invalid': errors.postalCode}"></Field>
            <span>{{$t('FORMS.PLACEHOLDERS.POSTAL_CODE')}}</span>
            <div class="invalid-feedback">{{errors.postalCode ? $t(errors.postalCode) : ''}}</div>
          </div>
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="province" name="province" placeholder=" " type="text" class="form-control" :class="{'is-invalid': errors.province}"></Field>
            <span>{{$t('FORMS.PLACEHOLDERS.PROVINCE')}}</span>
            <div class="invalid-feedback">{{errors.province ? $t(errors.province) : ''}}</div>
          </div>
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="country" name="country" placeholder=" " type="text" class="form-control" :class="{'is-invalid': errors.country}"></Field>
            <span>{{$t('FORMS.PLACEHOLDERS.COUNTRY')}}</span>
            <div class="invalid-feedback">{{errors.country ? $t(errors.country) : ''}}</div>
          </div>
<!--          <h5>Adres korespondencyjny</h5>
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="correspondenceStreet" name="correspondenceStreet" placeholder=" " type="text" class="form-control" :class="{'is-invalid': errors.correspondenceStreet}"></Field>
            <span>{{$t('FORMS.PLACEHOLDERS.STREET')}}</span>
            <div class="invalid-feedback">{{errors.correspondenceStreet ? $t(errors.correspondenceStreet) : ''}}</div>
          </div>
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="correspondenceBuildingNumber" name="correspondenceBuildingNumber" placeholder=" " type="text" class="form-control" :class="{'is-invalid': errors.correspondenceBuildingNumber}"></Field>
            <span>{{$t('FORMS.PLACEHOLDERS.BUILDING_NUMBER')}}</span>
            <div class="invalid-feedback">{{errors.correspondenceBuildingNumber ? $t(errors.correspondenceBuildingNumber) : ''}}</div>
          </div>
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="correspondenceApartmentNumber" name="correspondenceApartmentNumber" placeholder=" " type="text" class="form-control" :class="{'is-invalid': errors.correspondenceApartmentNumber}"></Field>
            <span>{{$t('FORMS.PLACEHOLDERS.APPARTMENT_NUMBER')}}</span>
            <div class="invalid-feedback">{{errors.correspondenceApartmentNumber ? $t(errors.correspondenceApartmentNumber) : ''}}</div>
          </div>
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="correspondenceCity" name="correspondenceCity" placeholder=" " type="text" class="form-control" :class="{'is-invalid': errors.correspondenceCity}"></Field>
            <span>{{$t('FORMS.PLACEHOLDERS.CITY')}}</span>
            <div class="invalid-feedback">{{errors.correspondenceCity ? $t(errors.correspondenceCity) : ''}}</div>
          </div>
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="correspondencePostalCode" name="correspondencePostalCode" placeholder=" " type="text" class="form-control" :class="{'is-invalid': errors.correspondencePostalCode}"></Field>
            <span>{{$t('FORMS.PLACEHOLDERS.POSTAL_CODE')}}</span>
            <div class="invalid-feedback">{{errors.correspondencePostalCode ? $t(errors.correspondencePostalCode) : ''}}</div>
          </div>
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="correspondenceProvince" name="correspondenceProvince" placeholder=" " type="text" class="form-control" :class="{'is-invalid': errors.correspondenceProvince}"></Field>
            <span>{{$t('FORMS.PLACEHOLDERS.PROVINCE')}}</span>
            <div class="invalid-feedback">{{errors.correspondenceProvince ? $t(errors.correspondenceProvince) : ''}}</div>
          </div>
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="correspondenceCountry" name="correspondenceCountry" placeholder=" " type="text" class="form-control" :class="{'is-invalid': errors.correspondenceCountry}"></Field>
            <span>{{$t('FORMS.PLACEHOLDERS.COUNTRY')}}</span>
            <div class="invalid-feedback">{{errors.correspondenceCountry ? $t(errors.correspondenceCountry) : ''}}</div>
          </div>-->
        </div>
        <div class="grid grid-nogutter justify-content-between">
          <Button :label="$t('GLOBALS.BUTTONS.BACK')" @click="prevPage"  icon="pi pi-angle-left"></Button>
          <Button :label="useSelectedCustomer ? $t('GLOBALS.BUTTONS.SAVE') : $t('GLOBALS.BUTTONS.CREATE')"  type="submit" icon="pi pi-angle-right"></Button>
        </div>
      </Form>
    </div>
  </div>
</template>

<script setup lang="ts">
import {ref, defineEmits, PropType} from 'vue';
import { Form, Field} from 'vee-validate';
import * as Yup from 'yup';
import { Customer } from './Customer';
import {onBeforeMount} from "@vue/runtime-core";

const emit = defineEmits(['complete', 'prevPage', 'update:newCustomer']);
const props = defineProps({
  newCustomer: {
      type: Object as PropType<Customer>,
      required: true
  },
  useSelectedCustomer: {
    type: Boolean
  },
  isBusinessClient: {
    type: Boolean
  }
});

onBeforeMount(() => {
  if(props.useSelectedCustomer) {
    street.value = props.newCustomer.street;
    buildingNumber.value = props.newCustomer.buildingNumber;
    apartmentNumber.value = props.newCustomer.apartmentNumber;
    postalCode.value = props.newCustomer.postalCode;
    city.value = props.newCustomer.postalCode;
    province.value = props.newCustomer.province;
    country.value = props.newCustomer.country;

  }
});

const street = ref();
const buildingNumber = ref();
const apartmentNumber = ref();
const city = ref();
const postalCode = ref();
const province = ref();
const country = ref();

const correspondenceStreet = ref();
const correspondenceBuildingNumber = ref();
const correspondenceApartmentNumber = ref();
const correspondenceCity = ref();
const correspondencePostalCode = ref();
const correspondenceProvince = ref();
const correspondenceCountry = ref();

const schema = Yup.object().shape({
  street: Yup.string().required('FORMS.VALIDATION_ERRORS.REQUIRED'),
  buildingNumber: Yup.string()
    .required('FORMS.VALIDATION_ERRORS.REQUIRED'),
  // apartmentNumber: Yup.string()
  //   .required('FORMS.VALIDATION_ERRORS.REQUIRED'),
  city: Yup.string().required('FORMS.VALIDATION_ERRORS.REQUIRED'),
  postalCode: Yup.string()
    .required('FORMS.VALIDATION_ERRORS.REQUIRED')
    .matches(/^[0-9]{2}-[0-9]{3}/, 'FORMS.VALIDATION_ERRORS.POSTAL_CODE'),
  province: Yup.string().required('FORMS.VALIDATION_ERRORS.REQUIRED'),
  country: Yup.string().required('FORMS.VALIDATION_ERRORS.REQUIRED')
});


const prevPage = () => {
  if(props.isBusinessClient) {
    emit('prevPage', {pageIndex: 2});
  } else {
    emit('prevPage', {pageIndex: 1});
  }

};

const complete = () => {

  let updatedNewCustomer: Customer = props.newCustomer;
  updatedNewCustomer.street = street.value;
  updatedNewCustomer.buildingNumber = buildingNumber.value;
  updatedNewCustomer.apartmentNumber = apartmentNumber.value;
  updatedNewCustomer.city = city.value;
  updatedNewCustomer.postalCode = postalCode.value;
  updatedNewCustomer.province = province.value;
  updatedNewCustomer.country = country.value;

  emit('update:newCustomer', updatedNewCustomer);
  emit('complete');

};
</script>

<style scoped>

</style>
