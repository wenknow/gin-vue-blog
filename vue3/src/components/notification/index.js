import { createApp } from 'vue'
import Message from './index.vue'
let key = 1
const createMessage = {
    options: {
        top: `24px`,
        bottom: '24px',
        duration: 3,
        maxCount: null,
        getContainer: document.body,
        placement: 'topRight',
        closeIcon: null,
    },
    default: {
        btn: null,
        bottom: null,
        class: null,
        description: null,
        duration: null,
        icon: null,
        key: null,
        message: null,
        placement: null,
        style: null,
        onClose: null,
        onClick: null,
        top: null,
        closeIcon: null,
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
        if (!this.vm) {
            this.create()
        }
        this.vm.add(options)
        return options.key
            // if (options.type != 'close' && options.duration == 0) {
            //     return this.close
            // }
    },
    close(key) {
        this.vm.close(key)
    },
    format(type, config) {
        const op = Object.assign({}, this.default, this.options, { type: type, getContainer: '', key: 'baymax-notification-' + key++ }, config, )
        return op
    },
    open(config) {
        const type = "info"
        return this.show(this.format(type, config))
    },
    success(config) {
        const type = "success"
        return this.show(this.format(type, config))
    },
    error(config) {
        const type = "error"
        return this.show(this.format(type, config, ))
    },
    info(config) {
        const type = "info"
        return this.show(this.format(type, config))
    },
    warning(config) {
        const type = "warning"
        return this.show(this.format(type, config))
    },
    loading(config) {
        const type = "loading"
        return this.show(this.format(type, config))
    },
}
export default createMessage