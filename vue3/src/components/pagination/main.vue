<template>
  <ul class="by-pagination" unselectable="unselectable">
    <li class="by-pagination-total-text" v-if="showTotal">
      {{ showTotal(total, [(mCurrent - 1) * mPageSize, mCurrent * mPageSize]) }}
    </li>
    <!-- 上一页 -->
    <li
      title="上一页"
      tabindex="0"
      @click="prev"
      class="by-pagination-prev"
      :class="{ 'by-pagination-disabled': mCurrent == 1 }"
    >
      <span class="by-pagination-item-link">prev</span>
    </li>

    <!-- 第一页 -->
    <li
      @click="jump(1)"
      v-if="mCurrent != 1"
      tabindex="0"
      class="by-pagination-item"
    >
      <span>1</span>
    </li>

    <!-- 向前快速跳转。显示条件：向前跳转5页后大于等于第一页- -->
    <li
      v-if="prevJumpShow"
      tabindex="0"
      @click="fastJump(-5)"
      class="by-pagination-jump-prev"
    >
      <div class="by-pagination-item-container by-pagination-item-link">
        <svg
          class="by-pagination-item-link-icon"
          data-icon="double-left"
          width="1em"
          height="1em"
          fill="currentColor"
          aria-hidden="true"
          viewBox="64 64 896 896"
          focusable="false"
        >
          <path
            d="M272.9 512l265.4-339.1c4.1-5.2.4-12.9-6.3-12.9h-77.3c-4.9 0-9.6 2.3-12.6 6.1L186.8 492.3a31.99 31.99 0 000 39.5l255.3 326.1c3 3.9 7.7 6.1 12.6 6.1H532c6.7 0 10.4-7.7 6.3-12.9L272.9 512zm304 0l265.4-339.1c4.1-5.2.4-12.9-6.3-12.9h-77.3c-4.9 0-9.6 2.3-12.6 6.1L490.8 492.3a31.99 31.99 0 000 39.5l255.3 326.1c3 3.9 7.7 6.1 12.6 6.1H836c6.7 0 10.4-7.7 6.3-12.9L576.9 512z"
          ></path>
        </svg>
        <span class="by-pagination-item-ellipsis">•••</span>
      </div>
    </li>

    <!-- 前两页 -->
    <li
      v-for="(item, index) of list.prevList"
      tabindex="0"
      @click="jump(item)"
      :key="'page' + item"
      class="by-pagination-item"
      :class="{
        'by-pagination-item-after-jump-prev': index == 0 && prevJumpShow,
      }"
    >
      <span>{{ item }}</span>
    </li>
    <!-- 当前页 -->
    <li class="by-pagination-item by-pagination-item-active" tabindex="0">
      <span>{{ mCurrent }}</span>
    </li>
    <!-- 后两页 -->
    <li
      v-for="(item, index) of list.nextList"
      :key="'page' + item"
      @click="jump(item)"
      tabindex="0"
      class="by-pagination-item"
      :class="{
        'by-pagination-item-before-jump-next':
          index == list.nextList.length - 1 && nextJumpShow,
      }"
    >
      <span>{{ item }}</span>
    </li>

    <!-- 向后快速跳转。显示条件：向后跳转5页后小于等于最后一页-->
    <li
      v-if="nextJumpShow"
      @click="fastJump(5)"
      tabindex="0"
      class="by-pagination-jump-next"
    >
      <div class="by-pagination-item-container by-pagination-item-link">
        <svg
          class="by-pagination-item-link-icon"
          data-icon="double-right"
          width="1em"
          height="1em"
          fill="currentColor"
          aria-hidden="true"
          viewBox="64 64 896 896"
          focusable="false"
        >
          <path
            d="M533.2 492.3L277.9 166.1c-3-3.9-7.7-6.1-12.6-6.1H188c-6.7 0-10.4 7.7-6.3 12.9L447.1 512 181.7 851.1A7.98 7.98 0 00188 864h77.3c4.9 0 9.6-2.3 12.6-6.1l255.3-326.1c9.1-11.7 9.1-27.9 0-39.5zm304 0L581.9 166.1c-3-3.9-7.7-6.1-12.6-6.1H492c-6.7 0-10.4 7.7-6.3 12.9L751.1 512 485.7 851.1A7.98 7.98 0 00492 864h77.3c4.9 0 9.6-2.3 12.6-6.1l255.3-326.1c9.1-11.7 9.1-27.9 0-39.5z"
          ></path>
        </svg>
        <span class="by-pagination-item-ellipsis">•••</span>
      </div>
    </li>

    <!-- 最后一页 -->
    <li
      v-if="mCurrent != lastPage"
      @click="jump(lastPage)"
      tabindex="0"
      class="by-pagination-item"
      :class="{ 'by-pagination-item-active': mCurrent == lastPage }"
    >
      <span>{{ lastPage }}</span>
    </li>

    <!-- 下一页一页 -->
    <li
      title="下一页"
      @click="next"
      class="by-pagination-next"
      tabindex="0"
      :class="{ 'by-pagination-disabled': mCurrent == lastPage }"
    >
      <span class="by-pagination-item-link">next</span>
    </li>
    <li class="by-pagination-options" v-if="showSizeChanger || showQuickJumper">
      <div class="by-pagination-options-size-changer" v-if="showSizeChanger">
        <select @change="mPageSize = $event.target.value">
          <option v-for="item of pageSizeOptions" :key="item" :value="item">
            <slot name="buildOptionText" :value="item">{{ item }}</slot>
          </option>
        </select>
      </div>
      <div v-if="showQuickJumper" class="by-pagination-options-quick-jumper">
        跳至<input
          type="text"
          @keyup.enter="jump($event.target.value)"
        />页<!---->
      </div>
    </li>
  </ul>
</template>
<script>
import { computed, ref } from "vue";
import utils from "../utils.js";
export default {
  name: "Pagination",
  props: {
    current: {
      type: Number,
    },
    pageSize: {
      type: Number,
    },
    defaultPageSize: {
      type: Number,
      default: 10,
    },
    disabled: {
      type: Boolean,
      default: false,
    },
    hideOnSinglePag: {
      type: Boolean,
      default: false,
    },
    itemRender: {
      type: Object,
    },
    pageSizeOptions: {
      type: Array,
      default: () => ["10", "20", "30", "40"],
    },
    showLessItems: {
      type: Boolean,
      default: false,
    },
    showQuickJumper: {
      type: Boolean,
      default: false,
    },
    showSizeChanger: {
      type: Boolean,
      default: false,
    },
    showTotal: {
      type: Function,
    },
    simple: {
      type: Boolean,
    },
    size: {
      type: String,
      default: "",
    },
    total: {
      type: Number,
      default: 0,
    },
  },
  setup(props, context) {
    // v-model:current

    const mSelfCurrent = ref(1);
    const mCurrent = computed({
      get: () => {
        if (!props.current) {
          return utils.limit(mSelfCurrent.value, 1, lastPage, false);
        }
        return props.current;
      },
      set: (val) => {
        val = parseInt(val);
        if (isNaN(val)) {
          return;
        }
        if (val == mSelfCurrent.value) {
          return;
        }
        console.log("mCurrent");
        mSelfCurrent.value = val;
        context.emit("update:current", val);
        context.emit("change", val, mPageSize.value);
      },
    });
    // v-model:pageSize
    const mSelfPageSize = ref(props.defaultPageSize);
    const mPageSize = computed({
      get: () => {
        if (!props.pageSize) {
          return mSelfPageSize.value;
        }
        return props.pageSize;
      },
      set: (val) => {
        val = parseInt(val);
        console.log(val);
        if (isNaN(val)) {
          return;
        }
        mSelfPageSize.value = val;
        context.emit("update:pageSize", val);
        context.emit("show-size-change", mCurrent.value, val);
      },
    });
    // 最后一页
    const lastPage = computed(() => {
      const last = Math.ceil(props.total / mPageSize.value);
      mCurrent.value = utils.limit(mCurrent.value, 1, last, false);
      return last;
    });
    const prevJumpShow = computed(() => {
      return mCurrent.value - 4 >= 1 && lastPage.value >= 7;
    });
    const nextJumpShow = computed(() => {
      return mCurrent.value + 4 <= lastPage.value && lastPage.value >= 7;
    });
    // 跳转到第几页
    const jump = (current) => {
      if (props.disabled) {
        return;
      }
      current = parseInt(current);
      if (isNaN(current)) {
        return;
      }
      console.log("jump", current);
      current = utils.limit(current, 1, lastPage.value, false);
      console.log(current);
      mCurrent.value = current;
    };
    // 上一页
    const prev = () => {
      if (mCurrent.value == 1) {
        return;
      }
      jump(mCurrent.value - 1);
    };
    // 下一页
    const next = () => {
      if (mCurrent.value == lastPage.value) {
        return;
      }
      jump(mCurrent.value + 1);
    };
    // 快速跳转几页
    const fastJump = (page) => {
      jump(mCurrent.value + page);
    };
    // 循环到几，是否为下列表
    const prevNextList = (count, isNext = true) => {
      let data = [];
      for (let i = 1; i <= count; i++) {
        if (isNext) {
          if (mCurrent.value + i < lastPage.value) {
            data.push(mCurrent.value + i);
          } else {
            break;
          }
        } else {
          if (mCurrent.value - i > 1) {
            data.push(mCurrent.value - i);
          } else {
            break;
          }
        }
      }
      if (!isNext) {
        return data.reverse();
      }
      return data;
    };
    const list = computed(() => {
      const count = 4;
      let prevList = prevNextList(2, false);
      let nextList = prevNextList(2);
      const prevLeng = prevList.length;
      const nextLeng = nextList.length;

      if (prevLeng != 2) {
        nextList = prevNextList(count - prevLeng);
      }

      if (nextList != 2) {
        prevList = prevNextList(count - nextLeng, false);
      }
      return {
        prevList,
        nextList,
      };
    });
    return {
      mCurrent,
      mPageSize,
      lastPage,
      jump,
      next,
      prev,
      fastJump,
      list,
      prevJumpShow,
      nextJumpShow,
    };
  },
};
</script>
<style lang="stylus" scoped>
.by-pagination {
  box-sizing: border-box;
  color: rgba(0, 0, 0, 0.65);
  font-size: 14px;
  font-variby: tabular-nums;
  line-height: 1.5715;
  font-feature-settings: 'tnum';
  list-style: none;
  margin: 0;
  padding: 0;

  li {
    outline: 0;
  }

  .by-pagination-item {
    display: inline-block;
    box-sizing: border-box;
    min-width: 32px;
    height: 32px;
    margin-right: 8px;
    font-family: Arial;
    line-height: 30px;
    text-align: center;
    vertical-align: middle;
    list-style: none;
    background-color: #fff;
    border: 1px solid #d9d9d9;
    border-radius: 2px;
    outline: 0;
    cursor: pointer;
    user-select: none;

    &:focus, &:hover {
      border-color: #1890ff;
      transition: all 0.3s;
    }

    span {
      display: block;
      padding: 0 6px;
      color: rgba(0, 0, 0, 0.65);
      transition: none;
    }
  }

  .by-pagination-total-text {
    display: inline-block;
    height: 32px;
    margin-right: 8px;
    line-height: 30px;
    vertical-align: middle;
  }

  .by-pagination-item-active {
    font-weight: 500;
    background: #fff;
    border-color: #1890ff;

    span {
      color: #1890ff;
    }
  }

  .by-pagination-next, .by-pagination-prev {
    outline: 0;
  }

  .by-pagination-prev .by-pagination-item-link, .by-pagination-next .by-pagination-item-link {
    display: block;
    height: 100%;
    font-size: 12px;
    text-align: center;
    background-color: #fff;
    border: 1px solid #d9d9d9;
    border-radius: 2px;
    outline: none;
    transition: all 0.3s;

    &:hover, &:focus {
      color: #1890ff;
      border-color: #1890ff;
    }
  }

  .by-pagination-jump-next, .by-pagination-jump-prev, .by-pagination-next, .by-pagination-prev {
    display: inline-block;
    min-width: 32px;
    height: 32px;
    color: rgba(0, 0, 0, 0.65);
    font-family: Arial;
    line-height: 32px;
    text-align: center;
    vertical-align: middle;
    list-style: none;
    border-radius: 2px;
    cursor: pointer;
    transition: all 0.3s;
  }

  .by-pagination-jump-prev .by-pagination-item-container, .by-pagination-jump-next .by-pagination-item-container {
    position: relative;
    height: 32px;
    outline: 0;

    &:hover {
      .by-pagination-item-ellipsis {
        opacity: 0;
      }

      .by-pagination-item-link-icon {
        opacity: 1;
        color: #1890ff;
      }
    }

    .by-pagination-item-ellipsis {
      position: absolute;
      top: 0;
      right: 0;
      bottom: 0;
      left: 0;
      display: block;
      margin: auto;
      color: rgba(0, 0, 0, 0.25);
      letter-spacing: 2px;
      text-align: center;
      text-indent: 0.13em;
      opacity: 1;
      transition: all 0.2s;
    }

    .by-pagination-item-link-icon {
      position: absolute;
      top: 0;
      right: 0;
      bottom: 0;
      left: 0;
      display: block;
      margin: auto;
      color: rgba(0, 0, 0, 0.25);
      letter-spacing: 2px;
      text-align: center;
      text-indent: 0.13em;
      opacity: 0;
      transition: all 0.2s;
      font-size: 12px;
    }
  }

  .by-pagination-jump-next, .by-pagination-jump-prev, .by-pagination-prev {
    margin-right: 8px;
  }

  .by-pagination-options {
    display: inline-block;
    margin-left: 16px;
    vertical-align: middle;

    .by-pagination-options-size-changer {
      display: inline-block;

      select {
        display: inline-block;
        width: auto;
        margin-right: 8px;
        box-sizing: border-box;
        margin: 0;
        padding: 0;
        color: rgba(0, 0, 0, 0.65);
        font-size: 14px;
        font-variant: tabular-nums;
        line-height: 1.5715;
        list-style: none;
        font-feature-settings: 'tnum';
        position: relative;
        display: inline-block;
        cursor: pointer;
        position: relative;
        background-color: #fff;
        border: 1px solid #d9d9d9;
        border-radius: 2px;
        transition: all 0.3s cubic-bezier(0.25, 0.8, 0.5, 1);
        width: 100%;
        height: 32px;
        padding: 0 11px;
        text-align: center;

        &:hover, &:focus {
          border-color: #40a9ff;
          box-shadow: 0 0 0 2px rgba(24, 144, 255, 0.2);
          outline: none;
        }

        option {
          flex: auto;
          overflow: hidden;
          white-space: nowrap;
          text-overflow: ellipsis;
          display: block;
          min-height: 32px;
          padding: 5px 12px;
          font-weight: 400;
          font-size: 14px;
          line-height: 22px;
        }
      }
    }

    .by-pagination-options-quick-jumper {
      display: inline-block;
      height: 32px;
      line-height: 32px;
      vertical-align: top;

      input {
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
        width: 50px;
        margin: 0 8px;

        &:focus {
          outline: 0;
          box-shadow: 0 0 0 2px rgba(24, 144, 255, 0.2);
        }

        &:hover, &:focus {
          border-color: #40a9ff;
          border-right-width: 1px !important;
        }
      }
    }
  }
}

.by-pagination-disabled {
  cursor: not-allowed;

  .by-pagination-item span, span {
    color: rgba(0, 0, 0, 0.25);
    background: transparent;
    border: none;
    cursor: not-allowed;
  }

  .by-pagination-item-link {
    background: transparent;
    border-color: transparent;

    &:hover, &:focus {
      color: rgba(0, 0, 0, 0.25) !important;
      border-color: #d9d9d9 !important;
      cursor: not-allowed !important;
    }
  }

  .by-pagination-item-active span {
    color: #fff;
  }
}

@media (max-width: 992px) {
  .by-pagination {
    .by-pagination-item-after-jump-prev, .by-pagination-item-before-jump-next {
      display: none;
    }
  }
}
</style>  