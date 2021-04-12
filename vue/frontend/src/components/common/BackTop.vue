<template>
    <div>
        <el-button
                :style='scrollToTopStyle'
                v-show='showScrollTop'
                @click="scrollToTop"
                icon="el-icon-arrow-up"
                type="primary"
                plain
                circle>
        </el-button>
    </div>
</template>

<script>
    export default {
        props: {
            top: {
                type: [Number, String],
                default: undefined
            },
            bottom: {
                type: [Number, String],
                default: undefined
            },
            left: {
                type: [Number, String],
                default: undefined
            },
            right: {
                type: [Number, String],
                default: undefined
            }
        },
        data() {
            return {
                showScrollTop:true,
                timer:null,
                scrollToTopStyle:{
                    position:'fixed',
                    top:'',
                    bottom:'',
                    left:'',
                    right:''
                }
            }
        },
        computed:{
            showTop:function(){
                let value = this.scrollTop > 200?true:false;
                return value;
            }
        },
        mounted() {
            this.initBtnPosition()
            this.initBindScroll()
        },
        methods:{
            isNumber (str) {
                if (!new RegExp('^[0-9]+([.]{1}[0-9]+)?$').test(str)) {
                    return false
                }
                return true
            },
            watchPosition () {
                if (![this.top, this.bottom, this.left, this.right].find(i => i !== undefined)) {
                    this.scrollToTopStyle.bottom = this.scrollToTopStyle.right = '14px'
                }
            },
            watchTop () {
                if (this.top !== undefined) {
                    this.scrollToTopStyle.top = this.isNumber(this.top) ? parseFloat(this.top) + 'px' : this.top
                }
            },
            watchBottom () {
                if (this.bottom !== undefined) {
                    this.scrollToTopStyle.bottom = this.isNumber(this.bottom) ? parseFloat(this.bottom) + 'px' : this.bottom
                }
            },
            watchLeft () {
                if (this.left !== undefined) {
                    this.scrollToTopStyle.left = this.isNumber(this.left) ? parseFloat(this.left) + 'px' : this.left
                }
            },
            watchRight () {
                if (this.right !== undefined) {
                    this.scrollToTopStyle.right = this.isNumber(this.right) ? parseFloat(this.right) + 'px' : this.right
                }
            },
            initBtnPosition () {
                this.watchTop()
                this.watchBottom()
                this.watchLeft()
                this.watchRight()
                this.watchPosition()
            },
            initBindScroll () {
                document.onscroll = ((oldScrollTopLength) => {
                    const clientHeight = document.documentElement.clientHeight
                    return () => {
                        const scrollTopLength = document.documentElement.scrollTop || document.body.scrollTop
                        if (!this.timer) {
                            this.showScrollToTop = scrollTopLength > clientHeight
                        }
                        if (scrollTopLength > oldScrollTopLength && this.timer) {
                            this.timer = clearInterval(this.timer)
                        }
                        oldScrollTopLength = scrollTopLength
                    }
                })(0)
            },

            scrollToTop () {
                this.timer = setInterval(() => {
                    const scrollTopLength = document.documentElement.scrollTop || document.body.scrollTop
                    if (scrollTopLength <= 0) {
                        this.timer = clearInterval(this.timer)
                        this.showScrollToTop = false
                    }
                    const spend = scrollTopLength / 5
                    document.documentElement.scrollTop = document.body.scrollTop = scrollTopLength - spend
                }, 30)
            }
        },
        components:{}
    }
</script>
<style scoped>
    #back-top {
        position: fixed;
    }
</style>