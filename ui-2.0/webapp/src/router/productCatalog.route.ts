import ProductCatalog from "@/views/ProductCatalog.vue";
import EmptyRouting from "@/views/EmptyRouting.vue";
import OfferDraftListView from "@/views/OfferDraftListView.vue";
import NewOfferDraft from "@/components/forms/create-draft-offer/NewOfferDraft.vue";
import DraftOfferStep1 from "@/components/forms/create-draft-offer/DraftOfferStep1.vue";
import DraftOfferStep2 from "@/components/forms/create-draft-offer/DraftOfferStep2.vue";
import DraftOfferStep3 from "@/components/forms/create-draft-offer/DraftOfferStep3.vue";
import {PagesEnum} from "@/services/permissions/pages-enum";
import NewPricing from "@/components/forms/create-pricing/NewPricing.vue";
import PricingStep1 from "@/components/forms/create-pricing/PricingStep1.vue";
import PricingStep2 from "@/components/forms/create-pricing/PricingStep2.vue";
import PricingView from "@/views/PricingListView.vue";

const productCatalogRoutes =
  {
    path: '/product_catalog',
    name: 'product_catalog',
    component: ProductCatalog,
    meta: {
      requiresAuth: true,
      pageType: PagesEnum.PRODUCT_CATALOG

    },
    children: [
      {
        path: 'pricing',
        name: 'pricing',
        component: EmptyRouting,
        meta: {
          requiresAuth: true,
          pageType: PagesEnum.PRODUCT_CATALOG
        },
        children: [
          {
            path: '',
            name: 'pricing_table',
            component: PricingView,
            meta: {
              requiresAuth: true,
              pageType: PagesEnum.PRODUCT_CATALOG
            }
          },
          {
            path: 'new_pricing',
            name: 'new_pricing',
            component: NewPricing,
            meta: {
              requiresAuth: true,
              pageType: PagesEnum.PRODUCT_CATALOG
            },
            children: [
              {
                path: '',
                name: 'pricing_step1',
                component: PricingStep1,
                meta: {
                  requiresAuth: true,
                  pageType: PagesEnum.PRODUCT_CATALOG
                }
              },
              {
                path: 'pricing_2',
                name: 'pricing_step2',
                component: PricingStep2,
                meta: {
                  requiresAuth: true,
                  pageType: PagesEnum.PRODUCT_CATALOG
                }
              },
            ]
          },
        ]
      },
      {
        path: 'offer_draft',
        name: 'offer_draft',
        component: EmptyRouting,
        meta: {
          // requiresAuth: true
          pageType: PagesEnum.PRODUCT_CATALOG
        },
        children: [
          {
            path: '',
            name: 'offer_draft_list',
            component: OfferDraftListView,
            meta: {
              // requiresAuth: true
              pageType: PagesEnum.PRODUCT_CATALOG
            },
          },
          {
            path: 'new_offer_draft',
            name: 'new_offer_draft',
            component: NewOfferDraft,
            meta: {
              // requiresAuth: true
              pageType: PagesEnum.PRODUCT_CATALOG
            },
            redirect: 'draftOfferStep1',
            children: [
              {
                path: '',
                name: 'new_offer_draft_step1',
                component: DraftOfferStep1,
                meta: {
                  // requiresAuth: true'
                  pageType: PagesEnum.PRODUCT_CATALOG
                }
              },
              {
                path: 'step2',
                name: 'new_offer_draft_step2',
                component: DraftOfferStep2,
                meta: {
                  // requiresAuth: true
                  pageType: PagesEnum.PRODUCT_CATALOG
                }
              },
              {
                path: 'step3',
                name: 'new_offer_draft_step3',
                component: DraftOfferStep3,
                meta: {
                  requiresAuth: true,
                  pageType: PagesEnum.PRODUCT_CATALOG
                }
              }
            ]
          },
        ]
      },
    ]
  };

export default productCatalogRoutes;

