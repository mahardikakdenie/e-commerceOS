<template>
  <div class="container login">
    <div class="card">
      <div class="row">
        <div class="col">
          <h2>Sign In</h2>
        </div>
      </div>
      <div class="row form--login">
        <form @submit.prevent.enter="login" action="">
          <div class="col form--email">
            <div class="form-group">
              <input
                v-model="form.email"
                type="text"
                class="form-control"
                id="exampleInputEmail1"
                aria-describedby="emailHelp"
                placeholder="Enter email"
              />
            </div>
          </div>
          <div class="col">
            <div class="form-group">
              <input
                v-model="form.password"
                type="password"
                class="form-control"
                placeholder="Enter Password"
              />
            </div>
          </div>
          <div class="col btn-login">
            <button type="submit" class="btn">Login</button>
          </div>
          <div class="col btn--register">
            <p>Belom Punya Akun ? <a href="">Buat Akun</a></p>
          </div>
          <div class="col btn--forgot">
            <p><a href="">Lupa Password</a></p>
          </div>
        </form>
      </div>
    </div>
  </div>
</template>

<script>
import axios from "axios";
export default {
  data: () => ({
    form: {
      email: "",
      password: "",
    },
    customer: {},
  }),

  computed: {
    isLogin() {
      return this.$store.state.auth.isLogin;
    },
  },
  methods: {
    login() {
      this.$store
        .dispatch("auth/login", {
          email: this.form.email,
          password: this.form.password,
          store_slug: this.$route.params.wilcard,
        })
        .then((res) => {
          if (res.data.meta.status) {
            // console.log(res);
            location.reload();
          } else {
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
              title: `${res.data.meta.message}`,
            });
          }
        });
    },
    checkLogout() {
      // if (localStorage.getItem("access_token") === null) {
      //   this.$router.push(`/${this.$route.params.wilcard}/`);
      //   const Toast = this.$swal.mixin({
      //     toast: true,
      //     position: "bottom-end",
      //     showConfirmButton: false,
      //     timer: 3000,
      //     timerProgressBar: true,
      //     didOpen: (toast) => {
      //       toast.addEventListener("mouseenter", this.$swal.stopTimer);
      //       toast.addEventListener("mouseleave", this.$swal.resumeTimer);
      //     },
      //     popup: "swal2-show",
      //     backdrop: "swal2-backdrop-show",
      //     icon: "swal2-icon-show",
      //   });
      //   Toast.fire({
      //     icon: "success",
      //     title: "cannot access this route",
      //   });
      //   this.$router.push(`/${this.$route.params.wilcard}/`);
      // }
      if (
        localStorage.getItem("store_id") !==
        localStorage.getItem("store_customer_id")
      ) {
        localStorage.removeItem("access_token");
        localStorage.setItem("isLogin", false);
        if (localStorage.getItem("isLogin") === "false") {
          // location.reload();
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
    me() {
      axios.defaults.headers.common.Authorization =
        "Bearer " + localStorage.getItem("access_token");
      axios.defaults.baseURL = process.env.VUE_APP_API_URL;
      const params = {
        entities: "Store",
      };
      axios.get("customer/me", { params: params }).then((res) => {
        if (res.data.meta.status) {
          console.log(res);
          localStorage.setItem("customer", JSON.stringify(res.data.data));
          if (localStorage.getItem("isLogin") === "true") {
            console.log(
              localStorage.getItem("store_customer_id") !==
                localStorage.getItem("store_id")
            );
            console.log(
              "store_customer_id",
              localStorage.getItem("store_customer_id")
            );
            console.log("store_jd", localStorage.getItem("store_id"));
            if (
              localStorage.getItem("store_customer_id") !=
              localStorage.getItem("store_id")
            ) {
              console.log("tidak cocok");

              localStorage.removeItem("access_token");
              localStorage.setItem("isLogin", false);
              location.reload();
            }
            localStorage.setItem("wild_route", res.data.data.store.Slug);
            this.$router.push(`/${localStorage.getItem("wild_route")}/`);
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
            title: `Selamat Datang ${res.data.data.username}`,
          });
          this.customer = res.data.data;
          console.log("Customer =>", this.customer);
          localStorage.setItem("store_customer_id", res.data.data.store_id);
        }
      });
    },
  },
};
</script>

<style></style>
