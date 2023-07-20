<template>
  <Dropdown v-model="selectedRole" :options="availableRole" optionLabel="roleName" placeholder="ROLE"
            @change="roleChanged"/>
</template>

<script setup lang="ts">
import Dropdown from 'primevue/dropdown';
import {RoleEnum} from "@/services/permissions/role-enum";
import {ref} from "vue";
import {LoginData} from "@/models/login-data";
import {useRouter} from "vue-router";
import {useToast} from "vue-toastification";
import {useUserStore} from "@/store/user.store";
import {PermissionsService} from "@/services/permissions/permissions.service";

const selectedRole = ref();

const permissionsService = new PermissionsService();
const toast = useToast();
const router = useRouter();

const availableRole = [
  {
    roleName: RoleEnum[RoleEnum.SUPER_ADMIN],
    login: "x",
    password: "x"
  },
  {
    roleName: RoleEnum[RoleEnum.ADMINISTRATOR_FULL],
    login: "adminfull",
    password: "x"
  },
  {
    roleName: RoleEnum[RoleEnum.ADMINISTRATOR_BASIC],
    login: "",
    password: ""
  },
  {
    roleName: RoleEnum[RoleEnum.TRADER],
    login: "",
    password: ""
  },
  {
    roleName: RoleEnum[RoleEnum.SUPER_AGENT],
    login: "agent",
    password: "x"
  },
  {
    roleName: RoleEnum[RoleEnum.AGENT],
    login: "",
    password: ""
  },
  {
    roleName: RoleEnum[RoleEnum.PROSUMER],
    login: "x",
    password: "x"
  }
];

function roleChanged() {
  console.log("Selected role:" + selectedRole.value);
  const loginData: LoginData = ({login: selectedRole.value.login, password: selectedRole.value.password});
  useUserStore().login(loginData, onSuccessLogin, onFailLogin);
}

const onSuccessLogin = () => {
  toast.success("Zmieniona role na: " + selectedRole.value.roleName);
  if (router.currentRoute.value.name === 'login') {
    router.push('home');
  } else {
    router.push(permissionsService.getAfterLoginPage());
  }
};

const onFailLogin = () => {
  toast.error("Błąd podczas logowania");
};

</script>

<style scoped>

</style>
