<template>
  <div class="card no-border m-3">
  <h5 class="card-header"> </h5>
    <div class="card-body">
      <Form @submit="complete" :validation-schema="schema" v-slot="{ errors }">
        <div class="p-fluid formgrid grid ">
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="email" name="email" placeholder=" " type="text" class="form-control" :class="{'is-invalid': errors.email}"></Field>
            <span>{{$t('FORMS.PLACEHOLDERS.EMAIL')}}</span>
            <div class="invalid-feedback">{{errors.email ? $t(errors.email): ''}}</div>
          </div>
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="phone" name="phone" placeholder=" " type="text" class="form-control" :class="{'is-invalid': errors.phone}"></Field>
            <span>{{$t('FORMS.PLACEHOLDERS.PHONE')}}</span>
            <div class="invalid-feedback">{{errors.phone ? $t(errors.phone) : ''}}</div>
          </div>
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
            <Field v-model="apartmentNumber " name="apartmentNumber " placeholder=" " type="text" class="form-control" :class="{'is-invalid': errors.apartmentNumber }"></Field>
            <span>{{$t('FORMS.PLACEHOLDERS.APARTMENT_NUMBER')}}</span>
            <div class="invalid-feedback">{{errors.apartmentNumber  ? $t(errors.apartmentNumber) : ''}}</div>
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
        </div>
        <div class="grid grid-nogutter justify-content-between">
          <Button :label="$t('GLOBALS.BUTTONS.BACK')" @click="prevPage"  icon="pi pi-angle-left"></Button>
          <Button v-if="!useSelectedProvider" :label="$t('GLOBALS.BUTTONS.ADD_ADMIN')"  @click="nextPage" icon="pi pi-angle-right"></Button>
          <Button :label="useSelectedProvider ? $t('GLOBALS.BUTTONS.SAVE') : $t('GLOBALS.BUTTONS.CREATE')"  type="submit" icon="pi pi-angle-right"></Button>
        </div>
      </Form>
    </div>
  </div>
</template>


<script setup lang="ts">
import { Provider } from '@/models/provider';
import { Form, Field } from "vee-validate";
import { PropType, ref } from 'vue';
import * as Yup from 'yup';
import {onBeforeMount} from "@vue/runtime-core";


const emit = defineEmits(['complete', 'prevPage', 'update:newProvider', 'nextPage' ]);

const props = defineProps({
    newProvider: {
        type: Object as PropType<Provider>,
        required: true
    },
  useSelectedProvider: {
    type: Boolean
  },
});

onBeforeMount(() => {
  street.value = props.newProvider.street;
  buildingNumber.value = props.newProvider.buildingNumber;
  city.value = props.newProvider.city;
  postalCode.value = props.newProvider.postalCode;
  province.value = props.newProvider.province;
  country.value = props.newProvider.country;
  phone.value = props.newProvider.phoneNumber;
  apartmentNumber.value = props.newProvider.apartmentNumber;

});

const street = ref();
const buildingNumber = ref();
const city = ref();
const postalCode = ref();
const province = ref();
const country = ref();
const phone = ref();
const email = ref();
const apartmentNumber = ref();
const schema = Yup.object().shape({
  email: Yup.string().required('FORMS.VALIDATION_ERRORS.REQUIRED'),
  buildingNumber: Yup.string().required('FORMS.VALIDATION_ERRORS.REQUIRED'),
  city: Yup.string().required('FORMS.VALIDATION_ERRORS.REQUIRED'),
  postalCode: Yup.string().required('FORMS.VALIDATION_ERRORS.REQUIRED'),
  province: Yup.string().required('FORMS.VALIDATION_ERRORS.REQUIRED'),
  country: Yup.string().required('FORMS.VALIDATION_ERRORS.REQUIRED'),
  street: Yup.string().required('FORMS.VALIDATION_ERRORS.REQUIRED'),
  phone: Yup.string().required('FORMS.VALIDATION_ERRORS.REQUIRED'),
});
const prevPage = () => {
  emit('prevPage', {pageIndex: 1});
};
const updateValues = () => {
  let updatedProvider : Provider = props.newProvider;
  updatedProvider.email = email.value;
  updatedProvider.phoneNumber = phone.value;
  updatedProvider.street = street.value;
  updatedProvider.buildingNumber = buildingNumber.value;
  updatedProvider.apartmentNumber = apartmentNumber.value;
  updatedProvider.city = city.value;
  updatedProvider.postalCode = postalCode.value;
  updatedProvider.province = province.value;
  updatedProvider.country = country.value;
  emit('update:newProvider', updatedProvider);
};
const complete = () => {
  updateValues();
  emit('complete', {isAdminAdded: false});
};
const nextPage = () => {
  updateValues();
  emit('nextPage', {pageIndex: 1});
};

</script>
