<template>
  <div class="project-wrap">
    <div class="occlusion-wrap">
      <div class="occlusion">
        <h4 class="title">Hello World</h4>
        <div class="content">There is no end to learning.</div>
      </div>
      <div class="waves-area">
        <svg
          class="waves-svg"
          xmlns="http://www.w3.org/2000/svg"
          xmlns:xlink="http://www.w3.org/1999/xlink"
          viewBox="0 24 150 28"
          preserveAspectRatio="none"
          shape-rendering="auto"
        >
          <defs>
            <path
              id="gentle-wave"
              d="M -160 44 c 30 0 58 -18 88 -18 s 58 18 88 18 s 58 -18 88 -18 s 58 18 88 18 v 44 h -352 Z"
            ></path>
          </defs>
          <g class="parallax">
            <use xlink:href="#gentle-wave" x="48" y="0"></use>
            <use xlink:href="#gentle-wave" x="48" y="3"></use>
            <use xlink:href="#gentle-wave" x="48" y="5"></use>
            <use xlink:href="#gentle-wave" x="48" y="7"></use>
          </g>
        </svg>
      </div>
    </div>
    <div class="project container">
      <div class="project-content">
        <template v-if="loading">
          <Skeleton
            class="project-item"
            style="padding: 20px"
            v-for="item of 3"
            :key="item"
            active
          />
        </template>
        <template v-else>
          <a
            :href="item.link"
            target="_blank"
            class="project-item"
            v-for="item of list"
            :key="item.id"
          >
            <div class="cover">
              <Image :src="item.imgUrl" class="img" fit="cover">
                  <template #placeholder>
                        <svg xmlns="http://www.w3.org/2000/svg" style="margin:auto;width: 100%;height: 100%;" width="200" height="200" viewBox="0 0 100 100" preserveAspectRatio="xMidYMid" display="block"><rect x="47" y="24" rx="3" ry="6" width="6" height="12" fill="#93dbe9"><animate attributeName="opacity" values="1;0" keyTimes="0;1" dur="1s" begin="-0.9166666666666666s" repeatCount="indefinite"/></rect><rect x="47" y="24" rx="3" ry="6" width="6" height="12" fill="#93dbe9" transform="rotate(30 50 50)"><animate attributeName="opacity" values="1;0" keyTimes="0;1" dur="1s" begin="-0.8333333333333334s" repeatCount="indefinite"/></rect><rect x="47" y="24" rx="3" ry="6" width="6" height="12" fill="#93dbe9" transform="rotate(60 50 50)"><animate attributeName="opacity" values="1;0" keyTimes="0;1" dur="1s" begin="-0.75s" repeatCount="indefinite"/></rect><rect x="47" y="24" rx="3" ry="6" width="6" height="12" fill="#93dbe9" transform="rotate(90 50 50)"><animate attributeName="opacity" values="1;0" keyTimes="0;1" dur="1s" begin="-0.6666666666666666s" repeatCount="indefinite"/></rect><rect x="47" y="24" rx="3" ry="6" width="6" height="12" fill="#93dbe9" transform="rotate(120 50 50)"><animate attributeName="opacity" values="1;0" keyTimes="0;1" dur="1s" begin="-0.5833333333333334s" repeatCount="indefinite"/></rect><rect x="47" y="24" rx="3" ry="6" width="6" height="12" fill="#93dbe9" transform="rotate(150 50 50)"><animate attributeName="opacity" values="1;0" keyTimes="0;1" dur="1s" begin="-0.5s" repeatCount="indefinite"/></rect><rect x="47" y="24" rx="3" ry="6" width="6" height="12" fill="#93dbe9" transform="rotate(180 50 50)"><animate attributeName="opacity" values="1;0" keyTimes="0;1" dur="1s" begin="-0.4166666666666667s" repeatCount="indefinite"/></rect><rect x="47" y="24" rx="3" ry="6" width="6" height="12" fill="#93dbe9" transform="rotate(210 50 50)"><animate attributeName="opacity" values="1;0" keyTimes="0;1" dur="1s" begin="-0.3333333333333333s" repeatCount="indefinite"/></rect><rect x="47" y="24" rx="3" ry="6" width="6" height="12" fill="#93dbe9" transform="rotate(240 50 50)"><animate attributeName="opacity" values="1;0" keyTimes="0;1" dur="1s" begin="-0.25s" repeatCount="indefinite"/></rect><rect x="47" y="24" rx="3" ry="6" width="6" height="12" fill="#93dbe9" transform="rotate(270 50 50)"><animate attributeName="opacity" values="1;0" keyTimes="0;1" dur="1s" begin="-0.16666666666666666s" repeatCount="indefinite"/></rect><rect x="47" y="24" rx="3" ry="6" width="6" height="12" fill="#93dbe9" transform="rotate(300 50 50)"><animate attributeName="opacity" values="1;0" keyTimes="0;1" dur="1s" begin="-0.08333333333333333s" repeatCount="indefinite"/></rect><rect x="47" y="24" rx="3" ry="6" width="6" height="12" fill="#93dbe9" transform="rotate(330 50 50)"><animate attributeName="opacity" values="1;0" keyTimes="0;1" dur="1s" begin="0s" repeatCount="indefinite"/></rect></svg>
                  </template>
                  <template #error>
                        <img
                          style="width: 100%;"
                          src="https://iconfont.alicdn.com/s/e8a7a96a-7f80-4f07-bc88-c162da451fe1_origin.svg"
                  />
                </template>
              </Image>
            </div>
            <div class="content">
              <span class="title">{{ item.name }}</span>
              <p class="info">{{ item.info }}</p>
            </div>
          </a>
        </template>
      </div>
    </div>
  </div>
</template>
<script>
import { getProjectList } from "./circle.js";
import Skeleton from "@/components/skeleton/skeleton";
import Image from "@/components/image/image";

export default {
  name: "Project",
  components: {
    Skeleton,
    Image
  },
  setup() {
    const { loading, list } = getProjectList();
    return {
      loading,
      list,
    };
  },
};
</script>
<style lang="stylus" scoped>
.project {
  .project-content {
    display: flex;
    justify-content: center;
    align-items: center;
    flex-wrap: wrap;
    margin: 0 auto;
    margin-top: 40px;
  }

  .project-item {
    background-color: #fff;
    margin: 18px;
    width: 300px;
    height: 400px;
    border: 1px solid #e8e8e8;
    border-radius: 8px;
    transition: all 0.6s;
    cursor: pointer;

    &:hover {
      box-shadow: 0 2px 20px rgba(0, 0, 0, 0.09);
      transform: translateY(-12px);
    }

    .cover {
      .img {
        width: 100%;
        height: 200px;
        border-top-left-radius: 8px;
        border-top-right-radius: 8px;
        object-fit: cover;
      }
    }

    .content {
      display: flex;
      flex-direction: column;
      padding: 23px 20px;

      .title {
        margin-bottom: 10px;
        width: 258px;
        overflow: hidden;
        text-overflow: ellipsis;
        white-space: nowrap;
        font-size: 18px;
        font-weight: 500;
        line-height: 28px;
        color: #383838;
      }

      .info {
        display: -webkit-box;
        -webkit-box-orient: vertical;
        -webkit-line-clamp: 4;
        overflow: hidden;
        text-overflow: ellipsis;
        font-size: 14px;
        line-height: 26px;
        color: gray;
        letter-spacing: 0.3px;
      }
    }
  }
}

@media (max-width: 776px) {
  .project {
    .project-content {
      padding: 10px;
      box-sizing: border-box;

      .project-item {
        width: 100%;
      }
    }
  }
}
</style>