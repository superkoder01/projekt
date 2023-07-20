<template>

  <div class="card m-3">
    <div class="card-body">
      <Form @submit="nextPage" :validation-schema="schema" v-slot="{errors}">
        <div class="p-fluid formgrid grid ">
          <div class="field col-12 md:col-6 mb-4">

            <Field v-model="selectedCustomerAccessPointMeterNumber" required name="selectedCustomerAccessPointMeterNumber" as="select" class="form-control"
                   :class="{'is-invalid': errors.selectedCustomerAccessPointMeterNumber}">
              <option v-for="accessPoint in installationsStore.getListDataHolder.elements" :key="accessPoint.id" :value="accessPoint.meterNumber">
                <span>{{ getAccessPointLabel(accessPoint)}}
                </span>
              </option>
            </Field>
            <span>{{$t("FORMS.PLACEHOLDERS.METER_NUMBER")}}</span>
            <div class="invalid-feedback">{{ errors.selectedCustomerAccessPointMeterNumber ? $t(errors.selectedCustomerAccessPointMeterNumber) : '' }}</div>
          </div>

          <div class="field col-12 md:col-6 mb-4" >
            <Field v-model="objectName" name="objectName" placeholder=" " type="text" class="form-control" :class="{'is-invalid': errors.objectName}"/>
            <span>{{$t("FORMS.PLACEHOLDERS.OBJECT_NAME")}}</span>
            <div class="invalid-feedback">{{ errors.objectName ? $t(errors.objectName) : '' }}</div>
          </div>

<!--          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="address.postCode" name="postCode" placeholder=" " type="text" class="form-control" :class="{'is-invalid': errors.postCode}"></Field>
            <span>Kod pocztowy</span>
            <div class="invalid-feedback">{{errors.postCode ? $t(errors.postCode) : ''}}</div>
          </div>
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="address.city" name="city" placeholder=" " type="text" class="form-control" :class="{'is-invalid': errors.city}"></Field>
            <span>{{$t('FORMS.PLACEHOLDERS.CITY')}}</span>
            <div class="invalid-feedback">{{errors.city ? $t(errors.city) : ''}}</div>
          </div>
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="address.street" name="street" placeholder=" " type="text" class="form-control" :class="{'is-invalid': errors.street}"></Field>
            <span>Ulica</span>
            <div class="invalid-feedback">{{errors.street ? $t(errors.street) : ''}}</div>
          </div>

          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="meterNumber" name="meterNumber" placeholder=" " type="text" class="form-control" :class="{'is-invalid': errors.meterNumber}"></Field>
            <span>Numer licznika</span>
            <div class="invalid-feedback">{{errors.meterNumber ? $t(errors.meterNumber) : ''}}</div>
          </div>
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="sapCode" name="sapCode" placeholder=" " type="text" class="form-control" :class="{'is-invalid': errors.sapCode}"></Field>
            <span>Numer PPE (kod PPE)</span>
            <div class="invalid-feedback">{{errors.sapCode ? $t(errors.sapCode) : ''}}</div>
          </div>-->
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="registrationNumber" name="registrationNumber" placeholder=" " type="text" class="form-control" :class="{'is-invalid': errors.registrationNumber}"></Field>
            <span>{{$t("FORMS.PLACEHOLDERS.REGISTRATION_NUMBER")}}</span>
            <div class="invalid-feedback">{{ errors.registrationNumber ? $t(errors.registrationNumber) : '' }}</div>
          </div>
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="connectionPower" name="connectionPower" placeholder=" " type="number" class="form-control"
                   :class="{'is-invalid': errors.connectionPower}"/>
            <span>{{$t("FORMS.PLACEHOLDERS.CONNECTION_POWER")}}</span>
            <div class="invalid-feedback">{{ errors.connectionPower ? $t(errors.connectionPower) : '' }}</div>
          </div>

          <div class="field col-12 md:col-6 mb-4" >
            <Field v-model="contractedPower" name="contractedPower" placeholder=" " type="number" class="form-control"
                   :class="{'is-invalid': errors.contractedPower}"/>
            <span>{{$t("FORMS.PLACEHOLDERS.CONTRACTED_POWER")}}</span>
            <div class="invalid-feedback">{{ errors.contractedPower ? $t(errors.contractedPower) : '' }}</div>
          </div>
          <div class="field col-12 md:col-6 mb-4" >
            <Field v-model="declaredAnnualElectricityConsumption" name="declaredEnergyUsage" placeholder=" " type="number" class="form-control"
                   :class="{'is-invalid': errors.declaredEnergyUsage}"/>
            <span>{{$t("FORMS.PLACEHOLDERS.DECLARED_ENERGY_USAGE")}}</span>
            <div class="invalid-feedback">{{ errors.declaredEnergyUsage ? $t(errors.declaredEnergyUsage) : '' }}</div>
          </div>

<!--          <div class="field col-12 md:col-6 mb-4">-->
<!--            <Field v-model="selectedPhase" required name="selectedPhase" as="select" class="form-control"-->
<!--                   :class="{'is-invalid': errors.selectedPhase}">-->
<!--              <option v-for="[phase, value] in phaseList.entries()" :key="phase" :value="value">{{ phase }}</option>-->
<!--            </Field>-->
<!--            <span>Faza</span>-->
<!--            <div class="invalid-feedback">{{ errors.selectedPhase }}</div>-->
<!--          </div>-->
        </div>
<!--        <div class="field col-12 md:col-6 mb-4">-->
<!--          <Field v-model="sourceType" name="sourceType" placeholder=" " type="text" class="form-control"-->
<!--                 :class="{'is-invalid': errors.sourceType}"/>-->
<!--          <span>Rodzaj źródła</span>-->
<!--          <div class="invalid-feedback">{{ errors.sourceType }}</div>-->
<!--        </div>-->

        <div class="field col-12 md:col-6 mb-4" >
          <Field v-model="sourcePower" name="sourcePower" placeholder=" " type="number" class="form-control"
                 :class="{'is-invalid': errors.sourcePower}"/>
          <span>{{$t("FORMS.PLACEHOLDERS.SOURCE_POWER")}}</span>
          <div class="invalid-feedback">{{ errors.sourcePower ? $t(errors.sourcePower) : '' }}</div>
        </div>



        <div class="grid grid-nogutter justify-content-between">
          <Button :label="$t('GLOBALS.BUTTONS.BACK')" @click="prevPage"  icon="pi pi-angle-left"></Button>
          <Button :label="$t('GLOBALS.BUTTONS.NEXT')" type="submit" icon="pi pi-angle-right" icon-pos="right"></Button>
        </div>
      </Form>
    </div>
  </div>

</template>

<script setup lang="ts">
import {defineEmits, onMounted, PropType, ref, watch} from "vue";
import {Field, Form} from 'vee-validate';
import * as Yup from 'yup';
import {onBeforeMount} from "@vue/runtime-core";
import { Address, UnitOfEnergy, UnitOfPower } from "@/models/billing/billing";
import {
  ContractForm,
  UnitAmount
} from "@/components/forms/create-contract/Contract";

import { useCustomerStore } from "@/store/customers.store";
import { Installation } from "@/components/forms/create-installation/Installation";
import { useInstallationsStore } from "@/store/installations.store";
import { useI18n } from "vue-i18n";
import factoryApi from "@/api/factory.api";

const i18n = useI18n();
const emit = defineEmits(['nextPage', 'prevPage', 'update:newContract']);
const installationsStore = useInstallationsStore();
const props = defineProps({

  newContract: {
    type: Object as PropType<ContractForm>,
    required: true
  },
  useSelectedContract: {
    type: Boolean
  }
});

onBeforeMount(() => {
  installationsStore.fetchSelectedCustomerAccessPoints(useCustomerStore().selectedCustomerData.id, true, null);

  if(props.useSelectedContract){
    connectionPower.value = props.newContract.serviceAccessPoints[0].connectionPower.amount;
    contractedPower.value = props.newContract.serviceAccessPoints[0].contractedPower.amount;
    sourcePower.value = props.newContract.serviceAccessPoints[0].sourcePower.amount;
    objectName.value = props.newContract.serviceAccessPoints[0].objectName;
    //registrationNumber.value = props.newContract.serviceAccessPoints[0].registrationNumber;
  }
});

const phaseList = new Map<string, number>([
  ["jednowazowa", 1],
  ["trójfazowa", 3]
]);

const selectedCustomerAccessPointMeterNumber = ref();
const selectedPhase = ref();
const declaredAnnualElectricityConsumption = ref(props.newContract.offer.payload.conditions.estimatedAnnualElectricityConsumption.amount);
const connectionPower = ref();
const contractedPower = ref();
const sourceType = ref();
const sourcePower = ref();
const address = ref<Address>({} as Address);
const selectedInstallation= ref({} as Installation);
const sapCode = ref();
const meterNumber = ref();
const registrationNumber = ref(); // numer ewidencyjny
const objectName = ref();
const schema = Yup.object().shape({
 // postCode: Yup.string().required(),
  objectName: Yup.string()
    .required('FORMS.VALIDATION_ERRORS.OBJECT_NAME'),
 // city: Yup.string().required(),
 // street: Yup.string().required(),
 // meterNumber: Yup.string().required(),
 // sapCode: Yup.string().required(),
  registrationNumber: Yup.string()
    .required('FORMS.VALIDATION_ERRORS.REGISTRATION_NUMBER'),
  connectionPower: Yup.number()
    .positive('FORMS.VALIDATION_ERRORS.POSITIVE_NUMBER')
    .required('FORMS.VALIDATION_ERRORS.CONNECTION_POWER'),
  contractedPower: Yup.number()
    .positive('FORMS.VALIDATION_ERRORS.POSITIVE_NUMBER')
    .required('FORMS.VALIDATION_ERRORS.CONTRACTED_POWER'),
  declaredEnergyUsage: Yup.number()
    .positive('FORMS.VALIDATION_ERRORS.POSITIVE_NUMBER')
    .required('FORMS.VALIDATION_ERRORS.DECLARED_ENERGY_USAGE'),
  sourcePower: Yup.number()
    .positive('FORMS.VALIDATION_ERRORS.POSITIVE_NUMBER')
    .required('FORMS.VALIDATION_ERRORS.SOURCE_POWER'),
  selectedCustomerAccessPointMeterNumber: Yup.string()
    .required('FORMS.VALIDATION_ERRORS.ACCESS_POINT_METER_NUMBER')
});

const prevPage = () => {
  emit('prevPage', {pageIndex:2});
};
watch(() => installationsStore.installations, () => {
  if(props.useSelectedContract)
    selectedCustomerAccessPointMeterNumber.value = props.newContract.serviceAccessPoints[0].meterNumber;
});
watch(() => selectedCustomerAccessPointMeterNumber.value, () => {
  selectedInstallation.value = installationsStore.installations.elements.find((installation: Installation) => {
    return installation.meterNumber == selectedCustomerAccessPointMeterNumber.value;
  });
},  {immediate:true});
const nextPage = () => {


  let updatedNewContract: ContractForm = props.newContract;
  console.log(selectedInstallation.value)
  updatedNewContract.serviceAccessPoints[0].objectName = objectName.value;
  updatedNewContract.serviceAccessPoints[0].address = selectedInstallation.value.city+" "+selectedInstallation.value.address;
  updatedNewContract.serviceAccessPoints[0].sapCode = selectedInstallation.value.sapCode;
  updatedNewContract.serviceAccessPoints[0].meterNumber = selectedInstallation.value.meterNumber;
  updatedNewContract.serviceAccessPoints[0].estimatedEnergyUsage = props.newContract.offer.payload.conditions.estimatedAnnualElectricityConsumption;
  updatedNewContract.serviceAccessPoints[0].declaredEnergyUsage = {amount: declaredAnnualElectricityConsumption.value, unit: UnitOfEnergy.MWh} as UnitAmount;

  updatedNewContract.serviceAccessPoints[0].connectionPower = {amount: connectionPower.value, unit: UnitOfPower.kW} as UnitAmount;
  updatedNewContract.serviceAccessPoints[0].contractedPower = {amount: contractedPower.value, unit: UnitOfPower.kW} as UnitAmount;


  // updatedNewContract.serviceAccessPoints[0].phase = selectedPhase.value.toString();

  updatedNewContract.serviceAccessPoints[0].sourcePower = {amount: sourcePower.value, unit: UnitOfPower.kWp} as UnitAmount;
  // updatedNewContract.serviceAccessPoints[0].sourceType = sourceType.value;

  emit('nextPage', {pageIndex: 2});
  emit('update:newContract', updatedNewContract);
};

function getAccessPointLabel(accessPoint : Installation){
  return i18n.t("FORMS.PLACEHOLDERS.METER_NUMBER") + ": " + accessPoint.meterNumber + "       "
    + i18n.t("FORMS.PLACEHOLDERS.SAP_CODE") + ": " + accessPoint.sapCode + "       "
    + i18n.t("FORMS.PLACEHOLDERS.ADDRESS") + ": " + accessPoint.city +", " + accessPoint.address;
}

</script>
<style scoped lang="scss">

</style>
