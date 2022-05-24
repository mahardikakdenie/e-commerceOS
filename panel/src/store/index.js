import Vue from "vue";
import Vuex from "vuex";
import auth from "./modules/auth";
import cart from "./modules/cart";
import store from "./modules/store";
import category from "./modules/store/category";
import product from "./modules/store/product";

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    auth,
    store,
    category,
    product,
    cart,
  },
});
