<template>
  <transition name="fade">
    <div class="sys-modal" v-if="show">
      <div class="sys-mask"></div>
      <div class="sys-mess-wrap">
        <div class="sys-mess concise">
          <div class="header">
            <Image :src="info.logo" class="logo" v-if="info.logo"> </Image>
            <h3 class="title" v-if="info.title">{{ info.title }}</h3>
          </div>
          <div class="body">
            <Preview
              :content="info.content"
              :tocNav="false"
              :dompurify="false"
              :markedOptions="options"
            />
          </div>
          <div class="footer">
            <button class="btn" @click="show = false">取消</button>
            <button class="btn know" @click="close">知道了</button>
          </div>
        </div>
      </div>
    </div>
  </transition>
</template>
<script>
import { ref, h, onMounted } from "vue";
import http from "@/utils/httpindex.js";
import Notification from "@/components/notification/index.js";
import Preview from "@/components/editor/preview";
import Image from "@/components/image/image.vue";
export default {
  name: "SystemMessage",
  components: {
    Preview,
    Image,
  },
  setup() {
    const info = ref([]);
    const show = ref(false);
    const options = {
      headerIds: false,
      breaks: true,
    };
    Notification.config({ top: "72px" });
    const getMessage = () => {
      http.get("/apis/sysmess/list").then((res) => {
        const type0 = res[0];
        const type1 = res[1];
        if (type0) {
          info.value = type0;
          const sys_info =JSON.parse(localStorage.getItem("sys_mess_info"));
          if (!sys_info || sys_info.id != type0.id) {
            show.value = true;
          } else if (sys_info.time != type0.updated_at) {
            const key = Notification.success({
              message: "公告更新提示",
              description: "点击查看",
              duration: 8,
              onClick: () => {
                show.value = true;
                Notification.close(key);
              },
              icon: h("img", {
                width: "60",
                style:
                  "position: absolute;width: 40px;border-radius: 50%;border: 2px solid rgba(223,223,223,0.3);",
                src:
                  "https://iconfont.alicdn.com/s/6ef764db-12ed-45a9-98cf-a88900c7c87c_origin.svg",
              }),
              // icon: h() <img src ={logo}
            });
          }
        }
        if (type1) {
          const logo = type1.logo
            ? type1.logo
            : "https://iconfont.alicdn.com/s/3e4fd9be-c989-4522-be47-b107b4cc757d_origin.svg";
          setTimeout(() => {
            Notification.success({
              message: type1.title,
              description: type1.content,
              duration: 6,
              icon: h("img", {
                width: "60",
                style:
                  "position: absolute;width: 40px;border-radius: 50%;border: 2px solid rgba(223,223,223,0.3);",
                src: logo,
              }),
              // icon: h() <img src ={logo}
            });
          }, 3500);
        }
      });
    };
    const close = () => {
      show.value = false;
      const mess_info= {
        id: info.value.id,
        time: info.value.updated_at,
      }
      localStorage.setItem("sys_mess_info",JSON.stringify(mess_info));
    };
    onMounted(() => {
      getMessage();
    });
    return {
      info,
      options,
      show,
      close,
    };
  },
};
</script>
<style lang="stylus" scoped>
.sys-modal {
  display: flex;

  .sys-mask {
    position: fixed;
    top: 0;
    right: 0;
    bottom: 0;
    left: 0;
    z-index: 1000;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.45);
  }

  .sys-mess-wrap {
    position: fixed;
    top: 0;
    right: 0;
    bottom: 0;
    left: 0;
    z-index: 1000;
    overflow: auto;
    outline: 0;
    display: flex;
    align-items: center;
    justify-content: center;

    .sys-mess {
      width: 480px;
      box-shadow: 0 0px 25px 5px rgba(46, 58, 89, 0.1);
      background: linear-gradient(180deg, #dcf4ff, #f4fcff);
      position: relative;
      max-width: 90vw;
      max-height: 90vh;
      padding-top: 20px;
      padding-bottom: 20px;

      .header {
        display: flex;
        justify-content: flex-start;
        align-items: center;

        .logo {
          width: 40px;
          height: 40px;
          margin-right: 10px;
        }

        .title {
          line-height: 24px;
          font-size: 18px;
          color: #666;
          font-weight: 700;

        }
      }

      .body {
        max-height: 50vh;
        overflow-y: auto;
        margin: 10px 0;

        &::-webkit-scrollbar {
          width: 4px;
          border-radius: 2px;
        }

        &::-webkit-scrollbar-thumb {
          width: 8px;
          border-radius: 2px;
          background-color: #ccc;
        }

        >>>p {
          color: #777;
          font-weight: 400;
          font-size: 13px;
          line-height: 22px;
          word-break: break-all;
          text-align: justify;
        }
      }

      .footer {
        display: flex;
        justify-content: center;
        flex-wrap: wrap;
        align-items: center;

        .btn {
          width: 100px;
          height: 45px;
          border: none;
          border-radius: 25px;
          outline: none;
          margin: 10px;
          color: #fff;
          background: #ddd;
          font-weight: 700;
          font-size: 15px;
          line-height: 45px;
          cursor: pointer;
        }

        .know {
          background: linear-gradient(180deg, #fec699, #fee5d2);
        }
      }
    }
  }
}

.fade-enter-active, .fade-leave-active {
  transition: opacity 0.5s ease;
}

.fade-enter-from, .fade-leave-to {
  opacity: 0;
}
</style>