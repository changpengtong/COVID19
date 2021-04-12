<template>
    <div>
        <el-tabs  type="border-card">
            <el-tab-pane label="Journal">
                <div class="graph" id="chart4">
                </div>
            </el-tab-pane>
        </el-tabs>
    </div>
</template>
<script>
    import $search from "./../util/search.js"
    let echarts = require('echarts');
    import $axios from"./../util/axios"
    export default {
        name: "bar",
        data() {
            return {
            }
        },
        mounted() {
            this.start_draw();
        },
        methods: {
            start_draw() {
                let d=[]
                $axios.get("/bar").then(response=>{
                    d=response.data
                    let group=[]
                    let allgroup=[]
                    let value=[]
                    let size=5
                    for(let i=0;i<size;i++) {
                        if(d[i]['group']==="") {
                            size++;
                            continue;
                        }
                        allgroup.push(d[i]['group'])
                        var ch = d[i]['group'].split(" ")
                        if(ch.length>1) {
                            d[i]['group'] = ch[0]
                        }
                        group.push(d[i]['group'].charAt(0).toUpperCase()+d[i]['group'].slice(1))
                        value.push(d[i]['value']/181778)
                    }
                    this.draw_bar(group,value,allgroup)
                })
            },
            draw_bar(group,value,allgroup) {
                let chart = document.getElementById("chart4");
                var initial = function () {
                    let chartParent = chart.parentNode.parentNode;
                    chart.style.width = chartParent.clientWidth - 30 + "px";
                    chart.style.margin = "0 auto";
                    chart.style.height = (chartParent.clientWidth *0.75 - 30) + "px";
                }
                initial();
                let seriesData = []
                for(let i=0;i<group.length;i++) {
                    seriesData.push({
                        name:allgroup[i],
                        value:value[i],
                    })
                }
                let myChart = echarts.init(chart);
                const option = {
                    title:{
                        text:"Article related/ total articles",
                        top:"0%",
                    },
                    tooltip : {
                        trigger: 'axis',
                        axisPointer : {      
                            type : 'line'      
                        },
                        formatter:function (params) {
                            return "Journal : "+params[0]['data']['name']+
                                "<br/>"+"Value: "+params[0]['data']['value']
                        }
                    },
                    grid: {
                        left: '3%',
                        right: '4%',
                        bottom: '3%',
                        containLabel: true
                    },
                    xAxis : [
                        {
                            type : 'category',
                            data : group,
                             axisTick: {
                                 alignWithLabel: true
                             },
                            inverse: false, 
                            axisLabel: {
                                textStyle: {
                                }
                            }
                        }
                    ],
                    yAxis : [
                        {
                            type : 'value',
                        }
                    ],
                    series : [
                        {
                            name:['test'],
                            type:'bar',
                            barWidth: '60%',
                            data:seriesData,
                            itemStyle: {
                                normal: {
                                    color: () => {
                                        let testcolor = ["rgb(129, 212, 250)","rgb(239, 154, 154)","rgb(128, 203, 196)","rgb(129, 199, 132)","rgb(156, 204, 101)","rgb(212, 225, 87)","rgb(255, 213, 79)","rgb(255, 167, 38)","rgb(255, 110, 64)","rgb(161, 136, 127)","rgb(248, 187, 208)","rgb(128, 203, 196)","rgb(209, 196, 233)"]
                                        let num = Math.floor(Math.random() * testcolor.length);
                                        return testcolor[num]
                                    }
                                }
                            }
                        }
                    ]
                };
                myChart.setOption(option);

                window.addEventListener('resize', function () {
                    initial();
                    myChart.resize()
                })
                myChart.on("click",function(params){
                    $search.to_display(params.data.name.toLowerCase())
                })
            }
        }
    }
</script>
<style>
</style>