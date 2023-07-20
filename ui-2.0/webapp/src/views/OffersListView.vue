<template>
    <DataTableWrapper :service="offersDataTableService" :global-filter-fields="['Payload.OfferDetails.Number']">

     <template v-slot:empty> {{$t('OFFERS_VIEW.OTHER.NO_OFFERS_FOUND')}} </template>
      <template v-slot:columns>
        <Column field="payload.offerDetails.title" :header="$t('OFFERS_VIEW.TABLE_HEADERS.OFFER_NAME')" sortable ></Column>
        <Column field="payload.offerDetails.number" :header="$t('OFFERS_VIEW.TABLE_HEADERS.OFFER_NUMBER')" sortable ></Column>
        <Column field="payload.offerDetails.type" :header="$t('OFFERS_VIEW.TABLE_HEADERS.OFFER_TYPE')" sortable ></Column>
        <Column field="payload.offerDetails.status" :header="$t('OFFERS_VIEW.TABLE_HEADERS.OFFER_STATE')" sortable >
          <template #body="{data}">
            <span class="status" :style="{'background-color': getStateColor(data.payload.offerDetails.status)}">{{$t("GLOBALS.STATUS."+data.payload.offerDetails.status)}}</span>
          </template>
        </Column>
        <Column :expander="true" headerStyle="width: 3rem" />
      </template>
      <template v-slot:expanded-columns="expandedData">
      <div style="padding-top: 15px" class="container">
          <h5> {{ $t("OFFERS_VIEW.DETAILS.DOCUMENT_DETAILS") }}</h5>
          <span class="detailed-data-dark-mode" style="width: 100%">

            <span>
              <span><Icon name="FileText" /></span>
              <span>
                <div class="label-small">{{ $t("OFFERS_VIEW.DETAILS.OFFER_NAME") }}:</div>
                {{ expandedData.expandedData.data.payload.offerDetails.title !== '' ? expandedData.expandedData.data.payload.offerDetails.title : '-' }}
              </span>
            </span>

            <span>
              <span><Icon name="FileText" /></span>
              <span>
                <div class="label-small">{{ $t("OFFERS_VIEW.DETAILS.NUMBER") }}:</div>
                {{ expandedData.expandedData.data.payload.offerDetails.number !== '' ? expandedData.expandedData.data.payload.offerDetails.number : '-' }}
              </span>
            </span>

            <span>
              <span><Icon name="Calendar" /></span>
              <span>
                <div class="label-small">{{ $t("OFFERS_VIEW.DETAILS.CREATION_DATE") }}:</div>
                {{ expandedData.expandedData.data.payload.offerDetails.creationDate !== '' ? expandedData.expandedData.data.payload.offerDetails.creationDate : '-' }}
              </span>
            </span>

            <span>
              <span><Icon name="FileText" /></span>
              <span>
                <div class="label-small">{{ $t("OFFERS_VIEW.DETAILS.TYPE") }}:</div>
                {{ expandedData.expandedData.data.payload.offerDetails.type !== '' ? expandedData.expandedData.data.payload.offerDetails.type : '-' }}
              </span>
            </span>

            <span>
              <span><Icon name="FileText" /></span>
              <span>
                <div class="label-small">{{ $t("OFFERS_VIEW.TABLE_HEADERS.OFFER_STATE") }}:</div>
                <Dropdown style="display: flex" class="dropdown" v-if="useSelectedCustomer && expandedData.expandedData.data.payload.offerDetails.status !== StatusEnum.ACCEPTED
                   && expandedData.expandedData.data.payload.offerDetails.status !== StatusEnum.REJECTED" v-model="selectedStatusOption" :placeholder="$t('GLOBALS.STATUS.'+expandedData.expandedData.data.payload.offerDetails.status)" :options="statusOptions" :change="changeStatus(expandedData.expandedData.data)" >
                  <template #option="slotProps">
                    <span>{{$t("GLOBALS.STATUS."+slotProps.option)}}</span>
                  </template>
                </Dropdown>
              <div v-if="!useSelectedCustomer || expandedData.expandedData.data.payload.offerDetails.status === StatusEnum.ACCEPTED || expandedData.expandedData.data.payload.offerDetails.status === StatusEnum.REJECTED  " style="color: white">{{$t('GLOBALS.STATUS.'+expandedData.expandedData.data.payload.offerDetails.status) }}</div>
              </span>
            </span>

            <span v-if="expandedData.expandedData.data.payload.priceList[0].zones != undefined">
              <span><Icon name="FileText" /></span>
              <span>
                <div class="label-small">{{ $t("OFFERS_VIEW.TABLE_HEADERS.PRICING_NAME") }}:</div>
                {{ expandedData.expandedData.data.payload.priceList[0].name !== '' ? expandedData.expandedData.data.payload.priceList[0].name : '-' }}
              </span>
            </span>

            <span v-if="expandedData.expandedData.data.payload.priceList[0].zones != undefined">
              <span><Icon name="FileText" /></span>
              <span>
                <div class="label-small">{{ $t("OFFERS_VIEW.TABLE_HEADERS.ZONE_PRICE") }}:</div>
                {{ expandedData.expandedData.data.payload.priceList[0].zones[0].name !== '' ? expandedData.expandedData.data.payload.priceList[0].zones[0].name : '-' }}
              </span>
            </span>

            <span v-if="expandedData.expandedData.data.payload.repurchase.id">
              <span><Icon name="Zap" /></span>
              <span>
                <div class="label-small">{{ $t("OFFERS_VIEW.TABLE_HEADERS.REPURCHASE_NAME") }}:</div>
                {{ expandedData.expandedData.data.payload.repurchase.name !== '' ? expandedData.expandedData.data.payload.repurchase.name : '-' }}
              </span>
            </span>

            <span v-if="expandedData.expandedData.data.payload.repurchase.id">
              <span><Icon name="Zap" /></span>
              <span>
                <div class="label-small">{{ $t("OFFERS_VIEW.TABLE_HEADERS.REPURCHASE_PRICE") }}:</div>
                {{ expandedData.expandedData.data.payload.repurchase.price.cost !== '' ? expandedData.expandedData.data.payload.repurchase.price.cost : '-' }} PLN/
                {{ expandedData.expandedData.data.payload.repurchase.price.unit !== '' ? expandedData.expandedData.data.payload.repurchase.price.unit : '-' }}
              </span>
            </span>

            <span v-if="expandedData.expandedData.data.payload.repurchase.id">
              <span><Icon name="Zap" /></span>
              <span>
                <div class="label-small">{{ $t("OFFERS_VIEW.TABLE_HEADERS.BILLING_PERIOD") }}:</div>
                {{ $t("OFFERS_VIEW.TABLE_HEADERS.BILLING_PERIOD_" + expandedData.expandedData.data.payload.conditions.billingPeriod.calendarUnit) }}
              </span>
            </span>

          </span>




        <div v-if="expandedData.expandedData.data.payload.priceList[0].commercialFee !== undefined">
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

          <div class="detail-row" v-bind:key="item.id" v-for="item in expandedData.expandedData.data.payload.priceList[0].commercialFee">
            <div class="detail-item">{{ item.from }}</div>
            <div class="detail-item">{{ item.to === undefined ? "-" : item.to }}</div>
            <div class="detail-item">{{ item.unit }}</div>
            <div class="detail-item">{{ item.price.calendarUnit === "" ? "-" : $t("ENUMS.CALENDAR_UNIT_TYPE." + toCalendarUnitTypeEnumKey(item.price.calendarUnit)) }}</div>
            <div class="detail-item">{{ item.price.cost === undefined ? "-" : item.price.cost }}</div>
            <div class="detail-item">{{ item.price.currency }}</div>
          </div>
         </span>
        </div>

        <Button @click="saveToPdf(expandedData.expandedData.data)" class="secondary"> {{ $t('GLOBALS.BUTTONS.SAVE_TO_PDF') }} <Icon name="Download"></Icon></Button>

        <Button class="secondary" @click="createContract(expandedData)" v-if="useSelectedCustomer && expandedData.expandedData.data.payload.offerDetails.status===StatusEnum.ACCEPTED"> {{$t('GLOBALS.BUTTONS.CREATE_CONTRACT')}} </Button>
        <Button class="secondary" @click="editOffer(expandedData.expandedData.data)" v-if="useSelectedCustomer && expandedData.expandedData.data.payload.offerDetails.status!==StatusEnum.ACCEPTED
                && expandedData.expandedData.data.payload.offerDetails.status!==StatusEnum.REJECTED"> {{$t('GLOBALS.BUTTONS.EDIT')}} </Button>
      </div>
      </template>

      <template v-slot:paginatorstart ></template>
    </DataTableWrapper>
  <template v-if="useSelectedCustomer">
    <Button :label="$t('GLOBALS.BUTTONS.ADD_NEW_OFFER')" @click="router.push('/customers/customer_tabs/offers/new_offer')"
            icon="pi pi-angle-right" icon-pos="right"></Button>
  </template>
</template>

<script setup lang="ts">
import DataTableWrapper from '@/components/layout/DataTableWrapper.vue';
import {useOffersStore} from "@/store/offers.store";
import {DataTableService} from "@/services/dataTableService";
import {PagingModel} from "@/services/model/paging.model";
import {LocalSpinner} from "@/services/model/localSpinner";
import {DataHolder} from "@/models/data-holder";
import {useContextStore} from "@/store/context.store";
import {useRouter} from "vue-router";
import {DefaultSortingModel} from "@/services/model/defaultSorting.model";
import {useToast} from "vue-toastification";
import { Contract, Offer } from "@/models/billing/billing";
import factoryApi from "@/api/factory.api";
import {ResponseError} from "@/models/request-response-api";
import { ref } from "vue";
import Dropdown from 'primevue/dropdown';
import { useI18n } from "vue-i18n";
import { StatusEnum } from "@/models/billing/enum/status.enum";
import {toCalendarUnitTypeEnumKey} from "@/models/billing/billing";

const toast = useToast();
const i18n = useI18n();
const offersStore = useOffersStore();
const router = useRouter();

const props = defineProps({
  useSelectedCustomer:{
    type: Boolean,
    default:false
  }
});
const currentPath = router.currentRoute.value.path;

const statusOptions = [StatusEnum.DRAFT,StatusEnum.SENT,StatusEnum.DELIVERED, StatusEnum.FINAL, StatusEnum.REJECTED, StatusEnum.ACCEPTED];
const selectedStatusOption = ref();

const offersDataTableService:DataTableService<Offer> = {
  getDefaultSorting():DefaultSortingModel{
    return new DefaultSortingModel( "desc",  "payload.offerDetails.creationDate");
  },
  fetchListData(pagination: PagingModel, lockScreen: boolean, localSpinner: LocalSpinner | null): void {
    if(props.useSelectedCustomer) {
      const selectedCustomerId = useContextStore().selectedCustomer.id;
      offersStore.fetchOffersForSelectedCustomer(selectedCustomerId, lockScreen, localSpinner, pagination);
    } else {
      offersStore.fetchOffers( lockScreen, localSpinner, pagination);
    }
  },
  getListDataHolder(): DataHolder<Offer> {
    return  offersStore.getListDataHolder;
  }
};

function saveToPdf(contract: Contract) {
  console.log('started saving');
  factoryApi.pdfConverterApi().downloadDocumentPdf(contract, true, null, pdfSuccess, pdfFailed, true);
}

function pdfSuccess(data:any):void{
  console.log("Offer pdf Success");
}

function pdfFailed(error : ResponseError):void{
  console.log("Offer pdf Failed: " + error.message );
  toast.error(error.message as string);
}

const createContract = (selectedOffer: any) => {
  useContextStore().selectedOffer = selectedOffer.expandedData.data;
  router.push('/customers/customer_tabs/offers/new_contract');
};

const editOffer = (selectedOffer: Offer) => {
  useContextStore().selectedOffer = selectedOffer;
  router.push(currentPath+"/edit_offer");
};

const changeStatus = (offer: Offer) => {
  if(selectedStatusOption.value != offer.payload.offerDetails.status && offer.id && selectedStatusOption.value){
    let updatedOffer: Offer = {} as Offer;
    updatedOffer.payload = offer.payload;
    updatedOffer.header = offer.header;
    updatedOffer.payload.offerDetails.status = selectedStatusOption.value;
    useOffersStore().updateOfferById(offer.id,updatedOffer, onSuccess, onFail);
    selectedStatusOption.value = "";
  }

};

const onSuccess = () => {
  toast.success("Zmieniono status");
};

const onFail = () => {
  toast.error("Nie udało się zmienić statusu");
};

function getStateColor(status: string) {
  if(status === 'ACCEPTED') return '#ccffcc';
  return '#ff9999';
}

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
