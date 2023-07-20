import { Customer } from "../components/forms/create-customer/Customer";
import {RoleEnum} from "@/services/permissions/role-enum";

export  declare class Emplayee extends Customer {
    id: number;
    nip: string;
    krs: string;
    workStartDate: Date;
    workEndDate: Date;
    status: boolean;
    extraInfo: string;
    role: string
    regon: string;
    supervisor: number;

}
