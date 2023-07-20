<template>
  <div>
    <div class="card no-border sticky">
      <Steps :model="items" :readonly="true" />
    </div>
    <router-view v-slot="{Component}" v-bind:passedRoleID="passedRoleID"  v-bind:useSelectedEmployee="useSelectedEmployee"
                 :previousPath="previousPath" v-model:newFunctionalUser="newFunctionalUser" @prevPage="prevPage($event)" @nextPage="nextPage($event)" @complete="complete">
      <keep-alive>
        <component :is="Component" />
      </keep-alive>
    </router-view>
  </div>

</template>

<script setup lang="ts">
import Steps from 'primevue/steps';
import { ref, PropType } from 'vue';
import { useRouter } from 'vue-router';
import { useToast } from 'vue-toastification';
import { useI18n } from 'vue-i18n';
import { FunctionalUser } from '@/components/forms/create-func-user/FunctionalUser';
import { RoleEnum } from '@/services/permissions/role-enum';
import { useWorkersStore } from '@/store/workers.store';
import {ResponseError} from "@/models/request-response-api";
import { computed } from '@vue/reactivity';
import {useContextStore} from "@/store/context.store";
import {useUserStore} from "@/store/user.store";
import {onBeforeMount} from "@vue/runtime-core";
import {Emplayee} from "@/models/emplayee";

const router = useRouter();
const toast = useToast();
const i18n = useI18n();



const currentPath = router.currentRoute.value.path;

const props = defineProps({
  passedRoleID: {
      type: Number as unknown as PropType<RoleEnum>,
      required: true
  },
  useSelectedEmployee: {
    type: Boolean
  },
  previousPath: {
    type: String,
    required: true
  },
});
onBeforeMount(() => {
  console.log(props.passedRoleID)
  console.log(props.useSelectedEmployee)

  console.log(props.previousPath)
  if(props.useSelectedEmployee){
    const employee = useContextStore().selectedEmployee;
    newFunctionalUser.value.providerId = employee.providerId;
    newFunctionalUser.value.roleId = employee.roleId;
    newFunctionalUser.value.id = employee.id;
    newFunctionalUser.value.workerId = employee.workerId;
    newFunctionalUser.value.firstName = employee.firstName;
    newFunctionalUser.value.lastName = employee.lastName;
    newFunctionalUser.value.phone = employee.phone;
    newFunctionalUser.value.blockchainAccAddress = employee.blockchainAccAddress;
    newFunctionalUser.value.customerTypeName = employee.customerTypeName;
    newFunctionalUser.value.pesel = employee.pesel;
    newFunctionalUser.value.email = employee.email;
    newFunctionalUser.value.nip = employee.nip;
    newFunctionalUser.value.krs = employee.krs;
    newFunctionalUser.value.workStartDate = employee.workStartDate;
    newFunctionalUser.value.workEndDate = employee.workEndDate;
    newFunctionalUser.value.street = employee.street;
    newFunctionalUser.value.buildingNumber = employee.buildingNumber;
    newFunctionalUser.value.apartmentNumber = employee.apartmentNumber;
    newFunctionalUser.value.city = employee.city;
    newFunctionalUser.value.postalCode = employee.postalCode;
    newFunctionalUser.value.province = employee.province;
    newFunctionalUser.value.country = employee.country;

  }

});
let items = ref();
function setItems(){
  if (props.passedRoleID == RoleEnum.SUPER_ADMIN){
    items = computed(()=> [
    {
      label: i18n.t('FORMS.HEADERS.PERSONAL_DATA'),
      to: currentPath + ""
    }])
  }
  else {
    items = computed(()=> [
    {
      label: i18n.t('FORMS.HEADERS.PERSONAL_DATA'),
      to: currentPath + ""
    },
    {
      label: i18n.t('FORMS.HEADERS.ADDRESS'),
      to: currentPath + "/address_2",
    }]);
  }
}
setItems();
const newFunctionalUser = ref<FunctionalUser>({} as FunctionalUser);

const nextPage = (event: any) => {
  router.push(items.value[event.pageIndex + 1].to);
};
const prevPage = (event: any) => {
  router.push(items.value[event.pageIndex - 1].to);
};

const goToParentPath = (roleEnum: RoleEnum) => {
      switch(parseInt(roleEnum.toString())){
        case RoleEnum.SUPER_ADMIN:
          return '/super_admins';
        case RoleEnum.AGENT:
        case RoleEnum.SUPER_AGENT:
          return '/employees';
        default:
          return '/home';
    }
};

const complete = () => {
  console.log(props.useSelectedEmployee)
    if(props.useSelectedEmployee) {
      console.log(props.useSelectedEmployee)
      const employee = useContextStore().selectedEmployee;
      const updatedEmployee = {} as Emplayee;
      updatedEmployee.id = employee.id;
      updatedEmployee.nip = newFunctionalUser.value.nip;
      updatedEmployee.krs = newFunctionalUser.value.krs;
      updatedEmployee.workStartDate = newFunctionalUser.value.workStartDate;
      updatedEmployee.workEndDate = newFunctionalUser.value.workEndDate;
      updatedEmployee.status = employee.status;
      updatedEmployee.extraInfo = employee.extraInfo;
      updatedEmployee.role = employee.role;
      updatedEmployee.regon = employee.regon;
      updatedEmployee.providerId = employee.providerId;
      updatedEmployee.firstName = newFunctionalUser.value.firstName;
      updatedEmployee.lastName = newFunctionalUser.value.lastName;
      updatedEmployee.phone = newFunctionalUser.value.phone;
      updatedEmployee.blockchainAccAddress = employee.blockchainAccAddress;
      updatedEmployee.pesel = newFunctionalUser.value.pesel;
      updatedEmployee.email = newFunctionalUser.value.email;
      updatedEmployee.street = newFunctionalUser.value.street;
      updatedEmployee.buildingNumber = newFunctionalUser.value.buildingNumber;
      updatedEmployee.apartmentNumber = newFunctionalUser.value.apartmentNumber;
      updatedEmployee.city = newFunctionalUser.value.city;
      updatedEmployee.postalCode = newFunctionalUser.value.postalCode;
      updatedEmployee.province = newFunctionalUser.value.province;
      updatedEmployee.country = newFunctionalUser.value.country;
      updatedEmployee.roleId = employee.roleId;
      updatedEmployee.supervisor = employee.supervisor;
      useWorkersStore().updateWorkerById(updatedEmployee.id, updatedEmployee, onSuccessEdit, onFailEdit);

    } else {
      newFunctionalUser.value.password = ''; //must be set to empty string
      if(useUserStore().workerId != undefined) {
        newFunctionalUser.value.supervisor = Number.parseInt(useUserStore().workerId);
        // newFunctionalUser.value.workerId = Number.parseInt(useUserStore().workerId);
        useWorkersStore().createWorker(newFunctionalUser.value, onSuccess, onFail);
        goToParentPath(props.passedRoleID);
      } else {
        console.log("workerId not defined");
      }
    }


};

const onSuccess = () => {
  toast.success(i18n.t('GLOBALS.TOASTS.ADDED_WORKER'));
  router.push('/home'); //TODO: Fix pathing after createing worker

};

const onFail = (error: ResponseError ) => {
  toast.error(i18n.t('GLOBALS.TOASTS.ERROR_POST')  + "\n" + error.status + " " + error.message);
};
const onSuccessEdit = () => {
  toast.success("Pomyślnie zaktualizowano pracownika");
  router.push(props.previousPath);

};

const onFailEdit = () => {
  toast.error("Edycja pracownika nie powiodła się");
};
</script>




<style scoped lang="scss">

</style>
