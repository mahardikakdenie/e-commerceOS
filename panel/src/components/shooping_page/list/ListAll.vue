<template>
  <div class="container list-product">
    <div class="row">
      <div class="col-12 d-flex justify-content-start">
        <h3 v-if="is_view">{{ title }}</h3>
      </div>

      <div class="col-lg-3 p-2" v-for="(item, i) in data" :key="i">
        <div class="card">
          <img
            v-if="item.link"
            :src="item.link.image"
            class="card-img-top img--size"
          />
          <img
            v-else
            src="https://images.unsplash.com/photo-1572454591674-2739f30d8c40?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MXx8Zmxvd2VyJTIwYm91cXVldHxlbnwwfHwwfHw%3D&auto=format&fit=crop&w=500&q=60"
            class="card-img-top"
          />
          <div class="card-body d-flex justify-content-start">
            <h3 class="card-title">{{ item.name.substr(0, 40) + " .. " }}</h3>
          </div>

          <div class="card-price">
            <p class="text-price">Rp.{{ item.price }}</p>
          </div>
          <div class="card-action">
            <button class="btn btn--list-product">Add To Cart</button>
          </div>
        </div>
      </div>
    </div>
    <div class="row">
      <div class="col-12 d-flex justify-content-center">
        <button class="btn primary-btn">Load More</button>
      </div>
    </div>
  </div>
</template>

<script>
// import corousel from "vue-owl-carousel";
export default {
  props: {
    title: {
      type: String,
      default: "",
    },
    data: {
      type: Array,
      default: null,
    },
    is_view: {
      type: Boolean,
      default: false,
    },
  },
  components: {},
  data: () => ({}),
};
</script>

<style>
.img--size {
  transition: 1s;
  cursor: pointer;
}
.img--size:hover {
  transform: scale(0.9);
}
</style>
