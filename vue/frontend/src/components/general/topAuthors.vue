<template>
  <el-table
      :data="tableData"
      :stripe="true"
  border
      style="width: 100%">
    <el-table-column
        prop="name"
        label="Name"
        width="150">
    </el-table-column>
    <el-table-column
        prop="affiliation"
        label="Affiliation"
        width="190">
    </el-table-column>
    <el-table-column
        prop="paper"
        label="Paper">
    </el-table-column>
  </el-table>
</template>

<script>
export default {
  data() {
    return {
      tableData: [{
        name: 'Bin Chen',
        affiliation:'Michigan State University',
        paper: '99'
      }, {
        name: 'Bin Chen',
        affiliation:'Michigan State University',
        paper: '99'
      }, {
        name: 'Bin Chen',
        affiliation:'Michigan State University',
        paper: '99'
      }, {
        name: 'Bin Chen',
        affiliation:'Michigan State University',
        paper: '99'
      }, {
        name: 'Bin Chen',
        affiliation:'Michigan State University',
        paper: '99'
      }, {
        name: 'Bin Chen',
        affiliation:'Michigan State University',
        paper: '99'
      }, {
        name: 'Bin Chen',
        affiliation:'Michigan State University',
        paper: '99'
      }, {
        name: 'Bin Chen',
        affiliation:'Michigan State University',
        paper: '99'
      }, ]
    }
  }
}
</script>

<!--<style scoped>-->
<!--/*/deep/ .el-table .gray-background {*/-->
<!--/*  background: #f2f2f2;*/-->
<!--/*}*/-->
<!--/*/deep/ .el-table .white-background {*/-->
<!--/*  background: #ffffff;*/-->
<!--/*}*/-->
<!--/*.entity-name {*/-->
<!--/*  !*text-decoration:underline;*!*/-->
<!--/*  color: black;*/-->
<!--/*}*/-->
<!--</style>-->