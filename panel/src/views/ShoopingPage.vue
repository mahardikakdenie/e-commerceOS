<template>
  <div class="shopping-page">
    <Product />
    <Discount />
    <ListProduct title="T-Shirt" :is_view="true" />
    <ListProduct title="Bouqet" :is_view="true" />
    <ListProductAll title="Products T-shirt" :is_view="true" />
    <ListProductAll title="Products Bouquet" :is_view="true" />
  </div>
</template>

<script>
import Product from "@/components/product/Product.vue";
import Discount from "@/components/shooping_page/discount/Discount.vue";
import ListProduct from "@/components/shooping_page/list/List.vue";
import ListProductAll from "@/components/shooping_page/list/ListAll";
export default {
  components: {
    Product,
    Discount,
    ListProduct,
    ListProductAll,
  },
};
</script>

<style></style>
