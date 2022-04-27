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
                  :items="4"
                  :nav="false"
                  :dots="false"
                  :autoplay="true"
                  class="product-thumbs-track ps-slider"
                >
                  <div
                    v-for="(v, index) in items"
                    :key="index"
                    :class="`pt ${
                      img_default === v.url ? 'active' : ''
                    } rounded-lg img--choice`"
                  >
                    <img :src="v.url" alt="" height="160" />
                  </div>
                </carousel>
              </div>
            </div>
            <div class="col-lg-6">
              <div class="product-details text-left">
                <div class="pd-title">
                  <span>asdas</span>
                  <h3>
                    Kacamata Photocromic Korea 2 In 1 Kacamata Lensa Photocromic
                    Anti Uv Wanita Dan Pria
                  </h3>
                </div>
                <div class="pd-desc">
                  <p>
                    Lorem ipsum dolor sit amet consectetur adipisicing elit.
                    Corporis, error officia. Rem aperiam laborum voluptatum vel,
                    pariatur modi hic provident eum iure natus quos non a sequi,
                    id accusantium! Autem.
                  </p>
                  <p>
                    Lorem ipsum dolor sit, amet consectetur adipisicing elit.
                    Quam possimus quisquam animi, commodi, nihil voluptate
                    nostrum neque architecto illo officiis doloremque et
                    corrupti cupiditate voluptatibus error illum. Commodi
                    expedita animi nulla aspernatur. Id asperiores blanditiis,
                    omnis repudiandae iste inventore cum, quam sint molestiae
                    accusamus voluptates ex tempora illum sit perspiciatis.
                    Nostrum dolor tenetur amet, illo natus magni veniam quia sit
                    nihil dolores. Commodi ratione distinctio harum voluptatum
                    velit facilis voluptas animi non laudantium, id dolorem
                    atque perferendis enim ducimus? A exercitationem recusandae
                    aliquam quod. Itaque inventore obcaecati, unde quam impedit
                    praesentium veritatis quis beatae ea atque perferendis
                    voluptates velit architecto?
                  </p>
                  <h4>$20</h4>
                </div>
                <div class="quantity">
                  <a class="cursor-pointer primary-btn pd-cart">Add To Cart</a>
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
    // img_default: {
    //   type: String,
    //   default: null,
    // },
    data: {
      type: Object,
      default: null,
    },
  },
  data: () => ({
    img_default:
      "https://images.unsplash.com/photo-1434389677669-e08b4cac3105?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8Nnx8Y2xvdGhlc3xlbnwwfHwwfHw%3D&auto=format&fit=crop&w=500&q=60",
    items: [
      {
        url: "https://images.unsplash.com/photo-1467043237213-65f2da53396f?ixlib=rb-1.2.1&ixid=MnwxMjA3fDB8MHxzZWFyY2h8N3x8Y2xvdGhlc3xlbnwwfHwwfHw%3D&auto=format&fit=crop&w=500&q=60",
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
  mounted() {
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
      this.img_default = url;
    },
    addCart() {
      this.$emit("add-cart");
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
</style>
