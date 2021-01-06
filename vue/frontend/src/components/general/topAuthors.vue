<template>
  <el-table
      :data="tableData"
      :stripe="true"
  border
      style="width: 100%">
    <el-table-column
        prop="name"
        label="Name"
        width="138">
      <template slot-scope="scope">
        <a :href="scope.row.authorUrl" style="color:#000000">{{scope.row.name}}</a>
      </template>
    </el-table-column>
    <el-table-column
        prop="affiliation"
        label="Affiliation"
        >
      <template slot-scope="scope">
        <a :href="scope.row.affiliationUrl" style="color:#000000">{{scope.row.affiliation}}</a>
      </template>
    </el-table-column>
  </el-table>
</template>

<script>
import data from "../../data/top10Authors.json"
export default {
  props:['author'],
  data() {
    return {
      tableData:data,
    }
  },
  methods: {
    // 发送检索请求
    getDetailPage(row) {
      row.authorUrl="/COVID19/"+row.authorUrl,
          row.affiliationUrl="/COVID19/"+row.affiliationUrl

    },
  },

}
</script>
