<template>
 <div class="display-visual">
<!--    < class="display-visual">-->
       <linefold></linefold>
        <br/>
        <radar></radar>
        <br/>
        <pie></pie>
        <br/>
        <bar></bar>
        <br/>
        <geo></geo>
        <br>
        <authorbar></authorbar>
        <br/>
 </div>

</template>

<script>
  //  import Vue from 'vue'
    import linefold from "./../graphs/linefold"
    import radar from "./../graphs/radar"
    import pie from "./../graphs/topicPie"
   import geo from "./../graphs/geo"
  import authorbar from "../graphs/authorbar"
  import bar from "../graphs/bar"

    export default {
        name: "DisplayVisual",
        components: {
            "linefold": linefold,
            "radar": radar,
            "pie":pie,
            "authorbar": authorbar,
            "geo":geo,
            "bar":bar,
            //  "back-top": backtop,
            // testbar
        },
    }
</script>

<style scoped>

</style>