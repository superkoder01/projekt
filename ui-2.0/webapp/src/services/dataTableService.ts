import {DataHolder} from "@/models/data-holder";
import {PagingModel} from "@/services/model/paging.model";
import {LocalSpinner} from "@/services/model/localSpinner";
import {DefaultSortingModel} from "@/services/model/defaultSorting.model";

export interface DataTableService<T> {

  getDefaultSorting(): DefaultSortingModel;

  fetchListData(pagination: PagingModel, lockScreen: boolean, localSpinner: LocalSpinner | null) : void;

  getListDataHolder() : DataHolder<T>;
}

