import ProvidersListView from "@/views/ProvidersListView.vue";
import ProviderTabsView from "@/views/ProviderTabsView.vue";
import AdminDetailsView from "@/views/AdminDetailsView.vue";
import {PagesEnum} from "@/services/permissions/pages-enum";
import {BreadcrumbElementType} from "@/router/breadcrumbElementType";
import AdministratorsListView from '@/views/AdministratorsListView.vue';
import ProviderDetailsView from '@/views/ProviderDetailsView.vue';
import EmptyRouting from "@/views/EmptyRouting.vue";
import NewProviderVue from "@/components/forms/create-provider/NewProvider.vue";
import ProviderDataStep1Vue from "@/components/forms/create-provider/ProviderDataStep1.vue";
import ProviderAddresStep2Vue from "@/components/forms/create-provider/ProviderAddresStep2.vue";
import ProviderAdminStep3Vue from "@/components/forms/create-provider/ProviderAdminStep3.vue";
import NewFunctionalUserVue from "@/components/forms/create-func-user/NewFunctionalUser.vue";
import FunctionalUserPersonalStep1Vue from "@/components/forms/create-func-user/FunctionalUserPersonalStep1.vue";
import {RoleEnum} from "@/services/permissions/role-enum";

const providersRoutes =
  {
    path: '/providers',
    name: 'providers',
    component: EmptyRouting,
    meta: {
      requiresAuth: true,
      pageType: PagesEnum.PROVIDERS
    },
    children: [
      {
        path: '',
        name: 'providers',
        component: ProvidersListView,
        meta: {
          requiresAuth: true,
          pageType: PagesEnum.PROVIDERS
        }
      },
      {
        path: 'new_provider',
        name: 'new_provider',
        component: NewProviderVue,
        meta: {
          requiresAuth: true,
          pageType: PagesEnum.PROVIDERS
        },
        children: [
          {
            path: '',
            name: 'provider_data_step',
            component: ProviderDataStep1Vue,
            meta: {
              requiresAuth: true,
              pageType: PagesEnum.PROVIDERS
            }
          },
          {
            path: 'address_2',
            name: 'provider_addres',
            component: ProviderAddresStep2Vue,
            meta: {
              requiresAuth: true,
              pageType: PagesEnum.PROVIDERS
            }
          },
          {
            path: 'admin_3',
            name: 'provider_admin',
            component: ProviderAdminStep3Vue,
            meta: {
              requiresAuth: true,
              pageType: PagesEnum.PROVIDERS
            }
          }
        ]
      },
      {
        path: 'provider_tabs',
        name: 'provider_tabs',
        component: ProviderTabsView,
        meta: {
          requiresAuth: true,
          pageType: PagesEnum.PROVIDERS,
          breadcrumbElement: BreadcrumbElementType.SELECTED_PROVIDE_NAME
        },
        children: [
          {
            path: 'edit_provider',
            name: 'edit_provider',
            component: NewProviderVue,
            props: { useSelectedProvider: true },
            meta: {
              requiresAuth: true,
              pageType: PagesEnum.PROVIDERS
            },
            children: [
              {
                path: '',
                name: 'edit_provider_data_step',
                component: ProviderDataStep1Vue,
                meta: {
                  requiresAuth: true,
                  pageType: PagesEnum.PROVIDERS
                }
              },
              {
                path: 'address_2',
                name: 'edit_provider_addres',
                component: ProviderAddresStep2Vue,
                meta: {
                  requiresAuth: true,
                  pageType: PagesEnum.PROVIDERS
                }
              },
              {
                path: 'admin_3',
                name: 'edit_provider_admin',
                component: ProviderAdminStep3Vue,
                meta: {
                  requiresAuth: true,
                  pageType: PagesEnum.PROVIDERS
                }
              }
            ]
          },
          {
            path: 'new_admin',
            name: 'new_admin',
            component: NewFunctionalUserVue,
            props: true,
            meta: {
              requiresAuth: true,
              pageType: PagesEnum.EMPLOYEES
            },
            children: [
              {
                path: '',
                name: 'admin_personal_step',
                component: FunctionalUserPersonalStep1Vue,
                props: true,
                meta: {
                  requiresAuth: true,
                  pageType: PagesEnum.EMPLOYEES
                }
              },
            ]
          },
          {
            path: '',
            name: 'info',
            component: ProviderDetailsView,
            props: { useSelectedProvider: true },
            meta: {
              requiresAuth: true,
              pageType: PagesEnum.PROVIDERS
            }
          },

          {
            path: 'admins',
            name: 'admins_router',
            component: EmptyRouting,
            meta: {
              requiresAuth: true,
              pageType: PagesEnum.PROVIDERS,
            },
            children: [
              {
                path: '',
                name: 'admins',
                component: AdministratorsListView,
                meta: {
                  requiresAuth: true,
                  pageType: PagesEnum.PROVIDERS,
                },
              },
              {
                path: 'provider_admins',
                name: 'provider_admins_router',
                component: EmptyRouting,
                meta: {
                  requiresAuth: true,
                  pageType: PagesEnum.PROVIDERS,
                },
                children: [
                  {
                    path: '',
                    name: 'provider_admins',
                    component: AdminDetailsView,
                    props: true,
                    meta: {
                      requiresAuth: true,
                      pageType: PagesEnum.PROVIDERS,
                      breadcrumbElement: BreadcrumbElementType.SELECTED_EMPLOYEE_NAME
                    },
                  },
                  {
                    path: 'edit_admin',
                    name: 'edit_admin',
                    component: NewFunctionalUserVue,
                    props: {passedRoleID: RoleEnum.ADMINISTRATOR_FULL, previousPath: '/providers/provider_tabs/admins', useSelectedEmployee: true},
                    meta: {
                      requiresAuth: true,
                      pageType: PagesEnum.EMPLOYEES
                    },
                    children: [
                      {
                        path: '',
                        name: 'edit_admin_personal_step',
                        component: FunctionalUserPersonalStep1Vue,
                        props: true,
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
          }
        ]
      },
    ]
  };

export default providersRoutes;
