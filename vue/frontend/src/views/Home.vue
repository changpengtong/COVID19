<template>
    <div class="home">
        <!------------------------------- layout ------------------------------->
        <el-row :gutter="20">
            <!------------------------------- 1st col ------------------------------->
            <el-col :span="8" class="first-col">

                <!------------------------------- total paper card ------------------------------->
                <el-row :span="8">
                    <el-card class="total-paper-card">
                        <div class="word">Total Paper</div>
                        <div class="number">301,667</div>
                    </el-card>    
                </el-row>

                <!------------------------------- welcome word card ------------------------------->
                <el-row :span="8">
                    <el-card class="welcome-card">
                        <img src="../image/welcome.png" class="welcome-image">
                    </el-card>
                </el-row>

                <div class="left-table">
                    <el-row :gutter="20">
                <!------------------------------- Table entity ------------------------------->
                
                    <el-col :span="12">
                    <el-card class="small-table-card">
                        <div class="entities-table" >
                            <div class="title">Top 10 Entities</div>
                            <el-table ref="singleTable" :data="dataList"
                            stripe height highlight-current-row fit
                            @current-change="handleCurrentChange"
                            style="width: 100%">
                                <el-table-column
                                type="index"
                                width="34">
                                </el-table-column>
                                    
                                <el-table-column
                                property="name"
                                label="Name">
                                </el-table-column>
                                
                                <el-table-column
                                property="paper"
                                label="Paper"
                                width="60">
                                </el-table-column>
                            </el-table>
                        </div>
                    </el-card>
                    </el-col>
                <!------------------------------- Table genes ------------------------------->
                    <el-col :span="12">
                    <el-card class="small-table-card">
                        <div class="genes-table">
                            <div class="title">Top 10 Genes</div>
                            <el-table ref="singleTable" :data="dataList"
                            stripe height highlight-current-row fit
                            @current-change="handleCurrentChange"
                            style="width: 100%">
                                <el-table-column
                                type="index"
                                width="34">
                                </el-table-column>
                                    
                                <el-table-column
                                property="name"
                                label="Name">
                                </el-table-column>
                                
                                <el-table-column
                                property="paper"
                                label="Paper"
                                width="60">
                                </el-table-column>
                            </el-table>
                        </div>
                    </el-card>
                    </el-col>
                    </el-row>
                </div>
            </el-col>
            <!------------------------------- 2nd col ------------------------------->
            
            <el-col :span="8" class="second-col">
                <!------------------------------- Table authors ------------------------------->
                <el-row :span="8">
                <el-card class="authors-table-card">
                    <div class="authors-table">
                        <div class="title">Top 10 Authors</div>
                        <el-table ref="singleTable" :data="tableData"
                        stripe height highlight-current-row fit
                        @current-change="handleCurrentChange"
                        style="width: 100%">
                            <el-table-column
                            type="index"
                            width="60">
                            </el-table-column>
                                
                            <el-table-column
                            property="name"
                            label="Name"
                            width="135">
                            </el-table-column>

                            <el-table-column
                            property="institution"
                            label="Institution">
                            </el-table-column>
                            
                            <el-table-column
                            property="paper"
                            label="Paper"
                            width="60">
                            </el-table-column>
                        </el-table>
                    </div>
                </el-card>
                </el-row>

                <el-row :gutter="20">
                <!------------------------------- Table drug ------------------------------->
                <el-col :span="12">
                <el-card class="small-table-card">
                    <div class="drugs-table">
                        <div class="title">Top 10 Drugs</div>
                        <el-table ref="singleTable" :data="dataList"
                        stripe height highlight-current-row fit
                        @current-change="handleCurrentChange"
                        style="width: 100%">
                            <el-table-column
                            type="index"
                            width="34">
                            </el-table-column>
                                
                            <el-table-column
                            property="name"
                            label="Name">
                            </el-table-column>
                            
                            <el-table-column
                            property="paper"
                            label="Paper"
                            width="60">
                            </el-table-column>
                        </el-table>
                    </div>
                </el-card>
                </el-col>
                <!------------------------------- Table diseases ------------------------------->
                <el-col :span="12">
                <el-card class="small-table-card">
                    <div class="diseases-table">
                        <div class="title">Top 10 Diseases</div>
                        <el-table ref="singleTable" :data="dataList"
                        stripe height highlight-current-row fit
                        @current-change="handleCurrentChange"
                        style="width: 100%">
                            <el-table-column
                            type="index"
                            width="34">
                            </el-table-column>
                                
                            <el-table-column
                            property="name"
                            label="Name">
                            </el-table-column>
                            
                            <el-table-column
                            property="paper"
                            label="Paper"
                            width="60">
                            </el-table-column>
                        </el-table>
                    </div>
                </el-card>
                </el-col>
                </el-row>
            </el-col>
            <!------------------------------- 3rd col ------------------------------->
            <!------------------------------- Table paper ------------------------------->
            <el-col :span="8" class="third-col">
            <el-card class="papers-table-card">
                <div class="papers-table">
                    <div class="title">Top 10 Papers about covid-19</div>
                    <el-table ref="singleTable" :data="tableData"
                    stripe height highlight-current-row fit
                    @current-change="handleCurrentChange"
                    style="width: 100%">
                        <el-table-column
                            type="index"
                            width="60">
                            </el-table-column>
                                
                            <el-table-column
                            property="title"
                            label="Title">
                            </el-table-column>

                            
                            <el-table-column
                            property="count"
                            label="Tweet_count"
                            width="110">
                            </el-table-column>
                    </el-table>
                </div>
            </el-card>
            </el-col>
        </el-row>
    </div>
</template>

<script>
    import data from "../data/sample.json"

    export default {
        name: 'home',
        data() {
            return {
                dataList: data
            };
        },
        methods: {
            getData() {
                fetch("../data/sample.json")
                    .then(response => response.json())
                    .then(data => (this.dataList = data));
            }
        }
    };
</script>

<style scoped>
    
    .home {
        margin-top: 40px;
        margin-left: 6%;
        margin-right: 4%;
    }

    .first-col {
        width: 30%;
    }

    /*------------------------------- total paper card -------------------------------*/
    .total-paper-card {
        height: 130px;
        text-align: center;
        margin-bottom: 20px;
    }

    .total-paper-card .word {
        font-size: 20px;
    }

    .total-paper-card .number {
        font-size: 45px;
    }

    /*------------------------------- welcome word card -------------------------------*/
    .welcome-card {
        height: 205px;
        margin-bottom: 20px;
    }

    .welcome-card .el-card__body {
        padding: 0px;
    }

    .welcome-card .welcome-image {
        width: 100%;
        display: block;
    }

    /*------------------------------- Table entity -------------------------------*/


    .small-table-card{
        text-align: center;
        font-size: 20px;
        
        height: 360px;
    }

    .title {
        font-weight: bold;
    }

    .second-col .authors-table-card {
        margin-bottom: 20px;
        height: 355px;
        text-align: center;
        font-size: 20px;
    }

    /*.el-table__row {*/
    /*    font-size: 0.5rem;*/
    /*}*/

    .el-table td, .el-table th {
        padding: 0;
    }

    .third-col .papers-table-card {
        height: 735px;
        text-align: center;
        font-size: 25px;
    }

</style>