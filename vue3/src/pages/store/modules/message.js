import http from '@/utils/httpindex.js'
const user = {
    state: {
        list: {},
    },
    mutations: {
        setMessList(state, data) { //写入token
            state.list = data
        },
        setPushMess(state, data) { //写入token
            state.list.data = state.list.data.concat(data.data)
            state.list.current_page = data.current_page
            state.list.last_page = data.last_page
        },
    },
    actions: {
        getMessList({ commit, state }, data) {
            // await http.get('/apis/user/logout')
            if (data.page == 1) {
                const res = http.get(`/apis/message/list?id=${data.id}&page=${data.page}`)
                commit('setMessList', res.data)
            } else if (state.list.current_page < state.list.last_page) {
                console.log(state.list.current_page, state.list.last_page)
                const res = http.get(`/apis/message/list?id=${data.id}&page=${data.page}`)
                commit('setPushMess', res.data)
            }
        },
    }
}
export default user