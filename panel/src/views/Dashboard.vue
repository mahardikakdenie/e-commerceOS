<template>
  <div class="home">
    <Product />
    <List />
    <Slider />
    <List />
  </div>
</template>

<script>
// @ is an alias to /src
// import HelloWorld from "@/components/HelloWorld.vue";
// import Navbar from "@/components/Navbar.vue";
import Product from "@/components/product/Product.vue";
import List from "@/components/product/list/List.vue";
import Slider from "@/components/product/slider/Slider.vue";

export default {
  name: "Home",
  components: {
    Product,
    Slider,
    List,
  },
};
</script>
