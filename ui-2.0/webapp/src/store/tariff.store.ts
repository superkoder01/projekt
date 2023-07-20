import {defineStore} from "pinia";
import {DataHolder} from "@/models/data-holder";
import {TariffGroup} from "@/components/forms/tariff-group/TariffGroup";

export const useTariffStore = defineStore({
  id: 'tariffStore',
  state: () => {
    return {
      tariff : Object(new DataHolder<TariffGroup>())
    };
  },
  actions: {
    setTariff (tariff:  DataHolder<TariffGroup>) {
      this.tariff = tariff;
    }
  },
  getters: {
    getListDataHolder ():DataHolder<TariffGroup> {
      return this.tariff;
    }
  }
});
