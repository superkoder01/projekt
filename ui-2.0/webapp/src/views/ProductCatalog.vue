


<template>
<div>
  <div v-if="mainView" style="margin-top: 20px">
<!--    <v-breadcrumbs :items="breadcrumbs"></v-breadcrumbs>-->
    <h2>Katalog produktów</h2>
    <div class="products-container">
      <router-link :to="'/product_catalog/pricing'">
        <div class="product">
          <span>Energia elektryczna</span>
          <xml version="1.0" encoding="iso-8859-1"/>
          <svg version="1.1" id="Layer_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px" viewBox="0 0 400 400" style="enable-background:new 0 0 400 400;" xml:space="preserve">
          <g>	<polygon points="157.055,0 90.798,196.319 164.417,196.319 88.344,400 289.571,159.509 218.405,159.509 311.656,0"/></g><g></g><g></g><g></g><g></g><g></g><g></g><g></g><g></g><g></g><g></g><g></g><g></g><g></g><g></g><g></g></svg>
        </div>
      </router-link>
      <router-link :to="'/product_catalog/offer_draft'">
        <div class="product">
          <span>Szablony ofert</span>
          <xml version="1.0" encoding="iso-8859-1"/>
          <svg version="1.1" id="Layer_1" xmlns="http://www.w3.org/2000/svg" xmlns:xlink="http://www.w3.org/1999/xlink" x="0px" y="0px" viewBox="0 0 400 400" style="enable-background:new 0 0 400 400;" xml:space="preserve">
          <g>	<polygon points="157.055,0 90.798,196.319 164.417,196.319 88.344,400 289.571,159.509 218.405,159.509 311.656,0"/></g><g></g><g></g><g></g><g></g><g></g><g></g><g></g><g></g><g></g><g></g><g></g><g></g><g></g><g></g><g></g></svg>
        </div>
      </router-link>
    </div>

  </div>
  <router-view v-if="!mainView"></router-view>
</div>
</template>

<script setup lang="ts">
import { useRouter } from "vue-router";
import { computed } from "vue";

const router = useRouter();

const mainView = computed(() => {
  console.log(router.currentRoute.value);
  if(router.currentRoute.value.matched.length > 1) {
    return false;
  } else {
    return true;
  }
});
</script>

<style scoped lang="scss">
@import '../styles/variables.scss';

$a-tags: 'a, a:active, a:hover, a:visited';
.products-container {
  margin-top: 20px;
  display: flex;
  gap: 10px;

  #{$a-tags} {
    color: white;
    text-decoration: none;
  }
}

.product {
  width: 350px;
  height: 200px;
  border-radius: 10px;
  background-color: var(--main-color);
  color: white;
  display: grid;
  align-items: center;;
  grid-template-columns: 2fr 2fr;
  padding: 10px 10px 10px 20px;
  overflow: hidden;
  position: relative;

  &:hover {
    cursor: pointer;
    opacity: 0.8;
  }

  svg {
    transform: scale(1.5, 1.5);
    height: 100%;
    width: auto;
    fill: var(--secondary-color);
    position: absolute;
    right: -20px;
  }

  span {
    transform: translateY(50%);
    font-size: 1.4em;
    font-weight: bold;
  }
}
</style>
