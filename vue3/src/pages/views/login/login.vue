<template>
  <teleport to="body">
    <transition name="fade">
      <div class="login-modal" v-show="mVisible">
        <div class="login-mask"></div>
        <div class="login-wrap">
          <div class="login concise">
            <span class="close" @click="close"><CloseOutlined /></span>
            <p class="title">欢迎登录</p>
            <el-tabs v-model="state.loginMethod" :stretch="true">
              <el-tab-pane label="免密登录" name="code"></el-tab-pane>
              <el-tab-pane label="密码登录" name="password"></el-tab-pane>
            </el-tabs>
            <el-form>
              <el-formItem
                      :label-width="state.formLabelWidth"
                      required>
                <el-input
                        v-model="state.params.email"
                        placeholder="请输入邮箱"
                        autocomplete="off"
                >
                </el-input>
              </el-formItem>
              <el-formItem
                      v-if="state.loginMethod === 'password'"
                      :label-width="state.formLabelWidth"
                      required>
                <el-input
                        type="password"
                        placeholder="请输入密码"
                        v-model="state.params.password"
                        autocomplete="off"
                ></el-input>
              </el-formItem>
              <el-formItem
                      v-if="state.loginMethod === 'code'"
                      :label-width="state.formLabelWidth"
                      required>
                <div class="bind_code margin_top">
                  <el-input
                          class="bind_code_input code"
                          v-model="state.params.code"
                          type="text"
                          maxlength="6"
                          placeholder="请输入验证码"
                          show-word-limit
                  />
                  <el-button
                          @click="sendEmailCode"
                          class="bind-code-gain"
                          :disabled="state.sendCodeIs"
                  >{{ state.sendCode }}</el-button>
                </div>
              </el-formItem>
            </el-form>
              <div class="login-btn">
                <span v-if="state.loginMethod === 'code'" class="login-hint">新的邮箱地址将会自动进行注册&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;&nbsp;</span>
                <el-button
                      class="btn quant"
                      :loading="state.btnLoading"
                      type="primary"
                      size="mini"
                      @click="handleOk"
              >登录</el-button>
              </div>

            <div class="mode">
              <span class="item qq" title="QQ" @click="partyLogin('qq')">
                <QqOutlined />
              </span>
              <span
                class="item github"
                title="GitHub"
                @click="partyLogin('github')"
              >
                <GithubOutlined />
              </span>
              <span
                class="item gitee"
                title="Gitee"
                @click="partyLogin('gitee')"
              >
                <svg
                  t="1607925359875"
                  class="icon"
                  viewBox="0 0 1024 1024"
                  version="1.1"
                  xmlns="http://www.w3.org/2000/svg"
                  p-id="2562"
                  width="24"
                  height="24"
                  fill="currentcolor"
                >
                  <path
                    d="M512 1024C230.4 1024 0 793.6 0 512S230.4 0 512 0s512 230.4 512 512-230.4 512-512 512z m259.2-569.6H480c-12.8 0-25.6 12.8-25.6 25.6v64c0 12.8 12.8 25.6 25.6 25.6h176c12.8 0 25.6 12.8 25.6 25.6v12.8c0 41.6-35.2 76.8-76.8 76.8h-240c-12.8 0-25.6-12.8-25.6-25.6V416c0-41.6 35.2-76.8 76.8-76.8h355.2c12.8 0 25.6-12.8 25.6-25.6v-64c0-12.8-12.8-25.6-25.6-25.6H416c-105.6 0-188.8 86.4-188.8 188.8V768c0 12.8 12.8 25.6 25.6 25.6h374.4c92.8 0 169.6-76.8 169.6-169.6v-144c0-12.8-12.8-25.6-25.6-25.6z"
                    p-id="2563"
                  ></path>
                </svg>
              </span>
            </div>
          </div>
        </div>
      </div>
    </transition>
  </teleport>
</template>
<script>
import {
  QqOutlined,
  GithubOutlined,
  CloseOutlined,
} from "@ant-design/icons-vue";
import { computed, reactive } from "vue";
import { useStore } from "vuex";
import { useRoute } from "vue-router";
import message from "@/components/message/index.js";
import http from '@/utils/httpindex.js'

export default {
  name: "login",
  components: {
    QqOutlined,
    GithubOutlined,
    CloseOutlined,
  },

  setup() {
    const route = useRoute();
    const store = useStore();
    const state = reactive({
      dialogModel: false,
      btnLoading: false,
      loading: false,
      loginMethod: "code",
      sendCode: "发送验证码",
      sendCodeIs: false,
      formLabelWidth: "10px",
      params: {
        email: "",
        code: "",
        password: "",
      }
    });

    // 手机验证码的倒计时
    const doLoop = (seconds) =>{
      seconds = seconds ? seconds : 60;
      state.sendCode = seconds + "s后获取";
      // state.code = code
      let countdown = setInterval(() => {
        if (seconds > 0) {
          state.sendCode = seconds + "s后获取";
          --seconds;
        } else {
          state.sendCode = "获取验证码";
          state.sendCodeIs = false;
          clearInterval(countdown);
        }
      }, 1000);
    };
    const emailReg = new RegExp(
        "^[a-z0-9]+([._\\-]*[a-z0-9])*@([a-z0-9]+[-a-z0-9]*[a-z0-9]+.){1,63}[a-z0-9]+$"
    ); //正则表达式
    const sendEmailCode = async ()=> {
      if (!state.params.email){
        message.error("邮箱不能为空！");
        return;
      } else if (!emailReg.test(state.params.email)) {
        message.error("请输入格式正确的邮箱！");
        return;
      }
      state.sendCodeIs = true;
      await http.post("/apis/user/sendEmailCode", state.params).then(function(res) {
        if (res){
          message.success("发送成功");
        }
      });

      setTimeout(() => {
        doLoop(60);
      }, 500);
    };

    const submit = async () =>{
      state.btnLoading = true;
      await http.post("/apis/user/login", state.params).then(function(res) {
          state.btnLoading = false;
          store.dispatch("setToken", res.token);
          store.dispatch("userInfo", res.user);
          close();
      });
    };
    const handleOk = ()=> {
      if (!state.params.email) {
        message.error("邮箱不能为空！");
        return;
      } else if (!emailReg.test(state.params.email)) {
        message.error("请输入格式正确的邮箱！");
        return;
      }
      if (state.loginMethod === "password") {
        if (!state.params.password) {
          message.error("密码不能为空！");
          return;
        }
      }
      if (state.loginMethod === "code") {
        if (!state.params.code) {
          message.error("验证码不能为空！");
          return;
        }
      }
      submit();
    };
    const partyLogin = (type) => {
      let url ='';
      if (type === 'github'){
        url = 'https://github.com/login/oauth/authorize?client_id=843089d866f8cab6d455';
      }
      window.location.href = url;
      localStorage.setItem("current_page", route.path);
    };
    const mVisible = computed(() => store.state.index.loginShow);
    const close = () => {
      store.commit("showLogin", false);
    };
    return {
      state,
      partyLogin,
      mVisible,
      close,
      handleOk,
      submit,
      sendEmailCode
    };
  },
};
</script>
<style lang="stylus" scoped>
.login-modal {
  .login-mask {
    position: fixed;
    top: 0;
    right: 0;
    bottom: 0;
    left: 0;
    z-index: 1000;
    height: 100%;
    background-color: rgba(0, 0, 0, 0.45);
  }
.btn {
  padding: 5px 20px;
  border-radius: 10px;
  color: rgba(0, 0, 0, 0.45);
  font-size: 14px;
  margin: 0 5px;
}
.login-btn{
    display: flex;
    align-items: center;
    justify-content: right;
}
.login-hint{
  margin: 0px 0px 0px 14px;
  font-size: 14px;
  color: red;
}
.quant {
    background: linear-gradient(148deg, #ddf5ff, #fffcf6);
}
  .login-wrap {
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

    .login {
      width: 430px;
      box-shadow: 0 0px 25px 5px rgba(46, 58, 89, 0.1);
      background: linear-gradient(180deg, #dcf4ff, #f4fcff);
      position: relative;

      &:hover {
        box-shadow: 0 30px 50px 20px rgba(46, 58, 89, 0.1);
      }
      .bind-code-gain {
        position: absolute;
        right: 10px;
        font-size: 14px;
        font-family: MicrosoftYaHei;
        color: #20aee5;
        line-height: 18px;
        cursor: pointer;
        padding-left: 10px;
        border-left: 1px solid #d8d8d8;
        border: 0;
        background: none;
        padding: 10px;
        border-radius: 0;
      }
      .bind_code {
        position: relative;
      }

      .code /deep/.el-input__suffix {
        right: 97px;
      }
      .close {
        position: absolute;
        right: 20px;
        top: 20px;
        font-size: 20px;
        color: #ccc;
        cursor: pointer;
        width: 30px;
        height: 30px;
        display: flex;
        justify-content: center;
        align-items: center;
        margin-bottom: 10px;
      }

      .title {
        font-size: 25px;
        text-align: center;
        line-height: 46px;
        color: #2e3a59;
        margin-bottom: 10px;
        font-weight: 700;
        font-family: GlowSansSC-ExtendedHeavy;
      }

      .alert {
        margin-bottom: 16px;
        box-sizing: border-box;
        margin: 0;
        color: rgba(0, 0, 0, 0.65);
        font-size: 14px;
        font-variant: tabular-nums;
        line-height: 1.5715;
        list-style: none;
        font-feature-settings: 'tnum';
        position: relative;
        padding: 8px 15px 8px 37px;
        word-wrap: break-word;
        border-radius: 20px;
        background-color: #fffbe6;
        position: relative;

        .icon {
          position: absolute;
          top: 12.0005px;
          left: 16px;
          color: #faad14;
        }
      }

      .mode {
        display: flex;
        align-items: center;
        justify-content: center;

        .item {
          display: flex;
          justify-content: center;
          align-items: center;
          width: 48px;
          height: 48px;
          border: 2px solid #efefef;
          border-radius: 50%;
          font-size: 24px;
          margin: 10px;
          transition: all 0.3s;
          cursor: pointer;
        }

        .qq {
          color: #7bd4ef;

          &:hover {
            background: rgba(123, 212, 239, 0.4);
            border-color: #7bd4ef;
          }
        }

        .github {
          color: #383838;

          &:hover {
            background: rgba(56, 56, 56, 0.4);
            border-color: #383838;
          }
        }

        .gitee {
          color: #fe7300;

          &:hover {
            background: rgba(254, 115, 0, 0.4);
            border-color: #fe7300;
          }
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

@media (max-width: 500px) {
  .login-modal {
    .login-wrap {
      .login {
        width: 90%;
        margin: 0 5%;

        .title {
          font-size: 20px;
          margin-bottom: 10px;
        }
      }
    }
  }
}
</style>