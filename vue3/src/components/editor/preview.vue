<template>
  <div class="marked">
    <div ref="preview" class="write article-concise" v-html="html"></div>
    <div class="pre-right">
      <AboutAuthor v-if="authorInfo.id" :authorInfo="authorInfo"></AboutAuthor>
      <AboutArticle v-if="authorInfo.id" :list="aboutArticleList"></AboutArticle>
    <div
      class="toc-wrap concise"
      v-if="tocNav && tocList"
      :class="{ 'toc-active': !tocIsShow }"
    >
      <a href="javascript:;" class="toc-close" @click="tocIsShow = !tocIsShow">
        <p class="menu" :class="{ 'menu-active': tocIsShow }"></p>
      </a>
      <transition :name="autoShow ? 'slide-fade' : ''">
        <div class="toc" v-show="tocIsShow">
          <div class="toc-top a-tag">
            <span class="toc-title">目录</span>
          </div>
          <ul class="toc-list">
            <li
              v-for="item of tocList"
              :key="item.id"
              @click="toTarget(item.id)"
              :class="{ acitve: item.id == currentTag }"
              :style="{ 'padding-left': item.tag - maxTitle + 'em' }"
              v-html="item.content"
            ></li>
          </ul>
        </div>
      </transition>
    </div>
    </div>
  </div>
</template>
<script>
import marked from "marked";
import { nextTick, onMounted, onUnmounted, ref, watch } from "vue";
import { throttle } from "throttle-debounce";
import DOMPurify from "dompurify";
import utils from "../utils";
import hljs from "./highlight.min.js";
import AboutAuthor from "@//components/about/author";
import AboutArticle from "@//components/about/article";
let rendererMD = new marked.Renderer();
export default {
  name: "MyMarked",
  props: {
    content: {
      // 初始化内容
      type: String,
      default: "",
    },
    // 配置
    markedOptions: {
      type: Object,
      default: () => ({}),
    },
    copyCode: {
      // 复制代码
      type: Boolean,
      default: false,
    },
    // 消毒
    dompurify: {
      type: Boolean,
      default: true,
    },
    copyBtnText: {
      // 复制代码按钮文字
      type: String,
      default: "复制代码",
    },
    imgView: {
      type: Boolean,
      default: false,
    },
    tocNav: {
      type: Boolean,
      default: false,
    },
    scroolOffset: {
      tyep: Number,
      default: 60,
    },
    getContainer: {
      tyep: Function,
      default: () => window,
    },
    width: {
      type: Number,
      default: 767,
    },
    authorInfo: {
      type: Object,
      default: () => ({})
    },
    aboutArticleList: {
      type: Object,
      default: () => ({})
    },
  },
  components: {
    AboutAuthor,
    AboutArticle
  },
  setup(props, context) {
    const html = ref("");
    const imgList = ref([]);
    const tocList = ref([]);
    const preview = ref(null);
    const tocIsShow = ref(true);
    const currentTag = ref(0);
    const idFre = "toc_";
    const maxTitle = ref(6);
    const autoShow = ref(true);
    const addImageClickListener = () => {
      nextTick(() => {
        console.log("注册图片点击事件");
        const imgs = preview.value.querySelectorAll("img");
        for (let i = 0, len = imgs.length; i < len; i++) {
          imgs[i].onclick = () => {
            const src = imgs[i].getAttribute("src");
            context.emit("click-img", src, imgList.value);
            console.log(src);
          };
        }
      });
    };
    const getImgList = () => {
      const b = /<img([\s]+|[\s]+[^<>]+[\s]+)src=("([^<>"']*)"|'([^<>"']*)')[^<>]*>/gi;
      const s = html.value.match(b);
      imgList.value = [];
      if (!s) return;
      for (let i = 0; i < s.length; i++) {
        const ss = s[i].match(b);
        imgList.value.push(RegExp.$3 + RegExp.$4);
        console.log(ss);
        console.log(RegExp.$3 + RegExp.$4);
      }
    };
    const getToc = () => {
      const r = /<h([1-6])([\S\s]*?)>([\S\s]*?)<\/h\1>/gi;
      tocList.value = [];
      let s = 0;
      if (!html.value) return;
      html.value = html.value.replace(r, (match, tag, attr, content) => {
        s++;
        tag = parseInt(tag);
        if (tag < maxTitle.value) {
          maxTitle.value = tag;
        }
        tocList.value.push({
          tag: tag,
          content: content,
          id: idFre + s,
        });
        // console.log(match)
        return `<h${tag}${attr} id="${idFre}${s}">${content}</h${tag}>`;
      });
    };
    const analysis = () => {
      rendererMD.link = function (href, title, text) {
        return (
          '<a href="' +
          href +
          '" title="' +
          text +
          '" target="_blank">' +
          text +
          "</a>"
        );
      };
      let DEMO_UID = 0;
      let SHOW_UID = 0;
      rendererMD.code = function (code, language) {
        // 提取language标识为 demo 的代码块重写
        if (language === "demo") {
          DEMO_UID += 2;
          // 页面中可能会有很多的示例，这里增加ID标识
          const id = "demo-mobai-template-" + DEMO_UID;
          // 将代码内容保存在template标签中
          const template = `<template type="text/demo" id="${id}">${code}</template>`;
          // 将template和自定义标签通过ID关联
          const sandbox = `<demo-mobai template="${id}"></demo-mobai>`;
          // 返回新的HTML
          return template + sandbox;
        }
        if (language === "show") {
          // 页面中可能会有很多的示例，这里增加ID标识
          const id = "demo-mobai-template-" + ++SHOW_UID;
          // 将代码内容保存在template标签中
          const template = `<template type="text/demo" id="${id}">${code}</template>`;
          // 将template和自定义标签通过ID关联
          const sandbox = `<demo-mobai template="${id}"></demo-mobai>`;
          // 返回新的HTML
          return template + sandbox;
        }

        // 其他标识的代码块依然使用代码高亮显示
        return `<div class="code-block"><span class="code-language">${language}</span><span class="copy-code el-icon-files">${
          props.copyBtnText
        }</span><pre rel="${language}"><code class="hljs ${language}">${
          hljs.highlightAuto(code).value
        }</code></pre></div>`;
      };
      html.value = marked(props.content, {
        sanitize: false,
        renderer: rendererMD,
        highlight: function (code) {
          return hljs.highlightAuto(code).value;
        },
        ...props.markedOptions,
      });
      // console.log(html.value)
      if (props.tocNav) {
        getToc();
      }
      if (props.dompurify) {
        html.value = DOMPurify.sanitize(html.value);
      }
      if (props.imgView) {
        getImgList();
        addImageClickListener();
      }
    };
    const onScroll = () => {
      const scrollTop = utils.getScroll(props.getContainer, true);
      for (let i = 0, len = tocList.value.length; i < len; i++) {
        const element = tocList.value[i];
        const tag = document.getElementById(element.id);
        if (!tag) break;
        const y =
          utils.getElementPosition(tag).y - scrollTop - props.scroolOffset;
        // 找到最后一个小于的元素
        if (y > 10) {
          break;
        } else {
          currentTag.value = element.id;
        }
      }
    };
    const throOnscroll = throttle(100, onScroll);
    const toTarget = (id) => {
      console.log("toTarget");
      const tag = document.getElementById(id);
      if (!tag) return;
      console.log(utils.getElementPosition(tag));
      const y = utils.getElementPosition(tag).y - props.scroolOffset;
      console.log(tag.offsetTop, props.scroolOffset);
      utils.scrollTo(y, props.getContainer);
    };
    const reSize = () => {
      const width = document.body.clientWidth;
      console.log(width);
      if (width < props.width) {
        tocIsShow.value = false;
        autoShow.value = false;
      } else {
        tocIsShow.value = true;
        autoShow.value = true;
      }
    };
    const throReSize = throttle(200, reSize);
    onMounted(() => {
      analysis();
      reSize();
      props.getContainer.addEventListener("scroll", throOnscroll);
      window.addEventListener("resize", throReSize);
    });
    onUnmounted(() => {
      props.getContainer.removeEventListener("scroll", throOnscroll);
      window.removeEventListener("resize", throReSize);
    });
    watch(
      () => props.content,
      () => {
        console.log("内容更新");
        analysis();
      }
    );
    return {
      html,
      preview,
      tocList,
      imgList,
      maxTitle,
      tocIsShow,
      toTarget,
      currentTag,
      autoShow,
    };
  },
};
</script>
<style lang="stylus" scoped>
@import '~@/assets/style/marked.css';

.marked {
  display: flex;
  flex-flow: row nowrap;
  position: relative;
  align-items: flex-start;

  .write {
    flex: 1 1 auto;
    width: 1%;
    overflow: hidden;
  }
  .pre-right{
    width: 25%;
    margin: 0 0 0 20px;

  .toc-wrap {
    position: sticky;
    top: 100px;
    flex-shrink: 0;
    transition: all 0.5s;

    .toc-close {
      position: absolute;
      top: 0;
      right: 0;
      font-size: 14px;
      color: #989898;
      cursor: pointer;
      z-index: 10;

      .menu {
        width: 16px;
        height: 16px;
        position: relative;
        margin: 14px auto;

        &:before, &:after {
          content: '';
          display: block;
          width: 16px;
          height: 2px;
          background: rgba(223, 223, 223, 0.8);
          border-radius: 1px;
          position: absolute;
          left: 0;
          transition: all 0.15s ease-in-out;
        }

        &:before {
          top: 2px;
          box-shadow: 0 5px rgba(223, 223, 223, 0.8);
        }

        &:after {
          bottom: 2px;
        }
      }

      .menu-active {
        &:before, &:after {
          background: #f974a1;
        }

        &:before {
          top: 7px;
          box-shadow: none;
          transform: rotate(225deg);
        }

        &:after {
          bottom: 7px;
          transform: rotate(135deg);
        }
      }
    }
  }
}

  .toc-active {
    position: fixed;
    right: 100px;
    bottom: 5px;
    top: auto !important;
    left: auto !important;
    transform: none !important;
    width: 40px;
    height: 40px;
    background: #fff;
    border-radius: 20px;
    box-shadow: 0px 0px 10px 3px rgba(46, 58, 89, 0.1);
    transition: all 0.5s;

    &:hover {
      box-shadow: 0px 0px 20px 6px rgba(46, 58, 89, 0.1);
    }

    .toc-close {
      width: 100%;
      height: 100%;
      display: flex;
      align-items: center;
      top: 0 !important;
      left: 0;
    }
  }

  .toc {
    width: 220px;
    border-left: 1px solid #efefee;
    padding-left: 10px;

    .toc-top {
      display: flex;
      justify-content: space-between;
      align-items: center;
      padding: 10px 0;

      .toc-title {
        font-size: 18px;
        color: #3da8f5;

        &:before {
          content: '#';
          color: #3da8f5;
          padding-right: 3px;
        }
      }
    }

    .toc-list {
      overflow-y: auto;

      li {
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
        margin-bottom: 10px;
        line-height: 1em;
        text-align: left;
        font-size: 14px;
        color: #8599ad;
        transition: 0.3s;
        cursor: pointer;

        a {
          color: inherit;
        }

        &:hover {
          color: #3da8f5;
          text-decoration: underline;
        }

        &:before {
          content: '- ';
        }
      }

      .acitve {
        color: #3da8f5;
        font-size: 15px;
      }
    }
  }

  .toc-tag {
    width: 40px;
    height: 40px;
    position: fixed;
    right: 20px;
    bottom: 85px;
    z-index: 999;
    background: #585d5d;
    display: flex;
    align-items: center;
    justify-content: center;
    flex-flow: column;
    transition: all 0.3s;
    cursor: pointer;

    &:hover {
      background-image: linear-gradient(to right, #8EC5FC, #9FACE6);

      i:nth-child(1) {
        transform: translateX(2px);
      }

      i:nth-child(3) {
        transform: translateX(-2px);
      }
    }

    i {
      display: block;
      width: 24px;
      height: 2px;
      background-color: hsla(0, 0%, 100%, 0.75);
      margin: 3px 0;
      transition: all 0.2s ease-in-out;
    }
  }
}

@media (max-width: 767px) {
  .marked {
    .toc-wrap {
      position: fixed;
      top: 50%;
      left: 50%;
      transform: translate(-50%, -50%);
      transition: all 0.5s;
      background: #fff;
      padding: 20px 16px;
      border-radius: 20px;
      box-shadow: 0 0px 25px 5px rgba(46, 58, 89, 0.1);
      box-sizing: border-box;
      background-color: #fff;
      transition: all 0.3s;
      z-index: 100;

      .toc-close {
        top: 15px;
        right: 20px;
      }
    }

    .toc-active {
      right: 1rem;
    }
  }
}

.slide-fade-enter-active {
  transition: all 0.3s cubic-bezier(1, 0.5, 0.8, 1);
}

.slide-fade-leave-active {
  transition: all 0.3s cubic-bezier(1, 0.5, 0.8, 1);
}

.slide-fade-leave-to, .slide-fade-enter {
  position: fixed;
  right: 100px;
  bottom: 5px;
  transform: scale(0.1) translateX(100%);
  transform-origin: bottom center;
  opacity: 0;
}
</style>