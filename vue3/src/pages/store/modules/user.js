import http from '@/utils/httpindex.js'
import Notification from '@/components/notification/index.js';

const user = {
    state: {
        info: {},
        token: '',
    },
    mutations: {
        setToken(state, data) { //写入token
            state.token = data
        },
        userInfo(state, data) { //写入个人信息
            state.info = data;
            Notification.success({
                message: ` 欢迎回来~${data.name}`,
                description: ``,
                icon: < img src = { `${data.head_url}` }
                width = "60"
                style = "position: absolute;width: 40px;border-radius: 50%;border: 2px solid rgba(223,223,223,0.3);" / >
            })
        },
        logOut(state) { //退出
            console.log(state.info)
            Notification.success({
                message: `退出成功~${state.info.name}`,
                description: `欢迎下次登陆！`,
                icon: < img src = { `${state.info.head_url}` }
                width = "60"
                style = "position: absolute;width: 40px;border-radius: 50%;border: 2px solid rgba(223,223,223,0.3);" / >
            })
            state.token = '';
            state.info = {}
        }
    },
    actions: {
        setToken({ commit }, data) {
            // localStorage.setItem("so_token", data);
            commit('setToken', data)
        },
        userInfo({ commit }, data) {
            commit('userInfo', data)
        },
        async logOut({ commit }) {
            commit('logOut')
            await http.post('/apis/user/logout');
        },
    }
}
export default user