import {ServiceTypeEnum} from "@/models/billing/enum/service-type.enum";
import {
  EstimatedAnnualElectricityConsumption,
} from "@/components/forms/create-offer/OfferForm";
import { CurrentSeller, FormServiceAccessPoint } from "@/components/forms/create-contract/Contract";

export interface PricingHolder{
  id: string,
  header: any,
  payload: Pricing
}


export interface Pricing {
  name: string,
  id: string,
  type: ServiceTypeEnum,
  osd: string,
  tariffGroup: string,
  zones: Zone[],
  commercialFee: CommercialFee[],
  price: Price
}

export interface OfferDraft {
  id: string,
  header: Header,
  payload: OfferDraftPayload
}

export interface Offer {
  id: string,
  header: Header,
  payload: OfferPayload
}

export interface Header {
  version: string,
  provider: string,
  content: Content
}

export interface Content {
  type: string,
  category: string
}

export interface OfferDraftPayload {
  offerDetails: OfferDraftDetails,
  conditions: OfferDraftConditions,
  priceList: OfferPriceList,
  repurchase: Repurchase
}

export interface OfferDraftDetails {
  title: string,
  type: string,
  creationDate: string,
  tariffGroup: TariffGroupType,
  agreementType: string
}

export interface OfferDraftConditions {
  duration: DeliveryPeriod,
  billingPeriod: Duration,
  invoiceDueDate: string,
  startDate: string,
  endDate: string
}

export interface Duration {
  calendarUnit: CalendarUnitType,
  number: string
}

export interface DeliveryPeriod {
  calendarUnit: CalendarUnitType,
  number: string,
}

export interface CommercialFee {
  from: string,
  to?: string,
  unit: CommercialFeeUnitType,
  price: Price,
}

export interface OfferPriceList {
  name: string,
  id: string,
  type: string,
  startDate: string,
  endDate: string,
  osd: string,
  tariffGroup: string,
  zones: Zone[],
  commercialFee: OfferCommercialFee[]
}

export interface OfferCommercialFee {
  from: string,
  to?: string,
  unit: CommercialFeeUnitType,
  price: Price,
}

export interface Price {
  calendarUnit?: CalendarUnitType,
  cost?: number,
  currency?: string
}
export interface RepurchasePrice {
  unit?: CommercialFeeUnitType,
  cost?: number,
  currency?: string
}

export interface Repurchase {
  name: string,
  type: string,
  id: string,
  price?: RepurchasePrice
}

export interface Zone {
  id: string,
  name: string,
  unit: string,
  cost?: number,
  currency: string
}

export interface Excise {
  calendarUnit: string,
  cost: number,
  currency: string
}

export interface Contract {
  id?: string,
  header: Header,
  payload: ContractPayload
}

export interface ContractPayload {
  contractDetails: ContractDetails,
  sellerDtls: Seller,
  customerDtls: Customer,
  conditions: ContractConditions,
  serviceAccessPoints: FormServiceAccessPoint[],
  priceList: OfferPriceList[],
  repurchase: Repurchase
}

export interface ContractConditions {
  signatureDate: string,
  startDate: string,
  endDate: string,
  duration: Duration,
  billingPeriod: Duration,
  invoiceDueDate: string,
  estimatedAnnualElectricityConsumption: UnitAmount
}

export interface UnitAmount {
  unit: string,
  amount: number
}

export interface OfferPayload {
  offerDetails: OfferDetails,
  conditions: Conditions,
  priceList: OfferPriceList[],
  repurchase: Repurchase,
  sellerDtls: Seller,
  customerDtls: Customer
}

export interface OfferDetails {
  title: string,
  type: string,
  number: string,
  offerDraftId: string,
  creationDate: string,
  status: string,
  customerId: string,
  tariffGroup: string,
  agreementType: string
}
export interface ContractDetails {
  title: string,
  type: string,
  number: string,
  offerId : string,
  creationDate: string,
  state: string,
  customerId: string,
  tariffGroup: string,
  agreementType: string,
  tpaParameter: string,
  clientType: string
}
export interface Seller {
  legalName: string,
  displayName: string,
  krs: string,
  nip: string,
  regon?: string,
  bankAccountNumber?: string,
  address: Address,
  contact: Contact,
}
export interface Address {
  street: string,
  postCode: string,
  city: string
}

export interface Contact {
  address: Address,
  phoneNumbers: PhoneNumber[],
  email: string,
  www: string
}
export interface PhoneNumber {
  type: string,
  number: string
}

export interface Customer {
  customerId: string,
  firstName: string,
  lastName: string,
  pesel?: string,
  displayName: string,
  address: Address,
  contact: Contact
}

export interface Conditions {
  duration: DeliveryPeriod,
  billingPeriod: Duration,
  invoiceDueDate: string,
  estimatedAnnualElectricityConsumption: EstimatedAnnualElectricityConsumption,
  estimatedAnnualElectricityProduction: EstimatedAnnualElectricityConsumption
  numberOfPPE: string,
  offerActivePeriod: OfferActivePeriod
}

export interface OfferActivePeriod {
  startDate: string,
  endDate: string
}
export enum CalendarUnitType {
  MONTH='month'
}

export function toCalendarUnitTypeEnumKey(value: CalendarUnitType): string{
  return Object.keys(CalendarUnitType)[Object.values(CalendarUnitType).indexOf(value)];
}

export enum CommercialFeeUnitType{
  KW = 'kW',
  KWP = "kWp"
}
export enum DateFormat{
  DATE_FORMAT = "DD-MM-YYYY",
  SEND_DATE_FORMAT="YYYY/MM/DD",
  SEND_DATE_FORMAT_WITH_TIME="YYYY/MM/DD HH:mm:ss"
  // SEND_DATE_FORMAT="DD/MM/YYYY"
}

export enum TariffGroupType{
  G11='G11'
}

export enum UnitOfPower {
  kW = 'kW',
  kWp = 'kWp'
}
export enum UnitOfEnergy {
  MWh ='MWh'
}
