<template>
  <div class="card m-3">
    <div class="card-body">
      <Form @submit="nextPage" :validation-schema="schema" v-slot="{errors}">
        <div class="p-fluid formgrid grid ">

          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="pricingName" name="pricingName" placeholder=" " type="text" class="form-control" :class="{'is-invalid': errors.pricingName}"/>
            <span>{{ $t('FORMS.PLACEHOLDERS.NAME') }}</span>
            <div class="invalid-feedback">{{ errors.pricingName ? $t(errors.pricingName) : '' }}</div>
          </div>
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="selectedTariff" required name="tariff" as="select" class="form-control" :class="{'is-invalid': errors.tariff}">
              <option v-for="tariff in tariffList" :key="tariff" :value="tariff.name">{{ tariff.name }}</option>
            </Field>
            <span>{{ $t('FORMS.PLACEHOLDERS.TARIFF_GROUP') }}</span>
            <div class="invalid-feedback">{{ errors.tariff ? $t(errors.tariff) : '' }}</div>
          </div>

          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="selectedServiceType" required name="serviceType" as="select" class="form-control" :class="{'is-invalid': errors.serviceType}">
              <option v-for="serviceType in serviceTypeList" :key="serviceType" :value="serviceType">
                {{ $t("ENUMS.SERVICE_TYPE." + toServiceTypeEnumKey(serviceType)) }}
              </option>
            </Field>
            <span>{{ $t('FORMS.PLACEHOLDERS.TYPE') }}</span>
            <div class="invalid-feedback">{{ errors.serviceType ? $t(errors.serviceType) : '' }}</div>

          </div>

          <!--          <div v-if="selectedServiceType===ServiceTypeEnum.SALE">-->
          <!--            <div class="field col-12 md:col-6 mb-4">-->
          <!--              <Field v-model="startDate" name="startDate" type="date" class="form-control"-->
          <!--                     :class="{'is-invalid': errors.startDate}"></Field>-->
          <!--              <span>{{ $t('FORMS.PLACEHOLDERS.START_DATE') }}</span>-->
          <!--              <div class="invalid-feedback">{{ errors.startDate }}</div>-->
          <!--            </div>-->
          <!--            <div class="field col-12 md:col-6 mb-4">-->
          <!--              <Field v-model="endDate" name="endDate" type="date" class="form-control"-->
          <!--                     :class="{'is-invalid': errors.endDate}"></Field>-->
          <!--              <span>{{ $t('FORMS.PLACEHOLDERS.END_DATE') }}</span>-->
          <!--              <div class="invalid-feedback">{{ errors.endDate }}</div>-->
          <!--            </div>-->
          <!--          </div>-->

        </div>
        <div class="grid grid-nogutter justify-content-between">
          <Button :label="$t('GLOBALS.BUTTONS.NEXT')" type="submit" icon="pi pi-angle-right" icon-pos="right"></Button>
        </div>
      </Form>
    </div>
  </div>
</template>

<script setup lang="ts">
import {defineEmits, PropType, ref} from 'vue';
import {Field, Form} from 'vee-validate';
import * as Yup from 'yup';
import {FormPricing} from "@/components/forms/create-pricing/Pricing";
import moment from 'moment';
import {ServiceTypeEnum, toServiceTypeEnumKey} from "@/models/billing/enum/service-type.enum";
import {DateFormat, TariffGroupType} from "@/models/billing/billing";

const emit = defineEmits(['nextPage', 'prevPage', 'update:newPricing']);

const props = defineProps({
  newPricing: {
    type: Object as PropType<FormPricing>,
    required: true
  }
});

const pricingName = ref();
const selectedTariff = ref();
const tariffList = ref([{id: 1, name: TariffGroupType.G11}]);
const serviceTypeList = ref<Array<ServiceTypeEnum>>([ServiceTypeEnum.SALE, ServiceTypeEnum.REPURCHASE_RDN, ServiceTypeEnum.REPURCHASE]);

const selectedServiceType = ref<ServiceTypeEnum>();


const schema = Yup.object().shape({
  pricingName: Yup.string().required('FORMS.VALIDATION_ERRORS.REQUIRED'),
  tariff: Yup.string().required('FORMS.VALIDATION_ERRORS.REQUIRED'),
  serviceType: Yup.string().required('FORMS.VALIDATION_ERRORS.REQUIRED'),

});

function serviceTypeChanged() {
  // Form.validate();
}

const nextPage = () => {

  let updatedNewPricing: FormPricing = props.newPricing;
  updatedNewPricing.tariffGroup = selectedTariff.value;
  updatedNewPricing.type = selectedServiceType.value;
  updatedNewPricing.name = pricingName.value;


  emit('nextPage', {pageIndex: 0});
  emit('update:newPricing', updatedNewPricing);
};

</script>
<style scoped lang="scss">

</style>
