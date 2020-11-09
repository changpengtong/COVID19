<template>
    <div class="wordcloud" id="icloud"></div>
</template>
<script>
    import $search from '../../util/search'
    import $axios from '../../util/axios'
    // import $drawGraphs from "./../util/drawGraphs.js"
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
                $axios.get("/AuthorWordCloudType").then(response=>{
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
                let color = ["orange","blue","yellow","green","red"]
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
                   // console.log(d)
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
                        name: 'WordCloud',
                        type: 'wordCloud',
                        size: ['90%', '90%'],
                        textRotation : [-90, 90],
                        shape: 'circle',
                      //  drawOutOfBound:true,
                        textStyle: {
                            normal:{
                                // color: () => {
                                //     let testcolor = ["rgb(129, 212, 250)","rgb(239, 154, 154)","rgb(128, 203, 196)","rgb(129, 199, 132)","rgb(156, 204, 101)","rgb(212, 225, 87)","rgb(255, 213, 79)","rgb(255, 167, 38)","rgb(255, 110, 64)","rgb(161, 136, 127)","rgb(248, 187, 208)","rgb(128, 203, 196)","rgb(209, 196, 233)"]
                                //     let num = Math.floor(Math.random() * testcolor.length);
                                //     return testcolor[num]
                                // }
                            },
                            emphasis:{
                                shadowBlur:10,
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
                    // $router.push({
                    //     path:"/displayInfo",
                    //     query:{query:params['data']['name'].toLowerCase()}
                    // })
                    let query = params['data']['name'].toLowerCase()
                    $search.to_display(query)
                })
            }
        }
    }
</script>
<style>
    .wordcloud {
        max-width:400px;
        max-height:400px;
    }
</style>