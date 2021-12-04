<template>
    <div>
        <NavBar url="/" logo="https://p.pstatp.com/origin/fede00031b9922f5244a" :section="list">
            <template #footer>
                <DropDown>
                   <img src="https://p.pstatp.com/origin/fede00031b9922f5244a"  class="avatar"/>
                    <template #menu>
                        <ul class="list">
                            <li>登录</li>
                            <li>注册</li>
                        </ul>
                    </template>
                </DropDown>
            </template>
        </NavBar>
        
    </div>
</template>
<script>
import NavBar from './main'
import DropDown from './drop-down';
export default {
    name:'demo',
    components:{
        NavBar,
        DropDown
    },
    data(){
        return {
            list:[{
                title:'主页',
                link:'/',
                key:'blog',
                slot:'blog',
                triangle:false,
            },
            
            {
                title:'博客',
                link:'/',
                key:'text',
                childSlot:'textSlot',
                children:[
                    {
                        title:'vue',
                        link:'/',
                        key:'vue'
                    },
                    {
                        title:'java',
                        link:'/',
                        key:'java'
                    }
                ]
                },
                {
                     title:'收集箱',
                link:'/',
                key:'text',
                childSlot:'textSlot',
                children:[
                    {
                        title:'音乐圈',
                        link:'/',
                        key:'vue'
                    },
                    {
                        title:'项目圈',
                        link:'/',
                        key:'java'
                    }
                ]
                }
            ]
        }
        
    }
}
</script>
<style lang="stylus" scoped>
.avatar{
    padding 0 10px
    width: 48px;
    height: 48px;
    margin: 0;
    padding: 0;
    color: #3da8f5;
    text-align: center;
    cursor: pointer;
    background: #fff;
    line-height: 44px;
    border: 2px solid rgba(223,223,223,0.3);
    box-sizing: border-box;
    border-radius: 50%;
    color: #c4c4c4;
}
.list{
    list-style: none;
    margin 0
    padding 0
    text-align center;
    &>li{
        padding: 10px 0px;
        display: block;
        border-radius: 4px;
        transition: all 0.3s;
        text-align: center;
        cursor: pointer;
        white-space: nowrap;
        &:hover{
            background-color: #edf5fc
        }
    }
}
</style>