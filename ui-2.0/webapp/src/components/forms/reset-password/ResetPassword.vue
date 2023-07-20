<template>
  <Form @submit="complete" :validation-schema="schema" v-slot="{errors}"  >
    <Field v-model="password" :placeholder="$t('FORMS.PLACEHOLDERS.NEW_PASSWORD')" name="password" type="password" class="form-control" :class="{'is-invalid': errors.password}">
    </Field>
    <div class="invalid-feedback">{{errors.password ? $t(errors.password) : ''}}</div>
    <Field v-model="passwordRetype"  :placeholder="$t('FORMS.PLACEHOLDERS.RETYPE_NEW_PASSWORD')" name="passwordRetype" type="password" class="form-control" :class="{'is-invalid': errors.passwordRetype}">
    </Field>
    <div class="invalid-feedback">{{errors.passwordRetype ? $t(errors.passwordRetype) : ''}}</div>
    <Button :label="$t('GLOBALS.BUTTONS.RESET_PASSWORD')" type="submit" ></Button>
  </Form>
</template>

<script setup lang="ts">
import {ref } from 'vue';
import { Form, Field, configure  } from 'vee-validate';
import * as Yup from 'yup';
import { useRouter } from 'vue-router';
import { useToast } from 'vue-toastification';
import { useI18n } from 'vue-i18n';

const toast = useToast();
const router = useRouter();
const pattern = /^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!@#$%^&*])(?=.{8,})/;
const password = ref();
const passwordRetype = ref();
const i18n = useI18n();

configure({
  validateOnBlur: true, // controls if `blur` events should trigger validation with `handleChange` handler
  validateOnChange: true, // controls if `change` events should trigger validation with `handleChange` handler
  validateOnInput: true, // controls if `input` events should trigger validation with `handleChange` handler
  validateOnModelUpdate: true, // controls if `update:modelValue` events should trigger validation with `handleChange` handler
});


const schema = Yup.object().shape({
  oldPassword: Yup.string().required('FORMS.VALIDATION_ERRORS.REQUIRED').matches(pattern,'FORMS.VALIDATION_ERRORS.PASSWORD_REGEX' ),
  password: Yup.string().required('FORMS.VALIDATION_ERRORS.REQUIRED').matches(pattern,'FORMS.VALIDATION_ERRORS.PASSWORD_REGEX' ),
  passwordRetype: Yup.string().required('FORMS.VALIDATION_ERRORS.REQUIRED').oneOf([Yup.ref('password'), null], 'FORMS.VALIDATION_ERRORS.PASSWORD_MATCH')
});

const complete = () => {
  console.log();
};

const onSuccess = () => {
  router.push('/home');
  toast.success(i18n.t('GLOBALS.TOASTS.PASSWORD_SET'));
};

const onFail = () => {
  toast.error(i18n.t('GLOBALS.TOASTS.ERROR_POST'));
};
</script>

<style scoped>

</style>
