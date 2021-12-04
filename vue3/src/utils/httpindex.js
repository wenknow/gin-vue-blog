import axios from 'axios'
import store from '@/pages/store'
import Notification from '@/components/notification/index.js';
var instance = axios.create({
    baseURL: process.env.VUE_APP_API_URL,
    timeout: 10000,
});
// 添加请求拦截器
instance.interceptors.request.use(function(config) {
    // 在发送请求之前做些什么
    config.headers['X-Token'] = store.state.user.token;
    return config;
}, function(error) {
    // 对请求错误做些什么
    return Promise.reject(error);
});

// 添加响应拦截器
instance.interceptors.response.use(function(response) {
        if (response.data.code === 200) {
            return Promise.resolve(response.data.data)
        } else if(response.data.code === 600){
            store.dispatch('logOut');
            Notification.error({
                message: '登录信息错误',
                description: response.data.msg,
            })
        } else {
            Notification.error({
                message: '错误提示',
                description: response.data.msg,
            })
        }
        // 打印错误信息
        return Promise.reject(response.data)
    },
    function(error) {
        // 对响应错误做点什么
        console.log(error.response)
        switch (error.response.status) {
            case 400:
                Notification.warning({
                    message: '错误提示',
                    description: error.response.data.message,
                    onClose() {
                        store.commit('showLogin')
                    }
                });
                break;
            case 401:
                Notification.warning({
                    message: '请求参数有误',
                    description: error.response.data.message,
                });
                break;
            case 403:
                Notification.warning({
                    message: '用户权限提示',
                    description: error.response.data.message,
                    onClose() {
                        store.commit('showLogin')
                    },
                });
                break;
            case 404:
                Notification.warning({
                    message: '访问路径不正确',
                    description: error.response.data.message,
                });
                break;
            case 422:
                Notification.warning({
                    message: '温馨提示',
                    description: error.response.data.message,
                    iconClass: 'el-icon-warning-outline',
                    onClose() {
                        store.dispatch("logOut")
                        store.commit('showLogin')
                    },
                });
                break;
            case 429:
                Notification.warning({
                    message: '温馨提示',
                    description: error.response.data.message,
                });
                break;
            case 500:
                Notification.error({
                    message: '网络提示',
                    description: '服务器连接失败，请稍后再试',
                });
                break;
            default:
                Notification.error({
                    message: '错误提示 ' + error.response.status,
                    description: error.response.data.message,
                });

        }
        return Promise.reject(error)
    });
export default instance;