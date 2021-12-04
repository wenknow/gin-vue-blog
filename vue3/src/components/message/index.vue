<template>
  <div class="ant-message">
    <span>
      <Message v-for="(item, key) of list" :key="key" :config="item" @end="end">
      </Message>
    </span>
  </div>
</template>
<script>
import Message from "./message";
export default {
  props: {
    defaults: Object,
  },
  components: {
    Message,
  },
  data() {
    return {
      list: {},
    };
  },
  methods: {
    end(key) {
      delete this.list[key];
    },
    add(config) {
      this.maxCount();
      if (config.type == "close") {
        this.close(config.key);
      } else {
        this.list[config.key] = config;
      }
    },
    close(key) {
      this.list[key].duration = 0;
    },
    maxCount() {
      const list = Object.keys(this.list);
      const length = list.length;
      if (this.defaults.maxCount && length >= this.defaults.maxCount) {
        list.forEach((key, index) => {
          //5 显示2个，删除0-3
          if (index <= length - this.defaults.maxCount) {
            this.close(key);
          }
        });
      }
    },
  },
};
</script>
<style lang="stylus" scoped>
.ant-message {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
  color: rgba(0, 0, 0, 0.65);
  font-size: 14px;
  font-variant: tabular-nums;
  line-height: 1.5715;
  list-style: none;
  font-feature-settings: 'tnum';
  position: fixed;
  top: 16px;
  left: 0;
  z-index: 1010;
  width: 100%;
  pointer-events: none;
}
</style>