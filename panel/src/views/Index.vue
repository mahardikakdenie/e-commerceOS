<template>
  <!-- <div class="home"> -->
  <div>
    <Navbar @logout="logout" />
    <router-view></router-view>
    <Footer />
  </div>
  <!-- </div> -->
</template>

<script>
import Navbar from "@/components/Navbar.vue";
import Footer from "@/components/footer/Footer.vue";
import axios from "axios";
export default {
  components: {
    Footer,
    Navbar,
  },
  data: () => ({
    customer: {},
    store: {},
    isLoad: false,
    count: 0,
  }),
  computed: {
    isLogin() {
      return this.$store.state.auth.isLogin;
    },
    computedMe() {
      return this.$store.state.auth.me;
    },
    computedStore() {
      return this.$store.state.store.store;
    },
  },
  mounted() {
    this.getStore();
    this.me();
    this.checkLogout();
  },
  methods: {
    getStore() {
      this.$store
        .dispatch("store/getStore", {
          wilcard: this.$route.params.wilcard,
        })
        .then((res) => {
          if (res.data.meta.status) {
            localStorage.setItem("store_id", res.data.data.id);
            // this.$store.commit("store/SET_RELOADED", false);
            const path = this.$route.path;
            const new_array_path = path.split("/");
            if (new_array_path.length > 3) {
              this.$route.path = `/himatif`;
              localStorage.setItem("reloaded", "false");
            }
            if (!this.$store.state.store.reloaded) {
              this.$store.commit("store/SET_RELOADED", true);
            }
          }
        });
    },
    me() {
      this.$store
        .dispatch("auth/me", {
          entities: "Store",
        })
        .then((res) => {
          if (res.data.meta.status) {
            localStorage.setItem("store_customer_id", res.data.data.store.ID);
          }
        });
    },
    checkLogout() {
      if (localStorage.getItem("access_token") === null) {
        this.$router.push(`/${this.$route.params.wilcard}/`);
        const Toast = this.$swal.mixin({
          toast: true,
          position: "bottom-end",
          showConfirmButton: false,
          timer: 3000,
          timerProgressBar: true,
          didOpen: (toast) => {
            toast.addEventListener("mouseenter", this.$swal.stopTimer);
            toast.addEventListener("mouseleave", this.$swal.resumeTimer);
          },
          popup: "swal2-show",
          backdrop: "swal2-backdrop-show",
          icon: "swal2-icon-show",
        });
        Toast.fire({
          icon: "success",
          title: `Selamat datang di ${localStorage
            .getItem("name_store")
            .toUpperCase()}`,
        });
      }
      if (
        localStorage.getItem("store_id") !==
          localStorage.getItem("store_customer_id") &&
        localStorage.getItem("access_token") !== null
      ) {
        localStorage.removeItem("access_token");
        localStorage.setItem("isLogin", false);
        if (localStorage.getItem("isLogin") !== "false") {
          location.reload();
        }
        const Toast = this.$swal.mixin({
          toast: true,
          position: "bottom-end",
          showConfirmButton: false,
          timer: 3000,
          timerProgressBar: true,
          didOpen: (toast) => {
            toast.addEventListener("mouseenter", this.$swal.stopTimer);
            toast.addEventListener("mouseleave", this.$swal.resumeTimer);
          },
          popup: "swal2-show",
          backdrop: "swal2-backdrop-show",
          icon: "swal2-icon-show",
        });
        Toast.fire({
          icon: "success",
          title: "You Logout",
        });
      }
    },
    logout() {
      axios.defaults.headers.common.Authorization =
        "Bearer " + localStorage.getItem("access_token");
      axios.defaults.baseURL = process.env.VUE_APP_API_URL;
      this.$store.dispatch("auth/logout").then((res) => {
        if (res.data.meta.status) {
          location.reload();
          const Toast = this.$swal.mixin({
            toast: true,
            position: "bottom-end",
            showConfirmButton: false,
            timer: 3000,
            timerProgressBar: true,
            didOpen: (toast) => {
              toast.addEventListener("mouseenter", this.$swal.stopTimer);
              toast.addEventListener("mouseleave", this.$swal.resumeTimer);
            },
            popup: "swal2-show",
            backdrop: "swal2-backdrop-show",
            icon: "swal2-icon-show",
          });
          Toast.fire({
            icon: "success",
            title: "Logout Success",
          });
        }
      });
    },
  },
};
</script>

<style></style>
