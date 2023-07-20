<template>
  <div class="container">
    <h2>{{ $t("SECTION_TITLES.ADMIN_DETAILS") }}</h2>
    <Button disabled label="Edytuj" @click="editProviderAdmin"></Button>
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
        <span><Icon name="User" /></span>
         <span>
          <div class="label-small">{{ $t("FORMS.PLACEHOLDERS.LOGIN") }}:</div>
          {{ employee?.login }}
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
        <span><Icon name="AtSign" /></span>
        <span>
          <div class="label-small">{{ $t("FORMS.PLACEHOLDERS.ROLE") }}:</div>
          <span v-if="employee.roleId"> {{ $t('GLOBALS.ROLES.'+RoleEnum[employee.roleId]) }} </span>
        </span>
      </span>
      <span>
        <span>
          <Icon v-if="employee?.isActive" name="CheckSquare" />
          <Icon v-if="!employee?.isActive" name="Square" />
        </span>
        <span>
          <div class="label-small">{{ $t("FORMS.PLACEHOLDERS.STATUS") }}:</div>
          {{ employee?.isActive ? $t('TABLES.SUPER_ADMINS.STATE.ACTIVE') : $t('TABLES.SUPER_ADMINS.STATE.INACTIVE') }}
        </span>
      </span>
    </div>
  </div>
</template>

<script setup lang="ts">
import { computed } from "vue";
import { useContextStore } from "@/store/context.store";
import { RoleEnum } from "@/services/permissions/role-enum";
import {useRouter} from "vue-router";

const contextStore = useContextStore();

const employee = computed(() => {
  return contextStore.selectedEmployee;
});
const router = useRouter();

const editProviderAdmin = () => {

  router.push({name: 'edit_admin_personal_step'});
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
