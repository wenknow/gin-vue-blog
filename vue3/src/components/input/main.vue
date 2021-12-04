<template>
  <span class="ant-input-affix-wrapper" :class="{ focused: focused }">
    <span class="ant-input-prefix">
      <slot name="prefix"></slot>
    </span>
    <input
      v-model="inputVal"
      @focus="focused = true"
      @blur="focused = false"
      :type="type"
      class="ant-input"
    />
    <span class="ant-input-suffix">
      <slot name="suffix"></slot>
    </span>
  </span>
</template>

<script>
import { computed, ref } from "vue";
export default {
  name: "Input",
  props: {
    type: {
      type: String,
      default: "text",
    },
    value: String,
  },
  setup(props, context) {
    const inputVal = computed({
      get: () => props.value || "",
      set: (val) => {
        context.emit("update:modelValue", val);
      },
    });
    const focused = ref(false);
    return {
      inputVal,
      focused,
    };
  },
};
</script>

<style lang="stylus" scoped>
.ant-input-affix-wrapper {
  box-sizing: border-box;
  margin: 0;
  font-variant: tabular-nums;
  list-style: none;
  font-feature-settings: 'tnum';
  position: relative;
  display: inline-flex;
  border: 1px solid #d9d9d9;
  border-radius: 2px;
  padding: 4px 11px;
  width: 100%;
  text-align: start;
  background-color: #fff;
  background-image: none;
  color: rgba(0, 0, 0, 0.65);
  font-size: 14px;
  line-height: 1.5715;

  &:hover {
    border-color: #40a9ff;
    border-right-width: 1px !important;
  }

  .ant-input-prefix, .ant-input-affix-wrapper .ant-input-suffix {
    display: flex;
    align-items: center;
    color: rgba(0, 0, 0, 0.65);
    white-space: nowrap;
  }

  .ant-input {
    box-sizing: border-box;
    margin: 0;
    font-variant: tabular-nums;
    list-style: none;
    font-feature-settings: 'tnum';
    position: relative;
    display: inline-block;
    width: 100%;
    padding: 4px 11px;
    color: rgba(0, 0, 0, 0.65);
    font-size: 14px;
    line-height: 1.5715;
    background-color: #fff;
    background-image: none;
    border: 1px solid #d9d9d9;
    border-radius: 2px;
    transition: all 0.3s;
  }

  .ant-input {
    position: relative;
    text-align: inherit;
    border: none;
    padding: 0;
  }

  .ant-input-prefix {
    margin-right: 4px;
  }

  .ant-input-suffix {
    margin-left: 4px;
  }
}

.focused {
  border-color: #40a9ff;
  box-shadow: 0 0 0 2px rgba(24, 144, 255, 0.2);
}
</style>
