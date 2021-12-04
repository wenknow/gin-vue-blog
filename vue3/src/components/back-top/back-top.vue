<template>
  <teleport to="#backTop">
    <transition name="fade">
      <div v-show="show" class="back-top" @click="back">
        <slot> </slot>
      </div>
    </transition>
  </teleport>
</template>

<script>
import { defineComponent, onUnmounted, ref } from "vue";
import utils from "../utils";
export default defineComponent({
  name: "ByBackTop",
  props: {
    targe: {
      type: Function,
      default: () => window,
    },
    visibilityHeight: {
      type: Number,
      default: 400,
    },
  },
  setup(props) {
    const node = document.createElement("div");
    node.id = "backTop";
    document.body.appendChild(node);
    const show = ref(false);
    const backTop = () => {
      let scrollTop = utils.getScroll(props.targe(), true);
      // console.log(scrollTop)
      if (scrollTop > props.visibilityHeight) {
        show.value = true;
      } else {
        show.value = false;
      }
    };
    document.addEventListener("scroll", backTop);
    const back = () => {
      utils.scrollTo(0, { getContainer: () => window });
    };
    onUnmounted(() => {
      document.body.removeChild(node);
      document.removeEventListener("scroll", backTop);
    });
    return {
      show,
      back,
    };
  },
});
</script>

<style lang="stylus" scoped>
.back-top {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
  color: rgba(0, 0, 0, 0.65);
  font-size: 14px;
  font-variant: tabular-nums;
  line-height: 1.5715;
  list-style: none;
  font-feature-settings: "tnum";
  position: fixed;
  right: 100px;
  bottom: 50px;
  z-index: 10;
  width: 40px;
  height: 40px;
  cursor: pointer;
}
@media (max-width: 768px) {
  .back-top {
    right: 1rem;
  }
}
</style>
