<template>
    <div class="topnav">
        <div class="_container">
            <div class="framework-width">
                <div id="icon">
                    <router-link to="/" tag="div">
                        <a href="#">COVID-19 PORTAL</a>
                    </router-link>
                </div>
                <router-link to="/author" tag="div">
                    <a class="link" href="#">ABOUT</a>
                </router-link>
                <router-link to="/" tag="div">
                    <a class="link" href="#">HOME</a>
                </router-link>
                <transition name="el-zoom-in-top">
                <search-box class="search-box" v-if="ifSearchBox"></search-box>
                </transition>
            </div>
        </div>
    </div>
</template>

<script>
    import SearchBox from "./../common/SearchBox.vue";
    export default {
        component: {
            "search-box":SearchBox
        },
        computed: {
            ifSearchBox() {
                return this.$route.name=="displayInfo"
            },
            routeName() {
                return this.$route.name;
            }
        }
    }

</script>

<style>
    * {
        box-sizing: border-box;
    }

    /*
    导航栏样式
     */
    .topnav {
        overflow: hidden;
        /* background-color: #e9e9e9; */
        background-color: #ffffff;
        /* box-shadow: 0px 1px 2px #aaaaaa; */
        /* position: fixed; */
        border-bottom: 1px solid #e1e1e1;
        top: 0;
        width: 100%;
    }

    /*
     * ---------------------------------------------
     *  暂时的网站标识
     */
    #icon {
        float: left;
        display: block;
        line-height: 80px;
        font-weight: bold;
    }
    #icon a {
        font-size: 1.7em;
        padding-left: .5em;
    }
    /*
     * ---------------------------------------------
    */

    .topnav a,
    .topnav .link {
        float: right;
        display: inline-block;
        color: #409EFF;
        font-weight: bold;
        text-decoration: none;
        text-align: center;
        padding: 20px 16px;
        font-size: 1em;
        line-height: 40px;
        border-bottom: 4px solid transparent;
    }

    .topnav .link:hover {
        /* border-bottom: 4px solid #409EFF; */
        border-bottom-color: #409EFF;
    }

    .topnav .search-box {
        float: right;
        display: inline-block;
        margin: 10px 16px;
        padding: 10px 0;
        /* height: 25px; */
        line-height: 40px;
    }
</style>