<template>
  <div class="container">

    <h2>{{ $t("SECTION_TITLES.EMPLOYEE_DETAILS_VIEW") }}</h2>
    <Button label="Edytuj" @click="editEmployee"></Button>
    <div class="detailed-data" style="margin-top: 20px">

        <span>
          <span><Icon name="User" /></span>
          <span>
            <div class="label-small">
              {{ $t("FORMS.PLACEHOLDERS.FULL_NAME") }}:
            </div>
            {{ employee?.firstName }} {{ employee?.lastName }}
          </span>
        </span>
      <span>
          <span><Icon name="AtSign" /></span>
          <span>
            <div class="label-small">{{ $t("FORMS.PLACEHOLDERS.EMAIL") }}:</div>
            {{ employee?.email }}
          </span>
        </span>
      <span>
          <span><Icon name="Smartphone" /></span>
          <span>
            <div class="label-small">{{ $t("FORMS.PLACEHOLDERS.PHONE") }}:</div>
            {{ employee?.phone }}
          </span>
        </span>
      <span>
          <span><Icon name="Contact" /></span>
          <span>
            <div class="label-small">{{ $t("FORMS.PLACEHOLDERS.ROLE") }}:</div>
            <span v-if="employee.role"> {{ $t('GLOBALS.ROLES.'+employee?.role) }} </span>
          </span>
        </span>
      <span>
          <span>
            <Icon v-if="employee?.status" name="CheckSquare" />
            <Icon v-if="!employee?.status" name="Square" />
          </span>
          <span>
            <div class="label-small">{{ $t("FORMS.PLACEHOLDERS.STATUS") }}:</div>
            {{ employee?.status ? "Aktywny" : "Nieaktywny" }}
          </span>
        </span>

    </div>

    <h2 style="margin-top: 20px">{{ $t("PROVIDERS.ADDRESS_DATA") }}</h2>
    <div class="detailed-data" style="margin-top: 20px">
      <span v-if="employee?.apartmentNumber"><span><Icon name="Home" /></span>
        <span>
          <div class="label-small">{{ $t("HOME_VIEW.FORMS.STREET") }}:</div>
          {{ getAddress(employee) }}
        </span>
      </span>

      <span>
        <span><Icon name="Inbox" /></span>
        <span>
          <div class="label-small">{{ $t("HOME_VIEW.FORMS.POSTAL_CODE") }}, {{ $t("HOME_VIEW.FORMS.CITY") }}:</div>
          {{ employee?.postalCode }}, {{ employee?.city }}
        </span>
      </span>
      <span>
        <span><Icon name="Map" /></span>
        <span>
          <div class="label-small">{{ $t("HOME_VIEW.FORMS.PROVINCE") }}:</div>
          {{ employee?.province }}</span>
      </span>

      <span>
        <span><Icon name="Globe" /></span>
        <span><div class="label-small">{{ $t("HOME_VIEW.FORMS.COUNTRY") }}:</div>
          {{ employee?.country }}
        </span>
      </span>
    </div>

    <h2 style="margin-top: 20px">{{ $t("PROVIDERS.ADDITIONAL_DATA") }}</h2>
    <div class="detailed-data" style="margin-top: 20px">
      <!-- TODO: What with premises number? -->
      <!-- TODO: Add blockchainAcc to Provider model  -->
      <span>
        <span>
          <Icon name="Link" /></span>
        <span>
          <div class="label-small">{{ $t("HOME_VIEW.FORMS.BLOCKCHAIN_ACC") }}:</div>
          {{ employee?.blockchainAccAddress }}
        </span>
      </span>

      <span>
        <span><Icon name="Calendar" /></span>
        <span><div class="label-small">{{ $t("HOME_VIEW.FORMS.WORK_START_DATE") }}:</div>
          {{ formatDate(employee?.workStartDate, DateFormat.DATE_FORMAT, '-') }}
        </span>
      </span>

      <span>
        <span><Icon name="None" /></span>
        <span><div class="label-small">{{ $t("HOME_VIEW.FORMS.KRS") }}:</div>
          {{ employee?.krs }}
        </span>
      </span>
      <span>
        <span><Icon name="Calendar" /></span>
        <span>
          <div class="label-small">{{ $t("HOME_VIEW.FORMS.WORK_END_DATE") }}:</div>
          {{ formatDate(employee?.workEndDate, DateFormat.DATE_FORMAT, '-') }}
        </span>
      </span>

      <span>
        <span><Icon name="None" /></span>
        <span>
          <div class="label-small">{{ $t("HOME_VIEW.FORMS.NIP") }}:</div>
          {{ employee?.nip }}
        </span>
      </span>
      <span>
        <span><Icon name="None" /></span>
        <span>
          <div class="label-small">{{ $t("HOME_VIEW.FORMS.REGON") }}:</div>
          {{ employee?.regon }}
        </span>
      </span>
      <span>
        <span><Icon name="None" /></span>
        <span><div class="label-small">{{ $t("HOME_VIEW.FORMS.EXTRA_INFO") }}:</div>
          {{ employee?.extraInfo }}
        </span>
      </span>


    </div>
  </div>
</template>

<script setup lang="ts">
import { Emplayee } from "@/models/emplayee";
import { computed } from "vue";
import { useContextStore } from "@/store/context.store";
import {useRouter} from "vue-router";
import {DateFormat} from "@/models/billing/billing";
import {formatDate} from '@/utils/date-formatter';
import {onBeforeMount} from "@vue/runtime-core";
import {useCustomerStore} from "@/store/customers.store";
import {useWorkersStore} from "@/store/workers.store";

const router = useRouter();

onBeforeMount(() => {
  const employeeId = useContextStore().selectedEmployee.id;
  useWorkersStore().fetchSelectedWorkerData(employeeId, true, null);
});
const employee = computed(() => {
  return useWorkersStore().selectedWorkerData;
});

//TODO:move to util
function getAddress(employee: Emplayee) {
  if (employee.apartmentNumber != null) {
    return (employee.street + " " + employee.buildingNumber + "/" + employee.apartmentNumber
    );
  } else {
    return employee.street + " " + employee.buildingNumber;
  }
}



const editEmployee = () => {
  const role = useContextStore().selectedEmployee.role;
  let name = '';
  if(role == 'ADMINISTRATOR_FULL') {
    name = 'edit_admin_full_personal_step';
  } else if (role == 'SUPER_AGENT'){
    name = 'edit_super_agent_personal_step';
  } else if (role == 'TRADER'){
    name = 'edit_trader_personal_step';
  } else if (role == 'AGENT'){
    name = 'edit_agent_personal_step';
  }
  router.push({name: name});
};

</script>

<style scoped lang="scss">
.personal-data-grid {
  display: grid;
  grid-template-columns: 40px 1fr 40px;
  max-width: 400px;

  span {
    height: 60px;
    position: relative;
    border-bottom: 1px solid #dfdfdf;
    display: flex;
    flex-direction: column;
    justify-content: center;

    .label-small {
      font-size: 14px;
      color: #ff0000;
    }

    .lucide-icon {
      color: #ff0000;
      box-sizing: content-box;
    }

    .edit-icon {
      cursor: pointer;
      border-radius: 10px;
      padding: 5px;

      &:hover {
        background-color: #ff0000;
        color: white;
      }
    }
  }
}
</style>
