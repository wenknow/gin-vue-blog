<template>
  <div class="blog-wrap">
    <div class="blog container">
      <div
        class="blog-left"
        v-infinite-scroll="getBlogTo"
        :infinite-scroll-distance="20"
        :infinite-scroll-disabled="loading"
        infinite-scroll-window="false"
      >
        <div class="article article-concise" v-for="item of list" :key="item.id">
          <router-link :to="'/blog/' + item.id" class="article-link">
            <div class="header">
              <span class="title">
                {{ item.title }}
              </span>
              <span class="date"
                >{{ formatTime(item.created_at) }}<FieldTimeOutlined
              /></span>
            </div>
            <div class="body">
              <div class="content" v-html="item.desc"></div>
            </div>
            <div class="footer">
              <div class="info-data">
                <router-link :to="'/blog/' + item.id"
                ><FireOutlined />{{ item.browse_count }}</router-link>
                <router-link :to="'/blog/' + item.id + '#comment'"
                ><LikeOutlined /> {{ item.good_count }}</router-link>
                <router-link :to="'/blog/' + item.id"
                ><MessageOutlined />{{ item.comment_count }}</router-link>
              </div>
              <div class="info-arti">
                <router-link
                        :to="'/?tag=' + value"
                        v-for="(value, index) of (item.tags.split(','))"
                        :key="index"
                >
                  <TagsOutlined v-if="index == 0" />
                  <TagOutlined v-else />
                  {{ value }}
                </router-link>
              </div>
            </div>
          </router-link>
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
              <img
                      style="width: 60%; margin: 0 20%"
                      src="https://iconfont.alicdn.com/s/e8a7a96a-7f80-4f07-bc88-c162da451fe1_origin.svg"
              />
            </template>
          </Image>
        </div>
        <template v-if="loading">
          <Skeleton
            v-for="item of 6"
            :key="'Skeleton' + item"
            class="concise"
            active
          ></Skeleton>
        </template>
        <p class="last" v-if="page_num * page_size >= all_num">
          —— 我可是有底线的！( •̥́ ˍ •̀ू ) ——
        </p>
      </div>
      <div class="blog-right">
        <div class="label concise">
          <router-link
                  :to="getUrl(item.name, '')"
                  class="label-item"
                  :class="{ 'label-item-active': $route.query.tag == item.name }"
                  v-for="item of labelList"
                  :key="item.tag_id"
          >
            {{ item.name }}
          </router-link>
          <Skeleton v-if="labelLoading" active></Skeleton>
        </div>
      </div>
    </div>
  </div>
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
import {
  getArticleList,
  getMainTagList,
  getBlogInfo,
} from "./blog.js";
import { watch } from "vue";
import { useRoute } from "vue-router";
import Skeleton from "@/components/skeleton/skeleton";
import utils from "@/components/utils";
export default {
  name: "Blog",
  components: {
    FieldTimeOutlined,
    TagsOutlined,
    TagOutlined,
    FireOutlined,
    MessageOutlined,
    LikeOutlined,
    Image,
    Skeleton,
  },
  setup() {
    const {
      getBlog,
      list,
      loading,
      page_num,
      page_size,
      all_num,
    } = getArticleList();
    const { labelList, labelLoading, getTag } = getMainTagList();
    const { count } = getBlogInfo();
    const route = useRoute();
    const name = route.name;
    const authorInfo = {
      name:"11212"
    };
    getBlog();
    getTag();
    const getBlogTo = () => {
      if (name == route.name) {
        getBlog();
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
        const name = route.name;
        if (name == "Blog") {
          page_num.value = 0;
          page_size.value = 10;
          all_num.value = 1;
          getBlog();
          getTag();
          utils.scrollTo(0, {});
        }
      }
    );
    return {
      list,
      getBlogTo,
      formatTime,
      loading,
      labelLoading,
      labelList,
      count,
      page_num,
      page_size,
      all_num,
      authorInfo,
      getUrl
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

    .article {
      display: flex;
      box-sizing: border-box;
      margin-bottom: 10px;
      height: 175px;

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
          a {
            margin: 0 5px;
            cursor: pointer;
            font-size: 14px;
            color: #bababa;

            span {
              margin-right: 4px;
            }
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