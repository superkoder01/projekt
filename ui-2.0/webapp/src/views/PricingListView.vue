<template>
  <DataTableWrapper :service="pricingDataTableService" :global-filter-fields="['']">
    <template v-slot:empty>{{ $t('PRICING_VIEW.OTHER.NO_PRICING_FOUND') }}</template>
    <template v-slot:columns>
      <Column field="payload.name" :header="$t('PRICING_VIEW.TABLE_HEADERS.NAME')" sortable></Column>
      <Column field="payload.type" :header="$t('PRICING_VIEW.TABLE_HEADERS.TYPE')" sortable>
        <template #body="{data}">
          <span class="status">{{ $t("ENUMS.SERVICE_TYPE." + toServiceTypeEnumKey(data.payload.type)) }}</span>
        </template>
      </Column>
      <Column field="payload.tariffGroup" :header="$t('PRICING_VIEW.TABLE_HEADERS.TARRIF_GROUP')" sortable></Column>

      <Column :expander="true" headerStyle="width: 3rem"/>
    </template>

    <template v-slot:expanded-columns="expandedData">
      <PricingDetails :darkMode="true" :pricing="expandedData.expandedData.data.payload"/>
    </template>

<!--    <template v-slot:expanded-columns="expandedData">-->

<!--      <span>-->
<!--          <div>{{ $t("PRICING_VIEW.TABLE_HEADERS.ID") }} </div>-->
<!--          {{ expandedData.expandedData.data.id }}-->
<!--        </span>-->

<!--      <span>-->
<!--          <div>{{ $t("PRICING_VIEW.TABLE_HEADERS.NAME") }} </div>-->
<!--          {{ expandedData.expandedData.data.payload.name }}-->
<!--        </span>-->

<!--      <span>-->
<!--          <div>{{ $t("PRICING_VIEW.TABLE_HEADERS.TYPE") }} </div>-->
<!--         {{ $t("ENUMS.SERVICE_TYPE." + toServiceTypeEnumKey(expandedData.expandedData.data.payload.type)) }}-->
<!--        </span>-->
<!--      <span>-->
<!--          <div>{{ $t("PRICING_VIEW.TABLE_HEADERS.TARRIF_GROUP") }}</div>-->
<!--          {{ expandedData.expandedData.data.payload.tariffGroup }}-->
<!--        </span>-->

<!--      <span>-->
<!--          <div>{{ $t("PRICING_VIEW.TABLE_HEADERS.START_DATE") }} </div>-->
<!--        {{ formatDate(expandedData.expandedData.data.payload.startDate, DateFormat.DATE_FORMAT, "-") }}-->
<!--        </span>-->

<!--      <span>-->
<!--          <div>{{ $t("PRICING_VIEW.TABLE_HEADERS.END_DATE") }} </div>-->
<!--        {{ formatDate(expandedData.expandedData.data.payload.endDate, DateFormat.DATE_FORMAT, "-") }}-->
<!--        </span>-->

<!--      <span>-->
<!--          <div>{{ $t("PRICING_VIEW.TABLE_HEADERS.OSD") }}</div>-->
<!--        {{ expandedData.expandedData.data.payload.osd }}-->
<!--        </span>-->

<!--      <span v-if="expandedData.expandedData.data.payload.type === ServiceTypeEnum.REPURCHASE">-->
<!--&lt;!&ndash;          <div>{{ $t("PRICING_VIEW.TABLE_HEADERS.PRICE") }}</div>&ndash;&gt;-->
<!--        <div class="detail-row">-->

<!--          <div>-->
<!--            <div class="detail-header">{{ $t("PRICING_VIEW.TABLE_HEADERS.FIXED_PRICE.UNIT") }}</div>-->
<!--            <div class="detail-item">{{ expandedData.expandedData.data.payload.price.calendarUnit }}</div>-->
<!--            </div>-->
<!--          <div>-->
<!--            <div class="detail-header">{{ $t("PRICING_VIEW.TABLE_HEADERS.FIXED_PRICE.COST") }}</div>-->
<!--            <div class="detail-item">{{ expandedData.expandedData.data.payload.price.cost }}</div>-->
<!--          </div>-->
<!--          <div>-->
<!--            <div class="detail-header">{{ $t("PRICING_VIEW.TABLE_HEADERS.FIXED_PRICE.CURRENCY") }}</div>-->
<!--           <div class="detail-item">{{ expandedData.expandedData.data.payload.price.currency }}</div>-->
<!--          </div>-->

<!--        </div>-->

<!--        </span>-->

<!--      <div v-if="expandedData.expandedData.data.payload.zones !== undefined">-->
<!--        <span>-->

<!--        <div class="detail-header">{{ $t("PRICING_VIEW.TABLE_HEADERS.ZONES") }}</div>-->

<!--            <div class="detail-headers">-->
<!--              <div class="detail-header"> {{ $t("PRICING_VIEW.TABLE_HEADERS.ZONE.ID") }}</div>-->
<!--              <div class="detail-header"> {{ $t("PRICING_VIEW.TABLE_HEADERS.ZONE.NAME") }}</div>-->
<!--              <div class="detail-header">{{ $t("PRICING_VIEW.TABLE_HEADERS.ZONE.UNIT") }}</div>-->
<!--              <div class="detail-header"> {{ $t("PRICING_VIEW.TABLE_HEADERS.ZONE.COST") }}</div>-->
<!--              <div class="detail-header"> {{ $t("PRICING_VIEW.TABLE_HEADERS.ZONE.CURRENCY") }}</div>-->
<!--            </div>-->


<!--          <div class="detail-row" v-bind:key="zone.id"-->
<!--               v-for="zone in expandedData.expandedData.data.payload.zones">-->

<!--            <div class="detail-item">{{ zone.id }}</div>-->
<!--            <div class="detail-item">{{ zone.name }}</div>-->
<!--            <div class="detail-item">{{ zone.unit }}</div>-->
<!--            <div class="detail-item">{{ zone.cost }}</div>-->
<!--            <div class="detail-item">{{ zone.currency }}</div>-->

<!--          </div>-->

<!--      </span>-->
<!--      </div>-->

<!--      <div v-if="expandedData.expandedData.data.payload.commercialFee !== undefined">-->
<!--        <span>-->
<!--          <div class="detail-header">{{ $t("PRICING_VIEW.TABLE_HEADERS.COMMERCIAL_FEE") }}</div>-->

<!--          <div class="detail-headers">-->
<!--             <div class="detail-header">{{ $t("PRICING_VIEW.TABLE_HEADERS.SINGLE_COMMERCIAL_FEE.FROM") }}</div>-->
<!--             <div class="detail-header">{{ $t("PRICING_VIEW.TABLE_HEADERS.SINGLE_COMMERCIAL_FEE.TO") }}</div>-->
<!--             <div class="detail-header">{{ $t("PRICING_VIEW.TABLE_HEADERS.SINGLE_COMMERCIAL_FEE.UNIT") }}</div>-->
<!--             <div class="detail-header">{{$t("PRICING_VIEW.TABLE_HEADERS.SINGLE_COMMERCIAL_FEE.PRICE.CALENDAR_UNIT")}}</div>-->
<!--             <div class="detail-header">{{ $t("PRICING_VIEW.TABLE_HEADERS.SINGLE_COMMERCIAL_FEE.PRICE.COST") }}</div>-->
<!--             <div class="detail-header">{{$t("PRICING_VIEW.TABLE_HEADERS.SINGLE_COMMERCIAL_FEE.PRICE.CURRENCY") }}</div>-->
<!--          </div>-->

<!--          <div class="detail-row" v-bind:key="item.id" v-for="item in expandedData.expandedData.data.payload.commercialFee">-->
<!--            <div class="detail-item">{{ item.from }}</div>-->
<!--            <div class="detail-item">{{ item.to === undefined ? "-" : item.To }}</div>-->
<!--            <div class="detail-item">{{ item.unit }}</div>-->
<!--            <div class="detail-item">{{ item.price.calendarUnit === "" ? "-" : $t("ENUMS.CALENDAR_UNIT_TYPE." + toCalendarUnitTypeEnumKey(item.price.calendarUnit)) }}</div>-->
<!--            <div class="detail-item">{{ item.price.cost === undefined ? "-" : item.price.cost }}</div>-->
<!--            <div class="detail-item">{{ item.price.currency }}</div>-->
<!--          </div>-->
<!--         </span>-->
<!--      </div>-->


<!--&lt;!&ndash;            <div>&ndash;&gt;-->
<!--&lt;!&ndash;              <span>&ndash;&gt;-->
<!--&lt;!&ndash;                <div>PAYLOAD</div>&ndash;&gt;-->
<!--&lt;!&ndash;                {{ expandedData.expandedData.data.Payload }}&ndash;&gt;-->
<!--&lt;!&ndash;              </span>&ndash;&gt;-->
<!--&lt;!&ndash;            </div>&ndash;&gt;-->


<!--    </template>-->
  </DataTableWrapper>


  <Button v-if="permissionService.canAccessFeature(FeatureEnum.MANAGE_PRICING)" :label="$t('GLOBALS.BUTTONS.ADD_PRICING')" @click="router.push('/product_catalog/pricing/new_pricing')"
          icon="pi pi-angle-right" icon-pos="right"></Button>
</template>

<script setup lang="ts">

import DataTableWrapper from '@/components/layout/DataTableWrapper.vue';
import {useRouter} from "vue-router";
import {DataTableService} from "@/services/dataTableService";
import {PagingModel} from "@/services/model/paging.model";
import {LocalSpinner} from "@/services/model/localSpinner";
import {DataHolder} from "@/models/data-holder";
import {usePricingStore} from "@/store/pricing.store";
import {ServiceTypeEnum, toServiceTypeEnumKey} from "@/models/billing/enum/service-type.enum";
import PricingDetails from "@/components/PricingDetails.vue";
import {DefaultSortingModel} from "@/services/model/defaultSorting.model";
import {PermissionsService} from "@/services/permissions/permissions.service";
import { FeatureEnum } from "@/services/permissions/feature-enum";
import {Pricing} from "@/models/billing/billing";

const router = useRouter();
const permissionService = new PermissionsService();

const pricingDataTableService: DataTableService<Pricing> = {
  getDefaultSorting():DefaultSortingModel{
    return new DefaultSortingModel( "",  "");
  },
  fetchListData(pagination: PagingModel, lockScreen: boolean, localSpinner: LocalSpinner | null): void {
    usePricingStore().fetchPricing(lockScreen, localSpinner, pagination);
  },
  getListDataHolder(): DataHolder<Pricing> {
    return usePricingStore().getListDataHolder;
  }
};

function test(st: string): ServiceTypeEnum {
  return st as ServiceTypeEnum;
}

</script>

<style scoped>

</style>
