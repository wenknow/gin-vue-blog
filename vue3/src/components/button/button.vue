<template>
  <button
    class="by-btn"
    :class="byBtn"
    :type="htmlType"
    :disabled="disabled"
    ref="ByButton"
    @click="clickAnimating"
  >
    <slot name="icon"></slot>
    <span>
      <slot></slot>
    </span>
  </button>
</template>
<script>
const color = document.createElement("style");
export default {
  name: "ByButton",
  props: {
    type: {
      default: "default",
      type: String,
      validator: function (value) {
        // 这个值必须匹配下列字符串中的一个
        return (
          ["primary", "dashed", "danger", "link", "default"].indexOf(value) !==
          -1
        );
      },
    },
    htmlType: {
      default: "",
      type: String,
    },
    disabled: {
      default: false,
      type: Boolean,
    },
    block: {
      default: false,
      type: Boolean,
    },
    shape: {
      default: "round",
      type: String,
      validator: function (value) {
        // 这个值必须匹配下列字符串中的一个
        return ["circle", "round"].indexOf(value) !== -1;
      },
    },
    size: {
      default: "default",
      type: String,
      validator: function (value) {
        // 这个值必须匹配下列字符串中的一个
        return ["large", "small", "default"].indexOf(value) !== -1;
      },
    },
  },
  data() {
    return {
      animating: false,
      animaId: 0,
      animatingColor: color,
    };
  },
  computed: {
    byBtn() {
      return {
        "by-btn-primary": this.type == "primary",
        "by-btn-dashed": this.type == "dashed",
        "by-btn-danger": this.type == "danger",
        "by-btn-link": this.type == "link",
        "by-block": this.block,
        "by-btn-click-animating": this.animating,
        "by-btn-lg": this.size == "large",
        "by-btn-sm": this.size == "small",
        "by-btn-circle": this.shape == "circle" && this.type != "link",
      };
    },
  },
  methods: {
    clickAnimating() {
      if (this.type != "link") {
        console.log(this.animatingColor);
        this.animatingColor.innerText = `.by-btn-click-animating::after {--by-btn--animation--color:${
          this.type != "danger" ? "rgb(54, 159, 249);" : "rgb(249, 110, 108);"
        }}`;
        document.body.appendChild(this.animatingColor);
        this.animating = false;
        setTimeout(() => {
          this.animating = true;
        }, 50);
        clearTimeout(this.animaId);
        this.animaId = setTimeout(() => {
          this.animating = false;
          this.animatingColor.innerText = "";
        }, 2000);
      }
    },
  },
};
</script>
<style lang="stylus" scoped>
.by-btn {
  line-height: 1.5715;
  position: relative;
  display: inline-block;
  font-weight: 400;
  white-space: nowrap;
  text-align: center;
  background-image: none;
  box-shadow: 0 2px 0 rgba(0, 0, 0, 0.015);
  cursor: pointer;
  transition: all 0.3s cubic-bezier(0.645, 0.045, 0.355, 1);
  user-select: none;
  touch-action: manipulation;
  height: 32px;
  padding: 0 15px;
  font-size: 14px;
  border-radius: 2px;
  color: rgba(0, 0, 0, 0.65);
  background-color: #fff;
  border: 1px solid #d9d9d9;
}

.by-btn:focus, .by-btn:hover {
  color: #40a9ff;
  background-color: #fff;
  border-color: #40a9ff;
}

.by-btn:active, .by-btn:focus {
  outline: 0;
}

.by-block {
  width: 100%;
}

.by-btn-primary {
  color: #fff;
  background-color: #1890ff;
  border-color: #1890ff;
  text-shadow: 0 -1px 0 rgba(0, 0, 0, 0.12);
  box-shadow: 0 2px 0 rgba(0, 0, 0, 0.045);
}

.by-btn-primary:focus, .by-btn-primary:hover {
  color: #fff;
  background-color: #40a9ff;
  border-color: #40a9ff;
}

.by-btn-dashed {
  color: rgba(0, 0, 0, 0.65);
  background-color: #fff;
  border-color: #d9d9d9;
  border-style: dashed;
}

.by-btn-danger {
  color: #fff;
  background-color: #ff4d4f;
  border-color: #ff4d4f;
  text-shadow: 0 -1px 0 rgba(0, 0, 0, 0.12);
  box-shadow: 0 2px 0 rgba(0, 0, 0, 0.045);
}

.by-btn-danger:focus, .by-btn-danger:hover {
  color: #fff;
  background-color: #ff7875;
  border-color: #ff7875;
}

.by-btn-link {
  color: #1890ff;
  background-color: transparent;
  border-color: transparent;
  box-shadow: none;
}

.by-btn-link:active, .by-btn-link:focus, .by-btn-link:hover {
  color: #40a9ff;
  border-color: transparent;
}

.by-btn[disabled] {
  color: rgba(0, 0, 0, 0.25);
  background-color: #f5f5f5;
  border-color: #d9d9d9;
  text-shadow: none;
  box-shadow: none;
  cursor: not-allowed;
}

.by-btn-link[disabled] {
  color: rgba(0, 0, 0, 0.25);
  background-color: transparent;
  border-color: transparent;
  text-shadow: none;
  box-shadow: none;
}

.by-btn-click-animating::after {
  position: absolute;
  top: 0;
  right: 0;
  bottom: 0;
  left: 0;
  display: block;
  border-radius: inherit;
  box-shadow: 0 0 0 0 #1890ff;
  box-shadow: 0 0 0 0 var(--by-btn--animation--color);
  opacity: 0.2;
  animation: fadeEffect 2s cubic-bezier(0.08, 0.82, 0.17, 1), waveEffect 0.4s cubic-bezier(0.08, 0.82, 0.17, 1);
  animation-fill-mode: forwards;
  content: '';
  pointer-events: none;
}

.by-btn-sm {
  height: 24px;
  padding: 0 7px;
  font-size: 14px;
  border-radius: 2px;
}

.by-btn-lg {
  height: 40px;
  padding: 0 15px;
  font-size: 16px;
  border-radius: 2px;
}

.by-btn-circle {
  min-width: 32px;
  padding-right: 0;
  padding-left: 0;
  text-align: center;
  border-radius: 50%;
}

@keyframes waveEffect {
  to {
    box-shadow: 0 0 0 #1890ff;
    box-shadow: 0 0 0 6px var(--by-btn--animation--color);
  }
}

@keyframes fadeEffect {
  to {
    opacity: 0;
  }
}
</style>