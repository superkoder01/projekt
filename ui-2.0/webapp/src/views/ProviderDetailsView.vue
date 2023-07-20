<template>
  <div class="container">
    <h2>{{ $t('SECTION_TITLES.PROVIDER_TABS') }}</h2>
    <Button v-if="permissionService.canAccessFeature(FeatureEnum.MANAGE_PARTNERS)"  label="Edytuj" @click="editProvider"></Button>
    
      <div class="detailed-data" style="margin-top: 20px">
        <span><span><Icon name="Info"/></span> <span><div class="label-small">{{ $t("FORMS.PLACEHOLDERS.TYPE") }}:</div>{{ provider.type }}</span></span>

        <span><span><Icon name="Building"/></span> <span><div class="label-small">{{ $t("FORMS.PLACEHOLDERS.NAME") }}:</div>{{ provider.name }}</span></span>

        <span><span><Icon name="AtSign"/></span> <span><div class="label-small">{{ $t("FORMS.PLACEHOLDERS.EMAIL") }}:</div>{{ provider.email }}</span></span>

        <span><span><Icon name="Phone"/></span> <span><div class="label-small">{{ $t("FORMS.PLACEHOLDERS.PHONE") }}:</div>{{ provider.phoneNumber }}</span></span>

        <span><span><Icon name="Album"/></span>
        <span><div class="label-small">{{ $t("HOME_VIEW.FORMS.NIP") }}:</div>{{ provider.nip }}</span></span>

        <span><span><Icon name="Archive"/></span>
        <span><div class="label-small">{{ $t("HOME_VIEW.FORMS.KRS") }}:</div>{{ provider.krs }}</span></span>

        <span><span><Icon name="Book"/></span>
        <span><div class="label-small">{{ $t("HOME_VIEW.FORMS.REGON") }}:</div>{{ provider.regon }}</span></span>

        <!-- TODO: Add date format  -->
        <span><span><Icon name="TimerReset"/></span>
        <span><div class="label-small">{{ $t("HOME_VIEW.FORMS.LICENSE_EXPIRATION_DATE") }}:</div>{{ formatDate(provider.licenseExpirationDate,DateFormat.DATE_FORMAT,'-') }}</span></span>

        <span><span><Icon v-if="provider.status" name="CheckSquare"/> <Icon v-if="!provider.status" name="Square"/> </span> <span><div class="label-small">{{ $t("FORMS.PLACEHOLDERS.STATUS") }}:</div>{{ provider.status ? 'Aktywny' : 'Nieaktywny' }}</span></span>
      </div>
      <h2 style="margin-top: 20px">{{$t('PROVIDERS.ADDRESS_DATA')}}</h2>
      <div class="detailed-data" style="margin-top: 20px">

        <!-- TODO: What with premises number? -->
        <span><span><Icon name="Home"/></span> <span><div class="label-small">{{ $t("FORMS.PLACEHOLDERS.ADDRESS") }}:</div>{{ getAddress(provider.street, provider.buildingNumber, provider.apartmentNumber) }} </span></span>

        <span><span><Icon name="Inbox"/></span>
        <span><div class="label-small">{{ $t("HOME_VIEW.FORMS.POSTAL_CODE") }}, {{$t("HOME_VIEW.FORMS.CITY")}}:</div>{{ provider.postalCode }}, {{ provider.city }}</span></span>

        <span><span><Icon name="Map"/></span>
        <span><div class="label-small">{{ $t("HOME_VIEW.FORMS.PROVINCE") }}:</div>{{ provider.province }}</span></span>

        <span><span><Icon name="Globe"/></span> <span><div class="label-small">{{ $t("HOME_VIEW.FORMS.COUNTRY") }}:</div> {{ provider.country }} </span></span>
    </div>

    <h2 style="margin-top: 20px">{{$t('PROVIDERS.ADDITIONAL_DATA')}}</h2>
      <div class="detailed-data" style="margin-top: 20px">
        <!-- TODO: What with premises number? -->
              <!-- TODO: Add blockchainAcc to Provider model  -->
        <span><span><Icon name="Link"/></span>
        <span><div class="label-small">{{ $t("HOME_VIEW.FORMS.BLOCKCHAIN_ACC") }}:</div>{{ provider?.blockchainAccAddress }}</span></span>
      </div>
  </div>

</template>

<script setup lang="ts">

import {computed} from "vue";
import {useContextStore} from "@/store/context.store";
import {formatDate} from '@/utils/date-formatter';
import {DateFormat} from "@/models/billing/billing";
import {useRouter} from "vue-router";
import {PermissionsService} from "@/services/permissions/permissions.service";
import {FeatureEnum} from "@/services/permissions/feature-enum";

const props = defineProps({
  useSelectedProvider: {
    type: Boolean,
    required: false
  }
});

const router = useRouter();
const context = useContextStore();
const permissionService = new PermissionsService();

const provider = computed(() => {
  if (props.useSelectedProvider == true) {
    return context.selectedProvider;
  } else {
    return context.currentLoggedProvider;
  }
})

function getAddress(street: string, building: string, apartment: string): string {
  if (apartment === undefined) {
    return street + ' ' + building;
  }
  return street + ' ' + building + '/' + apartment;
}

const editProvider = () => {

  router.push({name: 'edit_provider_data_step'});
};

</script>

<style scoped lang="scss">
</style>
