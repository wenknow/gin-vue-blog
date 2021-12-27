<template>
  <div class="about-article article-concise">
    <div class="about-article__title">我的分类</div>
    <ul v-for="item in cgList" :key="item.id" >
      <router-link :key="item.id" :to="`?cg_id=${item.id}`">
      <li class="article-item">
        <span class="cg-name">{{item.name}}</span>
        <span class="btn-edit" v-if="isAuthor"><EditOutlined  @click="submitArticleCgDialog(item.id, item.name)" /></span>
        <span class="btn-edit" v-if="isAuthor"><DeleteOutlined @click="delArticleCgDialog(item.id)" /></span>
      </li>
      </router-link>
    </ul>
  </div>
  <el-dialog title="修改我的分类" v-model="dialogEditVisible" width="400px">
    <el-form :model="editFrom">
      <el-form-item label="分类名称" >
        <el-input v-model="editFrom.name" ></el-input>
      </el-form-item>
    </el-form>
    <div  class="dialog-footer">
      <el-button @click="dialogEditVisible = false">取 消</el-button>
      <el-button type="primary" @click="editMyArticleCg()">确 定</el-button>
    </div>
  </el-dialog>
  <el-dialog v-model="dialogDelVisible" title="删除确认" width="400px">
    <el-button type="danger" @click="dialogDelVisible=false">取消</el-button>
    <el-button type="primary" @click="delMyArticleCg()">确认</el-button>
  </el-dialog>
</template>

<script>
  import http from "@/utils/httpindex.js";
  import {
    EditOutlined,
    DeleteOutlined,
  } from "@ant-design/icons-vue";
export default {
  props: {
    isAuthor: {
      type: Boolean,
      default: false
    },
    uid: {
      type: Number,
      default: 0
    },
  },
  components: {
    EditOutlined,
    DeleteOutlined,
  },
  data(){
    return{
      dialogEditVisible: false,
      dialogDelVisible: false,
      cgList: [],
      delCgId: 0,
      editFrom:{
        id: 0,
        name: ""
      }
    }
  },
  watch: {
    uid: {
      handler: function () {
        this.getCgList();
      },
      deep: true,
    },
  },

  methods:{
    submitArticleCgDialog(id, name){
      this.dialogEditVisible = true;
      this.editFrom.id = id;
      this.editFrom.name = name;
    },
    delArticleCgDialog(id){
      this.dialogDelVisible = true;
      this.delCgId = id;
    },
    getCgList() {
      http.post("/apis/blog/cg/list", {uid:this.uid}).then((res) => {
        this.cgList = res;
      });
    },
    async editMyArticleCg() {
      await http.post("/apis/blog/cg/update",this.editFrom).then(() => {
        this.dialogEditVisible = false;
        this.getCgList();
      });
    },
    async delMyArticleCg() {
      await http.post("/apis/blog/cg/del", {id:this.delCgId}).then(() => {
        this.dialogDelVisible = false;
        this.getCgList();
      });
    },
  }
}
</script>

<style lang='stylus' scoped>
  .cg-name{
    color:#333;
    margin-right: 10px;
  }
  .btn-edit{
    color:#bb99ff;
    margin-right: 5px;
  }
.about-article{


  .about-article__title{
    padding: 15px;
    font-size: 16px;
    border-bottom: 1px solid #eee;
  }
  
  .article-item{
    display: block;
    padding: 0.8rem 1.3rem;
    font-size: 14px;
    line-height: 1.3;
    color: #333;
    cursor: pointer;


    &:hover{
      background: hsla(0,0%,84.7%,.1);
    }

    .article__meta{
      display: flex;
      align-items: center;
      margin-top: 10px;
      color: #b2bac2;
      font-size: 14px;

      .meta-item{
        display: flex;
        align-items: center;

        &:not(:last-child){
          margin-right: 20px;
        }

        .meta-item__icon{
          width: 16px;
          height: 16px;
          margin-right: 5px;
        }
      }
    }
  }
}
</style>