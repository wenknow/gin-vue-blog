import { createStore } from 'vuex'
import index from './modules/index'
import user from './modules/user'
import message from './modules/message'
import createPersistedState from "vuex-persistedstate";
const dataState = createPersistedState({
    paths: ['user']
})
export default createStore({
    state: {},
    mutations: {},
    actions: {},
    modules: {
        index,
        user,
        message
    },
    plugins: [dataState]
})