import axios from "axios";

export default {
  namespaced: true,
  state: {
    token: null || localStorage.getItem("access_token"),
    isLogin: false,
    me: {},
  },
  mutations: {
    GET_TOKEN: (state, token) => {
      state.token = token;
    },
    SET_ISLOGIN: (state, isLogin) => {
      state.isLogin = isLogin;
    },
    SET_ME: (state, payload) => {
      state.me = payload;
    },
  },
  actions: {
    login: ({ commit }, payload) => {
      axios.defaults.baseURL = process.env.VUE_APP_API_URL;
      return new Promise((resolve, reject) => {
        const data = new FormData();
        data.append("username", payload.email);
        data.append("password", payload.password);
        data.append("store_slug", payload.store_slug);
        console.log("payload =>", payload);
        axios
          .post("customer_auth/login", data)
          .then((res) => {
            resolve(res);
            if (res.data.meta.status) {
              commit("SET_ISLOGIN", true);
              commit("GET_TOKEN", res.data.data.token);
              localStorage.setItem("access_token", res.data.data.token);
              localStorage.setItem("isLogin", "true");
            }
          })
          .catch((e) => {
            reject(e);
          });
      });
    },
    logout() {
      axios.defaults.headers.common.Authorization =
        "Bearer " + localStorage.getItem("access_token");
      axios.defaults.baseURL = process.env.VUE_APP_API_URL;
      return new Promise((resolve, reject) => {
        axios
          .post("customer_auth/logout", {})
          .then((res) => {
            resolve(res);
            localStorage.removeItem("access_token");
            localStorage.setItem("reloaded", "false");
            localStorage.setItem("isLogin", "false");
          })
          .catch((e) => {
            reject(e);
          });
      });
    },
    me({ commit }, payload) {
      axios.defaults.headers.common.Authorization =
        "Bearer " + localStorage.getItem("access_token");
      axios.defaults.baseURL = process.env.VUE_APP_API_URL;

      return new Promise((resolve, reject) => {
        const params = { ...payload };
        axios
          .get("customer/me", { params: params })
          .then((res) => {
            resolve(res);
            commit("SET_ME", res.data.data);
          })
          .catch((e) => {
            reject(e);
          });
      });
    },
  },
};
