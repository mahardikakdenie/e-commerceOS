<template>
  <nav class="navbar navbar-expand-lg navbar-light bg-grey">
    <div class="container">
      <h3 class="text-black cursor-pointer text--brand font-bold" href="#">
        {{ title }}
      </h3>
      <button
        class="navbar-toggler"
        type="button"
        data-bs-toggle="collapse"
        data-bs-target="#navbarSupportedContent"
        aria-controls="navbarSupportedContent"
        aria-expanded="false"
        aria-label="Toggle navigation"
      >
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarSupportedContent">
        <ul class="ms-auto navbar-nav">
          <li v-for="(item, i) in itemNavbar" :key="i" class="nav-item">
            <router-link
              :class="`nav-link ${$route.path === item.to ? 'active' : ''}`"
              aria-current="page"
              :to="item.to"
              @click="setNav(item)"
              >{{ item.text }}</router-link
            >
          </li>
          <li class="nav-item dropdown">
            <a
              class="nav-link dropdown-toggle"
              href="#"
              id="navbarScrollingDropdown"
              role="button"
              data-bs-toggle="dropdown"
              aria-expanded="false"
            >
              Product
            </a>
            <ul class="dropdown-menu" aria-labelledby="navbarScrollingDropdown">
              <li v-for="(item, i) in propsCategory" :key="i">
                <a class="dropdown-item" href="#">{{ item.name }}</a>
              </li>
              <li><hr class="dropdown-divider" /></li>
              <li
                :class="` ${
                  $route.path === `/${store_route}/shopping-page`
                    ? 'bg-primary'
                    : ''
                }`"
              >
                <a
                  @click="toRoute('shopping-page')"
                  :class="`dropdown-item dropdown-item-all`"
                  href="#"
                >
                  Semua Product</a
                >
              </li>
            </ul>
          </li>
          <!-- <li
            :class="`nav-item ${
              $route.path === `/${$route.params.wilcard}/cart` ? 'active' : ''
            } icon--cart cursor-pointer`"
          ></li> -->

          <li class="nav-item dropdown">
            <div class="is--login" v-if="isLogin === 'true'">
              <img
                src="https://upload.wikimedia.org/wikipedia/commons/thumb/d/d3/Microsoft_Account.svg/512px-Microsoft_Account.svg.png?20170218203212"
                alt=""
                class="img--user"
                srcset=""
              />
              <span
                class="nick--customer dropdown-toggle"
                id="dropdownMenuLink"
                data-bs-toggle="dropdown"
                aria-expanded="false"
                >{{ propsUser.email }}</span
              >
              <span class="nick--customer mobile--user">{{
                propsUser.email
              }}</span>
              <ul class="dropdown-menu" aria-labelledby="dropdownMenuLink">
                <li>
                  <a class="dropdown-item" href="#"
                    ><svg
                      xmlns="http://www.w3.org/2000/svg"
                      width="16"
                      height="16"
                      fill="currentColor"
                      class="bi bi-patch-check-fill"
                      viewBox="0 0 16 16"
                    >
                      <path
                        d="M10.067.87a2.89 2.89 0 0 0-4.134 0l-.622.638-.89-.011a2.89 2.89 0 0 0-2.924 2.924l.01.89-.636.622a2.89 2.89 0 0 0 0 4.134l.637.622-.011.89a2.89 2.89 0 0 0 2.924 2.924l.89-.01.622.636a2.89 2.89 0 0 0 4.134 0l.622-.637.89.011a2.89 2.89 0 0 0 2.924-2.924l-.01-.89.636-.622a2.89 2.89 0 0 0 0-4.134l-.637-.622.011-.89a2.89 2.89 0 0 0-2.924-2.924l-.89.01-.622-.636zm.287 5.984-3 3a.5.5 0 0 1-.708 0l-1.5-1.5a.5.5 0 1 1 .708-.708L7 8.793l2.646-2.647a.5.5 0 0 1 .708.708z"
                      />
                    </svg>
                    Pembayaran</a
                  >
                </li>
                <li>
                  <a @click="toRoute('cart')" class="dropdown-item" href="#">
                    <svg
                      xmlns="http://www.w3.org/2000/svg"
                      width="20"
                      height="20"
                      fill="currentColor"
                      class="bi bi-cart-check"
                      viewBox="0 0 16 16"
                    >
                      <path
                        d="M11.354 6.354a.5.5 0 0 0-.708-.708L8 8.293 6.854 7.146a.5.5 0 1 0-.708.708l1.5 1.5a.5.5 0 0 0 .708 0l3-3z"
                      />
                      <path
                        d="M.5 1a.5.5 0 0 0 0 1h1.11l.401 1.607 1.498 7.985A.5.5 0 0 0 4 12h1a2 2 0 1 0 0 4 2 2 0 0 0 0-4h7a2 2 0 1 0 0 4 2 2 0 0 0 0-4h1a.5.5 0 0 0 .491-.408l1.5-8A.5.5 0 0 0 14.5 3H2.89l-.405-1.621A.5.5 0 0 0 2 1H.5zm3.915 10L3.102 4h10.796l-1.313 7h-8.17zM6 14a1 1 0 1 1-2 0 1 1 0 0 1 2 0zm7 0a1 1 0 1 1-2 0 1 1 0 0 1 2 0z"
                      />
                    </svg>
                    <!-- <span class="badge--color">
                      3
                      <span class="visually-hidden">unread messages</span>
                    </span> -->
                    Cart
                    <span class="S badge badge--dekstop rounded-pill bg-danger"
                      >4</span
                    ></a
                  >
                </li>
                <li>
                  <a @click="logout" class="dropdown-item" href="#"
                    ><svg
                      xmlns="http://www.w3.org/2000/svg"
                      width="16"
                      height="16"
                      fill="currentColor"
                      class="bi bi-arrow-right-square-fill"
                      viewBox="0 0 16 16"
                    >
                      <path
                        d="M0 14a2 2 0 0 0 2 2h12a2 2 0 0 0 2-2V2a2 2 0 0 0-2-2H2a2 2 0 0 0-2 2v12zm4.5-6.5h5.793L8.146 5.354a.5.5 0 1 1 .708-.708l3 3a.5.5 0 0 1 0 .708l-3 3a.5.5 0 0 1-.708-.708L10.293 8.5H4.5a.5.5 0 0 1 0-1z"
                      />
                    </svg>
                    Logout</a
                  >
                </li>
              </ul>
            </div>
            <button
              v-else
              class="btn btn--login--width"
              @click="toRoute('login')"
              type="submit"
            >
              login
            </button>
          </li>
        </ul>
      </div>
    </div>
  </nav>
</template>

<script>
export default {
  props: {
    propsUser: {
      type: Object,
      default: null,
    },
    propsCategory: {
      type: Array,
      default: null,
    },
  },
  data: () => ({
    active: "Home",
    store_route: localStorage.getItem("wild_route"),
    itemNavbar: [
      {
        text: "Home",
        to: `/${localStorage.getItem("wild_route")}/`,
      },
      // {
      //   text: "Tentang",
      //   to: `/${localStorage.getItem("wild_route")}/shopping-page`,
      // },
      {
        text: "Contact",
        to: `/${localStorage.getItem("wild_route")}/contact`,
      },
    ],
    isLogin: localStorage.getItem("isLogin") || false,
    title: localStorage.getItem("name_store"),
    isDown: false,
    // icon: ``
  }),
  mounted() {
    console.log("wilcard :", this.$route.path);
  },
  methods: {
    setNav(item) {
      this.active = item.text;
      // this.$router.push(item.to);
    },
    toRoute(name) {
      this.$router.push(`/${localStorage.getItem("wild_route")}/${name}`);
    },
    logout() {
      this.$emit("logout");
    },
  },
};
</script>

<style>
.img--user {
  width: 30px;
  margin: -3px 5px 0 8px;
}
.is--login {
  margin: 9px 0 0 0;
}
.nick--customer {
  font-weight: bold;
  cursor: pointer;
}
.nick--customer:hover {
  color: #fff;
}
.dropdown-menu {
  left: 70px !important;
  min-width: 12rem;
  padding: 10px 2px 10px 1px;
}
.dropdown-item-all.active {
  background-color: red;
}
@media screen and (max-width: 400px) {
  .dropdown-menu {
    background: #a0a0a0;
    border: none;
    display: flex;
    justify-content: center;
  }
  .dropdown-toggle::after {
    display: none;
  }
  #dropdownMenuLink {
    display: none;
  }
}
@media screen and (min-width: 400px) {
  .mobile--user {
    display: none;
  }
  /* .badge--dekstop {
    margin: 2px 0 0 100px;
  } */
}
</style>
