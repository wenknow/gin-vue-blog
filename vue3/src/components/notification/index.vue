<template>
  <template v-for="(item, cKey) of types" :key="cKey">
    <div
      class="ant-notification"
      :class="`ant-notification-${cKey}`"
      :style="getStyle(cKey)"
    >
      <span>
        <Notification
          v-for="noti of item"
          :key="noti.key"
          :config="noti"
          @end="end"
        >
        </Notification>
      </span>
    </div>
  </template>
</template>
<script>
import Notification from "./notification";
export default {
  props: {
    defaults: Object,
  },
  components: {
    Notification,
  },
  data() {
    return {
      types: {
        topRight: {},
        topLeft: {},
        bottomRight: {},
        bottomLeft: {},
      },
    };
  },
  methods: {
    end(key) {
      for (let item in this.types) {
        if (this.types[item][key]) {
          delete this.types[item][key];
        }
      }
    },
    add(config) {
      this.maxCount(config.placement);
      if (config.type == "close") {
        this.close(config.key);
      } else {
        this.types[config.placement][config.key] = config;
      }
    },
    close(key) {
      for (let item in this.types) {
        if (this.types[item][key]) {
          this.types[item][key].duration = -1;
        }
      }
    },
    getStyle(type) {
      let style;
      if (type == "topRight" || type == "topLeft") {
        style = {
          top: this.defaults.top,
          bottom: "auto",
        };
        const key = type == "topRight" ? "right" : "left";
        style[key] = "0px";
      } else {
        style = {
          bottom: this.defaults.bottom,
          top: "auto",
        };
        const key = type == "bottomRight" ? "right" : "left";
        style[key] = "0px";
      }
      return style;
    },
    maxCount(type) {
      const list = Object.keys(this.types[type]);
      const length = list.length;
      if (this.defaults.maxCount && length >= this.defaults.maxCount) {
        list.forEach((key, index) => {
          //5 显示2个，删除0-3
          if (index <= length - this.defaults.maxCount) {
            this.types[type][key].duration = -1;
          }
        });
      }
    },
  },
  mounted() {
    // this.types[this.config.config.placement][this.config.config.key]=this.config.config
  },
};
</script>
<style lang="stylus" scoped>
.ant-notification {
  box-sizing: border-box;
  color: rgba(0, 0, 0, 0.65);
  font-size: 14px;
  line-height: 1.5715;
  font-feature-settings: 'tnum';
  position: fixed;
  z-index: 1010;
  width: 384px;
  max-width: calc(100vw - 32px);
  padding: 0px;
  font-variant: tabular-nums;
  list-style: none;
  margin: 0px 24px 0px 0px;
}

.ant-notification-bottomLeft, .ant-notification-topLeft {
  margin-right: 0px;
  margin-left: 24px;
}
</style>