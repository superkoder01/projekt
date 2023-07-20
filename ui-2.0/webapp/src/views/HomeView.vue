<template>
  <div class="container">
    <h2>{{ $t("HOME_VIEW.RDN_PRICES") }}</h2>
    <div class="placeholder-work-in-progress">
      <Icon name="Wrench" /> <span> {{ $t("HOME_VIEW.SITE_UNDER_CONSTRUCTION") }} </span>
    </div>
    <h2>{{ $t("HOME_VIEW.REPORTS") }}</h2>
    <div class="placeholder-work-in-progress">
      <Icon name="Wrench" /> <span> {{ $t("HOME_VIEW.SITE_UNDER_CONSTRUCTION") }} </span>
    </div>
    <h2>{{ $t("HOME_VIEW.SHORTCUTS") }}</h2>
    <div class="shortcut-container">
      <v-container v-if="setVisibility(PagesEnum.PRODUCT_CATALOG)"> <router-link :to="'/product_catalog'"><Button class="btn"><Icon name="BookOpen"> </Icon> {{ $t('SECTION_TITLES.PRODUCT_CATALOG') }} </Button></router-link></v-container>
      <v-container v-if="setVisibility(PagesEnum.OFFERS)"> <router-link :to="'/offers'"><Button class="btn"> <Icon name="Gem"> </Icon> {{ $t('SECTION_TITLES.OFFERS') }} </Button></router-link></v-container>
      <v-container v-if="setVisibility(PagesEnum.CONTRACTS)"> <router-link :to="'/contracts'"><Button class="btn"><Icon name="Album"> </Icon>  {{ $t('SECTION_TITLES.CONTRACTS') }} </Button></router-link></v-container>
      <v-container v-if="setVisibility(PagesEnum.CUSTOMERS)"> <router-link :to="'/customers'"><Button class="btn"><Icon name="Users"> </Icon>  {{ $t('SECTION_TITLES.CUSTOMERS') }} </Button></router-link></v-container>
      <v-container v-if="setVisibility(PagesEnum.INVOICES)"> <router-link :to="'/invoices'"><Button class="btn"> <Icon name="Coins"> </Icon> {{ $t('SECTION_TITLES.INVOICES') }} </Button></router-link></v-container>
    </div>
  </div>
</template>

<script setup lang="ts">
import {Customer} from "@/models/customer";
import {computed } from "vue";
import { useRouter } from "vue-router";
import {useContextStore} from "@/store/context.store";
import { PagesEnum } from "@/services/permissions/pages-enum";
import { PermissionsService } from "@/services/permissions/permissions.service";
const router = useRouter();
const permissionsService = new PermissionsService();
const customer = computed(() => {
   return useContextStore().currentLoggedUserData;
}
);

function setVisibility(pageType: PagesEnum): boolean {
  return permissionsService.canView(pageType);
}

</script>

<style scoped lang="scss">
.placeholder-work-in-progress {
  display: flex;
  justify-content: center;
  align-items: center;
  margin-bottom: 2em;
  margin-top: 0.5em;
  svg {
    margin-right: 5px;
  }
}

.shortcut-container {
  display: flex;
  text-align: center;
  justify-content: center;
  flex-wrap: wrap;
  width: 100%;

  .btn {
    width: 220px;
  }
}

.personal-data-grid {
  display: grid;
  grid-template-columns: 40px 1fr 40px;
  max-width: 400px;

  span {
    height: 60px;
    position: relative;
    border-bottom: 1px solid #DFDFDF;
    display: flex;
    flex-direction: column;
    justify-content: center;


    .label-small {
      font-size: 14px;
      color: #ff0000;
    }

    .lucide-icon {
      color: #ff0000;
      box-sizing: content-box;
    }

    .edit-icon {
      cursor: pointer;
      border-radius: 10px;
      padding: 5px;

      &:hover {
        background-color: #ff0000;
        color: white;
      }
    }
  }
}
.borders {
  margin-top: 1rem;
}
</style>

