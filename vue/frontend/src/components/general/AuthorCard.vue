<template xmlns:img="http://www.w3.org/1999/html">
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
            <a :href="profile.url">
              <span class="institution" v-if="profile.institution">{{profile.institution}}</span>
            </a>
          </div>
          <div class="info-item" v-if="profile.mail">
            <span class="info-item-title">Location:</span><span>{{profile.location}}</span>
          </div>
          <!-- mail -->
          <div class="info-item" v-if="profile.mail">
            <span class="info-item-title">Email:</span><span>{{profile.mail}}</span>
          </div>
        </div>
      </el-col>
    </el-row>
  </el-card>
</template>

<script>
export default {
  props:['card'],
  components: {
  },
  data() {
    return {
      profile:{
        name: '',
        lab: '',
        url: '',
        institution: '',
        mail: '',
      }
    }
  },
  watch:{
    card:function (newData) {
      this.profile.name = newData[0]["FullName"]
      this.profile.mail = newData[0]["Email"]
      this.profile.institution = newData[0]["Affiliation"].replace(/^,+/,"").replace(/,+$/,"")
      this.profile.location = newData[0]["Location"]
      this.profile.url = "/COVID19/#/institution/" + newData[0]["institution_id"]
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
.center {
  margin: 0 auto;
}
.info-card {}

.info-image img {
  object-fit: cover;
  border-radius: 50%;
  height: 120px;
  width: 120px;
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