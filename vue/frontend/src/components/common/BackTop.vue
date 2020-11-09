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
    // 组件代码来源：https://blog.rxliuli.com/p/a4fbddd2/

    export default {
        props: {
            // 定义上拉按钮容器的位置
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
            // window.addEventListener('scroll',this.getScrollTop);
            this.initBtnPosition()
            this.initBindScroll()
        },
        methods:{
            // getScrollTop(){
            //   this.scrollTop = window.pageYOffset || document.documentElement.scrollTop || document.body.scrollTop;
            // },
            // backtop(){
            //   document.body.scrollTop = 0;
            //   document.documentElement.scrollTop = 0;
            // },
            // runScroll(){
            //   this.scrollTo(document.documentElement,0,600);
            // },
            // scrollTo(element,to,duration){
            //   if (duration <= 0) return;
            //   var difference = to - element.scrollTop;
            //   var perTick = difference / duration * 10;
            //   setTimeout(function() {
            //     element.scrollTop = element.scrollTop + perTick;
            //     if (element.scrollTop == to) return;
            //     this.scrollTo(element,to,duration - 10);
            //   },10);
            // }
            isNumber (str) {
                // RegExp对象，用来规定在文本中检索的内容。test()方法检索参数字符串中的指定值，返回true || false
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
            /**
             * 初始化按钮的位置
             */
            initBtnPosition () {
                this.watchTop()
                this.watchBottom()
                this.watchLeft()
                this.watchRight()
                this.watchPosition()
            },
            initBindScroll () {
                // 监听窗口滚动
                document.onscroll = ((oldScrollTopLength) => {
                    const clientHeight = document.documentElement.clientHeight
                    return () => {
                        const scrollTopLength = document.documentElement.scrollTop || document.body.scrollTop
                        // 如果定时器不存在的话就正常计算滚动到顶部的图标是否存在
                        if (!this.timer) {
                            // 滚动到第二屏就显示
                            this.showScrollToTop = scrollTopLength > clientHeight
                        }
                        // 向下滚动时判断判断是否正在向上滚动，是的话就清除定时器，停在当前位置
                        if (scrollTopLength > oldScrollTopLength && this.timer) {
                            // 设置这个是因为有时候 clearInterval() 并不能清除这个属性，或许是 vuejs 组件中的属性特殊一点？
                            this.timer = clearInterval(this.timer)
                        }
                        oldScrollTopLength = scrollTopLength
                    }
                })(0)
            },
            /**
             * 回到顶部
             */
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