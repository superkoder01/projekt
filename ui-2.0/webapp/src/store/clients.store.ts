import {defineStore} from "pinia";
import {DataHolder} from "@/models/data-holder";
import {Customer} from "@/models/customer";

export const useClientsStore = defineStore({
  id: 'clientsStore',
  state: () => {
    return {
      customers: Object(new DataHolder<Customer>()),
      customer: {} as Customer
    };
  },
  actions: {
    setCustomers (customers:  DataHolder<Customer>) {
      this.customers = customers;
      console.log(this.customers);
    },
    setCurrentCustomerData(customer : Customer){
      this.customer = customer;
    }
  },
  getters: {
    listDataHolder ():DataHolder<Customer> {
      return this.customers;
    },
    singleData ():Customer{
      return this.customer;
    }
  }
});
