<template>
  <transition v-on="on" :class="{ 'collapse-transition': show }">
    <slot></slot>
  </transition>
</template>
<script>
import { ref } from "vue";
export default {
  name: "collapseTransition",
  setup() {
    const show = ref(false);
    return {
      show,
      on: {
        beforeEnter(el) {
          show.value = true;
          if (!el.dataset) el.dataset = {};
          el.dataset.oldPaddingTop = el.style.paddingTop;
          el.dataset.oldPaddingBottom = el.style.paddingBottom;
          el.style.height = "0";
          el.style.paddingTop = 0;
          el.style.paddingBottom = 0;
        },
        enter(el) {
          el.dataset.oldOverflow = el.style.overflow;
          if (el.scrollHeight !== 0) {
            el.style.height = el.scrollHeight + "px";
            el.style.paddingTop = el.dataset.oldPaddingTop;
            el.style.paddingBottom = el.dataset.oldPaddingBottom;
          } else {
            el.style.height = "";
            el.style.paddingTop = el.dataset.oldPaddingTop;
            el.style.paddingBottom = el.dataset.oldPaddingBottom;
          }
          el.style.overflow = "hidden";
        },
        afterEnter(el) {
          // for safari: remove class then reset height is necessary
          show.value = false;
          el.style.height = "";
          el.style.overflow = el.dataset.oldOverflow;
        },
        beforeLeave(el) {
          if (!el.dataset) el.dataset = {};
          el.dataset.oldPaddingTop = el.style.paddingTop;
          el.dataset.oldPaddingBottom = el.style.paddingBottom;
          el.dataset.oldOverflow = el.style.overflow;
          el.style.height = el.scrollHeight + "px";
          el.style.overflow = "hidden";
        },
        leave(el) {
          if (el.scrollHeight !== 0) {
            // for safari: add class after set height, or it will jump to zero height suddenly, weired
            show.value = true;
            el.style.height = 0;
            el.style.paddingTop = 0;
            el.style.paddingBottom = 0;
          }
        },
        afterLeave(el) {
          show.value = false;
          el.style.height = "";
          el.style.overflow = el.dataset.oldOverflow;
          el.style.paddingTop = el.dataset.oldPaddingTop;
          el.style.paddingBottom = el.dataset.oldPaddingBottom;
        },
      },
    };
  },
};
</script>
<style scoped>
.collapse-transition {
  transition: height 0.2s ease-in-out, padding-top 0.2s ease-in-out,
    padding-bottom 0.2s ease-in-out !important;
}
</style>