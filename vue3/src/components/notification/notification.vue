 <template>
  <transition name="noti-fade" @after-leave="handleAfterLeave">
    <div
      class="ant-notification-notice ant-notification-notice-closable"
      v-if="visible"
      @mouseleave="!config.duration ? () => {} : startTimer()"
      @mouseenter="!config.duration ? () => {} : clearTimer()"
    >
      <div class="ant-notification-notice-content">
        <div
          @click="config.onClick ? onClick() : () => {}"
          :class="
            config.icon || config.type
              ? 'ant-notification-notice-with-icon'
              : ''
          "
        >
          <!---->
          <component
            class="ant-notification-notice-icon"
            :class="'ant-notification-notice-icon-' + config.type"
            :is="config.icon ? config.icon : getDefaultIcon"
          ></component>
          <div class="ant-notification-notice-message">
            <!---->{{ config.message }}
          </div>
          <div class="ant-notification-notice-description">
            {{ config.description }}
          </div>

          <!---->
        </div>
        <span class="ant-notification-notice-btn" v-if="config.btn">
          <component :is="config.btn" @click="close" />
        </span>
      </div>
      <a tabindex="0" class="ant-notification-notice-close">
        <component
          @click="close"
          :is="
            config.closeIcon
              ? config.closeIcon
              : $options.components.CloseOutlined
          "
        ></component>
      </a>
    </div>
  </transition>
</template>

<script>
import {
  CheckCircleOutlined,
  CloseCircleOutlined,
  ExclamationCircleOutlined,
  LoadingOutlined,
  CloseOutlined,
} from "@ant-design/icons-vue";
export default {
  name: "Message",
  props: {
    config: Object,
  },
  components: {
    CheckCircleOutlined,
    CloseCircleOutlined,
    ExclamationCircleOutlined,
    LoadingOutlined,
    CloseOutlined,
  },
  computed: {
    getDefaultIcon() {
      let icon;
      if (this.config.type == "success") {
        icon = this.$options.components.CheckCircleOutlined;
      } else if (this.config.type == "error") {
        icon = this.$options.components.CloseCircleOutlined;
      } else if (this.config.type == "warning" || this.config.type == "info") {
        icon = this.$options.components.ExclamationCircleOutlined;
      } else if (this.config.type == "loading") {
        icon = this.$options.components.LoadingOutlined;
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
    onClick() {
      this.config.onClick();
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
  watch: {
    config: {
      handler: function (newVal) {
        if (newVal.duration < 0) {
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
.ant-notification-notice {
  position: relative;
  margin-bottom: 16px;
  line-height: 1.5;
  box-shadow: rgba(0, 0, 0, 0.15) 0px 4px 12px;
  padding: 16px 24px;
  overflow: hidden;
  background: rgb(255, 255, 255);
  border-radius: 2px;
  outline: none;

  .ant-notification-notice-content {
    .ant-notification-notice-with-icon {
      .ant-notification-notice-message {
        margin-bottom: 4px;
        margin-left: 48px;
        font-size: 16px;
      }

      .ant-notification-notice-description {
        margin-left: 48px;
        font-size: 14px;
      }

      .ant-notification-notice-icon {
        position: absolute;
        margin-left: 4px;
        font-size: 24px;
        line-height: 24px;
      }

      .ant-notification-notice-icon-info, .ant-notification-notice-icon-loading {
        color: #1890ff;
      }

      .ant-notification-notice-icon-warning {
        color: rgb(250, 173, 20);
      }

      .ant-notification-notice-icon-error {
        color: rgb(245, 34, 45);
      }

      .ant-notification-notice-icon-success {
        color: rgb(82, 196, 26);
      }
    }

    .ant-notification-notice-message {
      display: inline-block;
      margin-bottom: 8px;
      color: rgba(0, 0, 0, 0.85);
      font-size: 16px;
      line-height: 24px;
    }

    .ant-notification-notice-description {
      font-size: 14px;
    }
  }

  .ant-notification-notice-btn {
    float: right;
    margin-top: 16px;
  }

  .ant-notification-notice-close {
    position: absolute;
    top: 16px;
    right: 22px;
    color: rgba(0, 0, 0, 0.45);
    outline: none;

    .ant-notification-close-icon {
      font-size: 14px;
      cursor: pointer;
      outline: none;
    }

    .anticon {
      display: inline-block;
      color: inherit;
      font-style: normal;
      line-height: 0;
      text-align: center;
      text-transform: none;
      vertical-align: -0.125em;
      text-rendering: optimizelegibility;
      -webkit-font-smoothing: antialiased;
    }
  }
}

.noti-fade-enter-active {
  animation: noti-fadeIn 0.24s ease;
}

.noti-fade-leave-active {
  animation: noti-fadeOut 0.24s ease;
}

.ant-notification-topLeft, .ant-notification-bottomLeft {
  .fade-enter-active {
    animation: noti-leftFadeIn 0.24s ease;
  }
}

@keyframes noti-fadeIn {
  0% {
    left: 384px;
    opacity: 0;
  }

  100% {
    left: 0;
    opacity: 1;
  }
}

@keyframes noti-leftFadeIn {
  0% {
    right: 384px;
    opacity: 0;
  }

  100% {
    right: 0;
    opacity: 1;
  }
}

@keyframes noti-fadeOut {
  0% {
    max-height: 150px;
    margin-bottom: 16px;
    padding-top: 16px 24px;
    padding-bottom: 16px 24px;
    opacity: 1;
  }

  100% {
    max-height: 0;
    margin-bottom: 0;
    padding-top: 0;
    padding-bottom: 0;
    opacity: 0;
  }
}
</style>