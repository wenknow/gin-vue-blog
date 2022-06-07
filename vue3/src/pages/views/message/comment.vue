<template>
  <div class="comment">
      <div class="header">
          <!-- <Avatar :size="48" :src="avator" /> -->
           <Image
            :src="avator"
            class="user">
            
              <template #placeholder>
                <img style="width: 100%;height: 100%;" src="https://iconfont.alicdn.com/s/e2e10a5a-ad58-4591-9987-73c68c23a136_origin.svg"/>
              </template><template #error>
                <img style="width: 100%;height: 100%;" src="https://iconfont.alicdn.com/t/8daee7d9-2099-4294-ae70-3b58deaaaca5.png"/>
              </template>
          </Image>
          <textarea ref="editor" v-model="content" class="editor" placeholder="支持markdown编辑"></textarea>
      </div>
      <div class="body">
          <div class="preview">
              <span title="表情" @click="clicShow(true)" class="emoji-btn"><svg viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="16172" width="22" height="22"><path d="M512 1024a512 512 0 1 1 512-512 512 512 0 0 1-512 512zM512 56.888889a455.111111 455.111111 0 1 0 455.111111 455.111111 455.111111 455.111111 0 0 0-455.111111-455.111111zM312.888889 512A85.333333 85.333333 0 1 1 398.222222 426.666667 85.333333 85.333333 0 0 1 312.888889 512z" p-id="16173"></path><path d="M512 768A142.222222 142.222222 0 0 1 369.777778 625.777778a28.444444 28.444444 0 0 1 56.888889 0 85.333333 85.333333 0 0 0 170.666666 0 28.444444 28.444444 0 0 1 56.888889 0A142.222222 142.222222 0 0 1 512 768z" p-id="16174"></path><path d="M782.222222 391.964444l-113.777778 59.733334a29.013333 29.013333 0 0 1-38.684444-10.808889 28.444444 28.444444 0 0 1 10.24-38.684445l113.777778-56.888888a28.444444 28.444444 0 0 1 38.684444 10.24 28.444444 28.444444 0 0 1-10.24 36.408888z" p-id="16175"></path><path d="M640.568889 451.697778l113.777778 56.888889a27.875556 27.875556 0 0 0 38.684444-10.24 27.875556 27.875556 0 0 0-10.24-38.684445l-113.777778-56.888889a28.444444 28.444444 0 0 0-38.684444 10.808889 28.444444 28.444444 0 0 0 10.24 38.115556z" p-id="16176"></path></svg></span>
              <span title="预览" @click="clicShow(false)" class="preview-btn"><svg viewBox="0 0 1024 1024" version="1.1" xmlns="http://www.w3.org/2000/svg" p-id="17688" width="22" height="22"><path d="M502.390154 935.384615a29.538462 29.538462 0 1 1 0 59.076923H141.430154C79.911385 994.461538 29.538462 946.254769 29.538462 886.153846V137.846154C29.538462 77.745231 79.950769 29.538462 141.390769 29.538462h741.218462c61.44 0 111.852308 48.206769 111.852307 108.307692v300.268308a29.538462 29.538462 0 1 1-59.076923 0V137.846154c0-26.899692-23.355077-49.230769-52.775384-49.230769H141.390769c-29.420308 0-52.775385 22.331077-52.775384 49.230769v748.307692c0 26.899692 23.355077 49.230769 52.775384 49.230769h360.999385z" p-id="17689"></path><path d="M196.923077 216.615385m29.538461 0l374.153847 0q29.538462 0 29.538461 29.538461l0 0q0 29.538462-29.538461 29.538462l-374.153847 0q-29.538462 0-29.538461-29.538462l0 0q0-29.538462 29.538461-29.538461Z" p-id="17690"></path><path d="M649.846154 846.769231a216.615385 216.615385 0 1 0 0-433.230769 216.615385 216.615385 0 0 0 0 433.230769z m0 59.076923a275.692308 275.692308 0 1 1 0-551.384616 275.692308 275.692308 0 0 1 0 551.384616z" p-id="17691"></path><path d="M807.398383 829.479768m20.886847-20.886846l0 0q20.886846-20.886846 41.773692 0l125.321079 125.321079q20.886846 20.886846 0 41.773693l0 0q-20.886846 20.886846-41.773693 0l-125.321078-125.321079q-20.886846-20.886846 0-41.773693Z" p-id="17692"></path></svg></span>
          </div>
          <div>
            <span class="str-leng">还可以输入<strong>{{length-content.length}}</strong>字</span>
              <el-button type="primary" @click="submit" size="mini" title="Cmd|Ctrl+Enter"  plain>提交</el-button>
          </div>
      </div>
      <div class="footer">
          <div class="preview-wrap" v-show="preview">
              <Preview :content="format"/>
          </div>
          <div class="emoji-wrap" v-show="emojiShow">
              <i v-for="(item,key) of emoji" :key="key" :title="key.slice(1,-1)" @click="emojiAdd(key)">
                  <img :src="item"/>
              </i>
          </div>
        </div>
  </div>
</template>
<script>
import Image from "@/components/image/image";
import Preview from "@/components/editor/preview";
import message from "@/components/message/index.js";
import { computed, ref } from 'vue';
import emoji from './emoji.json'
import { useStore } from 'vuex';
export default {
  name: "Main",
   components:{
      Image,
      Preview
  },
  props:{
      url:{
          type:String,
          default:'https://iconfont.alicdn.com/s/e2e10a5a-ad58-4591-9987-73c68c23a136_origin.svg'
      },
      length:{
          type:Number,
          default:600
      },
    //   isLogin:{
    //       type:Boolean,
    //       default:false
    //   }
  },
  setup(props,context){
      const content=ref("")
      const emojiShow=ref(false)
      const preview=ref(true)
      const editor=ref(null)
      const store=useStore()
      const isLogin=computed(()=>{
          const token=store.state.user.token
          if(token){
              return true
          }
          return false
      })
      const avator=computed(()=>{
          const user=store.state.user.info
          if(user){
              return user.head_url
          }
          return props.url
      })
    //   提交
      const submit=()=>{
          const len=content.value.length
          if(len==0||len>props.length){
              message.warning('内容不能为空！');
              return
          }
          if(!isLogin.value){
              message.warning('未登录,请登录后提交！');
              store.commit('showLogin')
              return
          }
          context.emit('submit',content.value)
          content.value=""
      }
    //   添加表情
      const emojiAdd=(em)=>{
        let cursurPosition=editor.value.selectionStart;
        if(!cursurPosition){
            cursurPosition=content.value.length
        }
        console.log(cursurPosition)
        content.value=content.value.slice(0,cursurPosition)+em+content.value.slice(cursurPosition)
      }
    //   预览和表情显示：true为表情显示
      const clicShow=(show=false)=>{
          if(show){
              emojiShow.value=!emojiShow.value
              // preview.value=false
          }else{
              preview.value=!preview.value
              emojiShow.value=false
          }
      }
    const emotion=(re)=>{
        if(emoji[re]){
            return `![${re.slice(1,-1)}](${emoji[re]})`
        }
        return re
    }
    //   预览显示表情
      const format=computed(()=>{
          const cont=content.value.replace(/#(.){0,8}?;/gi,emotion)
          return cont
      })
    //   console.log(emoji)
      return {
          content,
          submit,
          emoji,
          emojiAdd,
          preview,
          emojiShow,
          clicShow,
          format,
          editor,
          avator
      }

  }
 
};
</script>
<style lang="stylus" scoped>
.comment
    border: 1px solid #f0f0f0;
    border-radius: 4px;
    margin-bottom: 10px;
    overflow: hidden;
    position: relative;
    padding: 10px;
    width: 100%;
    box-sizing: border-box;
    .header
        display flex
        .user  
            border: 2px solid rgba(223,223,223,0.3);
            box-sizing border-box
            width 48px
            height 48px
            background: #fff;
            line-height: 44px;
            border: 2px solid rgba(223,223,223,0.3);
            box-sizing: border-box;
            border-radius: 50%;
            color: #c4c4c4;
        .editor
            flex: 1;
            min-height: 8.75em;
            font-size: .875em;
            background: transparent;
            resize: vertical;
            transition: all .25s ease;
            border: none;
            resize: none;
            outline: none;
            padding: 10px 5px;
            max-width: 100%;
            font-size: .775em;
            display: block;
            font-size: 16px;
            line-height: 22px;
            color: #336a99;
    .body
        display flex;
        justify-content space-between
        align-items center
        .preview
            span
                cursor: pointer;
                display: inline-block;
                overflow: hidden;
                fill: #555;
                vertical-align: middle;
            .emoji-btn
                margin-right 10px
        .str-leng
            text-align: right;
            color: #b5b5b5;
            padding: 0 8px;
            font-size 14px
            strong
                text-align: right;
                color: #b5b5b5;
                font-family: georgia;
                font-size: 18px;
                padding: 0 2px;
                font-weight: 700;
        .submit
            transition-duration: .4s;
            text-align: center;
            color: #555;
            border: 1px solid #ededed;
            border-radius: .3em;
            display: inline-block;
            background: transparent;
            margin-bottom: 0;
            font-weight: 400;
            vertical-align: middle;
            touch-action: manipulation;
            cursor: pointer;
            white-space: nowrap;
            padding: .5em 1.25em;
            font-size: .875em;
            line-height: 1.42857143;
            user-select: none;
            outline: none;
            &:hover
                color: #3090e4;
                border-color: #3090e4;
    .footer
        /*padding 10px 0*/
        .emoji-wrap
            max-height 132px
            overflow-y: auto;
            &::-webkit-scrollbar {
                width: 4px;
                border-radius: 2px;
            }
            &::-webkit-scrollbar-thumb {
                width: 8px;
                border-radius: 2px;
                background-color: #e9e9e9;
            }
            i
                font-style: normal;
                width: 36px;
                cursor: pointer;
                text-align: center;
                display: inline-block;
                vertical-align: middle;
                padding 4px
                transition all .2s
                &:hover
                    box-shadow: 0px 0px 10px 3px rgba(46,58,89,0.1);
                    border-radius: 6px;
                img
                    max-width: 25px;
                    vertical-align: middle;
                    margin: 0 1px;
                    display: inline-block;
                    user-select: none

    

</style>