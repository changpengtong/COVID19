<template>
    <div>
        <el-tabs type="border-card">
            <el-tab-pane label="AllTab">
                <div class="graph" id="chart1"></div>
            </el-tab-pane>
        </el-tabs>
    </div>
</template>
<script>
    import $search from "./../util/search.js"
    import $axios from "./../util/axios"
    let echarts = require('echarts');
    export default {
        name: "linefold",
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
                let x1=[]
                let y1=[]
                let x2=[]
                let y2=[]
                $axios.get("/linefold").then(response=> {
                    //   console.log(response)
                    d = response.data
                    console.log(d)
                    let list = d
                    var date =null;
                    var cnt=0;
                    for(let i=0;i<list.length;i++) {
                        x1.push(list[i]['x'].slice(0,7))
                        y1.push(list[i]['y'])
                    }
                    for(let i=0;i<x1.length;i++) {
                        if(date==null) {
                            date = x1[i]
                            x2.push(date)
                        }
                        if(date===x1[i]) {
                            cnt+=y1[i]
                        } else if(date!==x1[i]){
                            y2.push(cnt)
                            date = x1[i]
                            x2.push(date)
                            cnt=y1[i]
                        }
                    }
                    // console.log(x1)
                    // console.log(y1)
                    this.draw_linefold(x2,y2);
                })
            },
            draw_linefold(x2,y2) {
                let chart = document.getElementById("chart1");

                // 给echart初始化宽度，替代系统的初始值100px。参考博客：https://unordered.org/timelines/5a0c74067b000000
                var initial = function () {
                    // 根据想要的效果调整的公式，不是固定的
                    // chart.style.width = (window.innerWidth/4)+80+"px";

                    // 获取 chart 的父父结点chartParent，再根据其算宽度
                    let chartParent = chart.parentNode.parentNode;
                    chart.style.width = chartParent.clientWidth - 30 + "px";
                    chart.style.margin = "0 auto";
                    chart.style.height = (chartParent.clientWidth *0.75 - 30) + "px";
                }
                initial();

                // 初始化图的宽度
                // chart.style.width = this.initial_style(chart);
                let myChart = echarts.init(chart);
                const option = {
                    title: {
                        text: "2020 Paper Trend"
                    },
                    tooltip: {
                        trigger: 'axis'
                    },
                    legend: {
                        // data:['cnn','rnn','lstm','svm','gbdt','perception']
                        data: ['line'],
                        y:"10%"
                    },
                    grid: {
                        left: '0%',
                        right: '3%',
                        bottom: '0%',
                        top:"18%",
                        containLabel: true
                    },
                    toolbox: {
                        feature: {
                            dataZoom: {
                                title: {
                                    zoom: 'Zoom in',
                                    back: 'Zoom out'
                                }
                            },
                            saveAsImage: {
                                title: "Save Image"
                            },
                        },
                    },
                    xAxis: {
                        type: 'category',
                        boundaryGap: false,
                        data: x2,
                        //data: ['1'],
                        axisLabel:{
                            interval:0,
                            rotate:60
                        }
                    },
                    yAxis: {
                        type: 'value'
                        // type:'log'
                    },
                    series:[{
                        data:y2,
                       // data:['123'],
                        type:'line',
                        smooth:true
                    }] ,
                    // color: ["#60acfc","#32d3eb","#5bc49f","#feb64d","#ff7c7c","#9287e7"]
                    color: ["#60acfc","#5bc49f","#9287e7"]
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
                    // chart.style.width = this.initial_style(chart);
                    myChart.resize()
                })
                // 这种写法无法实现自适应，不知道为什么
                // window.onresize = ("resize",function() {
                //     initial();
                //     echart.resize()
                // })
                myChart.on("click",function(params){
                    // console.log(params);
                    // window.open("https://www.baidu.com/baidu?wd="+params.seriesName);
                    // $router.push({
                    //     path:"/displayInfo",
                    //     query:{query:params['seriesName'].toLowerCase()}
                    // })
                    let query = params['seriesName'].toLowerCase();
                    $search.to_display(query);
                })
            }
        }
    }
</script>

<style scoped>
    .graph {
        width: 100%;
        min-height: 300px;
        /* height: 400px; */
    }
    .tabs {
        box-shadow:0 2px 4px 0 rgba(0,0,0,0.1), 0 3px 10px 0 rgba(0,0,0,0.15);
        padding:10px;
        /* border-radius: 5px; */
    }
</style>