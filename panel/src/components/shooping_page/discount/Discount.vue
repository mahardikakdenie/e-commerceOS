<template>
  <section id="discount" class="discount">
    <div class="container">
      <div class="row">
        <div class="col-12 mobile--title d-flex justify-content-start mb-2">
          <h2>Best Seller</h2>
        </div>
        <div
          class="hero-items--discount col-lg-3 col-sm-12 mobile-col"
          :class="`${i == 1 ? 'mobile-col-one' : ''}`"
          v-for="(item, i) in best_seller"
          :key="i"
        >
          <div class="card">
            <img
              v-if="item.link"
              :src="item.link.image"
              class="card-img-top"
              alt="..."
            />
            <img
              v-else
              src="https://images.unsplash.com/photo-1572454591674-2739f30d8c40?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MXx8Zmxvd2VyJTIwYm91cXVldHxlbnwwfHwwfHw%3D&auto=format&fit=crop&w=500&q=60"
              class="card-img-top"
              alt="..."
            />
            <div class="card-body d-flex justify-content-start">
              <h3 class="card-title">{{ item.name.substr(0, 40) + " .. " }}</h3>
            </div>
            <div class="card-price">
              <p class="ml-2 text-current-price">Rp.{{ item.price }}</p>
              <!-- <p class="text-price">Rp.{{ item.price }}</p> -->
            </div>
            <div class="card-price">
              <p class="ml-2 text-current-price">Stock : {{ item.stock }}</p>
              <!-- <p class="text-price">Rp.{{ item.price }}</p> -->
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
            v-for="(item, i) in best_seller"
            :key="i"
          >
            <div class="card card--mobile">
              <img
                @click="$router.push(`/product/${i}`)"
                src="https://images.unsplash.com/photo-1572454591674-2739f30d8c40?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MXx8Zmxvd2VyJTIwYm91cXVldHxlbnwwfHwwfHw%3D&auto=format&fit=crop&w=500&q=60"
                class="img--product"
                alt="..."
              />
              <div class="card-body d-flex justify-content-start">
                <h3 class="card-title">Title Product</h3>
              </div>
              <div class="card-price">
                <p class="ml-2 text-current-price">Rp.20000</p>
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
  props: {
    best_seller: {
      type: Array,
      required: true,
    },
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

<style>
.card .card-body .card-title {
  font-size: 18px;
  text-align: justify;
  text-justify: inter-word;
}
</style>
