<template>
    <div>
        <div class="graph" id="chart4"></div>
    </div>
</template>
<script>
    import $search from "../../util/search"
    let echarts = require('echarts');
    export default {
        props: ['bar'],
        data() {
            return {
                d:[]
            }
        },
        mounted() {
        },
        watch: {
            bar:function(newData) {
                this.d = newData
                this.start_draw()
            }
        },
        methods: {
            start_draw() {
                let group=[]
                let value=[]
                for(let i=this.d.length-1;i>=0;i--) {
                    group.push(this.d[i]['PubYear'])
                    value.push(parseInt(this.d[i]['NumberOfPapers']))
                }
                this.draw_bar(group,value)

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
                        axisPointer : {       
                            type : 'line'     
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
                            data:group,
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
                            data:seriesData
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