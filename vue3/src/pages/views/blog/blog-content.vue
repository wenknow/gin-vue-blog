<template>
  <div class="content-wrap" :style="myLoading?'padding-top:72px':''" >
    <div class="skeleton" v-if="myLoading">
        <Skeleton class="concise container" v-for="item of 2" :key="'myloading'+item" active/>
    </div>
    <div v-else class="content" >
      <div class="header" >
        <div class="info">
          <h4 class="title">{{ article.title }}</h4>
          <p class="desc">
            <span><FieldTimeOutlined />{{ formatTime(article.created_at) }}</span>
            <span>字数：{{ article.word_count }}</span>
            <span>阅读量：{{ article.browse_count }}</span>
            <router-link :to="'/blog/edit/' + article.id"> <el-button class="edit-btn" type="primary" size="mini" plain>编辑</el-button>
            </router-link>
          </p>
        </div>
      </div>
      <article-panel :userAction="userAction" :articleDetail="article"></article-panel>
      <div class="body container">
          <Preview
            :content="article.content"
            :tocNav="true"
            :dompurify="false"
            @click-img="clickImg"
            :imgView="true"
            :scroolOffset="72"
            :markedOptions="options"
            :authorInfo="authorInfo"
            :aboutArticleList="aboutArticleList"
          />
        <div class="comment-view">
          <div class="comment-list article-concise">
            <div class="skeleton"  v-if="myLoading">
              <Skeleton class="concise container" active/>
            </div>
            <template  v-else  >
              <Comment class="concise" @submit="submit" />
            </template>
            <div class="message" v-infinite-scroll="getMess" :infinite-scroll-distance="20" :infinite-scroll-disabled="loading" infinite-scroll-window="false">
              <MessageCard
                      @reply="reply"
                      @delete="deleteMess"
                      :item="item"
                      :author="authorInfo"
                      v-for="(item, index) of commentList"
                      :key="item.comment.id"
                      :class="'message-list-item-bg-' + ((index + 1) % 4)"
              />
              <div class="skeleton" v-if="loading">
                <Skeleton v-for="item of 3" :key="'Skeleton'+item" :class="'message-list-item-bg-'+(item+1)%4"  class="concise" avatar active></Skeleton>
              </div>
              <p class="last" v-if="page_num*page_size >= all_num">—— 我可是有底线的！( •̥́ ˍ •̀ू ) ——</p>
            </div>
          </div>
          <div class="pre-right"></div>
        </div>
    </div>
    </div>
  </div>
</template>
<script>
import { computed, ref, watch } from "vue";
import { useRoute } from "vue-router";
import http from "@/utils/httpindex.js";
import Skeleton from "@/components/skeleton/skeleton";
import utils from "@/components/utils";
import {
  FieldTimeOutlined,
} from "@ant-design/icons-vue";
import Preview from "@/components/editor/preview";
import ArticlePanel from '@/components/about/panel.vue';
import Comment from "../message/comment";
import MessageCard from "../message/message-card";
import { getMessList } from "../message/message.js";

export default {
  name: "BlogContent",
  components: {
    FieldTimeOutlined,
    Preview,
    Comment,
    MessageCard,
    Skeleton,
    ArticlePanel,
  },
  setup() {
    const route = useRoute();
    const article = ref({});
    const authorInfo = ref({});
    const userAction = ref({});
    const aboutArticleList = ref({});
    const myLoading=ref(true)
    const name=route.name
    const imgShow=ref(false)
    const imgSrc=ref('')
    const imgList=ref([]);
    const article_id = computed(()=>{
      return parseInt(route.params.id, 10);
    });
    const url=computed(()=>{
      console.log(process.env)
      return process.env.VUE_APP_HOME_URL+'/blog/'+article_id.value;
    });
    const options = {
      headerIds: false,
      breaks: true,
    };
    const logo="https://p.pstatp.com/origin/137d10002a95149f1fef7"
    const desc="Baymax•记录美好学习时光，分享学习点点滴滴"
    const getArticle = () => {
      myLoading.value=true;
      http.post("/apis/blog/detail",{id:article_id.value})
        .then((res) => {
            myLoading.value=false;
            article.value = res.detail;
            userAction.value = res.user_action;
            document.title=res.detail.title+'-文知道开发者的博客平台'
            document.querySelector('meta[name="keywords"]').setAttribute('content', res.detail.tags)
            document.querySelector('meta[name="description"]').setAttribute('content', res.detail.desc)

            http.post("/apis/user/author",{id:article.value.uid}).then((res) => {authorInfo.value = res});
        });
    };
    const getAboutArticle = async () => {
        await http.post("/apis/blog/relatedList",{id:article_id.value}).then((res) => {
            aboutArticleList.value = res;
        });
    };
    const formatTime = (value) => {
      return utils.timestampToTime(value, true);
    };
    getArticle();
    getAboutArticle();
    const {
      commentList,
      getMess,
      loading,
      page_num,
      page_size,
      all_num,
      messageAdd,
      replyAdd,
      deleteMess,
    } = getMessList(article_id.value);
    getMess();
    const submit = (content) => {
      const data = {
        content: content,
        article_id: article_id.value,
      };
      messageAdd(data);
    };
    const reply = (data) => {
      data.article_id = article_id.value;
      replyAdd(data);
    };
    const clickImg=(src,list)=>{
        imgSrc.value=src
        imgList.value=list
        imgShow.value=true

    }
    watch(
      () => route.params.id,
      (id) => {
        const newName=route.name;
        if(newName==name){
            console.log(id);
            getArticle();
            getMess(id);
        }
      }
    );
    const qzone=computed(()=>{
      return (
        "http://sns.qzone.qq.com/cgi-bin/qzshare/cgi_qzshare_onekey?url=" +
        url.value +
        "&title=" +
        article.value.title +
        "&desc=" +
        (article.value.desc?article.value.desc:desc) +
        "&summary=" +
        article.value.title +
        "&site=" +
        article.value.source +
        "&pics=" +
        (article.value.image?article.value.image:logo) +
        ""
      );

    })
    const qq=computed(()=>{
         return (
        "http://connect.qq.com/widget/shareqq/index.html?url=" +
        url.value +
        "&title=" +
        article.value.title +
        "&source=" +
        article.value.source +
        "&desc=" +
        (article.value.desc?article.value.desc:desc) +
        "&pics=" +
        (article.value.image?article.value.image:logo) +
        '&summary="' +
        article.value.title +
        '"'
      );
    })
    const weibo=computed(()=>{
      return (
        "https://service.weibo.com/share/share.php?url=" +
        url.value +
        "&title=" +
        article.value.title +
        "&pic=" +
        (article.value.image?article.value.image:logo) +
        "&appkey=" +
        article.value.weibokey +
        ""
      );
      
    })
    return {
      formatTime,
      article,
      authorInfo,
      userAction,
      aboutArticleList,
      options,
      submit,
      reply,
      commentList,
      getMess,
      loading,
      myLoading,
      page_num,
      page_size,
      all_num,
      deleteMess,
      clickImg,
      imgShow,
        imgSrc,
        imgList,
        url,
        qzone,
        qq,
        weibo
    };
  },
};
</script>
<style lang="stylus" scoped>
  .edit-btn {
    background: linear-gradient(315deg, #e4eef3, #fffcf6);
    padding: 5px 20px;
    border-radius: 10px;
    color: rgba(0, 0, 0, 0.45);
    margin: 0 30px 0 0;
  }
.comment-view{
    display: flex;
    flex-flow: row nowrap;
    position: relative;
    align-items: flex-start;
    margin-top: 15px;

  .comment-list{
    flex: 1 1 auto;
    width: 1%;
    overflow: hidden;
  }
  .pre-right{
    width: 25%;
    margin: 10px 0 0 20px;
  }
}
.content-wrap {
  background-color: #f6f8f9;
  .content{
      padding-top 0
  }
  .header {
    position: relative;
    overflow: hidden;
    margin-top: 72px;
    border-radius: 20px;
    padding 0 10px
    border-top-left-radius: 0;
    border-top-right-radius: 0;
    .cover {
      position: absolute;
      top: 0;
      left: 0;
      width: 100%;
      height: 100%;
      background-position: 50%;
      background-size: cover;
      background-image: url('https://p6-tt-ipv6.byteimg.com/img/pgc-image/cca4898a577c49efb8cd22805444a629~tplv-obj.image'); // p.pstatp.com/origin/fede00031b9922f5244a);
      z-index: 0;
      transition all .3s

      &:after {
        content: '';
        position: absolute;
        top: 0;
        left: 0;
        width: 100%;
        height: 100%;
        background-image: linear-gradient(-180deg, rgba(45, 58, 111, 0.5), rgba(0, 0, 0, 0.5));
      }
    }
    

    .info {
      position: relative;
      width: 100%;
      height: 100%;
      z-index: 1;
      display: flex;
      justify-content: center;
      align-content: center;
      flex-direction: column;
      text-align: center;

      .title {
        font-size: 24px;
        font-weight: 500;
        line-height: 1.5em;
        overflow: hidden;
        display: -webkit-box;
        -webkit-line-clamp: 2;
        -webkit-box-orient: vertical;
        word-break: break-all;
        margin: 10px 0;
      }

      .desc {
        font-size: 14px;
        margin-bottom: 10px;
        display: flex;
        justify-content: center;
        align-items: center;
        flex-wrap: wrap;

        span {
          margin: 0 10px;
        }
      }
    }
  }
  .body{
      .share{
        display: flex;
        align-items: center;
        justify-content: center;
        margin-top 30px
        a {
          margin 10px
          width 45px
          height 45px
          border-radius 50%
          display: flex;
          align-items: center;
          justify-content: center;
          box-sizing border-box
        }
        .qzone{
          &:hover{
            border 2px solid rgba(255, 205, 0,.3);
            background rgba(255, 205, 0,.2);
          }
        }
        .qq{
          &:hover{
            border 2px solid rgba(27, 193, 250,.3);
            background rgba(27, 193, 250,.2);
          }
        }
        .weibo{
          &:hover{
            border 2px solid rgba(231, 31, 25,.3);
            background rgba(231, 31, 25,.2);
          }
        }
      }
    }
  .copyright {
    .blog-info {
      border-left: 2px solid #bababa;
      padding-left: 10px;

      p {
        color: #bababa;
        font-size: 14px;
        font-weight: 500;
        line-height: 1.5em;
        text-overflow: ellipsis;
        white-space: nowrap;
        overflow: hidden;
        a {
          color: #bababa;
          font-size: 14px;
        }
      }
    }

    .article-prev-next {
      display: flex;
      justify-content: space-between;
      margin: 10px 0;
      .prev, .next {
        text-overflow: ellipsis;
        overflow: hidden;
        white-space: nowrap;
        font-size: 14px;
        padding: 0 16px;
        position: relative;
        flex: 1;
        color #3a8ee6
        
        .anticon {
          position: absolute;
          top: 50%;
          transform: translateY(-50%);
        }
      }

      .prev {
        .anticon {
          left: 0;
        }
      }

      .next {
        text-align: right;

        .anticon {
          right: 0;
        }
      }
    }
  }

  .concise {
    padding: 16px 20px;
    border-radius: 20px;
    margin-top: 10px;
    box-sizing: border-box;
    background-color: #fff;
    transition: box-shadow .3s;
  }

  .message {
      .last{
        text-align: center;
        color: #bababa;
        font-size: 14px;
        margin: 10px 0;
    }
    >>>.message-list-item-bg-0 {
      background: linear-gradient(180deg, #d2e5f9, #f8fbfe);
    }

    >>>.message-list-item-bg-1 {
      background: linear-gradient(180deg, #dcf4ff, #f4fcff);
    }

    >>>.message-list-item-bg-2 {
      background: linear-gradient(180deg, #fff0ce, #fffcf6);
    }

    >>>.message-list-item-bg-3 {
      background: linear-gradient(180deg, #ffe4cf, #fff7f1);
    }
  }
}
@media (max-width:776px) {
  .content-wrap{
    .desc{
        padding: 0 10px;
        box-sizing: border-box;
    }
    .skeleton{
      padding 0 10px
      box-sizing border-box
    }
  }
  
}
@media (max-width: 576px) {
    .content-wrap .concise {
        padding: 20px 10px
    }
}
</style>