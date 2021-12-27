<template>
  <div class="blog-wrap">
    <div class="blog container">
      <div class="blog-left" v-infinite-scroll="getBlogTo" :infinite-scroll-distance="20" :infinite-scroll-disabled="state.loading" infinite-scroll-window="false">
        <div class="about-user article-concise">
          <Image :src="authorInfo.head_url ? authorInfo.head_url: 'https://bit-images.bj.bcebos.com/bit-new/file/20210323/4inn.png'" class="about-user-img">
          </Image>
          <div class="info-box">
            <h1 class="username">{{authorInfo.name}}</h1>
            <div class="intro">
              <svg  width="21" height="18">
                <path fill="#72777B" fill-rule="evenodd" d="M4 4h13a1 1 0 0 1 1 1v9a1 1 0 0 1-1 1H4a1 1 0 0 1-1-1V5a1 1 0 0 1 1-1zm9 6a2 2 0 1 0 0-4 2 2 0 0 0 0 4zm3 3a3 3 0 0 0-6 0h6zM5 7v1h4V7H5zm0 2.5v1h4v-1H5zM5 12v1h4v-1H5z"></path>
              </svg>
              <span class="user-content">{{authorInfo.desc ? authorInfo.desc: '人的高贵在于不断的学习'}}</span>
            </div>
          </div>
        </div>
        <div class="article-list article-concise">
          <el-tabs class="article-list-tab" v-model="state.tab_value" @tab-click="handleTabClick">
            <el-tab-pane label="我的文章" name="a"></el-tab-pane>
            <el-tab-pane label="我的草稿" name="f"></el-tab-pane>
          </el-tabs>
          <div class="article" v-for="item of state.list" :key="item.id">
            <div class="article-link">
            <router-link :to="'/blog/' + item.id" >
              <div class="header">
                <span class="title">{{item.cg_name ? '【'+item.cg_name+'】' : ''}}
                  {{ item.title }}
                </span>
                <span class="date"
                >{{ formatTime(item.created_at) }}<FieldTimeOutlined
                /></span>
              </div>
              <div class="body">
                <div class="content" v-html="item.desc"></div>
              </div>
            </router-link>
            <div class="footer">
                <div class="info-data">
                  <span><FireOutlined />{{ item.browse_count }}</span>
                  <span><LikeOutlined /> {{ item.good_count }}</span>
                  <span><MessageOutlined />{{ item.comment_count }}</span>
                </div>
                <div class="info-arti">
                  <span v-for="(value, index) of (item.tags.split(','))" :key="index">
                    <TagsOutlined v-if="index == 0" />
                    <TagOutlined v-else />
                    {{ value }}
                  </span>
                </div>
                <el-button class="del-btn" v-if="state.isAuthor" type="primary" size="mini" @click="delArticleDialog(item.id)" plain>删除</el-button>
              </div>
            </div>
            <Image
                    class="image concise"
                    v-if="item.img_url"
                    fit="cover"
                    :src="item.img_url"
                    :lazy="true"
            >
              <template #placeholder>
                <svg xmlns="http://www.w3.org/2000/svg" style="margin:auto;" width="200" height="200" viewBox="0 0 100 100" preserveAspectRatio="xMidYMid" display="block"><rect x="47" y="24" rx="3" ry="6" width="6" height="12" fill="#93dbe9"><animate attributeName="opacity" values="1;0" keyTimes="0;1" dur="1s" begin="-0.9166666666666666s" repeatCount="indefinite"/></rect><rect x="47" y="24" rx="3" ry="6" width="6" height="12" fill="#93dbe9" transform="rotate(30 50 50)"><animate attributeName="opacity" values="1;0" keyTimes="0;1" dur="1s" begin="-0.8333333333333334s" repeatCount="indefinite"/></rect><rect x="47" y="24" rx="3" ry="6" width="6" height="12" fill="#93dbe9" transform="rotate(60 50 50)"><animate attributeName="opacity" values="1;0" keyTimes="0;1" dur="1s" begin="-0.75s" repeatCount="indefinite"/></rect><rect x="47" y="24" rx="3" ry="6" width="6" height="12" fill="#93dbe9" transform="rotate(90 50 50)"><animate attributeName="opacity" values="1;0" keyTimes="0;1" dur="1s" begin="-0.6666666666666666s" repeatCount="indefinite"/></rect><rect x="47" y="24" rx="3" ry="6" width="6" height="12" fill="#93dbe9" transform="rotate(120 50 50)"><animate attributeName="opacity" values="1;0" keyTimes="0;1" dur="1s" begin="-0.5833333333333334s" repeatCount="indefinite"/></rect><rect x="47" y="24" rx="3" ry="6" width="6" height="12" fill="#93dbe9" transform="rotate(150 50 50)"><animate attributeName="opacity" values="1;0" keyTimes="0;1" dur="1s" begin="-0.5s" repeatCount="indefinite"/></rect><rect x="47" y="24" rx="3" ry="6" width="6" height="12" fill="#93dbe9" transform="rotate(180 50 50)"><animate attributeName="opacity" values="1;0" keyTimes="0;1" dur="1s" begin="-0.4166666666666667s" repeatCount="indefinite"/></rect><rect x="47" y="24" rx="3" ry="6" width="6" height="12" fill="#93dbe9" transform="rotate(210 50 50)"><animate attributeName="opacity" values="1;0" keyTimes="0;1" dur="1s" begin="-0.3333333333333333s" repeatCount="indefinite"/></rect><rect x="47" y="24" rx="3" ry="6" width="6" height="12" fill="#93dbe9" transform="rotate(240 50 50)"><animate attributeName="opacity" values="1;0" keyTimes="0;1" dur="1s" begin="-0.25s" repeatCount="indefinite"/></rect><rect x="47" y="24" rx="3" ry="6" width="6" height="12" fill="#93dbe9" transform="rotate(270 50 50)"><animate attributeName="opacity" values="1;0" keyTimes="0;1" dur="1s" begin="-0.16666666666666666s" repeatCount="indefinite"/></rect><rect x="47" y="24" rx="3" ry="6" width="6" height="12" fill="#93dbe9" transform="rotate(300 50 50)"><animate attributeName="opacity" values="1;0" keyTimes="0;1" dur="1s" begin="-0.08333333333333333s" repeatCount="indefinite"/></rect><rect x="47" y="24" rx="3" ry="6" width="6" height="12" fill="#93dbe9" transform="rotate(330 50 50)"><animate attributeName="opacity" values="1;0" keyTimes="0;1" dur="1s" begin="0s" repeatCount="indefinite"/></rect></svg>
              </template>
              <template #error>
                <img style="width: 60%; margin: 0 20%" src="https://iconfont.alicdn.com/s/e8a7a96a-7f80-4f07-bc88-c162da451fe1_origin.svg"/>
              </template>
            </Image>
          </div>
        </div>
        <template v-if="state.loading">
          <Skeleton
                  v-for="item of 6"
                  :key="'Skeleton' + item"
                  class="concise"
                  active
          ></Skeleton>
        </template>
        <p class="last" v-if="state.page_num * state.page_size >= state.all_num">
          —— 我可是有底线的！( •̥́ ˍ •̀ू ) ——
        </p>
      </div>
      <div class="blog-right">
        <AboutAuthor :authorInfo="authorInfo"></AboutAuthor>
        <ArticleCg :uid="state.params.uid" :isAuthor="state.isAuthor"></ArticleCg>
      </div>
    </div>
  </div>
  <el-dialog v-model="state.dialogDelVisible" title="删除确认" width="400px">
    <el-button type="danger" @click="state.dialogDelVisible=false">取消</el-button>
    <el-button type="primary" @click="delMyArticle()">确认</el-button>
  </el-dialog>
</template>
<script>
  import {
    FieldTimeOutlined,
    TagOutlined,
    TagsOutlined,
    FireOutlined,
    MessageOutlined,
    LikeOutlined,
  } from "@ant-design/icons-vue";
  import Image from "@/components/image/image";
  import AboutAuthor from "@/components/about/author";
  import ArticleCg from "@/components/about/articleCg.vue";
  import { watch, ref, reactive, onMounted } from "vue";
  import { useRoute } from "vue-router";
  import { useStore } from "vuex";
  import Skeleton from "@/components/skeleton/skeleton";
  import utils from "@/components/utils";
  import http from "@/utils/httpindex.js";
  export default {
    name: "Center",
    components: {
      FieldTimeOutlined,
      TagsOutlined,
      TagOutlined,
      FireOutlined,
      MessageOutlined,
      LikeOutlined,
      Image,
      Skeleton,
      AboutAuthor,
      ArticleCg,
    },
    setup() {
      const route = useRoute();
      const store = useStore();
      const state = reactive({
        list : [],
        loading: true,
        isAuthor: false,
        dialogDelVisible: false,
        page_num: 0,
        page_size: 10,
        all_num: 1,
        route_name: '',
        tab_value: 'a',
        delArticleId: 0,
        params: {
          page_num: 0,
          tag: '',
          uid: 0,
          public_is:true,
          cg_id: 0,
        },
    });
      const authorInfo = ref({});
      const getArticleList = ()=>{
        if (route.query.tag) {
          state.query.tag = decodeURI(encodeURI(route.query.tag).replace(/%20/g,'+'));
        }
        if (state.page_num * state.page_size < state.all_num) {
          state.page_num += 1;
          state.params.page_num = state.page_num;
          state.loading = true;
          http.post(`/apis/blog/userList`, state.params)
            .then(function(res) {
              if (res.page_num === 1) {
                state.list = res.list
              } else {
                state.list = state.list.concat(res.list)
              }
              state.loading = false;
              state.page_num = res.page_num;
              state.page_size = res.page_size;
              state.all_num = res.all_num;
            })
        }
      };
      const delArticleDialog = (id) =>{
        state.dialogDelVisible = true;
        state.delArticleId = id;
      };
      const delMyArticle = () => {
        http.post("/apis/blog/del", {id:state.delArticleId}).then(() => {
          state.dialogDelVisible = false;
          state.page_num = 0;
          state.page_size = 10;
          state.all_num = 1;
          getArticleList();
        });
      };
      const getAuthorInfo = async () => {
        await http.post("/apis/user/author",{id:state.params.uid}).then((res) => {
          authorInfo.value = res;
        });
      };
      const handleTabClick = (tab) =>{
        state.list= [];
        state.page_num = 0;
        state.params.cg_id = 0;
        if (tab.props.name === 'a'){
          state.params.public_is = true;
        }else {
          state.params.public_is = false;
        }
        state.page_num = 0;
        state.page_size = 10;
        state.all_num = 1;
        getArticleList();
      };

      const getBlogTo = () => {
        if (state.route_name == route.name) {
          getArticleList();
        }
      };
      const formatTime = (value) => {
        return utils.timestampToTime(value, true);
      };
      const getUrl = (tag, search) => {
        if (!search) {
          search = route.query.search;
        }
        if (!search) {
          search = '';
        }
        return '/?tag='+ tag + '&search='+search;
      };
      watch(
        () => route.query,
        () => {
          state.params.uid = parseInt(route.params.id, 10);
          state.params.cg_id = parseInt(route.query.cg_id, 10);
          if (route.name === "Center") {
            state.page_num = 0;
            state.page_size = 10;
            state.all_num = 1;
            state.params.cg_id = parseInt(route.query.cg_id, 10);
            getArticleList();
            utils.scrollTo(0, {});
          }
        }
      );
      onMounted(() => {
        state.route_name = route.name;
        state.params.uid = parseInt(route.params.id, 10);
        state.params.cg_id = parseInt(route.query.cg_id, 10);
        if (store.state.user.user.id === state.params.uid){
          state.isAuthor = true;
        }
        getAuthorInfo();
        getArticleList();
      });
      return {
        state,
        getBlogTo,
        formatTime,
        authorInfo,
        getUrl,
        handleTabClick,
        delArticleDialog,
        delMyArticle,
      };
    },
  };
</script>
<style lang="stylus" scoped>
  .blog-wrap {
    background-color: #f6f8f9;
  }

  .blog {
    display: flex;
    flex-flow: row nowrap;
    padding: 0 10px;
    padding-top: 72px;
    box-sizing: border-box;

  .blog-right {
    width: 20%;
    min-width: 24px;
    flex-shrink: 0;
    margin-left: 20px;

  .class, .info, .label {
    margin-top: 10px;
  }

  .class {
  .serach-wrap {
    display: flex;
    align-items: center;
    padding: 8px 0 8px 10px;
    box-sizing: border-box;
    height: 30px;
    border: 1px solid #ccc;
    border-radius: 15px;

  .serach {
    flex: 1 1 auto;
    outline: none;
    border: 0;
    height: 20px;
    color: rgba(0, 0, 0, 0.45);
  }

  .serach-btn {
    color: #ccc;
    height: 30px;
    width: 30px;
    border: 1px solid #ccc;
    box-sizing: border-box;
    border-radius: 15px;
    display: flex;
    justify-content: center;
    align-items: center;
    transition: all 0.3s;
    cursor: pointer;

  &:hover {
     background: linear-gradient(180deg, #dcf4ff, #f4fcff);
   }
  }
  }

  .class-item {
    text-align: center;
    color: rgba(0, 0, 0, 0.45);
    display: block;
    padding: 6px 5px;
    border-bottom: 1px solid #fafafa;
    margin-top: 5px;
    font-size: 14px;
    border-radius: 10px;
    transition: all 0.3s;

  &:hover {
     background: linear-gradient(180deg, #fff0ce, #fffcf6);
   }
  }

  .class-item-active {
    background: linear-gradient(180deg, #dcf4ff, #f4fcff);
  }
  }

  .label {
  .label-item {
    display: inline-block;
    padding: 5px 10px;
    border-radius: 10px;
    color: rgba(0, 0, 0, 0.45);
    font-size: 14px;
    margin: 5px;
    background: linear-gradient(180deg, #dcf4ff, #f4fcff);

  &:hover {
     box-shadow: 0 0px 10px 5px rgba(46, 58, 89, 0.1);
   }
  }

  .label-item-active {
    background: linear-gradient(180deg, #fff0ce, #fffcf6);
  }
  }
  }

.blog-left {
    flex: 1 1 auto;
    width: 47%;
  .about-user{
    display: flex;
    padding: 1rem;
    background-color: #fff;
    .info-box{
      display: flex;
      flex-direction: column;
      justify-content: center;
        .username{
          margin: 0;
          padding: 0;
          font-size: 1.8rem;
          font-weight: 600;
          line-height: 1.2;
          color: #000;
        }
        .intro{
          margin-top: 10px;
        }
    }
  }
  .about-user-img{
    flex: 0 0 auto;
    margin-right: 2.4rem;
    width: 7.5rem;
    height: 7.5rem;
    background-color: #f9f9f9;
    border-radius: 50%;
    position: relative;
    object-fit: cover;
  }

  .article {
    display: flex;
    box-sizing: border-box;
    margin-bottom: 10px;
    height: 165px;
    border-bottom: 1px solid #ebebeb;
  .image {
    /*width: 150px;*/
    height: 150px;
    overflow: hidden;
    float: right;
    margin: 0 10px 10px 0;
    padding: 0;
  }
  .article-link{
    width: 85%;
  }

  .header {
    padding-bottom: 15px;
    position: relative;
    border-bottom: 1px solid #fafafa;
    width: 90%;
  .title {
    color: #336a99;
    font-size: 20px;
    font-weight: 500;
    line-height: 1.5em;
    overflow: hidden;
    display: -webkit-box;
    -webkit-line-clamp: 2;
    -webkit-box-orient: vertical;
    word-break: break-all;
  }

  .date {
    position: absolute;
    bottom: 0;
    left: 0;
    transform: translateY(50%);
    font-size: 14px;
    color: #bebebe;

  .anticon {
    margin-left: 10px;
  }
  }
  }

  .body {
    padding: 15px 0;
    overflow: hidden;
    border-bottom: 1px solid #fafafa;
    width: 90%;
    height: 39px;
  .content {
    font-size: 14px;
    line-height: 24px;
    color: #bebebe;
    cursor: pointer;
    /*text-align: justify;*/
  }
  }

  .footer {
    display: flex;
    /*justify-content: space-between;*/
    align-items: center;
    flex-wrap: wrap;
    width: 90%;
  .info-arti, .info-data {
  span {
    margin: 0 5px;
    cursor: pointer;
    font-size: 14px;
    color: #bababa;

  }
  }
  }
  }

  .last {
    text-align: center;
    color: #bababa;
    font-size: 14px;
    margin: 10px 0;
  }
  }
  }

  @media (max-width: 776px) {
    .blog {
      flex-direction: column;

    .left {
      margin: 0;
      width: 100%;
    }

    .right {
    .article {
    .body {
      flex-direction: column;

    .image {
      width: 100%;
      margin: 0;
    }
  }
  }
  }
  }
  }
</style>