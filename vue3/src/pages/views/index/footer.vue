<template>
    <footer class="footer" >
        <div class="container">
            <div class="foot-row service">
                <p class="title">wenknow•文知道-开发者的博客平台</p>
                    <ul class="link">
                        <li><router-link to="/">主页</router-link></li>
                        <li><router-link to="/message">留言</router-link></li>
                    </ul>
            </div>
            <div class="foot-row provide">
            </div>
            <div class="foot-row publicity">
                <span><a target="_blank" href="http://beian.miit.gov.cn/">沪ICP备2021012701号-1</a></span>
                <span><a class="a-flex" target="_blank" href="http://www.beian.gov.cn/portal/registerSystemInfo" >沪公网安备 34010002020462号</a></span>
            </div>
            <hr/>
            <div class="foot-row bottom">
                <p class="copyright">© 2021 - {{year}} <a href="/">wenknow</a> All Rights Reserved.</p>
                <div class="share">
                     <a href="http://wpa.qq.com/msgrd?v=3&uin=1365896679&site=qq&menu=yes" target="_blank" class="btn qq">
                            <QqOutlined/>
                        </a>
                        <a href="https://github.com/wenknow" target="_blank" class="btn github">
                            <GithubOutlined />
                        </a>
                </div>
            </div>
        </div>
    </footer>
    
</template>
<script>
import {
  QqOutlined,
  GithubOutlined,
} from "@ant-design/icons-vue";
export default {
    name:'IndexFooter',
      components: {
    QqOutlined,
    GithubOutlined,
  },
    props:{
      route:String
    },
    data(){
        return{
            year:2020,
        }
    },
   created(){
       let date=new Date();
       this.year=date.getFullYear()
   }
}
</script>
<style lang="stylus" scoped>
    .footer
        padding: 10px 30px;
        font-size 12px
        color #a7a7a7
        // background:rgba(248, 249, 250, 0.9)
        background:#fff
        font-weight:bolder
        letter-spacing: 2px;
        padding:10px
        line-height 26px
        a
            color #a7a7a7
            font-size 12px
            transition all .3s
            &:hover
                text-decoration underline 
                color #3a8ee6
        .container
            padding 0 10px
            box-sizing: border-box;
            .foot-row
                display: flex;
                justify-content space-between
                align-items: center;
                flex-wrap wrap 
            .service
                font-size 16px
                .link
                    li
                        
                        float right
                        margin 0 20px
                        a
                            font-size 14px
            .provide
                .a-flex
                    display: flex;
                    justify-content:stretch;
                    align-items: center;
                    .ypy
                        height 22px
                .search-wrap 
                    width: 300px;
                    height 22px
                    position: relative;
                    padding 2px 0
                    .search-ipt 
                        height 20px
                        position absolute 
                        right 0
                        top 0
                        width: 240px;
                        color: #3a8ee6;
                        background: none;
                        border: none;
                        border-bottom: 2px solid #f5f5f5;
                        font-family: Roboto, serif;
                        font-size: 1rem;
                        line-height: 1.75;
                        transition all .4s
                        &:focus
                            width 100%
                    .icon 
                        outline:none;
                        background-color: transparent;
                        border: none;
                        padding 7px
                        line-height 0
                        position: absolute;
                        right: 0;
                        color: #f5f5f5;
                        cursor: pointer;
           
            hr
                background-color #f0f0f0
            .bottom
                .share
                    display: flex;
                    align-items: center;
                    .btn
                        padding: 5px 10px;
                        border-radius: 10px;
                        color: rgba(0, 0, 0, 0.45);
                        font-size: 14px;
                        margin: 5px;
                        display inline-block
                        background: linear-gradient(180deg, #dcf4ff, #f4fcff);
                        &:hover
                            background: linear-gradient(180deg, #fff0ce, #fffcf6);
                    .gitee
                        display: inline-block;
                        color: inherit;
                        font-style: normal;
                        line-height: 0;
                        text-align: center;
                        text-transform: none;
                        vertical-align: -0.125em;
                        text-rendering: optimizeLegibility;
                        -webkit-font-smoothing: antialiased;
                        -moz-osx-font-smoothing: grayscale;
                        height: 34px;
                        box-sizing: border-box;
                        display: flex;
                        align-items: center;

                    

            
     @media(max-width: 700px)
        .footer
            padding 10px
            font-size 12px
            letter-spacing: 2px;
            line-height 24px
            .container
                padding 0!important
                .foot-row
                    justify-content center!important
                .service
                    .title
                        font-size 14px
                    .link
                        li
                            margin 0 5px!important
                .provide
                    .search-wrap
                        padding 1px 0!important
                        .search-ipt 
                            right 50%!important
                            transform translateX(50%)
                .publicity
                    span 
                        margin 0 5px
                .bottom
                    .copyright
                        letter-spacing: 1px;
                
</style>