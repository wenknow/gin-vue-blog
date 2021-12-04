import http from '@/utils/httpindex.js'
const formatDate = () => {
    var d = new Date(),
        month = "" + (d.getMonth() + 1),
        day = "" + d.getDate(),
        year = d.getFullYear();

    if (month.length < 2) month = "0" + month;
    if (day.length < 2) day = "0" + day;

    return [year, month, day].join("-");
}
const user = {
    state: {
        list: [],
        musicList: [],
        friendList: [],
        dailyEnglish: {},
        loginShow: false
    },
    mutations: {
        setCarousels(state, data) { //写入token
            state.list = data
        },
        setMusicList(state, data) { //写入token
            state.musicList = data
        },
        setFriendList(state, data) { //写入token
            state.friendList = data
            console.log(state.friendList)
        },
        setDailyEnglish(state, data) { //写入token
            state.dailyEnglish = {
                name: data.content,
                lrc: data.note,
                url: data.tts,
                cover: data.picture,
                bg: data.picture2
            }
        },
        showLogin(state, show = true) {
            state.loginShow = show
        }
    },
    actions: {
        async getCarousels({ commit }) {
            // await http.get('/apis/user/logout')
            const res = await http.get('/apis/show/list')
            commit('setCarousels', res.data)
        },
        async getMusicList({ commit }) {
            // await http.get('/apis/user/logout')
            const res = await http.get('/apis/music/list')
            commit('setMusicList', res.data)
        },
        async getFriendList({ commit }) {
            // await http.get('/apis/user/logout')
            const res = await http.get('/apis/link/list?type=0')
            commit('setFriendList', res.data)
        },
        async getDailyEnglish({ commit }) {
            // await http.get('/apis/user/logout')
            const res = await http.get('/english/index.php?c=dailysentence&m=getdetail&title=' + formatDate())
            commit('setDailyEnglish', res)
        },
    }
}
export default user