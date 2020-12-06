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
                        <span id="name" class="info-item-title">{{profile.name}}</span>
                    </div>
                    <!-- affilliation-->
                    <div class="info-item" v-if="profile">
                        <span v-if="profile.lab">{{profile.lab}}, </span>
<!--                        <router-link :to=affiliation.institution.url >-->
                            <a :href="profile.url">
                           <span class="institution" v-if="profile.institution">{{profile.institution}}</span>
<!--                        </router-link>-->
                            </a>
                    </div>
                    <!-- mail -->
                    <div class="info-item" v-if="profile.mail">
                        <span class="info-item-title">Email:</span><span>{{profile.mail}}</span>
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
    props:['card'],
    components: {
      // "info-card": CardHelper
    },
    data() {
      return {
            profile:{
                name: 'bin chen',
                lab: 'COVID Lab',
                url: 'google.com',
                institution: 'msu',
                mail: '12345@qq.com',
            }
      }
    },
    watch:{
        card:function (newData) {
            console.log("1")
            console.log(newData)
            this.profile.name = newData[0]["ForeName"]+" "+newData[0]["LastName"]
            this.profile.mail = ""
            this.profile.institution = newData[0]["Affiliation"]
            this.profile.lab=""
            this.profile.url = "/COVID19/#/institution/"+newData[0]["Affiliation"]
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