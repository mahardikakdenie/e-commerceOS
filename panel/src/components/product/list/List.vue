<template>
  <div class="container product--list">
    <div
      :class="`row ${(i + 1) % 2 === 0 ? 'd-flex flex-row-reverse' : ''}`"
      v-for="(item, i) in list"
      :key="i"
    >
      <div class="col-lg-6">
        <div class="card">
          <img
            v-if="item.link"
            :src="item.link.image"
            width="100%"
            height="543"
            class="card-img-top img--brand"
            alt="..."
          />
        </div>
      </div>
      <div class="col-lg-6">
        <div class="__brand">
          <div class="title--brand">
            <h1 class="--brand" @click="$router.push(`product/${item.id}`)">
              {{ item.name }}
            </h1>
          </div>
          <div class="paragrap--brand">
            <p>
              {{ item.description }}
            </p>
          </div>
          <div class="price--brand">
            <h3 class="font-bold">Rp.{{ item.price }}</h3>
          </div>
          <div class="action--brand">
            <button class="btn action--brand__buy">Buy</button>
            <button class="btn action--brand__add--cart">Add To Cart</button>
          </div>
        </div>
      </div>
    </div>
    <div class="button--loadmore">
      <button
        class="btn primary-btn btn-lg"
        @click="$router.push('/shopping-page')"
      >
        More {{ title }}
      </button>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    list: {
      type: Array,
      default: null,
    },
    title: {
      type: String,
      default: "",
    },
  },
  methods: {},
};
</script>

<style scoped>
.--brand {
  text-shadow: 5px 5px 10px #e7ab3c;
  font-weight: bold;
  cursor: pointer;
  transition: 1s;
}
.--brand:hover {
  color: #e7ab3c;
  font-weight: bold;
  cursor: pointer;
  transform: scale(0.9);
}
</style>
