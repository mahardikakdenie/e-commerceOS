<template>
  <div>
    <List :propsCart="computedCart" />
  </div>
</template>

<script>
export default {
  components: {
    List: () => import("../components/chart/List.vue"),
  },
  computed: {
    computedCart() {
      return this.$store.state.cart.data;
    },
  },
  mounted() {
    this.getDataCart();
  },
  methods: {
    getDataCart() {
      this.$store
        .dispatch("cart/getDataCart", {
          entities: "Product,Store,Customer",
        })
        .then((res) => {
          if (res.data.meta.status) {
            console.log("data => ", res);
          }
        });
    },
  },
};
</script>

<style></style>
