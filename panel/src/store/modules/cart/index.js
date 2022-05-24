import axios from "axios";

export default {
  namespaced: true,
  state: {
    data: [],
  },
  getters: {},
  mutations: {
    GET_CART: (state, payload) => (state.data = payload),
  },
  actions: {
    getDataCart: ({ commit }, payload) => {
      axios.defaults.headers.common.Authorization =
        "Bearer " + localStorage.getItem("access_token");
      axios.defaults.baseURL = process.env.VUE_APP_API_URL;

      return new Promise((resolve, reject) => {
        const params = {
          ...payload,
        };
        axios
          .get(`cart_customer/carts`, {
            params: params,
          })
          .then((res) => {
            commit("GET_CART", res.data.data);
            resolve(res);
            console.log("payload => ", payload);
          })
          .catch((e) => {
            reject(e);
          });
      });
    },
  },
};
