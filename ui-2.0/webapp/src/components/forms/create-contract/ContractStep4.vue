<template>

  <div class="card m-3">
    <div class="card-body">
      <Form @submit="complete" :validation-schema="schema" v-slot="{errors}">
        <div class="p-fluid formgrid grid ">

          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="noticePeriod" required name="noticePeriod" as="select" class="form-control"
                   :class="{'is-invalid': errors.noticePeriod}">
              <option v-for="period in noticePeriodList" :key="period" :value="period">{{ $t("FORMS.PLACEHOLDERS."+period+"_MONTHS") }}</option>
            </Field>
            <span>{{$t("FORMS.PLACEHOLDERS.NOTICE_PERIOD")}}</span>
            <div class="invalid-feedback">{{ errors.noticePeriod }}</div>
          </div>
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="startDate" name="startDate" type="date" class="form-control"
                   :class="{'is-invalid': errors.startDate}"></Field>
            <span>{{$t("FORMS.PLACEHOLDERS.START_DATE")}}</span>
            <div class="invalid-feedback">{{ errors.startDate }}</div>
          </div>

          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="bankAccountNumber" name="bankAccountNumber" placeholder=" " type="number" class="form-control"
                   :class="{'is-invalid': errors.bankAccountNumber}"/>
            <span>{{$t("FORMS.PLACEHOLDERS.BANK_ACCOUNT_NUMBER")}}</span>
            <div class="invalid-feedback">{{ errors.name }}</div>
          </div>
        </div>

        <div class="grid grid-nogutter justify-content-between">
          <Button :label="$t('GLOBALS.BUTTONS.BACK')" @click="prevPage"  icon="pi pi-angle-left"></Button>
          <Button :label="$t('GLOBALS.BUTTONS.CREATE')" type="submit" icon="pi pi-angle-right"></Button>
        </div>
      </Form>
    </div>
  </div>

</template>

<script setup lang="ts">
import {computed, defineEmits, onMounted, PropType, ref, watch} from "vue";
import {Field, Form} from 'vee-validate';
import * as Yup from 'yup';
import { ContractForm } from "@/components/forms/create-contract/Contract";
import moment from "moment";
import {onBeforeMount} from "@vue/runtime-core";

const emit = defineEmits(['complete', 'prevPage', 'update:newContract']);

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
  if(props.useSelectedContract) {
    noticePeriod.value = props.newContract.serviceAccessPoints[0].currentSeller.noticePeriod;
    startDate.value = moment(props.newContract.startDate).format('YYYY-MM-DD');
    bankAccountNumber.value = props.newContract.bankAccountNumber;
  }
});

const noticePeriodList = ref(['0','1', '2', '3', '4', '5', '6']);
const noticePeriod = ref('0');
const bankAccountNumber = ref();
const startDate = ref<string>((moment(new Date(new Date().getFullYear(), new Date().getMonth()+Number(noticePeriod.value)+1, 1))).format('YYYY-MM-DD'));

watch(() => noticePeriod.value, ()=> {
  startDate.value = (moment(new Date(new Date().getFullYear(), new Date().getMonth()+Number(noticePeriod.value)+1, 1))).format('YYYY-MM-DD');

});

const schema = Yup.object().shape({
  noticePeriod: Yup.string()
    .required(),
  startDate: Yup
    .date()
    .when('noticePeriod',(noticePeriod, schema) =>  {
      if(noticePeriod != '0')
        return schema.min(new Date(new Date().getFullYear(), new Date().getMonth()+Number(noticePeriod)+1, 1));
    }),
  bankAccountNumber: Yup.string()
    .required('FORMS.VALIDATION_ERRORS.BANK_ACCOUNT_NUMBER')

});



const prevPage = () => {
  emit('prevPage', {pageIndex: 3});
};

const complete = () => {

  let updatedNewContract: ContractForm = props.newContract;

  updatedNewContract.bankAccountNumber = bankAccountNumber.value;
  updatedNewContract.serviceAccessPoints[0].currentSeller.noticePeriod = noticePeriod.value;
  updatedNewContract.startDate = startDate.value;

  emit('complete');
  emit('update:newContract', updatedNewContract);
};



</script>
<style scoped lang="scss">

</style>
