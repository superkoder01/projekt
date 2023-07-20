import {createRouter, createWebHistory, NavigationGuardNext, RouteLocationNormalized, RouteRecordRaw} from 'vue-router';
import HomeView from '../views/HomeView.vue';
import LoginView from '../views/LoginView.vue';
import ContractsView from '../views/ContractsListView.vue';
import InvoicesView from '../views/InvoicesListView.vue';
import OffersView from '../views/OffersListView.vue';
import DocumentsView from '../views/DocumentsView.vue';
import InstallationsView from '../views/InstallationsListView.vue';
import UserManagementView from '../views/UserManagementView.vue';
import EmployeeDetailsView from '../views/EmployeeDetailsView.vue';
import PricingView from '../views/PricingListView.vue';
import TariffView from '../views/TariffListView.vue';
import SuperAdminsListView from '@/views/SuperAdminsListView.vue';
import NewFunctionalUserVue from '@/components/forms/create-func-user/NewFunctionalUser.vue';
import FunctionalUserPersonalStep1Vue from '@/components/forms/create-func-user/FunctionalUserPersonalStep1.vue';
import FunctionalUserAddressStep2Vue from '@/components/forms/create-func-user/FunctionalUserAddressStep2.vue';
import ConfigurationView from '@/views/ConfigurationView.vue';
import providersRoutes from "@/router/providers.route";
import AccountActivateVue from '@/views/AccountActivate.vue';
import productCatalogRoutes from "@/router/productCatalog.route";
import {PagesEnum} from "@/services/permissions/pages-enum";
import FAQViewVue from '@/views/FAQView.vue';
import HelpView from '@/views/HelpView.vue';
import TermsConditionsView from '@/views/TermsConditionsView.vue';
import PrivacyPolicyView from '@/views/PrivacyPolicyView.vue';
import ResetPasswordVue from '@/components/forms/reset-password/ResetPassword.vue';
import SettingsTabsView from '@/views/SettingsTabsView.vue';
import customersRoute from "@/router/customers.route";
import EmptyRouting from "@/views/EmptyRouting.vue";
import {RoleEnum} from "@/services/permissions/role-enum";
import AdminDetailsView from '@/views/AdminDetailsView.vue';
import {useUserStore} from "@/store/user.store";
// import { useUserStore } from '@/store/user.store';
// import { inject } from 'vue';
// import { LoggerService } from '@/services/logger/logger.service';
// import { LogLevel } from '@/services/logger/log-level';

const routes: Array<RouteRecordRaw> = [
  {
    path: '/',
    component: HomeView,
    meta: {
      requiresAuth: true,
      pageType: PagesEnum.CUSTOMERS
    }
  },
  {
    path: '/home',
    name: 'home',
    component: HomeView,
    meta: {
      requiresAuth: true,
      pageType: PagesEnum.HOME
    }
  },
  {
    path: '/login',
    name: 'login',
    component: LoginView,
    meta: {
      requiresAuth: false,
    }
  },
  {
    path: '/register/:activateCode',
    name: 'register',
    component: AccountActivateVue,
  },
  {
    path: '/reset_password',
    name: 'reset_password',
    component: ResetPasswordVue,
    meta: {
      requiresAuth: true,
    }
  },
  {
    path: '/contracts',
    name: 'contracts',
    component: ContractsView,
    meta: {
      requiresAuth: true,
      pageType: PagesEnum.CONTRACTS
    }
  },
  providersRoutes,
  {
    path: '/invoices',
    name: 'invoices',
    component: InvoicesView,
    meta: {
      requiresAuth: true,
      pageType: PagesEnum.INVOICES
    }
  },
  {
    path: '/offers',
    name: 'offers',
    component: OffersView,
    meta: {
      requiresAuth: true,
      pageType: PagesEnum.OFFERS
    }
  },
  {
    path: '/documents',
    name: 'documents',
    component: DocumentsView,
    meta: {
      requiresAuth: true,
      pageType: PagesEnum.DOCUMENTS
    }
  },
  {
    path: '/installations',
    name: 'installations',
    component: InstallationsView,
    meta: {
      requiresAuth: true,
    }
  },
  {
    path: '/settings',
    name: 'settings',
    component: SettingsTabsView,
    meta: {
      requiresAuth: true,
    }
  },
  {
    path: '/employees',
    name: 'employees',
    component: UserManagementView,
    meta: {
      requiresAuth: true,
      pageType: PagesEnum.EMPLOYEES
    },
    children: [
      {
        path: 'new_trader',
        name: 'new_trader',
        component: NewFunctionalUserVue,
        props: {passedRoleID: RoleEnum.TRADER, previousPath: '/employees'},
        meta: {
          requiresAuth: true,
          pageType: PagesEnum.EMPLOYEES
        },
        children: [
          {
            path: '',
            name: 'trader_personal_step',
            component: FunctionalUserPersonalStep1Vue,
            props: true,
            meta: {
              requiresAuth: true,
              pageType: PagesEnum.EMPLOYEES
            }
          },
          {
            path: 'address_2',
            name: 'trader_address',
            component: FunctionalUserAddressStep2Vue,
            meta: {
              requiresAuth: true,
              pageType: PagesEnum.EMPLOYEES
            }
          },
        ]
      },
      {
        path: 'new_super_agent',
        name: 'new_super_agent',
        component: NewFunctionalUserVue,
        props: {passedRoleID: RoleEnum.SUPER_AGENT, previousPath: '/employees'},
        meta: {
          requiresAuth: true,
          pageType: PagesEnum.EMPLOYEES
        },
        children: [
          {
            path: '',
            name: 'super_agent_personal_step',
            component: FunctionalUserPersonalStep1Vue,
            props: true,
            meta: {
              requiresAuth: true,
              pageType: PagesEnum.EMPLOYEES
            }
          },
          {
            path: 'address_2',
            name: 'super_agent_address',
            component: FunctionalUserAddressStep2Vue,
            meta: {
              requiresAuth: true,
              pageType: PagesEnum.EMPLOYEES
            }
          },
        ]
      },
      {
        path: 'new_agent',
        name: 'new_agent',
        component: NewFunctionalUserVue,
        props: {passedRoleID: RoleEnum.AGENT, previousPath: '/employees'},
        meta: {
          requiresAuth: true,
          pageType: PagesEnum.EMPLOYEES
        },
        children: [
          {
            path: '',
            name: 'agent_personal_step',
            component: FunctionalUserPersonalStep1Vue,
            props: true,
            meta: {
              requiresAuth: true,
              pageType: PagesEnum.EMPLOYEES

            }
          },
          {
            path: 'address_2',
            name: 'agent_address',
            component: FunctionalUserAddressStep2Vue,
            meta: {
              requiresAuth: true,
              pageType: PagesEnum.EMPLOYEES
            }
          },
        ]
      },
      {
        path: 'details',
        name: 'employee_details_view_router',
        props: true,
        component: EmptyRouting,
        meta: {
          requiresAuth: true,
          pageType: PagesEnum.EMPLOYEES
        },
        children: [
          {
            path: '',
            name: 'employee_details_view',
            props: true,
            component: EmployeeDetailsView,
            meta: {
              requiresAuth: true,
              pageType: PagesEnum.EMPLOYEES
            }
          },
          {
            path: 'edit_trader',
            name: 'edit_trader',
            component: NewFunctionalUserVue,
            props: {passedRoleID: RoleEnum.TRADER, previousPath: '/employees', useSelectedEmployee: true},
            meta: {
              requiresAuth: true,
              pageType: PagesEnum.EMPLOYEES
            },
            children: [
              {
                path: '',
                name: 'edit_trader_personal_step',
                component: FunctionalUserPersonalStep1Vue,
                props: true,
                meta: {
                  requiresAuth: true,
                  pageType: PagesEnum.EMPLOYEES
                }
              },
              {
                path: 'address_2',
                name: 'edit_trader_address',
                props: true,
                component: FunctionalUserAddressStep2Vue,
                meta: {
                  requiresAuth: true,
                  pageType: PagesEnum.EMPLOYEES
                }
              },
            ]
          },
          {
            path: 'edit_agent',
            name: 'edit_agent',
            component: NewFunctionalUserVue,
            props: {passedRoleID: RoleEnum.AGENT, previousPath: '/employees/details', useSelectedEmployee: true},
            meta: {
              requiresAuth: true,
              pageType: PagesEnum.EMPLOYEES
            },
            children: [
              {
                path: '',
                name: 'edit_agent_personal_step',
                component: FunctionalUserPersonalStep1Vue,
                props: true,
                meta: {
                  requiresAuth: true,
                  pageType: PagesEnum.EMPLOYEES
                }
              },
              {
                path: 'address_2',
                name: 'edit_agent_address',
                props: true,
                component: FunctionalUserAddressStep2Vue,
                meta: {
                  requiresAuth: true,
                  pageType: PagesEnum.EMPLOYEES
                }
              },
            ]
          },
          {
            path: 'edit_super_agent',
            name: 'edit_super_agent',
            component: NewFunctionalUserVue,
            props: {passedRoleID: RoleEnum.SUPER_AGENT, previousPath: '/employees/details', useSelectedEmployee: true},
            meta: {
              requiresAuth: true,
              pageType: PagesEnum.EMPLOYEES
            },
            children: [
              {
                path: '',
                name: 'edit_super_agent_personal_step',
                component: FunctionalUserPersonalStep1Vue,
                props: true,
                meta: {
                  requiresAuth: true,
                  pageType: PagesEnum.EMPLOYEES
                }
              },
              {
                path: 'address_2',
                name: 'edit_super_agent_address',
                props: true,
                component: FunctionalUserAddressStep2Vue,
                meta: {
                  requiresAuth: true,
                  pageType: PagesEnum.EMPLOYEES
                }
              },
            ]
          },
          {
            path: 'edit_admin_full',
            name: 'edit_admin_full',
            component: NewFunctionalUserVue,
            props: {passedRoleID: RoleEnum.ADMINISTRATOR_FULL, previousPath: '/employees/details', useSelectedEmployee: true},
            meta: {
              requiresAuth: true,
              pageType: PagesEnum.EMPLOYEES
            },
            children: [
              {
                path: '',
                name: 'edit_admin_full_personal_step',
                component: FunctionalUserPersonalStep1Vue,
                props: true,
                meta: {
                  requiresAuth: true,
                  pageType: PagesEnum.EMPLOYEES
                }
              },
              {
                path: 'address_2',
                name: 'edit_admin_full_address',
                props: true,
                component: FunctionalUserAddressStep2Vue,
                meta: {
                  requiresAuth: true,
                  pageType: PagesEnum.EMPLOYEES
                }
              },
            ]
          },
        ]
      }

    ]
  },

  {
    path: '/super_admins',
    name: 'super_admins',
    component: EmptyRouting,
    meta: {
      requiresAuth: true,
      pageType: PagesEnum.SUPER_ADMINS
    },
    children: [
      {
        path: '',
        name: 'super_admins',
        component: SuperAdminsListView,
        meta: {
          requiresAuth: true,
          pageType: PagesEnum.SUPER_ADMINS
        },
      },
      {
        path: 'details',
        name: 'super_admin_details',
        props: true,
        component: AdminDetailsView,
        meta: {
          requiresAuth: true,
          pageType: PagesEnum.SUPER_ADMINS
        },
      },
      {// add new super admin steps
        path: 'new_super_admin',
        name: 'new_super_admin',
        component: NewFunctionalUserVue,
        props: {passedRoleID: RoleEnum.SUPER_ADMIN, previousPath: '/providers/provider_tabs/admins'},
        meta: {
          requiresAuth: true,
          checkRequiredProps: {
            requiredProps: 'passedRoleID',
            redirectTo: '/super_admins'
          },
          pageType: PagesEnum.SUPER_ADMINS
        },
        // beforeEnter:(to, from, next) => checkIfPropsIsSetOrRedirect(to, 'passedRoleID', '/super_admins', next),
        children: [
          {
            path: '',
            name: 'super_admin_personal_step',
            component: FunctionalUserPersonalStep1Vue,
            props: true,
            meta: {
              requiresAuth: true,
              pageType: PagesEnum.SUPER_ADMINS
            }
          },
        ]
      },
    ]
  },

  customersRoute,
  {
    path: '/pricing',
    name: 'pricing',
    component: PricingView,
    meta: {
      requiresAuth: true
    }
  },
  {
    path: '/tariff',
    name: 'tariff',
    component: TariffView,
    meta: {
      requiresAuth: true
    }
  },

  productCatalogRoutes,
/*  {
    path: '/product_catalog/tariff_pricing/new_tariff_group',
    name: 'new_tariff_group',
    component: newTariffGroup,
    meta: {
      requiresAuth: true,
      pageType: PagesEnum.PRODUCT_CATALOG
    },
    children: [
      {
        path: '',
        name: 'general_step_1',
        component: GeneralStep1,
        meta: {
          requiresAuth: true,
          pageType: PagesEnum.PRODUCT_CATALOG
        }
      },
      {
        path: 'fees_2',
        name: 'fees',
        component: FeesStep2,
        meta: {
          requiresAuth: true,
          pageType: PagesEnum.PRODUCT_CATALOG
        }
      },
    ]
  },*/


  // {
  //   path: '/product_catalog/offer_draft/new_offer_draft',
  //   name: 'new_offer_draft',
  //   component: NewOfferDraft,
  //   meta: {
  //     // requiresAuth: true
  //   },
  //   children:[
  //     {
  //       path: '',
  //       name: 'draft_offer_step1',
  //       component: DraftOfferStep1,
  //       meta: {
  //         // requiresAuth: true
  //       }
  //     },
  //     {
  //       path: 'offer_draft_2',
  //       name: 'draft_offer_step2',
  //       component: DraftOfferStep2,
  //       meta: {
  //         // requiresAuth: true
  //       }
  //     },
  //     {
  //       path: 'offer_draft_3',
  //       name: 'draft_offer_step3',
  //       component: DraftOfferStep3,
  //       meta: {
  //         requiresAuth: true
  //       }
  //     }
  //   ]
  // },

  {
    path: '/configuration',
    name: 'configuration',
    component: ConfigurationView,
    meta: {
      requiresAuth: true,
      pageType: PagesEnum.CONFIGURATION
    }
  },
  {
    path: '/history',
    name: 'history',
    component: () => import(/* webpackChunkName: "history" */ '../views/HistoryView.vue'),
    meta: {
      requiresAuth: true,
      pageType: PagesEnum.HISTORY
    }
  },
  {
    path: '/about',
    name: 'about',
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () => import(/* webpackChunkName: "about" */ '../views/AboutView.vue'),
    meta: {
      requiresAuth: true,
    }
  },
  {
    path: '/faq',
    name: 'faq',
    component: FAQViewVue,
    meta: {
      requiresAuth: true,
    }
  },
  {
    path: '/help',
    name: 'help',
    component: HelpView,
    meta: {
      requiresAuth: true,
    }
  },
  {
    path: '/terms_conditions',
    name: 'terms_conditions',
    component: TermsConditionsView,
    meta: {
      requiresAuth: true,
    }
  },
  {
    path: '/privacy_policy',
    name: 'privacy_policy',
    component: PrivacyPolicyView,
    meta: {
      requiresAuth: true,
    }
  },
];

const router = createRouter({
  scrollBehavior(to, from, savedPosition) {
    if (savedPosition) {
      return savedPosition
    } else {
      return { top: 0 }
    }
  },
  history: createWebHistory(process.env.BASE_URL),
  routes
});

function checkIfPropsIsSetOrRedirect(to: RouteLocationNormalized, requiredProp: string, redirectTo: string, next: NavigationGuardNext) {
  if (to.params['passedRoleID'] !== undefined) {
    next();
  } else {
    next(redirectTo);
  }
}

// const logger = inject<LoggerService>('logger') as LoggerService;

// router.beforeEach((to, from, next) => {
//   logger.logToConsole(LogLevel.DEBUG, 'Router', 'go form:' + JSON.stringify(from.name) + ' to:' + JSON.stringify(to.name));
//   if (to.matched.some(record => record.meta.requiresAuth)) {
//     if (useUserStore().isLoggedIn) {
//       next();
//       return;
//     }
//     logger.logToConsole(LogLevel.INFO, 'Router', 'user not logged in, redirecting to login page');
//     next('/login');
//   } else {
//     logger.logToConsole(LogLevel.WARNING, 'Router', 'requiresAuth not defined in router configuration');
//     next();
//   }
// });

export default router;
