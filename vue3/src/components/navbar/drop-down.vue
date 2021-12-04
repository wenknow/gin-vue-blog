<template>
  <div
    class="drop-down"
    :class="{ 'drop-down-float': float }"
    @mouseleave="hoverToggle"
    @mouseenter="hoverToggle"
  >
    <span class="drop-down-content" @click="clickToggle">
      <slot :visible="visible"></slot>
    </span>
    <transition :name="transiName">
      <div class="drop-down-menu" v-if="menu" v-show="visible">
        <slot name="menu"></slot>
      </div>
    </transition>
  </div>
</template>
<script>
import { ref } from "vue";
export default {
  props: {
    float: {
      type: Boolean,
      default: true,
    },
    mode: {
      type: String,
      default: "hover",
    },
    menu: {
      type: Boolean,
      default: true,
    },
    transiName: {
      type: String,
      default: "fade",
    },
  },
  setup(props) {
    const visible = ref(false);
    const clickToggle = () => {
      if (props.mode == "click") {
        visible.value = !visible.value;
      }
    };
    const hoverToggle = (e) => {
      if (props.mode == "click") {
        return;
      }
      if (e.type == "mouseleave") {
        visible.value = false;
      } else {
        visible.value = true;
      }
    };
    return {
      visible,
      clickToggle,
      hoverToggle,
    };
  },
};
</script>
<style lang="stylus" scoped>
.drop-down {
  height: 100%;

  .drop-down-content {
    display: flex;
    align-items: center;
    height: 100%;
  }

  .drop-down-menu {
    background-color: #fff;
    box-shadow: 0 0 2px rgba(0, 0, 0, 0.08), 0 14px 40px -10px rgba(0, 0, 0, 0.2);
    border-radius: 8px;
    padding: 8px;
    min-width: 100px;
  }
}

.drop-down-float {
  position: relative;

  .drop-down-menu {
    position: absolute;
    left: 50%;
    transform: translateX(-50%);
  }
}

.fade-enter-active, .fade-leave-active {
  transition: opacity 0.5s ease;
}

.fade-enter-from, .fade-leave-to {
  opacity: 0;
}

.left-enter-active, .left-leave-active {
  transition: all 0.5s;
  max-height: 4rem;
  max-width: 4rem;
}

.left-enter-from, .left-leave-to {
  opacity: 0;
  max-height: 0;
  max-width: 0;
}
</style>