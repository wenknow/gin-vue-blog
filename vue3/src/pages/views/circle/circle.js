import { ref } from 'vue'
import http from '@/utils/httpindex.js'
// import message from "@/components/message/index.js";

function getProjectList() {
    const list = ref([])
    const loading = ref(true)
    http.get('/apis/link/list?type=1')
        .then((res) => {
            list.value = res
            loading.value = false
        })
    return {
        list,
        loading
    }
}


export {
    getProjectList,
}