import { ref } from 'vue'
import { useRoute } from "vue-router";
import http from '@/utils/httpindex.js'

function getUserArticleList(uid) {
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
        if (page_num.value * page_size.value < all_num.value) {
            page_num.value += 1
            loading.value = true
            http.post(`/apis/blog/userList`, {page_num:page_num.value, tag:tag, uid:uid})
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

function getUserTagList(uid) {
    const labelList = ref([])
    const labelLoading = ref(true)
    const getTag = () => {
        http.post('/apis/blog/userList', {uid:uid})
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

export {
    getUserArticleList,
    getUserTagList,
}