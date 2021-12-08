// 代码压缩
const IS_PROD = process.env.NODE_ENV === 'production';
module.exports = {
    // 由于部分插件，导致ie下空白
    transpileDependencies: ['vue-savedata', 'vue-baberrage'],
    publicPath: '/',
    // 输出目录，
    outputDir: 'dist',
    //多页面配置
    pages: {
        index: {
            entry: 'src/pages/main.js',
            template: `public/index.html`,
            filename: 'index.html',
            chunks: ["chunk-vendors", "chunk-common", "index", 'runtime~index'],
            title: '文知道-开发者的博客平台'
        }
    },
    devServer: {
        sockHost: "localhost",
        disableHostCheck: true,
        port: 8080, // 端口号
        host: "0.0.0.0",
        https: false, // https:{type:Boolean}
        open: true, //配置自动启动浏览器
        proxy: {
            "/apis": {
                target: process.env.VUE_APP_API_URL, // 需要请求的地址
                changeOrigin: true, // 是否跨域
                pathRewrite: {
                    "^/apis": ""
                }
            },
        }
    },
    productionSourceMap: !IS_PROD,
    configureWebpack: config => {
        // 用cdn方式引入，则构建时要忽略相关资源
        config.externals = {
            "echarts": "echarts",
        };

        if (IS_PROD) {
            // 代码压缩
            config.optimization.minimizer[0].options.terserOptions.compress.warnings = false
            config.optimization.minimizer[0].options.terserOptions.compress.drop_console = true
            config.optimization.minimizer[0].options.terserOptions.compress.drop_debugger = true
            config.optimization.minimizer[0].options.terserOptions.compress.pure_funcs = ['console.log']
        }
    },
    chainWebpack: config => {
        if (IS_PROD) {
            config.plugins.delete('prefetch')
                // 压缩代码
            config.optimization.minimize(true);
        }
    }
}