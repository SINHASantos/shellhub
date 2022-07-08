import { Module } from "vuex";
import { State } from "./../index";
import * as apiPublicKey from '../api/public_keys';

export interface PublicKeysState {
  publicKeys: Array<any>,
  publicKey: any,
  numberPublicKeys: number,
  page: number,
  perPage: number,
}

export const publicKeys: Module<PublicKeysState, State> = {
  namespaced: true,
  state: {
    publicKeys: [],
    publicKey: {},
    numberPublicKeys: 0,
    page: 1,
    perPage: 10,
  },

  getters: {
    list: (state) => state.publicKeys,
    get: (state) => state.publicKey,
    getNumberPublicKeys: (state) => state.numberPublicKeys,
    getPage: (state) => state.page,
    getPerPage: (state) => state.perPage,
  },

  mutations: {
    setPublicKeys: (state, res) => {
      state.publicKeys = res.data;
      state.numberPublicKeys = parseInt(res.headers['x-total-count'], 10);
    },

    setPublicKey: (state, res) => {
      state.publicKey = res.data;
    },

    setPagePerpage: (state, data) => {
      state.page = data.page;
      state.perPage = data.perPage;
    },

    resetPagePerpage: (state) => {
      state.page = 1;
      state.perPage = 10;
    },

    removePublicKey: (state, fingerprint) => {
      state.publicKeys.splice(state.publicKeys.findIndex((d) => d.fingerprint === fingerprint), 1);
    },

    clearListPublicKeys: (state) => {
      state.publicKeys = [];
      state.numberPublicKeys = 0;
    },

    clearObjectPublicKey: (state) => {
      state.publicKey = {};
    },
  },

  actions: {
    post: async (context, data) => {
      await apiPublicKey.postPublicKey(data);
    },

    fetch: async (context, data) => {
      try {
        const res = await apiPublicKey.fetchPublicKeys(data.page, data.perPage, "");
        if (res.data.length) {
          context.commit('setPublicKeys', res);
          context.commit('setPagePerpage', data);
          return true;
        } 
        return false;
      } catch (error) {
        context.commit('clearListPublicKeys');
        throw error;
      }
    },

    refresh: async (context) => {
      try {
        const res = await apiPublicKey.fetchPublicKeys(
          context.state.page,
          context.state.perPage,
          "",
        );
        context.commit('setPublicKeys', res);
      } catch (error) {
        context.commit('clearListPublicKeys');
        throw error;
      }
    },

    get: async (context, id) => {
      try {
        const res = await apiPublicKey.getPublicKey(id);
        context.commit('setPublicKey', res);
      } catch (error) {
        context.commit('clearObjectPublicKey');
        throw error;
      }
    },

    put: async (context, data) => {
      await apiPublicKey.putPublicKey(data);
    },

    resetPagePerpage: async (context) => {
      context.commit('resetPagePerpage');
    },

    remove: async (context, fingerprint) => {
      await apiPublicKey.removePublicKey(fingerprint);
    },
  },
};
