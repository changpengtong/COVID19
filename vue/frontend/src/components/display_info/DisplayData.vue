<template>
    <div class="data-board">
        <el-tabs value="papers">
            <el-tab-pane name="papers" label="PAPER">
                <!-- <span slot="label"><i class="el-icon-document"></i> PAPER</span> -->
                <!-- 展示论文 -->
                <div class="show-pane">
                    <item-paper></item-paper>
                </div>
            </el-tab-pane>
        </el-tabs>

    </div>
</template>

<script>
    /* eslint-disable vue/no-unused-components */
    import textExpandShrink from "../textExpandShrink"
    import ItemPaper from "./display_item/ItemPaper";

    export default {
        name: "DisplayData",
        components: {
            "text-expand-shrink": textExpandShrink,
            "item-paper": ItemPaper,
        },
        data() {
            return {
            }
        },
        computed: {
            result() {
                return this.$bus.get_result();
            },
        },
        methods: {
        }
    }
</script>

<style>
    /* override------------------------------ */
    /* .el-tabs__nav-wrap::after, */
    .el-tabs__active-bar {
        height: 4px;
    }
    /* -------------------------------------- */
    .show-pane {
        margin: 5px auto 20px auto;
    }

    .data-board {
        margin: 0 auto;
        padding: 5px;
    }

    .data-item {
        margin: 0.7em;
        margin-bottom: 1em;
        padding-bottom: 1em;
        border-bottom: 0.1em solid rgba(150, 150, 150, 0.2)
    }

    .author-span-margin {
        margin: 0.2em;
    }
</style>