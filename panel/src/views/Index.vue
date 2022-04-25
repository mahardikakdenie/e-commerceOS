<template>
  <!-- <div class="home"> -->
  <div>
    <Navbar />
    <router-view></router-view>
    <Footer />
  </div>
  <!-- </div> -->
</template>

<script>
import Navbar from "@/components/Navbar.vue";
import Footer from "@/components/footer/Footer.vue";
export default {
  components: {
    Footer,
    Navbar,
  },
};
</script>

<style></style>
