<template>
  <div class="card">
    <TabMenu :model="views" />
    <router-view v-bind:provider="provider" />
  </div>
</template>

<script setup lang="ts">
import { useProvidersStore } from '@/store/providers.store';
import TabMenu from "primevue/tabmenu";
import { computed, ref } from "vue";
import { useI18n } from "vue-i18n";


const provider = computed(() => useProvidersStore().selectedProviderData);
const i18n = useI18n();
const views = ref([
  {
    label: i18n.t('TABLES.PROVIDER_ADMINS.INFO'),
    to: '/providers/provider_tabs'
  },
  {
    label: i18n.t('TABLES.PROVIDER_ADMINS.ADMINS'),
    to: '/providers/provider_tabs/admins'
  },
]);

</script>

<style scoped>

</style>
