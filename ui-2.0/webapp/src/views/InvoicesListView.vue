<template>
  <DataTableWrapper :service="invoicesDataTableService" :global-filter-fields="['city']">

    <template v-slot:empty>{{ $t('INVOICES_VIEW.OTHER.NO_INVOICES_FOUND' ) }} </template>
    <template v-slot:columns>
      <Column field="payload.invoiceDetails.number" :header="$t('INVOICES_VIEW.TABLE_HEADERS.NUMBER')"  sortable></Column>
      <Column field="payload.invoiceDetails.billingStartDt" :header="$t('INVOICES_VIEW.TABLE_HEADERS.BILLING_START_DATE')" sortable></Column>
      <Column field="payload.iaymentDetails.payDueDt" :header="$t('INVOICES_VIEW.TABLE_HEADERS.PAY_DUE_DATE')" sortable></Column>
      <Column field="payload.invoiceDetails.iissueDt" :header="$t('INVOICES_VIEW.TABLE_HEADERS.ISSUE_DATE')" sortable></Column>
      <Column field="payload.invoiceDetails.status" :header="$t('INVOICES_VIEW.TABLE_HEADERS.STATUS')" sortable>
        <template #body="{data}">
          <span class="status"
                :style="{'background-color': getStateColor(data.payload.invoiceDetails.status)}">{{ data.payload.invoiceDetails.status }}</span>
        </template>
      </Column>
      <Column :expander="true" headerStyle="width: 3rem"> </Column>
    </template>
    <template v-slot:expanded-columns="expandedData">
    <div class="container">
      <div>
      <h5> {{ $t('INVOICES_VIEW.DETAILS.DETAILS' ) }} </h5>
        <div class="icon-holder">
          <Icon name="FileText"></Icon>
          <span>
            <div>{{ $t('INVOICES_VIEW.DETAILS.TYPE' ) }} </div>
            {{ expandedData.expandedData.data.payload.invoiceDetails.type}}
          </span>
        </div>
        <div class="icon-holder">
          <Icon name="FileText"></Icon>
          <span>
            <div>{{ $t('INVOICES_VIEW.DETAILS.NUMBER' ) }} </div>
            {{ expandedData.expandedData.data.payload.invoiceDetails.number }}
          </span>
        </div>
        <div class="icon-holder">
          <Icon name="FileText"></Icon>
          <span>
            <div>{{ $t('INVOICES_VIEW.DETAILS.CATEGORY' ) }} </div>
            {{ expandedData.expandedData.data.payload.invoiceDetails.catg }}
          </span>
        </div>
        <div class="icon-holder">
          <Icon name="FileText"></Icon>
          <span>
            <div>{{ $t('INVOICES_VIEW.DETAILS.CREATION_DATE' ) }} </div>
            {{ expandedData.expandedData.data.payload.invoiceDetails.issueDt }}
          </span>
        </div>
        <div class="icon-holder">
          <Icon name="FileText"></Icon>
          <span>
            <div>{{ $t('INVOICES_VIEW.DETAILS.BILLING_PERIOD' ) }} </div>
            {{ expandedData.expandedData.data.payload.invoiceDetails.billingStartDt }} -  {{ expandedData.expandedData.data.payload.invoiceDetails.billingEndDt }}
          </span>
        </div>
      </div>

      <div>
      <h5> {{ $t('INVOICES_VIEW.DETAILS.SELLER' ) }} </h5>
        <div class="icon-holder">
          <Icon name="Zap"></Icon>
          <span>
            <div>{{ $t('INVOICES_VIEW.DETAILS.SELLER_NAME' ) }} </div>
            {{ expandedData.expandedData.data.payload.sellerDetails.displayName }}
          </span>
        </div>
        <div class="icon-holder">
          <Icon name="Zap"></Icon>
          <span>
            <div>{{ $t('INVOICES_VIEW.DETAILS.STREET' ) }} </div>
            {{ expandedData.expandedData.data.payload.sellerDetails.contact.address.street }}
          </span>
        </div>
        <div class="icon-holder">
          <Icon name="Zap"></Icon>
          <span>
            <div>{{ $t('INVOICES_VIEW.DETAILS.POSTAL_CODE' ) }} </div>
            {{ expandedData.expandedData.data.payload.sellerDetails.contact.address.postCode }}
          </span>
        </div>
        <div class="icon-holder">
          <Icon name="Zap"></Icon>
          <span>
            <div>{{ $t('INVOICES_VIEW.DETAILS.CITY' ) }} </div>
            {{ expandedData.expandedData.data.payload.sellerDetails.contact.address.city }}
          </span>
        </div>
        <div class="icon-holder">
          <Icon name="Zap"></Icon>
          <span>
            <div>{{ $t('INVOICES_VIEW.DETAILS.EMAIL' ) }} </div>
            {{ expandedData.expandedData.data.payload.sellerDetails.contact.email }}
          </span>
        </div>
      </div>



      <div>
      <h5> {{ $t('INVOICES_VIEW.DETAILS.PAYMENT_DETAILS' ) }}  </h5>
        <div class="icon-holder">
          <Icon name="Banknote"></Icon>
          <span>
            <div>{{ $t('INVOICES_VIEW.DETAILS.TITLE' ) }} </div>
            {{ expandedData.expandedData.data.payload.paymentDetails.paymentTitle }}
          </span>
        </div>
        <div class="icon-holder">
          <Icon name="Banknote"></Icon>
          <span>
            <div>{{ $t('INVOICES_VIEW.DETAILS.PAYMENT_DUE_DATE' ) }} </div>
            {{ expandedData.expandedData.data.payload.paymentDetails.paymentDueDt }}
          </span>
        </div>
        <div class="icon-holder">
          <Icon name="Banknote"></Icon>
          <span>
            <div>{{ $t('INVOICES_VIEW.DETAILS.ACCOUNT_NUMBER' ) }} </div>
            {{ expandedData.expandedData.data.payload.paymentDetails.bankDetails.account }}
          </span>
        </div>
      </div>
    </div>
    <Button @click="saveToPdf(expandedData.expandedData.data)"> {{ $t('GLOBALS.BUTTONS.SAVE_TO_PDF') }} <Icon name="Download"></Icon></Button>
    </template>

    <template v-slot:paginatorstart></template>

  </DataTableWrapper>
</template>

<script setup lang="ts">
import DataTableWrapper from '@/components/layout/DataTableWrapper.vue';
import {DataTableService} from "@/services/dataTableService";
import {PagingModel} from "@/services/model/paging.model";
import {LocalSpinner} from "@/services/model/localSpinner";
import {DataHolder} from "@/models/data-holder";
import {useInvoicesStore} from "@/store/invoices.store";
import {useContextStore} from "@/store/context.store";
import {DefaultSortingModel} from "@/services/model/defaultSorting.model";
import {Invoice} from "@/models/invoice/invoice";
import factoryApi from "@/api/factory.api";
import {useToast} from "vue-toastification";
import {ResponseError} from "@/models/request-response-api";

const props = defineProps({
  expandedData: {
    type: Object
  },
  useSelectedCustomer: {
    type: Boolean,
    default: false
  }
});

const toast = useToast();
const invoiceStore = useInvoicesStore();

const invoicesDataTableService: DataTableService<Invoice> = {
  getDefaultSorting():DefaultSortingModel{
    return new DefaultSortingModel( "desc",  "payload.invoiceDetails.issueDt");
  },
  fetchListData(pagination: PagingModel, lockScreen: boolean, localSpinner: LocalSpinner | null): void {
    if (props.useSelectedCustomer) {
      const selectedCustomerId = useContextStore().selectedCustomer.id;
      invoiceStore.fetchInvoicesForCustomer(selectedCustomerId, lockScreen, localSpinner, pagination);
    } else {
      invoiceStore.fetchInvoices(lockScreen, localSpinner, pagination);
    }
  },
  getListDataHolder(): DataHolder<Invoice> {
    return invoiceStore.getListDataHolder;
  }
};

function saveToPdf(invoice: Invoice) {
  console.log('started saving');
  factoryApi.pdfConverterApi().downloadDocumentPdf(invoice, true, null, pdfSuccess, pdfFailed, true);
}

function pdfSuccess(data:any):void{
  console.log("Invoice pdf Success");
}

function pdfFailed(error : ResponseError):void{
  console.log("Invoice pdf Failed: " + error.message );
  toast.error(error.message as string);
}

function getStateColor(status: string) {
  if (status === 'accepted') return '#ccffcc';
  return '#ff9999';
}

</script>

<style scoped lang="scss">
h5 {
  padding-left: 12px;
}
.container {
  display: flex;
  flex-wrap: wrap;
}
.container>div{
  flex-grow: 1;
  width: 33%;
}
.icon-holder{
  display: flex;
  align-items: center;
}
</style>
