<template>
  <div class="by-image" ref="ByImage">
    <slot v-if="loading" name="placeholder">
      <div class="by-image__placeholder">加载中……</div>
    </slot>
    <slot v-else-if="error" name="error">
      <div class="by-image__error">图片加载失败</div>
    </slot>
    <img
      class="by-image__inner"
      v-else
      :src="src"
      :style="{ 'object-fit': fit }"
      :alt="alt"
    />
  </div>
</template>
<script>
import { ref, onMounted, watch, onUnmounted } from "vue";
import { throttle } from "throttle-debounce";
import utils from "../utils";
export default {
  name: "ByImage",
  props: {
    src: {
      type: String,
      request: true,
    },
    fit: {
      type: String,
    },
    alt: {
      type: String,
    },
    referrerPolicy: {
      type: String,
    },
    lazy: {
      type: Boolean,
      default: false,
    },
    scrollContainer: {
      type: Function,
      default: () => window,
    },
    offset: {
      type: Number,
      default: 0,
    },
  },
  emits: ["load", "error"],
  setup(props, context) {
    let loading = ref(true);
    let error = ref(false);
    const ByImage = ref(null);
    const loadImage = () => {
      loading.value = true;
      error.value = false;
      const image = new Image();
      image.src = props.src;
      image.onload = (e) => {
        loading.value = false;
        context.emit("load", e);
      };
      image.onerror = (e) => {
        loading.value = false;
        error.value = true;
        context.emit("error", e);
      };
    };
    const lazyImage = () => {
      const container = props.scrollContainer();
      if (utils.isInContainer(ByImage.value, container, props.offset)) {
        loadImage();
        container.removeEventListener("scroll", throttleImg);
        return true;
      }
      return false;
    };
    const throttleImg = throttle(200, lazyImage);
    onMounted(() => {
      // console.log(ByImage.value )
      if (props.lazy) {
        if (!lazyImage()) {
          const container = props.scrollContainer();
          container.addEventListener("scroll", throttleImg);
        }
      } else {
        loadImage();
      }
    });
    onUnmounted(() => {
      const container = props.scrollContainer();
      container.removeEventListener("scroll", throttleImg);
    });
    watch(
      () => props.src,
      () => {
        console.log("更新链接");
        loadImage();
      }
    );
    return {
      loading,
      error,
      ByImage,
    };
  },
};
</script>
<style lang="stylus" scoped>
.by-image {
  position: relative;
  display: inline-block;
  overflow: hidden;

  .by-image__error, .by-image__inner, .by-image__placeholder {
    width: 100%;
    height: 100%;
  }

  .by-image__error, .by-image__placeholder {
    display: flex;
    justify-content: center;
    align-items: center;
    font-size: 14px;
    color: #c0c4cc;
    vertical-align: middle;
    background: #f5f7fa;
  }
}
</style>