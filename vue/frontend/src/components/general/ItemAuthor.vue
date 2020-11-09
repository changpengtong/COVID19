<template>
    <div class="display-view container">
        <div v-for="(card,index) in authors" :key="index" is="author-card" v-bind:card="card">
<!--            <author-card :card="card"></author-card>-->
        </div>
    </div>
    
</template>

<script>
    import AuthorCard from "../general/AuthorCard";
    export default {
        name: "ItemAuthor",
        data() {
            return {
                authors:[{
                    'ForeName':'a',
                    'LastName':'b',
                    'lab':'COVID Lab',
                    'url':'google.com',
                    'institution':'msu',
                    'mail':'12345@qq.com',
                    'paper':123,
                    'quotes':14,
                },{
                    'ForeName':'c',
                    'LastName': 'd',
                    'lab':'FOVID Lab',
                    'url':'google.com',
                    'institution':'msu',
                    'mail':'67890@qq.com',
                    'paper':12222,
                    'quotes':145,
                }],
            }
        },
        components: {
            'author-card': AuthorCard,
        }
    }
</script>

<style scoped>
    .container {
        max-width: 1200px;
        width: 100%;
        margin: 0 auto;
        /*padding: .5em;*/
    }
    .display-view {
        width: 100%;
        /* max-width: 970px; */
        margin: 0 auto;
    }

</style>