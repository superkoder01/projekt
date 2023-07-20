<template>
  <div class="card no-border m-3">
    <h5 class="card-header">{{ $t("FORMS.HEADERS.PROVIDER_DATA") }}</h5>
    <div class="card-body">
      <Form @submit="nextPage" :validation-schema="baseSchema" v-slot="{ errors }">
        <div class="p-fluid formgrid grid">
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="name" name="name" placeholder=" " type="text" class="form-control" :class="{ 'is-invalid': errors.name }"></Field>
            <span>{{ $t('FORMS.PLACEHOLDERS.PROVIDER_NAME') }}</span>
            <div class="invalid-feedback">{{ errors.name ? $t(errors.name) : "" }}</div>
          </div>
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="krs" name="krs" placeholder=" " type="text" class="form-control" :class="{ 'is-invalid': errors.krs }" ></Field>
            <span>{{ $t('FORMS.PLACEHOLDERS.KRS') }}</span>
            <div class="invalid-feedback">{{ errors.krs ? $t(errors.krs) : "" }}</div>
          </div>
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="regon" name="regon" placeholder=" " type="text" class="form-control" :class="{ 'is-invalid': errors.regon }" :validateOnInput="true"></Field>
            <span>{{ $t('FORMS.PLACEHOLDERS.REGON') }}</span>
            <div class="invalid-feedback">{{ errors.regon ? $t(errors.regon) : "" }}</div>
          </div>
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="nip" name="nip" placeholder=" " type="text" class="form-control" :class="{ 'is-invalid': errors.nip }" :validateOnInput="true"></Field>
            <span>{{ $t('FORMS.PLACEHOLDERS.NIP') }}</span>
            <div class="invalid-feedback">{{ errors.nip ? $t(errors.nip) : "" }}</div>
          </div>
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="type" name="type" placeholder=" " type="text" class="form-control" :class="{ 'is-invalid': errors.type }"></Field>
            <span>{{ $t('FORMS.PLACEHOLDERS.TYPE') }}</span>
            <div class="invalid-feedback">{{ errors.type ? $t(errors.type) : "" }}</div>
          </div>
        </div>
        <div class="field col-12 md:col-6 mb-4">
          <Field v-model="blockchainAccAddress" name="blockchainAccAddress" placeholder=" " type="text" class="form-control" :class="{ 'is-invalid': errors.blockchainAccAddress }"></Field>
          <span>{{ $t('FORMS.PLACEHOLDERS.BLOCKCHAIN_ACC') }}</span>
          <div class="invalid-feedback">
            {{ errors.blockchainAccAddress ? $t(errors.blockchainAccAddress) : "" }}
          </div>
        </div>
        <div class="field col-12 md:col-6 mb-4">
          <Field v-model="endDate" name="endDate" placeholder=" " onfocus="this.type='date'" type="text" class="form-control" :class="{ 'is-invalid': errors.workEndDate }"></Field>
          <span>{{ $t('FORMS.PLACEHOLDERS.END_DATE') }}</span>
          <div class="invalid-feedback">
            {{ errors.workEndDate ? $t(errors.workEndDate) : "" }}
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
import {PropType, ref} from "vue";
import {Field, Form} from "vee-validate";
import * as Yup from "yup";
import {Provider} from "@/models/provider";
import {onBeforeMount} from "@vue/runtime-core";
import moment from "moment";
import {useRouter} from "vue-router";
import {validateNIP, validateREGON} from "@/utils/validators";
import {useContextStore} from "@/store/context.store";
import { YupSequentialStringSchema } from "@/utils/yup-utils";
import factoryApi from "@/api/factory.api";

const props = defineProps({
  newProvider: {
    type: Object as PropType<Provider>,
    required: true,
  },
  useSelectedProvider: {
    type: Boolean
  },
});
const emit = defineEmits(["nextPage", "prevPage", "update:newProvider"]);
const router = useRouter();
const context = useContextStore();

onBeforeMount(() => {
  name.value = props.newProvider.name;
  krs.value = props.newProvider.krs;
  regon.value = props.newProvider.regon;
  nip.value = props.newProvider.nip;
  type.value = props.newProvider.type;
  blockchainAccAddress.value = props.newProvider.blockchainAccAddress;
  endDate.value = moment(new Date(props.newProvider.licenseExpirationDate)).format('YYYY-MM-DD');
});

const name = ref();
const krs = ref();
const regon = ref();
const nip = ref();
const type = ref();
const typeList = ref([
  {
    id: 1,
    name: 'SOE'
  }]);
const blockchainAccAddress = ref();
const endDate = ref<string>((moment(new Date())).format('YYYY-MM-DD'));
const providerApi = factoryApi.providersApi();
const baseSchema = Yup.object().shape({
  name: YupSequentialStringSchema([ 
    Yup.string().required("FORMS.VALIDATION_ERRORS.REQUIRED"),
    Yup.string().test('name-already-exists', 'FORMS.VALIDATION_ERRORS.NAME_ALREADY_IN_USE', async (name) => {
        if(name != undefined) {
          return await providerApi.checkFieldAvailability(name, 'name', false, null);
        }
        return true;
      })
  ]),  
  krs: YupSequentialStringSchema([
    Yup.string().required("FORMS.VALIDATION_ERRORS.REQUIRED"),
    Yup.string().test('krs-already-exists', 'FORMS.VALIDATION_ERRORS.KRS_ALREADY_IN_USE', async (krs) => {
        if(krs != undefined) {
          return await providerApi.checkFieldAvailability(krs, 'krs', false, null);
        }
        return true;
      })
   ]),  
  regon: YupSequentialStringSchema([
    Yup.string().required("FORMS.VALIDATION_ERRORS.REQUIRED"),
    Yup.string().test('validate-regon', 'FORMS.VALIDATION_ERRORS.INVALID_REGON',
      function (value) {
        if (value != undefined && context.allowExtendedValidation) {
          return validateREGON(value);
        } else {
          return true;
        }
      }),
    Yup.string().test('regon-already-exists', 'FORMS.VALIDATION_ERRORS.REGON_ALREADY_IN_USE', async (regon) => {
        if(regon != undefined) {
          return await providerApi.checkFieldAvailability(regon, 'regon', false, null);
        }
        return true;
      })
   ]),      
  nip: YupSequentialStringSchema([ 
    Yup.string().required("FORMS.VALIDATION_ERRORS.REQUIRED"),
    Yup.string().test('validate-nip', 'FORMS.VALIDATION_ERRORS.INVALID_NIP',
      function (value) {
        if (value != undefined && context.allowExtendedValidation) {
          return validateNIP(value);
        } else {
          return true;
        }
      }),
    Yup.string().test('nip-already-exists', 'FORMS.VALIDATION_ERRORS.NIP_ALREADY_IN_USE', async (nip) => {
        if(nip != undefined) {
          return await providerApi.checkFieldAvailability(nip, 'nip', false, null);
        }
        return true;
      })
  ]),    
  type: Yup.string().required("FORMS.VALIDATION_ERRORS.REQUIRED"),
  endDate: Yup.string()
    .required("FORMS.VALIDATION_ERRORS.REQUIRED")
    .matches(
      /^\d{4}-(0[1-9]|1[012])-(0[1-9]|[12][0-9]|3[01])$/,
      "FORMS.VALIDATION_ERRORS.DATE_FORMAT"
    ),
});

const goBack = () => {
  router.push('/providers/provider_tabs');
};

const nextPage = () => {
  let updatedProvider: Provider = props.newProvider;
  updatedProvider.name = name.value;
  updatedProvider.krs = krs.value;
  updatedProvider.regon = regon.value;
  updatedProvider.nip = nip.value;
  updatedProvider.type = type.value;
  updatedProvider.blockchainAccAddress = blockchainAccAddress.value;
  if (endDate.value) {
    let dateEnd = endDate.value.toString().split("-");
    updatedProvider.licenseExpirationDate = new Date(
      Number(dateEnd[0]),
      Number(dateEnd[1]) - 1,
      Number(dateEnd[2]) + 1,
      1,
      0,
      0
    );
  }
  emit("nextPage", {pageIndex: 0});

  emit("update:newProvider", updatedProvider);
};
</script>

<style scoped></style>
