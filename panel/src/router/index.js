import Vue from "vue";
import VueRouter from "vue-router";
// import { component } from "vue/types/umd";
import Home from "../views/Dashboard.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "",
    redirect: "/",
    name: "Dashboard",
    component: () => import("../views/Index.vue"),
    children: [
      {
        path: "/",
        name: "About",
        component: Home,
      },
      {
        path: "/shopping-page",
        name: "Shooping",
        component: () => import("../views/ShoopingPage.vue"),
      },
      {
        path: "/contact",
        name: "Contact",
        component: () => import("../views/Contact.vue"),
      },
      {
        path: "/product/:id",
        name: "Show Product",
        component: () => import("../views/pages/product/_Show.vue"),
      },
    ],
  },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

export default router;
