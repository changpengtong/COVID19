<template>
    <div class="footer _bg-grey">
        <div class="_container">
            <span>&copy;<router-link to="/" tag="a">Covid-19 Portal</router-link>. All Rights Reserved </span>
            <!-- <router-link to="/about" tag="a" class="item">contact us</router-link> -->
        </div>
    </div>
</template>

<script>
    export default {
        name: "Footer"
    }
</script>

<style>
    .footer {
        text-transform: capitalize;
        text-align: center;
    }
    .item {
        margin: 0 .5em;
    }
</style>