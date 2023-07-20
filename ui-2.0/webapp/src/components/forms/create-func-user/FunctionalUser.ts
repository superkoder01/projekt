export interface FunctionalUser{
    login: string;
    password: string;
    providerId: number;
    roleId: number;
    id: number;
    workerId: number;

    firstName: string;
    lastName: string;
    phone: string;
    blockchainAccAddress: string;
    customerTypeName: string;
    pesel: string;
    email: string;

    nip: string;
    krs: string;
    workStartDate: Date;
    workEndDate: Date;

    street: string;
    buildingNumber: string;
    apartmentNumber: string;
    city: string;
    postalCode: string;
    province: string;
    country: string;

    // TODO:
    supervisor: number
}
