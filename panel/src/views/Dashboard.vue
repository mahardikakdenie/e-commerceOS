<template>
  <div class="home">
    <Navbar />
    <Product />
  </div>
</template>

<script>
// @ is an alias to /src
// import HelloWorld from "@/components/HelloWorld.vue";
import Navbar from "@/components/Navbar.vue";
import Product from "@/components/banner/Product.vue";

export default {
  name: "Home",
  components: {
    Product,
    Navbar,
  },
};
</script>
