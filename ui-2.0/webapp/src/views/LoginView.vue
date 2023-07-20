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
        <!--        <v-form ref="form" v-model="valid" lazy-validation @submit.prevent="doLogin" :disabled="isFormDisabled">-->
        <div class="form-container">
          <span class="icon-input-grid">
            <Icon name="User"/>
              <InputText type="text" id="login" v-model="loginData.login" v-on:keypress.enter="goToPass"></InputText>
          </span>
          <span class="icon-input-grid">
            <Icon name="Lock"/>
            <InputText v-model="loginData.password" id="password" type="password" v-on:keypress.enter="doLogin"></InputText>
          </span>
          <span @click="goToForgotPassword" style="color: #2AFD88"
                class="colored-header float-right padding clickable">{{ $t("LOGIN.forgot_password") }}</span><br>
          <!--            <v-btn type="submit" class="btn&#45;&#45;green btn&#45;&#45;rounded right" color="primary"-->
          <!--                   style="max-width: 125px !important; background-color: #2AFD88 !important" block large-->
          <!--                   :loading="loading">{{ $t("LOGIN.log_in") }}-->
          <!--            </v-btn>-->
          <Button style="width: 100%; justify-content: center;" class="secondary" @click="doLogin">{{$t("LOGIN.log_in")}}</Button>
        </div>
        <!--        </v-form>-->
        <!--        <v-theme-provider :theme="{theme}">-->
        <!--          <v-btn @click="doTest()">test</v-btn>-->
        <!--        </v-theme-provider>-->
        <!--        <v-switch v-model="$vuetify" hint="This toggles the global state of the Vuetify theme" inset label="Vuetify Theme Dark" persistent-hint></v-switch>-->

      </div>
    </div>
    </div>
  </div>
</template>

<script setup lang="ts">

import {LoginData} from '../models/login-data';
import {onMounted, reactive} from 'vue';
import LangSwitch from '../components/lang/LangSwitch.vue';
import {useRouter} from 'vue-router';
import {useToast} from 'vue-toastification';
import {useUserStore} from "@/store/user.store";
import {PermissionsService} from "@/services/permissions/permissions.service";

const loginData: LoginData = reactive({login: '', password: ''});

const router = useRouter();
const toast = useToast();
const permissionService = new PermissionsService();

onMounted(() => {
  document.getElementById('login')?.focus();
});

function doLogin() {
  console.log('ooooooooo: login->' + loginData.login + ' password->' + loginData.password);
  useUserStore().login(loginData, onSuccessLogin, onFailLogin);
}
function goToPass(){
  document.getElementById('password')?.focus();
}

function goToForgotPassword(){
  console.log('goToForgotPassword');
}
const onSuccessLogin = () => {
  toast.success("Zalogowano poprawnie");
  router.push(permissionService.getAfterLoginPage());
};

const onFailLogin = () => {
  toast.error("Błąd podczas logowania");
};

</script>

<style scoped lang="scss">
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
