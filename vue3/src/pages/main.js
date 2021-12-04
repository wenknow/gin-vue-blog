import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import http from '@/utils/httpindex'
import '@/assets/iconfont/iconfont.js'
import InfiniteScroll from '@/components/infinite-scroll/index.js'
import '@/components/editor/show.js'
import ElementPlus from "element-plus";
import 'element-plus/lib/theme-chalk/index.css';
import VMdEditor from '@kangc/v-md-editor';
import '@kangc/v-md-editor/lib/style/base-editor.css';
import githubTheme from '@kangc/v-md-editor/lib/theme/github.js';
import '@kangc/v-md-editor/lib/theme/style/github.css';
import hljs from '@/components/editor/highlight.min.js';

VMdEditor.use(githubTheme, {
    Hljs: hljs,
});

router.beforeEach((to, from, next) => {
    if (to.meta.title) {
        document.title = to.meta.title
    }
    next()
})

const app = createApp(App);
app.use(ElementPlus);
app.use(store);
app.use(router);
app.use(VMdEditor);
app.use(InfiniteScroll);
app.config.globalProperties.$get = http.get;
app.config.globalProperties.$post = http.post;
http.defaults.baseURL = process.env.VUE_APP_URL;
app.mount('#app');