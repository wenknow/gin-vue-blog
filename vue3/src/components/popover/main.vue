<template>
  <component
    class="popover-host"
    ref="popoverHost"
    tabindex="0"
    @click="clickShow"
    @focus="focusShow"
    @blur="blurHide"
    @mouseenter="enterShow"
    @mouseleave="leaveHide"
    :is="defaultSlot"
  ></component>
  <teleport to="body">
    <div
      class="ant-popover"
      ref="popover"
      :class="'ant-popover-placement-' + placement"
    >
      <div class="ant-popover-content">
        <div class="ant-popover-arrow"><!----></div>
        <div class="ant-popover-inner" role="tooltip">
          <div>
            <div class="ant-popover-title"></div>
            <div class="ant-popover-inner-content"></div>
          </div>
        </div>
      </div>
    </div>
  </teleport>
</template>
<script>
import { computed, ref, watch } from "vue";
export default {
  name: "Popover",
  props: {
    arrowPointAtCenter: {
      type: Boolean,
      default: false,
    },
    autoAdjustOverflow: {
      type: Boolean,
      default: true,
    },
    defaultVisible: {
      type: Boolean,
      default: false,
    },
    getPopupContainer: {
      type: Function,
      default: () => document.body,
    },
    mouseEnterDelay: {
      type: Number,
      default: 1,
    },
    mouseLeaveDelay: {
      type: Number,
      default: 0.1,
    },
    overlayClassName: {
      type: String,
    },
    overlayStyle: {
      type: Object,
    },
    placement: {
      type: String,
      default: "top",
    },
    trigger: {
      type: String,
      default: "hover",
    },
    visible: {
      type: Boolean,
      default: false,
    },
    destroyTooltipOnHie: {
      type: Boolean,
      default: false,
    },
    align: {
      type: Object,
    },
  },
  setup(props, context) {
    const mVisible = ref(props.visible);
    const popover = ref(null);
    const popoverHost = ref(null);
    watch(
      () => mVisible.value,
      (val) => {
        context.emit("update:visible", val);
      }
    );
    watch(
      () => props.visible,
      (val) => {
        mVisible.value = val;
      }
    );
    const getPosition = () => {
      let style = {
        left: 0,
        top: 0,
      };
      const host = popoverHost.value;
      const hostHeight = host.offsetHeight;
      const hostWidht = host.offsetWidth;
      const hostPosition = getElementPosition(host);
      console.log(hostHeight, hostWidht, hostPosition);
      return style;
    };
    const getElementPosition = (e) => {
      var x = 0,
        y = 0;
      while (e != null) {
        x += e.offsetLeft;
        y += e.offsetTop;
        e = e.offsetParent;
      }
      return { x: x, y: y };
    };
    const defaultSlot = computed(() => {
      const slot = context.slots.default ? context.slots.default() : "";
      if (!slot || !slot.length) {
        return;
      }
      return slot[0];
    });
    console.log(context.slots.default());
    const clickShow = () => {
      if (props.trigger == "click") {
        console.log("click");
        console.log(popoverHost.value);
        getPosition();
        popoverHost.value.focus();
        mVisible.value = !mVisible.value;
      }
    };
    const focusShow = () => {
      if (props.trigger == "focus") {
        console.log("focusShow");
        mVisible.value = true;
      }
    };
    const blurHide = () => {
      if (props.trigger != "hover") {
        console.log("blurHide");
        mVisible.value = false;
      }
    };
    const enterShow = () => {
      if (props.trigger == "hover") {
        console.log("enterShow");
      }
    };
    const leaveHide = () => {
      if (props.trigger == "hover") {
        console.log("leaveHide");
      }
    };
    return {
      mVisible,
      popover,
      popoverHost,
      defaultSlot,
      focusShow,
      blurHide,
      clickShow,
      enterShow,
      leaveHide,
    };
  },
};
</script>
<style lang="stylus" scoped>
.popover-host {
  outline: none;
}

.ant-popover {
  box-sizing: border-box;
  margin: 0;
  padding: 0;
  color: rgba(0, 0, 0, 0.65);
  font-size: 14px;
  font-variant: tabular-nums;
  line-height: 1.5715;
  list-style: none;
  font-feature-settings: 'tnum';
  position: absolute;
  top: 0;
  left: 0;
  z-index: 1030;
  font-weight: normal;
  white-space: normal;
  text-align: left;
  cursor: auto;
  user-select: text;
}

.ant-popover::after {
  position: absolute;
  background: rgba(255, 255, 255, 0.01);
  content: '';
}

.ant-popover-hidden {
  display: none;
}

.ant-popover-placement-top, .ant-popover-placement-topLeft, .ant-popover-placement-topRight {
  padding-bottom: 10px;
}

.ant-popover-placement-right, .ant-popover-placement-rightTop, .ant-popover-placement-rightBottom {
  padding-left: 10px;
}

.ant-popover-placement-bottom, .ant-popover-placement-bottomLeft, .ant-popover-placement-bottomRight {
  padding-top: 10px;
}

.ant-popover-placement-left, .ant-popover-placement-leftTop, .ant-popover-placement-leftBottom {
  padding-right: 10px;
}

.ant-popover-inner {
  background-color: #fff;
  background-clip: padding-box;
  border-radius: 2px;
  box-shadow: 0 2px 8px rgba(0, 0, 0, 0.15);
  box-shadow: 0 0 8px rgba(0, 0, 0, 0.15) \9;
}

@media screen and (-ms-high-contrast: active), (-ms-high-contrast: none) {
  /* IE10+ */
  .ant-popover-inner {
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.15);
  }
}

.ant-popover-title {
  min-width: 177px;
  min-height: 32px;
  margin: 0;
  padding: 5px 16px 4px;
  color: rgba(0, 0, 0, 0.85);
  font-weight: 500;
  border-bottom: 1px solid #f0f0f0;
}

.ant-popover-inner-content {
  padding: 12px 16px;
  color: rgba(0, 0, 0, 0.65);
}

.ant-popover-message {
  position: relative;
  padding: 4px 0 12px;
  color: rgba(0, 0, 0, 0.65);
  font-size: 14px;
}

.ant-popover-message > .anticon {
  position: absolute;
  top: 8px;
  font-size: 14px;
}

.ant-popover-message-title {
  padding-left: 22px;
}

.ant-popover-buttons {
  margin-bottom: 4px;
  text-align: right;
}

.ant-popover-buttons button {
  margin-left: 8px;
}

.ant-popover-arrow {
  position: absolute;
  display: block;
  width: 8.48528137423857px;
  height: 8.48528137423857px;
  background: transparent;
  border-style: solid;
  border-width: 4.242640687119285px;
  transform: rotate(45deg);
}

.ant-popover-placement-top > .ant-popover-content > .ant-popover-arrow, .ant-popover-placement-topLeft > .ant-popover-content > .ant-popover-arrow, .ant-popover-placement-topRight > .ant-popover-content > .ant-popover-arrow {
  bottom: 6.2px;
  border-top-color: transparent;
  border-right-color: #fff;
  border-bottom-color: #fff;
  border-left-color: transparent;
  box-shadow: 3px 3px 7px rgba(0, 0, 0, 0.07);
}

.ant-popover-placement-top > .ant-popover-content > .ant-popover-arrow {
  left: 50%;
  transform: translateX(-50%) rotate(45deg);
}

.ant-popover-placement-topLeft > .ant-popover-content > .ant-popover-arrow {
  left: 16px;
}

.ant-popover-placement-topRight > .ant-popover-content > .ant-popover-arrow {
  right: 16px;
}

.ant-popover-placement-right > .ant-popover-content > .ant-popover-arrow, .ant-popover-placement-rightTop > .ant-popover-content > .ant-popover-arrow, .ant-popover-placement-rightBottom > .ant-popover-content > .ant-popover-arrow {
  left: 6px;
  border-top-color: transparent;
  border-right-color: transparent;
  border-bottom-color: #fff;
  border-left-color: #fff;
  box-shadow: -3px 3px 7px rgba(0, 0, 0, 0.07);
}

.ant-popover-placement-right > .ant-popover-content > .ant-popover-arrow {
  top: 50%;
  transform: translateY(-50%) rotate(45deg);
}

.ant-popover-placement-rightTop > .ant-popover-content > .ant-popover-arrow {
  top: 12px;
}

.ant-popover-placement-rightBottom > .ant-popover-content > .ant-popover-arrow {
  bottom: 12px;
}

.ant-popover-placement-bottom > .ant-popover-content > .ant-popover-arrow, .ant-popover-placement-bottomLeft > .ant-popover-content > .ant-popover-arrow, .ant-popover-placement-bottomRight > .ant-popover-content > .ant-popover-arrow {
  top: 6px;
  border-top-color: #fff;
  border-right-color: transparent;
  border-bottom-color: transparent;
  border-left-color: #fff;
  box-shadow: -2px -2px 5px rgba(0, 0, 0, 0.06);
}

.ant-popover-placement-bottom > .ant-popover-content > .ant-popover-arrow {
  left: 50%;
  transform: translateX(-50%) rotate(45deg);
}

.ant-popover-placement-bottomLeft > .ant-popover-content > .ant-popover-arrow {
  left: 16px;
}

.ant-popover-placement-bottomRight > .ant-popover-content > .ant-popover-arrow {
  right: 16px;
}

.ant-popover-placement-left > .ant-popover-content > .ant-popover-arrow, .ant-popover-placement-leftTop > .ant-popover-content > .ant-popover-arrow, .ant-popover-placement-leftBottom > .ant-popover-content > .ant-popover-arrow {
  right: 6px;
  border-top-color: #fff;
  border-right-color: #fff;
  border-bottom-color: transparent;
  border-left-color: transparent;
  box-shadow: 3px -3px 7px rgba(0, 0, 0, 0.07);
}

.ant-popover-placement-left > .ant-popover-content > .ant-popover-arrow {
  top: 50%;
  transform: translateY(-50%) rotate(45deg);
}

.ant-popover-placement-leftTop > .ant-popover-content > .ant-popover-arrow {
  top: 12px;
}

.ant-popover-placement-leftBottom > .ant-popover-content > .ant-popover-arrow {
  bottom: 12px;
}
</style>