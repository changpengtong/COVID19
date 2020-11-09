<template>
  <div id="dimension-item">
    <div 
      class="title _bd-bottom-blue"
      v-if="data.dimension"
    >{{data.dimension}}
    </div>
    <ul class="content">
      <li 
        v-for="(detail, index) in data.data"
        :key="index"
        class="_bd-bottom-grey"
      >{{detail}}</li>
    </ul>
    <div 
      class="dimension-graph"
      v-if="data.graphName">
      <component 
        :is="graphType"
        :items="getGraphData"
        :keyWord="data.graphName"
        :key="data.graphName"></component>
    </div>
  </div>
</template>

<script>
  export default {
    name: "DimensionItem",
    props: {
      data: {
        type: Object,
        default () {
          return {
            "test": ""
          }
        }
      },
    },
    computed: {
      graphType() {
        let tmpArr = this.data.graphName.split("_");
        return tmpArr[tmpArr.length-1];
      },
      getGraphData() {
        let gName = this.data.graphName;
        return (this.$bus.get_result_graph())[gName];
      }
    },
    components: {
      "pie": () => import("./../../../components/graphs/pie.vue"),
      // "relation": () => import("./../../../components/graphs/graph_relation.vue"),
      // "entity": () => import("./../../../components/graphs/graph_entity.vue"),
      // "linefold": () => import("./../../../components/graphs/graph_linefold.vue"),
      // "bar": () => import("./../../../components/graphs/graph_bar.vue")
    },
    methods: {
      // getGraphData() {
      //   let gName = this.data.graphName;
      //   return (this.$bus.get_result_graph())[gName];
      // }
    }
  }
</script>

<style lang="less" scoped>
ul {
  list-style: none;
  margin: .2rem .1rem;
  padding: 1rem;
}

#dimension-item {
  .title {
    font-size: 1.2rem;
    font-weight: bold;
    text-transform: uppercase;
  }

  .content {
    text-transform: capitalize;

    li {
      margin: 10px auto;
    }
  }
}
</style>