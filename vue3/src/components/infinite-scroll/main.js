import { throttle } from "throttle-debounce";
import utils from '../utils'
import { nextTick, } from 'vue'
const getScrollOptions = (el) => {
    const option = {
        delay: 200,
        distance: 0,
        disabled: false,
        immediate: true,
    }
    for (let name in option) {
        const val = el.getAttribute(`infinite-scroll-${name}`)
        if (val) {
            if (name == 'disabled' || name == 'immediate') {
                option[name] = val == 'true' ? true : false
            } else {
                option[name] = parseInt(val)
            }

        }
    }
    return option

}
const handleScroll = (el, cb) => {
    const { distance, disabled } = getScrollOptions(el)
    if (disabled) return
        // const isWindow = target === window;
        // const top=isWindow?'scrollTop'
    const isWindow = el.getAttribute('infinite-scroll-window')
    if (isWindow) {
        el = window
    }
    const scrollTop = utils.getScroll(el, true)
    el = el === window ? document.documentElement : el
    const elHeight = el.scrollHeight
    const clientHeight = el.clientHeight
    if (elHeight - scrollTop - clientHeight <= distance) {
        cb()
    }

}
const InfiniteScroll = {
    async mounted(el, binding) {
        await nextTick()
        const cb = binding.value
        let container = el
        const { delay } = getScrollOptions(el)
        const isWindow = el.getAttribute('infinite-scroll-window')
        if (isWindow) {
            container = window
        }
        const onScroll = throttle(delay, handleScroll.bind(null, el, cb))
        el['SCOPE'] = {
            onScroll,
            container
        }
        container.addEventListener("scroll", onScroll);

    },
    unmounted(el) {
        const { container, onScroll } = el['SCOPE']
        container.removeEventListener("scroll", onScroll);
    },
}
export default InfiniteScroll