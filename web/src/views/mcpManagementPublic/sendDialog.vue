<template>
  <div class="add-dialog">
    <el-dialog
      title="添加MCP服务"
      :visible.sync="dialogVisible"
      width="50%"
      :show-close="false"
      :close-on-click-modal="false"
    >
      <div>
        <el-form
          :model="ruleForm"
          :rules="rules"
          ref="ruleForm"
          label-width="130px"
        >
          <el-form-item label="MCP名称">
            <div>{{detail.name}}</div>
          </el-form-item>
          <el-form-item label="服务来源">
            <div>{{detail.from}}</div>
          </el-form-item>
          <el-form-item label="功能描述" class="description-text">
            <div>{{detail.desc}}</div>
          </el-form-item>
          <el-form-item label="MCP ServerURL" prop="serverUrl">
            <el-input v-model="ruleForm.serverUrl"></el-input>
          </el-form-item>
          <el-form-item label=" " style="text-align: right">
            <el-button
              type="primary"
              size="mini"
              @click="handleTools"
              :disabled="isGetMCP"
              >获取MCP工具</el-button
            >
          </el-form-item>
        </el-form>
        <el-divider v-if="mcpList.length > 0"></el-divider>
        <ul class="mcpList" v-if="mcpList.length > 0">
          <li v-for="(item, index) in mcpList" :key="index">{{ item.name }}</li>
        </ul>
      </div>
      <span slot="footer" class="dialog-footer">
        <el-button
          type="primary"
          size="mini"
          :disabled="mcpList.length === 0"
          @click="submitForm('ruleForm')"
        >
          确定发送
        </el-button>
        <el-button @click="handleCancel" size="mini">取 消</el-button>
      </span>
    </el-dialog>
  </div>
</template>
<script>
import { getTools, setCreate } from "@/api/mcp.js";
import { isValidURL } from "@/utils/util";

export default {
  props: ["dialogVisible","detail"],
  data() {
    const validateUrl = (rule, value, callback) => {
      if (!isValidURL(value)) {
        callback(new Error("请再次检查Server Url格式"));
      } else {
        callback();
      }
    };
    return {
      mcpList: [],
      ruleForm: {
        serverUrl: "",
      },
      rules: {
        serverUrl: [
          {
            required: true,
            message: "请输入服务Server Url",
            trigger: "blur",
          },
          { validator: validateUrl, trigger: "blur" },
        ],
      },
    };
  },
  methods: {
    handleCancel() {
      this.clearForm()
      this.$emit("handleClose", false);
    },
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          const params = {
            name: this.detail.name,
            from: this.detail.from,
            sseUrl: this.ruleForm.serverUrl,
            desc: this.detail.desc,
            mcpSquareId: this.detail.mcpSquareId
          }
          setCreate(params).then((res) => {
            if(res.code === 0){
              this.$message.success("发布成功")
              this.clearForm()
              this.$refs["ruleForm"].clearValidate();
              this.handleCancel();
              // 更新发送按钮状态
              this.$emit("getIsCanSendStatus");
            }
          })
        }
      })
    },
    clearForm(){
      this.ruleForm.serverUrl = ''
      this.mcpList = []
    },
    handleTools() {
      this.$refs['ruleForm'].validate((valid) => {
        if (valid) {
          getTools({
            serverUrl: this.ruleForm.serverUrl,
          }).then((res) => {
            this.mcpList = res.data.tools || []
          })
        }
      });
    },
  },
  computed: {
    isGetMCP() {
      if (!isValidURL(this.ruleForm.serverUrl)) {
        return true;
      } else {
        return false;
      }
    },
  },
};
</script>
<style lang="scss" scoped>
.add-dialog {
  .el-button.is-disabled {
    &:active {
      background: transparent !important;
      border-color: #ebeef5 !important;
    }
  }
  .mcpList {
    list-style: none;
    li {
      padding: 10px;
      border: 1px solid #ddd;
      border-radius: 5px;
      margin-bottom: 10px;
      background: #fff;
    }
  }
  .description-text .el-form-item__content{
    line-height: 24px!important;
    padding: 10px 0;
  }
}
</style>