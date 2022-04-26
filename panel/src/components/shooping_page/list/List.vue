<template>
  <div class="container list-product">
    <div class="row">
      <div class="col-12 d-flex justify-content-start">
        <h3 v-if="is_view">{{ title }}</h3>
      </div>
      <!-- Destop Version -->
      <corousel
        class="hero-items-dekstop"
        :dots="true"
        :items="4"
        :nav="false"
        :autoplay="true"
      >
        <div class="col-lg-12 p-2" v-for="i in 8" :key="i">
          <div class="card">
            <img
              src="https://images.unsplash.com/photo-1572454591674-2739f30d8c40?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MXx8Zmxvd2VyJTIwYm91cXVldHxlbnwwfHwwfHw%3D&auto=format&fit=crop&w=500&q=60"
              class="card-img-top"
            />
            <div class="card-body d-flex justify-content-start">
              <h3 class="card-title">Title Product</h3>
            </div>

            <div class="card-price">
              <p class="text-price">Rp.200000</p>
            </div>
            <div class="card-action">
              <button class="btn btn--list-product">Add To Cart</button>
            </div>
          </div>
        </div>
      </corousel>
      <!-- Mobile Version -->
      <corousel
        class="hero-items-mobile"
        :dots="true"
        :items="1"
        :nav="false"
        :autoplay="true"
      >
        <div class="col-lg-12 p-2" v-for="i in 8" :key="i">
          <div class="card">
            <img
              src="https://images.unsplash.com/photo-1572454591674-2739f30d8c40?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MXx8Zmxvd2VyJTIwYm91cXVldHxlbnwwfHwwfHw%3D&auto=format&fit=crop&w=500&q=60"
              class="card-img-top"
            />
            <div class="card-body d-flex justify-content-start">
              <h3 class="card-title">Title Product</h3>
            </div>

            <div class="card-price">
              <p class="text-price">Rp.200000</p>
            </div>
            <div class="card-action">
              <button class="btn btn--list-product">Add To Cart</button>
            </div>
          </div>
        </div>
      </corousel>
    </div>
  </div>
</template>

<script>
import corousel from "vue-owl-carousel";
export default {
  props: {
    title: {
      type: String,
      default: "",
    },
    is_view: {
      type: Boolean,
      default: false,
    },
  },
  components: {
    corousel,
  },
  data: () => ({}),
};
</script>

<style></style>
