import {PagesEnum} from "@/services/permissions/pages-enum";
import {useUserStore} from "@/store/user.store";
import {RoleEnum} from "@/services/permissions/role-enum";
import {SidebarConfig, SidebarElement} from "@/services/permissions/sidebar.config";
import {FeatureEnum} from "./feature-enum";

export class PermissionsService {

  canAccessFeature(feature: FeatureEnum): boolean {

    switch (feature) {
      case FeatureEnum.MANAGE_SUPER_ADMIN:
        return this.canDo([RoleEnum.SUPER_ADMIN]);
      case FeatureEnum.MANAGE_FULL_ADMIN:
        return this.canDo([RoleEnum.SUPER_ADMIN]);
      case FeatureEnum.MANAGE_ADMIN:
        return this.canDo([RoleEnum.ADMINISTRATOR_FULL]);
      case FeatureEnum.MANAGE_TRADER:
        return this.canDo([RoleEnum.ADMINISTRATOR_FULL, RoleEnum.ADMINISTRATOR_BASIC]);
      case FeatureEnum.MANAGE_SUPER_AGENT:
        return this.canDo([RoleEnum.ADMINISTRATOR_FULL, RoleEnum.ADMINISTRATOR_BASIC, RoleEnum.TRADER]);
      case FeatureEnum.MANAGE_AGENT:
        return this.canDo([RoleEnum.ADMINISTRATOR_FULL, RoleEnum.ADMINISTRATOR_BASIC, RoleEnum.TRADER, RoleEnum.SUPER_AGENT]);
      case FeatureEnum.MANAGE_PROSUMER:
        return this.canDo([RoleEnum.ADMINISTRATOR_FULL, RoleEnum.ADMINISTRATOR_BASIC, RoleEnum.TRADER, RoleEnum.SUPER_AGENT, RoleEnum.AGENT]);
      case FeatureEnum.RESEND_ACTIVATION_LINK:
        return this.canDo([RoleEnum.SUPER_ADMIN, RoleEnum.ADMINISTRATOR_FULL, RoleEnum.ADMINISTRATOR_BASIC, RoleEnum.TRADER]);

      case FeatureEnum.MANAGE_PRICING:
        return this.canDo([RoleEnum.SUPER_ADMIN, RoleEnum.ADMINISTRATOR_FULL, RoleEnum.ADMINISTRATOR_BASIC]);

      case FeatureEnum.MANAGE_OFFER_DRAFTS:
        return this.canDo([RoleEnum.SUPER_ADMIN, RoleEnum.ADMINISTRATOR_FULL, RoleEnum.ADMINISTRATOR_BASIC]);

      case FeatureEnum.MANAGE_OFFERS:
        return this.canDo([RoleEnum.SUPER_ADMIN, RoleEnum.ADMINISTRATOR_FULL, RoleEnum.ADMINISTRATOR_BASIC, RoleEnum.SUPER_AGENT, RoleEnum.AGENT]);

      case FeatureEnum.MANAGE_PARTNERS:
        return this.canDo([RoleEnum.SUPER_ADMIN]);
    }
    return false;
  }


  canView(page: PagesEnum): boolean {
    switch (page) {
      case PagesEnum.PROVIDERS:
        return this.canDo([RoleEnum.SUPER_ADMIN]);
      case PagesEnum.USERS:
        return this.canDo([RoleEnum.SUPER_ADMIN]);
      case PagesEnum.CONTRACTS:
        return this.canDo([RoleEnum.ADMINISTRATOR_FULL, RoleEnum.ADMINISTRATOR_BASIC, RoleEnum.TRADER, RoleEnum.PROSUMER]);
      case PagesEnum.INVOICES:
        return this.canDo([RoleEnum.ADMINISTRATOR_FULL, RoleEnum.ADMINISTRATOR_BASIC, RoleEnum.TRADER, RoleEnum.PROSUMER]);
      case PagesEnum.OFFERS:
        return this.canDo([RoleEnum.ADMINISTRATOR_FULL, RoleEnum.ADMINISTRATOR_BASIC, RoleEnum.TRADER, RoleEnum.SUPER_AGENT, RoleEnum.AGENT]);
      case PagesEnum.HOME:
        return this.canDo([RoleEnum.ADMINISTRATOR_FULL, RoleEnum.ADMINISTRATOR_BASIC, RoleEnum.TRADER, RoleEnum.SUPER_AGENT, RoleEnum.AGENT, RoleEnum.PROSUMER]);
      case PagesEnum.DOCUMENTS:
        return this.canDo([RoleEnum.PROSUMER]);
      case PagesEnum.CUSTOMERS:
        return this.canDo([RoleEnum.ADMINISTRATOR_FULL, RoleEnum.ADMINISTRATOR_BASIC, RoleEnum.TRADER, RoleEnum.SUPER_AGENT, RoleEnum.AGENT]);
      case PagesEnum.EMPLOYEES:
        return this.canDo([RoleEnum.SUPER_ADMIN, RoleEnum.ADMINISTRATOR_FULL, RoleEnum.ADMINISTRATOR_BASIC, RoleEnum.TRADER, RoleEnum.SUPER_AGENT, RoleEnum.AGENT]);
      case PagesEnum.PRODUCT_CATALOG:
        return this.canDo([RoleEnum.SUPER_ADMIN, RoleEnum.ADMINISTRATOR_FULL, RoleEnum.TRADER, RoleEnum.AGENT, RoleEnum.SUPER_AGENT]);
      case PagesEnum.SUPER_ADMINS:
        return this.canDo([RoleEnum.SUPER_ADMIN]);
      case PagesEnum.CONFIGURATION:
        return this.canDo([RoleEnum.SUPER_ADMIN]);
      case PagesEnum.HISTORY:
        return this.canDo([RoleEnum.SUPER_ADMIN]);
      case PagesEnum.INSTALATIONS:
        return this.canDo([RoleEnum.ADMINISTRATOR_FULL, RoleEnum.ADMINISTRATOR_BASIC]);
      case PagesEnum.SETTINGS:
        return this.canDo([RoleEnum.ADMINISTRATOR_FULL]);
    }
    return false;
  }

  private canDo(roles: Array<RoleEnum>): boolean {
    return roles.includes(useUserStore().currentRole);
  }

  sidebarElementOrder(): PagesEnum[] {
    const currentRole = useUserStore().currentRole;
    switch (currentRole) {
      case RoleEnum.SUPER_ADMIN:
        return [PagesEnum.PROVIDERS, PagesEnum.SUPER_ADMINS, PagesEnum.CONFIGURATION, PagesEnum.HISTORY];
      case RoleEnum.ADMINISTRATOR_FULL:
        return [PagesEnum.HOME, PagesEnum.CUSTOMERS, PagesEnum.OFFERS, PagesEnum.USERS, PagesEnum.PROVIDERS, PagesEnum.CONTRACTS, PagesEnum.INSTALATIONS, PagesEnum.INVOICES, PagesEnum.DOCUMENTS, PagesEnum.EMPLOYEES, PagesEnum.PRODUCT_CATALOG, PagesEnum.SUPER_ADMINS, PagesEnum.SETTINGS];
      case RoleEnum.ADMINISTRATOR_BASIC:
        return [PagesEnum.HOME, PagesEnum.CUSTOMERS, PagesEnum.OFFERS, PagesEnum.USERS, PagesEnum.PROVIDERS, PagesEnum.CONTRACTS, PagesEnum.INSTALATIONS, PagesEnum.INVOICES, PagesEnum.DOCUMENTS, PagesEnum.EMPLOYEES, PagesEnum.PRODUCT_CATALOG, PagesEnum.SUPER_ADMINS];
      case RoleEnum.SUPER_AGENT:
        return [PagesEnum.HOME, PagesEnum.CUSTOMERS, PagesEnum.OFFERS, PagesEnum.USERS, PagesEnum.PROVIDERS, PagesEnum.CONTRACTS, PagesEnum.INVOICES, PagesEnum.DOCUMENTS, PagesEnum.EMPLOYEES, PagesEnum.PRODUCT_CATALOG];
      case RoleEnum.AGENT:
        return [PagesEnum.HOME, PagesEnum.CUSTOMERS, PagesEnum.OFFERS, PagesEnum.USERS, PagesEnum.PROVIDERS, PagesEnum.CONTRACTS, PagesEnum.INVOICES, PagesEnum.DOCUMENTS, PagesEnum.EMPLOYEES, PagesEnum.PRODUCT_CATALOG];
      case RoleEnum.TRADER:
        return [PagesEnum.HOME, PagesEnum.CUSTOMERS, PagesEnum.OFFERS, PagesEnum.USERS, PagesEnum.PROVIDERS, PagesEnum.CONTRACTS, PagesEnum.INVOICES, PagesEnum.DOCUMENTS, PagesEnum.EMPLOYEES, PagesEnum.PRODUCT_CATALOG];
      case RoleEnum.PROSUMER:
        return [PagesEnum.HOME, PagesEnum.INVOICES, PagesEnum.DOCUMENTS];
    }
    //TODO: remove
    return [PagesEnum.HOME, PagesEnum.USERS, PagesEnum.PROVIDERS, PagesEnum.INVOICES, PagesEnum.CONTRACTS, PagesEnum.OFFERS, PagesEnum.DOCUMENTS, PagesEnum.CUSTOMERS, PagesEnum.EMPLOYEES, PagesEnum.PRODUCT_CATALOG];
  }

  createSideBar(): SidebarElement[] {
    const sidebarConfig = new SidebarConfig();
    const rerVal = Array<SidebarElement>();
    if (useUserStore().isLoggedIn) {
      console.log("createSideBar");

      const pagesEnums = this.sidebarElementOrder();
      for (const page of pagesEnums) {
        if (this.canView(page)) {
          const configForPage = sidebarConfig.getConfigForPage(page);
          if (configForPage !== undefined) {
            rerVal.push(configForPage);
          }
        }
      }
    }
    console.log(rerVal);
    return rerVal;
  }

  getAfterLoginPage(): string {
    const sidebarConfig = new SidebarConfig();
    if (useUserStore().isLoggedIn) {
      console.log("createSideBar");

      const pagesEnums = this.sidebarElementOrder();
      for (const page of pagesEnums) {
        if (this.canView(page)) {
          const configForPage = sidebarConfig.getConfigForPage(page);
          if (configForPage !== undefined) {
            return configForPage.href;
          }
        }
      }
    }
    return "/home";
  }

}
