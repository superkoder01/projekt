import newCustomer from "@/components/forms/create-customer/newCustomer.vue";
import { PagesEnum } from "@/services/permissions/pages-enum";
import PersonalStep1 from "@/components/forms/create-customer/PersonalStep1.vue";
import AddressStep2 from "@/components/forms/create-customer/AddressStep2.vue";
import CustomerStep3 from "@/components/forms/create-customer/CustomerStep3.vue";
import CustomersView from "@/views/CustomersListView.vue";
import CustomerTabsView from "@/views/CustomerTabsView.vue";
import { BreadcrumbElementType } from "@/router/breadcrumbElementType";
import CustomerDetailsView from "@/views/CustomerDetailsView.vue";
import OffersView from "@/views/OffersListView.vue";
import ContractsView from "@/views/ContractsListView.vue";
import InvoicesView from "@/views/InvoicesListView.vue";
import InstallationsView from "@/views/InstallationsListView.vue";
import NewOffer from "@/components/forms/create-offer/NewOffer.vue";
import OfferStep1 from "@/components/forms/create-offer/OfferStep1.vue";
import OfferStep2 from "@/components/forms/create-offer/OfferStep2.vue";
import NewContract from "@/components/forms/create-contract/NewContract.vue";
import ContractStep1 from "@/components/forms/create-contract/ContractStep1.vue";
import ContractStep2 from "@/components/forms/create-contract/ContractStep2.vue";
import ContractStep3 from "@/components/forms/create-contract/ContractStep3.vue";
import ContractStep4 from "@/components/forms/create-contract/ContractStep4.vue";
import NewInstallationVue from "@/components/forms/create-installation/NewInstallation.vue";
import InstallationStep1Vue from "@/components/forms/create-installation/InstallationStep1.vue";
import InstallationStep2Vue from "@/components/forms/create-installation/InstallationStep2.vue";
import EmptyRouting from "@/views/EmptyRouting.vue";

const customersRoute =
    {
      path: '/customers',
      name: 'customer_routing',
      component: EmptyRouting,
      meta: {
        requiresAuth: true,
        pageType: PagesEnum.CUSTOMERS
      },
      children: [
        {
          path: '',
          name: 'customers',
          component: CustomersView,
          meta: {
            requiresAuth: true,
            pageType: PagesEnum.CUSTOMERS
          },
        },
        {
          path: 'new_customer',
          name: 'new_customer',
          props: true,
          component: newCustomer,
          meta: {
            requiresAuth: true,
            pageType: PagesEnum.CUSTOMERS
          },
          children: [
            {
              path: '',
              name: 'personal_step_1',
              component: PersonalStep1,
              meta: {
                requiresAuth: true,
                pageType: PagesEnum.CUSTOMERS
              }
            },
            {
              path: 'address_step_2',
              name: 'address',
              component: AddressStep2,
              meta: {
                requiresAuth: true,
                pageType: PagesEnum.CUSTOMERS
              }
            },
          ]
        },
        {
          path: 'new_business_customer',
          name: 'new_business_customer',
          component: newCustomer,
          props: {isBusinessClient: true},
          meta: {
            requiresAuth: true,
            pageType: PagesEnum.CUSTOMERS
          },
          children: [
            {
              path: '',
              name: 'business_step_1',
              component: CustomerStep3,
              meta: {
                requiresAuth: true,
                pageType: PagesEnum.CUSTOMERS
              }
            },
            {
              path: 'personal_step_2',
              name: 'personal_step_2',
              component: PersonalStep1,
              meta: {
                requiresAuth: true,
                pageType: PagesEnum.CUSTOMERS
              }
            },
            {
              path: 'address_step_2',
              name: 'business_address_step_3',
              component: AddressStep2,
              meta: {
                requiresAuth: true,
                pageType: PagesEnum.CUSTOMERS
              }
            },
          ]
        },
        {
          path: 'customer_tabs',
          name: 'customer_tabs',
          // route level code-splitting
          // this generates a separate chunk (about.[hash].js) for this route
          // which is lazy-loaded when the route is visited.
          component: CustomerTabsView,
          meta: {
            requiresAuth: true,
            pageType: PagesEnum.CUSTOMERS,
            breadcrumbElement: BreadcrumbElementType.SELECTED_CUSTOMER_NAME
          },
          children: [

            {
              path: '',
              name: 'customer_tabs_info',
              component: CustomerDetailsView,
              meta: {
                requiresAuth: true,
                pageType: PagesEnum.CUSTOMERS
              }
            },
            {
              path: 'edit_customer',
              name: 'edit_customer',
              props: {useSelectedCustomer: true},
              component: newCustomer,
              meta: {
                requiresAuth: true,
                pageType: PagesEnum.CUSTOMERS
              },
              children: [
                {
                  path: '',
                  name: 'edit_personal_step_1',
                  component: PersonalStep1,
                  props: {useSelectedCustomer: true},
                  meta: {
                    requiresAuth: true,
                    pageType: PagesEnum.CUSTOMERS
                  }
                },
                {
                  path: 'address_step_2',
                  name: 'edit_address',
                  component: AddressStep2,
                  props: {useSelectedCustomer: true},
                  meta: {
                    requiresAuth: true,
                    pageType: PagesEnum.CUSTOMERS
                  }
                },
              ]
            },
            {
              path: 'offers',
              name: 'customer_tabs_offers_routing',
              component: EmptyRouting,
              props: {useSelectedCustomer: true},
              meta: {
                requiresAuth: true,
                pageType: PagesEnum.CUSTOMERS
              },
              children: [
                {
                  path: '',
                  name: 'customer_tabs_offers',
                  component: OffersView,
                  props: {useSelectedCustomer: true},
                  meta: {
                    requiresAuth: true,
                    pageType: PagesEnum.CUSTOMERS
                  },
                },
                {
                  path: 'new_offer',
                  name: 'new_offer',
                  component: NewOffer,
                  meta: {
                    requiresAuth: true,
                    pageType: PagesEnum.CUSTOMERS
                  },
                  children: [
                    {
                      path: '',
                      name: 'offer_step1',
                      component: OfferStep1,
                      meta: {
                        requiresAuth: true,
                        pageType: PagesEnum.CUSTOMERS
                      }
                    },
                    {
                      path: 'offer_step2',
                      name: 'offer_step2',
                      component: OfferStep2,
                      meta: {
                        requiresAuth: true,
                        pageType: PagesEnum.CUSTOMERS
                      }
                    },
                  ]
                },
                {
                  path: 'edit_offer',
                  name: 'edit_offer',
                  component: NewOffer,
                  props: {useSelectedOffer: true},
                  meta: {
                    requiresAuth: true,
                    pageType: PagesEnum.CUSTOMERS
                  },
                  children: [
                    {
                      path: '',
                      name: 'edit_offer_step1',
                      component: OfferStep1,
                      meta: {
                        requiresAuth: true,
                        pageType: PagesEnum.CUSTOMERS
                      }
                    },
                    {
                      path: 'offer_step2',
                      name: 'edit_offer_step2',
                      component: OfferStep2,
                      meta: {
                        requiresAuth: true,
                        pageType: PagesEnum.CUSTOMERS
                      }
                    },
                  ]
                },
                {
                  path: 'new_contract',
                  name: 'new_contract',
                  component: NewContract,
                  meta: {
                    requiresAuth: true,
                    pageType: PagesEnum.CUSTOMERS
                  },
                  children: [
                    {
                      path: '',
                      name: 'contract_step1',
                      component: ContractStep1,
                      meta: {
                        requiresAuth: true,
                        pageType: PagesEnum.CUSTOMERS
                      }
                    },
                    {
                      path: 'contract_step2',
                      name: 'contract_step2',
                      component: ContractStep2,
                      meta: {
                        requiresAuth: true,
                        pageType: PagesEnum.CUSTOMERS
                      }
                    },
                    {
                      path: 'contract_step3',
                      name: 'contract_step3',
                      component: ContractStep3,
                      meta: {
                        requiresAuth: true,
                        pageType: PagesEnum.CUSTOMERS
                      }
                    },
                    {
                      path: 'contract_step4',
                      name: 'contract_step4',
                      component: ContractStep4,
                      meta: {
                        requiresAuth: true,
                        pageType: PagesEnum.CUSTOMERS
                      }
                    },
                  ]
                },
              ]
            },
            {
              path: 'contracts',
              name: 'customer_tabs_contracts_router',
              component: EmptyRouting,
              meta: {
                requiresAuth: true,
                pageType: PagesEnum.CUSTOMERS
              },
              children: [
                {
                  path: '',
                  name: 'customer_tabs_contracts',
                  component: ContractsView,
                  props: {useSelectedCustomer: true},
                  meta: {
                    requiresAuth: true,
                    pageType: PagesEnum.CUSTOMERS
                  }
                },
                {
                  path: 'edit_contract',
                  name: 'edit_contract',
                  component: NewContract,
                  props: {useSelectedContract: true},
                  meta: {
                    requiresAuth: true,
                    pageType: PagesEnum.CUSTOMERS
                  },
                  children: [
                    {
                      path: '',
                      name: 'edit_contract_step1',
                      component: ContractStep1,
                      meta: {
                        requiresAuth: true,
                        pageType: PagesEnum.CUSTOMERS
                      }
                    },
                    {
                      path: 'contract_step2',
                      name: 'edit_contract_step2',
                      component: ContractStep2,
                      meta: {
                        requiresAuth: true,
                        pageType: PagesEnum.CUSTOMERS
                      }
                    },
                    {
                      path: 'contract_step3',
                      name: 'edit_contract_step3',
                      component: ContractStep3,
                      meta: {
                        requiresAuth: true,
                        pageType: PagesEnum.CUSTOMERS
                      }
                    },
                    {
                      path: 'contract_step4',
                      name: 'edit_contract_step4',
                      component: ContractStep4,
                      meta: {
                        requiresAuth: true,
                        pageType: PagesEnum.CUSTOMERS
                      }
                    },
                  ]
                },
              ]
            },

            {
              path: 'invoices',
              name: 'customer_tabs_invoices',
              component: InvoicesView,
              props: {useSelectedCustomer: true},
              meta: {
                requiresAuth: true,
                pageType: PagesEnum.CUSTOMERS
              }
            },
            {
              path: 'installations',
              name: 'customer_tabs_installations_routing',
              component: EmptyRouting,
              props: {useSelectedCustomer: true},
              meta: {
                requiresAuth: true,
                pageType: PagesEnum.CUSTOMERS
              },
              children: [
                {
                  path: '',
                  name: 'customer_tabs_installations',
                  component: InstallationsView,
                  props: {useSelectedCustomer: true},
                  meta: {
                    requiresAuth: true,
                    pageType: PagesEnum.CUSTOMERS
                  },
                },
                {
                  path: 'new_installation',
                  name: 'new_installation',
                  component: NewInstallationVue,
                  meta: {
                    requiresAuth: true
                  },
                  children:[
                    {
                      path: '',
                      name: 'installation_data_step_1',
                      component: InstallationStep1Vue ,
                      props: true,
                      meta: {
                        requiresAuth: true
                      }
                    },
                    {
                      path: 'address_2',
                      name: 'installation_address',
                      component: InstallationStep2Vue,
                      meta: {
                        requiresAuth: true
                      }
                    },
                  ]
                },
              ]
            }
          ]
        }
      ]
    };

export default customersRoute
