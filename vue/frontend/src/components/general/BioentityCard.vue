<template>
    <el-card class="info-card" :body-style="{ padding: '0px' }" shadow="hover">
        <el-row :gutter="24">
            <el-col :md="24" :lg="6">
                <div class="info-img">
                    <img src="https://encrypted-tbn0.gstatic.com/images?q=tbn%3AANd9GcTvQtNbrzI2mYCjOqaq5nBvzy6-FNDlNpWMYg&usqp=CAU">
                </div>
            </el-col>
            <el-col :md="24" :lg="16">
                <div class="info-text">
                    <div class="info-item-title">
<!--                        <span style="font-weight:normal">{{tableData.entity}}</span>-->
                        <span class="info-item"></span><span>{{tableData.type}}</span>
                    </div>
                    <div id="name" class="info-item-title">
<!--                      <span>Name: </span>-->
                      <span>{{tableData.name}}</span>
                    </div>
                </div>
            </el-col>
        </el-row>
    </el-card>
</template>

<script>
// import $search from "../../util/search"

    export default {
        props:['name'],
        data() {
            return {

                tableData: {
                    type:'',
                    name: '',
                }

            }
        },
        watch:{
            name:function (newData) {
              console.log("bioentity card is: ")
                console.log(newData)
              this.tableData.name=newData
            }
        }
      // watch:{
      //   card:function (newData) {
      //     console.log("bio-card")
      //     console.log(newData)
      //     this.tableData.name = newData["bioentity"]
      //     this.tableData.type = newData["type"]
      //   }
      // }

    }
</script>

<style scoped>

    .show-developer {
        margin: 1.5em auto;
    }

    .info-card {
        max-width: 60em;
        /*width: 50em;*/
        /* display: inline-block; */
        margin: 2em auto;
        padding: 1em 1em;
    }
    .info-card {
        /*margin: 0 auto;*/
        /*padding: 1em;*/
    }
    img {
        object-fit: cover;
        border-radius: 50%;
        height: 120px;
        width: 120px;
        display: block;
        margin: auto;
        /*margin-top: 1em;*/

    }
    .info-text{
        margin: 0;
        position: absolute;
        top: 50%;
        -ms-transform: translateY(-50%);
        transform: translateY(-50%);
        /*-ms-transform: translateY(-50%);*/
        /*transform: translateY(-50%);*/
    }
    .info-item-name {
        margin: .5em 0;
        /*font-weight: normal;*/
        text-transform: capitalize;
        font-weight: 500;

    }
    .info-item-title {
      font-weight: normal;
      text-transform: capitalize;
      margin-right: .5em;
    }
    #type {
        font-size: 1.5em;
        font-weight: 500;
        margin: 0;
        position: absolute;
        top: 50%;

    }
    #name {
      font-size: 1.5em;
      font-weight: 500;
      margin: 0;
      /*position: absolute;*/
      top: 50%;
      -ms-transform: translateY(-50%);
      transform: translateY(-50%);
    }

</style>