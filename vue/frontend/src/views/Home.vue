<template>
  <div class="home">
    <el-row :gutter="20" type="flex" justify="center">
      <el-col :span="20">
        <el-row :gutter="20" class="first-row">
          <el-col :span="12" class="first-col">
            <!------------------------------- total paper card ------------------------------->
            <el-row>
              <el-card class="total-paper-card">
                <div class="word">Total Paper</div>
                <div class="number">301,667</div>
              </el-card>
            </el-row>

            <!------------------------------- welcome word card ------------------------------->
            <el-row>
              <el-card class="welcome-card">
                <!-- <img src="../image/bg.png" class="welcome-image"> -->
                <div class="word">
                  <h3>Welcome!</h3>
                  <p class="first">
                    COVID-19 Portal is a dense trove of coronavirus research
                    trends. Updated regularly, this website is an atlas to help
                    researchers and people who are interested in cutting-edge
                    research on COVID-19 find influential researchers, papers,
                    and institutions.
                  </p>
                  <p class="second">
                    Our datasets include over 2,200,000 research paper, covering
                    Kaggle, PubMed, PubTator center.
                  </p>
                </div>
              </el-card>
            </el-row>
          </el-col>

          <el-col :span="12" class="second-col">
            <!------------------------------- Table authors ------------------------------->

            <el-card class="authors-table-card">
              <div class="authors-table">
                <div class="title">Top 10 Authors</div>
                <el-table
                    ref="singleTable"
                    :data="dataAuthors"
                    stripe
                    height
                    highlight-current-row
                    fit
                    @current-change="handleCurrentChange"
                    @row-click="getInstitutionDetailPage"
                    style="width: 100%"
                >
                  <el-table-column type="index" width="34"> </el-table-column>

                  <el-table-column property="name" label="Name" width="135">
                  </el-table-column>

                  <el-table-column property="institution" label="Institution">
                  </el-table-column>

                  <el-table-column property="paper" label="Paper" width="60">
                  </el-table-column>
                </el-table>
              </div>
            </el-card>

          </el-col>
        </el-row>

        <el-row :gutter="20">
          <div class="second-row">
            <!------------------------------- Table genes ------------------------------->
            <el-col :span="8">
              <small-table
                  :tableName="tableName[1]"
                  :dataList="dataGenes"
              ></small-table>
            </el-col>
            <!------------------------------- Table drug ------------------------------->
            <el-col :span="8">
              <small-table
                  :tableName="tableName[2]"
                  :dataList="dataDrugs"
              ></small-table>
            </el-col>
            <!------------------------------- Table diseases ------------------------------->
            <el-col :span="8">
              <small-table
                  :tableName="tableName[3]"
                  :dataList="dataDiseases"
              ></small-table>
            </el-col>
          </div>
        </el-row>
      </el-col>

      <el-col :span="12" class="third-col">
        <!------------------------------- Table paper ------------------------------->
        <el-card class="papers-table-card">
          <div class="papers-table">
            <div class="title">Top 10 Papers about covid-19</div>
            <el-table
                ref="singleTable"
                :data="dataPapers"
                stripe
                height
                highlight-current-row
                fit
                @current-change="handleCurrentChange"
                @row-click="getPaperDetailPage"
                style="width: 100%"
            >
              <el-table-column type="index" width="36"> </el-table-column>

              <el-table-column property="name" label="Title"> </el-table-column>

              <el-table-column property="tweet" label="Tweet" width="60">
              </el-table-column>
            </el-table>
          </div>
        </el-card>
      </el-col>
    </el-row>
  </div>
</template>

<script>
import dataEntities from "../data/top10Entities.json";
import dataDiseases from "../data/top10Diseases.json";
import dataDrugs from "../data/top10Drugs.json";
import dataGenes from "../data/top10Genes.json";
import dataPapers from "../data/top10Papers.json";
import dataAuthors from "../data/top10Authors.json";
import SmallTable from "../components/home/SmallTable.vue";

export default {
  name: "home",

  data() {
    return {
      dataEntities: dataEntities,
      dataDiseases: dataDiseases,
      dataDrugs: dataDrugs,
      dataGenes: dataGenes,
      dataPapers: dataPapers,
      dataAuthors: dataAuthors,

      tableName: [
        "Entities",
        "Genes",
        "Drugs",
        "Diseases",
        "Papers",
        "Authors",
      ],
    };
  },
  components: {
    "small-table": SmallTable,
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
        path: "/author/" + row.name,
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

<style scoped>
.home {
  margin-top: 40px;
  margin-left: 5%;
  margin-right: 4%;
}


/*------------------------------- total paper card -------------------------------*/
.home .total-paper-card {
  height: 130px;
  text-align: center;
  margin-bottom: 20px;
}

.home .total-paper-card .word {
  font-size: 20px;
  font-weight: 600;
  font-family: 'Slabo 27px', serif;
}

.home .total-paper-card .number {
  margin-top:-10px;
  font-size: 47px;
  color: #409eff;
  font-weight: 600;
  font-family: 'Playfair Display', serif;
}

/*------------------------------- welcome word card -------------------------------*/
.home .welcome-card {
  height: 230px;
  margin-bottom: 20px;
  background-color: #EDF4FF;

}

.home .welcome-card .el-card__body {
  padding: 0px;
}

.home .welcome-card .welcome-image {
  width: 100%;
  display: block;
}

.home .first-col .welcome-card {
  background-image: url("../image/welcome.png");
  background-repeat: no-repeat, repeat;
  background-size: 350px 150px;

}

.home .first-col .welcome-card .word h3 {
  padding-top: 18px;
  text-align: center;
  font-family: 'Merienda', cursive;
  font-weight: 250;

}

.home .first-col .welcome-card .word .first {
  padding-top: 10px;
  padding-left: 50px;
  text-align: left;
  font-size: 15px;
  font-weight: 600;
  color: #102443;
  font-family: 'Piazzolla', serif;
}
.home .first-col .welcome-card .word .second {
  padding-top: 15px;
  padding-left: 30px;
  text-align: left;
  font-size: 12px;
  font-weight: 500;
  color: #4876bf;
  font-family: 'Satisfy', cursive;
}

/*------------------------------- Tables -------------------------------*/

.home .el-table {
  padding-top: 10px;
}

.home .el-table__row {
  font-size: 13px;
}
.home .authors-table-card .el-table__row {
  font-size: 12px;
}

.home .el-table--striped .el-table__body tr.el-table__row--striped td {
  background-color: #edf4ff;
}

.home .el-table td,
.el-table th {
  padding: 0;
}

.home .title {
  font-weight: bold;
  font-family: 'Slabo 27px', serif;
}

.second-col .authors-table-card {
  margin-bottom: 20px;
  height: 380px;
  text-align: center;
  font-size: 20px;
}
.second-col .authors-table-card .title{
  font-size: 21px;
}

.third-col .papers-table-card {
  height: 745px;
  text-align: center;
  font-size: 25px;
}

.home .third-col .papers-table-card .el-table__row {
  font-size: 14px;
  height: 20px;
}

.home .third-col .papers-table-card .el-table .cell {
  height: 47px;
}
</style>