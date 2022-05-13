import Vue from "vue";
import Vuex from "vuex";
import auth from "./modules/auth";
import store from "./modules/store";
import category from "./modules/store/category";
Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    auth,
    store,
    category,
  },
});
