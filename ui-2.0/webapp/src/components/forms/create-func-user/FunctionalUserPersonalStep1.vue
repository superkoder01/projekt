<template>
  <div class="card no-border m-3">
    <h5 v-if="!checkIfSuperAdmin" class="card-header">{{ $t('FORMS.HEADERS.PERSONAL_DATA') }} </h5>
    <div class="card-body">
      <Form @submit="nextPage" :validation-schema="setValidation(props.passedRoleID)" v-slot="{errors}">
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
            <Field v-model="email" name="email" placeholder=" " type="text" class="form-control" :class="{'is-invalid': errors.email}"></Field>
            <span>{{ $t('FORMS.PLACEHOLDERS.EMAIL') }} ({{ $t('FORMS.PLACEHOLDERS.LOGIN') }})</span>
            <div class="invalid-feedback">{{ errors.email ? $t(errors.email) : '' }}</div>
          </div>
          <div class="field col-12 md:col-6 mb-4" v-if="checkIfAgent">
            <Field v-model="workStartDate" name="workStartDate" placeholder=" " onfocus="this.type='date'" type="text" class="form-control" :class="{'is-invalid': errors.workStartDate}"></Field>
            <span>{{ $t('FORMS.PLACEHOLDERS.WORK_START_DATE') }}</span>
            <div class="invalid-feedback">{{ errors.workStartDate ? $t(errors.workStartDate) : '' }}</div>
          </div>
          <div class="field col-12 md:col-6 mb-4" v-if="checkIfAgent">
            <Field v-model="workEndDate" name="workEndDate" placeholder=" " onfocus="this.type='date'" type="text" class="form-control" :class="{'is-invalid': errors.workEndDate}"></Field>
            <span>{{ $t('FORMS.PLACEHOLDERS.WORK_END_DATE') }}</span>
            <div class="invalid-feedback">{{ errors.workEndDate ? $t(errors.workEndDate) : '' }}</div>
          </div>
          <div class="field col-12 md:col-6 mb-4" v-if="checkIfAgent">
            <Field v-model="nip" name="nip" placeholder=" " type="text" class="form-control" :class="{'is-invalid': errors.nip}" :validateOnInput="true"></Field>
            <span>{{ $t('FORMS.PLACEHOLDERS.NIP') }}</span>
            <div class="invalid-feedback">{{ errors.nip ? $t(errors.nip) : '' }}</div>
          </div>
          <div class="field col-12 md:col-6 mb-4" v-if="checkIfAgent">
            <Field v-model="pesel" name="pesel" placeholder=" " type="text" class="form-control" :class="{'is-invalid': errors.pesel}" :validateOnInput="true"></Field>
            <span>{{ $t('FORMS.PLACEHOLDERS.PESEL') }}</span>
            <div class="invalid-feedback">{{ errors.pesel ? $t(errors.pesel) : '' }}</div>
          </div>
        </div>
        <div class="grid grid-nogutter justify-content-between">
          <Button :label="$t('GLOBALS.BUTTONS.BACK')" @click="goBack" type="submit" icon="pi pi-angle-left" icon-pos="left"></Button>
          <Button :label="checkIfAdmin ? useSelectedEmployee ? $t('GLOBALS.BUTTONS.SAVE') : $t('GLOBALS.BUTTONS.CREATE') : $t('GLOBALS.BUTTONS.NEXT')"
                  type="submit" icon="pi pi-angle-right" icon-pos="right"></Button>
        </div>
      </Form>
    </div>
  </div>
</template>

<script setup lang="ts">

import {ref, defineEmits, PropType, onMounted} from 'vue';
import {Form, Field} from 'vee-validate';
import * as Yup from 'yup';
import {FunctionalUser} from '@/components/forms/create-func-user/FunctionalUser';
import {RoleEnum} from '@/services/permissions/role-enum';
import {useRouter} from "vue-router";

import moment from "moment";
import {onBeforeMount} from "@vue/runtime-core";
import {useContextStore} from "@/store/context.store";
import {validateNIP, validatePESEL} from '@/utils/validators';
import {start} from "@popperjs/core";

onMounted(() => {
  console.log("props:" + props.passedRoleID + " " + checkIfAgent);
  console.log(props.selectedProviderId);
});

const router = useRouter();
const context = useContextStore();

const emit = defineEmits(['nextPage', 'prevPage', 'update:newFunctionalUser', 'complete']);
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
    type: String,
    required: true
  },
  selectedProviderId: {
  }
});

onBeforeMount(() => {
  if (props.useSelectedEmployee) {
    firstName.value = props.newFunctionalUser.firstName;
    lastName.value = props.newFunctionalUser.lastName;
    workStartDate.value = moment(new Date(props.newFunctionalUser.workStartDate)).format('YYYY-MM-DD');
    workEndDate.value = moment(new Date(props.newFunctionalUser.workEndDate)).format('YYYY-MM-DD');
    nip.value = props.newFunctionalUser.nip;
    pesel.value = props.newFunctionalUser.pesel;
    phone.value = props.newFunctionalUser.phone;
    email.value = props.newFunctionalUser.email;
  }
});
const checkIfSuperAdmin = (props.passedRoleID == RoleEnum.SUPER_ADMIN);
const checkIfFullAdmin = (props.passedRoleID == RoleEnum.ADMINISTRATOR_FULL);
const checkIfBasicAdmin = (props.passedRoleID == RoleEnum.ADMINISTRATOR_BASIC);
const checkIfAdmin = (checkIfBasicAdmin || checkIfFullAdmin || checkIfSuperAdmin);
const checkIfAgent = (props.passedRoleID == RoleEnum.AGENT || props.passedRoleID == RoleEnum.SUPER_AGENT);

const firstName = ref();
const lastName = ref();
const workStartDate = ref<string>((moment(new Date())).format('YYYY-MM-DD'));
const workEndDate = ref<string>((moment(new Date())).format('YYYY-MM-DD'));
const nip = ref();
const pesel = ref();
const phone = ref();
const email = ref();
let enableValidation = true;

const superAdminSchema = Yup.object().shape({
  email: Yup.string()
    .required('FORMS.VALIDATION_ERRORS.REQUIRED')
    .email('FORMS.VALIDATION_ERRORS.INVALID_MAIL'),
  firstName: Yup.string()
    .required('FORMS.VALIDATION_ERRORS.REQUIRED')
    .matches(/^[a-zA-ZąĄćĆęĘłŁńŃóÓśŚźŹżŻ\s]*$/, 'FORMS.VALIDATION_ERRORS.INVALID_FNAME'),
  lastName: Yup.string()
    .required('FORMS.VALIDATION_ERRORS.REQUIRED')
    .matches(/^[a-zA-ZąĄćĆęĘłŁńŃóÓśŚźŹżŻ_.-\s]*$/, 'FORMS.VALIDATION_ERRORS.INVALID_SNAME'),
  phone: Yup.number()
    .required('FORMS.VALIDATION_ERRORS.REQUIRED')
    .typeError('FORMS.VALIDATION_ERRORS.PHONE.NaN'),
});
const agentSchema = superAdminSchema.shape({
  workStartDate: Yup.string()
    .required('FORMS.VALIDATION_ERRORS.REQUIRED')
    .matches(/^\d{4}-(0[1-9]|1[012])-(0[1-9]|[12][0-9]|3[01])$/, 'FORMS.VALIDATION_ERRORS.DATE_FORMAT'),
  workEndDate: Yup.string()
    .required('FORMS.VALIDATION_ERRORS.REQUIRED')
    .matches(/^\d{4}-(0[1-9]|1[012])-(0[1-9]|[12][0-9]|3[01])$/, 'FORMS.VALIDATION_ERRORS.DATE_FORMAT'),
  nip: Yup.string()
    .required('FORMS.VALIDATION_ERRORS.REQUIRED')
    .test('validate-nip', 'FORMS.VALIDATION_ERRORS.INVALID_NIP',
      function (value) {
        if (value != undefined && context.allowExtendedValidation) {
          return validateNIP(value);
        } else {
          return true;
        }
      }),
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
});

function setValidation(roleEnum: RoleEnum) {
  if (enableValidation)
    switch (roleEnum) {
      case RoleEnum.SUPER_ADMIN:
        return superAdminSchema;
      case RoleEnum.AGENT:
      case RoleEnum.SUPER_AGENT:
        return agentSchema;
      default:
        return superAdminSchema;
    }
}

const goBack = () => {
  enableValidation = false;
  router.push(props.previousPath);
};
const nextPage = () => {
  console.log("workStartDate:" + workStartDate.value + " " + workEndDate.value);
  let updatedNewFunctionalUser: FunctionalUser = props.newFunctionalUser;
  if (workStartDate.value) {
    let dateStart = workStartDate.value.toString().split('-');
    console.log(dateStart)
    updatedNewFunctionalUser.workStartDate = new Date(Number(dateStart[0]), Number(dateStart[1]) , Number(dateStart[2]) );
  }
  if (workEndDate.value) {
    let dateEnd = workEndDate.value.toString().split('-');
    updatedNewFunctionalUser.workEndDate = new Date(Number(dateEnd[0]), Number(dateEnd[1]), Number(dateEnd[2]) );
  }

  updatedNewFunctionalUser.firstName = firstName?.value;
  updatedNewFunctionalUser.lastName = lastName?.value;
  updatedNewFunctionalUser.nip = nip?.value;
  updatedNewFunctionalUser.pesel = pesel?.value;
  updatedNewFunctionalUser.phone = phone?.value;
  updatedNewFunctionalUser.email = email.value;
  updatedNewFunctionalUser.login = email.value;  
  if(props.selectedProviderId){
    updatedNewFunctionalUser.providerId = useContextStore().selectedProvider.id;
  }
  else{
    updatedNewFunctionalUser.providerId = useContextStore().currentLoggedProvider.id;
  }  
  updatedNewFunctionalUser.roleId = parseInt(props.passedRoleID.toString());
  console.log(updatedNewFunctionalUser);
  emit('update:newFunctionalUser', updatedNewFunctionalUser);

  if (props.passedRoleID == RoleEnum.SUPER_ADMIN || props.passedRoleID == RoleEnum.ADMINISTRATOR_BASIC || props.passedRoleID == RoleEnum.ADMINISTRATOR_FULL) {
    emit('complete');
  } else {
    emit('nextPage', {pageIndex: 0});
  }

};
</script>

<style scoped>
button {
  margin-right: 0.4rem;
}
</style>
