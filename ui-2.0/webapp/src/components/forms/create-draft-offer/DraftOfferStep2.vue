<template>
  <div class="card m-3">
    <div class="card-body">
      <Form @submit="nextPage" :validation-schema="schema" v-slot="{errors}">
        <!--        <div class="grid grid-nogutter justify-content-between">-->
        <div class="responsive-grid">
          <div class="grid-column">
            <div class="radio-container">
              <div class="field-radiobutton">
                <RadioButton id="saleEmpty" name="saleEmpty" :value="ServiceTypeEnum.NONE" v-model="selectedSaleType"/>
                <label for="saleEmpty">{{ $t('FORMS.PLACEHOLDERS.NONE') }} </label>
              </div>
              <div class="field-radiobutton" disabled>
                <RadioButton id="saleRdn" name="saleRdn" :value="ServiceTypeEnum.SALE" v-model="selectedSaleType"
                             disabled="true"/>
                <label for="saleRdn" disabled="true">{{ $t('FORMS.PLACEHOLDERS.RDN_SALE') }}</label>
              </div>
              <div class="field-radiobutton">
                <RadioButton id="sale" name="sale" :value="ServiceTypeEnum.SALE_RDN" v-model="selectedSaleType"/>
                <label for="sale">{{ $t('FORMS.PLACEHOLDERS.FIXED_RATE_SALE') }}</label>
              </div>
            </div>
            <div class="p-fluid formgrid grid ">
              <div v-if="selectedSaleType === ServiceTypeEnum.SALE_RDN" class="field col-12 md:col-6 mb-4">
                <Field v-model="selectedPricingSaleId" name="selectedPricingSale" as="select" class="form-control"
                       :class="{'is-invalid': errors.selectedPricingSale}">
                  <option v-for="pricingSale in pricingStore.getPricingSALE" :key="pricingSale.id" :value="pricingSale.id">
                    {{ pricingSale.name }}
                  </option>
                </Field>
                <span>{{ $t('FORMS.PLACEHOLDERS.SALE_PRICING') }}</span>
                <div class="invalid-feedback">{{ errors.selectedPricingSale }}</div>
              </div>
            </div>
            <div style="width: 100%" v-if="(selectedSaleType === ServiceTypeEnum.SALE || selectedSaleType === ServiceTypeEnum.SALE_RDN) && selectedPricingSale != undefined">
              <span style="width: 100%">
                <pricing-details :pricing="selectedPricingSale"/>
              </span>
            </div>
          </div>
          <div class="grid-column">
            <div class="radio-container">
              <div class="field-radiobutton">
                <RadioButton id="purchaseEmpty" name="purchaseEmpty" :value="ServiceTypeEnum.NONE" v-model="selectedPurchaseType"/>
                <label for="purchaseEmpty">{{ $t('FORMS.PLACEHOLDERS.NONE') }}</label>
              </div>
              <div class="field-radiobutton">
                <RadioButton id="purchaseRdn" name="purchaseRdn" :value="ServiceTypeEnum.REPURCHASE_RDN" v-model="selectedPurchaseType"/>
                <label for="purchaseRdn">{{ $t('FORMS.PLACEHOLDERS.RDN_REPURCHASE') }}</label>
              </div>
              <div class="field-radiobutton">
                <RadioButton id="purchaseRdn" name="purchaseRdn" :value="ServiceTypeEnum.REPURCHASE" v-model="selectedPurchaseType"/>
                <label for="purchaseRdn">{{ $t('FORMS.PLACEHOLDERS.FIXED_RATE_REPURCHASE') }}</label>
              </div>
            </div>
            <div class="p-fluid formgrid grid ">
              <div v-if="selectedPurchaseType === ServiceTypeEnum.REPURCHASE"
                   class="field col-12 md:col-6 mb-4">
                <Field v-model="selectedPricingRepurchaseId" name="selectedPricingRepurchase" as="select"
                       class="form-control" :class="{'is-invalid': errors.selectedPricingSale}">
                  <option v-for="pricing in pricingStore.getPricingREPURCHASE" :key="pricing.id" :value="pricing.id">
                    {{ pricing.name }}
                  </option>
                </Field>
                <span>{{ $t('FORMS.PLACEHOLDERS.REPURCHASE_PRICING') }}</span>
                <div class="invalid-feedback">{{ errors.selectedPricingSale }}</div>
              </div>
              <div v-if="selectedPurchaseType === ServiceTypeEnum.REPURCHASE_RDN"
                   class="field col-12 md:col-6 mb-4">
                <Field v-model="selectedPricingRepurchaseRDNId" name="selectedPricingRepurchaseRDN" as="select"
                       class="form-control" :class="{'is-invalid': errors.selectedPricingSale}">
                  <option v-for="pricingRepurchase in pricingStore.getPricingREPURCHASE_RDN" :key="pricingRepurchase.id"
                          :value="pricingRepurchase.id">{{ pricingRepurchase.name }}
                  </option>
                </Field>
                <span>{{ $t('FORMS.PLACEHOLDERS.RDN_REPURCHASE_PRICING')  }} </span>
                <div class="invalid-feedback">{{ errors.selectedPricingSale }}</div>
              </div>
            </div>
            <div style="width: 100%" v-if="(selectedPurchaseType === ServiceTypeEnum.REPURCHASE || selectedPurchaseType === ServiceTypeEnum.REPURCHASE_RDN) && selectedPricingRepurchase != undefined">
              <span style="width: 100%">
                <pricing-details :pricing="selectedPricingRepurchase"/>
              </span>
            </div>
          </div>

        </div>
        <div class="grid grid-nogutter justify-content-between">
          <Button :label="$t('GLOBALS.BUTTONS.BACK')" @click="goBack" type="submit" icon="pi pi-angle-left"
                  icon-pos="left"></Button>
          <Button :label="$t('GLOBALS.BUTTONS.NEXT')" type="submit" icon="pi pi-angle-right" icon-pos="right"></Button>
<!--          <Button label="Next" icon="pi pi-angle-right" icon-pos="right" @click="test"></Button>-->
        </div>
      </Form>
    </div>
  </div>
</template>

<script setup lang="ts">
import {computed, defineEmits, PropType, ref} from 'vue';
import {Field, Form} from 'vee-validate';
import * as Yup from 'yup';
import {onBeforeMount} from "@vue/runtime-core";
import {FormOfferDraft} from "@/components/forms/create-draft-offer/OfferDraft";
import {ServiceTypeEnum} from "@/models/billing/enum/service-type.enum";
import {usePricingStore} from "@/store/pricing.store";
import {PagingModel} from "@/services/model/paging.model";
import {Pricing} from "@/models/billing/billing";
import PricingDetails from "@/components/PricingDetails.vue";

const emit = defineEmits(['nextPage', 'prevPage', 'update:updatedOfferDraft']);

const props = defineProps({
  newOfferDraft: {
    type: Object as PropType<FormOfferDraft>,
    required: true
  }
});

onBeforeMount(() => {
  console.log();
});

onBeforeMount(() => {
  const paging = new PagingModel([]);
  paging.limit = 10000;
  pricingStore.fetchPricing(true, null, paging);
});


const selectedSaleType = ref<ServiceTypeEnum>(ServiceTypeEnum.NONE);
const selectedPurchaseType = ref<ServiceTypeEnum>(ServiceTypeEnum.NONE);

// const selectedServiceList = ref<Array<ServiceTypeEnum>>(new Array<ServiceTypeEnum>());

const selectedPricingSaleId = ref();
const selectedPricingRepurchaseRDNId = ref();
const selectedPricingRepurchaseId = ref();

const selectedPricingSale = computed<Pricing | undefined>(() => {
  if (selectedPricingSaleId.value != undefined) {
    return pricingStore.getPricingSALE?.find(pricing => {
      return pricing.id === selectedPricingSaleId.value;
    });
  }
  return undefined;
});

const selectedPricingRepurchase = computed<Pricing | undefined>(() => {
  if(selectedPurchaseType.value === ServiceTypeEnum.REPURCHASE) {
    if (selectedPricingRepurchaseId.value != undefined) {
      return pricingStore.getPricingREPURCHASE?.find(pricing => {
        return pricing.id === selectedPricingRepurchaseId.value;
      });
    }
  } else if(selectedPurchaseType.value === ServiceTypeEnum.REPURCHASE_RDN){
    if (selectedPricingRepurchaseRDNId.value != undefined) {
      return pricingStore.getPricingREPURCHASE_RDN?.find(pricing => {
        return pricing.id === selectedPricingRepurchaseRDNId.value;
      });
    }
  }
  return undefined;
});

const pricingStore = usePricingStore();

const schema = Yup.object().shape({
  //TODO:
  // title: Yup.string().required('Title is required'),
  // type: Yup.string().required('Type is required'),
  // agreementType: Yup.string().required('AgreementType is required'),
});

const nextPage = () => {
  let updatedOfferDraft: FormOfferDraft = props.newOfferDraft;
  updatedOfferDraft.price = pricingStore.getPricingSALE?.find(pricing => {
    return pricing.id === selectedPricingSaleId.value;
  });
  updatedOfferDraft.repurchase = pricingStore.getPricingREPURCHASE_RDN?.find(pricing => {
    return pricing.id === selectedPricingRepurchaseRDNId.value;
  });

  emit('nextPage', {pageIndex: 1});
  emit('update:updatedOfferDraft', updatedOfferDraft);
};

const goBack = () => {
  emit('prevPage', {pageIndex: 1});
};

const test = () => {
  console.log("selectedSaleType:" + selectedSaleType.value);
  console.log("selectedPricingSaleId:" + JSON.stringify(selectedPricingSaleId.value));

  console.log("selectedPurchaseType:" + selectedPurchaseType.value);
  console.log("selectedPricingRepurchaseId:" + JSON.stringify(selectedPricingRepurchaseId.value));
  console.log("selectedPricingRepurchaseRDNId:" + JSON.stringify(selectedPricingRepurchaseRDNId.value));
  console.log("selectedPricingRepurchase:" + JSON.stringify(selectedPricingRepurchase.value));

};


// function addSelectedService(type: ServiceTypeEnum) {
//   console.log(type);
//   if (selectedServiceList.value.indexOf(type) < 0) {
//     selectedServiceList.value.push(type);
//   }
// }
//
// function removeService(serviceType: ServiceTypeEnum) {
//   const index = selectedServiceList.value.indexOf(serviceType, 0);
//   console.log(index);
//   console.log(serviceType);
//   console.log(ServiceTypeEnum.REPURCHASE);
//   if (index > -1) {
//     selectedServiceList.value.splice(index, 1);
//   }
// }

</script>


<style scoped lang="scss">
.responsive-grid {
  display: grid;
  width: 100%;
  grid-template-columns: 1fr 1fr;
  column-gap: 10px;

  .grid-column {
    display: flex;
    flex-direction: column;
    align-items: center;

    .radio-container {
      display: flex;
      flex-direction: column;
      align-items: flex-start;
    }

    .formgrid {
      width: 100% !important;
      margin-top: 2rem;
    }
  }
}

@media (max-width: 850px) {
  .responsive-grid {
    grid-template-columns: 1fr;

    .grid-column {

      .radio-container {
        align-items: center;
      }
    }
  }
}
</style>
