import { Address, Offer } from "@/models/billing/billing";

export interface ContractForm {
  offer: Offer,
  bankAccountNumber?: string,
  serviceAccessPoints: FormServiceAccessPoint[],
  tpaParameter: string,
  clientType: string,
  registrationNumber: string,
  startDate: string
}

export interface UnitAmount {
  unit: string,
  amount: number
}

export interface CurrentSeller {
  name: string,
  noticePeriod: string
}

export interface ServiceAccessPoint {
  accountId: number,
  address: string,
  city: string,
  id: number,
  meterNumber: string,
  providerId: number,
  sapCode: string
}

export interface FormServiceAccessPoint {
  objectName: string,
  address: string,
  sapCode: string,
  meterNumber: string,
  osd: Osd,
  tariffGroup: string,
  declaredEnergyUsage: UnitAmount,
  estimatedEnergyUsage: UnitAmount,
  connectionPower: UnitAmount,
  contractedPower: UnitAmount,
  currentSeller: CurrentSeller,
  phase: string,
  sourceType: string,
  sourcePower: UnitAmount
}

export interface  Osd {
  name: string,
  branch: string
}
