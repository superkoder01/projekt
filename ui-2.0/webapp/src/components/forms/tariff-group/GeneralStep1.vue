<template>

  <div class="card no-border m-3">
    <h5 class="card-header">{{ $t('FORMS.HEADERS.GENERAL_INFO') }}</h5>
    <div class="card-body">
      <Form @submit="nextPage" :validation-schema="schema" v-slot="{errors}">
        <div class="p-fluid formgrid grid ">
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="selectedDso" required name="dso" as="select"  class="form-control" :class="{'is-invalid': errors.dso}">
              <option v-for="dso in dsoList.elements" :key="dso.id" :value="dso.id" >{{dso.name}}</option>
            </Field>
            <span>{{$t('FORMS.PLACEHOLDERS.DSO')}}</span>
            <div class="invalid-feedback">{{errors.dso}}</div>
          </div>
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="selectedTariff" required name="tariff" as="select" class="form-control" :class="{'is-invalid': errors.tariff}">
              <option v-for="tariff in tariffList" :key="tariff" :value="tariff.name">{{tariff.name}}</option>
            </Field>
            <span>{{$t('FORMS.PLACEHOLDERS.TARIFF_TYPE')}}</span>
            <div class="invalid-feedback">{{errors.tariff}}</div>
          </div>
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="name" name="name" placeholder=" " type="text" class="form-control" :class="{'is-invalid': errors.name}"></Field>
            <span>{{$t('FORMS.PLACEHOLDERS.NAME')}}</span>
            <div class="invalid-feedback">{{errors.name}}</div>
          </div>
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="startDate"  name="startDate" placeholder=" " type="date" class="form-control" :class="{'is-invalid': errors.startDate}"></Field>
            <span>{{$t('FORMS.PLACEHOLDERS.START_DATE')}}</span>
            <div class="invalid-feedback">{{errors.startDate}}</div>
          </div>
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="endDate"  name="endDate" placeholder=" " type="date" class="form-control" :class="{'is-invalid': errors.endDate}"></Field>
            <span>{{$t('FORMS.PLACEHOLDERS.END_DATE')}}</span>
            <div class="invalid-feedback">{{errors.endDate}}</div>
          </div>
        </div>
        <div class="grid grid-nogutter justify-content-between">
          <Button label="Next" type="submit" icon="pi pi-angle-right" icon-pos="right"></Button>
        </div>
      </Form>
    </div>
  </div>
</template>

<script setup lang="ts">
import {ref, defineEmits, PropType, onMounted} from 'vue';
import {TariffGroup} from "@/components/forms/tariff-group/TariffGroup";
import { Form, Field } from 'vee-validate';
import * as Yup from 'yup';
import {DataHolder} from "@/models/data-holder";
import {TariffGroupType} from "@/models/billing/billing";
import { useTariffGroupStore } from '@/store/tariffGroup.store';

const emit = defineEmits(['nextPage', 'prevPage', 'update:newTariffGroup']);

const props = defineProps({

  newTariffGroup: {
    type: Object as PropType<TariffGroup>,
    required: true
  }
});


const tariffGroupStore = useTariffGroupStore();
onMounted(()=> {
    tariffGroupStore.fetchDistributionNetworkOperator(true, null);
    dsoList.value = tariffGroupStore.distributionNetworkOperator;
  }
);

const selectedDso=ref();
const dsoList = ref<DataHolder<TariffGroup>>(Object(DataHolder));
const selectedTariff = ref();
const tariffList = ref([{id:1 ,name:TariffGroupType.G11}]);
const name = ref();
const startDate = ref(new Date());
const endDate = ref(new Date());

const schema = Yup.object().shape({
  dso: Yup.number().required('Dso is required'),
  tariff: Yup.string().required('Tariff is required'),
  name: Yup.string().required('Name is required'),
  startDate: Yup.string()
    .required('Start date is required')
    .matches(/^\d{4}-(0[1-9]|1[012])-(0[1-9]|[12][0-9]|3[01])$/, 'Start date must be a valid date in the format YYYY-MM-DD'),
  endDate: Yup.string()
    .required('End date is required')
    .matches(/^\d{4}-(0[1-9]|1[012])-(0[1-9]|[12][0-9]|3[01])$/, 'End date must be a valid date in the format YYYY-MM-DD')
  ,
});

const nextPage = () => {
  let dateStart = startDate.value.toString().split('-');
  let dateEnd = endDate.value.toString().split('-');

  let updatedNewTariffGroup: TariffGroup = props.newTariffGroup;
  updatedNewTariffGroup.distributionNetworkOperatorID = selectedDso.value;
  updatedNewTariffGroup.endDate = new Date(Number(dateEnd[0]), Number(dateEnd[1])-1, Number(dateEnd[2])+1, 1, 0, 0);
  updatedNewTariffGroup.startDate = new Date(Number(dateStart[0]), Number(dateStart[1])-1, Number(dateStart[2])+1, 1, 0, 0);
  updatedNewTariffGroup.name = name.value;
  updatedNewTariffGroup.tariffGroupLabelName = selectedTariff.value;

  emit('nextPage', {pageIndex: 0});
  emit('update:newTariffGroup', updatedNewTariffGroup);
};

</script>

<style scoped>

</style>
