<template>
    <div>
        <div class="graph" id="chart4"></div>
    </div>
</template>
<script>
    import $search from "../../util/search.js"
    import $axios from "../../util/axios";
    // import echarts from 'vue-echarts/components/ECharts'
    let echarts = require('echarts');
    // import $axios from"./../../../util/axios"
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
                $axios.get("/InstitutionPaperByYear").then(response=>{
                    d=response.data
                    let group=[]
                    let value=[]
                    for(let i=d.length-11;i<d.length;i++) {
                        group.push(d[i]['PubYear'])
                        value.push(d[i]['Number_of_Papers'])
                    }
                    // val = d[0]['value']/10;
                    this.draw_bar(group,value)
                })
            },
            draw_bar(group,value) {
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
                        group:group[i],
                        value:value[i],
                    })
                }
                console.log(seriesData)
                let myChart = echarts.init(chart);
                const option = {
                    color:['#686868'],
                    tooltip : {
                        trigger: 'axis',
                        axisPointer : {            // 坐标轴指示器，坐标轴触发有效
                            type : 'line'        // 默认为直线，可选为：'line' | 'shadow'
                        },
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
                            data:group,//["2013", "2014", "2015", "2016", "2017", "2018", "2019", "2020"],
                            //data : group,
                            axisTick: {
                                alignWithLabel: true
                            },
                            axisLabel: {
                                interval:0,
                                rotate:45
                            }
                        }
                    ],
                    yAxis : [
                        {
                            type : 'value',
                            position:'right',
                            axisLine: {
                                show:false
                            }
                        }
                    ],
                    series:[
                        {
                            name:'Paper',
                            type:'bar',
                            barWidth: '40%',
                            data:seriesData//[119,139,137,205,219,253,389,289]
                        }
                    ]
                };
                myChart.setOption(option);
                // 当多个图形渲染时，调用以下方法只能实现最后一个图的自适应效果
                // window.onresize = function(){
                //     console.log("window resize!")
                //     initial();
                //     myChart.resize();
                //     console.log("myChart resize!")
                // }
                // 生成多个图时实现全部都能自适应。参考博客：https://blog.csdn.net/qq_25816185/article/details/82414529
                window.addEventListener('resize', function () {
                    initial();
                    myChart.resize()
                })
                // 这种写法无法实现自适应，不知道为什么
                // window.onresize = ("resize",function() {
                //     initial();
                //     echart.resize()
                // })
                myChart.on("click",function(params){
                    console.log(params);
                    // window.open("https://www.baidu.com/baidu?wd="+params.seriesName);
                    // $router.push({
                    //     path:"/displayInfo",
                    //     query:{query:params.data.name.toLowerCase()}
                    // })
                    $search.to_display(params.data.name.toLowerCase())
                })
            }
        }
    }
</script>
<style>
</style>