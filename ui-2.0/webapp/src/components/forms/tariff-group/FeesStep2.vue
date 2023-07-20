<template>
  <div class="card no-border m-3">
    <div class="card-body">
      <Form @submit="complete" v-slot="{errors}">
        <div class="p-fluid formgrid grid ">
          <div v-for="(fee, index) of feeList" :key="fee.id" class="field col-12 md:col-6 mb-4">
            <Field v-model="definedFees[index]" :rules="isRequired" :name="`fee${index}`" placeholder=" " type="text" class="form-control" :class="{'is-invalid': errors[`fee${index}`]}"></Field>
            <span>{{fee.name}}</span>
            <div class="invalid-feedback">{{errors[`fee${index}`]}}</div>
          </div>
        </div>
        <div class="grid grid-nogutter justify-content-between">
          <Button label="Back" @click="prevPage"  icon="pi pi-angle-left"></Button>
          <Button label="Complete" type="submit" icon="pi pi-angle-right"></Button>
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
import { useTariffGroupStore } from '@/store/tariffGroup.store';

const emit = defineEmits(['complete', 'prevPage','update:newTariffGroup']);
const props = defineProps({

  newTariffGroup: {
    type: Object as PropType<TariffGroup>,
    required: true
  }
});
const isRequired = Yup.number().required('This field is required');
const tariffGroupStore = useTariffGroupStore();
onMounted(
  ()=>{
    tariffGroupStore.fetchFees(true, null);
    feeList.value = tariffGroupStore.fees;
  }
);

const feeList = ref<[{id:number, name:string}]>([{id:-1, name:''}]);

const definedFees = ref([]);

const prevPage = () => {
  emit('prevPage', {pageIndex: 1});
};

const complete = () => {

  let fees:{nameId:number, price: number}[] = [];
  feeList.value.forEach((fee, index) => {
    fees.push({nameId: fee.id, price: Number(definedFees.value[index])});
  });
  let updatedNewTariffGroup: TariffGroup = props.newTariffGroup;

  updatedNewTariffGroup.fees = fees;
  emit('complete');
  emit('update:newTariffGroup', updatedNewTariffGroup);
};
</script>

<style scoped>

</style>
