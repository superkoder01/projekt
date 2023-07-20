<template>
  <div class="card">
    <TabMenu :model="views" />
    <router-view />
  </div>
</template>

<script setup lang="ts">
import TabMenu from "primevue/tabmenu";
import {computed, ref} from "vue";
import { useI18n } from "vue-i18n";

const i18n = useI18n();

const views = computed(() => [
  {
    label: i18n.t('SECTION_TITLES.INFO'),
    to: '/customers/customer_tabs'
  },
  {
    label: i18n.t('SECTION_TITLES.OFFERS'),
    to: '/customers/customer_tabs/offers'
  },
  {
    label: i18n.t('SECTION_TITLES.CONTRACTS'),
    to: '/customers/customer_tabs/contracts'
  },
  {
    label: i18n.t('SECTION_TITLES.INVOICES'),
    to: '/customers/customer_tabs/invoices'
  },
  {
    label: i18n.t('SECTION_TITLES.INSTALATIONS'),
    to: '/customers/customer_tabs/installations'
  }
]);

</script>

<style  scoped lang="scss">

</style>
