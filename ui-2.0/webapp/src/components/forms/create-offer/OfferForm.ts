import { OfferDraft } from "@/models/billing/billing";

export interface OfferForm {
  estimatedAnnualElectricityConsumption: EstimatedAnnualElectricityConsumption,
  estimatedAnnualElectricityProduction: EstimatedAnnualElectricityConsumption;

  numberOfPPE: number,
  startDate: Date,
  endDate: Date
  offerDraft: OfferDraft
}

export interface EstimatedAnnualElectricityConsumption {
  unit: string,
  amount: number
}


