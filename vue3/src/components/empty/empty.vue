<template>
  <div class="by-empty">
    <div class="by-empty-image" :style="imageStyle">
      <component
        v-for="(item, index) of imageSlot"
        :key="'imageSlot' + index"
        :is="item"
      ></component>
    </div>
    <div class="by-empty-description" v-if="description || $slots.description">
      <component
        v-for="(item, index) of descSlot"
        :key="'descSlot' + index"
        :is="item"
      ></component>
    </div>
    <div class="by-empty-footer" v-if="$slots.default">
      <slot></slot>
    </div>
  </div>
</template>
<script>
import { computed, h } from "vue";
import EmptyDefault from "./empty-default";
import EmptySimple from "./empty-simple";

export default {
  name: "Empty",
  props: {
    description: {
      type: [String, Boolean],
    },
    imageStyle: {
      type: Object,
    },
    image: {
      type: String,
    },
  },
  components: {
    EmptyDefault,
    EmptySimple,
  },
  setup(props, context) {
    const imageSlot = computed(() => {
      let slot = ["EmptyDefault"];
      if (props.image) {
        if (props.image == "EmptySimple") {
          slot[0] = "EmptySimple";
        } else {
          // <img src={props.image}/>
          slot[0] = h("img", {
            src: props.image,
          });
        }
      }
      const actions = context.slots["image"] ? context.slots["image"]() : "";
      if (actions || actions.length) {
        slot = actions;
      }
      return slot;
    });
    const descSlot = computed(() => {
      let slot = [];
      if (props.description) {
        slot.push(h("p", props.description));
      }
      const actions = context.slots["description"]
        ? context.slots["description"]()
        : "";
      console.log(actions);
      if (actions || actions.length) {
        slot = actions;
      }
      return slot;
    });
    return {
      imageSlot,
      descSlot,
    };
  },
};
</script>
<style lang="stylus" scoped>
.by-empty {
  margin: 0 8px;
  font-size: 14px;
  line-height: 22px;
  text-align: center;

  .by-empty-image {
    height: 100px;
    margin-bottom: 8px;

    svg {
      height: 100%;
      margin: auto;
    }

    img {
      height: 100%;
      vertical-align: middle;
      border-style: none;
    }
  }

  .by-empty-description {
    margin: 0;
  }

  .by-empty-footer {
    margin-top: 16px;
  }
}
</style>