import {ClientTypeEnum} from "@/models/billing/enum/client-type.enum";

export declare class Customer {

  id: number;
  providerId: number;
  workerId: number
  customerTypeName: string;
  status: boolean;

  firstName: string;
  lastName: string;

  country: string;
  province: string;
  city: string;
  postalCode: string;
  buildingNumber: string;
  apartmentNumber: string;
  street: string;
  email: string;
  phone: string;
  pesel: string;

}
