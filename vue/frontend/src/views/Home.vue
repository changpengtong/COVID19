<template>
  <div class="home">
    <div class="context">
      <div class="search">
        <search-wrapper></search-wrapper>
      </div>

      <div class="table">
        <el-row>
          <!------------------------------- Table Institutions ------------------------------->
          <el-col :span="6">
            <small-table
              :tableName="tableName[0]"
              :dataList="dataInstitutions"
              :type="1"
            ></small-table>
          </el-col>
          <!------------------------------- Table Bioentities ------------------------------->
          <el-col :span="6">
            <small-table
              :tableName="tableName[1]"
              :dataList="dataEntities"
              :type="2"
            ></small-table>
          </el-col>
          <!------------------------------- Table Authors ------------------------------->
          <el-col :span="6">
            <small-table
              :tableName="tableName[2]"
              :dataList="dataAuthors"
              :type="3"
            ></small-table>
          </el-col>
        </el-row>
      </div>
    </div>
  </div>
</template>

<script>
import dataEntities from "../data/top5BioEntities.json";
import dataInstitutions from "../data/top5Institutions.json";
import dataAuthors from "../data/top5Authors.json";
import SmallTable from "../components/home/SmallTable.vue";
import SearchWrapper from "../components/home/SearchWrapper";

export default {
  name: "home",

  data() {
    return {
      dataEntities: dataEntities,
      dataAuthors: dataAuthors,
      dataInstitutions: dataInstitutions,

      tableName: ["Institutions", "Bioentities", "Authors"],
    };
  },
  components: {
    "small-table": SmallTable,
    "search-wrapper": SearchWrapper,
  },
  methods: {
    getData() {
      fetch("../data/sample.json")
        .then((response) => response.json())
        .then((data) => (this.dataList = data));
    },
    getInstitutionDetailPage(row) {
      this.$router.push("path").catch((error) => {
        if (error.name != "NavigationDuplicated") {
          throw error;
        }
      });
      this.$router.push({
        path: "/institution/" + row.name,
      });
    },
    getAuthorDetailPage(row) {
      this.$router.push("path").catch((error) => {
        if (error.name != "NavigationDuplicated") {
          throw error;
        }
      });
      this.$router.push({
        path: "/itemAuthor/" + row.name,
      });
    },
    getPaperDetailPage(row) {
      this.$router.push({
        path: "/displayInfo",
        query: {
          query: row.name,
        },
      });
    },
  },
};
</script>

<style>
.home {
  background: url(../assets/img/home/blue-background.png) no-repeat center
    center fixed;
  -webkit-background-size: cover;
  -moz-background-size: cover;
  -o-background-size: cover;
  background-size: cover;
  background-attachment: fixed;
  height: 700px;
}

.home .context {

}

.home .search {
  padding-top: 190px;
  padding-left: 25%;
}

.home .table {
  padding-top: 35px;
  padding-left: 22%;
  padding-right: 9%;
  position: fixed;
}

.home .el-col-6 {
  margin-right: 50px;
}
.home .home-search-pane {
  display: block;
  width: 100%;
  padding-bottom: -100px;
  padding-top: 15px;
}

.home .home-search-pane .search-box {
  padding-bottom: 10px;
}

.home .home-search-pane .radios .el-radio-group {
  padding-top: 13px;
  padding-left: 15%;
}
.home .search-box .search-box-auto {
  position: center;
  width: 770px;
}

.home .search-box .el-col-6 {
  min-width: 100px;
  padding-right: 100px;
}

.home .search-box .el-input__inner {
  border-radius: 20px 20px;
}

.home .search-box .el-icon-search {
  margin-top: 4px;
  font-size: 17px;
  margin-right: 10px;
  border: 0px;
}
</style>