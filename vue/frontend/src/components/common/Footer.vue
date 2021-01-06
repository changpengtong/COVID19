<template>
  <div class="footer _bg-grey">
    <div class="_container">
      <span
        >&copy;<router-link to="/" tag="a">Covid-19 Portal</router-link>. All
        Rights Reserved
      </span>
      <!-- <router-link to="/about" tag="a" class="item">contact us</router-link> -->
    </div>
  </div>
</template>

<script>
export default {
  name: "Footer",
};
</script>

<style>
.footer {
  /*position: fixed;*/
  left: 0;
  bottom: 0;
  width: 100%;
  background-color: #edf4ff !important;

  padding-top: 10px;
  text-align: center;
  text-transform: capitalize;

  min-height: 50px;
}
</style>