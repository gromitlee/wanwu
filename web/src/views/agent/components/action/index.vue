<template>
    <div class="action">
        <el-row>
            <el-col :span="24" class="left-col">
                <div class="left-col-header rl">
                    <i class="el-icon-arrow-left  back-icon" @click="goBack"></i>
                    <span class="header-title">{{actionId?'编辑':'创建'}} actions</span>
                </div>
                <div class="action-form">
                    <div class="block prompt-box">
                        <p class="block-title required-label rl">API身份认证</p>
                        <div class="rl" @click="preAuthorize">
                            <!--<el-input class="name-input" v-model="basicForm.apiKey" @blur="listenerUpdate" @focus="apiKeyOnFocus" maxlength="10" placeholder="" ></el-input>-->
                            <div class="api-key">{{basicForm.apiKey}}</div>
                            <img class="auth-icon" src="./auth.png" />
                        </div>
                    </div>

                    <div class="block prompt-box">
                        <p class="block-title required-label rl">Schema</p>
                        <div class="rl">
                            <!--<div class="flex" style="margin-bottom: 10px">
                                <el-input class="name-input" v-model="basicForm.name" @blur="listenerUpdate" maxlength="10" placeholder="应用名称" ></el-input>
                                <el-button style="margin-left: 10px" type="primary" size="small">导入</el-button>
                                <el-select style="margin-left: 100px;width: 200px">
                                    <el-option label="模板样例导入" value="json"></el-option>
                                    <el-option label="JSON样例导入" value="json"></el-option>
                                    <el-option label="YAML样例导入" value="json"></el-option>
                                </el-select>
                            </div>-->
                            <div class="flex" style="margin-bottom: 10px">
                                <el-select v-model="basicForm.example" placeholder="选择样例" style="width:100%;" @change="exampleChange">
                                    <!--<el-option label="模板样例导入" value="json"></el-option>-->
                                    <el-option label="JSON样例导入" value="json"></el-option>
                                    <el-option label="YAML样例导入" value="yaml"></el-option>
                                </el-select>
                            </div>
                            <el-input class="schema-textarea" v-model="basicForm.schema" @blur="listenerUpdate"  placeholder="请输入对应API的openapi3.0结构，可以选择示例了解详情。" type="textarea" ></el-input>
                        </div>
                    </div>

                    <div class="block prompt-box">
                        <p class="block-title required-label rl">可用API</p>
                        <div class="api-list">
                            <el-table
                                    :data="apiList"
                                    border
                                    size="mini"
                                    class="api-table"
                                    >
                                <el-table-column
                                        prop="name"
                                        label="Name"
                                        width="180">
                                </el-table-column>
                                <el-table-column
                                        prop="method"
                                        label="Method"
                                        width="180">
                                </el-table-column>
                                <el-table-column
                                        prop="path"
                                        label="Path">
                                </el-table-column>
                            </el-table>
                        </div>
                    </div>


                    <div class="block prompt-box">
                        <p class="block-title  rl">隐私政策</p>
                        <el-input class="name-input" v-model="basicForm.privacy" placeholder="填写API对应的隐私政策url链接" ></el-input>
                    </div>
                    <!--<div class="block prompt-box">
                        <p class="block-title ">actions获取结果</p>
                        <div class="rl">
                            <el-table
                                    :data="tableData"
                                    border
                                    style="width: 100%">
                                <el-table-column
                                        prop="name"
                                        label="名称"
                                        width="180">
                                </el-table-column>
                                <el-table-column
                                        prop="type"
                                        label="方法"
                                        width="180">
                                </el-table-column>
                                <el-table-column
                                        prop="url"
                                        label="路径">
                                </el-table-column>
                                <el-table-column
                                        fixed="right"
                                        label="操作"
                                        width="100">
                                    <template slot-scope="scope">
                                        <el-button type="text" size="small">测试</el-button>
                                    </template>
                                </el-table-column>
                            </el-table>
                        </div>
                    </div>

                    <div class="block recommend-box">
                        <p class="block-title">隐私条款</p>
                        <el-input class="desc-input " v-model="basicForm.description" placeholder="填写URL"></el-input>
                    </div>-->
                </div>
            </el-col>

        </el-row>

        <!--认证弹窗-->
        <el-dialog
            title="认证"
            :visible.sync="dialogVisible"
            width="600px"
            append-to-body
            :close-on-click-modal="false"
            @close="beforeAuthFormClose"
            >
            <div class="action-form">
                <el-form :rules="rules" ref="form" :inline="false" :model="authForm">
                    <el-form-item label="认证类型">
                        <el-radio-group v-model="authForm.type">
                            <el-radio label="none">None</el-radio>
                            <el-radio label="apiKey">API Key</el-radio>
                            <!--<el-radio label="3">OAuth</el-radio>-->
                        </el-radio-group>
                    </el-form-item>
                    <!--API Key-->
                    <div v-if="authForm.type === 'apiKey'">
                        <el-form-item label="API key" prop="apiKey">
                            <el-input class="desc-input " v-model="authForm.apiKey" placeholder="API key" clearable></el-input>
                        </el-form-item>
                        <el-form-item label="Auth类型">
                            <el-radio-group v-model="authForm.authType">
                                <!--<el-radio label="1">Basic</el-radio>
                                <el-radio label="2">Bearer</el-radio>-->
                                <el-radio label="custom">Custom</el-radio>
                            </el-radio-group>
                        </el-form-item>
                        <el-form-item label="Custom Header Name" prop="customHeaderName">
                            <el-input class="desc-input " v-model="authForm.customHeaderName" placeholder="Custom Header Name" clearable></el-input>
                        </el-form-item>
                    </div>
                    <!--OAuth-->
                    <!--<div v-if="authForm.type === '3'">
                        <el-form-item label="Client ID">
                            <el-input class="desc-input " v-model="basicForm.description" placeholder="Client ID"></el-input>
                        </el-form-item>
                        <el-form-item label="Client Secret">
                            <el-input class="desc-input " v-model="basicForm.description" placeholder="Client Secret"></el-input>
                        </el-form-item>
                        <el-form-item label="Authentication URL">
                            <el-input class="desc-input " v-model="basicForm.description" placeholder="Authentication URL"></el-input>
                        </el-form-item>
                        <el-form-item label="Token URL">
                            <el-input class="desc-input " v-model="basicForm.description" placeholder="Token URL"></el-input>
                        </el-form-item>
                        <el-form-item label="Scope">
                            <el-input class="desc-input " v-model="basicForm.description" placeholder="Scope"></el-input>
                        </el-form-item>
                        <el-form-item label="Token Exchange Method">
                            <el-radio-group v-model="authForm.authType">
                                <el-radio label="1">Default（POST request）</el-radio>
                                <el-radio label="2">Bearer</el-radio>
                            </el-radio-group>
                        </el-form-item>
                    </div>-->

                </el-form>

            </div>

            <span slot="footer" class="dialog-footer">
                <el-button @click="beforeAuthFormClose">取 消</el-button>
                <el-button type="primary" @click="listenerApiKey">确 定</el-button>
            </span>
        </el-dialog>

    </div>
</template>

<script>
    import { batchUpload,createApp,getAppDetail,updateApp } from '@/api/chat'
    import { addAction,updateAction,deleteAction,getActionDetail } from '@/api/cubm'
    import { addActionInfo,editActionInfo,getActionInfo } from "@/api/agent";
    import {schemaConfig} from './schema.conf'

    export default {
        props:['assistantId'],
        data(){
            return{
                actionId:'',
                apiList:[],//{name: "getWeatherNow", method: "GET", path: "/weather/now.json"}
                basicForm:{
                    example:'',
                    apiKey:'None',
                    schema:'',
                    privacy:''
                },
                //认证表单
                dialogVisible: false,
                authForm:{
                    type:'none',
                    authType:'custom',
                    apiKey:'',
                    customHeaderName:'',
                },
                rules: {
                    apiKey: [{required: true, message: '请输入', trigger: 'blur'}],
                    customHeaderName: [{required: true, message: '请输入', trigger: 'blur'}],
                },
                schemaConfig: schemaConfig

            }
        },
        methods:{
            setActionData(data){
                this.initData()
                switch (data.type){
                    case 'create':
                        break;
                    case 'update':
                        this.actionId = data.actionId
                        this.getActionDetail()
                        break;
                }
            },
            initData(){
                this.actionId = ''
                this.apiList = []
                this.basicForm ={
                    example:'',
                    apiKey:'None',
                    schema:'',
                    privacy:''
                }
                this.authForm ={
                    type:'none',
                    apiKey:'',
                    authType:'custom',
                    customHeaderName:'',
                }
            },
            async getActionDetail(){
                let res = await getActionInfo({actionId:this.actionId})
                if(res.code === 0){
                    let data = res.data
                    this.basicForm.schema = data.schema
                    this.authForm.type = data.apiAuth.type
                    this.authForm.apiKey = data.apiAuth.apiKey
                    this.authForm.customHeaderName = data.apiAuth.customHeaderName
                    this.apiList = data.list
                    this.matchApiType()
                }
            },
            matchApiType(){
                if(this.authForm.type === 'apiKey' && this.authForm.apiKey && this.authForm.customHeaderName){
                    this.basicForm.apiKey = 'Api Key'
                }
                if(this.authForm.type === 'none'){
                    this.basicForm.apiKey = 'None'
                }
            },
            listenerUpdate(){
                if(!(this.basicForm.schema && this.authForm.type)){
                    return
                }
                if(this.actionId){
                    this.doUpdateAction()
                }else{
                    this.doCreateAction()
                }
            },
            async doCreateAction(){
                let res = await addActionInfo(this.makeParams())
                if (res.code === 0) {
                    this.actionId = res.data.actionId
                    this.getActionDetail()
                }
            },
            async doUpdateAction(){
                let res = await editActionInfo(this.makeParams({actionId:this.actionId}))
                if(res.code === 0){
                    this.getActionDetail()
                }
            },
            makeParams(otherParams){
                let params = {
                    assistantId: this.assistantId,
                    schema: this.basicForm.schema,
                    apiAuth:{
                        type:this.authForm.type
                    },
                    ...otherParams
                }
                switch (this.authForm.type) {
                    case 'none':
                        break;
                    case 'apiKey':
                        params.apiAuth.authType = this.authForm.authType
                        params.apiAuth.apiKey = this.authForm.apiKey
                        params.apiAuth.customHeaderName = this.authForm.customHeaderName
                        break;
                }
                return params
            },

            exampleChange(value){
                this.basicForm.schema = this.schemaConfig[value]
                this.listenerUpdate()
            },
            beforeAuthFormClose(){
                this.$refs['form'].clearValidate()
                this.dialogVisible = false
            },
            listenerApiKey(){
                this.$refs.form.validate(async (valid) => {
                    if (!valid) return;
                    this.matchApiType()
                    this.listenerUpdate()
                    this.dialogVisible = false
                })
            },
            preAuthorize(){
              this.dialogVisible =true
            },
            goBack(){
                this.listenerUpdate()
                this.$emit('closeAction')
                //this.$router.go(-1)
            },
        },

    }
</script>

<style lang="scss" scoped>
/deep/.el-radio__input.is-checked .el-radio__inner{
    border-color: #D33A3A !important;
    background: transparent !important;
}
/deep/.el-radio__input.is-checked .el-radio__inner::after{
    background: #eb0a0b !important;
    width: 7px !important;
    height: 7px !important;
}
::selection {
    color: #1a2029;
    background: #c8deff;
}
.left-col{
    // background-color: #fafafa;
    overflow: auto;
    height: 100%;
    .left-col-header{
        width: 100%;
        padding: 30px 40px;
        text-align: center;
        .back-icon{
            position: absolute;
            left: 35px;
            font-size: 14px;
            cursor: pointer;
            border-radius: 20px;
            border: 1px solid #e1e1e1;
            padding: 6px;
            color: #444;
            &:hover{
                font-weight: bold;
            }
        }
        .header-title{
            font-size: 18px;
            font-weight: bold;
            color: #303133;
        }
        .bt-box{
            position: absolute;
            width: 160px;
            height: 30px;
            right: 20px;
            top:0;
            bottom: 0;
            margin: auto;
            .del-bt{margin-left: 10px;}
        }
    }
    .action-form{
        padding: 0 40px;
        /deep/.schema-textarea{
            .el-textarea__inner{
                height: 200px!important;
            }
        }
        .api-key{
            background-color: transparent !important;
            border: 1px solid #d3d7dd !important;
            padding:0 15px;
            -webkit-appearance: none;
            background-image: none;
            border-radius: 4px;
            box-sizing: border-box;
            color: #606266;
            display: inline-block;
            height: 40px;
            line-height: 40px;
            outline: 0;
            transition: border-color .2s cubic-bezier(.645,.045,.355,1);
            width: 100%;
        }
        .auth-icon{
            position: absolute;
            right: 0;
            height: 39px;
            top: 0;
            cursor: pointer;
            border-left: 1px solid #d3d7dd;
            padding: 7px 9px;
        }
    }

}
.right-col{
    height: 100%;
    // background-color: #f6f7f9;
    .right-title{
        line-height: 84px;
        font-size: 18px;
        font-weight: bold;
        text-align: center;
        color: #303133;
    }
    .smart-center{
        min-width: 0;
        height: calc(100% - 84px);
        flex:1;
        background-size: 100% 100%;
        position: relative;
    }
}
/*通用*/
.action{
    position: relative;
    height:100%;
    /deep/.el-input__inner, /deep/.el-textarea__inner {
        background-color: transparent !important;
        border: 1px solid #d3d7dd!important;
        font-family: 'Microsoft YaHei', Arial, sans-serif;
        padding: 15px;
    }
    .flex{
        width: 100%;
        display: flex;
        justify-content: space-between;
    }
    .block{
        margin-bottom: 20px;
        .block-title{
            line-height: 30px;
            font-size: 15px;
            font-weight: bold;
        }
        .required-label::after{
            content:'*';
            position: absolute;
            color: #eb0a0b;
            font-size: 20px;
            margin-left: 4px;
        }
        .block-tip{
            color: #919eac;
        }
    }
    .el-input__count {
        color: #909399;
        // background: #fafafa;
        position: absolute;
        font-size: 12px;
        bottom: 5px;
        right: 10px;
    }
}
    .action-form /deep/ .el-form-item__label{
        display: block;
        width: 100%;
        text-align: left;
    }
    .api-list{
        .api-table /deep/ .el-table__body  tr td,
        .api-table /deep/ .el-table__header  tr th,
        .api-table /deep/ .el-table__body tr:hover > td{
            background-color: transparent !important;
        }
    }

</style>
