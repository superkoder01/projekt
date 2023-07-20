<template>
  <span>

    <DataTableWrapper :service="contractsDataTableService" :global-filter-fields="['payload.contractDetails.number']">

      <template v-slot:empty>{{ $t('CONTRACTS_VIEW.OTHER.NO_CONTRACTS_FOUND') }}</template>
      <template v-slot:columns>
        <Column field="payload.contractDetails.number" :header="$t('CONTRACTS_VIEW.TABLE_HEADERS.NUMBER')" sortable></Column>
        <Column field="payload.contractDetails.type" :header="$t('CONTRACTS_VIEW.TABLE_HEADERS.TYPE')"  sortable></Column>
        <Column field="payload.contractDetails.state" :header="$t('CONTRACTS_VIEW.TABLE_HEADERS.STATE')"  sortable>
          <template #body="{data}">
            <span class="status"
                  :style="{'background-color': getStateColor(data.payload.contractDetails.state)}">{{ $t('GLOBALS.STATUS.'+data.payload.contractDetails.state) }}</span>
          </template>
        </Column>
        <Column :expander="true"> </Column>
      </template>
      <template v-slot:expanded-columns="expandedData">
      <div style="padding-top: 15px; display: flex; flex-direction: column" class="container">
        <h5> {{ $t('CONTRACTS_VIEW.DETAILS.DETAILS') }} </h5>
        <span class="detailed-data-dark-mode" style="width: 100%">

          <span>
            <span><Icon name="FileText" /></span>
            <span>
              <div class="label-small">{{ $t("CONTRACTS_VIEW.DETAILS.NAME") }}:</div>
              {{ expandedData.expandedData.data.payload.contractDetails.title !== '' ? expandedData.expandedData.data.payload.contractDetails.title : '-' }}
            </span>
          </span>

          <span>
            <span><Icon name="FileText" /></span>
            <span>
              <div class="label-small">{{ $t("CONTRACTS_VIEW.DETAILS.NUMBER") }}:</div>
              {{ expandedData.expandedData.data.payload.contractDetails.number !== '' ? expandedData.expandedData.data.payload.contractDetails.number : '-' }}
            </span>
          </span>

          <span>
            <span><Icon name="Calendar" /></span>
            <span>
              <div class="label-small">{{ $t("CONTRACTS_VIEW.DETAILS.CREATION_DATE") }}:</div>
              {{ expandedData.expandedData.data.payload.contractDetails.creationDate !== '' ? expandedData.expandedData.data.payload.contractDetails.creationDate : '-' }}
            </span>
          </span>

          <span>
            <span><Icon name="FileText" /></span>
            <span>
              <div class="label-small">{{ $t("CONTRACTS_VIEW.DETAILS.TYPE") }}:</div>
              {{ expandedData.expandedData.data.payload.contractDetails.type !== '' ? expandedData.expandedData.data.payload.contractDetails.type : '-' }}
            </span>
          </span>

          <span>
            <span><Icon name="Calendar" /></span>
            <span>
              <div class="label-small">{{ $t("CONTRACTS_VIEW.DETAILS.START_DATE") }}:</div>
              {{ expandedData.expandedData.data.payload.conditions.startDate !== '' ? expandedData.expandedData.data.payload.conditions.startDate : '-' }}
            </span>
          </span>

          <span>
            <span><Icon name="Calendar" /></span>
            <span>
              <div class="label-small">{{ $t("CONTRACTS_VIEW.DETAILS.END_DATE") }}:</div>
              {{ expandedData.expandedData.data.payload.conditions.endDate !== '' ? expandedData.expandedData.data.payload.conditions.endDate : '-' }}
            </span>
          </span>

          <span>
            <span><Icon name="Calendar" /></span>
            <span>
              <div class="label-small">{{ $t("CONTRACTS_VIEW.DETAILS.ACTIVATION_DATE") }}:</div>
              {{ expandedData.expandedData.data.payload.conditions.signatureDate !== '' ? expandedData.expandedData.data.payload.conditions.signatureDate : '-' }}
            </span>
          </span>

          <span>
            <span><Icon name="FileText" /></span>
            <span>
              <div class="label-small">{{ $t("CONTRACTS_VIEW.DETAILS.STATUS") }}:</div>
              <Dropdown v-if="useSelectedCustomer && expandedData.expandedData.data.payload.contractDetails.state !== StatusEnum.ACCEPTED && expandedData.expandedData.data.payload.contractDetails.state !== StatusEnum.REJECTED" style="display: flex" v-model="selectedStatusOption" :placeholder="$t('GLOBALS.STATUS.'+expandedData.expandedData.data.payload.contractDetails.state)" :options="statusOptions" :change="changeStatus(expandedData.expandedData.data)" >
                     <template #option="slotProps">
                    <span>{{$t("GLOBALS.STATUS."+slotProps.option)}}</span>
                  </template>
               </Dropdown>
              <div v-if="!useSelectedCustomer || expandedData.expandedData.data.payload.contractDetails.state === StatusEnum.ACCEPTED || expandedData.expandedData.data.payload.contractDetails.state === StatusEnum.REJECTED  " style="color: white">{{$t('GLOBALS.STATUS.'+expandedData.expandedData.data.payload.contractDetails.state) }}</div>
            </span>
          </span>

          <span>
            <span><Icon name="Zap" /></span>
            <span>
              <div class="label-small">{{ $t("CONTRACTS_VIEW.DETAILS.PPE_CODE") }}:</div>
              {{ expandedData.expandedData.data.payload.serviceAccessPoints[0].sapCode !== '' ? expandedData.expandedData.data.payload.serviceAccessPoints[0].sapCode : '-' }}
            </span>
          </span>

          <span>
            <span><Icon name="Zap" /></span>
            <span>
              <div class="label-small">{{ $t("CONTRACTS_VIEW.DETAILS.PPE_ADDRESS") }}:</div>
              {{ expandedData.expandedData.data.payload.serviceAccessPoints[0].address !== '' ? expandedData.expandedData.data.payload.serviceAccessPoints[0].address : '-' }}
            </span>
          </span>

          <span v-if="expandedData.expandedData.data.payload.priceList[0]">
            <span><Icon name="Coins" /></span>
            <span>
              <div class="label-small">{{ $t("CONTRACTS_VIEW.DETAILS.PRICING_NAME") }}:</div>
              {{ expandedData.expandedData.data.payload.priceList[0].name !== '' ? expandedData.expandedData.data.payload.priceList[0].name : '-' }}
            </span>
          </span>

          <span v-if="expandedData.expandedData.data.payload.priceList[0]">
            <span><Icon name="Coins" /></span>
            <span>
              <div class="label-small">{{ $t("CONTRACTS_VIEW.DETAILS.ZONE_NAME") }}:</div>
              {{ expandedData.expandedData.data.payload.priceList[0].zones[0].name !== '' ? expandedData.expandedData.data.payload.priceList[0].zones[0].name : '-' }}
            </span>
          </span>

          <span v-if="expandedData.expandedData.data.payload.priceList[0]">
            <span><Icon name="Coins" /></span>
            <span>
              <div class="label-small">{{ $t("CONTRACTS_VIEW.DETAILS.COMMERCIAL_FEE") }}:</div>
              {{ expandedData.expandedData.data.payload.priceList[0].commercialFee.cost
                }}{{ expandedData.expandedData.data.payload.priceList[0].commercialFee.currency
                }}/{{ expandedData.expandedData.data.payload.priceList[0].commercialFee.calendarUnit !== '' ? $t('ENUMS.CALENDAR_UNIT_TYPE.' + expandedData.expandedData.data.payload.priceList[0].commercialFee.calendarUnit.toUpperCase()) : '' }}
            </span>
          </span>

          <span v-if="expandedData.expandedData.data.payload.repurchase.id">
            <span><Icon name="Coins" /></span>
            <span>
              <div class="label-small">{{ $t("CONTRACTS_VIEW.DETAILS.REPURCHASE_NAME") }}:</div>
              {{ expandedData.expandedData.data.payload.repurchase.name !== '' ? expandedData.expandedData.data.payload.repurchase.name : '-' }}
            </span>
          </span>

          <span v-if="expandedData.expandedData.data.payload.repurchase.id">
            <span><Icon name="Coins" /></span>
            <span>
              <div class="label-small">{{ $t("CONTRACTS_VIEW.DETAILS.REPURCHASE_PRICE") }}:</div>
              {{ expandedData.expandedData.data.payload.repurchase.price.cost
                }} PLN/{{ expandedData.expandedData.data.payload.repurchase.price.unit }}
            </span>
          </span>

          <span>
            <span><Icon name="Calendar" /></span>
            <span>
              <div class="label-small">{{ $t("CONTRACTS_VIEW.DETAILS.BILLING_PERIOD") }}:</div>
              {{ expandedData.expandedData.data.payload.conditions.billingPeriod.calendarUnit !== '' ? $t('ENUMS.CALENDAR_UNIT_TYPE.' + expandedData.expandedData.data.payload.conditions.billingPeriod.calendarUnit.toUpperCase())  : '-' }}
            </span>
          </span>
        </span>
        <Button style="align-self: start;" class="secondary" @click="saveToPdf(expandedData.expandedData.data)"> {{ $t('GLOBALS.BUTTONS.SAVE_TO_PDF') }} <Icon name="Download"></Icon></Button>
         <Button class="secondary" @click="editContract(expandedData.expandedData.data)" v-if="useSelectedCustomer && expandedData.expandedData.data.payload.contractDetails.state!==StatusEnum.ACCEPTED
                && expandedData.expandedData.data.payload.contractDetails.state!==StatusEnum.REJECTED"> {{$t('GLOBALS.BUTTONS.EDIT')}} </Button>
      </div>
      </template>

      <template v-slot:paginatorstart></template>

    </DataTableWrapper>
  </span>
</template>

<script setup lang="ts">
import DataTableWrapper from '@/components/layout/DataTableWrapper.vue';
import {DataTableService} from "@/services/dataTableService";
import {PagingModel} from "@/services/model/paging.model";
import {LocalSpinner} from "@/services/model/localSpinner";
import {DataHolder} from "@/models/data-holder";
import {useContractsStore} from "@/store/contracts.store";
import {useContextStore} from "@/store/context.store";
import {DefaultSortingModel} from "@/services/model/defaultSorting.model";
import factoryApi from "@/api/factory.api";
import {useToast} from "vue-toastification";
import { Contract, DateFormat, Offer } from "@/models/billing/billing";
import {ResponseError} from "@/models/request-response-api";
import { ref } from "vue";
import { StatusEnum } from "@/models/billing/enum/status.enum";
import { useOffersStore } from "@/store/offers.store";
import Dropdown from 'primevue/dropdown';
import { useI18n } from "vue-i18n";
import { formatSendDate } from "@/utils/date-formatter";
import {useRouter} from "vue-router";

const props = defineProps({
  expandedData: {
    type: Object
  },
  useSelectedCustomer: {
    type: Boolean,
    default: false
  }
});
const router = useRouter();
const toast = useToast();
const contractStore = useContractsStore();
const currentPath = router.currentRoute.value.path;

const statusOptions = ref([StatusEnum.DRAFT,StatusEnum.SENT,StatusEnum.DELIVERED, StatusEnum.FINAL, StatusEnum.REJECTED, StatusEnum.ACCEPTED]);
const selectedStatusOption = ref();


const contractsDataTableService: DataTableService<Contract> = {

  getDefaultSorting():DefaultSortingModel{
    return new DefaultSortingModel( "desc",  "payload.contractDetails.creationDate");
  },
  fetchListData(pagination: PagingModel, lockScreen: boolean, localSpinner: LocalSpinner | null): void {
    if (props.useSelectedCustomer) {
      const selectedCustomerId = useContextStore().selectedCustomer.id;
      contractStore.fetchContractsForCustomer(selectedCustomerId, lockScreen, localSpinner, pagination);
    } else {
      contractStore.fetchContracts(lockScreen, localSpinner, pagination);
    }
  },
  getListDataHolder(): DataHolder<Contract> {
    return contractStore.getListDataHolder;
  }
};

function saveToPdf(contract: Contract) {
  console.log('started saving');
  factoryApi.pdfConverterApi().downloadDocumentPdf(contract, true, null, pdfSuccess, pdfFailed, true);
}

function pdfSuccess(data:any):void{
  console.log("Contract pdf Success");
}

function pdfFailed(error : ResponseError):void{
  console.log("Contract pdf Failed: " + error.message );
  toast.error(error.message as string);
}

const editContract = (selectedContract: Contract) => {
  useContextStore().selectedContract = selectedContract;
  router.push(currentPath+"/edit_contract");
};

const changeStatus = (contract: Contract) => {
  if(selectedStatusOption.value != contract.payload.contractDetails.state && contract.id && selectedStatusOption.value){
    let updatedContract: Contract = {} as Contract;
    updatedContract.payload = contract.payload;
    updatedContract.header = contract.header;
    updatedContract.payload.contractDetails.state = selectedStatusOption.value;
    if(selectedStatusOption.value === StatusEnum.ACCEPTED) {
      updatedContract.payload.conditions.signatureDate = formatSendDate(new Date().toString(),DateFormat.SEND_DATE_FORMAT_WITH_TIME);
    }
    useContractsStore().updateContractById(contract.id,updatedContract, onSuccess, onFail);
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
  if (status === 'ACCEPTED') return '#ccffcc';
  return '#ff9999';
}
</script>

<style scoped lang="scss">

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

  @media screen and (max-width: 950px) {
    .detailed-data-dark-mode {
      grid-template-columns: 1fr !important;
    }
  }
</style>
