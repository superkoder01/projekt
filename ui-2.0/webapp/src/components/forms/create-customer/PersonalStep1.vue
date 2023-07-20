<template>
  <div class="card no-border m-3">
    <h5 class="card-header">{{ $t('FORMS.HEADERS.PERSONAL_DATA') }} </h5>
    <div class="card-body">
      <Form @submit="nextPage" :validation-schema="schema" v-slot="{errors}">
        <div class="p-fluid formgrid grid ">
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="firstName" name="firstName" placeholder=" " type="text" class="form-control" :class="{'is-invalid': errors.firstName}"></Field>
            <span>{{ $t('FORMS.PLACEHOLDERS.FIRST_NAME') }}</span>
            <div class="invalid-feedback">{{ errors.firstName ? $t(errors.firstName) : '' }}</div>
          </div>
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="lastName" name="lastName" placeholder=" " type="text" class="form-control" :class="{'is-invalid': errors.lastName}"></Field>
            <span>{{ $t('FORMS.PLACEHOLDERS.LAST_NAME') }}</span>
            <div class="invalid-feedback">{{ errors.lastName ? $t(errors.lastName) : '' }}</div>
          </div>
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="phone" name="phone" placeholder=" " type="text" class="form-control" :class="{'is-invalid': errors.phone}"></Field>
            <span>{{ $t('FORMS.PLACEHOLDERS.PHONE') }}</span>
            <div class="invalid-feedback">{{ errors.phone ? $t(errors.phone) : '' }}</div>
          </div>
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="type" name="type" placeholder=" " as="select" required class="form-control" :class="{'is-invalid': errors.type}">
              <option v-for="type in typeList" :key="type.id" :value="type.name">
                {{ $t("GLOBALS.CUSTOMER_TYPES." + type.name) }}
              </option>
            </Field>
            <span>{{ $t('FORMS.PLACEHOLDERS.TYPE') }}</span>
            <div class="invalid-feedback">{{ errors.type ? $t(errors.type) : '' }}</div>
          </div>
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="pesel" name="pesel" placeholder=" " type="text" class="form-control" :class="{'is-invalid': errors.pesel}" :validateOnInput="true"></Field>
            <span>{{ $t('FORMS.PLACEHOLDERS.PESEL') }}</span>
            <div class="invalid-feedback">{{ errors.pesel ? $t(errors.pesel) : '' }}</div>
          </div>
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="email" name="email" placeholder=" " type="text" class="form-control" :class="{'is-invalid': errors.email}" :validateOnInput="true"></Field>
            <span>{{$t('FORMS.PLACEHOLDERS.EMAIL')}}</span>
            <div class="invalid-feedback">{{errors.email ? $t(errors.email): ''}}</div>
          </div>
        </div>
        <div class="grid grid-nogutter justify-content-between">
<!--          <Button :label="$t('GLOBALS.BUTTONS.BACK')" @click="goBack" type="submit" icon="pi pi-angle-left" icon-pos="left"></Button>-->
          <Button :label="$t('GLOBALS.BUTTONS.BACK')" @click="prevPage"  icon="pi pi-angle-left"></Button>
          <Button :label="$t('GLOBALS.BUTTONS.NEXT')" type="submit" icon="pi pi-angle-right" icon-pos="right"></Button>
        </div>
      </Form>
    </div>
  </div>
</template>

<script setup lang="ts">
import {ref, defineEmits, PropType, onActivated} from 'vue';
import {Form, Field} from 'vee-validate';
import * as Yup from 'yup';
import {Customer} from './Customer';
import {ClientTypeEnum} from "@/models/billing/enum/client-type.enum";
import {onBeforeMount} from "@vue/runtime-core";
import {useRouter} from "vue-router";
import { CustomerTypeEnum } from "@/models/billing/enum/customer-type.enum";
import {validatePESEL} from '@/utils/validators';
import {useContextStore} from "@/store/context.store";
import factoryApi from '@/api/factory.api';
import {YupSequentialStringSchema} from "@/utils/yup-utils";

const context = useContextStore();


const emit = defineEmits(['nextPage', 'prevPage', 'update:newCustomer']);

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
  if (props.useSelectedCustomer) {
    firstName.value = props.newCustomer.firstName;
    lastName.value = props.newCustomer.lastName;
    phone.value = props.newCustomer.phone;
    type.value = props.newCustomer.customerTypeName;
    pesel.value = props.newCustomer.pesel;
    email.value = props.newCustomer.email;
  }
});

const firstName = ref();
const lastName = ref();
const phone = ref();
const type = ref();
const selectedCustomerType = ref();

const nip = ref();
const regon = ref();
const krs = ref();

const typeList = ref([
  {
    id: 1,
    name: ClientTypeEnum.PROSUMER
  },
  {
    id: 2,
    name: ClientTypeEnum.CONSUMER
  },
  {
    id: 3,
    name: ClientTypeEnum.PRODUCER
  }]);
const pesel = ref();
const email = ref();
const router = useRouter();
const customerApi = factoryApi.customerApi();


const schema = Yup.object().shape({
  firstName: Yup.string()
    .required('FORMS.VALIDATION_ERRORS.REQUIRED'),
  lastName: Yup.string()
    .required('FORMS.VALIDATION_ERRORS.REQUIRED'),
  phone: Yup.string()
    .required('FORMS.VALIDATION_ERRORS.REQUIRED')
    .matches(/^\d+$/, 'FORMS.VALIDATION_ERRORS.PHONE.NaN')
    .min(8, 'FORMS.VALIDATION_ERRORS.PHONE.INVALID_LENGTH')
    .max(13, 'FORMS.VALIDATION_ERRORS.PHONE.INVALID_LENGTH'),
  type: Yup.string()
    .required('FORMS.VALIDATION_ERRORS.REQUIRED'),
  pesel: Yup.string()
    .required('FORMS.VALIDATION_ERRORS.REQUIRED')
    .test('validate-pesel', 'FORMS.VALIDATION_ERRORS.INVALID_PESEL',
      function (value) {
        if (value != undefined && context.allowExtendedValidation) {
        return validatePESEL(value);
        } else {
          return true;
        }
      }),
  email: YupSequentialStringSchema([
      Yup.string().required('FORMS.VALIDATION_ERRORS.REQUIRED'),
      Yup.string().email('FORMS.VALIDATION_ERRORS.EMAIL.INCORRECT_FORMAT'),
      Yup.string().test('email-already-exists', 'FORMS.VALIDATION_ERRORS.EMAIL.ALREADY_IN_USE', async (email) => {
        if(email != undefined) {
          return await customerApi.checkEmailAvailability(email, false, null);
        }
        return true;
      })
    ]
  )

});

const prevPage = () => {
  if(props.isBusinessClient) {
    emit('prevPage', {pageIndex: 1});
  } else {
    router.push("/customers")
  }
};

const nextPage = () => {

  let updatedNewCustomer: Customer = props.newCustomer;

  updatedNewCustomer.firstName = firstName.value;
  updatedNewCustomer.lastName = lastName.value;
  updatedNewCustomer.phone = phone.value;
  updatedNewCustomer.customerTypeName = type.value;
  updatedNewCustomer.pesel = pesel.value;
  updatedNewCustomer.email = email.value;
  updatedNewCustomer.login = email.value;

  if(props.isBusinessClient) {
    emit('nextPage', {pageIndex: 1});
  } else {
    emit('nextPage', {pageIndex: 0});
  }

  emit('update:newCustomer', updatedNewCustomer);
};
const goBack = () => {

  router.push({name: 'customer_tabs_info'});
};
</script>

<style scoped>

</style>
