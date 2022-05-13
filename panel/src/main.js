import Vue from "vue";
import App from "./App.vue";
import router from "./router";
import "bootstrap/dist/css/bootstrap.min.css";
import "bootstrap/dist/js/bootstrap.js";
import "./assets/css/index.css";
import store from "./store";
import VueSweetalert2 from "vue-sweetalert2";

// If you don't need the styles, do not connect
import "sweetalert2/dist/sweetalert2.min.css";
Vue.config.productionTip = false;
Vue.use(VueSweetalert2);

new Vue({
  router,

  store,
  render: (h) => h(App),
}).$mount("#app");
