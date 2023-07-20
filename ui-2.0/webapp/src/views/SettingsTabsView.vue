<template>
  <div class="card">
    <TabView class="tabview-custom" ref="providersTab" lazy>
      <TabPanel>
        <template #header>
          <span> {{ $t('TABLES.PROVIDER_ADMINS.INFO') }} </span>
        </template>
        <ProviderDetailsView/>
      </TabPanel>
      <!-- <TabPanel>
        <template #header>
          <span>Empty Tab </span>
        </template>
      </TabPanel> -->
    </TabView>
  </div>
</template>

<script setup lang="ts">

import TabView from 'primevue/tabview';
import TabPanel from 'primevue/tabpanel';
import ProviderDetailsView from "@/views/ProviderDetailsView.vue";

</script>

<style scoped>

</style>
