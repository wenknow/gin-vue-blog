import { createRouter, createWebHistory } from 'vue-router'
const name = "文知道开发者的博客平台"

const routes = [
    // {
    //     path: '/',
    //     name: 'Home',
    //     component: Index,
    //     meta: {
    //         title: '首页-' + name,
    //         keepAlive: true, // 需要被缓存
    //         requiresAuth: false
    //     }
    // },
    {
        path: '/',
        name: 'Blog',
        component: () =>
            import ('../views/blog/main.vue'),
        meta: {
            title: '博客-' + name,
            keepAlive: true, // 需要被缓存
            requiresAuth: false
        }
    },
    {
        path: '/center/:id',
        name: 'Center',
        component: () =>
            import ('../views/user/center.vue'),
        meta: {
            title: '博客-' + name,
            keepAlive: true, // 需要被缓存
            requiresAuth: false
        }
    },
    {
        path: '/blog/:id',
        name: 'BlogContent',
        component: () =>
            import ('../views/blog/blog-content.vue'),
        meta: {
            title: 'loading……',
            keepAlive: true // 需要被缓存
        }
    },
    {
        path: '/blog/edit/:id',
        name: 'BlogEdit',
        component: () =>
            import ('../views/blog/blog-edit.vue'),
        meta: {
            title: 'loading……',
            keepAlive: true // 需要被缓存
        }
    },
    {
        path: '/message',
        name: 'Message',
        component: () =>
            import ('../views/message/main.vue'),
        meta: {
            // keepAlive: true, // 需要被缓存
            title: '留言圈-' + name,
        }
    },
    {
        path: '/login',
        name: 'Login',
        component: () =>
            import ('../views/login/login.vue'),
        meta: {
            title: '登录-' + name,
        }
    },
    {
        path: '/project',
        name: 'Project',
        component: () =>
            import ('../views/circle/project.vue'),
        meta: {
            title: '项目圈-' + name,
        }
    },
    {
        // 会匹配所有路径
        path: '/:pathMatch(.*)*',
        name: 'Windmill',
        component: () =>
            import ('../views/windmill.vue'),
        meta: {
            title: '哎呀页面不见了-' + name
        }
    },
]

const router = createRouter({
    history: createWebHistory(),
    scrollBehavior() {
        // return desired position
        return { top: 0 }
    },
    routes
})

export default router