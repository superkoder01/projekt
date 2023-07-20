<template>
  <div class="card no-border m-3">
  <h5 class="card-header">{{ $t('FORMS.HEADERS.ADDRESS') }}</h5>
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
        </div>
        <div class="grid grid-nogutter justify-content-between">
          <Button :label="$t('GLOBALS.BUTTONS.BACK')" @click="prevPage"  icon="pi pi-angle-left"></Button>
          <Button :label="useSelectedEmployee ? $t('GLOBALS.BUTTONS.SAVE') : $t('GLOBALS.BUTTONS.CREATE')" type="submit" icon="pi pi-angle-right" icon-pos="right"></Button>
        </div>
      </Form>
    </div>
  </div>
</template>

<script setup lang="ts">
import {ref, defineEmits, PropType} from 'vue';
import { Form, Field} from 'vee-validate';
import * as Yup from 'yup';
import { FunctionalUser } from '@/components/forms/create-func-user/FunctionalUser';
import {onBeforeMount} from "@vue/runtime-core";
import {RoleEnum} from "@/services/permissions/role-enum";


const emit = defineEmits(['complete', 'prevPage', 'update:newFunctionalUser']);

const props = defineProps({
  newFunctionalUser: {
      type: Object as PropType<FunctionalUser>,
      required: true
  },
  passedRoleID: {
    type: Number as unknown as PropType<RoleEnum>,
    required: true
  },
  useSelectedEmployee: {
    type: Boolean
  },
  previousPath: {
    type: String
  },
});

onBeforeMount(() => {
  console.log(props)
  street.value = props.newFunctionalUser.street;
  buildingNumber.value = props.newFunctionalUser.buildingNumber;
  apartmentNumber.value = props.newFunctionalUser.apartmentNumber;
  city.value = props.newFunctionalUser.city;
  postalCode.value = props.newFunctionalUser.postalCode;
  province.value = props.newFunctionalUser.province;
  country.value = props.newFunctionalUser.country;

});

const street = ref();
const buildingNumber = ref();
const apartmentNumber = ref();
const city = ref();
const postalCode = ref();
const province = ref();
const country = ref();
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
  emit('prevPage', {pageIndex: 1});
};

const complete = () => {

  let updatedNewFunctionalUser: FunctionalUser = props.newFunctionalUser;
  updatedNewFunctionalUser.street = street.value;
  updatedNewFunctionalUser.buildingNumber = buildingNumber.value;
  updatedNewFunctionalUser.apartmentNumber = apartmentNumber.value;
  updatedNewFunctionalUser.city = city.value;
  updatedNewFunctionalUser.postalCode = postalCode.value;
  updatedNewFunctionalUser.province = province.value;
  updatedNewFunctionalUser.country = country.value;

  emit('complete');
  emit('update:newFunctionalUser', updatedNewFunctionalUser);
};

</script>

<style scoped>
button {
  margin-left: 1rem;
}

</style>
