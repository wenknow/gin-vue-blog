import { ref } from 'vue'
import { useRoute } from "vue-router";
import http from '@/utils/httpindex.js'

function getArticleList() {
    const page_num = ref(0)
    const page_size = ref(10)
    const all_num = ref(1)
    const list = ref([])
    const loading = ref(false)
    const route = useRoute()
    const getBlog = () => {
        let tag = '';
        if (route.query.tag) {
            tag = decodeURI(encodeURI(route.query.tag).replace(/%20/g,'+'));
        }
        let search = '';
        if (route.query.search) {
            search = decodeURI(encodeURI(route.query.search).replace(/%20/g,'+'));
        }
        if (page_num.value * page_size.value < all_num.value) {
            page_num.value += 1
            loading.value = true
            http.post(`/apis/blog/list`, {page_num:page_num.value, tag:tag, keyword:search})
                .then(function(res) {
                    if (page_num.value === 1) {
                        list.value = res.list
                    } else {
                        list.value = list.value.concat(res.list)
                    }
                    loading.value = false
                    page_num.value = res.page_num
                    page_size.value = res.page_size
                    all_num.value = res.all_num
                })
        }
    }
    return {
        page_num,
        page_size,
        all_num,
        list,
        loading,
        getBlog
    }

}

function getMainTagList() {
    const labelList = ref([])
    const labelLoading = ref(true)
    const route = useRoute()
    const getTag = () => {
        let search = '';
        if (route.query.search) {
            search = decodeURI(encodeURI(route.query.search).replace(/%20/g, '+'));
        }
        http.post('/apis/blog/tag/main', {keyword: search})
            .then(function (res) {
                labelLoading.value = false
                labelList.value = res
            })
        return {
            labelList,
            labelLoading
        }
    }
    return {
        labelList,
        labelLoading,
        getTag
    }

}

function getBlogInfo() {
    const count = ref(0)
    http.get("/apis/blog/info")
        .then(function(res) {
            count.value = res;
        })
    return {
        count
    }
}

function getTagList() {
    const tagList = ref([])
    const tagLoading = ref(true)
    http.post('/apis/blog/tag/list')
        .then(function(res) {
            tagLoading.value = false
            tagList.value = res
        })
    return tagList
}

function getCgList(uid) {
    const cgList = ref([])
    http.post('/apis/blog/cg/list', {uid:uid})
        .then(function(res) {
            cgList.value = res
        });
    return cgList
}

function  guessTag(params) {
    const tagArr = ref([])
    const loading = ref(true)
    http.post('/apis/blog/tag/guess',params)
        .then(function(res) {
            loading.value = false
            tagArr.value = res
        })
    return tagArr
}


export {
    getArticleList,
    getMainTagList,
    getBlogInfo,
    getTagList,
    getCgList,
    guessTag,
}