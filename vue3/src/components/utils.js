const utils = {
    //限制
    limit(index, min, max, cycle = true) {
        // index -= 1
        index = parseInt(index)
        if (index < min) {
            index = cycle ? max : min
        } else if (index > max) {
            index = cycle ? min : max
        }
        return index
    },
    //// 时间 格式化成 2018-12-12 12:12:00
    timestampToTime(timestamp, dayMinSecFlag) {
        let date = new Date(timestamp);
        let Y = date.getFullYear() + "-";
        let M =
            (date.getMonth() + 1 < 10
                ? "0" + (date.getMonth() + 1)
                : date.getMonth() + 1) + "-";
        let D =
            date.getDate() < 10 ? "0" + date.getDate() + " " : date.getDate() + " ";
        let h =
            date.getHours() < 10 ? "0" + date.getHours() + ":" : date.getHours() + ":";
        let m =
            date.getMinutes() < 10
                ? "0" + date.getMinutes() + ":"
                : date.getMinutes() + ":";
        let s =
            date.getSeconds() < 10 ? "0" + date.getSeconds() : date.getSeconds();
        if (!dayMinSecFlag) {
            return Y + M + D;
        }
        return Y + M + D + h + m + s;
    },

    // 时间格式化位00：00
    formatTime(value) {
        let time = "";
        let s = value.split(":");
        let i = 0;
        for (; i < s.length - 1; i++) {
            time += s[i].length == 1 ? "0" + s[i] : s[i];
            time += ":";
        }
        time += s[i].length == 1 ? "0" + s[i] : s[i];
        return time;
    },
    // 把毫秒变成时分秒
    transTime(value) {
        let time = "";
        let h = parseInt(value / 3600);
        value %= 3600;
        let m = parseInt(value / 60);
        let s = parseInt(value % 60);
        if (h > 0) {
            time = this.formatTime(h + ":" + m + ":" + s);
        } else {
            time = this.formatTime(m + ":" + s);
        }
        return time;
    },

    isInContainer(el, container, offset = 0) {
            //元素位置
        const elTop = el.offsetTop
            //滚动位置
        const containerTop = this.getScroll(container, true)
        const isWindow = container === window
        const containerHeight = isWindow ? document.documentElement.clientHeight : container.clientHeight
        if ((elTop > containerTop) && (elTop < containerTop + containerHeight + offset)) {
            return true
        }
        return false

    },
    getScroll(target, top) {
        if (typeof window === 'undefined') {
            return 0;
        }

        const prop = top ? 'pageYOffset' : 'pageXOffset';
        const method = top ? 'scrollTop' : 'scrollLeft';
        const isWindow = target === window;
        let ret = isWindow ? target[prop] : target[method];
        // ie6,7,8 standard mode
        if (isWindow && typeof ret !== 'number') {
            ret = window.document.documentElement[method];
        }
        return ret;
    },
    easeInOutCubic(t, b, c, d) {
        const cc = c - b;
        t /= d / 2;
        if (t < 1) {
            return (cc / 2) * t * t * t + b;
        }
        return (cc / 2) * ((t -= 2) * t * t + 2) + b;
    },
    scrollTo(y, options) {
        const { getContainer = () => window, callback, duration = 450 } = options
        const container = getContainer();
        const scrollTop = this.getScroll(container, true);
        const startTime = Date.now();
        const frameFunc = () => {
            const timestamp = Date.now();
            const time = timestamp - startTime;
            const nextScrollTop = this.easeInOutCubic(time > duration ? duration : time, scrollTop, y, duration);
            if (container === window) {
                window.scrollTo(window.pageXOffset, nextScrollTop);
            } else {
                container.scrollTop = nextScrollTop;
            }
            if (time < duration) {
                requestAnimationFrame(frameFunc);
            } else if (typeof callback === 'function') {
                callback();
            }
        };
        requestAnimationFrame(frameFunc);

    },
    getElementPosition(e) {
        var x = 0,
            y = 0;
        while (e != null) {
            x += e.offsetLeft;
            y += e.offsetTop;
            e = e.offsetParent;
        }
        return { x: x, y: y };
    }
}
export default utils