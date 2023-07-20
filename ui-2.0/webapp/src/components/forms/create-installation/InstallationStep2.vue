<template>
  <div class="card no-border m-3">
    <h5 class="card-header">{{ $t('FORMS.HEADERS.ADDRESS') }} </h5>
    <div class="card-body">
      <Form @submit="complete" :validation-schema="schema" v-slot="{errors}">
        <div class="p-fluid formgrid grid ">
          <div class="field col-12 md:col-6 mb-4">
                <Field v-model="address" name="address" placeholder=" " type="text" class="form-control" :class="{'is-invalid': errors.address}"></Field>
                <span>{{$t('FORMS.PLACEHOLDERS.ADDRESS')}}</span>
                <div class="invalid-feedback">{{errors.address ? $t(errors.address) : ''}}</div>
            </div>
            <div class="field col-12 md:col-6 mb-4">
                <Field v-model="city" name="city" placeholder=" " type="text" class="form-control" :class="{'is-invalid': errors.city}"></Field>
                <span>{{$t('FORMS.PLACEHOLDERS.CITY')}}</span>
                <div class="invalid-feedback">{{errors.city ? $t(errors.city) : ''}}</div>
            </div>
        </div>
        <div class="grid grid-nogutter justify-content-between">
          <Button :label="$t('GLOBALS.BUTTONS.BACK')" @click="prevPage"  icon="pi pi-angle-left"></Button>
          <Button :label="$t('GLOBALS.BUTTONS.CREATE')" type="submit" icon="pi pi-angle-right" icon-pos="right"></Button>
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

const address = ref();
const city = ref();

const emit = defineEmits(['complete', 'prevPage', 'update:newInstallation']);

const props = defineProps({
  newInstallation: {
    type: Object as PropType<Installation>,
    required: true
  }
});

const schema = Yup.object().shape({
  address: Yup.string().required('FORMS.VALIDATION_ERRORS.REQUIRED'),
  city: Yup.string().required('FORMS.VALIDATION_ERRORS.REQUIRED'),
});

const prevPage = () => {
  emit('prevPage', {pageIndex: 1});
};

const complete = () => {

  let updatedInstallation: Installation = props.newInstallation;

  updatedInstallation.address = address.value;
  updatedInstallation.city = city.value;


  emit('update:newInstallation', updatedInstallation);
  emit('complete');
};

</script>

<style scoped>

</style>
