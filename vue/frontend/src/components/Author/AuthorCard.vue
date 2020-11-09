<template xmlns:img="http://www.w3.org/1999/html">
    <!-- 开发者名片 -->
    <el-card class="info-card" :body-style="{ padding: '0px' }" shadow="hover">
        <el-row :gutter="24">
            <el-col :md="24" :lg="6">
                <div class="info-image" >
                    <img src="../../assets/img/author.png">
                </div>
            </el-col>
            <el-col :md="24" :lg="16">
                <div class="info-text">
                    <div class="info-item">
                        <span id="name" class="info-item-title">{{name}}</span>
                    </div>
                    <!-- affilliation-->
                    <div class="info-item" v-if="affiliation">
                        <span v-if="affiliation.lab">{{affiliation.lab}}, </span>
                        <router-link :to=affiliation.institution.url >
                           <span class="institution" v-if="affiliation.institution">{{affiliation.institution.name}}</span>
                        </router-link>
                    </div>
                    <!-- mail -->
                    <div class="info-item" v-if="mail">
                        <span class="info-item-title">Email:</span><span>{{mail}}</span>
                    </div>
                    <!-- key words -->
<!--                    <div class="info-item" >-->
<!--                    <span class="words" v-for="(key, index) of keywords"-->
<!--                          :key="index"><a :href="key.word">{{key.word}}&nbsp;&nbsp;&nbsp;&nbsp;</a></span>-->
<!--                    </div>-->
                </div>
            </el-col>
        </el-row>
    </el-card>
</template>

<script>
// import http from "../../../util/axios";


export default {
    name: "about",
    components: {
      // "info-card": CardHelper
    },
    data() {
      return {

            name: "Bin Chen",
            mail: "bin.chen@hc.msu.edu",
          affiliation:{ lab:"Bin Chen Lab", institution:{name:"Michigan State University",
              url:"/institution"}},
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
         /*width: 50em;*/
        /* display: inline-block; */
        margin: 2em auto;
        padding: 1em 1em;
    }
    .center {
        margin: 0 auto;
    }
    .info-card {
         /*margin: 0 auto;*/
        /*padding: 1em;*/
    }
    .info-image img {
        object-fit: cover;
        border-radius: 50%;
        height: 120px;
        width: 120px;
        display: block;
        margin: auto;
        /*margin-top: 1em;*/

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
    .institution{
        font-weight: normal;
        text-transform: capitalize;
        margin-right: .5em;
        text-decoration:underline;
    }
    .words{
        white-space: pre-wrap;
        margin-right: 1rem;
    }
    #name {
        font-size: 1.5em;
        margin-bottom: .5em;
        font-weight: 500;
    }
    .info-descript {
        color: rgba(0,0,0,.5);
        margin-bottom: .5em;
    }
</style>