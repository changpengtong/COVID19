<template>
  <div class="footer _bg-grey">
    <div class="_container">
      <p ><span>This Portal is Funded by</span> NSF RAPID (2028717)<span> and is built by </span>Yan Zhan, Sanjana Tripathi, Aoshen Wan, Jiarong Yang,</p>
      <p>Yijing Chen, Mona Sachdev, and Ying Ding
       <span>from</span> <a href="https://www.ischool.utexas.edu/" tag="a"> the School of Information at University of Texas at Austin</a>.</p>
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

  padding-top: 20px;
  padding-bottom: 20px;
  text-align: center;
  text-transform: capitalize;

  min-height: 50px;
}
._container span{
  color: #303133;
  /*font-size: small;*/
}
/*._container p span{*/
/*  color: #434343;*/
/*}*/
/*._container p{*/
/*  color: #409eff;*/
/*}*/
</style>