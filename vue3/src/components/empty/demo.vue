<template>
<div>
    <Empty>
    </Empty>
     <Empty
    image="https://gw.alipayobjects.com/mdn/miniapp_social/afts/img/A*pevERLJC9v0AAAAAAAAAAABjAQAAAQ/original"
    :image-style="{
      height: '60px',
    }"
  >
    <template #description>
      <span> Customize <a href="#API">Description</a> </span>
    </template>
    <button type="primary">
      Create Now
    </button>
  </Empty>
</div>
</template>
<script>
import Empty from './empty'
export default {
    components:{
        Empty
    }
}
</script>