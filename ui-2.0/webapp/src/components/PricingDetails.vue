<template>
<div class="container">

  <span class="detailed-data" :class="darkMode ? 'detailed-data-dark-mode': 'detailed-data'" style="width: 100%">

        <span>
          <span><Icon name="None" /></span>
          <span>
            <div class="label-small">{{ $t("PRICING_VIEW.TABLE_HEADERS.ID") }}:</div>
            {{ props.pricing.id !== '' ?  props.pricing.id : "-"}}
          </span>
        </span>

        <span>
          <span><Icon name="None" /></span>
          <span>
            <div class="label-small">{{ $t("PRICING_VIEW.TABLE_HEADERS.NAME") }}:</div>
            {{ props.pricing.name }}
          </span>
        </span>

        <span>
          <span><Icon name="None" /></span>
          <span>
            <div class="label-small">{{ $t("PRICING_VIEW.TABLE_HEADERS.TYPE") }}:</div>
            {{ $t("ENUMS.SERVICE_TYPE." + toServiceTypeEnumKey(props.pricing.type)) }}
          </span>
        </span>

        <span>
          <span><Icon name="None" /></span>
          <span>
            <div class="label-small">{{ $t("PRICING_VIEW.TABLE_HEADERS.TARRIF_GROUP") }}:</div>
            {{ props.pricing.tariffGroup }}
          </span>
        </span>

        <!-- <span>
          <span><Icon name="None" /></span>
          <span>
            <div class="label-small">{{ $t("PRICING_VIEW.TABLE_HEADERS.OSD") }}:</div>
            {{ props.pricing.osd }}
          </span>
        </span> -->

      </span>

      <div v-if="props.pricing.type === ServiceTypeEnum.REPURCHASE">
          <div class="detail-row">

            <div>
              <div class="detail-header">{{ $t("PRICING_VIEW.TABLE_HEADERS.FIXED_PRICE.UNIT") }}</div>
              <div class="detail-item">{{ props.pricing.price.calendarUnit }}</div>
              </div>
            <div>
              <div class="detail-header">{{ $t("PRICING_VIEW.TABLE_HEADERS.FIXED_PRICE.COST") }}</div>
              <div class="detail-item">{{ props.pricing.price.cost }}</div>
            </div>
            <div>
              <div class="detail-header">{{ $t("PRICING_VIEW.TABLE_HEADERS.FIXED_PRICE.CURRENCY") }}</div>
            <div class="detail-item">{{ props.pricing.price.currency }}</div>
            </div>
          </div>
          </div>
      <div v-if="props.pricing.zones !== undefined">
          <span>
          <h4>{{ $t("PRICING_VIEW.TABLE_HEADERS.ZONES") }}</h4>
              <div class="detail-headers five-cols" :class="darkMode ? 'detail-headers-dark-mode' : 'detail-headers'">
                <div class="detail-header"> {{ $t("PRICING_VIEW.TABLE_HEADERS.ZONE.ID") }}</div>
                <div class="detail-header"> {{ $t("PRICING_VIEW.TABLE_HEADERS.ZONE.NAME") }}</div>
                <div class="detail-header">{{ $t("PRICING_VIEW.TABLE_HEADERS.ZONE.UNIT") }}</div>
                <div class="detail-header"> {{ $t("PRICING_VIEW.TABLE_HEADERS.ZONE.COST") }}</div>
                <div class="detail-header"> {{ $t("PRICING_VIEW.TABLE_HEADERS.ZONE.CURRENCY") }}</div>
              </div>
            <div class="detail-row five-cols" v-bind:key="zone.id"
                v-for="zone in props.pricing.zones">
              <div class="detail-item">{{ zone.id }}</div>
              <div class="detail-item">{{ zone.name }}</div>
              <div class="detail-item">{{ zone.unit }}</div>
              <div class="detail-item">{{ zone.cost }}</div>
              <div class="detail-item">{{ zone.currency }}</div>
            </div>
        </span>
      </div>
      <div v-if="props.pricing.commercialFee !== undefined">
          <span>
            <h4>{{ $t("PRICING_VIEW.TABLE_HEADERS.COMMERCIAL_FEE") }}</h4>
            <div class="detail-headers six-cols" :class="darkMode ? 'detail-headers-dark-mode' : 'detail-headers'">
              <div class="detail-header">{{ $t("PRICING_VIEW.TABLE_HEADERS.SINGLE_COMMERCIAL_FEE.FROM") }}</div>
              <div class="detail-header">{{ $t("PRICING_VIEW.TABLE_HEADERS.SINGLE_COMMERCIAL_FEE.TO") }}</div>
              <div class="detail-header">{{ $t("PRICING_VIEW.TABLE_HEADERS.SINGLE_COMMERCIAL_FEE.UNIT") }}</div>
              <div class="detail-header">{{$t("PRICING_VIEW.TABLE_HEADERS.SINGLE_COMMERCIAL_FEE.PRICE.CALENDAR_UNIT")}}</div>
              <div class="detail-header">{{ $t("PRICING_VIEW.TABLE_HEADERS.SINGLE_COMMERCIAL_FEE.PRICE.COST") }}</div>
              <div class="detail-header">{{$t("PRICING_VIEW.TABLE_HEADERS.SINGLE_COMMERCIAL_FEE.PRICE.CURRENCY") }}</div>
            </div>
            <div class="six-cols" :class="darkMode ? 'detail-row-dark-mode' : 'detail-row '" v-bind:key="item.id" v-for="item in props.pricing.commercialFee">
              <div class="detail-item">{{ item.from }}</div>
              <div class="detail-item">{{ item.to === undefined ? "-" : item.to }}</div>
              <div class="detail-item">{{ item.unit }}</div>
              <div class="detail-item">{{ item.price.calendarUnit === "" ? "-" : $t("ENUMS.CALENDAR_UNIT_TYPE." + toCalendarUnitTypeEnumKey(item.price.calendarUnit)) }}</div>
              <div class="detail-item">{{ item.price.cost === undefined ? "-" : item.price.cost }}</div>
              <div class="detail-item">{{ item.price.currency }}</div>
            </div>
          </span>
      </div>
</div>

</template>

<script setup lang="ts">

import {PropType} from "vue";
import {Pricing} from "@/models/billing/billing";
import {ServiceTypeEnum, toServiceTypeEnumKey} from "@/models/billing/enum/service-type.enum";
import {toCalendarUnitTypeEnumKey, DateFormat} from "@/models/billing/billing";

const props = defineProps({
  pricing: {
    type: Object as PropType<Pricing>,
    required: true
  },
  darkMode: {
    type: Boolean,
    required: true,
    default: false
  }
});

</script>

<style lang="scss" scoped>
@import '../styles/variables.scss';

h4 {
  margin-top: 10px;
}
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

.detail-row {
  display: grid;
  padding: 5px;

  &:nth-child(even) {
    background-color: rgba(228, 228, 228, 1);
  }
}

.detail-row-dark-mode {
  display: grid;
  padding: 5px;

  &:nth-child(even) {
    background-color: var(--main-darker-color);
  }
}

.detailed-data {
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
      -webkit-box-shadow: 0px 4px 6px 0px rgba(228, 228, 228, 1);
      -moz-box-shadow: 0px 4px 6px 0px rgba(228, 228, 228, 1);
      box-shadow: 0px 4px 6px 0px rgba(228, 228, 228, 1);
      padding: 10px 15px;
      padding-left: -10px;
      border-radius: 10px;
      display: flex;
      align-items: center;

      svg {
        filter: drop-shadow(0px 2px 2px var(--secondary-color));
      }

      span:first-child {
        width: 0px;
        display: none;
      }

      span:last-child {
        text-align: start;

        .label-small {
          font-weight: 600;
          font-size: 0.9em;
          color: var(--main-color);
        }
      }
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
        width: 0px;
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

@media screen and (max-width: 950px) {
    .detailed-data {
      grid-template-columns: 1fr !important;
    }
    .detailed-data-dark-mode {
      grid-template-columns: 1fr !important;
    }
  }
</style>
