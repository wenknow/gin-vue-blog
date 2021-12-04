<template>
    <div class="article-suspended-panel">
        <div :id="articleDetail.id" :badge="goodCount"
             :class="isGood ? 'like-btn panel-btn  active like-adjust with-badge':'like-btn panel-btn like-adjust with-badge' "
             @click="updateUserArticleGood"
        ></div>

        <div :badge="commentCount"
             class="comment-btn panel-btn comment-adjust  with-badge"
             @click="scrollIntoComment"
        ></div>

    </div>
</template>
<script>
import http from "@/utils/httpindex.js";
export default {
    name: "ArticlePanel",
    props: {
        userAction: {
            type: Object,
            default: () => ({})
        },
        articleDetail: {
            type: Object,
            default: () => ({})
        }
    },
    data() {
        return {
            isGood :false,
            goodCount: 0,
            commentCount: 0
        };
    },
    created() {
        this.isGood = this.userAction.is_good;
        this.goodCount = this.articleDetail.good_count;
        this.commentCount = this.articleDetail.comment_count;
    },
    methods: {
        // 跳转到评论区
        scrollIntoComment() {
            let offsetTop = document.querySelector('.comment-view').offsetTop;
            window.scrollTo({
                top: offsetTop-80
            })
        },
        updateUserArticleGood() {
            let url = "/apis/blog/good/add";
            if (this.isGood){
                url = "/apis/blog/good/del";
            }
            http.post(url, {article_id: this.articleDetail.id}).then(()=>{
                this.isGood = !this.isGood;
                if (this.isGood){
                    this.goodCount++;
                }else {
                    this.goodCount--;
                }
            });
        },
    }
}
</script>
<style lang="stylus" scoped>
    .article-suspended-panel {
        position: fixed;
        /*margin-left: 16rem;*/
        top: 16rem;
        left: calc((100vw - 1550px)/2 - 78px);
    }

    .panel-btn.like-btn.like-adjust {
        background-position: 53% 46%;
    }

    .panel-btn.like-btn {
        background-image: url('~@/assets/images/zan-1.svg');
    }

    .panel-btn.like-btn.active {
        background-image: url('~@/assets/images/zan-active-1.svg');
    }

    .panel-btn {
        position: relative;
        margin-bottom: .75rem;
        width: 3rem;
        height: 3rem;
        background-color: #fff;
        background-position: 50%;
        background-repeat: no-repeat;
        border-radius: 50%;
        box-shadow: 0 2px 4px 0 rgba(0, 0, 0, .04);
        cursor: pointer;
    }

    .panel-btn.comment-btn.comment-adjust {
        background-position: 50% 55%;
    }

    .panel-btn.comment-btn {
        background-image: url('~@/assets/images/comment-1.svg');
    }

    .panel-btn.with-badge:after {
        content: attr(badge);
        position: absolute;
        top: 0;
        left: 75%;
        padding: .1rem .4rem;
        font-size: 1rem;
        text-align: center;
        line-height: 1;
        white-space: nowrap;
        color: #fff;
        background-color: #b2bac2;
        border-radius: .7rem;
        transform-origin: left top;
        transform: scale(.75);
    }

</style>