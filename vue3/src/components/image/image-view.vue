<template>
  <teleport to="body">
    <transition name="fade">
      <div
        class="by-image-view"
        ref="ByImageView"
        tabindex="-1"
        :style="{ 'z-index': zIndex }"
        v-show="value"
      >
        <div class="by-image-mask"></div>
        <span class="by-image-btn by-image-close" @click="value = false">
          <IconFont type="iconGroup-2" />
        </span>
        <span class="by-image-btn by-image-prev" @click="next(-1)">
          <IconFont type="iconGroup-" />
        </span>
        <span class="by-image-btn by-image-next" @click="next(1)">
          <IconFont type="iconGroup-1" />
        </span>
        <div class="by-image-btn by-image-view-actions">
          <div class="by-image-view-actions-inner">
            <IconFont type="iconGroup-4" @click="handleActions(-0.2)" />
            <IconFont type="iconGroup-3" @click="handleActions(0.2)" />
            <IconFont
              :type="isFull ? 'iconGroup-6' : 'iconGroup-5'"
              @click="isFull = !isFull"
            />
            <IconFont type="iconGroup-8" @click="rotationActions(-90)" />
            <IconFont type="iconGroup-7" @click="rotationActions(90)" />
          </div>
        </div>
        <div class="by-image-view-canvas">
          <!-- <template
            v-for="(item, index) of urlList"
            :key="'by-image-view' + index"
          >
            <img
              v-if="index == currentIndex"
              @touchstart="onMousedown($event, false)"
              @mousedown="onMousedown"
              ref="ByImage"
              :style="imageStyle"
              :src="item"
            />
          </template> -->
          <img
            v-if="urlList.length"
            v-show="value"
            @touchstart="onMousedown($event, false)"
            @mousedown="onMousedown"
            ref="ByImage"
            :style="imageStyle"
            :src="urlList[currentIndex]"
          />
        </div>
      </div>
    </transition>
  </teleport>
</template>
<script>
import IconFont from "../icon-font/icon-font";
import { ref, reactive, watch, computed, onUnmounted } from "vue";
import { throttle } from "throttle-debounce";
import utils from "../utils.js";
export default {
  name: "ByImageView",
  props: {
    url: {
      type: String,
      default: "",
    },
    urlList: {
      type: Array,
      default: () => [],
    },
    zIndex: {
      type: Number,
      default: 2000,
    },
    show: {
      type: Boolean,
      default: true,
    },
  },
  components: {
    IconFont,
  },
  setup(props, context) {
    const transform = reactive({
      scale: 1,
      rotate: 0,
      left: 0,
      top: 0,
      transition: false,
    });
    const currentIndex = ref(0);
    const isFull = ref(false);
    const ByImageView = ref(null);
    const isPc = ref(true);
    // 图片拖动行为
    let thrimgPosition;
    const onMousedown = (e, pc = true) => {
      console.log("onMousedown");
      isPc.value = pc;
      e.preventDefault();
      e = isPc.value ? e : e.changedTouches[0];
      const startX = e.clientX;
      const startY = e.clientY;
      const offsetX = transform.left;
      const offsetY = transform.top;
      const imgPosition = (ev) => {
        ev = isPc.value ? ev : ev.changedTouches[0];
        console.log(ev.clientX, startX);
        transform.transition = true;
        transform.left = offsetX + ev.clientX - startX;
        transform.top = offsetY + ev.clientY - startY;
      };
      thrimgPosition = throttle(200, imgPosition);
      removeDocumentEvents();
      addDocumentEvents();
    };
    const onEnd = () => {
      console.log("onEnd");
      removeDocumentEvents();
    };
    const addDocumentEvents = () => {
      document.addEventListener(
        isPc.value ? "mousemove" : "touchmove",
        thrimgPosition
      );
      document.addEventListener(isPc.value ? "mouseup" : "touchend", onEnd);
    };
    const removeDocumentEvents = () => {
      document.removeEventListener(
        isPc.value ? "mousemove" : "touchmove",
        thrimgPosition
      );
      document.removeEventListener(isPc.value ? "mouseup" : "touchend", onEnd);
    };
    const keyDown = (e) => {
      e.preventDefault();
      const keyCode = e.keyCode;
      switch (keyCode) {
        // ESC
        case 27:
          value.value = false;
          break;
        // SPACE
        case 32:
          isFull.value = !isFull.value;
          break;
        // LEFT_ARROW
        case 37:
          next(-1);
          break;
        // UP_ARROW
        case 38:
          handleActions(0.2);
          break;
        // RIGHT_ARROW
        case 39:
          next(1);
          break;
        // DOWN_ARROW
        case 40:
          handleActions(-0.2);
          break;
      }
    };
    const Rolling = (e) => {
      const delta = e.wheelDelta ? e.wheelDelta : -e.detail;
      if (delta > 0) {
        handleActions(0.2);
      } else {
        handleActions(-0.2);
      }
    };
    const thrKeyDown = throttle(200, keyDown);
    const thrRolling = throttle(200, Rolling);
    // 放大缩小行为
    const handleActions = (sca) => {
      sca = parseFloat((sca + transform.scale).toFixed(3));
      console.log(sca);
      if (sca <= 0.2) {
        sca = 0.2;
      }
      transform.transition = true;
      transform.scale = sca;
    };
    //旋转行为
    const rotationActions = (deg) => {
      transform.transition = true;
      transform.rotate += deg;
    };
    const resetStyle = () => {
      transform.scale = 1;
      transform.rotate = 0;
      transform.top = 0;
      transform.left = 0;
      transform.transition = false;
    };
    // 上，下张图片行为
    const next = (next) => {
      resetStyle();
      currentIndex.value = utils.limit(
        currentIndex.value + next,
        0,
        props.urlList.length - 1
      );
    };
    console.log(props.url);

    // 获取当前图片[位置
    watch(
      () => props.url,
      () => {
        console.log("image-viw 更新链接");
        props.urlList.some((item, index) => {
          if (item == props.url) {
            currentIndex.value = index;
            return true;
          }
        });
      }
    );
    onUnmounted(() => {
      document.removeEventListener("keydown", thrKeyDown);
      document.removeEventListener("mousewheel", thrRolling);
    });
    let prevOverflow;
    const value = computed({
      get: () => {
        console.log(props.show);
        if (props.show) {
          prevOverflow = document.body.style.overflow;
          document.body.style.overflow = "hidden";
          resetStyle();
          document.addEventListener("keydown", thrKeyDown);
          document.addEventListener("mousewheel", thrRolling);
          ByImageView.value.focus();
        } else {
          console.log(prevOverflow);
          document.body.style.overflow = prevOverflow;
          document.removeEventListener("keydown", thrKeyDown);
          document.removeEventListener("mousewheel", thrRolling);
          removeDocumentEvents();
        }
        return props.show;
      },
      set: (val) => {
        context.emit("update:show", val);
      },
    });
    // 图片样式
    const imageStyle = computed(() => {
      return {
        transform: `scale(${transform.scale}) rotate(${transform.rotate}deg)`,
        "margin-left": `${transform.left}px`,
        "margin-top": `${transform.top}px`,
        "max-height": isFull.value ? "none" : "100%",
        "max-width": isFull.value ? "none" : "100%",
        transition: transform.transition ? "transform 0.3s" : "",
      };
    });

    return {
      currentIndex,
      isFull,
      isPc,
      imageStyle,
      ByImageView,
      onMousedown,
      next,
      handleActions,
      rotationActions,
      value,
    };
  },
};
</script>
<style lang="stylus" scoped>
.by-image-view {
  position: fixed;
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;

  .by-image-mask {
    position: absolute;
    width: 100%;
    height: 100%;
    top: 0;
    left: 0;
    opacity: 0.5;
    background: #000;
  }

  .by-image-btn {
    position: absolute;
    z-index: 1;
    display: flex;
    align-items: center;
    justify-content: center;
    border-radius: 50%;
    opacity: 0.8;
    cursor: pointer;
    box-sizing: border-box;
    user-select: none;
    cursor: pointer;
  }

  .by-image-close {
    top: 40px;
    right: 40px;
    width: 40px;
    height: 40px;
    font-size: 40px;
  }

  .by-image-prev, .by-image-next {
    top: 50%;
    transform: translateY(-50%);
    width: 44px;
    height: 44px;
    font-size: 24px;
    color: #fff;
    background-color: #606266;
    border-color: #fff;
  }

  .by-image-prev {
    left: 40px;
  }

  .by-image-next {
    right: 40px;
  }

  .by-image-view-actions {
    left: 50%;
    bottom: 30px;
    transform: translateX(-50%);
    width: 282px;
    height: 44px;
    padding: 0 23px;
    background-color: #606266;
    border-color: #fff;
    border-radius: 22px;

    svg {
      cursor: pointer;
    }
  }

  .by-image-view-actions-inner {
    width: 100%;
    height: 100%;
    text-align: justify;
    cursor: default;
    font-size: 23px;
    color: #fff;
    display: flex;
    align-items: center;
    justify-content: space-around;
  }

  .by-image-view-canvas {
    width: 100%;
    height: 100%;
    display: flex;
    justify-content: center;
    align-items: center;
    cursor: pointer;
  }
}

.fade-enter-active, .fade-leave-active {
  transition: opacity 0.2s ease;
}

.fade-enter-from, .fade-leave-to {
  opacity: 0;
}
</style>