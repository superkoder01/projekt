<template>

  <div class="card m-3">
    <div class="card-body">
      <Form @submit="nextPage" :validation-schema="schema" v-slot="{errors}">
        <div class="p-fluid formgrid grid ">
          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="selectedOsd" required name="selectedOsd" as="select" class="form-control"
                   :class="{'is-invalid': errors.selectedOsd}">
              <option v-for="osd in osdList" :key="osd" :value="osd">{{ osd }}</option>
            </Field>
            <span>OSD</span>
            <div class="invalid-feedback">{{ errors.selectedOsd ? $t(errors.selectedOsd) : '' }}</div>
          </div>

          <div class="field col-12 md:col-6 mb-4" v-if="selectedOsd === 'PGE'">
            <Field v-model="selectedBranch" required name="selectedBranch" as="select" class="form-control"
                   :class="{'is-invalid': errors.selectedBranch}">
              <option v-for="branch in branchList" :key="branch" :value="branch">{{ branch }}</option>
            </Field>
            <span>{{$t("FORMS.PLACEHOLDERS.BRANCH")}}</span>
            <div class="invalid-feedback">{{ errors.selectedBranch ? $t(errors.selectedBranch) : '' }}</div>
          </div>

          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="tariffGroup" name="tariffGroup" placeholder=" " type="text" class="form-control"
                   :class="{'is-invalid': errors.tariffGroup}" disabled="true"/>
            <span>{{$t("FORMS.PLACEHOLDERS.TARIFF_GROUP")}}</span>
            <div class="invalid-feedback">{{ errors.tariffGroup }}</div>
          </div>

          <div class="field col-12 md:col-6 mb-4">
            <Field v-model="invoiceDueDate" name="invoiceDueDate" type="text" class="form-control"
                   :class="{'is-invalid': errors.invoiceDueDate}" disabled></Field>
            <span>{{$t("FORMS.PLACEHOLDERS.INVOICE_DUE_DATE")}}</span>
            <div class="invalid-feedback">{{ errors.invoiceDueDate }}</div>
          </div>



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
import {defineEmits, PropType, ref} from 'vue';
import {Field, Form} from 'vee-validate';
import * as Yup from 'yup';
import { ContractForm, Osd } from "@/components/forms/create-contract/Contract";
import {onBeforeMount} from "@vue/runtime-core";


const emit = defineEmits(['nextPage', 'prevPage', 'update:newContract']);

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
    selectedOsd.value = props.newContract.serviceAccessPoints[0].osd.name;
    selectedBranch.value = props.newContract.serviceAccessPoints[0].osd.branch;
    console.log(props.newContract.serviceAccessPoints)
  }
});
const osdList = ref(['PGE', 'Energa', 'Enea', 'Tauron']);
const selectedOsd = ref();
const branchList = ref(['Warszawa', 'Lublin','Białystok', 'Łódź Teren', 'Łódź miasto', 'Zamość', 'Rzeszów', 'Skarżysko- Kamienna']);
const selectedBranch = ref();
const tariffGroup = ref(props.newContract.offer.payload.offerDetails.tariffGroup);
const invoiceDueDate = ref(props.newContract.offer.payload.conditions.invoiceDueDate);
const schema = Yup.object().shape({
  selectedOsd: Yup.string()
    .required('FORMS.VALIDATION_ERRORS.OSD'),
  selectedBranch: Yup.string()
    .when('selectedOsd', {
      is: 'PGE', then: Yup.string().required('FORMS.VALIDATION_ERRORS.BRANCH')
    })
});


const prevPage = () => {
  emit('prevPage', {pageIndex: 1});
};

const nextPage = () => {
  let updatedNewContract: ContractForm = props.newContract;
  if(props.useSelectedContract) {
    updatedNewContract.serviceAccessPoints[0].osd.name = selectedOsd.value;
  } else {
    updatedNewContract.serviceAccessPoints[0].osd = {} as Osd;
    updatedNewContract.serviceAccessPoints[0].osd.name = selectedOsd.value;
    updatedNewContract.serviceAccessPoints[0].tariffGroup = props.newContract.offer.payload.offerDetails.tariffGroup;
  }
  if(selectedOsd.value === 'PGE') {
    updatedNewContract.serviceAccessPoints[0].osd.branch = selectedBranch.value;
  } else{
    updatedNewContract.serviceAccessPoints[0].osd.branch ='';
  }


  emit('nextPage', {pageIndex: 1});
  emit('update:newContract', updatedNewContract);
};

</script>
<style scoped lang="scss">

</style>
