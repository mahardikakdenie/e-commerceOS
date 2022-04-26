<template>
  <div>
    <div class="footer">
      <footer>
        <div class="container">
          <div class="row">
            <div class="col-sm-6 col-md-3 item">
              <h2 class="text-left">Himatif Store</h2>
              <ul class="">
                <li>
                  <a href="#"
                    >“My new website attracted a lot of attention in the best
                    possible way. I am forever grateful.”</a
                  >
                </li>
                <li class="mt-6">
                  <a href="">@ARS University, 2022</a>
                </li>
              </ul>
            </div>
            <div class="col-sm-6 col-md-3 item">
              <h2 class="text-left">Social Media</h2>
              <ul class="__list">
                <li>
                  <a href="#"
                    ><img
                      style="margin-right: 8px"
                      :src="ig"
                      alt=""
                    />@himatif_store</a
                  >
                </li>
                <li>
                  <a href="#"
                    ><img
                      class="mr-2"
                      style="margin-right: 8px"
                      :src="tw"
                      alt=""
                    />@himatif_store</a
                  >
                </li>
              </ul>
            </div>
            <div class="col-sm-6 col-md-3 item">
              <h2 class="text-left">Payment Method</h2>
              <ul class="__list">
                <li>
                  <a href="#"
                    ><img
                      style="margin-right: 8px"
                      :src="paypal"
                      alt=""
                    />Paypal</a
                  >
                </li>
              </ul>
            </div>
            <div class="col-sm-6 col-md-3 item">
              <h2 class="text-left">Contact</h2>
              <ul class="__list">
                <li>
                  <a href="#"
                    ><img
                      style="margin-right: 8px"
                      :src="telp"
                      alt=""
                    />089621174934</a
                  >
                </li>
                <li>
                  <a href="#"
                    ><img
                      class="mr-2"
                      style="margin-right: 8px"
                      :src="wa"
                      alt=""
                    />089621174934</a
                  >
                </li>
              </ul>
            </div>
          </div>
          <p class="copyright">Company Name © 2018</p>
        </div>
      </footer>
    </div>
  </div>
</template>

<script>
import ig from "../../assets/images/ig.png";
import paypal from "../../assets/images/paypal.png";
import tw from "../../assets/images/tw.png";
import telp from "../../assets/images/telp.png";
import wa from "../../assets/images/wa.png";
export default {
  data: () => ({
    ig: ig,
    tw: tw,
    paypal: paypal,
    telp: telp,
    wa: wa,
  }),
};
</script>

<style></style>
