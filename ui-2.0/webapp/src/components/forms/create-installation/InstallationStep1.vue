<template>
  <div class="card no-border m-3">
    <h5 class="card-header">{{ $t('FORMS.HEADERS.INSTALLATION_DATA') }} </h5>
    <div class="card-body">
      <Form @submit="nextPage" :validation-schema="schema" v-slot="{errors}">
        <div class="p-fluid formgrid grid ">
          <div class="field col-12 md:col-6 mb-4">
                <Field v-model="sapCode" name="sapCode" placeholder=" " type="text" class="form-control" :class="{'is-invalid': errors.sapCode}"></Field>
                <span>{{$t('FORMS.PLACEHOLDERS.SAP_CODE')}}</span>
                <div class="invalid-feedback">{{errors.sapCode ? $t(errors.sapCode) : ''}}</div>
            </div>
            <div class="field col-12 md:col-6 mb-4">
                <Field v-model="meterNumber" name="meterNumber" placeholder=" " type="text" class="form-control" :class="{'is-invalid': errors.meterNumber}"></Field>
                <span>{{$t('FORMS.PLACEHOLDERS.METER_NUMBER')}}</span>
                <div class="invalid-feedback">{{errors.meterNumber ? $t(errors.meterNumber) : ''}}</div>
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
import {ref, defineEmits, PropType} from 'vue';
import { Form, Field } from 'vee-validate';
import * as Yup from 'yup';
import { Installation } from './Installation';

const sapCode = ref();
const meterNumber = ref();

const emit = defineEmits(['nextPage', 'prevPage', 'update:newInstallation']);

const props = defineProps({
  newInstallation: {
    type: Object as PropType<Installation>,
    required: true
  }
});

const schema = Yup.object().shape({
  sapCode: Yup.string().required('FORMS.VALIDATION_ERRORS.REQUIRED'),
  meterNumber: Yup.string().required('FORMS.VALIDATION_ERRORS.REQUIRED'),
});

const nextPage = () => {

  let updatedInstallation: Installation = props.newInstallation;

    updatedInstallation.sapCode = sapCode.value;
    updatedInstallation.meterNumber = meterNumber.value;


  emit('nextPage', {pageIndex: 0});
  emit('update:newInstallation', updatedInstallation);
};

</script>

<style scoped>

</style>
