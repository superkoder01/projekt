<template>
  <div class="footer">
    <span @click="router.push('/faq')">{{ $t('SECTION_TITLES.FAQ') }}</span>
    <span @click="router.push('/help')">{{ $t('SECTION_TITLES.HELP') }}</span>
    <span @click="router.push('/terms_conditions')">{{ $t('SECTION_TITLES.TERMS_CONDITIONS') }}</span>
    <span @click="router.push('/privacy_policy')">{{ $t('SECTION_TITLES.PRIVACY_POLICY') }}</span>
    <span class="right"> {{ $t('GLOBALS.LAYOUT.VERSION')}} {{ version }} </span>
    <span @click="goToPage"  class="right">Go to</span>
  </div>
</template>

<script setup lang="ts">
import { useRouter } from 'vue-router';
const router = useRouter();
const version = process.env.VUE_APP_VERSION;
function goToPage() {
  window.open('https://google.com/', '_blank')?.focus();
}
</script>

<style scoped lang="scss">
.footer {
  width: 100%;
  min-height: 50px;
  background-color: white;

  padding: 20px 0 20px 8%;

  span {
    float:left;
    margin-left:15px;
  }
  span:hover {
    cursor: pointer;
  }
  .right {
    float: right;
    padding-right: 20px;
  }
}
</style>
