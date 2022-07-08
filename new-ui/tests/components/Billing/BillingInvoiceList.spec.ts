import { createVuetify } from "vuetify";
import { flushPromises, mount, VueWrapper } from "@vue/test-utils";
import { beforeEach, describe, expect, it, vi } from "vitest";
import BillingInvoiceList from "../../../src/components/Billing/BillingInvoiceList.vue";
import { createStore } from "vuex";
import { key } from "../../../src/store";
import routes from "../../../src/router";

const invoices = [
  {
    paid: false,
    amountDue: 3040,
    dueDate: 1614983421,
    status: 'open',
    pdf: 'test.pdf',
    url: 'inv_url',
  },
  {
    paid: true,
    amountDue: 2030,
    dueDate: 1614983421,
    status: 'paid',
    pdf: '---',
    url: '---',
  },
];

const headers = [
  {
    text: "Status",
    value: "status",
    align: "center",
    sortable: false,
  },
  {
    text: "Due date",
    value: "dueDate",
    align: "center",
    sortable: false,
  },
  {
    text: "Amount",
    value: "amountDue",
    align: "center",
    sortable: false,
  },
  {
    text: "PDF",
    value: "pdf",
    align: "center",
    sortable: false,
  },
  {
    text: "URL",
    value: "url",
    align: "center",
    sortable: false,
  },
];

const store = createStore({
  state: {
    invoices,
    invoicesLength: invoices.length,
    defaultPerPage: 3,
  },
  getters: {
    'billing/getInvoices': (state) => state.invoices,
    'billing/defaultPerPage': (state) => state.defaultPerPage,
    'billing/getInvoicesLength': (state) => state.invoices.length,
  },
  actions: {
    'billing/getPagination': vi.fn(),
  },
});
describe('BillingInvoiceList', () => {
  let wrapper: VueWrapper<any>;
  const vuetify = createVuetify();

  beforeEach(() => {
    wrapper = mount(BillingInvoiceList, {
      global: {
        plugins: [[store, key], routes, vuetify],
      },
      shallow: true,
    });
  });

  ///////
  // Component Rendering
  //////

  it("Is a Vue instance", () => {
    expect(wrapper).toBeTruthy();
  });
  it("Renders the component", () => {
    expect(wrapper.html()).toMatchSnapshot();
  });

  ///////
  // Data checking
  //////
  it("Data is defined", () => {
    expect(wrapper.vm.$data).toBeDefined();
  });

  it('Compares data with default value', () => {
    expect(wrapper.vm.headers).toStrictEqual(headers);
  });
  it('Process data in the computed', () => {
    expect(wrapper.vm.invoiceList).toStrictEqual(invoices);
  });

  // todo check the pagination
});