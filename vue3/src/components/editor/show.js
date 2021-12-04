import marked from 'marked'
import hljs from './highlight.min.js'
const TAG_NAME = 'demo-mobai'
    // 此处是可能产生例外的语句
customElements.define(TAG_NAME, class DemoSandbox extends HTMLElement {
    constructor() {
        super()
            // 使用影子DOM
        this.shadow = this.attachShadow({
                mode: 'open'
            })
            // 获取关联的代码块模板的ID
        const templateId = this.getAttribute('template')
        const $template = document.getElementById(templateId)
        if (!templateId) {
            return
        }
        // 获取代码块内容
        const template = $template.innerHTML
        console.log(templateId)
        let id = parseInt(templateId.split('demo-mobai-template-')[1]);
        console.log(id % 2 == 0)
        if (id % 2 == 0) {
            // 用获取到的代码块来填充影子DOM的HTML
            let code = marked('``` html  \n' + template + '\n```', {
                sanitize: false,
                highlight: function(code) {
                    return hljs.highlightAuto(code).value;
                },
            })
            this.shadow.innerHTML = `
            <style>
                :host {
                    display:block;
                    width:100%;
                    padding: 0;
                    border: 1px solid #f0f0f0;
                    color: #414240;
                    font-size: 1rem;
                    position: relative;
                    margin: 10px 0;
                    min-height: 36px;
                }
                :host:before {
                    content: " ";
                    position: absolute;
                    -webkit-border-radis: 50%;
                    border-radius: 50%;
                    background: #ff6058;
                    width: 12px;
                    height: 12px;
                    left: 15px;
                    margin-top: 10px;
                    -webkit-box-shadow: 20px 0 #ffbd2b, 40px 0 #3cef57;
                    box-shadow: 20px 0 #ffbd2b, 40px 0 #3cef57;
                    z-index: 2;
                }
                :host:after {
                    content: "demo";
                    position: absolute;
                    top:0px;
                    left: 50%;
                    z-index: 2;
                    color:#3da8f5;
                    font-weight:bold;
                    transform: translateX(-50%);
                    font-size: 20px;
                    line-height:32px
                }
                * {
                    box-sizing: border-box;
                }
                
                #demo-run {
                    padding:20px;
                    background-color:white;
                    border-top: 32px solid #ecf5ff;
                    border-radius: 6px;
                    overflow-x: auto;
                    overflow-y: hidden;
                    position:relative;
                }
                #demo-code {
                    padding:20px;
                    border-top: 1px solid #eaeefb;
                    font-size: 85%;
                    font-family: "Operator Mono SSm A","Operator Mono SSm B","Operator Mono","Source Code Pro",Menlo,Consolas,Monaco,monospace;
                    line-height: 1.4;
                    background-color:#fefefe; 
                }
                #demo-code code{
                    display: block;
                    overflow-x: auto;
                }
                #demo-open {
                    width:100%;
                    -webkit-appearance: none;
                    border:none;
                    border-top: 1px solid #eaeefb;
                    text-align:center;
                    padding: 10px 20px;
                    font-size: 14px;
                    cursor: pointer;
                    outline: 0;
                    transition: background-color .3s;
                    color: #3da8f5;
                    background-color:#fff
                }
                #demo-open:hover,
                #demo-open:active {
                    background-color: #f5f7fa;
                }
            </style>
            <div id="demo-run">${template}</div>
            <div id="demo-code" hidden>${code}</div>
            <button id="demo-open">查看源码</button>
            <style>
            .hljs{display:block;overflow-x:auto}.hljs-comment,.hljs-meta{color:#969896}.hljs-emphasis,.hljs-quote,.hljs-string,.hljs-strong,.hljs-template-variable,.hljs-variable{color:#df5000}.hljs-keyword,.hljs-selector-tag,.hljs-type{color:#a71d5d}.hljs-attribute,.hljs-bullet,.hljs-literal,.hljs-number,.hljs-symbol{color:#0086b3}.hljs-name,.hljs-section{color:#63a35c}.hljs-tag{color:#333}.hljs-attr,.hljs-selector-attr,.hljs-selector-class,.hljs-selector-id,.hljs-selector-pseudo,.hljs-title{color:#795da3}.hljs-addition{color:#55a532;background-color:#eaffea}.hljs-deletion{color:#bd2c00;background-color:#ffecec}.hljs-link{text-decoration:underline}.hljs-comment,.hljs-quote{color:#998}.hljs-keyword,.hljs-selector-tag,.hljs-subst{font-weight:700}.hljs-literal,.hljs-number,.hljs-tag .hljs-attr,.hljs-template-variable,.hljs-variable{color:teal}.hljs-doctag,.hljs-string{color:#d14}.hljs-section,.hljs-selector-id,.hljs-title{color:#900;font-weight:700}.hljs-subst{font-weight:400}.hljs-class .hljs-title,.hljs-type{color:#458;font-weight:700}.hljs-link,.hljs-regexp{color:#009926}.hljs-bullet,.hljs-symbol{color:#990073}.hljs-built_in,.hljs-builtin-name{color:#0086b3}.hljs-deletion{background:#fdd}.hljs-addition{background:#dfd}.hljs-emphasis{font-style:italic}
            </style>
            `
            const co = this.shadow.getElementById("demo-code")
            this.shadow.getElementById("demo-open").addEventListener(
                "click", (function() {
                    co.hasAttribute("hidden") ? co.removeAttribute("hidden") : co.setAttribute("hidden", "")
                }));
        } else {
            this.shadow.innerHTML = `
            <div id="demo-run">${template}</div>
            `
        }


        // 移除掉关联的template节点
        // 移除掉关联的template节点
        $template.parentNode.removeChild($template)
            // 处理 script
            // 1. 查找影子DOM中刚才填充的script节点
        const scripts = Array.from(this.shadow.querySelectorAll('script'))
        console.log(scripts)
            // 2. 创建一个用来保存影子DOM根节点的Script
        const $globalDefines = document.createElement('script')
            // 3. 创建一个自执行函数，将代码包裹起来
        $globalDefines.innerHTML = `(function(){
        const $component = document.querySelector('${TAG_NAME}[template="${templateId}"]');
        const $shadowDocument = $component.shadowRoot;
        `
        let arr = []
            // 4. 拼合所有Script
        for (let i = 0; i < scripts.length; i++) {
            // 全局替换document为新的$shadowDocument
            if (scripts[i].src) {
                // 创建
                const $sc = document.createElement('script')
                $sc.setAttribute("type", "text/javascript");
                $sc.setAttribute('src', scripts[i].src);
                this.shadow.appendChild($sc)
                arr.push(
                    //通过Promise来解决，所有js都加载成功后，在将代码添加到Shadow DOM
                    new Promise(function(resolve) {
                        //js 加载完成回执行
                        $sc.onload = function() {
                            console.log($sc)
                            resolve();
                        };
                    })
                )
                this.shadow.getElementById('demo-run').removeChild(scripts[i])
                continue
            }

            $globalDefines.innerHTML += `{
                ${scripts[i].textContent.replace(/(document)\.(getElementById|querySelector|querySelectorAll|getElementsByClassName|getElementsByName|getElementsByTagName)/gm, '$shadowDocument.$2').replace(/\r\n?/gm, '')}
            }`
                // 移除旧节点
            this.shadow.getElementById('demo-run').removeChild(scripts[i])
        }
        $globalDefines.innerHTML += `})();`

        Promise.all(arr).then(() => {
            console.log('js加载成功');
            this.shadow.appendChild($globalDefines)
        })
    }
})