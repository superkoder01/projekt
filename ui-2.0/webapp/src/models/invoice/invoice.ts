import { Address, Contact, Customer, Seller } from "../billing/billing";

export interface Invoice{
    header: any,
    payload: ProsumentPayload,
}
export interface ProsumentPayload {
    invoiceDetails: InvoiceDetails,
    sellerDetails: Seller,
    customerDetails: Customer,
    paymentDetails: PaymentDetails,
    payerDetails: PayerDetails,
    ppeDetails: PpeItem[],
    activeEnergyConsumed: ActiveEnergyConsumed,
    activeEnergyProduced: ActiveEnergyProduced,
    depositSummary: PpeDeposit,
    excessSalesBalance: ExcessSalesBalance,
    sellSummary: SellSummary,
    paymentSummary: PaymentSummary,
    energyValueAnnualBalance: EnergyAnnualBalance,
    energyAmountAnnualBalance: EnergyAnnualBalance,
    ppeSummary: PpeSummary
}
export interface InvoiceDetails {
    number: string,
    issueDt: Date,
    serviceDt: string,
    type: string,
    customerId: string,
    billingStartDt: string,
    billingEndDt: string, 
    catg: string,
    status: string, 
    timeZone: string,
}
export interface PaymentDetails {
    bankDetails: BankDetails,
    paymentTitle: string, 
    paymentDueDt: Date, 
}
export interface PayerDetails {
    customerId: string,
    firstName: string, 
    lastName: string, 
    displayName: string, 
    nip: string, 
    regon: string, 
    address: Address,
    contact: Contact
}
export interface PpeItem {
    ppeCode: string,
    ppeName: string,
    ppeObName: string,
    address: Address,
    tariffGroup: string,
    contractedPower: ContractedPower,
}
export interface ActiveEnergyConsumed {
    energySell: EnergyReading,
    energyDistribution: EnergyReading
}
export interface ActiveEnergyProduced {
    meters: MeterProduction[],
    summary: MeterProductionSummary
}
export interface PpeDeposit {
    deposit: PpeDepositItem[],
    ppeSummaryTotal: PpeDepositSummary
}
export interface ExcessSalesBalance {
    itemName: string,
    itemCode: string,
    grossVal: number
}
export interface SellSummary {
    items: SellSummaryItem[],
    total: Total
}
export interface PaymentSummary {
    items: SellSummaryItem[],
    total: Total
} 
export interface EnergyAnnualBalance{
    history: DepositHistory[]
}
export interface PpeSummary {
    items: PpeSummaryItem[],
    total: PpeSummaryTotal
}
export interface BankDetails {
    account: string
}
export interface ContractedPower{
    value: number,
    unit: string
}
export interface EnergyReading {
    meters: Meter[],
    exciseTax: number,
    subtotal: ConsumptionSubtotal
}
export interface MeterProduction {
    meterNumber: string,
    items: MeterProductionItem[]
}
export interface MeterProductionSummary {
    production: number,
    netVal: number,
    taxVal: number,
    grossVal: number
}
export interface PpeDepositItem {
    ppeCode: number,
    ppeSummary: PpeDepositSummary
}
export interface PpeDepositSummary {
    depositCurrent: number, 
    depositConsumed: number, 
    depositNext: number
}
export interface SellSummaryItem {
    itemName: string,
    itemCode: string, 
    vatRate: number, 
    netVal: number,
    taxVal: number, 
    grossVal: number
}
export interface Total {
    netVal: number,
    taxVal: number,
    grossVal: number
}
export interface DepositHistory {
    ppeCode: string,
    items: EnergyAnnualBalanceItem[]
}
export interface PpeSummaryItem {
    ppeCode: string,
    value: number, 
    energyConsumed: number, 
    energyProduced: number
}
export interface PpeSummaryTotal {
    value: number, 
    energyConsumed: number, 
    energyProduced: number
}
export interface Meter {
    meterNumber: string,
    items: Item[]
}
export interface ConsumptionSubtotal {
    amount: number, 
    netValue: number 
}
export interface MeterProductionItem {
    itemName: string, 
    itemCode: string, 
    dateFrom: string,
    dateTo: string,
    production: number, 
    vatRate: number,
    netVal: number,
    taxVal: number,
    grossVal: number
}
export interface EnergyAnnualBalanceItem {
    itemName: string, 
    itemCode: string, 
    periods: number[]
}
export interface Item {
    itemName: string,
    itemCode: string, 
    prevMeterRead: MeterReading,
    currMeterRead: MeterReading,
    factor: number,
    consumption: number,
    netUnitPrice: number,
    netVal: number, 
    vatRate: number
}
export interface MeterReading {
    dt: string, 
    value: number, 
    readType: string
}