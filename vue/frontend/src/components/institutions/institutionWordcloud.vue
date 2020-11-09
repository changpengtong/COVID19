<template>
    <el-card class="box-card" shadow="hover">
    <div class="wordcloud" id="icloud"></div>
        </el-card>
</template>
<script>
    import $search from '../../util/search'
    import $axios from '../../util/axios'
    require('echarts-wordcloud');
    let echarts = require('echarts');
    export default {
        data() {
            return {
            }
        },
        mounted() {
            this.start_draw();
        },
        methods: {
            start_draw() {
                let id = 'icloud'
                let d=[]
                $axios.get("/InstitutionWordCloud").then(response=>{
                    d=response.data
                    this.draw_wordcloud(id,d);
                })
            },
            draw_wordcloud(id,data) {
                let chart = document.getElementById(id);
                let myChart = echarts.init(chart);
               // let d=[]
                 console.log(data)
                let d = []
                let index=0
               // let seriesData = []
                let color = ["#577590","#43AA8B","#E76F51","#F4A261","#E9C46A"]
              for(let key in data) {
                    let tempData = data[key]
                    console.log(key)
                    console.log(tempData)
                    for(let i=0;i<tempData.length;i++) {
                        var map = {}
                        map['name'] = tempData[i]['Mention']
                        map['value'] = tempData[i]['occurences']
                        map['textStyle'] = {
                            normal:{
                                color: color[index]
                            },
                            emphasis:{
                                shadowBlur:10,
                                shadowColor:'#999'
                            }
                        }
                        d.push(map)
                    }
                    index++;
                }
                 console.log(d)
                const option = {
                    tooltip: {
                        trigger: 'item',
                        // formatter:
                    },
                    toolbox: {
                        show: false,
                        orient: 'vertical',
                        x: 'right',
                        y: 'center',
                    },
                    series: [{
                        left: 'center',
                        top: 'center',
                        width: '70%',
                        height: '80%',
                        right: null,
                        bottom: null,
                        name: 'WordCloud',
                        type: 'wordCloud',
                        // size: ['90%', '90%'],
                        textRotation : [-90, 90],
                        rotationStep: 30,
                        // the larger the grid size, the bigger the gap between words.

                        gridSize: 12,
                        shape: 'circle',
                      //  drawOutOfBound:true,
                        textStyle: {
                            normal:{
                                fontWeight: 'bold',
                            },
                            emphasis:{
                                fontWeight:'bolder',
                                fontSize:30,
                                shadowBlur:55,
                                shadowColor:'#999'
                            }
                        },
                        autoSize: {
                            enable: true,
                            minSize: 8
                        },
                        data: d
                    }],
                };
                myChart.setOption(option);
                window.addEventListener('resize', function () {
                    // initial();
                    myChart.resize()
                })
                myChart.on("click",function(params){
                    let query = params['data']['name'].toLowerCase()
                    $search.to_display(query)
                })
            }
        }
    }
</script>
<style>
    .wordcloud {
        max-width:450px;
        max-height:450px;
        padding: 0px ;
    }
    .box-card {
        width: 330px;
        margin-top: 2.5rem;
        padding: 0px;
    }
    .el-card__body{
        padding: 0 0 0 0;
    }
</style>