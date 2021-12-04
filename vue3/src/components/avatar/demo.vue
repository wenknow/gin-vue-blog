<template>
  <div>
    <div>
      <ByAvatar :size="64">
        <template #icon><UserOutlined /></template>
      </ByAvatar>
      <ByAvatar size="large">
        <template #icon><UserOutlined /></template>
      </ByAvatar>
      <ByAvatar>
        <template #icon><UserOutlined /></template>
      </ByAvatar>
      <ByAvatar size="small">
        <template #icon><UserOutlined /></template>
      </ByAvatar>
    </div>
    <br />
    <div>
      <ByAvatar shape="square" :size="64">
        <template #icon><UserOutlined /></template>
      </ByAvatar>
      <ByAvatar shape="square" size="large">
        <template #icon><UserOutlined /></template>
      </ByAvatar>
      <ByAvatar shape="square">
        <template #icon><UserOutlined /></template>
      </ByAvatar>
      <ByAvatar shape="square" size="small">
        <template #icon><UserOutlined /></template>
      </ByAvatar>
    </div>
     <div>
    <ByAvatar>
      <template #icon>
        <UserOutlined />
      </template>
    </ByAvatar>
    <ByAvatar>U</ByAvatar>
    <ByAvatar>USER</ByAvatar>
    <ByAvatar src="https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png" />
    <ByAvatar style="color: #f56a00; backgroundColor: #fde3cf">U</ByAvatar>
    <ByAvatar style="backgroundColor:#87d068">
      <template #icon>
        <UserOutlined />
      </template>
    </ByAvatar>
  </div>
   <div>
    <ByAvatar
      size="large"
      :style="{ backgroundColor: color, verticalAlign: 'middle' }"
    >
      {{ avatarValue }}
    </ByAvatar>
    <button
      :style="{ marginLeft: 16, verticalAlign: 'middle' }"
      @click="changeValue"
    >
      改变
    </button>
  </div>
  </div>
</template>
<script>
import { UserOutlined } from '@ant-design/icons-vue';
import ByAvatar from './avatar'
const UserList = ['U', 'Lucy', 'Tom', 'Edward'];
const colorList = ['#f56a00', '#7265e6', '#ffbf00', '#00a2ae'];

export default {
  components: {
    UserOutlined,
    ByAvatar
  },
  data() {
    return {
      avatarValue: UserList[0],
      color: colorList[0],
    };
  },
  methods: {
    changeValue() {
      const index = UserList.indexOf(this.avatarValue);
      this.avatarValue = index < UserList.length - 1 ? UserList[index + 1] : UserList[0];
      this.color = index < colorList.length - 1 ? colorList[index + 1] : colorList[0];
    },
  },
};
</script>