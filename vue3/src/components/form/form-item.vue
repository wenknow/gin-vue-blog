<template>
  <div class="form-item">
    <div class="form-item-label"></div>
    <div class="form-item-control-wrapper">
      <div class="form-item-children">
        <slot></slot>
      </div>
      <div class="form-explain"></div>
    </div>
  </div>
</template>
<script>
export default {
  name: "FormItem",
  props: {
    label: String,
    namme: String,
  },
};
</script>
<style lang="stylus" scoped>
.form-item {
  box-sizing: border-box;
  padding: 0;
  color: rgba(0, 0, 0, 0.65);
  font-size: 14px;
  font-variant: tabular-nums;
  line-height: 1.5715;
  list-style: none;
  font-feature-settings: 'tnum';
  margin: 0 0 24px;
  vertical-align: top;

  .form-item-label {
    display: inline-block;
    overflow: hidden;
    line-height: 39.9999px;
    white-space: nowrap;
    text-align: right;
    vertical-align: middle;
  }

  .form-item-control-wrapper {
    .form-item-children {
      position: relative;
      line-height: 40px;
      zoom: 1;
    }

    .form-explain {
      color: #f5222d;
      margin-bottom: -1px;
      clear: both;
      min-height: 23px;
      margin-top: -2px;
      font-size: 14px;
      line-height: 1.5715;
      transition: color 0.3s cubic-bezier(0.215, 0.61, 0.355, 1);
    }
  }
}
</style>