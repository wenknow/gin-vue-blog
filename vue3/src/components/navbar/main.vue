<template>
  <header
    class="navbar"
    :class="{
      'navbar-folding': folding,
      'navbar-open': visible && folding,
      'navbar-fixed': fixed,
    }"
    :style="arriveHeight ? style.arrive : style.normal"
  >
    <nav class="nav-main">
      <div class="nav-header">
        <router-link :to="url">
          <slot name="logo">
            <span
              v-if="logo"
              class="nav-brand"
              :style="{ 'background-image': `url(${logo})` }"
            >
            </span>
          </slot>
        </router-link>
<!--        <router-link :to="url">-->
<!--          <slot name="title">-->
<!--            <span class="nav-title">{{ title }}</span>-->
<!--          </slot>-->
<!--        </router-link>-->
      </div>
      <collapse-transition>
        <ul
          class="nav-section"
          v-show="foldShow"
          :style="{ 'justify-content': sectionPosition }"
        >
          <li v-for="item of section" :key="item.key">
            <slot :name="item.slot" :item="item">
              <DropDown
                :menu="item.children ? true : false"
                :float="folding ? false : true"
                :mode="folding ? 'click' : 'hover'"
              >
                <template v-slot="{ visible }">
                  <router-link
                    :to="item.link"
                    v-if="!item.children"
                    style="display: block; width: 100%; height: 100%"
                  >
                    <span
                      class="nav-section-item"
                      :class="{ active: $route.path == item.link }"
                      @click="close"
                    >
                      {{ item.title }}
                    </span>
                  </router-link>
                  <span class="nav-section-item" v-else>{{ item.title }}</span>
                  <span
                    v-if="item.children"
                    class="nav-triangle nav-arrow-icon"
                    :style="{ transform: `rotate(${visible ? 180 : 0}deg)` }"
                  ></span>
                </template>
                <template #menu>
                  <ul class="nav-section-children">
                    <li
                      v-for="chil of item.children"
                      :key="chil.key"
                      @click="close"
                    >
                      <slot :name="chil.slot || item.childSlot" :item="chil">
                        <router-link
                          :to="chil.link"
                          custom
                          v-slot="{ navigate }"
                        >
                          <span
                            class="nav-section-children-item"
                            @click="navigate"
                            role="link"
                            >{{ chil.title }}</span
                          >
                        </router-link>
                      </slot>
                    </li>
                  </ul>
                </template>
              </DropDown>
            </slot>
          </li>
        </ul>
      </collapse-transition>

      <div class="nav-footer-wrap">
        <slot name="footer"> </slot>
        <slot name="toggle">
          <span class="navbar-toggle-btn" @click="toggle">
            <p class="menu" :class="{ 'menu-active': visible }"></p>
          </span>
        </slot>
      </div>
    </nav>
  </header>
</template>
<script>
import { computed, onMounted, onUnmounted, ref } from "vue";
import CollapseTransition from "./text";
import { throttle } from "throttle-debounce";
import DropDown from "./drop-down";
// import { useRoute } from 'vue-router';
export default {
  name: "Navbar",
  components: {
    DropDown,
    CollapseTransition,
  },
  props: {
    url: {
      type: String,
      default: "/",
    },
    logo: {
      type: String,
      default: "",
    },
    title: {
      type: String,
      default: "小白组件库",
    },
    section: {
      type: Object,
    },
    sectionPosition: {
      type: String,
      default: "center",
    },
    width: {
      type: Number,
      default: 767,
    },
    height: {
      type: Number,
      default: 120,
    },
    fixed: {
      type: Boolean,
      default: true,
    },
    autoClose: {
      type: Boolean,
      default: true,
    },
    style: {
      type: Object,
      default: () => {
        return {
          normal: {
            "background-color": "#fff",
          },
          arrive: {
            "background-color": "#fff",
            "box-shadow": "0 4px 30px rgba(0,0,0,.07)",
            // 'backdrop-filter': 'blur(24px)'
          },
        };
      },
    },
    getContainer: {
      type: Function,
      default: () => window,
    },
  },
  setup(props) {
    //   是否显示
    const visible = ref(false);
    //   是否折叠
    const folding = ref(false);
    const arriveHeight = ref(false);
    const getScrollTop=()=>{
        const ele = props.getContainer();
        let scroll_top;
        if(ele==window){
             scroll_top = document.documentElement.scrollTop||document.body.scrollTop;
        }else{
            scroll_top=ele.scrollTop
        }
        return scroll_top;
    }
    const scroll = () => {
      const scrollTop = getScrollTop()
      if (scrollTop >= props.height) {
        arriveHeight.value = true;
      } else {
        arriveHeight.value = false;
      }
    };
    const reSize = () => {
      const width = document.body.clientWidth;
      if (width < props.width) {
        folding.value = true;
      } else {
        folding.value = false;
      }
    };
    const throReSize = throttle(200, reSize);
    const throScroll = throttle(200, scroll);
    const toggle = () => {
      visible.value = !visible.value;
    };
    const foldShow = computed(() => {
      if (!folding.value) {
        return true;
      } else {
        return visible.value;
      }
    });
    const close = () => {
      if (folding.value && props.autoClose) {
        setTimeout(() => {
          visible.value = false;
        }, 300);
      }
    };
    onMounted(() => {
      scroll();
      reSize();
      props.getContainer().addEventListener("scroll", throScroll);
      window.addEventListener("resize", throReSize);
    });
    onUnmounted(() => {
      props.getContainer().removeEventListener("scroll", throScroll);
      window.removeEventListener("resize", throReSize);
    });
    return {
      toggle,
      visible,
      folding,
      arriveHeight,
      foldShow,
      close,
    };
  },
};
</script>
<style lang="stylus" scoped>
.navbar {
  width: 100%;
  height: 72px;
  z-index: 233;
  transition: all 0.3s;
  box-sizing: border-box;
  color: rgba(0, 0, 0, 0.35);

  ul {
    list-style: none;
    margin: 0;
    padding: 0;
  }

  a {
    text-decoration: none;
  }

  .nav-main {
    display: flex;
    height: 100%;
    width: 70%;
    padding-left: 10px;
    margin: 0 auto;
    position: relative;
    min-height: 50px;
    box-sizing: border-box;

    .nav-header {
      display: flex;
      align-items: center;
      margin-right: 30px;

      .nav-brand {
        display: inline-block;
        width: 50px;
        height: 36px;
        background-image: url('https://g.alicdn.com/teambition-site/site/images/header/logo-one.svg'); // g.alicdn.com/teambition-site/site/images/header/logo-one.svg);
        background-repeat: no-repeat;
        background-position: center;
        background-size: contain;
        text-indent: -9999px;
      }

      .nav-title {
        font-style: normal;
        font-weight: 400;
        font-size: 14px;
        line-height: 14px;
        color: #1b9aee;
        border-left: 1px solid #ccc;
        padding-left: 15px;
        margin-left: 15px;
      }
    }

    .nav-section {
      display: flex;
      justify-content: center;
      align-items: center;
      margin: 0;
      &>li {
        height: 100%;
        padding: 0 15px;
        a{
          color inherit
        }
        .nav-section-item {
          cursor: pointer;
          display: flex;
          align-items: center;
          height: 100%;
          position: relative;
          padding-left: 4px;
          padding-right: 4px;
          // color: #383838;
        }

        .active {
          border-bottom: 2px solid #1b9aee;
          color: #1b9aee;
        }

        &:hover {
          .nav-section-item {
            color: #1b9aee;
          }

          .nav-section-children-wrap {
            display: block;
          }
        }

        .nav-triangle {
          width: 20px;
          height: 20px;
          background: url('https://g.alicdn.com/teambition-site/site/images/footer/triangle.png') center no-repeat; // g.alicdn.com/teambition-site/site/images/footer/triangle.png) center no-repeat;
          text-align: center;
          background-size: contain;
        }

        .nav-arrow-icon {
          display: inline-block;
          width: 10px;
          margin-left: 8px;
          height: 20px;
          line-height: 20px;
          text-align: center;
          font-size: 12px;
          color: #bfbfbf;
          transition: all 0.3s;
        }

        .nav-section-children {
          &>li {
            .nav-section-children-item {
              padding: 10px 0px;
              display: block;
              border-radius: 4px;
              transition: all 0.3s;
              text-align: center;
              cursor: pointer;
              white-space: nowrap;

              &:hover {
                background-color: #edf5fc;
              }
            }
          }
        }
      }
    }

    .nav-footer-wrap {
      display: flex;
      align-items: center;
      height: 100%;
      justify-content: flex-end;
      flex: 1;
      text-align: right;

      .nav-footer {
        display: flex;
        position: relative;
        align-items: center;
        justify-content: flex-end;
        margin: 0;
        height: 100%;
      }

      .navbar-toggle-btn {
        position: relative;
        display: none;
        width: 48px;
        height: 48px;
        margin: 0;
        padding: 0;
        color: #3da8f5;
        text-align: center;
        align-items: center;
        cursor: pointer;
        background: #fff;
        line-height: 44px;
        border: 2px solid rgba(223, 223, 223, 0.3);
        box-sizing: border-box;
        border-radius: 50%;
        color: #c4c4c4;

        .menu {
          width: 16px;
          height: 16px;
          position: relative;
          margin: 14px auto;

          &:before, &:after {
            content: '';
            display: block;
            width: 16px;
            height: 2px;
            background: rgba(223, 223, 223, 0.8);
            border-radius: 1px;
            position: absolute;
            left: 0;
            transition: all 0.15s ease-in-out;
          }

          &:before {
            top: 2px;
            box-shadow: 0 5px rgba(223, 223, 223, 0.8);
          }

          &:after {
            bottom: 2px;
          }
        }

        .menu-active {
          &:before, &:after {
            background: #f974a1;
          }

          &:before {
            top: 7px;
            box-shadow: none;
            transform: rotate(225deg);
          }

          &:after {
            bottom: 7px;
            transform: rotate(135deg);
          }
        }
      }
    }
  }

  .navbar-collapse {
    background-color: #fff;
    padding-right: 16px;
    padding-left: 16px;
    border-top: 1px solid transparent;
    box-shadow: inset 0 1px 0 rgba(255, 255, 255, 0.1);

    .navbar-collapse-item {
      color: #383838;
      display: flex;
      font-size: 16px;
      line-height: 22px;
      text-align: center;
      align-items: center;
      justify-content: space-between;
      position: relative;
      padding: 25px 0;
      border-bottom: 1px solid #f5f5f5;
      font-size: 16px;
      line-height: 22px;
      text-align: center;
    }
  }
}

.navbar-fixed {
  position: fixed;
  top: 0;
}

.navbar-open {
  background: #fff !important;
}

.navbar-folding {
  .nav-footer {
    display: none !important;
  }

  .nav-main {
    padding: 0 16px;
    position: relative;

    .nav-section {
      position: absolute;
      width: 100%;
      top: 0;
      left: 0;
      background: #fff;
      display: block;
      transition: all 0.3s;
      padding: 8px 36px;
      box-sizing: border-box;
      overflow-y: auto;
      max-height: 100vh;
      padding-top: 72px;
      z-index: -1;

      &::-webkit-scrollbar {
        width: 0 !important;
      }

      &>li {
        display: block;
        cursor: pointer;
        margin: 0;
        padding: 0;
        position: relative;

        >>>.drop-down {
          margin-right: 0px;

          .drop-down-content {
            .nav-section-item {
              display: block;
              color: #383838;
              font-size: 16px;
              line-height: 22px;
              text-align: center;
              padding: 25px 0;
              border-bottom: 1px solid #f5f5f5;
              margin: 0 auto;
              width: 100%;
            }

            .nav-triangle {
              position: absolute;
              right: 0;
              width: 10px;
              height: 10px;
              line-height: 24px;
              text-align: center;
              color: #a6a6a6;
              -webkit-transition: all 0.3s;
              transition: all 0.3s;
            }
          }

          .drop-down-menu {
            box-shadow: none;
            padding: 0;
            padding-top: 8px;
            display: block;

            .nav-section-children>li {
              margin: 0 33px;
              border-bottom: 1px solid #f5f5f5;

              .nav-section-children-item {
                padding: 23px 0;
              }
            }
          }
        }
      }
    }
  }

  .nav-footer-wrap {
    flex: 1 1 auto;

    .navbar-toggle-btn {
      display: inline-block !important;
    }
  }
}
</style>