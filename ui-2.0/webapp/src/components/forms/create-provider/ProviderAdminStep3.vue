<template>
  <div class="card no-border m-3">
  <h5 class="card-header"> </h5>
    <div class="card-body">
      <Form @submit="complete" :validation-schema="schema" v-slot="{ errors }" >
        <div class="p-fluid formgrid grid ">
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="firstName" name="firstName" placeholder=" " type="text" class="form-control" :class="{'is-invalid': errors.firstName}"></Field>
            <span>{{$t('FORMS.PLACEHOLDERS.FIRST_NAME')}}</span>
            <div class="invalid-feedback">{{errors.firstName ? $t(errors.firstName): ''}}</div>
          </div>
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="lastName" name="lastName" placeholder=" " type="text" class="form-control" :class="{'is-invalid': errors.lastName}"></Field>
            <span>{{$t('FORMS.PLACEHOLDERS.LAST_NAME')}}</span>
            <div class="invalid-feedback">{{errors.lastName ? $t(errors.lastName): ''}}</div>
          </div>
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

        </div>
        <div class="grid grid-nogutter justify-content-between">
          <Button :label="$t('GLOBALS.BUTTONS.BACK')"  @click="prevPage"  icon="pi pi-angle-left"></Button>
          <Button :label="$t('GLOBALS.BUTTONS.CREATE')" type="submit" icon="pi pi-angle-right"></Button>
        </div>
      </Form>
    </div>
  </div>
</template>


<script setup lang="ts">
import { FunctionalUser } from '../create-func-user/FunctionalUser';
import { PropType, ref } from 'vue';
import { Form, Field } from "vee-validate";
import * as Yup from 'yup';
import { RoleEnum } from '@/services/permissions/role-enum';
import {onBeforeMount} from "@vue/runtime-core";

const emit = defineEmits(['complete', 'prevPage', 'update:newProviderAdmin']);

const props = defineProps({
    newProviderAdmin : {
        type: Object as PropType<FunctionalUser>,
        required: true
    }
});

const firstName = ref();
const lastName = ref();
const email = ref();
const phone = ref();


const schema = Yup.object().shape({
    firstName: Yup.string().required('FORMS.VALIDATION_ERRORS.REQUIRED'),
    lastName: Yup.string().required('FORMS.VALIDATION_ERRORS.REQUIRED'),
    email: Yup.string().required('FORMS.VALIDATION_ERRORS.REQUIRED'),
    phone: Yup.string().required('FORMS.VALIDATION_ERRORS.REQUIRED'),
});
const complete = () => {
    let updatedProviderAdmin : FunctionalUser = props.newProviderAdmin;
    updatedProviderAdmin.firstName = firstName.value;
    updatedProviderAdmin.lastName  = lastName.value;
    updatedProviderAdmin.email = email.value;
    updatedProviderAdmin.phone = phone.value;
    updatedProviderAdmin.password = ''; //must be set as empty string
    updatedProviderAdmin.roleId = RoleEnum.ADMINISTRATOR_FULL;
    updatedProviderAdmin.login = email.value;
    emit('update:newProviderAdmin', updatedProviderAdmin);
    emit('complete', {isAdminAdded: true});
};
const prevPage = () => {
  emit('prevPage', {pageIndex: 2});
};
</script>
