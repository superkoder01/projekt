import {CommercialFee} from "@/models/billing/billing";
import {ServiceTypeEnum} from "@/models/billing/enum/service-type.enum";

export interface FormPricing {
  name: string,
  type: ServiceTypeEnum,
  tariffGroup: string,
  zones: number,
  commercialFee: CommercialFee[],
  fixedPrice: number
}
