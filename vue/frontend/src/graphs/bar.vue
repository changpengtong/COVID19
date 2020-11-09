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
   // import echarts from 'vue-echarts/components/ECharts'
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
                  //  let val =0;
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
                   // val = d[0]['value']/10;
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
                        // itemStyle: {
                        //     normal: {
                        //         color: () => {
                        //             let testcolor = ["rgb(129, 212, 250)","rgb(239, 154, 154)","rgb(128, 203, 196)","rgb(129, 199, 132)","rgb(156, 204, 101)","rgb(212, 225, 87)","rgb(255, 213, 79)","rgb(255, 167, 38)","rgb(255, 110, 64)","rgb(161, 136, 127)","rgb(248, 187, 208)","rgb(128, 203, 196)","rgb(209, 196, 233)"]
                        //             let num = Math.floor(Math.random() * testcolor.length);
                        //             return testcolor[num]
                        //         }
                        //     }
                        // }
                    })
                }
                let myChart = echarts.init(chart);
                const option = {
                    title:{
                        text:"Article related/ total articles",
                        // left:'center',
                        top:"0%",
                    },
                    tooltip : {
                        trigger: 'axis',
                        axisPointer : {            // 坐标轴指示器，坐标轴触发有效
                            type : 'line'        // 默认为直线，可选为：'line' | 'shadow'
                        },
                        formatter:function (params) {
                           // console.log(params)
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
                           // data:["map", "lines", "bar", "line", "pie", "scatter", "candlestick", "radar", "heatmap", "treemap"],
                            data : group,
                             axisTick: {
                                 alignWithLabel: true
                             },
                            inverse: false,  // 反向排列 默认升序 和sortNumber方法配合使用
                            axisLabel: {
                               // rotate:45,
                                textStyle: {
                                   // fontSize:'16'
                                }
                            }
                        }
                    ],
                    yAxis : [
                        {
                            type : 'value',
                        }
                    ],
                  //  series:seriesData
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
                          //  data: [120, 200, 150, 80, 70, 110, 130]
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