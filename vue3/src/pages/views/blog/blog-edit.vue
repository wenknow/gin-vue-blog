<template>
  <div class="editor-header">
    <el-input placeholder="请输入文章标题..." class="title-input" v-model="state.params.title" maxlength="60"
              @input="onInputChange"></el-input>
    <div class="status-text">{{state.autoSaveStatus}}</div>
    <el-button type="primary" @click="onClickSubmit">发布</el-button>
  </div>
    <v-md-editor
            v-model="state.params.content"
            :disabled-menus="[]"
            @upload-image="handleUploadImage"
            @change="onInputChange"
            style="height: 100vh"
            ref="md"
    ></v-md-editor>
  <el-dialog title="发布文章" v-model="state.dialogFormVisible" width="550px" >
    <el-form :model="state.form" >
      <el-form-item label="我的分类" required>
        <el-select v-model="state.params.cg_name"  filterable allow-create clearable placeholder="请选择">
          <el-option
                  v-for="item in state.cgList"
                  :key="item.name"
                  :label="item.name"
                  :value="item.name">
          </el-option>
        </el-select>
        <span class="select-hint">可填写新的分类将自动创建。</span>
      </el-form-item>
      <el-form-item label="所属标签" required>
        <el-select v-model="state.params.tag_id_arr" :multiple-limit=3 multiple filterable placeholder="请选择">
          <el-option
                  v-for="item in state.tagList"
                  :key="item.id"
                  :label="item.name"
                  :value="item.id">
          </el-option>
        </el-select>
        <span class="select-hint">最多可以选择3个。</span>
      </el-form-item>
<!--      <el-form-item label="封面图片">-->
<!--        <div class="img-del" v-if="state.params.img_url">-->
<!--          <i class="el-icon-close" :onclick="handleAvatarRemove" circle></i>-->
<!--        </div>-->
<!--        <el-upload-->
<!--                class="avatar-uploader"-->
<!--                name="image"-->
<!--                :action="state.uploadPath"-->
<!--                :headers="state.uploadHeaders"-->
<!--                :show-file-list="false"-->
<!--                :on-success="handleAvatarSuccess"-->
<!--                :before-upload="beforeAvatarUpload">-->
<!--          <img v-if="state.params.img_url" :src="state.params.img_url" class="avatar">-->
<!--          <i v-else class="el-icon-plus avatar-uploader-icon"></i>-->
<!--        </el-upload>-->
<!--      </el-form-item>-->
      <el-form-item >
        <el-button @click="state.dialogFormVisible = false">取 消</el-button>
        <el-button type="primary" :onclick="submitArticle">确 定</el-button>
      </el-form-item>
    </el-form>
  </el-dialog>

</template>
<script>
  import {onMounted, reactive} from "vue";
  import { useStore } from "vuex";
  import { useRoute } from "vue-router";
  import message from "@/components/message/index.js";
  import http from "@/utils/httpindex.js";

  import {
    getCgList,
    getTagList,
  } from "./blog.js";
  export default {
  name: "BlogEdit",
  components:{
  },
  setup() {
    const state = reactive({
      params: {
        id: 0,
        cg_name: "",
        title: "",
        content: "",
        img_url: "",
        tag_id_arr: []
      },
      dialogFormVisible: false,
      form: {},
      formLabelWidth: '120px',
      tagList: [],
      cgList: [],
      autoSaveStatus: "文章将自动保存至草稿箱",
      uploadPath: "/apis/system/uploadImg",
      uploadHeaders: {},
      userInfo:{}
    });
    const store = useStore();
    const route = useRoute();
    onMounted(() => {
      state.userInfo = store.state.user.user;
      if (!state.userInfo){
        window.location.replace("/blog");
      }
      state.params.id = parseInt(route.params.id, 10);
      if (state.params.id){
        getArticleDetail();
      }
      state.cgList = getCgList(state.userInfo.id);
      state.tagList = getTagList();
    });
    // 绑定@imgAdd event
    const handleUploadImage = (event, insertImage, files) => {
      let formData = new FormData();
      let fileName= files[0].name;
      formData.append('image', files[0], fileName);
      http.post('/apis/system/uploadImg',formData).then(function(res) {
        insertImage({
          url: res.url,
          desc: fileName,
        });
      });
    };
      const onInputChange = () => {
        autoSaveArticle();
      };
      const autoSaveArticle = async () => {

        state.autoSaveStatus = "保存中...";
        await http.post('/apis/blog/save',state.params).then(function(res) {
          if (res){
            state.params.id = res;
            state.autoSaveStatus = "保存成功";
          }else {
            state.autoSaveStatus = "自动保存失败";
          }
        });
      };

      const submitArticle = async () => {
        await http.post('/apis/blog/publish',state.params).then(function(res) {
          if (res){
            state.dialogFormVisible = false;
            window.location.replace("/blog/" + state.params.id);
          }
        });
      };
      const onClickSubmit =async () => {
        state.dialogFormVisible = true;
        if ((state.params.title !== "" || state.params.contentText !== "") && state.params.tag_id_arr.length === 0) {
          await http.post("/apis/blog/tag/guess", { title: state.params.title, content: state.params.content}).then((res) => {
            state.params.tag_id_arr = res;
          });
        }
      };

    const getArticleDetail = async ()=> {
      await http.post("/apis/blog/detail", {id:state.params.id}).then((res) => {
          state.params.title = res.detail.title;
          state.params.content = res.detail.content;
          state.params.cg_name = res.detail.cg_name;
          state.params.img_url =  res.detail.img_url;
          if (res.tagList){
            for (const v of res.tagList) {
              state.params.tag_id_arr.push(v.tagId);
            }
          }
       });
    };

      const handleAvatarSuccess = (res, file) => {
        console.log(file);
        if (res.code === 200){
          state.params.img_url = res.url;
        }else {
          message.error(res.msg);
        }
      };
      const handleAvatarRemove = () => {
        state.params.img_url = "";
        console.log(state.params.img_url)
      };
      const beforeAvatarUpload = (file) => {
        const isJPG = file.type === 'image/jpeg';
        const isLt1M = file.size / 1024 / 1024 < 1;
        if (!isJPG) {
          message.error("上传图片只能是 JPG 格式!");
        }
        if (!isLt1M) {
          message.error("上传图片大小不能超过 1MB!");
        }
        return isJPG && isLt1M;
      };
    return {
      state,
      handleUploadImage,
      onClickSubmit,
      onInputChange,
      handleAvatarSuccess,
      beforeAvatarUpload,
      handleAvatarRemove,
      submitArticle
    };
  }
};
</script>

<style lang="stylus" scoped>
  .status-text {
    margin-left: 8px;
    margin-right: 8px;
    font-size: 14px;
    white-space: nowrap;
    color: #c9cdd4;
    cursor: default;
    user-select: none;
  }

  .editor-header {
    display: flex;
    align-items: center;
    padding: 0 27px;
    height: 4.5rem;
    background-color: #fff;
    border-bottom: 1px solid #ddd;
    z-index: 100;
  }

  .title-input {
    font-size: 28px;
    color: #1d2129;
    border: none;
  }

  .el-input__inner {
    border: 0;
  }

  .avatar-uploader .el-upload {
    border: 1px dashed #d9d9d9;
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;
  }

  .avatar-uploader .el-upload:hover {
    border-color: #409EFF;
  }

  .avatar-uploader-icon {
    font-size: 28px;
    color: #8c939d;
    width: 320px;
    height: 180px;
    line-height: 178px;
    text-align: center;
  }

  .avatar {
    height: 180px;
    display: block;
  }
  .img-del {
    margin-left: 240px;
  }
  .select-hint{
    color:red;
    font-size: 13px;
    margin-left: 12px;
  }
</style>
