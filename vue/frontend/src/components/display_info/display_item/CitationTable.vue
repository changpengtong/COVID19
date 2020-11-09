<template>
    <el-table
            :data="tableData"
            style="width: 100%"
            :header-cell-style="{
                'background-color': '#ffffff'
            }"
    >
        <el-table-column
                prop="text"
                align="left"
        >
        </el-table-column>
        <el-table-column
                prop="all"
                label="All"
                align="right"
        >
        </el-table-column>
        <el-table-column
                prop="since2015"
                label="Since 2015"
                align="right"
        >
        </el-table-column>
    </el-table>
</template>

<script>
    export default {
        name: "CitationTable",
        data() {
            return {
                tableData: [{
                    text: "Citations",
                    all:1998,
                    since2015:1486
                }, {
                    text: "h-index",
                    all:21,
                    since2015:19,
                }, {
                    text: "i10-index",
                    all:35,
                    since2015: 30
                }]
            }
        },

    }


</script>

<style scoped>
    .el-table__row>td {
        border:none;
        border-color: black;
    }

    .el-table::before {
        height: 0px
    }


</style>