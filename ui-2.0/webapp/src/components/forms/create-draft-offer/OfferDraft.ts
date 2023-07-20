import { DeliveryPeriod, Duration, Pricing, TariffGroupType } from "@/models/billing/billing";

export interface FormOfferDraft {
  title: string,
  type: string,
  agreementType: string,
  price?: Pricing,
  repurchase?: Pricing,
  duration: DeliveryPeriod,
  billingPeriod: Duration,
  invoiceDueDate: number,
  tariffGroup: TariffGroupType,
  startDate: string,
  endDate: string
}
