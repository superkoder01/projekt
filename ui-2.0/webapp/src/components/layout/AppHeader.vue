<template>
  <!--  <div class="navbar navbar-expand-lg background bg-primary">-->
  <!--    <div class="container-fluid">-->
  <!--      <Image class="navbar-brand" :src="require('../../assets/logo-new.svg')"/>-->
  <!--      <button class="navbar-toggler" type="button" data-bs-toggle="collapse" data-bs-target="#navbarNavAltMarkup"-->
  <!--              aria-controls="navbarNavAltMarkup" aria-expanded="false" aria-label="Toggle navigation">-->
  <!--        <span class="navbar-toggler-icon"></span>-->
  <!--      </button>-->
  <!--      &lt;!&ndash;      <AutoLogOut/>&ndash;&gt;-->
  <!--      <div class="collapse navbar-collapse" id="navbarNavAltMarkup">-->
  <!--        <div class="navbar-nav">-->
  <!--          <span>AppHeader</span>-->
  <!--          <LangSwitch/>-->
  <!--          <Button></Button>-->
  <!--        </div>-->
  <!--      </div>-->
  <!--    </div>-->
  <!--  </div>-->
  <nav class="navbar navbar-expand-lg navbar-dark background">
    <div class="navbar-container">
      <div class="container-fluid d-flex justify-content-between">
        <span class="d-flex" style="align-items: center">
          <Image @click="goToHomePage" class="navbar-brand" :src="props.logo" alt="Image" height="36" />

            <div class="navbar-nav left">
              <i class="pi pi-user me-3" style="font-size: 1.5rem;" />
              <span><span style="font-weight:bold">{{ contextStore.currentLoggedUserData.firstName }} {{ contextStore.currentLoggedUserData.lastName }}</span><span v-if="userStore.isLoggedIn"> -  {{ $t(currentRole) }} </span> </span>
              <i class="pi pi-sun ms-4" style="font-size: 1.5rem; padding-right: 6px;" />
              <span style="padding-right: 6px;">{{ String(today.getDate()).padStart(2, '0') }}-{{ String(today.getMonth() + 1).padStart(2, '0') }}-{{today.getFullYear()}} </span>
            </div>

          <div v-if="showRoleChanger==='true'">
            <RoleChanger/>
          </div>
          <span class="text"> {{$t('GLOBALS.LAYOUT.PROVIDER')}}: {{ contextStore.loggedProvider.name }}</span>
        </span>
          <div class="navbar-nav" style="align-items: center">
  <!--          <div class="nav-link">-->

  <!--          </div>-->

            <LangSwitch class="nav-link mx-1"/>
            <span style="transform: translateY(-4px)">
              <Button icon="pi pi-bell" class="nav-link p-button-rounded p-button-text p-button-lg mx-1" ></Button>
            </span>

            <Button icon=" pi pi-power-off" style="transform: translateY(-4px)" class="nav-link p-button-rounded p-button-text p-button-lg mx-1"  @click="logOut()"></Button>

          </div>
          <div class="mobile-header">
            <LangSwitch class="nav-link mx-1" style="transform: translateY(7px)"/>
            <div class="switch-mobile" @focusout="closeDropdown" tabindex="0" @click="toggleDropdown()">
<!--    <font-awesome-icon icon="globe" class="icon"/>-->

            <span class="dropdown-btn">
              <Icon name="Menu"/>
            </span>
              <span class="options-container" v-if="dropdown">
                <span class="option">
                  <span class="text"> {{ contextStore.currentLoggedUserData.firstName }} {{ contextStore.currentLoggedUserData.lastName }} </span>
                </span>
                <span class="option">
                  <span class="text" v-if="userStore.isLoggedIn">{{$t('GLOBALS.ROLES.' + RoleEnum[userStore.currentRole])}}</span>
                </span>

                <span class="option">
                  <span class="icon-flex">
                    <Icon name="Bell"/>
                    <Icon v-if="userStore.isLoggedIn" @click="logOut()" name="Power"/>

                    <!-- <Button icon=" pi pi-power-off" style="transform: translateY(-4px)" class="nav-link p-button-rounded p-button-text p-button-lg mx-1"  @click="logOut()"></Button> -->
                  </span>
                </span>

              </span>

            </div>
            </div>
      </div>
        <div class="bottom-container">
          <!-- <span class="breadcrumb">Home / Jan Kowalski</span>
          <AutoLogOut class="right" style="top: 7vh;" /> -->
          <h2>{{$t("SECTION_TITLES." + currentRouteName?.toUpperCase())}}</h2>
          <breadcrumbs-component />
        </div>
        <span style="position: absolute; right: -10px; bottom: 15px" class="p-input-icon-left p-input-icon-right mx-5 desktop-search-bar" >
                <i class="pi pi-times-circle" style="transform: translateY(8px)" @click="globalFilter.clearFilter()"/>
                <InputText type="text" v-model="globalFilter.filter" :placeholder="$t('GLOBALS.PLACEHOLDERS.SEARCH')" />
                <i class="pi" style="transform: translateY(8px)" :class="{'pi-search': !globalFilter.isLoading, 'pi-spin pi-spinner': globalFilter.isLoading}"/>
        </span>
      <div class="container-fluid mobile-search">
        <span class="p-input-icon-left p-input-icon-right mx-5 " >
              <i class="pi pi-times-circle" style="transform: translateY(8px)" @click="globalFilter.clearFilter()"/>
              <InputText type="text" v-model="globalFilter.filter" :placeholder="$t('GLOBALS.PLACEHOLDERS.SEARCH')" />
              <i class="pi" style="transform: translateY(8px)" :class="{'pi-search': !globalFilter.isLoading, 'pi-spin pi-spinner': globalFilter.isLoading}"/>
            </span>
      </div>
    </div>
  </nav>

</template>

<script setup lang="ts">
import LangSwitch from '../../components/lang/LangSwitch.vue';
import BreadcrumbsComponent from '../features/BreadcrumbsComponent.vue';

import { useRouter } from 'vue-router';
import {useGlobalFilterStore} from "../../store/global-filter.store";
import { computed, onMounted, ref } from 'vue';
import {useUserStore} from "@/store/user.store";
import RoleChanger from "@/components/layout/RoleChanger.vue";
import {RoleEnum} from "@/services/permissions/role-enum";
import { useContextStore } from '@/store/context.store';
import { useThemeStore } from '@/store/theme.store';
import { PermissionsService } from '@/services/permissions/permissions.service';


const permissonsService = new PermissionsService();
const router = useRouter();
const userStore = useUserStore();
const theme = useThemeStore();
const dropdown = ref(false);
const today = new Date();
const contextStore = useContextStore();

const props = defineProps({
  logo: {
    type: String,
    required: true,
    default: 'data:image/gif;base64,R0lGODlhAQABAAAAACH5BAEKAAEALAAAAAABAAEAAAICTAEAOw=='
  }
});
const toggleDropdown = () => {
  dropdown.value = !dropdown.value;
};
const showRoleChanger = process.env.VUE_APP_SHOW_ROLE_CHANGER;
const closeDropdown = () => {
  dropdown.value = false;
};

const currentRouteName = computed(() => {
  console.log(router.currentRoute.value);
  return router.currentRoute.value.name;
});

const globalFilter = useGlobalFilterStore();
const currentRole = computed(() => "GLOBALS.ROLES."+RoleEnum[userStore.currentRole]);
console.log(currentRole);
function logOut(){
  router.push("/login");
  useUserStore().loggedOut();
}
function goToHomePage(){
  router.push(permissonsService.getAfterLoginPage());
}
</script>

<style scoped lang="scss">
@import '../../styles/app.scss';

.navbar-brand:hover{
  cursor: pointer;
}

.navbar-dark .navbar-nav .nav-link {
  color: rgba(255,255,255,.8) !important;
}

.navbar-dark .navbar-nav .nav-link:focus, .navbar-dark .navbar-nav .nav-link:hover {
  color: var(--secondary-color) !important;
}
</style>
