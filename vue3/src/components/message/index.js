import { createApp } from 'vue'
import Message from './index.vue'
let key = 1
const createMessage = {
    options: {
        top: `24px`,
        duration: 3,
        maxCount: null,
        getContainer: document.body
    },
    vm: null,
    config(opts) {
        this.options = Object.assign({}, this.options, opts)
        return this
    },

    create() {
        const app = createApp(Message, {
            defaults: this.options
        })
        const mountNode = document.createElement('div')
        this.options.getContainer.appendChild(mountNode)
        this.vm = app.mount(mountNode)
            // return new Noty(params)
    },

    show(options = {}) {
        console.log(this.mess)
        if (!this.vm) {
            this.create()
        }
        this.vm.add(options)
        return options.key
    },
    close(key) {
        this.vm.close(key)
    },
    format(type, config, duration, onClose) {
        console.log(typeof config)
        if (typeof config != 'object') {
            config = {
                type: type,
                content: config,
                duration: duration || this.options.duration,
                onClose: onClose || null,
                key: `by-message-${key++}`
            }
        } else {
            config = {
                type: type,
                content: config.content,
                duration: config.duration || this.options.duration,
                onClose: config.onClose || null,
                icon: config.icon || null,
                key: config.key || `by-message-${key++}`
            }
        }
        return config

    },
    open(config, duration, onClose) {
        console.log()
        const type = "info"
        return this.show(this.format(type, config, duration, onClose))
    },
    success(config, duration, onClose) {
        const type = "success"
        return this.show(this.format(type, config, duration, onClose))
    },
    error(config, duration, onClose) {
        const type = "error"
        return this.show(this.format(type, config, duration, onClose))
    },
    info(config, duration, onClose) {
        const type = "info"
        return this.show(this.format(type, config, duration, onClose))
    },
    warning(config, duration, onClose) {
        const type = "warning"
        return this.show(this.format(type, config, duration, onClose))
    },
    loading(config, duration, onClose) {
        const type = "loading"
        return this.show(this.format(type, config, duration, onClose))
    },
}
export default createMessage