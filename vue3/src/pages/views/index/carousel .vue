<template>
    <div class="carousel"  >
        <div class="cover" :class="{'cover-title':coverShow}" :style="{'background-image':list[current]?`url(${list[current].img_url})`:false,'justify-content':list[current]?list[current].link:false}">
             <slot>
                <transition name="fade" v-for="(item,index) of list" :key="item.id" v-show="index==current">
                <div class="carousel-item" :style="{'align-items':item.link}">
                    <div class="title" v-html="item.title"></div>
                    <div class="content" v-html="item.info"></div>
                </div>
            </transition>
               
             </slot>
        </div>
       
        <LeftOutlined class="prev" @click="swit(-1)" />
        <RightOutlined class="next" @click="swit(1)"/>
        <div class="waves-area">
          <svg
            class="waves-svg"
            xmlns="http://www.w3.org/2000/svg"
            xmlns:xlink="http://www.w3.org/1999/xlink"
            viewBox="0 24 150 28"
            preserveAspectRatio="none"
            shape-rendering="auto"
          >
            <defs>
              <path
                id="gentle-wave"
                d="M -160 44 c 30 0 58 -18 88 -18 s 58 18 88 18 s 58 -18 88 -18 s 58 18 88 18 v 44 h -352 Z"
              ></path>
            </defs>
            <g class="parallax">
              <use xlink:href="#gentle-wave" x="48" y="0"></use>
              <use xlink:href="#gentle-wave" x="48" y="3"></use>
              <use xlink:href="#gentle-wave" x="48" y="5"></use>
              <use xlink:href="#gentle-wave" x="48" y="7"></use>
             
            </g>
          </svg>
        </div>
    </div>
</template>
<script>
import { computed, ref } from 'vue'
import {LeftOutlined,RightOutlined} from '@ant-design/icons-vue'
import utils from '@/components/utils'
export default {
    name:'Carousel',
    components:{
        LeftOutlined,
        RightOutlined
    },
    props:{
        list:{
            type:Array,
        },
        position:{
            tyep:String,
            default:'center'
        },
        interval:{
            type:Number,
            default:10000
        }
    },
    setup(props){
        const current=ref(0)
        const listLen=computed(()=>{
            return props.list.length
        })
        let id
        const autoPlay=()=>{
            id=setInterval(()=>{
                // console.log(current.value++)
                const index=current.value+1
                current.value=utils.limit(index,0,listLen.value-1)
                // console.log(current.value)
            },props.interval)
        }
       const coverShow=computed(()=>{
           let show=true 
           const item=props.list[current.value]
           if(!item){
               return false
           }
           if(item.title||item.info){
               show=false
           }
           return show
       })
        const swit=(val)=>{
             clearInterval(id)
             const index=current.value+val
             current.value=utils.limit(index,0,listLen.value-1)
             autoPlay()
        }
        autoPlay()
        return {
            current,
            listLen,
            swit,
            coverShow
        }
    }
}
</script>
<style lang="stylus" scoped>
    .carousel
        position relative
       
        &:hover
                .next,.prev
                    opacity: 1;
        
        .cover
            padding: 40px;
            box-sizing: border-box;
            display flex
            align-items center
            // background-size: cover;
            // background-repeat: no-repeat;
            // background-position: bottom;
            overflow: hidden;
            position: absolute;
            top: 0;
            left: 0;
            width: 100%;
            height: 100%;
            background-position: 50%;
            background-size: cover;
            background-image: url('https://p6-tt-ipv6.byteimg.com/img/pgc-image/cca4898a577c49efb8cd22805444a629~tplv-obj.image'); // p.pstatp.com/origin/fede00031b9922f5244a);
            z-index: 0;            
            transition: all 2.5s;
            &:after 
                content: '';
                position: absolute;
                top: 0;
                left: 0;
                width: 100%;
                height: 100%;
                z-index -1
                background-color:rgba(0, 0, 0, 0.3);
            
            .carousel-item
                display: flex;
                flex-direction: column;
                padding 20px 80px
                .title
                    font-size: 24px;
                    color: #ffffff;
                .content
                    font-size: 14px;
                    color: #ffffff;
                    line-height: 24px;
                    margin-top: 10px;
                    font-family: PingFang SC;
        .cover-title
            &:after
                background-color transparent
             
        .btn
            display: inline-block;
            text-decoration: none;
            padding: 0 32px;
            height: 48px;
            line-height: 48px;
            text-align: center;
            border-radius: 26.5px;
            font-size: 20px;
            transition: all .3s;
            box-shadow: 0 2px 8px rgba(0,0,0,.2);
        .prev,.next
            position: absolute
            top:50%
            transform translateY(-50%)
            font-size 20px
            padding 5px
            opacity: 0;
            transition opacity .5s;
            cursor: pointer;
            color #fff
            outline: none;
        .next
            right 10px
        .prev
            left 10px

.fade-enter-active{
    transition: opacity 1.5s ease,transform 2s ease;

}
.fade-leave-active {
    transition: opacity 0s ease;
}

.fade-enter-from,
.fade-leave-to {
  opacity: 0;
  transform scale(.5)
}
@media (max-width:767px) 
    .carousel
        &-item
            padding 20px 40px
            .title 
                font-size 16px!important
            .content 
                font-size 12px !important
    

</style>