<template>
  <svg :class="iconStyle" class="icon" aria-hidden="true" fill="currentColor">
    <use :xlink:href="iconName"></use>
  </svg>
</template>
 
<script>
export default {
  name: "icon",
  props: {
    type: {
      type: String,
      required: true,
    },
    iconStyle: String,
  },
  computed: {
    iconName() {
      return `#${this.type}`;
    },
  },
};
</script>
<style lang="stylus" scoped>
.icon {
  width: 1em;
  height: 1em;
  vertical-align: -0.15em;
  fill: currentColor;
  overflow: hidden;
}
</style>