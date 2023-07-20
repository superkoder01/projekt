<template>

  <div class="card m-3">
    <div class="card-body">

      <Form @submit="nextPage" :validation-schema="schema" v-slot="{errors}" >
        <div class="p-fluid formgrid grid ">

          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="selectedOfferDraftId" required name="selectedOfferDraftId" as="select" class="form-control"
                   :class="{'is-invalid': errors.selectedOfferDraftId}">
              <option v-for="offerDraft in offerDraftStore.getListDataHolder.elements" :key="offerDraft" :value="offerDraft.id">{{ offerDraft.payload.offerDetails.title }}</option>
            </Field>
            <span>{{$t("FORMS.PLACEHOLDERS.DRAFT_OFFER")}}</span>
            <div class="invalid-feedback">{{ errors.selectedOfferDraftId ? $t(errors.selectedOfferDraftId) : '' }}</div>
          </div>
          <div class="col-12 md:col-6 mb-4" v-if="selectedOfferDraft">
            <h3>Szczegóły</h3>

            <span class="detailed-data" style="width: 100%">
              <span>
                <span><Icon name="None" /></span>
                <span>
                  <div class="label-small">{{ $t("PRICING_VIEW.TABLE_HEADERS.NAME") }}:</div>
                  {{ selectedOfferDraft.payload.offerDetails.title }}
                </span>
              </span>

              <span>
                <span><Icon name="None" /></span>
                <span>
                  <div class="label-small">{{ $t("PRICING_VIEW.TABLE_HEADERS.TYPE") }}:</div>
                  {{ selectedOfferDraft.payload.offerDetails.type }}
                </span>
              </span>

              <span>
                <span><Icon name="None" /></span>
                <span>
                  <div class="label-small">{{ $t("FORMS.PLACEHOLDERS.CONTRACT_TYPE") }}:</div>
                  {{ selectedOfferDraft.payload.offerDetails.agreementType }}
                </span>
              </span>

              <span>
                <span><Icon name="None" /></span>
                <span>
                  <div class="label-small">{{ $t("PRICING_VIEW.TABLE_HEADERS.TARRIF_GROUP") }}:</div>
                  {{ selectedOfferDraft.payload.offerDetails.tariffGroup }}
                </span>
              </span>

            </span>

            <h3>Cennik sprzedaży</h3>
            <span class="detailed-data">
              <span>
                <span><Icon name="None" /></span>
                <span>
                  <div class="label-small">{{ $t("OFFERS_VIEW.TABLE_HEADERS.PRICING_NAME") }}:</div>
                  {{ selectedOfferDraft.payload.priceList.name }}
                </span>
              </span>
            </span>
            <!-- <div>Opłata handlowa: {{selectedOfferDraft.payload.priceList.commercialFee}} </div> -->

            <span v-if="selectedOfferDraft.payload.priceList.commercialFee !== undefined">
              <h5>{{ $t("PRICING_VIEW.TABLE_HEADERS.COMMERCIAL_FEE") }}</h5>
              <div class="detail-headers six-cols">
                <div class="detail-header">{{ $t("PRICING_VIEW.TABLE_HEADERS.SINGLE_COMMERCIAL_FEE.FROM") }}</div>
                <div class="detail-header">{{ $t("PRICING_VIEW.TABLE_HEADERS.SINGLE_COMMERCIAL_FEE.TO") }}</div>
                <div class="detail-header">{{ $t("PRICING_VIEW.TABLE_HEADERS.SINGLE_COMMERCIAL_FEE.UNIT") }}</div>
                <div class="detail-header">{{$t("PRICING_VIEW.TABLE_HEADERS.SINGLE_COMMERCIAL_FEE.PRICE.CALENDAR_UNIT")}}</div>
                <div class="detail-header">{{ $t("PRICING_VIEW.TABLE_HEADERS.SINGLE_COMMERCIAL_FEE.PRICE.COST") }}</div>
                <div class="detail-header">{{$t("PRICING_VIEW.TABLE_HEADERS.SINGLE_COMMERCIAL_FEE.PRICE.CURRENCY") }}</div>
              </div>
              <div class="detail-row six-cols" v-bind:key="item.id" v-for="item in selectedOfferDraft.payload.priceList.commercialFee">
                <div class="detail-item">{{ item.from }}</div>
                <div class="detail-item">{{ item.to === undefined ? "-" : item.to }}</div>
                <div class="detail-item">{{ item.unit }}</div>
                <div class="detail-item">{{ item.price.calendarUnit === "" ? "-" : $t("ENUMS.CALENDAR_UNIT_TYPE." + toCalendarUnitTypeEnumKey(item.price.calendarUnit)) }}</div>
                <div class="detail-item">{{ item.price.cost === undefined ? "-" : item.price.cost }}</div>
                <div class="detail-item">{{ item.price.currency }}</div>
              </div>
              <h5>{{ $t("PRICING_VIEW.TABLE_HEADERS.ZONES") }}</h5>
                <div class="detail-headers five-cols">
                  <div class="detail-header"> {{ $t("PRICING_VIEW.TABLE_HEADERS.ZONE.ID") }}</div>
                  <div class="detail-header"> {{ $t("PRICING_VIEW.TABLE_HEADERS.ZONE.NAME") }}</div>
                  <div class="detail-header">{{ $t("PRICING_VIEW.TABLE_HEADERS.ZONE.UNIT") }}</div>
                  <div class="detail-header"> {{ $t("PRICING_VIEW.TABLE_HEADERS.ZONE.COST") }}</div>
                  <div class="detail-header"> {{ $t("PRICING_VIEW.TABLE_HEADERS.ZONE.CURRENCY") }}</div>
                </div>
                <div class="detail-row five-cols" v-bind:key="zone.id"
                    v-for="zone in selectedOfferDraft.payload.priceList.zones">
                  <div class="detail-item">{{ zone.id }}</div>
                  <div class="detail-item">{{ zone.name }}</div>
                  <div class="detail-item">{{ zone.unit }}</div>
                  <div class="detail-item">{{ zone.cost }}</div>
                  <div class="detail-item">{{ zone.currency }}</div>
                </div>
            </span>

            <span v-if="selectedOfferDraft.payload.repurchase.type == 'rdn'">
              <h4>Cennik odkupu</h4>
              <span class="detailed-data">
                <span>
                  <span><Icon name="None" /></span>
                  <span>
                    <div class="label-small">{{ $t("OFFERS_VIEW.TABLE_HEADERS.REPURCHASE_NAME") }}:</div>
                    {{ selectedOfferDraft.payload.repurchase.name }}
                  </span>
                </span>

                  <span>
                    <span><Icon name="None" /></span>
                    <span>
                      <div class="label-small">{{ $t("OFFERS_VIEW.TABLE_HEADERS.COMMERCIAL_FEE") }}:</div>
                      {{typeof selectedOfferDraft.payload.repurchase.price === 'number'? selectedOfferDraft.payload.repurchase.price : '-'}}
                    </span>
                  </span>
              </span>
            </span>
          </div>
        </div>
        <div class="grid grid-nogutter justify-content-between">
          <Button :label="$t('GLOBALS.BUTTONS.NEXT')" type="submit" icon="pi pi-angle-right" icon-pos="right"></Button>
        </div>
      </Form>
    </div>
  </div>

</template>

<script setup lang="ts">
import { defineEmits, PropType, ref, watch } from "vue";
import {Field, Form} from 'vee-validate';
import * as Yup from 'yup';
import {onBeforeMount} from "@vue/runtime-core";
import {FormPricing} from "@/components/forms/create-pricing/Pricing";
import {useDraftOffersStore} from "@/store/draftOffers.store";
import {OfferForm} from "@/components/forms/create-offer/OfferForm";
import { PagingModel } from "@/services/model/paging.model";
import {toCalendarUnitTypeEnumKey, DateFormat} from "@/models/billing/billing";
import factoryApi from "@/api/factory.api";


const emit = defineEmits(['nextPage', 'prevPage', 'update:newOffer']);

const props = defineProps({

  newOffer: {
    type: Object as PropType<OfferForm>,
    required: true
  },
  useSelectedOffer: {
    type: Boolean
  }
});
const offerDraftStore = useDraftOffersStore();



onBeforeMount(() => {
  const paging = new PagingModel([]);
  paging.limit = 10000;
  offerDraftStore.fetchDraftOffers(true, null, paging);
});

const selectedOfferDraftId = ref();
const selectedOfferDraft = ref();

const schema = Yup.object().shape({
  selectedOfferDraftId: Yup.string()
    .required('FORMS.VALIDATION_ERRORS.REQUIRED')
});
watch(() => offerDraftStore.draftOffers, ()=> {
  if(props.useSelectedOffer)
    selectedOfferDraftId.value = props.newOffer.offerDraft.id;
});
watch(() => selectedOfferDraftId.value, () => {

  selectedOfferDraft.value = offerDraftStore.getOfferDraft.find(offerDraft => {
    return offerDraft.id === selectedOfferDraftId.value;
  });
  console.log(selectedOfferDraft.value)
});

const nextPage = () => {

  let updatedNewOffer: OfferForm = props.newOffer;
  updatedNewOffer.offerDraft = selectedOfferDraft.value;
  emit('nextPage', {pageIndex: 0});
  emit('update:newOffer', updatedNewOffer);
};

</script>
<style scoped lang="scss">
@import '../../../styles/variables.scss';

h4, h5 {
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

.detailed-data {
    display: grid;
    grid-template-columns: 1fr 1fr;
    grid-column-gap: 10px;
    grid-row-gap: 10px;
    margin-bottom: 10px;

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

  @media (max-width: 1150px) {
    .detailed-data {
      grid-template-columns: 1fr;
    }
  }

</style>
