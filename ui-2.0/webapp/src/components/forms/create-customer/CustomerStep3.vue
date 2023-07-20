<template>
  <div class="card no-border m-3">
    <h5 class="card-header">{{ $t('FORMS.HEADERS.PERSONAL_DATA') }} </h5>
    <div class="card-body">
      <Form @submit="nextPage" :validation-schema="schema" v-slot="{errors}">

        <div class="p-fluid formgrid grid " >
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="nip" name="nip" placeholder=" " type="number" class="form-control" :class="{'is-invalid': errors.nip}"></Field>
            <span>{{$t('FORMS.PLACEHOLDERS.NIP')}}</span>
            <div class="invalid-feedback">{{errors.nip ? $t(errors.nip) : ''}}</div>
          </div>
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="regon" name="regon" placeholder=" "  type="number" class="form-control" :class="{'is-invalid': errors.regon}"></Field>
            <span>{{$t('FORMS.PLACEHOLDERS.REGON')}}</span>
            <div class="invalid-feedback">{{errors.regon ? $t(errors.regon) :  ''}}</div>
          </div>
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="krs" name="phone" placeholder=" " type="number" class="form-control" :class="{'is-invalid': errors.krs}"></Field>
            <span>{{$t('FORMS.PLACEHOLDERS.KRS')}}</span>
            <div class="invalid-feedback">{{errors.krs ? $t(errors.krs) : ''}}</div>
          </div>
<!--          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="companyType" name="companyType" placeholder=" " type="text" class="form-control" :class="{'is-invalid': errors.companyType}"></Field>
            <span>{{$t('FORMS.PLACEHOLDERS.COMPANY_TYPE')}}</span>
            <div class="invalid-feedback">{{errors.companyType ? $t(errors.companyType): ''}}</div>
          </div>-->

        </div>
        <div class="grid grid-nogutter justify-content-between">
          <Button :label="$t('GLOBALS.BUTTONS.BACK')" @click="prevPage"  icon="pi pi-angle-left"></Button>
          <Button :label="$t('GLOBALS.BUTTONS.NEXT')" type="submit" icon="pi pi-angle-right" icon-pos="right"></Button>
        </div>
      </Form>
    </div>
  </div>
</template>

<script setup lang="ts">
import {ref, defineEmits, PropType} from 'vue';
import { Form, Field } from 'vee-validate';
import * as Yup from 'yup';
import { Customer } from './Customer';
import { CustomerTypeEnum } from "@/models/billing/enum/customer-type.enum";
import router from "@/router";

const emit = defineEmits(['nextPage', 'prevPage', 'update:newCustomer']);

const props = defineProps({
  newCustomer: {
    type: Object as PropType<Customer>,
    required: true
  }
});

const nip = ref();
const regon = ref();
const krs = ref();
const districtCourt = ref();
const companyType = ref();
const typeList = ref([
  {
    id: 1,
    name: 'PROSUMER'
  },
  {
    id: 2,
    name: 'CONSUMER'
  },
  {
    id: 3,
    name: 'PRODUCER'
  }]);
const pesel = ref();
const email = ref();

const schema = Yup.object().shape({
  nip: Yup.number().required('FORMS.VALIDATION_ERRORS.REQUIRED'),
  regon: Yup.number(),
  krs: Yup.number(),
  // companyType: Yup.string().required('FORMS.VALIDATION_ERRORS.REQUIRED'),
});

const prevPage = () => {
  router.push("/customers")

};


const nextPage = () => {

  let updatedNewCustomer: Customer = props.newCustomer;

  updatedNewCustomer.nip = nip.value;
  updatedNewCustomer.regon = regon.value;
  updatedNewCustomer.krs = krs.value;
  // updatedNewCustomer.lineOfBusiness = companyType.value;

  emit('nextPage', {pageIndex: 0});
  emit('update:newCustomer', updatedNewCustomer);
};

</script>

<style scoped>

</style>
