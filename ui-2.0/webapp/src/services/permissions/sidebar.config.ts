import {PagesEnum} from "@/services/permissions/pages-enum";

export class SidebarConfig{

  config = new Map<PagesEnum,SidebarElement>() ;

  constructor() {
    this.config.set(PagesEnum.USERS, this.createUsers());
    this.config.set(PagesEnum.PROVIDERS, this.createProviders());
    this.config.set(PagesEnum.CONTRACTS, this.createContracts());
    this.config.set(PagesEnum.INVOICES, this.createInvoices());
    this.config.set(PagesEnum.OFFERS, this.createOffers());
    this.config.set(PagesEnum.DOCUMENTS, this.createDocuments());
    this.config.set(PagesEnum.HOME, this.createHome());
    this.config.set(PagesEnum.EMPLOYEES, this.createEmployees());
    this.config.set(PagesEnum.CUSTOMERS, this.createCustomers());
    this.config.set(PagesEnum.PRODUCT_CATALOG, this.createProductCatalog());
    this.config.set(PagesEnum.SUPER_ADMINS, this.createSuperAdmins());
    this.config.set(PagesEnum.CONFIGURATION, this.createConfiguration());
    this.config.set(PagesEnum.HISTORY, this.createHistory());
    this.config.set(PagesEnum.INSTALATIONS, this.createInstallation());
    this.config.set(PagesEnum.SETTINGS, this.createSettings());
  }

  getConfigForPage(page: PagesEnum): SidebarElement | undefined{
    return this.config.get(page);
  }

  private createProviders(): SidebarElement{
    const retVal = new SidebarElement();
    retVal.href = '/providers';
    retVal.title = "Providers";
    // retVal.title = this.i18n.t('SIDEBAR_MENU.TITLE');
    retVal.icon = new SidebarIcon('Ghost');
    return retVal;
  }

  private createUsers(): SidebarElement{
    const retVal = new SidebarElement();
    retVal.href = '/users';
    retVal.title = 'Users';
    retVal.icon = new SidebarIcon('Contact');
    return retVal;
  }
  private createContracts(): SidebarElement{
    const retVal = new SidebarElement();
    retVal.href = '/contracts';
    retVal.title = 'Contracts';
    retVal.icon = new SidebarIcon('Album');
    return retVal;
  }
  private createInvoices(): SidebarElement{
    const retVal = new SidebarElement();
    retVal.href = '/invoices';
    retVal.title = 'Invoices';
    retVal.icon = new SidebarIcon('Coins');
    return retVal;
  }
  private createOffers(): SidebarElement{
    const retVal = new SidebarElement();
    retVal.href = '/offers';
    retVal.title = 'Offers';
    retVal.icon = new SidebarIcon('Gem');
    return retVal;
  }
  private createDocuments(): SidebarElement{
    const retVal = new SidebarElement();
    retVal.href = '/documents';
    retVal.title = 'Documents';
    retVal.icon = new SidebarIcon('FileText');
    return retVal;
  }

  private createHome(): SidebarElement{
    const retVal = new SidebarElement();
    retVal.href = '/home';
    retVal.title = 'Home';
    retVal.icon = new SidebarIcon('Home');
    return retVal;
  }

  private createCustomers(): SidebarElement{
    const retVal = new SidebarElement();
    retVal.href = '/customers';
    retVal.title = 'Customers';
    retVal.icon = new SidebarIcon('Users');
    return retVal;
  }

  private createEmployees(): SidebarElement{
    const retVal = new SidebarElement();
    retVal.href = '/employees';
    retVal.title = 'Employees';
    retVal.icon = new SidebarIcon('Zap');
    return retVal;
  }
  private createProductCatalog(): SidebarElement{
    const retVal = new SidebarElement();
    retVal.href = '/product_catalog';
    retVal.title = 'ProductCatalog';
    retVal.icon = new SidebarIcon('BookOpen');
    return retVal;
  }
  private createSuperAdmins(): SidebarElement {
    const retVal = new SidebarElement();
    retVal.href = '/super_admins';
    retVal.title = "SuperAdmins";
    retVal.icon = new SidebarIcon('ShieldCheck');
    return retVal;
  }
  private createConfiguration(): SidebarElement {
    const retVal = new SidebarElement();
    retVal.href = '/configuration';
    retVal.title = "Configuration";
    retVal.icon = new SidebarIcon('Wrench');
    return retVal;
  }
  private createHistory(): SidebarElement {
    const retVal = new SidebarElement();
    retVal.href = '/history';
    retVal.title = "History";
    retVal.icon = new SidebarIcon('History');
    return retVal;
  }
  private createInstallation(): SidebarElement {
    const retVal = new SidebarElement();
    retVal.href = '/installations';
    retVal.title = 'Installations';
    retVal.icon = new SidebarIcon('Sun');
    return retVal;
  }
  private createSettings(): SidebarElement {
    const retVal = new SidebarElement();
    retVal.href = '/settings';
    retVal.title = 'Settings';
    retVal.icon = new SidebarIcon('Settings');
    return retVal;
  }
}

export class SidebarElement {
  href = "" ;
  title ="";
  icon = new SidebarIcon('');

}

export class SidebarIcon {
  element: string;
  constructor(element: string) {
    this.element = element;
  }
}
