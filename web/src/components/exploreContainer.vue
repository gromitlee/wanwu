<template>
  <el-container class="explore-container">
    <el-aside width="220px"v-if="showAside">
        <slot name="aside">
            <div class="aside-title">
                <span class="title" v-if="!isButton">{{ asideTitle }}</span>
                <div v-else class="title">
                  <el-button type="primary" @click="handleClick" size="mini" ><span class="el-icon-circle-plus-outline addIcon"></span>{{asideTitle}}</el-button>
                </div>
            </div>
        <div class="aside-content">
          <slot name="aside-content"></slot>
        </div>
        </slot>
    </el-aside>
    <el-main>
        <div class="page-wrapper" style="height:100%;overflow:hidden;">
          <slot name="main-content"></slot>
        </div>
    </el-main>
  </el-container>
</template>

<script>
export default {
  name:'commonLayout',
  props:{
    asideTitle: {
      type: String,
      default: ''
    },
    isButton:{
      type:Boolean,
      default:false
    },
    showAside:{
      type:Boolean,
      default:true
    }
  },
  data() {
    return {
    }
  },
  methods:{
    handleClick(){
      this.$emit('handleBtnClick')
    }
  }
}
</script>
<style lang="scss" scoped>
.explore-container{
    width:100%;
    height:100%;
    .el-aside{
      background:#fff;
      border-radius:10px 0 10px 10px;
      .aside-title{
        height:60px;
        padding:12px 15px 0;
        .title{
          display: block;
          font-weight:800;
          font-size: 14px;
          width:100%;
          height: 100%;
          text-align: center;
          line-height: 48px;
          border-bottom: 1px solid #EBEBEB;
          .addIcon{
            margin-right: 8px;
          }
        }
      }
      .aside-content{
        height: calc(100% -60px);
      }
    }
    .el-main{
      padding: 0;
      border-radius:8px 8px 8px 8px;
      /*border-left: 0.5px solid #e6e6e6;*/
    }
}
</style>