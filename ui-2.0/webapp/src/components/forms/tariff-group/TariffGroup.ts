export interface TariffGroup {
  distributionNetworkOperatorID: number;
  tariffGroupLabelName: string;
  name: string;
  startDate: Date;
  endDate: Date;
  fees: {nameId: number, price:number}[];
}
