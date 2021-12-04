<template>
  <div class="by-skeleton" :class="skeletonClass">
    <template v-if="loading">
      <div class="by-skeleton-header" v-if="avatar">
        <span
          class="by-skeleton-avatar"
          :style="avatarClaSty.style"
          :class="avatarClaSty.aClass"
        ></span>
      </div>
      <div class="by-skeleton-content">
        <h3 class="by-skeleton-title" :style="titleStyle" v-if="title"></h3>
        <ul class="by-skeleton-paragraph" v-if="paragraph">
          <li
            v-for="(item, index) of parahList"
            :key="'paragraph' + index"
            :style="{ width: item ? `${item}%` : false }"
          ></li>
        </ul>
      </div>
    </template>
    <slot v-else></slot>
  </div>
</template>
<script>
import { computed } from "vue";
export default {
  name: "Skeleton",
  props: {
    active: {
      type: Boolean,
      default: false,
    },
    avatar: {
      type: [Boolean, Object],
      default: false,
    },
    loading: {
      type: Boolean,
      default: true,
    },
    paragraph: {
      type: [Boolean, Object],
      default: true,
    },
    title: {
      type: [Boolean, Object],
      default: true,
    },
  },
  setup(props) {
    const skeletonClass = computed(() => {
      const sClass = {};
      if (props.avatar) {
        sClass["by-skeleton-with-avatar"] = true;
      }
      if (props.active) {
        sClass["by-skeleton-active"] = true;
      }
      return sClass;
    });
    const avatarClaSty = computed(() => {
      const aClass = {};
      let style;
      aClass["by-skeleton-avatar-large"] = true;
      aClass["by-skeleton-avatar-circle"] = true;
      if (typeof props.avatar == "object") {
        if (typeof props.avatar.size == "number") {
          aClass["by-skeleton-avatar-large"] = false;
          style = {
            width: `${props.avatar.size}px`,
            height: `${props.avatar.size}px`,
            lineHeight: `${props.avatar.size}px`,
            fontSize: `${props.avatar.size / 2}px`,
          };
        } else {
          aClass["by-skeleton-avatar-large"] = false;
          aClass["by-skeleton-avatar-" + props.avatar.size] = true;
        }
        if (props.avatar.shape) {
          aClass["by-skeleton-avatar-circle"] = false;
          aClass["by-skeleton-avatar-" + props.avatar.shape] = true;
        }
      }
      return {
        aClass,
        style,
      };
    });
    const titleStyle = computed(() => {
      const style = {
        width: 38,
      };
      if (typeof props.title == "object") {
        style.width = props.title.width;
      }
      style.width = style.width + "%";
      return style;
    });
    const parahList = computed(() => {
      let list = new Array(3).fill(false);
      if (typeof props.paragraph == "object") {
        console.log("object");
        if (props.paragraph.rows) {
          list = new Array(props.paragraph.rows).fill(false);
        }
        if (props.paragraph.width instanceof Array) {
          list = props.paragraph.width;
        } else if (props.paragraph.width) {
          list[list.length - 1] = props.paragraph.width;
        }
      }
      return list;
    });
    return {
      skeletonClass,
      avatarClaSty,
      titleStyle,
      parahList,
    };
  },
};
</script>
<style lang="stylus">
.by-skeleton {
  display: table;
  width: 100%;

  .by-skeleton-header {
    display: table-cell;
    padding-right: 16px;
    vertical-align: top;

    .by-skeleton-avatar {
      display: inline-block;
      vertical-align: top;
      background: #f2f2f2;
      width: 32px;
      height: 32px;
      line-height: 32px;
    }

    .by-skeleton-avatar-circle {
      border-radius: 50%;
    }

    .by-skeleton-avatar-square {
      border-radius: 2px;
    }

    .by-skeleton-avatar-large {
      width: 40px;
      height: 40px;
      line-height: 40px;
    }

    .by-skeleton-avatar-samll {
      width: 24px;
      height: 24px;
      line-height: 24px;
    }
  }

  .by-skeleton-content {
    display: table-cell;
    width: 100%;
    vertical-align: top;

    .by-skeleton-title {
      width: 100%;
      height: 16px;
      margin: 0;
      margin-top: 16px;
      background: #f2f2f2;
    }

    .by-skeleton-paragraph {
      margin: 0;
      margin-top: 24px;
      padding: 0;
      list-style: none;

      li {
        width: 100%;
        height: 16px;
        list-style: none;
        background: #f2f2f2;
      }

      li+li {
        margin-top: 16px;
      }

      li:last-child:not(:first-child):not(:nth-child(2)) {
        width: 61%;
      }
    }
  }
}

.by-skeleton-with-avatar {
  .by-skeleton-content {
    .by-skeleton-title {
      margin-top: 12px;
    }

    .by-skeleton-paragraph {
      margin-top: 28px;
    }
  }
}

.by-skeleton-active {
  .by-skeleton-header .by-skeleton-avatar, .by-skeleton-content .by-skeleton-title, .by-skeleton-content .by-skeleton-paragraph>li {
    background: linear-gradient(90deg, #f2f2f2 25%, #e6e6e6 37%, #f2f2f2 63%);
    background-size: 400% 100%;
    animation: by-skeleton-loading 1.4s ease infinite;
  }
}

@keyframes by-skeleton-loading {
  0% {
    background-position: 100% 50%;
  }

  100% {
    background-position: 0 50%;
  }
}
</style>