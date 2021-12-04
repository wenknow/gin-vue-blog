 <template>
  <transition name="by-mess" @after-leave="handleAfterLeave">
    <div class="ant-message-notice" v-if="visible">
      <div
        class="ant-message-notice-content"
        @mouseleave="!config.duration ? () => {} : startTimer()"
        @mouseenter="!config.duration ? () => {} : clearTimer()"
      >
        <div
          class="ant-message-custom-content"
          :class="'ant-message-' + config.type"
        >
          <component
            :is="config.icon ? config.icon : getDefaultIcon"
          ></component>
          <span>{{ config.content }}</span>
        </div>
      </div>
    </div>
  </transition>
</template>

<script>
import {
  CheckCircleOutlined,
  CloseCircleOutlined,
  ExclamationCircleOutlined,
  LoadingOutlined,
} from "@ant-design/icons-vue";
import { h } from "vue";
export default {
  name: "Message",
  props: {
    config: {
      type: Object,
      default: function () {
        return {
          top: `100px`,
          duration: 3000,
          type: "success",
        };
      },
    },
  },
  components: {
    CheckCircleOutlined,
    CloseCircleOutlined,
    ExclamationCircleOutlined,
    LoadingOutlined,
  },
  computed: {
    getDefaultIcon() {
      let icon;
      if (this.config.type == "success") {
        icon = this.$options.components.CheckCircleOutlined;
      } else if (this.config.type == "error") {
        icon = this.$options.components.CloseCircleOutlined;
      } else if (this.config.type == "warning") {
        icon = this.$options.components.ExclamationCircleOutlined;
      } else if (this.config.type == "loading") {
        icon = this.$options.components.LoadingOutlined;
      } else {
        icon = this.$options.components.ExclamationCircleOutlined;
      }
      return icon;
    },
  },
  data() {
    return {
      visible: false,
      timer: null,
    };
  },
  mounted() {
    this.visible = true;
    if (this.config.duration != 0) {
      this.setTimer(); // 组件挂载的时候开始计时，计时完成移除
    }
  },
  methods: {
    setTimer() {
      // 组件停留时间duration
      this.timer = setTimeout(() => {
        this.close();
      }, this.config.duration * 1000);
    },
    close() {
      this.visible = false;
      if (this.config.onClose) {
        this.config.onClose();
      }
      clearTimeout(this.timer);
    },
    clearTimer() {
      //鼠标移入提醒框时，清除定时器，提醒框持久化显示
      console.log("清除定时器，不会自动消失");
      clearTimeout(this.timer);
    },
    startTimer() {
      // 鼠标离开提醒框，定时器开始计时
      console.log("开始定时器");
      this.setTimer();
    },
    handleAfterLeave() {
      console.log("动画执行完成");
      // 移除动画执行完成后，销毁组件，并将其从父节点移除
      this.$emit("end", this.config.key);
    },
  },
  beforeUnmount() {
    console.log("组件死亡");
  },
  render() {
    console.log("render");
    // `<div><slot></slot></div>`
    return h("div", {}, this.$slots.default());
  },
  watch: {
    config: {
      handler: function (newVal) {
        if (newVal.duration == 0) {
          this.close();
        }
        this.clearTimer();
        this.setTimer();
      },
      deep: true,
    },
  },
};
</script>

<style lang="stylus">
.ant-message-notice {
  padding: 8px;
  text-align: center;

  .ant-message-notice-content {
    display: inline-block;
    padding: 10px 16px;
    background: #fff;
    border-radius: 2px;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.15);
    pointer-events: all;

    .anticon {
      position: relative;
      top: 1px;
      margin-right: 8px;
      font-size: 16px;
    }
  }

  .ant-message-info .anticon, .ant-message-loading .anticon {
    color: #1890ff;
  }

  .ant-message-warning .anticon {
    color: rgb(250, 173, 20);
  }

  .ant-message-error .anticon {
    color: rgb(245, 34, 45);
  }

  .ant-message-success .anticon {
    color: rgb(82, 196, 26);
  }
}

.by-mess-enter-active {
  animation: by-mess-In 0.3s ease;
}

.by-mess-leave-active {
  animation: by-mess-Out 0.3s ease;
}

@keyframes by-mess-In {
  0% {
    max-height: 0;
    padding: 0;
    opacity: 0;
  }

  100% {
    max-height: 150px;
    padding: 8;
    opacity: 1;
  }
}

@keyframes by-mess-Out {
  0% {
    max-height: 150px;
    padding: 8px;
    opacity: 1;
  }

  100% {
    max-height: 0;
    padding: 0;
    opacity: 0;
  }
}
</style>