import Vue from "vue";
import Vuex from "vuex";
import auth from "./modules/auth";
import cart from "./modules/cart";
import store from "./modules/store";
import category from "./modules/store/category";
import product from "./modules/store/product";
import all_product from "./modules/product";
import product_image from "./modules/product/image";

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    auth,
    store,
    category,
    product,
    cart,
    all_product,
    product_image,
  },
});
