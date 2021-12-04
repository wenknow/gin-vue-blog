<template>
  <span
    ref="byAvatar"
    class="by-avatar"
    :style="avararStyle"
    :class="avatarClass"
  >
    <span
      ref="byAvatarString"
      class="by-avatar-string"
      v-if="$slots.default"
      :style="stringStyle"
    >
      <slot></slot>
    </span>
    <img v-if="src" :src="src" />
    <slot name="icon" v-else>
      <component :is="icon"></component>
    </slot>
  </span>
</template>
<script>
import { computed, onMounted, onUpdated, reactive, ref, nextTick } from "vue";
export default {
  name: "ByAvatar",
  props: {
    icon: {
      type: Function,
    },
    shape: {
      type: String,
    },
    size: {
      type: [Number, String],
    },
    src: {
      type: String,
    },
    srcset: {
      type: String,
    },
    alt: {
      type: String,
    },
    loadError: {
      type: Function,
    },
  },
  setup(props, context) {
    const byAvatar = ref(null);
    const byAvatarString = ref(null);
    const avatarClass = computed(() => {
      const cla = {};
      if (props.src) {
        cla["by-avatar-image"] = true;
      }
      if (typeof props.size == "string") {
        const name = "by-avatar-" + props.size;
        cla[name] = true;
      }
      if (props.icon || context.slots.icon) {
        cla["by-avatar-icon"] = true;
      }
      if (props.shape == "square") {
        cla["by-avatar-square"] = true;
      }

      return cla;
    });
    const avararStyle = computed(() => {
      let style;
      if (typeof props.size == "number") {
        style = {
          width: `${props.size}px`,
          height: `${props.size}px`,
          lineHeight: `${props.size}px`,
          fontSize: `${props.size / 2}px`,
        };
      }
      return style;
    });
    const stringStyle = reactive({});
    const updateString = () => {
      if (context.slots["default"]) {
        console.log("更新模板");
        const avatarWidth = byAvatar.value.clientWidth;
        const stringWidth = byAvatarString.value.clientWidth;
        console.log(byAvatarString.value.clientWidth);
        let modeZoom;
        // 内接正方形
        modeZoom = Math.sqrt(Math.pow(avatarWidth / 2, 2) * 2) / stringWidth;
        if (props.shape) {
          //防止占满，去除4像素
          modeZoom = (avatarWidth - 4) / stringWidth;
        }
        console.log(avatarWidth, stringWidth);
        const zoom = avatarWidth > stringWidth ? 1 : modeZoom;
        stringStyle["transform"] = `scale(${zoom}) translateX(-50%)`;
      }
    };
    onMounted(() => {
      nextTick(() => {
        updateString();
      });
    });
    onUpdated(() => {
      nextTick(() => {
        updateString();
      });
    });
    return {
      avatarClass,
      avararStyle,
      byAvatar,
      byAvatarString,
      stringStyle,
    };
  },
  mounted() {
    console.log("mounted");
  },
};
</script>
<style lang="stylus" scoped>
.by-avatar {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
  color: rgba(0, 0, 0, 0.65);
  font-size: 14px;
  font-variant: tabular-nums;
  line-height: 1.5715;
  list-style: none;
  font-feature-settings: 'tnum';
  position: relative;
  display: inline-block;
  overflow: hidden;
  color: #fff;
  white-space: nowrap;
  text-align: center;
  vertical-align: middle;
  background: #ccc;
  width: 32px;
  height: 32px;
  line-height: 32px;
  border-radius: 50%;

  img {
    display: block;
    width: 100%;
    height: 100%;
    object-fit: cover;
  }

  .by-avatar-string {
    position: absolute;
    left: 50%;
    transform-origin: 0 center;
  }
}

.by-avatar-icon {
  font-size: 18px;
}

.by-avatar-image {
  background: transparent;
}

.by-avatar-large {
  width: 40px;
  height: 40px;
  line-height: 40px;
  border-radius: 50%;
}

.by-avatar-small {
  width: 24px;
  height: 24px;
  line-height: 24px;
  border-radius: 50%;
}

.by-avatar-square {
  border-radius: 2px;
}
</style>