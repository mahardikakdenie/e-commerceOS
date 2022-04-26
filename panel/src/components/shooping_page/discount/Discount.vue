<template>
  <section id="discount" class="discount">
    <div class="container">
      <div class="row">
        <div class="col-12 mobile--title d-flex justify-content-start mb-2">
          <h2>Discount Products</h2>
        </div>
        <div
          class="hero-items--discount col-lg-3 col-sm-12 mobile-col"
          :class="`${i == 1 ? 'mobile-col-one' : ''}`"
          v-for="i in 4"
          :key="i"
        >
          <div class="card">
            <img
              src="https://images.unsplash.com/photo-1572454591674-2739f30d8c40?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MXx8Zmxvd2VyJTIwYm91cXVldHxlbnwwfHwwfHw%3D&auto=format&fit=crop&w=500&q=60"
              class="card-img-top"
              alt="..."
            />
            <div class="card-body d-flex justify-content-start">
              <h3 class="card-title">Title Product</h3>
            </div>
            <div class="card-price">
              <p class="text-price">Rp.200000</p>
            </div>
            <div class="card-action">
              <button class="btn btn--cart">Add To Cart</button>
            </div>
          </div>
        </div>
        <!-- Mobile Version -->
        <carousel
          class="hero-items--discount-mobile"
          :dots="true"
          :items="1"
          :nav="false"
          :autoplay="true"
        >
          <div
            class="col-lg-3 col-sm-12 mobile-col"
            :class="`${i == 1 ? 'mobile-col-one' : ''}`"
            v-for="i in 4"
            :key="i"
          >
            <div class="card card--mobile">
              <img
                src="https://images.unsplash.com/photo-1572454591674-2739f30d8c40?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MXx8Zmxvd2VyJTIwYm91cXVldHxlbnwwfHwwfHw%3D&auto=format&fit=crop&w=500&q=60"
                class="img--product"
                alt="..."
              />
              <div class="card-body d-flex justify-content-start">
                <h3 class="card-title">Title Product</h3>
              </div>
              <div class="card-start">
                <img :src="require('../../../assets/images/star.png')" alt="" />
                <span class="text">4.5</span>
              </div>
              <div class="card-price">
                <p class="text-price">Rp.200000</p>
              </div>
              <div class="card-action">
                <button class="btn btn--cart">Add To Cart</button>
              </div>
            </div>
          </div>
        </carousel>
      </div>
    </div>
  </section>
</template>

<script>
import carousel from "vue-owl-carousel";
export default {
  components: {
    carousel,
  },
  data: () => ({
    innerHeight: 0,
    is_mobile: false,
  }),
  mounted() {
    this.innerHeight = window.innerWidth;
    if (this.innerHeight < 400) this.is_mobile = true;
    else return (this.is_mobile = false);
    console.log("Mobile :", this.is_mobile);
  },
};
</script>

<style></style>
