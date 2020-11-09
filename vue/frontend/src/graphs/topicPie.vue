<template>
    <div>
        <el-tabs type="border-card">
            <el-tab-pane label="AllTab">
                <div class="graph" id="chart3"></div>
            </el-tab-pane>
        </el-tabs>
    </div>
</template>
<script>
    import $search from "./../util/search.js"
    import $axios from "./../util/axios"
    let echarts = require('echarts');
    export default {
        name: "pie",
        data() {
            return {
                // tabItem:"Paper"
            }
        },
        mounted() {
            this.start_draw();
        },
        methods: {
            start_draw() {
                let list=[]
                var legendData = [];
                var seriesData = [];
                var selected = {};
                var name;
                $axios.get("/topk").then(response=> {
                    //   console.log(response)
                    // allData = response.data
                    // console.log(allData)
                    // let list = d
                    list=response.data
                    // console.log(list)
                    // name.push(nameList[Math.round(Math.random() * nameList.length - 1)]);
                    for(let i=0;i<list.length;i++) {
                        // nameList.push(list[i]['nameList'])
                        name=list[i]['text']
                        // y1.push(list[i]['y'])
                        legendData.push(name)
                        // console.log("name",name)
                        // console.log("legendData",legendData)
                        seriesData.push({
                            name: name,
                            value: list[i]['size']
                        });
                        selected[name] = i < 11;
                        // console.log("selected", selected[name])
                    }
                    // console.log(x1)
                    // console.log(legendData)
                    this.draw_pie(legendData,seriesData)
                })

            },
            draw_pie(legendData,seriesData) {
                let chart = document.getElementById("chart3");
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
                let myChart = echarts.init(chart);
                const option = {
                    title: {
                        text: "Related topic",
                        y: 'top'
                    },
                    tooltip: {
                        trigger: 'item',
                        formatter: "{b} : {c} ({d}%)"
                    },
                    legend: {

                        type: 'scroll', // 图例过多加上滚动条
                        //   orient: 'vertical',
                        show:true,
                        //   right: 10,
                        top:"10%",
                        //   bottom: 0,
                        data: legendData,
                        // selected: data.selected
                    },
                    color: ["rgb(255, 138, 128)","rgb(225, 190, 231)","rgb(130, 177, 255)","rgb(128, 222, 234)","rgb(129, 199, 132)","rgb(255, 183, 77)","rgb(245, 124, 0)","rgb(248, 187, 208)"],
                    calculable: true,
                    series: [
                        {
                            type: 'pie',
                            radius: '55%',
                            center: ['50%', '65%'],
                            itemStyle: {
                                emphasis: {
                                    label: {
                                        show: true,
                                        textStyle: {
                                            fontWeight: 'bold'
                                        },
                                        shadowBlur: 10,
                                        shadowOffsetX: 0,
                                        shadowColor: 'rgba(0, 0, 0, 0.5)'
                                    }
                                }
                            },
                            data: seriesData,
                            roseType: 'radius'  // 玫瑰图
                        }
                    ]
                };
                myChart.setOption(option);
                // 生成多个图时实现全部都能自适应。参考博客：https://blog.csdn.net/qq_25816185/article/details/82414529
                window.addEventListener('resize', function () {
                    initial();
                    myChart.resize()
                })
                myChart.on("click",function(params){
                    let query = params.data.name.toLowerCase()
                    $search.to_display(query);
                })
            }
        }
    }
</script>
<style>
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