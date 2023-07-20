import { TariffGroupApi } from './tarrifGroup.api';
import {WorkersApi} from './workerks.api';
import {UserApi} from "@/api/user.api";
import {CustomerApi} from "@/api/customer.api";
import {UsersApi} from "@/api/users.api";
import {ProvidersApi} from "@/api/providers.api";
import {OffersApi} from "@/api/offers.api";
import {InvoicesApi} from "@/api/invoices.api";
import {ContractsApi} from "@/api/contracts.api";
import {DraftOffersApi} from "@/api/draftOffers.api";
import {InstallationsApi} from '@/api/installations.api';
import {PricingApi} from "@/api/pricing.api";
import {PdfConverterApi} from "@/api/pdfConverter.api";

class ApiFactory {

  private static instance: ApiFactory;

  private readonly _userApi = new UserApi();
  private readonly _usersApi = new UsersApi();
  private readonly _customerApi = new CustomerApi();
  private readonly _providersApi = new ProvidersApi();
  private readonly _offersApi = new OffersApi();
  private readonly _invoicesApi= new InvoicesApi();
  private readonly _contractsApi= new ContractsApi();
  private readonly _workersApi = new WorkersApi();
  private readonly _draftOffersApi = new DraftOffersApi();
  private readonly _installationsApi = new InstallationsApi();
  private readonly _pricingApi = new PricingApi();
  private readonly _pdfConverterApi = new PdfConverterApi();
  private readonly _tarrifGroupApi = new TariffGroupApi();

  private constructor() {
 //
  }

  public static getInstance(): ApiFactory {
    if (!ApiFactory.instance) {
      ApiFactory.instance = new ApiFactory();
    }
    return ApiFactory.instance;
  }

  public userApi(): UserApi{
    return this._userApi;
  }

  public usersApi(): UsersApi{
    return this._usersApi;
}

  public customerApi(): CustomerApi{
    return this._customerApi;
  }

  public providersApi(): ProvidersApi{
    return this._providersApi;
  }

  public invoicesApi(): InvoicesApi{
    return this._invoicesApi;
  }

  public contractsApi(): ContractsApi{
    return this._contractsApi;
  }

  public offersApi(): OffersApi {
    return this._offersApi;
  }
  public workersApi(): WorkersApi {
    return this._workersApi;
  }
  public draftOffersApi(): DraftOffersApi {
    return this._draftOffersApi;
  }
  public installationsApi() : InstallationsApi {
    return this._installationsApi;
  }
  public pricingApi(): PricingApi {
    return this._pricingApi;
  }
  public pdfConverterApi(): PdfConverterApi {
    return this._pdfConverterApi;
  }
  public tarrifGroupApi(): TariffGroupApi {
    return this._tarrifGroupApi;
  }
  
}

export default
ApiFactory.getInstance();

