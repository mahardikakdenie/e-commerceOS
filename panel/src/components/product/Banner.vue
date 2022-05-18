<template>
  <carousel
    class="hero-items"
    :dots="true"
    :items="1"
    :nav="false"
    :autoplay="true"
  >
    <div class="product overflow-hidden">
      <div class="container product__banner">
        <div class="row row-cols-2">
          <div class="col-lg-7"></div>
          <div class="col-lg-5">
            <div class="p-3 text-end">
              <div class="text--product title--product-sm">
                {{ title }}
              </div>
              <p class="paragrap--product my-6 title--product-sm">
                Lorem ipsum, dolor sit amet consectetur adipisicing elit.
                Doloremque debitis delectus cupiditate tenetur ipsam minus
              </p>
              <button class="btn btn--product btn--product-sm btn-lg">
                Explore
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
    <div class="product overflow-hidden">
      <div class="container product__banner">
        <div class="row row-cols-2">
          <div class="col-lg-7"></div>
          <div class="col-lg-5">
            <div class="p-3 text-end">
              <div class="text--product title--product-sm">
                {{ title }}
              </div>
              <p class="paragrap--product my-6 title--product-sm">
                Lorem ipsum, dolor sit amet consectetur adipisicing elit.
                Doloremque debitis delectus cupiditate tenetur ipsam minus
              </p>
              <button class="btn btn--product btn--product-sm btn-lg">
                Explore
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
  </carousel>
</template>

<script>
import carousel from "vue-owl-carousel";
export default {
  components: {
    carousel,
  },
  props: {
    title: {
      type: String,
      default: "",
    },
  },
};
</script>

<style></style>
