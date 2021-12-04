<template>
  <div
    class="by-slider"
    :class="{ 'by-slider-vertical': vertical, 'by-slider-disabled': disabled }"
    ref="slider"
    tabindex="-1"
    @touchstart="disabled ? dis() : onMousedown($event, false)"
    @mousedown="disabled ? dis() : onMousedown($event)"
    @mouseup="disabled ? dis() : onMouseup($event)"
    @keydown="disabled ? dis() : onKeydown($event)"
    @keyup="disabled ? dis() : onKeyup($event)"
    @focus="disabled ? dis() : onFocus($event)"
    @blur="disabled ? dis() : onBlur($event)"
  >
    <div class="by-slider-rail"></div>
    <div class="by-slider-track" :style="track"></div>
    <div class="by-slider-step"></div>
    <div class="by-slider-handle" :class="{'by-handle-show':handleShow}" ref="sliderHandle" tabindex="0" :style="handle"></div>
    <div class="by-slider-mark"></div>
  </div>
</template>
<script>
import { throttle } from "throttle-debounce";
export default {
  name: "BuSlider",
  props: {
    value:Number,
    disabled: {
      default: false,
      type: Boolean,
    },
    max: {
      default: 100,
      type: Number,
    },
    min: {
      default: 0,
      type: Number,
    },
    vertical: {
      default: false,
      type: Boolean,
    },
    reverse: {
      default: false,
      type: Boolean,
    },
    step:{
        default:1,
        type:Number
    },
    handleShow:{
      type:Boolean,
      default:true
    }
  },
  computed: {
    track() {
      let style = {};
      const position = ((this.sliderValue - this.min) / (this.max - this.min)) * 100;
      if (this.vertical) {
        style = {
          bottom: this.reverse ? "auto" : `0%`,
          top: this.reverse ? `0%` : "auto",
          height: `${position}%`,
        };
      } else {
        style = {
          left: this.reverse ? "auto" : `0%`,
          right: this.reverse ? `0%` : "auto",
          width: `${position}%`,
        };
      }
      return style;
    },
    handle() {
      let style = {};
      const position = ((this.sliderValue - this.min) / (this.max - this.min)) * 100;
      //   const rePosition=((1-this.value-this.min)/(this.max-this.min))*100
      if (this.vertical) {
        style = {
          bottom: this.reverse ? "auto" : `${position}%`,
          top: this.reverse ? `${position}%` : "auto",
          transform: `${this.reverse ?  "translateY(0%)" : "translateY(50%)"}`,
        };
      } else {
        style = {
          left: this.reverse ? "auto" : `${position}%`,
          right: this.reverse ? `${position}%` : "auto",
          transform: `${this.reverse ? "translateX(50%)" : "translateX(-50%)"}`,
        };
      }
      return style;
    },
  },
  data() {
    return {
      sliderValue:this.min,
      isDown: false,
      isPc: false,
      model:true,
    };
  },
  watch:{
      value(newV){
          if(this.step!=1){
                newV=this.onStep(newV)
          }
          if(this.model){
            this.sliderValue=this.limit(newV)
          }
      },
      sliderValue(newV){
        this.$emit('change',newV)
        this.$emit('update:value',newV)
      }
  },
  created(){
    if(this.value){
      this.sliderValue=this.limit(this.value)
    }
  },
  methods: {
    dis() {
        console.log('111')
    },
    onMousedown(e, pc = true) {
      this.isPc = pc;
      this.removeDocumentEvents();
      this.addDocumentTouchEvents();
      console.log(e)
      this.onMove(e);
      this.onMouseup()
      e.stopPropagation();
      e.preventDefault();
    },
    onMouseup(){
        // const handle=this.$refs.sliderHandle
        // handle.focus()
        this.$refs.slider.focus()

    },
    onKeydown(e) {
      console.log("onKeydown");
       e.preventDefault();
       this.model=false
        const keyCode = e.keyCode;
        switch(keyCode){
          // LEFT_ARROW
          case 37:
            this.sliderValue=this.limit(this.sliderValue-this.step)
            break;
          // UP_ARROW
          case 38:
            this.sliderValue=this.limit(this.sliderValue+this.step)
            break;
          // RIGHT_ARROW
          case 39:
            this.sliderValue=this.limit(this.sliderValue+this.step)
            break;
          // DOWN_ARROW
          case 40:
            this.sliderValue=this.limit(this.sliderValue-this.step)
            break; 
        }
    },
    onKeyup(){
      this.model=true
      this.$emit('after-change',this.sliderValue)
    },
    onFocus() {
      console.log("onFocus");
    },
    onBlur() {
      console.log("onBlur");
    },
    onEnd() {
      console.log("end");
      this.model=true
      this.$emit('after-change',this.sliderValue)
      this.removeDocumentEvents();
    },
    onMove(e) {
      this.model=false
      e= this.isPc ?e:e.changedTouches[0]
      const slider = this.$refs.slider;
      const width = slider.clientWidth;
      const height = slider.clientHeight;
      const offsetX = e.clientX - slider.getBoundingClientRect().left;
      const offsetY = e.clientY - slider.getBoundingClientRect().top;
      const info = this.max - this.min;
      console.log(height,offsetY)
      let elePosition;
      if (this.vertical) {
        elePosition =this.min+(1- (offsetY / height)) * info;
      } else {
        elePosition =this.min+(offsetX / width) * info;
      }
      elePosition =this.reverse
        ?  this.max-elePosition+this.min
        : elePosition
      if(this.step!=1){
          elePosition=this.onStep(elePosition)
      }
      elePosition= this.limit(elePosition)
      this.sliderValue =elePosition;
      console.log("move",this.sliderValue);
    },
    limit(value){
        if(value>this.max){
            value=this.max
        }else if(value<this.min){
            value=this.min
        }
        return parseInt(value) 
    },
    onStep(value){
        let i=parseInt(value/this.step)*this.step
        if(value>i+this.step/2){
          value=i+this.step
        }else{
          value=i
        }
        return value
    }
    ,
    addDocumentTouchEvents() {
      this.test = throttle(100,this.onMove);
      document.addEventListener(
        this.isPc ? "mousemove" : "touchmove",
        this.test
      );
      document.addEventListener(this.isPc ? "mouseup" : "touchend", this.onEnd);
    },
    removeDocumentEvents() {
      document.removeEventListener(
        this.isPc ? "mousemove" : "touchmove",
        this.test
      );
      document.removeEventListener(
        this.isPc ? "mouseup" : "touchend",
        this.onEnd
      );
    },
  },
  unmounted(){
    this.removeDocumentEvents()
  }
};
</script>
<style lang="stylus" scoped>
div:focus {
  outline: none;
}

.by-slider {
  box-sizing: border-box;
  color: rgba(0, 0, 0, 0.65);
  font-size: 14px;
  font-variant: tabular-nums;
  line-height: 1.5715;
  list-style: none;
  font-feature-settings: 'tnum';
  position: relative;
  height: 12px;
  margin: 14px 6px 10px;
  padding: 4px 0;
  cursor: pointer;
  touch-action: none;

  &:hover {
    .by-slider-rail {
      background-color: #e1e1e1;
    }

    .by-slider-track {
      background-color: #69c0ff;
    }
    .by-handle-show{
      transform scale(1)
      opacity:1;
    }
  }

  .by-slider-rail, .by-slider-track {
    position: absolute;
    height: 4px;
    border-radius: 2px;
    transition: background-color 0.3s;
  }

  .by-slider-rail {
    width: 100%;
    background-color: #f5f5f5;
  }

  .by-slider-track {
    background-color: #91d5ff;
  }

  .by-slider-step {
    position: absolute;
    width: 100%;
    height: 4px;
    background: transparent;
  }

  .by-slider-handle {
    position: absolute;
    width: 14px;
    height: 14px;
    margin-top: -5px;
    background-color: #fff;
    border: 2px solid #91d5ff;
    border-radius: 50%;
    box-shadow: 0;
    cursor: pointer;
    box-sizing: border-box;
    transition: border-color 0.3s,opacity .2s,box-shadow 0.6s, transform 0.3s cubic-bezier(0.18, 0.89, 0.32, 1.28);

    &:focus {
      border-color: #46a6ff;
      outline: none;
      box-shadow: 0 0 0 5px rgba(24, 144, 255, 0.2);
      transform scale(1)
      opacity:1;
    }

    &:hover {
      border-color: #46a6ff;
    }
  }
  .by-handle-show{
    opacity:0;
    transform scale(0)
  }

  .by-slider-mark {
    position: absolute;
    top: 14px;
    left: 0;
    width: 100%;
    font-size: 14px;
  }
}

.by-slider-vertical {
  width: 12px;
  height: 100%;
  margin: 6px 10px;
  padding: 0 4px;

  .by-slider-rail {
    width: 4px;
    height: 100%;
  }

  .by-slider-track {
    width: 4px;
  }

  .by-slider-step {
    width: 4px;
    height: 100%;
  }

  .by-slider-handle {
    margin-top: -6px;
    margin-left: -5px;
  }

  .by-slider-mark {
    top: 0;
    left: 12px;
    width: 18px;
    height: 100%;
  }
}

.by-slider-disabled {
  cursor: not-allowed;

  .by-slider-track {
    background-color: rgba(0, 0, 0, 0.25) !important;
  }

  .by-slider-handle {
    background-color: #fff;
    border-color: rgba(0, 0, 0, 0.25) !important;
    box-shadow: none;
    cursor: not-allowed;
  }
}
</style>