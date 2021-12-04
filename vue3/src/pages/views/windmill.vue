<template>
    <router-link  to="/blog">
      <div class="canvas a-tag blog-animation">
        <img src="https://iconfont.alicdn.com/s/e8a7a96a-7f80-4f07-bc88-c162da451fe1_origin.svg" width="400">
      </div>
    </router-link>
</template>
<script>
export default {
    name:'Windmill'
}
</script>
<style lang="stylus" scoped>
.canvas
  height 100vh
  display flex
  flex-direction column
  justify-content center
  align-items center
  img 
    width 100%
    max-width 400px
  p
    margin-top 20px
  
</style>
