import axios from "axios";

export default {
  namespaced: true,
  state: {
    categories: [],
  },
  mutations: {
    SET_CATEGORIES: (state, payload) => (state.categories = payload),
  },
  getters: {},
  actions: {
    getCategories: ({ commit }, payload) => {
      axios.defaults.headers.common.Authorization =
        "Bearer " + localStorage.getItem("access_token");
      axios.defaults.baseURL = process.env.VUE_APP_API_URL;

      return new Promise((resolve, reject) => {
        const params = { ...payload };
        axios
          .get("category-store", {
            params: params,
          })
          .then((res) => {
            resolve(res);
            commit("SET_CATEGORIES", res.data.data);
          })
          .catch((e) => {
            reject(e);
          });
      });
    },
  },
};
