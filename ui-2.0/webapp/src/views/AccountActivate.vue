<template>
  <div class="login-container">
    <div class="whirls">
      <div class="login-data-fade">
      <div class="lang">
        <LangSwitch/>
      </div>
      <div class="justify">
        <div class="logo">
          <img src="../assets/logo-new.svg"/>
        </div>
        <div class="form-container">

        <Form @submit="complete" :validation-schema="schema" v-slot="{errors}"  >
          <span class="icon-input-grid">
            <Icon name="Lock"/>
              <Field v-model="password" :placeholder="$t('FORMS.PLACEHOLDERS.PASSWORD')" name="password" type="password" class="form-control" :class="{'is-invalid': errors.password}">
              </Field>
            </span>
              <div class="invalid-feedback">{{errors.password ? $t(errors.password) : ''}}</div>
            <span class="icon-input-grid">
            <Icon name="Lock"/>
              <Field v-model="passwordRetype"  :placeholder="$t('FORMS.PLACEHOLDERS.RETYPE_PASSWORD')" name="passwordRetype" type="password" class="form-control" :class="{'is-invalid': errors.passwordRetype}">
              </Field>
            </span>
              <div class="invalid-feedback">{{errors.passwordRetype ? $t(errors.passwordRetype) : ''}}</div>
              <Button class="secondary" style="width: 100%; justify-content: center;" :label="$t('GLOBALS.BUTTONS.ACTIVATE_ACCOUNT')" type="submit" ></Button>
        </Form>
      </div>
    </div>
    </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import {ref } from 'vue';
import { Form, Field, configure  } from 'vee-validate';
import * as Yup from 'yup';
import { useUserStore } from '@/store/user.store';
import { useRouter } from 'vue-router';
import { useToast } from 'vue-toastification';
import { useI18n } from 'vue-i18n';

const toast = useToast();
const router = useRouter();
const pattern = /^(?=.*[a-z])(?=.*[A-Z])(?=.*[0-9])(?=.*[!@#$%^&*])(?=.{8,})/;
const password = ref();
const passwordRetype = ref();
const currentPath = router.currentRoute.value;
const i18n = useI18n();

configure({
  validateOnBlur: true, // controls if `blur` events should trigger validation with `handleChange` handler
  validateOnChange: true, // controls if `change` events should trigger validation with `handleChange` handler
  validateOnInput: true, // controls if `input` events should trigger validation with `handleChange` handler
  validateOnModelUpdate: true, // controls if `update:modelValue` events should trigger validation with `handleChange` handler
});


const schema = Yup.object().shape({
  password: Yup.string().required('FORMS.VALIDATION_ERRORS.REQUIRED').matches(pattern,'FORMS.VALIDATION_ERRORS.PASSWORD_REGEX' ),
  passwordRetype: Yup.string().required('FORMS.VALIDATION_ERRORS.REQUIRED').oneOf([Yup.ref('password'), null], 'FORMS.VALIDATION_ERRORS.PASSWORD_MATCH')
});

const complete = () => {
  useUserStore().activateUser({newPassword: password.value, newPasswordRetype: passwordRetype.value}, currentPath.params.activateCode.toString(), onSuccess, onFail);
};

const onSuccess = () => {
  router.push('/login');
  toast.success(i18n.t('GLOBALS.TOASTS.PASSWORD_SET'));
};

const onFail = () => {
  toast.error(i18n.t('GLOBALS.TOASTS.ERROR_POST'));
};
</script>

<style lang="scss" scoped>
@import '../styles/variables.scss';

.login-container {
  width: 100%;
  height: 100vh;
  padding: 0;
  position: absolute;
  top: 0;
  left: 0;
  object-fit: cover;
  background: linear-gradient(135deg, $--main-color 0%, $--main-darker-color 100%);
  overflow: hidden;
  box-sizing: border-box;

  .whirls {
    position: relative;
    width: 100%;
    height: 100%;
    background-color: var(--main-color);
    background-image:  url("../assets/login-bg.svg");
    background-position: right center;
    -webkit-background-size: cover;
    -moz-background-size: cover;
    -o-background-size: cover;
    background-size: cover;
    background-blend-mode: color-dodge;
    overflow: hidden;
    box-sizing: border-box;
  }

  .login-data-fade {
    position: relative;
    width: 100%;
    height: 100%;
    overflow: hidden;
    padding: 8em 0 0 2em;
    background: $--main-color;
    background-image: linear-gradient(90deg, rgba(0,44,80,1) 0%, rgba(0,44,80,1) 30%, rgba(18,83,137,0) 100%);
    overflow: hidden;
    box-sizing: border-box;
  }
}


.padding {
  margin-bottom: 10px;
  padding-bottom: 10px;
}

.form-container {
  margin-top: 6em;
  margin-left: 5em;
  max-width: 600px;
  z-index: 1000;
}

.lang {
  position: fixed;
  z-index: 1000;
  right: 15px;
  top: 15px;
}


    .logo {
      max-width: 400px;
      text-align: center;
      overflow: hidden;
      z-index: 1000;

      img {
        height: 100%;
        width: auto;
        z-index: 1000;
        overflow: hidden;
      }
    }

  .colored-header {
    color: #2AFD88;
  }

  .float-right {
    float: right;
  }

  .clickable {
    cursor: pointer;
  }

  .blue-btn {
    background: $--main-darker-color !important;
  }

.icon-input-grid {
  display: grid;
  grid-template-columns: 40px 1fr;
  align-items: center;
  margin-bottom: 20px;

  svg {
    transform: translateY(-5px);
    // color: $brand-green;
    color: #2AFD88;
  }
}

@media only screen and (max-width: 768px) {
  .grid-container {
    grid-template-columns: 1fr;

    .grid-img {
      display: none;
    }
  }
}
</style>
