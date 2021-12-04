import { ref } from 'vue'
import http from '@/utils/httpindex.js'
import message from "@/components/message/index.js";

function getMessList(id) {
    const page_num = ref(0)
    const page_size = ref(10)
    const all_num = ref(1)
    const commentList = ref([])
    const loading = ref(false)
    const getMess = async(my) => {
        if (my) {
            page_num.value = 0
            page_size.value = 10
            commentList.value = []
        }
        if (page_num.value * page_size.value < all_num.value) {
            page_num.value += 1
            loading.value = true
            await http.post(`/apis/blog/comment/list`, {page_num:page_num.value, article_id:id})
                .then((res) => {
                    loading.value = false
                    if (res.list) {
                        commentList.value = commentList.value.concat(res.list)
                    }
                    page_num.value = res.page_num
                    page_size.value = res.page_size
                    all_num.value = res.all_num
                })
        }

    }

    function replyAdd(data) {
        data['ip'] = window._DEFAULT_CITY.ip
        data['address'] = window._DEFAULT_CITY.nation + '-' + window._DEFAULT_CITY.province + '-' + window._DEFAULT_CITY.city
        http.post("/apis/blog/comment/reply", data)
            .then(function(res) {
                message.success(`回复成功！`);
                for (let i = 0, len = commentList.value.length; i < len; i++) {
                    if (commentList.value[i].comment.id == res.to_sup_comment_id) {
                        // bug:当列表回复为空时！不能添加的回复列表
                        if (!commentList.value[i]['reply_list']) {
                            commentList.value[i]['reply_list'] = []
                        }
                        commentList.value[i].reply_list.unshift(res)
                        break
                    }
                }
            }).catch(() => {
                message.error(`回复失败！`);
            })
    }

    function messageAdd(data, isMess = true) {
        const mess = isMess ? '留言' : '评论';
        data['ip'] = window._DEFAULT_CITY.ip;
        data['address'] = window._DEFAULT_CITY.nation + '-' + window._DEFAULT_CITY.province + '-' + window._DEFAULT_CITY.city
        http.post("/apis/blog/comment/add", data)
            .then(function(res) {
                message.success(`${mess}成功！`);
                if (commentList.value.length % 20 == 0) {
                    commentList.value.pop()
                }
                commentList.value.unshift(res)
            })
            .catch(() => {
                message.error(`${mess}失败！`);
            })
    }

    function deleteMess(info) {
        http.post("/apis/blog/comment/del", {
                id: info.id,
            })
            .then(function() {
                message.success(`删除成功！`);
                refreshList();
            })
            .catch(function() {
                message.success(`删除失败！`);
            });
    }
    
    function refreshList() {
        loading.value = true
        http.post(`/apis/blog/comment/list`, {page_num:page_num.value, article_id:id})
            .then((res) => {
                loading.value = false
                commentList.value = res.list
                page_num.value = res.page_num
                page_size.value = res.page_size
                all_num.value = res.all_num
            })
    }
    return {
        page_num,
        page_size,
        all_num,
        commentList,
        loading,
        getMess,
        messageAdd,
        replyAdd,
        deleteMess,
    }
}


export {
    getMessList,
}