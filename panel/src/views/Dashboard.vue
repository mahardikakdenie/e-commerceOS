<template>
  <div class="home">
    <Product title="Himatif Store Website" />
    <!-- <div v-for="(item, i) in computedCategories" :key="i"> -->
    <!-- <Slider v-if="i !== 0" /> -->
    <List :list="computedProduct" title="Product" />
    <!-- </div> -->
  </div>
</template>

<script>
import Product from "@/components/product/Banner.vue";
import List from "@/components/product/list/List.vue";
// import Slider from "@/components/product/slider/Slider.vue";

export default {
  name: "Home",
  components: {
    Product,
    // Slider,
    List,
  },
  data: () => ({
    dataClothes: [
      {
        url: "https://images.unsplash.com/photo-1467043237213-65f2da53396f?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8N3x8Y2xvdGhlc3xlbnwwfHwwfHw%3D&auto=format&fit=crop&w=500&q=60",
      },
      {
        url: "https://images.unsplash.com/photo-1434389677669-e08b4cac3105?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8Nnx8Y2xvdGhlc3xlbnwwfHwwfHw%3D&auto=format&fit=crop&w=500&q=60",
      },
    ],
    dataBucket: [
      {
        url: "https://images.unsplash.com/photo-1572454591674-2739f30d8c40?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MXx8Zmxvd2VyJTIwYm91cXVldHxlbnwwfHwwfHw%3D&auto=format&fit=crop&w=500&q=60",
      },
      {
        url: "https://images.unsplash.com/photo-1565695951564-007d8f297e48?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8NXx8Zmxvd2VyJTIwYm91cXVldHxlbnwwfHwwfHw%3D&auto=format&fit=crop&w=500&q=60",
      },
    ],
  }),
  computed: {
    computedCategories() {
      return this.$store.state.category.categories;
    },
    computedProduct() {
      return this.$store.state.product.views;
    },
  },
  mounted() {
    this.getCategories();
    this.getBestSeller();
    console.log("ini Dashboard");
  },
  methods: {
    getCategories() {
      this.$store
        .dispatch("category/getCategories", {
          // entities: "User,Store",
          store_id: localStorage.getItem("store_id"),
        })
        .then((res) => {
          console.log(res);
          if (res.data.meta.status) {
            console.log(res);
          }
        });
    },
    getBestSeller() {
      this.$store
        .dispatch("product/getBestSeller", {
          page: 1,
          per_page: 2,
          entities: "Category",
          store_id: localStorage.getItem("store_id"),
          is_seller: 1,
        })
        .then((res) => {
          console.log("ini best seller => ", res.data.data);
        });
    },
  },
};
</script>
