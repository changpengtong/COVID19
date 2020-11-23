import $bus from "./bus"
import $axios from "./axios"
import $router from './router'
// import { type } from "os";
// let $echarts = require('echarts');
// import $echarts from 'echarts'

export default {
    author_query: (query) => {
        $bus.update_query_text(query);
        $bus.clean_last_result();
        console.log("author请求!!!")
        $bus.$emit("changeLoading",true)
        $axios.get("/author/"+query)
            .then(res=> {
                console.log(res.data)
                $bus.receive_result(res.data)
            }).catch(err=> {
                console.log("Error" + err);
                $bus.$emit("changeLoading",false)
        })
    },

    bioentity_query: (query) => {
        $bus.update_query_text(query);
        $bus.clean_last_result();
        $bus.$emit("changeLoading",true)
        $axios.get("/bioentity/"+query)
            .then(res=> {
                console.log(res.data)
                $bus.receive_result(res.data)
            }).catch(err=> {
            console.log("Error" + err);
            $bus.$emit("changeLoading",false)
        })
    },

    institution_query: (query) => {
        $bus.update_query_text(query);
        $bus.clean_last_result();
        $bus.$emit("changeLoading",true)
        $axios.get("/institution/"+this.$route.params.id)
            .then(res=> {
                console.log(res.data)
                $bus.receive_result(res.data)
            }).catch(err=> {
            console.log("Error" + err);
            $bus.$emit("changeLoading",false)
        })
    },

    to_display: (query) => {
        $router.push({
            path: "/displayInfo",
            query: {
                query: query
            }
        })
    },
    search_query: (query) => {
        $bus.update_query_text(query);
        $bus.clean_last_result();

        console.log("发送请求");

        // 开始 loading……
        $bus.$emit("changeLoading", true)

        // 向114的 MINA 发送请求
        $axios.get("/search?keyword="+query)
            // 开发环境，请求图的静态数据
            // $axios.get("/api/answer_liushi.json")
            .then(res => {
                console.log("数据已接收");
                console.log(res.data)
                $bus.receive_result(res.data);
                // this.$bus.$emit("changeLoading", false)
            })
            .catch(err => {
                console.log("Error" + err);
                $bus.$emit("changeLoading", false)
            });
    }
}
