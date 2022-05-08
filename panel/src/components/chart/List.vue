<template>
  <div class="container cart">
    <!-- <div class="row">
      <div class="col"> -->
    <div class="card mb-3" v-for="(v, i) in list" :key="i" style="">
      <div class="row g-0">
        <div class="col-md-2">
          <img
            :src="v.image"
            class="img-fluid rounded-start controll--img"
            alt="..."
          />
        </div>
        <div class="col-md-10">
          <div class="card-body">
            <h4 class="card-title">
              {{ v.title }}
            </h4>
            <div>
              <select
                v-model="v.variant"
                class="form-select"
                aria-label="Default select example"
              >
                <option selected>Variant</option>
                <option value="1">One</option>
                <option value="2">Two</option>
                <option value="3">Three</option>
              </select>
            </div>
            <div class="card--info-action">
              <div class="card--info-action-price">
                <p class="ml-2 text-current-price">
                  {{ v.price }}, x {{ v.mount }} || Rp.{{ totalPrice(i) }}
                </p>
              </div>
            </div>
            <div class="d-flex action--cart">
              <button @click="methodPlus(i)" class="btn btn-primary btn--">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  width="16"
                  height="16"
                  fill="currentColor"
                  class="bi bi-bookmark-plus"
                  viewBox="0 0 16 16"
                >
                  <path
                    d="M2 2a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v13.5a.5.5 0 0 1-.777.416L8 13.101l-5.223 2.815A.5.5 0 0 1 2 15.5V2zm2-1a1 1 0 0 0-1 1v12.566l4.723-2.482a.5.5 0 0 1 .554 0L13 14.566V2a1 1 0 0 0-1-1H4z"
                  />
                  <path
                    d="M8 4a.5.5 0 0 1 .5.5V6H10a.5.5 0 0 1 0 1H8.5v1.5a.5.5 0 0 1-1 0V7H6a.5.5 0 0 1 0-1h1.5V4.5A.5.5 0 0 1 8 4z"
                  />
                </svg>
              </button>
              <button @click="minusMethod(i)" class="btn btn-primary btn--">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  width="16"
                  height="16"
                  fill="currentColor"
                  class="bi bi-bookmark-dash"
                  viewBox="0 0 16 16"
                >
                  <path
                    fill-rule="evenodd"
                    d="M5.5 6.5A.5.5 0 0 1 6 6h4a.5.5 0 0 1 0 1H6a.5.5 0 0 1-.5-.5z"
                  />
                  <path
                    d="M2 2a2 2 0 0 1 2-2h8a2 2 0 0 1 2 2v13.5a.5.5 0 0 1-.777.416L8 13.101l-5.223 2.815A.5.5 0 0 1 2 15.5V2zm2-1a1 1 0 0 0-1 1v12.566l4.723-2.482a.5.5 0 0 1 .554 0L13 14.566V2a1 1 0 0 0-1-1H4z"
                  />
                </svg>
              </button>
              <button @click="deleteMethod(i)" class="btn btn-danger">
                <svg
                  xmlns="http://www.w3.org/2000/svg"
                  width="16"
                  height="16"
                  fill="currentColor"
                  class="bi bi-trash"
                  viewBox="0 0 16 16"
                >
                  <path
                    d="M5.5 5.5A.5.5 0 0 1 6 6v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm2.5 0a.5.5 0 0 1 .5.5v6a.5.5 0 0 1-1 0V6a.5.5 0 0 1 .5-.5zm3 .5a.5.5 0 0 0-1 0v6a.5.5 0 0 0 1 0V6z"
                  />
                  <path
                    fill-rule="evenodd"
                    d="M14.5 3a1 1 0 0 1-1 1H13v9a2 2 0 0 1-2 2H5a2 2 0 0 1-2-2V4h-.5a1 1 0 0 1-1-1V2a1 1 0 0 1 1-1H6a1 1 0 0 1 1-1h2a1 1 0 0 1 1 1h3.5a1 1 0 0 1 1 1v1zM4.118 4 4 4.059V13a1 1 0 0 0 1 1h6a1 1 0 0 0 1-1V4.059L11.882 4H4.118zM2.5 3V2h11v1h-11z"
                  />
                </svg>
              </button>
            </div>
          </div>
        </div>
      </div>
    </div>
    <!-- </div>
    </div> -->
  </div>
</template>

<script>
export default {
  data: () => ({
    mount_price: 0,
    mount: 3,
    list: [
      {
        title: "Jersey RRQ HOSHI",
        variant: "1",
        price: 120000,
        mount: 1,
        image:
          "https://images.tokopedia.net/img/cache/700/VqbcmM/2022/1/25/c2288826-35d6-4184-9555-61c4e20343c3.jpg.webp?ect=4g",
      },
      {
        title: "Jersey Evos",
        variant: "1",
        price: 40000,
        mount: 1,
        image:
          "https://images.tokopedia.net/img/cache/700/VqbcmM/2022/3/19/8ec93675-0a8a-4a26-aa83-4faf58b650d2.jpg.webp?ect=4g",
      },
      {
        title: "Jersey Onic",
        variant: "1",
        price: 60000,
        mount: 1,
        image:
          "https://images.tokopedia.net/img/cache/700/product-1/2021/4/17/942954/942954_fddd31ae-eba1-48b7-94cc-8af55c102164.jpg.webp?ect=4g",
      },
    ],
  }),
  computed: {
    summaryPrice(i) {
      let price = 0;
      price = this.list[i].price * this.list[i].mount;

      return price;
    },
  },
  methods: {
    methodPlus(i) {
      this.list[i].mount++;
    },
    minusMethod(i) {
      if (this.list[i].mount <= 1) {
        this.deleteMethod(i);
        return;
      }
      this.list[i].mount--;
    },
    deleteMethod(index) {
      console.log(index);
      this.list.splice(index, 1);
    },
    totalPrice(i) {
      let price = 0;
      price = this.list[i].price * this.list[i].mount;
      return price;
    },
  },
};
</script>

<style></style>
