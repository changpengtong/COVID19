<template xmlns:img="http://www.w3.org/1999/html">
    <el-card class="info-card" :body-style="{ padding: '0px' }" shadow="hover">
        <el-row :gutter="24">
            <el-col :md="24" :lg="6">
                <div class="info-image" >
                    <img src="./../../assets/img/institution.png">
                </div>
            </el-col>
            <el-col :md="24" :lg="16">
                <div class="info-text">
                    <div class="info-item">
                        <span id="name" class="info-item-title">{{iname}}</span>
                    </div>
                </div>
            </el-col>
        </el-row>
    </el-card>
</template>

<script>
export default {
    props:['name'],
    components: {
    },
    data() {
      return {
            iname: "",
      }
    },
    watch:{
        name:function (newData) {
            this.iname = newData[0]['institution']
        }
    }

  }
</script>
<style scoped>

    .show-developer {
        margin: 1.5em auto;
    }

    .info-card {
        max-width: 60em;
        margin: 2em auto;
        padding: 1em 1em;
    }
    .info-card {
    }
    .info-image img {
        object-fit: cover;
        border-radius: 50%;
        height: 10em;
        width: 10em;
        display: block;
        margin: auto;
    }
    .info-text{
        margin: .5em 0;
    }
    .info-item {
        margin: .5em 0;
    }
    .info-item-title {
        font-weight: normal;
        text-transform: capitalize;
        margin-right: .5em;
    }
    #name {
        font-size: 1.5em;
        font-weight: 500;
        margin: 0;
        position: absolute;
        top: 50%;
        -ms-transform: translateY(-50%);
        transform: translateY(-50%);
    }

</style>