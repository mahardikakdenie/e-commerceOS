import axios from "axios";

export default {
  namespaced: true,
  state: {
    wildcard: "",
    reloaded: false,
    store: {},
  },
  mutations: {
    SET_STORE: (store, payload) => {
      store.store = payload;
    },
    GET_WILDCARD: (state, wildcard) => {
      state.wildcard = wildcard;
    },
    SET_RELOADED: (state, reloaded) => {
      state.reloaded = reloaded;
    },
  },
  getters: {},
  actions: {
    getStore: ({ commit }, payload) => {
      axios.defaults.baseURL = process.env.VUE_APP_API_URL;
      return new Promise((resolve, reject) => {
        const params = { ...payload };
        axios
          .get(`store-global/${payload.wilcard}`, { params })
          .then((res) => {
            if (res.data.meta.status) {
              commit("SET_STORE", res.data.data);
              resolve(res);
              commit("GET_WILDCARD", res.data.data.slug);
              localStorage.setItem("wild_route", res.data.data.slug);
              localStorage.setItem("name_store", res.data.data.name);
              localStorage.setItem("  ", res.data.data.id);
            } else {
              console.log(res.data.meta.message);
              this.$router.push(`/${localStorage.getItem("wild_route")}/`);
            }
          })
          .catch((e) => {
            reject(e);
          });
      });
    },
  },
};
