<template>
  <!-- Hero Section Begin -->
  <section class="hero-section">
    <!-- <div class=" owl-carousel"></div> -->
    <carousel
      class="hero-items"
      :dots="true"
      :items="1"
      :nav="false"
      :autoplay="true"
    >
      <!-- <template slot="prev"><span class="prev">prev</span></template> -->
      <div
        class="single-hero-items set-bg"
        :style="`background-image: url('https://images.unsplash.com/photo-1517487881594-2787fef5ebf7?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MTd8fHdlYnNpdGUlMjBjbG90aGVzfGVufDB8fDB8fA%3D%3D&auto=format&fit=crop&w=500&q=60');opcity:0.2`"
      >
        <div class="container">
          <div class="row d-flex">
            <div class="col-lg-12 justify-content-center">
              <span>Bag,kids</span>
              <h1 style="color: #424233" class="">
                Himatif Website Store Clothes
              </h1>
              <p>
                Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do
                eiusmod tempor incididunt ut labore et dolore
              </p>
              <a href="#" class="primary-btn">Shop Now</a>
            </div>
          </div>
        </div>
      </div>
      <div
        class="single-hero-items set-bg"
        :style="`background-image: url('https://images.unsplash.com/photo-1527176930608-09cb256ab504?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MTd8fGJhY2tncm91bmQlMjBzdG9yZXxlbnwwfHwwfHw%3D&auto=format&fit=crop&w=500&q=60')`"
      >
        <div class="container">
          <div class="row">
            <div class="col-lg-5">
              <span>Bag,kids</span>
              <h1>Black friday</h1>
              <p>
                Lorem ipsum dolor sit amet, consectetur adipisicing elit, sed do
                eiusmod tempor incididunt ut labore et dolore
              </p>
              <a href="#" class="primary-btn">Shop Now</a>
            </div>
          </div>
        </div>
      </div>

      <!-- <template slot="next"><span class="next">next</span></template> -->
    </carousel>
  </section>
  <!-- Hero Section End -->
</template>

<script>
import carousel from "vue-owl-carousel";
import hero2 from "../../../assets/images/hero-2.jpg";
import hero1 from "../../../assets/images/hero-1.jpg";
export default {
  components: {
    carousel,
  },
  data: () => ({
    hero1: hero1,
    hero2: hero2,
  }),
};
</script>
