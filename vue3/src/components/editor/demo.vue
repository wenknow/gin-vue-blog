<template>
    <preview :content="content" :tocNav="true" @click-img="clickImg" :dompurify="false" :markedOptions="options"/>
    <input v-model="content" />
    <ImageView v-model:show="show" :urlList="list" :url="url"/>
</template>
<script>
import preview from './preview'
import ImageView from '../image/image-view'
export default {
    components:{
        preview,
        ImageView
    },
    data(){
        return {
            content:localStorage.getItem('markdown'),
            url:"",
            list:[],
            show:false,
            options:{
                headerIds:false,
                breaks:true
            }
        }
    },
    methods:{
        clickImg(src,list){
            this.src=src
            this.list=list
            console.log(list)
            this.show=true
        }
    }
}
</script>