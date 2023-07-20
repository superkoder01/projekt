<template>
  <div class="container">
    <h2>{{ $t('SECTION_TITLES.CUSTOMER_TABS') }}</h2>
    <Button label="Edytuj" @click="editCustomer"></Button>
    <div class="detailed-data" style="margin-top: 20px">
      <span><span><Icon name="User"/></span> <span><div class="label-small">{{ $t("FORMS.PLACEHOLDERS.FULL_NAME") }}:</div>{{ customer.firstName }} {{ customer.lastName }}</span></span>
      <span><span><Icon name="AtSign"/></span> <span><div class="label-small">{{ $t("FORMS.PLACEHOLDERS.EMAIL") }}:</div>{{ customer.email }}</span></span>
      <span><span><Icon name="Smartphone"/></span> <span><div class="label-small">{{ $t("FORMS.PLACEHOLDERS.PHONE") }}:</div> {{ customer.phone }}</span></span>

      <!-- @formatter:off -->
        <span v-if="customer.bankAccNumber"><span><Icon name="Banknote"/></span>
        <span><div class="label-small">{{ $t("HOME_VIEW.FORMS.BANK_ACCOUNT_NUMBER") }}:</div>{{ customer.bankAccNumber }}</span></span>

        <span v-if="customer.pesel"><span><Icon name="Contact"/></span>
        <span><div class="label-small">{{ $t("HOME_VIEW.FORMS.PESEL") }}:</div>{{ customer.pesel }}</span></span>

        <span v-if="customer.nip"><span><Icon name="Album"/></span>
        <span><div class="label-small">{{ $t("HOME_VIEW.FORMS.NIP") }}:</div>{{ customer.nip }}</span></span>

        <span v-if="customer.regon"><span><Icon name="Book"/></span>
        <span><div class="label-small">{{ $t("HOME_VIEW.FORMS.REGON") }}:</div>{{ customer.regon }}</span></span>

        <!-- @formatter:on -->
    </div>
    <h2 style="margin-top: 20px">{{$t('PROVIDERS.ADDRESS_DATA')}}</h2>
    <div class="detailed-data" style="margin-top: 20px">
      <span><span><Icon name="Home"/></span>
        <span><div class="label-small">{{ $t("HOME_VIEW.FORMS.STREET") }}:</div>{{getAddress(customer)}}</span></span>

        <span><span><Icon name="Inbox"/></span>
          <span><div class="label-small">{{ $t("HOME_VIEW.FORMS.POSTAL_CODE") }}, {{ $t("HOME_VIEW.FORMS.CITY") }}:</div>{{ customer.postalCode }}, {{ customer.city }}</span></span>

        <span><span><Icon name="Map"/></span>
        <span><div class="label-small">{{ $t("HOME_VIEW.FORMS.PROVINCE") }}:</div>{{ customer.province }}</span></span>

        <span><span><Icon name="Globe"/></span>
        <span><div class="label-small">{{ $t("HOME_VIEW.FORMS.COUNTRY") }}:</div>{{ customer.country }}</span></span>

    </div>
  </div>
</template>

<script setup lang="ts">
import {Customer} from "@/models/customer";
import {computed, onMounted} from "vue";
import {useCustomerStore} from "@/store/customers.store";
import EmptySpan from "@/components/utils/EmptySpan.vue";
import {useContextStore} from "@/store/context.store";
import {useRouter} from "vue-router";

const customer = computed(() => useCustomerStore().selectedCustomerData);
const router = useRouter();
const currentPath = router.currentRoute.value.path;

onMounted(() => {
  const customerId = useContextStore().selectedCustomer.id;
  useCustomerStore().fetchSelectedCustomerData(customerId, true, null);
});

//TODO:move to util
function getAddress(customer: Customer) {
  if (customer.apartmentNumber != null) {
    return customer.street + " " + customer.buildingNumber + "/" + customer.apartmentNumber;
  } else {
    return customer.street + " " + customer.buildingNumber;
  }
}

const editCustomer = () => {
  router.push(currentPath+"/edit_customer");
}


</script>

<style scoped lang="scss">

</style>
