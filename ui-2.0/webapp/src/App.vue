<template>
  <div class="page-container">
    <loading-screen/>
    <app-header :logo="themeLogo" v-if="useUserStore().isLoggedIn" />
    <div class="main-container">
      <div class="content">
        <router-view/>
      </div>
    </div>
    <app-footer class="footer"/>
    <app-sidebar/>
  </div>


</template>

<script setup lang="ts">
import LoadingScreen from './components/LoadingScreen.vue';
import AppHeader from './components/layout/AppHeader.vue';
import AppSidebar from './components/layout/AppSidebar.vue';
import {computed, inject, onMounted, ref} from 'vue';
import {LoggerService} from './services/logger/logger.service';
import {createRouterBeforeEach} from './router/before_each';
import ThemeService from './services/theme.service';

import "primevue/resources/themes/bootstrap4-light-blue/theme.css";
import "primeicons/primeicons.css";
import "primevue/resources/primevue.min.css";
import AppFooter from './components/layout/AppFooter.vue';
import { useUserStore} from './store/user.store';
import { useThemeStore} from './store/theme.store';
const logger = inject<LoggerService>('logger') as LoggerService;

const themeService = new ThemeService();

const themeColours = themeService.getUserTheme();

const themeLogo = computed(() => {
  themeService.getLogoURL();
  return useThemeStore().getLogoURL;
});

const changeColours = (propertyName, propertyValue) => {
  const rootElement = document.querySelector(':root') as HTMLElement;
  rootElement.style.setProperty(propertyName, propertyValue);
};

onMounted(() => {
  createRouterBeforeEach(logger);
  themeColours.forEach(colorVariable => {
    changeColours(colorVariable.propertyName, colorVariable.propertyValue);
  });
});

// useRouter().beforeEach((to, from, next) => {
//   logger.logToConsole(LogLevel.DEBUG, LogObjectTypeEnum.ROUTER, 'go form:' + JSON.stringify(from.name) + ' to:' + JSON.stringify(to.name));
//   if (to.matched.some(record => record.meta.requiresAuth)) {
//     if (useUserStore().isLoggedIn) {
//       next();
//       return;
//     }
//     logger.logToConsole(LogLevel.INFO, LogObjectTypeEnum.ROUTER, 'user not logged in, redirecting to login page');
//     next('/login');
//   } else {
//     logger.logToConsole(LogLevel.WARNING, LogObjectTypeEnum.ROUTER, 'requiresAuth not defined in router configuration');
//     next();
//   }
// });

</script>

<style lang="scss">
@import "./styles/app.scss";


</style>
