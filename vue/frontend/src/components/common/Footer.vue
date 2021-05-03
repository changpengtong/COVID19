<template>
  <div class="footer _bg-grey">
    <div class="_container">
      <p><span>This Portal is Funded by</span> NSF RAPID (2028717), Suit Endowment Fund and Mary R. Boyvey Dean’s Excellence Fund, and is hosted by the School of Information at University of Texas at Austin</p>
      <p>Website Organized by Aoshen Wan, Yan Zhan, Sanjana Tripathi, Jiarong Yang and Mona Sachdev</p>
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
  clear: both;
  position: relative;
  bottom: 1%;
  height: 90%;
  width: 100%;
  background-color: #edf4ff !important;

  text-align: center;
  text-transform: capitalize;
}
._container span{
  color: #303133;
}
</style>