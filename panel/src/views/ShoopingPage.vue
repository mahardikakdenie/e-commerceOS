<template>
  <div class="shopping-page">
    <Product title="Product" />
    <Discount :best_seller="best_seller" />
    <!-- <ListProduct title="T-Shirt" :is_view="true" />
    <ListProduct title="Bouqet" :is_view="true" /> -->
    <!-- <ListProductAll title="Products T-shirt" :is_view="true" /> -->
    <ListProductAll title="Products" :data="data" :is_view="true" />
  </div>
</template>

<script>
import Product from "@/components/product/Banner.vue";
import Discount from "@/components/shooping_page/discount/Discount.vue";
// import ListProduct from "@/components/shooping_page/list/List.vue";
import ListProductAll from "@/components/shooping_page/list/ListAll";
export default {
  components: {
    Product,
    Discount,
    // ListProduct,
    ListProductAll,
  },
  data: () => ({
    data: [],
    best_seller: [],
  }),
  mounted() {
    this.getDataProduct();
    this.bestSeller();
  },
  methods: {
    getDataProduct() {
      this.$store
        .dispatch("all_product/getDataProduct", {
          entities: "Category,Store,User",
        })
        .then((res) => {
          if (res.data.meta.status) {
            this.data = res.data.data;
          }
        });
    },
    bestSeller() {
      this.$store
        .dispatch("all_product/getDataProduct", {
          entities: "Category,Store,User",
          is_seller: "1",
          page: 1,
          per_page: 4,
          store_id: localStorage.getItem("store_id"),
        })
        .then((res) => {
          if (res.data.meta.status) {
            this.best_seller = res.data.data;
            console.log("data all shopping pages Best => ", res);
          }
        });
    },
  },
};
</script>

<style></style>
