import Vue from 'vue';

export default new Vue({
    methods: {
        get_query_text() {
            return this.query_data.query_text;
        },
        return_testdata() {
            return this.test_data;
        },
        update_query_text(val) {
            if (val !== "") {
                this.query_data.query_text = val;
            }
        },
        receive_result(val) {
            if(val === "   [delete] Deleting directory /var/www/html/ai/MINA_STEP1/classes/main") {
                console.log("error: 114 server is not running!");
                this.$notify.error({
                    title: 'ERROR',
                    message: "The system is under maintenance",
                    offset: 100
                });
            }
            else if(val === "114Wrongnull") {
                console.log(val);
                this.$notify.error({
                    title: 'ERROR',
                    message: "The system is under maintenance",
                    offset: 100,
                    duration: 0
                });
            }
            else {
                this.query_data.query_result = val;
            }
            this.$bus.$emit("changeLoading", false)
        },
        get_result() {
            return this.query_data.query_result;
        },
        get_result_graph() {
            return this.query_data.query_result;
        },
        clean_last_result() {
            this.query_data.query_result = "";
        },
        highlight_query(data) {
            let query = this.get_query_text().toUpperCase();
            let replaceReg = new RegExp(query, "gi");
            let replaceStr = "<span style=\"text-decoration:underline;display:inline-block;margin:0 .1em;color:#409EFF;\">" + query + "</span>";
            for(let type in data) {
                if(type === "visual") continue;
                for(let item in data[type]) {
                    for(let item_data in data[type][item]) {
                        let info = data[type][item][item_data]
                        if(typeof(info) === "string") {
                            data[type][item][item_data] = data[type][item][item_data].replace(replaceReg, replaceStr);
                        }
                        else if(typeof(info) === "object") {
                            data[type][item][item_data].forEach(e => {
                                e.replace(replaceReg, replaceStr);
                            });
                        }
                    }
                }
            }
        }
    },
    data() {
        return {
            data: "origin data",
            query_data: {
                query_text: "",
                type: "",
                seq: "",
                query_result: {}
            },
        }
    }
});
