<template>
  <div>
    <List />
  </div>
</template>

<script>
export default {
  components: {
    List: () => import("../components/chart/List.vue"),
  },
};
</script>

<style></style>
