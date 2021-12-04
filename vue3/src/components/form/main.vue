<template>
  <form class="validate-form-container">
    <slot>
      <FormItem></FormItem>
    </slot>
  </form>
</template>

<script>
import { defineComponent, onUnmounted } from "vue";
import mitt from "mitt";
import FormItem from "./form-item";
export const emitter = mitt();
export default defineComponent({
  emits: ["form-submit"],
  components: {
    FormItem,
  },
  setup(props, context) {
    let funcArr = [];
    const submitForm = () => {
      const result = funcArr.map((func) => func()).every((result) => result);
      context.emit("form-submit", result);
    };
    const callback = (func) => {
      if (func) {
        funcArr.push(func);
      }
    };
    emitter.on("form-item-created", callback);
    onUnmounted(() => {
      emitter.off("form-item-created", callback);
      funcArr = [];
    });
    return {
      submitForm,
    };
  },
});
</script>
