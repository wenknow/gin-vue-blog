<template>
    <div class="authLogin container" >
        <div class="message-auth">
            <h4 class="title">授权登录加载中，请耐心等待...</h4>
        </div>
    </div>
</template>
<script>
    import {onMounted} from "vue";
    import { useRoute } from "vue-router";
    import { useStore } from "vuex";
    import http from "@/utils/httpindex.js";
    export default {
    name:'Auth',
    setup(){
        const store = useStore();
        const route = useRoute();
        const goAuthLogin = () => {
            let method = route.query.method;
            let code = route.query.code;
            if (method === 'github'){
                http.post('/apis/user/githubLogin',{code:code}).then(function(res) {
                    if (res){
                        store.dispatch("setToken", res.token);
                        store.dispatch("userInfo", res.user);
                        window.location.replace("/");
                    }
                });
            }
        };

        onMounted(() => {
            goAuthLogin();
        });
    }
}
</script>
<style lang="stylus" scoped>
    .authLogin {
        margin: 0 auto;
    }
    .message-auth {
        padding: 100px 50px;
        display: flex;
        justify-content: center;
        flex-direction: column;
        align-items: center;
        font-weight: 800;
        font-size: 24px;
        line-height: 110%;
        margin: 90px 0 28px;
    }
</style>