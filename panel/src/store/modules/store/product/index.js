import axios from "axios";
export default {
  namespaced: true,
  state: {
    views: [],
    product_category: [],
  },
  getters: {},
  mutations: {
    SET_VIEWS: (state, payload) => (state.views = payload),
  },
  actions: {
    getBestSeller: ({ commit }, payload) => {
      axios.defaults.headers.common.Authorization =
        "Bearer " + localStorage.getItem("access_token");
      axios.defaults.baseURL = process.env.VUE_APP_API_URL;

      return new Promise((resolve, reject) => {
        const params = { ...payload };
        axios
          .get("product-customer", { params: params })
          .then((res) => {
            commit("SET_VIEWS", res.data.data);
            resolve(res);
          })
          .catch((e) => {
            reject(e);
          });
      });
    },
    getProductCategory: ({ commit }, payload) => {
      axios.defaults.headers.common.Authorization =
        "Bearer " + localStorage.getItem("access_token");
      axios.defaults.baseURL = process.env.VUE_APP_API_URL;

      return new Promise((resolve, reject) => {
        const params = { ...payload };
        axios
          .get(`product-customer/${payload.slug}`, { params: params })
          .then((res) => {
            commit("SET_VIEWS", res.data.data);
            resolve(res);
          })
          .catch((e) => {
            reject(e);
          });
      });
    },
  },
};
