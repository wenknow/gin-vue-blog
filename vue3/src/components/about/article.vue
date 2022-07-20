<template>
  <div class="about-article article-concise">
    <div class="about-article__title">相关文章</div>
    <a v-for="item in list" :key="item.id" :href="'/blog/'+ item.id" target="_blank">
      <div class="article-item">
        <div class="article__title">{{ item.title }}</div>
        <div class="article__meta">
          <div class="meta-item">
           <span><LikeOutlined />{{ item.good_count }}</span>
          </div>
          <div class="meta-item">
            <span><MessageOutlined />{{ item.comment_count }}</span>
          </div>
        </div>
      </div>
    </a>
  </div>
</template>

<script>
  import {
    LikeOutlined,
    MessageOutlined
  } from "@ant-design/icons-vue";
export default {
  name: "AboutArticle",
  props: {
    list: {
      type: Array,
      default: () => []
    }
  },
  components: {
    MessageOutlined,
    LikeOutlined,
  },
}
</script>

<style lang="stylus" scoped>
.about-article{

  .about-article__title{
    padding: 15px;
    font-size: 16px;
    color: #2e3a59;
    font-weight: 700;
    border-bottom: 1px solid #eee;
  }
  
  .article-item{
    padding: 15px;
    font-size: 14px;
    line-height: 1.3;
    color: #333;
    cursor: pointer;

    &:hover{
      background: hsla(0,0%,84.7%,.1);
    }

    .article__meta{
      display: flex;
      align-items: center;
      margin-top: 10px;
      color: #b2bac2;
      font-size: 14px;

      .meta-item{
        display: flex;
        align-items: center;

        &:not(:last-child){
          margin-right: 20px;
        }
        span {
          margin-right: 4px;
        }
      }
    }
  }
}
</style>