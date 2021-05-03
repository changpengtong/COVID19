import $bus from "./bus"
import $axios from "./axios"

export default {
    author_query: (query) => {
        $bus.update_query_text(query);
        $bus.clean_last_result();
        $bus.$emit("changeLoading",true)
        $axios.get("/author/"+query)
            .then(res=> {
                $bus.receive_result(res.data)
            }).catch(
                $bus.$emit("changeLoading",false)
        )
    },

    bioentity_query: (query) => {
        $bus.update_query_text(query);
        $bus.clean_last_result();
        $bus.$emit("changeLoading",true)
        $axios.get("/bioentity/"+query)
            .then(res=> {
                $bus.receive_result(res.data)
            }).catch(
            $bus.$emit("changeLoading",false)
        )
    },

    institution_query: (query) => {
        $bus.update_query_text(query);
        $bus.clean_last_result();
        $bus.$emit("changeLoading",true)
        $axios.get("/institution/"+this.$route.params.id)
            .then(res=> {
                $bus.receive_result(res.data)
            }).catch(
            $bus.$emit("changeLoading",false)
        )
    },
}
