import { defineStore } from 'pinia';
import {DataHolder} from "@/models/data-holder";
import {Customer} from "@/models/customer";
import factoryApi from "@/api/factory.api";
import {LocalSpinner} from "@/services/model/localSpinner";
import { Customer as CustomerModel} from '@/components/forms/create-customer/Customer';
import { PagingModel } from '@/services/model/paging.model';

export const useCustomerStore = defineStore({
  id: 'customersStore',
  state: () => {
    return {
      customers: Object(new DataHolder<Customer>()),
      currentCustomerData: {} as Customer,
      selectedCustomerData: {} as Customer
    };
  },
  actions: {
    setCustomers (customers:  DataHolder<Customer>) {
      this.customers = customers;
      console.log(this.customers);
    },
    setCurrentCustomerData(customer : Customer){
      this.currentCustomerData = customer;
    },
    setSelectedCustomerData(customer : Customer){
      this.selectedCustomerData = customer;
    },
    fetchCustomersData(lockScreen: boolean, localSpinner: LocalSpinner | null, pagination: PagingModel){
      factoryApi.customerApi().getCustomersData(lockScreen, localSpinner, pagination).then(
        (response) => {
          if (response.error === null) {
            this.customers = response.data;
          } else {
            //TODO: error handling
          }
        }
      );
    },
    fetchCurrentCustomerData(lockScreen: boolean, localSpinner: LocalSpinner | null){
      factoryApi.customerApi().getCurrentCustomerData(lockScreen, localSpinner).then(
        (response)=>{
          if(response.error === null) {
            if(response.data !== undefined)
            this.currentCustomerData = response.data;
          } else {
            //TODO: error handling
          }
        }
      );
    },
    fetchSelectedCustomerData(customerId: number, lockScreen: boolean, localSpinner: LocalSpinner | null) {
      factoryApi.customerApi().getCustomerDataById(customerId, lockScreen, localSpinner).then(
        (response) => {
          if (response.error === null) {
            if (response.data !== undefined)
              this.selectedCustomerData = response.data;
          } else {
            //TODO: error handling
          }
        }
      );
    },
    createCustomer(customer: CustomerModel, successCallback: () => void , failCallback: () => void) {
      factoryApi.customerApi().saveNewCustomer(customer, successCallback, failCallback);
    },
    updateCustomerById(customerId: number, customer: Customer, successCallback: () => void , failCallback: () => void) {
      factoryApi.customerApi().updateCustomerById(customerId, customer, successCallback, failCallback);
    },
    deleteCustomer(customerId: number, successCallback: () => void , failCallback: () => void) {
      factoryApi.customerApi().deleteCustomerById(customerId, successCallback, failCallback);
    }
  },
  getters: {
    listDataHolder ():DataHolder<Customer> {
      return this.customers;
    },
  }
});
