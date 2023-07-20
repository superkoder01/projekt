<template>
  <DataTableWrapper :service="offersDraftDataTableService" :global-filter-fields="['payload.offerDetails.number']">

    <template v-slot:empty>{{ $t('OFFERS_VIEW.OTHER.NO_OFFERS_FOUND') }}</template>
    <template v-slot:columns>
      <Column field="payload.offerDetails.title" :header="$t('OFFERS_VIEW.TABLE_HEADERS.OFFER_NAME')" sortable/>
      <Column field="payload.offerDetails.type" :header="$t('OFFERS_VIEW.TABLE_HEADERS.OFFER_TYPE')" sortable/>
      <Column field="payload.offerDetails.creationDate" :header="$t('OFFERS_VIEW.TABLE_HEADERS.OFFER_CREATION_DATE')" sortable></Column>
      <Column field="payload.offerDetails.tariffGroup" :header="$t('OFFERS_VIEW.TABLE_HEADERS.TARRIF')" sortable/>
      <Column field="payload.offerDetails.agreementType" :header="$t('OFFERS_VIEW.TABLE_HEADERS.CONTRACT_TYPE')" sortable/>
      <Column :expander="true" headerStyle="width: 3rem"/>
    </template>
    <template v-slot:expanded-columns="expandedData">
      <div class="container">

        <div class="detailed-data-dark-mode" style="width: 100%">
          <span>
            <span><Icon name="FileText" /></span>
            <span>
              <div class="label-small">{{ $t("OFFERS_VIEW.TABLE_HEADERS.OFFER_NAME") }}:</div>
              {{ expandedData.expandedData.data.payload.offerDetails.title !== '' ? expandedData.expandedData.data.payload.offerDetails.title : '-' }}
            </span>
          </span>

          <span>
            <span><Icon name="FileText" /></span>
            <span>
              <div class="label-small">{{ $t("OFFERS_VIEW.TABLE_HEADERS.PRICING_NAME") }}:</div>
              {{ expandedData.expandedData.data.payload.offerDetails.name !== '' ? expandedData.expandedData.data.payload.offerDetails.name : '-' }}
            </span>
          </span>

          <span>
            <span><Icon name="Calendar" /></span>
            <span>
              <div class="label-small">{{ $t("OFFERS_VIEW.TABLE_HEADERS.OFFER_CREATION_DATE") }}:</div>
              {{ expandedData.expandedData.data.payload.offerDetails.creationDate !== '' ? expandedData.expandedData.data.payload.offerDetails.creationDate : '-' }}
            </span>
          </span>

          <span>
            <span><Icon name="FileText" /></span>
            <span>
              <div class="label-small">{{ $t("OFFERS_VIEW.TABLE_HEADERS.OFFER_TYPE") }}:</div>
              {{ expandedData.expandedData.data.payload.offerDetails.type !== '' ? expandedData.expandedData.data.payload.offerDetails.type : '-' }}
            </span>
          </span>

          <span>
            <span><Icon name="FileText" /></span>
            <span>
              <div class="label-small">{{ $t("OFFERS_VIEW.TABLE_HEADERS.REPURCHASE_NAME") }}:</div>
              {{ expandedData.expandedData.data.payload.repurchase.name !== '' ? expandedData.expandedData.data.payload.repurchase.name : '-' }}
            </span>
          </span>

          <span>
            <span><Icon name="Coins" /></span>
            <span>
              <div class="label-small">{{ $t("OFFERS_VIEW.TABLE_HEADERS.REPURCHASE_PRICE") }}:</div>
              {{ expandedData.expandedData.data.payload.repurchase.price.cost !== '' ? expandedData.expandedData.data.payload.repurchase.price.cost : '-' }}PLN/
              {{ expandedData.expandedData.data.payload.repurchase.price.unit !== '' ? expandedData.expandedData.data.payload.repurchase.price.unit : '-' }}
            </span>
          </span>

          <span>
            <span><Icon name="Calendar" /></span>
            <span>
              <div class="label-small">{{ $t("OFFERS_VIEW.TABLE_HEADERS.BILLING_PERIOD") }}:</div>
              {{ expandedData.expandedData.data.payload.conditions.billingPeriod.calendarUnit !== '' ? $t("OFFERS_VIEW.TABLE_HEADERS.BILLING_PERIOD_" + expandedData.expandedData.data.payload.conditions.billingPeriod.calendarUnit) : '-' }}
            </span>
          </span>


        </div>
      </div>
<!--      <span>-->
<!--          <div>{{ $t("OFFERS_VIEW.TABLE_HEADERS.ZONE_PRICE") }}</div>-->
<!--          {{ expandedData.expandedData.data.Payload.PriceList.Zones[0].Name }}-->
<!--        </span>-->
<!--      <span>-->
<!--          <div>{{ $t("OFFERS_VIEW.TABLE_HEADERS.COMMERCIAL_FEE") }}</div>-->
<!--          {{ expandedData.expandedData.data.Payload.PriceList.CommercialFee }}-->
<!--        </span>-->

      <div v-if="expandedData.expandedData.data.payload.priceList.zones !== undefined">
        <span>

        <div class="detail-header">{{ $t("PRICING_VIEW.TABLE_HEADERS.ZONES") }}</div>

            <div class="detail-headers detail-headers-dark-mode five-cols">
              <div class="detail-header"> {{ $t("PRICING_VIEW.TABLE_HEADERS.ZONE.ID") }}</div>
              <div class="detail-header"> {{ $t("PRICING_VIEW.TABLE_HEADERS.ZONE.NAME") }}</div>
              <div class="detail-header">{{ $t("PRICING_VIEW.TABLE_HEADERS.ZONE.UNIT") }}</div>
              <div class="detail-header"> {{ $t("PRICING_VIEW.TABLE_HEADERS.ZONE.COST") }}</div>
              <div class="detail-header"> {{ $t("PRICING_VIEW.TABLE_HEADERS.ZONE.CURRENCY") }}</div>
            </div>


          <div class="detail-row" v-bind:key="zone.id"
               v-for="zone in expandedData.expandedData.data.payload.priceList.zones">

            <div class="detail-item">{{ zone.id }}</div>
            <div class="detail-item">{{ zone.name }}</div>
            <div class="detail-item">{{ zone.unit }}</div>
            <div class="detail-item">{{ zone.cost }}</div>
            <div class="detail-item">{{ zone.currency }}</div>

          </div>

      </span>
      </div>

      <div v-if="expandedData.expandedData.data.payload.priceList.commercialFee !== undefined">
        <span>
          <div class="detail-header">{{ $t("OFFERS_VIEW.TABLE_HEADERS.COMMERCIAL_FEE") }}</div>

          <div class="detail-headers detail-headers-dark-mode six-cols">
             <div class="detail-header">{{ $t("PRICING_VIEW.TABLE_HEADERS.SINGLE_COMMERCIAL_FEE.FROM") }}</div>
             <div class="detail-header">{{ $t("PRICING_VIEW.TABLE_HEADERS.SINGLE_COMMERCIAL_FEE.TO") }}</div>
             <div class="detail-header">{{ $t("PRICING_VIEW.TABLE_HEADERS.SINGLE_COMMERCIAL_FEE.UNIT") }}</div>
             <div class="detail-header">{{$t("PRICING_VIEW.TABLE_HEADERS.SINGLE_COMMERCIAL_FEE.PRICE.CALENDAR_UNIT")}}</div>
             <div class="detail-header">{{ $t("PRICING_VIEW.TABLE_HEADERS.SINGLE_COMMERCIAL_FEE.PRICE.COST") }}</div>
             <div class="detail-header">{{$t("PRICING_VIEW.TABLE_HEADERS.SINGLE_COMMERCIAL_FEE.PRICE.CURRENCY") }}</div>
          </div>

          <div class="detail-row" v-bind:key="item.id" v-for="item in expandedData.expandedData.data.payload.priceList.commercialFee">
            <div class="detail-item">{{ item.from }}</div>
            <div class="detail-item">{{ item.to === undefined ? "-" : item.to }}</div>
            <div class="detail-item">{{ item.unit }}</div>
            <div class="detail-item">{{ item.price.calendarUnit === "" ? "-" : $t("ENUMS.CALENDAR_UNIT_TYPE." + toCalendarUnitTypeEnumKey(item.price.calendarUnit)) }}</div>
            <div class="detail-item">{{ item.price.cost === undefined ? "-" : item.price.cost }}</div>
            <div class="detail-item">{{ item.price.currency }}</div>
          </div>
         </span>
      </div>

<!--                  <div>-->
<!--                    <span>-->
<!--                      <div>PAYLOAD</div>-->
<!--                      {{ expandedData.expandedData.data.Payload }}-->
<!--                    </span>-->
<!--                  </div>-->

    </template>

    <template v-slot:paginatorstart></template>

  </DataTableWrapper>
  <Button  v-if="permissionService.canAccessFeature(FeatureEnum.MANAGE_OFFER_DRAFTS)" :label="$t('GLOBALS.BUTTONS.ADD_OFFER')" @click="router.push('/product_catalog/offer_draft/new_offer_draft')"
          icon="pi pi-angle-right" icon-pos="right"></Button>
</template>

<script setup lang="ts">

import {onBeforeMount} from "@vue/runtime-core";
import {useDraftOffersStore} from "@/store/draftOffers.store";
import {DataTableService} from "@/services/dataTableService";
import {PagingModel} from "@/services/model/paging.model";
import {LocalSpinner} from "@/services/model/localSpinner";
import {DataHolder} from "@/models/data-holder";
import DataTableWrapper from '@/components/layout/DataTableWrapper.vue';
import {OfferDraft, toCalendarUnitTypeEnumKey} from "@/models/billing/billing";

import router from "@/router";
import {DefaultSortingModel} from "@/services/model/defaultSorting.model";
import {PermissionsService} from "@/services/permissions/permissions.service";
import { FeatureEnum } from "@/services/permissions/feature-enum";


const draftOffers = useDraftOffersStore();
const permissionService = new PermissionsService();

onBeforeMount(() => {
  useDraftOffersStore().fetchDraftOffers(true, null);
});

const offersDraftDataTableService: DataTableService<OfferDraft> = {
  getDefaultSorting():DefaultSortingModel{
    return new DefaultSortingModel( "desc",  "payload.offerDetails.creationDate");
  },
  fetchListData(pagination: PagingModel, lockScreen: boolean, localSpinner: LocalSpinner | null): void {
    draftOffers.fetchDraftOffers(lockScreen, localSpinner, pagination);
  },
  getListDataHolder(): DataHolder<OfferDraft> {
    return draftOffers.getListDataHolder;
  }
};

</script>

<style scoped lang="scss">
.five-cols {
  grid-template-columns: repeat(5, 1fr);
}

.six-cols {
  grid-template-columns: repeat(6, 1fr);
}

.detail-headers {
  display: grid;
  justify-content: space-between;
  background: var(--main-color);
  color: var(--header-text-color);
  border-radius: 10px 10px 0 0;
  padding: 5px;

  .detail-header {
    text-align: center;
  }
}

.detail-headers-dark-mode {
  background: var(--main-lighter-color);
}

.detail-item {
  text-align: center;
}

.detail-row-dark-mode {
  display: grid;
  padding: 5px;

  &:nth-child(even) {
    background-color: var(--main-darker-color);
  }
}

.detailed-data-dark-mode {
    display: grid;
    grid-template-columns: 1fr 1fr;
    grid-column-gap: 10px;
    grid-row-gap: 10px;

    span {
      border: 0;
    }

    & > span {
      width: 100%;
      min-height: 65px;
      box-sizing: border-box;
      -webkit-box-shadow: 0px 4px 6px 0px rgb(29, 29, 29);
      -moz-box-shadow: 0px 4px 6px 0px rgba(29, 29, 29);
      box-shadow: 0px 4px 6px 0px rgba(29, 29, 29);
      padding: 10px 15px;
      padding-left: -10px;
      border-radius: 10px;
      display: flex;
      align-items: center;
      background-color: var(--main-lighter-color);

      svg {
        filter: drop-shadow(0px 2px 2px var(--secondary-color));
      }

      span:first-child {
        width: 50px;
      }

      span:last-child {
        text-align: start;

        .label-small {
          font-weight: 600;
          font-size: 0.9em;
          color: var(--secondary-color);
        }
      }
    }
  }

.container {
  display: flex;
  flex-wrap: wrap;
}

Button {
  margin-top: 2rem;
}

@media screen and (max-width: 950px) {
    .detailed-data-dark-mode {
      grid-template-columns: 1fr !important;
    }
  }

</style>

<style>
div.p-dropdown.p-component.p-inputwrapper.dropdown > div.p-dropdown-trigger > span {
  border-bottom: none;
}
</style>
