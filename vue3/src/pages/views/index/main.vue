<template>
    <div class="index">
        <Carousel class="crousel" :list="list">

        </Carousel>
        <Introduce v-for="item of info" :key="item.key" :reverse="item.reverse" :url="item.url" :title="item.title" :content="item.content" :info="item.info" :class="item.key">
            <template #info="{info}" v-if="item.key=='study'" >
                <div class="study-item">
                    <div>
                         <span  class="image">
                            <img :src="info.url"/>
                        </span>
                        <p class="title">
                            {{info.title}}
                        </p>
                    </div>
                    <DoubleRightOutlined class="icon-right"/>
                </div>
            </template>
        </Introduce>
    </div>
</template>
<script>
import Carousel from './carousel ';
import { useStore } from "vuex";
import { computed } from 'vue';
import {DoubleRightOutlined} from '@ant-design/icons-vue'
import Introduce from './introduce'
export default {
    name:'Index',
    components:{
        Carousel,
        Introduce,
        DoubleRightOutlined
    },
    setup(){
        const store = useStore();
        store.dispatch('getCarousels')
        const list=computed(()=>store.state.index.list)
        const info=[
            {
                url:'https://iconfont.alicdn.com/t/7dc3c4f4-0805-477d-9f7e-20f3ec164107.png',
                title:'关于博客',
                content:'<p>这是一次重零开始的尝试、是自主学习的开始、是发现兴趣源头。</p><p>重零开始至博客2.0的重构，学无止境，继续出发！采用最新VUE 3.0架构，脱离UI库，自主编写UI组件！大三上完成博客前端重构</p><p>博客布局及样式参考——Teambition</p><p>组件样式来自于——Element、Ant</p><p>音乐播放器部分样式参考——APlayer</p>',
                key:'about',
                info:[
                    {
                        title:'Vue',
                        url:'https://iconfont.alicdn.com/t/7dc3c4f4-0805-477d-9f7e-20f3ec164107.png'
                    },
                     {
                        title:'Laravel',
                        url:'https://iconfont.alicdn.com/t/7dc3c4f4-0805-477d-9f7e-20f3ec164107.png'
                    },
                    {
                        title:'ByUI',
                        url:'https://iconfont.alicdn.com/t/7dc3c4f4-0805-477d-9f7e-20f3ec164107.png'
                    }
                ]
            },
            {
                url:'https://iconfont.alicdn.com/t/56f85d77-57a9-435b-b3ae-ec344ff25691.png',
                title:'学习目标',
                content:'<p>只有不断的学习才能进步，而循序渐进才是我的学习思想。</p><p>有了明确的学习目标才能我有动力前进，在本项目重构之前，也荒废了一段时间，沉浸在自己的完成博客的喜悦中，殊不知这才是一个开始！</p>',
                key:'study',
                reverse:true,
                info:[
                    {
                        title:'熟悉Vue生态',
                        url:'https://iconfont.alicdn.com/t/7dc3c4f4-0805-477d-9f7e-20f3ec164107.png'
                    },
                     {
                        title:'编写Vue UI库',
                        url:'https://iconfont.alicdn.com/t/7dc3c4f4-0805-477d-9f7e-20f3ec164107.png'
                    },
                    {
                        title:'了解Vue原理',
                        url:'https://iconfont.alicdn.com/t/7dc3c4f4-0805-477d-9f7e-20f3ec164107.png'
                    },
                    {
                        title:'待续...',
                        url:'https://iconfont.alicdn.com/t/7dc3c4f4-0805-477d-9f7e-20f3ec164107.png'
                    }
                ]
            },
            {
                url:'https://iconfont.alicdn.com/t/f524cdcc-89ee-4236-8e81-7e010a250d53.png',
                title:'关于我',
                content:'<p>本着无聊的态度学了了Vue,也由此迷上了前端！</p><p>Baymax,00后程序小白！思维呆板！做事执着！学习能力一般，除非感兴趣！</p><p>掌握技能：Vue、Laravel、Java、Php、Mysql、Js、Ps  ……</p><p>可以在以下场所找到我！</p>',
                key:'about-me',
                info:[
                    {
                        title:'Gitee',
                        url:'https://iconfont.alicdn.com/t/7dc3c4f4-0805-477d-9f7e-20f3ec164107.png',
                        link:'https://gitee.com/baymaxsjj'
                    },
                     {
                        title:'GitHub',
                        url:'https://iconfont.alicdn.com/t/7dc3c4f4-0805-477d-9f7e-20f3ec164107.png',
                        link:'https://github.com/baymaxsjj'
                    },
                    {
                        title:'Csdn',
                        url:'https://iconfont.alicdn.com/t/7dc3c4f4-0805-477d-9f7e-20f3ec164107.png',
                        link:'https://blog.csdn.net/weixin_45294607'
                    },
                    {
                        title:'待续...',
                        url:'https://iconfont.alicdn.com/t/7dc3c4f4-0805-477d-9f7e-20f3ec164107.png',
                        link:'/'
                    }
                ]
            },
        ]
        return {
            list,
            info
        }
    }
}
</script>
<style lang="stylus" scoped>
.index
    .crousel
        height 100vh;
    .about,.about-me
        background-color #fff
    .study
        background: #f6f8f9;
        .study-item
            color: #336a99;
            display: flex;
            align-items: center;
            // flex-wrap wrap
            margin-right: 10px;
           
            .image
                img
                    width: 30px;
                    height: 30px;
                    margin-bottom: 8px;
            .title
                font-size: 14px;
            .icon-right
                margin 5px

@media (max-width:767px) 
    .index
        .crousel
            height: 450px;
            border-radius 10px
        .study
            .study-item
                height: 100px;
                width: 120px;
                justify-content center
@media (max-width:576px) 
    .index
        .crousel
            height: 300px;
            border-radius 10px
            
            

</style>