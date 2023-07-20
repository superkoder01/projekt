export enum ServiceTypeEnum {
  NONE= " none",
  SALE="sale",
  SALE_RDN="sale_rdn",
  REPURCHASE="fixed",
  REPURCHASE_RDN="rdn"

}

export function toServiceTypeEnumKey(value: ServiceTypeEnum): string{
  return Object.keys(ServiceTypeEnum)[Object.values(ServiceTypeEnum).indexOf(value)];
}
