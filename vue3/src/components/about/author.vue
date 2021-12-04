<template>
    <div class="info concise">
        <div>
            <img :src="authorInfo.head_url" class="avatar"/>
            <p class="name">{{ authorInfo.name }}</p>
        </div>
        <div class="about">
            <div class="about-count">
                <FileTextOutlined />
                <span class="btn quant">文章数： {{ authorInfo.all_article_count }}</span>
            </div>
            <div class="about-count">
                <LikeOutlined />
                <span class="btn quant">被赞数： {{ authorInfo.all_good_count }}</span>
            </div>
            <div class="about-count">
                <FireOutlined />
                <span class="btn quant">被阅读数： {{ authorInfo.all_browse_count }}</span>
            </div>
        </div>
    </div>
</template>
<script>
import {
    FileTextOutlined,
    FireOutlined,
    LikeOutlined,
} from "@ant-design/icons-vue";
export default {
    name: "AboutAuthor",
    props: {
        authorInfo: {
            type: Object,
            default: () => ({})
        },
    },
    components: {
        FileTextOutlined,
        FireOutlined,
        LikeOutlined,
    },
}
</script>
<style lang="stylus" scoped>
.info {
    display: flex;

    .avatar {
        padding: 0 10px;
        width: 80px;
        height: 80px;
        margin: 0;
        padding: 0;
        color: #3da8f5;
        text-align: center;
        cursor: pointer;
        background: #fff;
        line-height: 44px;
        border: 2px solid #dcf4ff
        box-sizing: border-box;
        border-radius: 50%;
        color: #c4c4c4;
        transition all .3s
    }
    &:hover{
    .avatar{
        transform rotate(360deg)!important
    }
    }
    .name {
        font-size: 16px;
        color: #2e3a59;
        font-weight: 700;
        text-align: center;
        font-family: GlowSansSC-ExtendedHeavy;
    }

    .about {
        margin-left: 15px;
        .btn {
            /*padding: 5px 10px;*/
            border-radius: 10px;
            color: rgba(0, 0, 0, 0.45);
            font-size: 14px;
            margin: 0 5px;
        }

        .quant {
            background: linear-gradient(135deg, #fff0ce, #fffcf6);
        }
        .about-count{
            color: #ada593;
            margin-top: 13px;
        }
    }
}
</style>