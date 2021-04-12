<template>
    <el-card class="box-card" shadow="hover">
      <div class="biograph" id="biograph"></div>
    </el-card>
</template>
<script>
import bioStyle from "../../data/bioentity-style.json";

export default {
  props:['bioentity_data'],
  components: {
  },
  data() {
    return {
      data: [{
        "data": {
          "name": "Not Found",
          "id": "-1",
          "type": "undefined"
        }
      },],
      biostyle: [],
    };
  },
  watch: {
    bioentity_data:function(newData) {
      console.log(newData)
      if (newData != null) {
        this.data = newData;
      }
      this.biostyle = bioStyle;
      var cy = this.$cytoscape({
        container: document.getElementById('biograph'),
          style: this.biostyle,
          elements: this.data,
          layout: { name: 'random' }
        });
      cy.nodes().on('click', (evt) => {
        console.log(evt)
      });
      cy.edges().on('click', (evt) => {
        console.log(evt)
      });   
    }
  },

  methods: {
    handleClick() {
    },
  },

  mounted: function () {
  },
};
</script>
<style>
    .biograph {
      min-height: 300px;
        max-width:450px;
        max-height:450px;
        padding: 0px ;
    }
    .box-card {
        width: 330px;
        margin-top: 2.5rem;
        padding: 0px;
    }
    .el-card__body{
        padding: 0 0 0 0;
    }
</style>