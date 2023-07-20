import {RoleEnum} from "@/services/permissions/role-enum";

export declare class User{
  id: number;
  customerAccountId: number;
  workerId: number;
  providerId: number;
  roleId: RoleEnum;
  isActive: boolean;
  login: string;
  addedDate : string;
}
