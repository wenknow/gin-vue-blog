<template>
  <div class="ant-comment">
    <div class="ant-comment-inner">
      <div class="ant-comment-avatar">
        <slot name="avatar">
          {{ avatar }}
        </slot>
      </div>
      <div class="ant-comment-content">
        <div class="ant-comment-content-author">
          <span class="ant-comment-content-author-name">
            <slot name="author">
              {{ author }}
            </slot>
          </span>
          <span class="ant-comment-content-author-time">
            <slot name="datetime">
              {{ datetime }}
            </slot>
          </span>
        </div>
        <div class="ant-comment-content-detail">
          <slot name="content">
            {{ content }}
          </slot>
        </div>
        <ul class="ant-comment-actions" v-if="$slots['actions']">
          <li
            v-for="(item, index) of $slots['actions']()"
            :key="'comment-action-' + index"
          >
            <component :is="item"></component>
          </li>
        </ul>
      </div>
    </div>
    <div class="ant-comment-nested">
      <slot></slot>
    </div>
    <!---->
  </div>
</template>
<script>
export default {
  name: Comment,
  props: {
    actions: Array,
    author: String,
    avatar: String,
    content: String,
    prefixCls: String,
    datetime: String,
  },
};
</script>
<style lang="stylus" scoped>
.ant-comment {
  position: relative;
  text-align: left;
}

.ant-comment-inner {
  display: flex;
  padding-bottom: 16px;
}

.ant-comment-avatar {
  position: relative;
  flex-shrink: 0;
  margin-right: 12px;
  cursor: pointer;
}

.ant-comment-avatar img {
  width: 32px;
  height: 32px;
  border-radius: 50%;
}

.ant-comment-content {
  position: relative;
  flex: 1 1 auto;
  min-width: 1px;
  font-size: 14px;
  word-wrap: break-word;
}

.ant-comment-content-author {
  display: flex;
  flex-wrap: wrap;
  justify-content: flex-start;
  margin-bottom: 4px;
  font-size: 14px;
}

.ant-comment-content-author > a, .ant-comment-content-author > span {
  padding-right: 8px;
  font-size: 12px;
  line-height: 18px;
}

.ant-comment-content-author-name {
  color: rgba(0, 0, 0, 0.45);
  font-size: 14px;
  transition: color 0.3s;
}

.ant-comment-content-author-name > * {
  color: rgba(0, 0, 0, 0.45);
}

.ant-comment-content-author-name > *:hover {
  color: rgba(0, 0, 0, 0.45);
}

.ant-comment-content-author-time {
  color: #ccc;
  white-space: nowrap;
  cursor: auto;
}

.ant-comment-content-detail {
  p {
    white-space: pre-wrap;
  }
}

.ant-comment-actions {
  margin-top: 12px;
  padding-left: 0;
}

.ant-comment-actions > li {
  color: rgba(0, 0, 0, 0.45);
}

.ant-comment-actions > li > span {
  padding-right: 10px;
  color: rgba(0, 0, 0, 0.45);
  font-size: 12px;
  cursor: pointer;
  transition: color 0.3s;
  user-select: none;
}

.ant-comment-actions > li > span:hover {
  color: #595959;
}

.ant-comment-nested {
  margin-left: 44px;
}
</style>