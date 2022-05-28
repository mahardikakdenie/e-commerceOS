<template>
  <!-- Product Shop Section Begin -->
  <section class="product-shop spad page-details">
    <div class="container">
      <div class="row">
        <div class="col-lg-12">
          <div class="row">
            <div class="col-lg-6">
              <div class="product-pic-zoom">
                <img class="product-big-img" :src="img_default" alt="" />
              </div>
              <div class="product-thumbs">
                <carousel
                  v-if="images.length > 0"
                  :items="5"
                  :nav="false"
                  :dots="false"
                  :autoplay="true"
                  class="product-thumbs-track ps-slider"
                >
                  <div
                    v-for="(v, index) in images"
                    :key="index"
                    :class="`pt ${
                      img_default === v.url ? 'active' : ''
                    } rounded-lg img--choice`"
                    @click="clickImg(v.url)"
                  >
                    <img :src="v.url" alt="" height="160" />
                  </div>
                  <div>
                    <img
                      v-if="computedShow.link"
                      :class="`pt ${
                        computedShow.link.image === img_default ? 'active' : ''
                      } rounded-lg img--choice`"
                      @click="clickImg(computedShow.link.image)"
                      :src="computedShow.link.image"
                      height="160"
                      alt=""
                    />
                  </div>
                </carousel>
              </div>
            </div>
            <div class="col-lg-6">
              <div class="product-details text-left">
                <div v-if="computedShow.entities" class="pd-title">
                  <span>{{ computedShow.entities.category.Name }} </span>
                  <h3>
                    {{ computedShow.name }}
                  </h3>
                </div>

                <div class="pd-desc">
                  <p></p>
                  <p></p>
                  <p>
                    {{ computedShow.description }}
                  </p>
                  <div class="card variants" style="">
                    <!-- <img src="..." class="card-img-top" alt="..." /> -->
                    <div class="card-body">
                      <h3 class="card-title">Pilih Variant</h3>
                      <div class="img">
                        <img
                          v-for="i in 4"
                          :key="i"
                          style="margin-right: 10px"
                          width="80"
                          class="img--variant"
                          src="https://images.tokopedia.net/img/cache/100-square/VqbcmM/2021/12/16/63023bc9-f402-43be-b96c-f95d5eaa2a01.jpg"
                          alt=""
                        />
                      </div>
                      <form action="">
                        <div class="mb-3">
                          <label for="exampleInputEmail1" class="form-label"
                            >Catatan Kiriman</label
                          >
                          <input
                            type="email"
                            class="form-control"
                            id="exampleInputEmail1"
                            aria-describedby="emailHelp"
                          />
                          <div id="emailHelp" class="form-text">
                            *Catatan Untuk Penjual
                          </div>
                        </div>
                      </form>
                    </div>
                  </div>
                  <p></p>
                  <h4>Rp.{{ computedShow.price }}</h4>
                  <h4>Stock : {{ computedShow.stock }}</h4>
                </div>

                <div class="quantity">
                  <button
                    style="width: 100%"
                    class="btn primary-btn rounded btn-full"
                  >
                    Add To Cart
                  </button>
                </div>
              </div>
            </div>
          </div>
        </div>
      </div>
    </div>
  </section>
  <!-- Product Shop Section End -->
</template>

<script>
import carousel from "vue-owl-carousel";
export default {
  components: {
    carousel,
  },
  props: {
    imgs: {
      type: Array,
      default: null,
    },

    data: {
      type: Object,
      default: null,
    },
  },
  data: () => ({
    images: [],
    img_default:
      "https://images.unsplash.com/photo-1434389677669-e08b4cac3105?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8Nnx8Y2xvdGhlc3xlbnwwfHwwfHw%3D&auto=format&fit=crop&w=500&q=60",
    items: [
      {
        url: "https://images.tokopedia.net/img/cache/200-square/VqbcmM/2021/7/30/55994347-7eaf-4d11-b9fe-6fc9183ed3f9.jpg.webp?ect=4g",
      },
      {
        url: "https://images.unsplash.com/photo-1434389677669-e08b4cac3105?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8Nnx8Y2xvdGhlc3xlbnwwfHwwfHw%3D&auto=format&fit=crop&w=500&q=60",
      },
      {
        url: "https://images.unsplash.com/photo-1572454591674-2739f30d8c40?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8MXx8Zmxvd2VyJTIwYm91cXVldHxlbnwwfHwwfHw%3D&auto=format&fit=crop&w=500&q=60",
      },
      {
        url: "https://images.unsplash.com/photo-1565695951564-007d8f297e48?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8NXx8Zmxvd2VyJTIwYm91cXVldHxlbnwwfHwwfHw%3D&auto=format&fit=crop&w=500&q=60",
      },
    ],
  }),
  computed: {
    computedShow() {
      return this.$store.state.all_product.show;
    },
    computedImage() {
      return this.$store.state.product_image.data;
    },
  },
  mounted() {
    this.getDataShowProduct();
    this.getDataProductImage();
    // this.isDefault();
    // for (const item in this.data.galeries) {
    //   if (Object.hasOwnProperty.call(this.data.galeries, item)) {
    //     const img = this.data.galeries[item];
    //     if (img.is_default) {
    //       this.img_default = img.photo;
    //     }
    //   }
    // }
  },
  methods: {
    clickImg(url) {
      console.log("before", this.computedShow.link.image);
      this.img_default = url;
      console.log("before", this.computedShow.link.image);
    },
    addCart() {
      this.$emit("add-cart");
    },
    getDataShowProduct() {
      this.$store
        .dispatch("all_product/getDataShowProduct", {
          id: this.$route.params.id,
          store_id: localStorage.getItem("store_id"),
          entities: "Category,User,Store",
        })
        .then((res) => {
          if (res.data.meta.status) {
            this.img_default = res.data.data.link.image;
            console.log("show data => ", res.data.data);
          }
        });
    },
    getDataProductImage() {
      this.$store
        .dispatch("product_image/getDataProductImage", {
          product_id: this.$route.params.id,
        })
        .then((res) => {
          if (res.data.meta.status) {
            this.images = res.data.data;
            console.log("show data Image=> ", this.images);
          }
        });
    },
    isDefault() {
      this.img_default = this.computedShow.link.image;
    },
  },
};
</script>

<style scoped>
.product-thumbs .pt {
  margin-right: 14px;
}
.cursor-pointer {
  cursor: pointer;
}
.text-left {
  text-align: left;
}
.product-thumbs-track .img--choice:hover {
  border: 1px solid #e7ab3c;
}
.variants {
  margin-bottom: 20px;
  width: 100%;
}

.variants .img {
  padding: 10px;
  display: flex;
  justify-content: center;
}

.variants .img .img--variant {
  transition: 1s;
}

.variants .img .img--variant:hover {
  border: 1px solid #e7ab3c;
  cursor: pointer;
}
</style>
