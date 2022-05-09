<template>
  <nav class="navbar navbar-expand-lg navbar-light bg-grey">
    <div class="container">
      <h3 class="text-black cursor-pointer text--brand font-bold" href="#">
        Himatif
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
          <li
            :class="`nav-item ${
              $route.path === '/cart' ? 'active' : ''
            } icon--cart cursor-pointer`"
          >
            <svg
              @click="$router.push('/cart')"
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
            <span class="badge--color">
              3
              <span class="visually-hidden">unread messages</span>
            </span>
          </li>

          <li class="nav-item">
            <button
              class="btn btn--login--width"
              @click="$router.push('/login')"
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
  data: () => ({
    active: "Home",
    itemNavbar: [
      {
        text: "Home",
        to: "/",
      },
      {
        text: "Product",
        to: "/shopping-page",
      },
      {
        text: "Contact",
        to: "/contact",
      },
    ],
  }),
  mounted() {},
  methods: {
    setNav(item) {
      this.active = item.text;
      // this.$router.push(item.to);
    },
  },
};
</script>

<style></style>
