<template>
  <div></div>
</template>

<script>
export default {
  components: {},
  mounted() {
    this.$router.push(`/${localStorage.getItem("wild_route")}/`);
  },
};
</script>

<style></style>
