import Vue from "vue";
import VueRouter from "vue-router";
// import { component } from "vue/types/umd";
import Home from "../views/Dashboard.vue";

Vue.use(VueRouter);

const routes = [
  {
    path: "/:wilcard",
    name: "Dashboard",
    component: () => import("../views/Index.vue"),
    children: [
      {
        path: "",
        name: "About",
        component: Home,
        meta: {
          requiresAuth: true,
        },
      },
      {
        path: "shopping-page",
        name: "Shooping",
        component: () => import("../views/ShoopingPage.vue"),
        meta: {
          requiresAuth: true,
        },
      },
      {
        path: "contact",
        name: "Contact",
        component: () => import("../views/Contact.vue"),
        meta: {
          requiresAuth: true,
        },
      },
      {
        path: "cart",
        name: "Cart",
        component: () => import("../views/Cart.vue"),
        meta: {
          requiresAuth: true,
        },
      },
      {
        path: "login",
        name: "Login",
        component: () => import("../views/Login.vue"),
        meta: {
          requiresVisitor: true,
        },
      },
      {
        path: "register",
        name: "Register",
        component: () => import("../views/Register.vue"),
        meta: {
          requiresVisitor: true,
        },
      },
      {
        path: "product/:id",
        name: "Show Product",
        component: () => import("../views/pages/product/_Show.vue"),
        meta: {
          requiresAuth: true,
        },
      },
    ],
  },
  {
    path: "/",
    name: "error",
    component: () => import("../views/Blank.vue"),
    meta: {
      error: true,
    },
  },
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes,
});

// router.beforeEach((to, from, next) => {
//   if (to.matched.some((record) => record.meta.requiresAuth)) {
//     // this route requires auth, check if logged in
//     // if not, redirect to login page.
//     // if (!localStorage.getItem("token")) {
//     //   next({ name: "Login" });
//     // }
//   } else if (to.matched.some((record) => record.meta.requiresVisitor)) {
//     // this route requires auth, check if logged in
//     // if not, redirect to login page.
//     if (localStorage.getItem("access_token")) {
//       // console.log(localStorage.getItem("access_item"));
//       // if (to.name === "Login") {
//       //   next({ name: "About" });
//       // }
//       // next({
//       //   path: "/",
//       //   query: { redirect: to.fullPath },
//       // });
//     } else {
//       next();
//     }
//   } else {
//     next(); // make sure to always call next()!
//   }
// });

export default router;
