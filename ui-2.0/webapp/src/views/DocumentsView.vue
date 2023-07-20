<template>
  <div class="card">

    <TabView class="tabview-custom" ref="tabview4">
      <TabPanel>
        <template #header>
          <span> {{ $t('SECTION_TITLES.OFFERS') }}</span>
        </template>
        <offers-view/>
      </TabPanel>
      <TabPanel>
        <template #header>
          <span> {{ $t('SECTION_TITLES.CONTRACTS') }}</span>
        </template>
        <contracts-view/>
      </TabPanel>

    </TabView>
  </div>
</template>

<script setup lang="ts">
import TabView from 'primevue/tabview';
import TabPanel from 'primevue/tabpanel';
import OffersView from "@/views/OffersListView.vue";
import ContractsView from "@/views/ContractsListView.vue";</script>

<style scoped lang="scss">

</style>
