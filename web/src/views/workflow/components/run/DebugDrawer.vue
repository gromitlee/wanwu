<template>
  <div class="debug" v-if="visable">
    <div class="config-header">
      <i class="el-icon-close close-icon" @click="preClose"></i>
      <span class="header-name">工作流调试</span>
    </div>
    <div class="form">
      <div class="form-item" v-for="(n, i) in startNode.data.outputs" :key="i">
        <div class="form-item--label">
          <span class="name">{{ n.name }}</span>
          <span class="required" v-if="n.required">*</span>
          <el-tooltip
            class="item"
            effect="dark"
            :content="n.desc || '暂无描述'"
            placement="top-start"
          >
            <i class="el-icon-question desc-icon"></i>
          </el-tooltip>
          <span class="type">{{ n.type }}</span>
          <el-select
            size="mini"
            v-model="n.upLoadType"
            v-if="n.type === 'fileUrl'"
            placeholder="请选择"
            @change="handleUrlChange(i)"
          >
            <el-option :label="'上传文件'" :value="'0'"> </el-option>
            <el-option :label="'输入Url'" :value="'1'"> </el-option>
          </el-select>
        </div>
        <div v-if="n.type !== 'fileUrl'">
          <el-input
            v-if="
              n.type != 'array' ||
              (n.type === 'fileUrl' && n.upLoadType === '1')
            "
            class="form-item--input"
            v-model="n.value.content"
            placeholder="请填写输入参数值"
          ></el-input>
          <codeEditor
            v-else
            style="height: 200px; overflow: auto"
            :value="n.value.content"
            :language="'json'"
            :n="i"
            @handleChange="handleChange"
            :theme="'vs'"
          />
        </div>
        <div v-else style="padding-top: 10px">
          <el-input
            v-if="n.upLoadType === '1'"
            class="form-item--input"
            v-model="n.value.content"
            placeholder="请填写输入参数值"
          ></el-input>
          <Upload
            v-if="n.upLoadType === '0'"
            :ref="'upload' + i"
            :index="i"
            @handleUploadSuccess="handleUploadSuccess"
            @handleDisabled="handleDisabled"
            @handleCancel="handleCancel"
          ></Upload>
        </div>
      </div>

      <div
        class="params-content"
        v-if="startNode.data.settings.staticAuthToken"
      >
        <div class="params-content-item">
          <span>token：</span>
          <span>{{
            startNode.data.settings.staticAuthToken.slice(0, 6) + "******"
          }}</span>
        </div>
      </div>
    </div>
    <div class="btns">
      <el-button
        size="mini"
        type="primary"
        @click="preRun"
        :disabled="runDisabled"
      >
        <i class="el-icon-caret-right"></i>&nbsp;&nbsp;开始运行
      </el-button>
      <br />
      <el-button size="mini" @click="preCancelRun">
        <i class="el-icon-video-pause"></i>&nbsp;&nbsp;取消调试
      </el-button>
      <br />
      <el-button
        v-if="isStream && sessionStatus !== -1"
        size="mini"
        @click="stopSse"
        type="danger"
      >
        <i class="el-icon-video-pause"></i>&nbsp;&nbsp;终止输出
      </el-button>
    </div>

    <div class="result answer-content" v-if="isStream">
      <p class="result-title" v-show="runResponse">运行结果:</p>
      <div
        v-if="showDSBtn(runResponse)"
        class="deepseek"
        @click.prevent.stop="toggle($event)"
      >
        <img src="@/assets/imgs/sikao.png" alt="" />
        {{ thinkText }}
        <i
          v-bind:class="{
            'el-icon-arrow-down': !isOpen,
            'el-icon-arrow-up': isOpen,
          }"
        ></i>
      </div>
      <div v-bind:class="{ 'ds-res': showDSBtn(runResponse) }">
        <div class="answer_content_box" v-html="replaceHTML(runResponse)"></div>
      </div>
    </div>
  </div>
</template>

<script>
import sseMethod from "@/mixins/sseMethod.js";
import codeEditor from "@/views/ArrayEditor/index.vue";
import Upload from "./upload.vue";
import { getQueryString } from "@/utils/util.js";

export default {
  components: { codeEditor, Upload },
  mixins: [sseMethod],
  data() {
    return {
      source: [], // 储存目前上传的接口
      isStream: getQueryString("isStream"),
      visable: false,
      startNode: {},
      nodeStatusObj: {
        success: {
          style: "green",
          label: "成功",
        },
        failed: {
          style: "red",
          label: "失败",
        },
        init: {
          style: "blue",
          label: "等待",
        },
      },
      nodeTypeObj: {
        StartNode: "开始节点",
        ApiNode: "API节点",
        PythonNode: "代码节点",
        EndNode: "结束节点",
      },
      nodeResult: {},
      runDisabled: false,
      thinkText: "",
      isOpen: true,
    };
  },
  created() {},
  methods: {
    // 上传功能新增代码 ************* start  **********************
    handleUploadSuccess(obj) {
      // ** obj
      // ** obj.index 下标 第几个参数
      // ** obj.url 返回的 downloadUrl
      // ** obj.fileId 返回的 fileId
      // ** obj.file 选中的文件信息
      this.startNode.data.outputs[obj.index].value.content = obj.fileId;
    },
    // 上传中触发事件,返回目前是否在上传文件
    handleDisabled(val) {
      // val: true 代表正在上传
      this.runDisabled = val;
    },
    // 储存目前正在上传的接口信息
    handleCancel(obj) {
      this.source.push(obj.source);
    },
    handleUrlChange(i) {
      this.startNode.data.outputs[i].value.content = "";
    },
    // ******************* end  **********************
    preClose() {
      this.visable = false;
    },
    openDialog(workflowSchema) {
      this.runResponse = "";
      if (workflowSchema.nodes) {
        let validate = true;
        let message = "";
        workflowSchema.nodes.forEach((item) => {
          if (
            item.validate &&
            (JSON.parse(item.validate).inputValidate === false ||
              JSON.parse(item.validate).outputValidate === false)
          ) {
            validate = false;
            message = `${item.name}:  ${JSON.parse(item.validate).message}`;
          }
        });
        if (!validate) {
          this.$message.error(message);
          return;
        }
      }
      let startNode = workflowSchema.nodes.filter((n) => {
        /*n.data.outputs.forEach(m=>{
                        m.value.content = '11'
                    })*/
        return n.type === "StartNode";
      })[0];
      startNode.data.outputs.forEach((n) => {
        if (n.type === "fileUrl") {
          n.upLoadType = "0";
        }
      });
      this.startNode = startNode;
      this.visable = true;
    },
    setArray(data) {
      try {
        let arr = JSON.parse(data);
        // arr = arr.map((item)=>{
        //   return String(item)
        // })
        return arr;
      } catch (error) {
        return [data];
      }
    },
    handleChange(value, i) {
      this.startNode.data.outputs[i].value.content = value;
    },
    preRun() {
      let params = {};
      this.startNode.data.outputs.forEach((n) => {
        params[n.name] =
          n.type == "array" && n.value.content
            ? this.setArray(n.value.content)
            : n.value.content;
      });

      this.runDisabled = true;
      this.$emit("doDebug", params);
    },
    setDebugResult(data) {
      this.nodeResult = data;
      this.runDisabled = false;
    },
    preCancelRun() {
      this.stopSse();
      if (this.source.length > 0) {
        this.source.forEach((item) => {
          item.cancel();
        });
        this.source = [];
      }
      this.$emit("cancelRun");
    },
    stopSse() {
      if (this.sessionStatus !== -1) {
        this.setStoreSessionStatus(-1);
        this.sseOnCloseCallBack();
        this._print && this._print.stop();
      }
    },
    toggle(event) {
      // console.log(event.target.className);
      const name = event.target.className;
      if (
        name === "deepseek" ||
        name === "el-icon-arrow-up" ||
        name === "el-icon-arrow-down"
      ) {
        // this.showDs = !this.showDs;
        this.isOpen = !this.isOpen;

        let elm = null;
        if (name === "el-icon-arrow-up" || name === "el-icon-arrow-down") {
          elm = event.target.parentNode.parentNode
            .getElementsByClassName("answer_content_box")[0]
            .getElementsByTagName("section")[0];
        } else {
          elm = event.target.parentNode
            .getElementsByClassName("answer_content_box")[0]
            .getElementsByTagName("section")[0];
        }
        if (!this.isOpen) {
          elm.className = "hideDs";
        } else {
          elm.className = "";
        }
      }
    },
    showDSBtn(data) {
      const pattern = /<\/?think>/;
      const matches = data.match(pattern);
      if (!matches) {
        return false;
      }
      return true;
    },
    replaceHTML(data) {
      let _data = data;
      var a = new RegExp("<think>");
      var b = new RegExp("</think>");
      if (b.test(data)) {
        this.thinkText = "已深度思考";
      } else {
        if (this.sessionStatus == -1) {
          if (a.test(data) && !b.test(data)) {
            this.thinkText = "思考已停止";
          }
        } else {
          this.thinkText = "思考中...";
        }
      }

      // 如果没有返回前缀，则补上
      if (b.test(data) && !a.test(data)) {
        _data = "<think>" + data;
      }
      _data = _data.replace(/<think>\s*\n/g, "<think>");

      return _data.replace(/think>/g, "section>");
    },
    scrollBottom() {
      var a = document.getElementsByClassName("answer-content")[0];
      if (a) {
        a.scrollTop = a.scrollHeight;
      }
    },
  },
};
</script>

<style lang="scss" scoped>
.debug {
  position: absolute;
  display: flex;
  flex-direction: column;
  width: 450px;
  height: 100%;
  top: 0;
  bottom: 0;
  right: 0;
  padding: 20px;
  border-left: 1px solid #ddd;
  background-color: #fff;
}
.config-header {
  position: relative;
  margin-top: 40px;
  .close-icon {
    position: absolute;
    right: 0;
    top: 0;
    font-size: 16px;
    padding: 2px;
    cursor: pointer;
  }
  .header-icon {
    vertical-align: middle;
    width: 26px;
    margin-right: 10px;
    object-fit: contain;
  }
  .header-name {
  }
  .desc {
    padding: 10px 0;
    border-bottom: 1px solid #e8e9eb;
  }
}
.form {
  margin-top: 20px;
  .form-item {
    margin: 20px 0;
    .form-item--label {
      margin-bottom: 10px;
      .name {
        font-size: 12px;
      }
      .required {
        color: #e60001;
        margin-left: 3px;
      }
      .desc-icon {
        margin: 0 5px;
        font-size: 16px;
      }
      .type {
        padding: 0 4px;
        text-align: center;
        color: #5c5f66;
        border-radius: 4px;
        background: #e8e9eb;
        font-size: 12px;
        font-weight: 400;
        line-height: 20px;
      }
    }
    .form-item--input {
    }
  }
}
/deep/.el-select {
  float: right;
  width: 99px;

  .el-input .el-input__inner {
    height: 28px;
    font-size: 12px;
  }
}
.run-bt {
  margin-top: 10px;
  width: 100%;
  i {
    font-size: 20px;
  }
}
.btns {
  display: flex;
  flex-direction: column;
  margin-top: 20px;
  .el-button {
    margin-left: 0;
  }
}
.cancel-bt {
  margin-top: 10px;
  width: 100%;
  margin-left: 0;
  i {
    font-size: 20px;
  }
}
.result {
  height: calc(100% - 320px);
  overflow: auto;
  padding: 10px 0;
  .result-title {
    font-size: 16px;
    font-weight: 600;
    margin-bottom: 10px;
  }
  .result-content {
    margin-top: 30px;
    .result-item {
      margin: 10px 0;
      div {
      }
    }
  }
  .green {
    color: #67c23a;
  }
  .red {
    color: #e60001;
  }
}

.params-content {
  margin-top: 4px;
  .params-content-item {
    padding: 4px 0;
    display: flex;
    background-color: #e8e9eb;
    border-radius: 4px;
    color: #5c5f66;
    span {
      margin-left: 5px;
      font-size: 12px;
      display: inline-block;
      vertical-align: middle;
    }
  }
}

.answer-content {
  // width: calc(100% - 30px);
  position: relative;
  margin-left: 14px;
  color: #333;
  // bottom: 40px;
  white-space: normal !important;
  li {
    display: revert !important;
  }
  img {
    max-width: 100%;
  }
  .answer_content_box {
    // line-height: 3;
    // white-space: pre-wrap;

    /deep/ p {
      margin-bottom: 20px;
    }
    /deep/li {
      margin: 0 0 10px 0;
      p {
        margin-bottom: 0;
      }
      // line-height: 1;
    }
    & /deep/ :last-child {
      margin-bottom: 0 !important;
    }
  }
}
.ds-res {
  /deep/ section {
    color: #8b8b8b;
    position: relative;
    font-size: 13px;
    padding-left: 15px;
    margin-bottom: 10px;
    font-weight: 700;
    // white-space: pre-wrap;
    * {
      font-size: 13px;
      // white-space: pre-wrap;
    }
  }
  /deep/ section::before {
    content: "";
    position: absolute;
    height: 100%;
    width: 2px;
    background: #ddd;
    left: 1px;
  }
  /deep/.hideDs {
    display: none;
  }
}

.deepseek {
  position: relative;
  display: flex;
  align-items: center;
  justify-content: space-between;
  font-size: 13px;
  color: rgb(1, 3, 56);
  font-weight: bold;
  background: #ebeff8;
  cursor: pointer;
  padding: 0 10px;
  padding-left: 15px;
  border-radius: 5px;
  margin-bottom: 5px;
  font-size: 12px;
  width: 150px;
  img {
    width: 12px;
    height: 12px;
  }
}
</style>
