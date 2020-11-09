<template>
    <div>
        <el-tabs type="border-card">
            <el-tab-pane label="AllTab">
                <div class="graph" id="chart2"></div>
            </el-tab-pane>
        </el-tabs>
    </div>
</template>
<script>
   // import $search from "./../util/search.js"
    import $axios from "./../util/axios"
  //   import 'echarts/lib/component/legend'
    let echarts = require('echarts');
    export default {
        name: "radar",
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
                let d=[]
                let value=[]
                let name=[]
                let d1=[]
                $axios.get("/radar").then(response=> {
                    //   console.log(response)
                    d = response.data
                    console.log(d)
                    let list = d
                    for (let i = 0; i < list.length; i++) {
                        value.push(list[i]['value'])
                        name.push(list[i]['group'])
                        var m = {}
                        m['text'] = list[i]['group']
                        m['max'] = 110000
                        d1[i] = m
                    }
                    // console.log(value)
                    // console.log(name)
                    // console.log(d1)
                    this.draw_radar(name,value,d1)
                })
            },
            draw_radar(name,value,d1) {
                let chart = document.getElementById("chart2");
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

                let series=[]
                for (let i=0; i<1; i++) {
                    //console.log(seriesData[i]);
                    series.push({
                        name:'data',
                        type: 'radar',
                        symbol: 'none', //拐角点是否显示。默认显示圆形的点 | "rect"显示方形点
                        symbolSize:5,  //拐角点的大小，等于0不显示
                        lineStyle: {
                            width: 1 //数据线的宽度
                        },
                        emphasis: {
                            areaStyle: {
                                // color: 'rgba(0,250,0,0.3)', //鼠标移动到数据区域内时显示的颜色
                            }
                        },
                        areaStyle:{
                            normal:{
                                opacity:0.3 //数据区域的透明度，范围0~1
                            }
                        },
                        data:[{
                            value:value,
                            name: 'item',
                            label: { //数据label，显示点的值。如果symbol没有显示，这个不显示
                                normal: {
                                    show:true,
                                    formatter:function(params) {
                                        return params.value;
                                    }
                                }
                            },
                        }]
                    });
                }
                name[0]='item'
                console.log(series);
                const option = {
                    title: {
                        text: "Top 5 Sources",
                        top: 10,
                        left: 10
                    },
                    tooltip: {
                        trigger: 'item',
                        backgroundColor : 'rgba(0,0,250,0.2)'
                    },
                    legend: {
                        type: 'scroll',
                        bottom: 10,
                        data: name,
                    },
                    radar: {
                        indicator:d1,
                    },
                    triggerEvent:true,  //配置文字绑定点击事件！
                                        //雷达图点击事件参考：https://www.cnblogs.com/sk-3/p/6673598.html
                    series : series, //具体格式见上方series的处理过程
                    color:["#60acfc","#32d3eb","#5bc49f","#feb64d","#ff7c7c","#9287e7"]
                   // color: ["#60acfc"] //每一个数据区域的颜色，循环显示
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
                // myChart.on("click",function(params){
                //     // console.log(params);
                //     // window.open("https://www.baidu.com/baidu?wd="+params.seriesName);
                //     if(params['componentType'] == 'radar'){
                //         // $router.push({
                //         //     path:"/displayInfo",
                //         //     query:{query:params['name'].toLowerCase()}
                //         // })
                //         $search.to_display(params['name'].toLowerCase())
                //     }
                //     if(params['componentType'] == 'series'){
                //         // $router.push({
                //         //     path:"/displayInfo",
                //         //     query:{query:params['data']['name'].toLowerCase()}
                //         // })
                //         $search.to_display(params['data']['name'].toLowerCase())
                //     }
                // })
            },
        }
    }
</script>
<style>
</style>