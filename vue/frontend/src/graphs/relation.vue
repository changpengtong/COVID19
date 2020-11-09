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
    import $graphs from "./../util/drawGraphs.js"
    let echarts = require('echarts');
    export default {
        name: "relation",
        data() {
            return {
                // graph_title: "专家合作网络"
            }
        },
        mounted() {
            this.start_draw();
        },
        methods: {
            start_draw() {
                this.draw_relation();
            },
            draw_relation() {
                let graph_title = graph_data.title;
                let data = graph_data.data;
                // 专家信息
                let graph_detail = graph_data.detail;
                // 初始化
                let chart = document.getElementById("chart1");
                var initial = function () {
                    let chartParent = chart.parentNode.parentNode;
                    chart.style.width = chartParent.clientWidth - 30 + "px";
                    chart.style.margin = "0 auto";
                    chart.style.height = (chartParent.clientWidth *0.75 - 30) + "px";
                    // chart.style.height = (chartParent.clientHeight) + "px";
                }
                initial();
                // 数据中点的个数
                let nodes = data.node;
                let node_num = nodes.length;
                // 色盘
                // let colors = ["#D53A35", "#334B5C", "#D48265", "#9FDABF","#8E8DBE", "#7A306C", "#7FAE90", "#DE9325", "#CFB2A9", "#797B7F"]
                // let colors = [ "rgba(90, 125, 124, 1)", "#12569B","#DE9325", "#CFB2A9", "#797B7F"]
                let colors = ["rgba(255,99,132,0.8)",  "rgba(75,192,192,0.8)", "rgba(153,102,255,0.8)",  "rgba(255,206,86,0.8)", "rgba(54,162,235,0.8)","rgba(255,159,64,0.8)"]
                // let colors = ["rgba(93, 97, 140, 1)", "rgba(242, 153, 68, 1)", "rgba(146, 94, 136, 1)", "rgba(141, 161, 185, 1)", "rgba(153, 151, 174, 1)"]
                // echarts 实例初始化
                let myChart = echarts.init(chart);
                // 计算节点大小
                let min_size = myChart.getWidth() / node_num * 0.8;
                let max_size = myChart.getWidth() / node_num * 4;
                let minus = max_size - min_size;
                let option = {
                    title: {
                        text: graph_title
                    },
                    animationDurationUpdate: 1500,
                    animationEasingUpdate: 'quinticInOut',
                    toolbox: {
                        show: true,
                        feature: {
                            saveAsImage: {
                                show: true,
                                type: "png",
                                pixelRatio: 3
                            },
                            restore: {
                                show: true,
                            }
                        }
                    },
                    tooltip:{
                        show: true
                    },
                    series: [{
                        type: 'graph',
                        // --------------------------
                        // 布局类型
                        // layout: "none",
                        // layout: 'circular',
                        // circular: {
                        //   rotateLabel: true,
                        // },
                        layout: 'force',
                        force: {
                            initLayout: "circular",
                            // 斥力
                            repulsion: 300,
                            gravity: 0.25,
                            // edgeLength: [(myChart.getWidth()/node_num), (myChart.getWidth() / node_num +200)]
                            edgeLength: [(myChart.getWidth()/12)+min_size, (myChart.getWidth()/8)+min_size]
                        },
                        // --------------------------
                        // 位置
                        top: "15%",
                        bottom: "10%",
                        right: "15%",
                        left: "15%",
                        progressiveThreshold: 500,
                        roam: true,
                        focusNodeAdjacency: true,
                        // 点的数据  ---------------------------------
                        data: data.node.map(function (node, index) {
                            let symbolSize = min_size + (node.value / 100 * minus);
                            let color_index = ((node.value/10) > colors.length)? ((node.value/10)%colors.length) : (node.value/10);
                            color_index = parseInt(color_index)
                            return {
                                name: node.name,
                                value: node.value,
                                symbolSize: symbolSize,
                                // symbolSize: node_num>20 ? (chart.parentNode.parentNode.clientWidth / 17) : (chart.parentNode.parentNode.clientWidth / 10),
                                itemStyle: {
                                    color: colors[color_index],
                                    opacity: 0.9
                                },
                                tooltip: {
                                    position: ["1%", "8%"],
                                    formatter: function (params) {
                                        if (params.dataType === "edge") {
                                            return params.data.source + " + " + params.data.target;
                                        } else if (params.dataType === "node") {
                                            return $graphs.get_detail(node.name, graph_detail)
                                        }
                                    }
                                }
                            };
                        }),
                        // 边的数据  ---------------------------------
                        links: data.link.map(function (link) {
                            return {
                                source: link.source,
                                target: link.target
                            };
                        }),
                        // 边两端标记
                        edgeSymbol: ['none', 'none'],
                        edgeSymbolSize: 5,
                        label: {
                            show: true,
                            position: "right",
                            color: "#000"
                        },
                        lineStyle: {
                            width: 1,
                            curveness: 0.3,
                            opacity: 0.9
                        },
                        // 强调样式
                        emphasis: {
                            label: {
                                show: true,
                                position: 'inside',
                                fontSize: 18,
                                fontWeight: "bold",
                                color: "#000",
                                textBorderColor: "#fff",
                                textBorderWidth: 2
                            },
                            itemStyle: {
                                shadowColor: 'rgba(0, 0, 0, 0.5)',
                                shadowBlur: 5,
                                opacity: 1
                            },
                            lineStyle: {
                                width: 1.5,
                                color: "#000",
                                opacity: 1,
                                // shadowColor: 'rgba(0, 0, 0, 0.7)',
                                // shadowBlur: 5,
                            },
                            edgeLabel: {
                                show: false,
                                color: "#000",
                                // formatter: function (params) {
                                //   return params.data.source + "-->" + params.data.target;
                                // }
                            }
                        }
                    }]
                };
                myChart.setOption(option);
                window.addEventListener('resize', function () {
                    initial();
                    // chart.style.width = this.initial_style(chart);
                    myChart.resize()
                })
                myChart.on("click", function (params) {
                    if (params.dataType === "node") {
                        // $router.push({
                        //   path: "/displayInfo",
                        //   query: {query: params.data.name}
                        // })
                        $search.to_display(params.data.name)
                    }
                })
            }
        }
    }
</script>

<style scoped>
    .graph {
        width: 100%;
        min-height: 300px;
        /* max-height: 400px; */
    }
    .tabs {
        box-shadow: 0 2px 4px 0 rgba(0, 0, 0, 0.1), 0 3px 10px 0 rgba(0, 0, 0, 0.15);
        padding: 10px;
        /* border-radius: 5px; */
    }
</style>