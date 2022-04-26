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
    ],
  },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

export default router;
