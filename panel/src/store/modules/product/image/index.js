import axios from "axios";

export default {
  namespaced: true,
  state: {
    data: [],
  },
  mutations: {
    GET_DATA: (state, payload) => (state.data = payload),
  },
  actions: {
    getDataProductImage: ({ commit }, payload) => {
      axios.defaults.headers.common.Authorization =
        "Bearer " + localStorage.getItem("access_token");
      axios.defaults.baseURL = process.env.VUE_APP_API_URL;

      return new Promise((resolve, reject) => {
        const params = {
          ...payload,
        };
        axios
          .get(`product-image/${payload.product_id}`, {
            params: params,
          })
          .then((res) => {
            commit("GET_DATA", res.data.data);
            resolve(res);
          })
          .catch((e) => {
            reject(e);
          });
      });
    },
  },
};
