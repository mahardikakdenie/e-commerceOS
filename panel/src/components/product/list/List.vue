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
            :src="item.url"
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
            <h1>Title Baju</h1>
          </div>
          <div class="paragrap--brand">
            <p>
              Lorem ipsum dolor sit amet consectetur adipisicing elit. Fugit
              officia commodi totam aperiam optio corrupti nesciunt, est at
              natus odio velit id excepturi ipsam maxime aut voluptate eius sed
              ea quidem molestias qui soluta doloribus! Dolorem cum reiciendis
              eum ratione.
            </p>
          </div>
          <div class="price--brand">
            <h3 class="font-bold">Rp.20000</h3>
          </div>
          <div class="action--brand">
            <button class="btn action--brand__buy btn-secondary">Buy</button>
            <button class="btn btn-secondary action--brand__add--cart">
              Add To Cart
            </button>
          </div>
        </div>
      </div>
    </div>
  </div>
</template>

<script>
export default {
  data: () => ({
    list: [
      {
        url: "https://images.unsplash.com/photo-1467043237213-65f2da53396f?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8N3x8Y2xvdGhlc3xlbnwwfHwwfHw%3D&auto=format&fit=crop&w=500&q=60",
      },
      {
        url: "https://images.unsplash.com/photo-1434389677669-e08b4cac3105?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8Nnx8Y2xvdGhlc3xlbnwwfHwwfHw%3D&auto=format&fit=crop&w=500&q=60",
      },
    ],
  }),
};
</script>

<style></style>
