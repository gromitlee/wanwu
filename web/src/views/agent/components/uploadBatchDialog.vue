<template>
    <div class="fileUpload">
        <el-dialog
                custom-class="upload-dialog"
                :visible.sync="dialogVisible"
                width="800px"
                append-to-body
                :before-close="handleClose">
                <div  v-loading="loading" element-loading-background="rgba(255, 255, 255, 0.5)">
                    <div class="dialog-body">
                        <p class="upload-title">{{$t('common.fileUpload.uploadFile')}}</p>
                        <el-upload
                                :class="['upload-box']"
                                drag
                                action=""
                                :show-file-list="false"
                                :auto-upload="false"
                                :limit="2"
                                :accept="tipsArr"
                                :file-list="fileList"
                                :on-change="uploadOnChange"
                                >
                                <div v-if="fileUrl" class="echo-img-box">
                                    <div class="echo-img">
                                        <img v-if="fileType === 'image/*'"  :src="'/user/api'+fileUrl" />
                                        <video v-if="fileType === 'video/*'" id="video" muted loop playsinline>
                                            <source :src = 'fileUrl' type="video/mp4">
                                            {{$t('common.fileUpload.videoTips')}}
                                        </video>
                                        <audio v-if="fileType === 'audio/*'" id="audio" controls>
                                            <source :src="fileUrl" type="video/mp3">
                                            <source :src="fileUrl" type="audio/ogg">
                                            <source :src="fileUrl" type="audio/mpeg">
                                            {{$t('common.fileUpload.audioTips')}}
                                        </audio>
                                        <div v-if="fileType === 'doc/*'" class="docFile">
                                            <img :src="require('@/assets/imgs/fileicon.png')" />
                                        </div>
                                        <p>文件名称: {{fileList[0]['name']}}</p>
                                        <p>文件大小: {{fileList[0]['size'] > 1024 ? (fileList[0]['size'] / (1024 * 1024 )).toFixed(2) + ' MB' : fileList[0]['size'] + ' bytes' }}</p>
                                    </div>
                                    <!--<i  class="el-icon-close" @click.stop="clearFile"></i>-->
                                    <div class="tips">
                                        <el-progress
                                            :percentage="file.percentage"
                                            v-if="file.percentage !== 100"
                                            :status="file.progressStatus"
                                            max="100"
                                            style="width:360px;margin:0 auto;"
                                        ></el-progress>
                                        <p>{{$t('common.fileUpload.limitTips')}}<span style="color: red"> {{$t('common.fileUpload.click')}} </span>{{$t('common.fileUpload.refreshTips')}}</p>
                                    </div>
                                </div>
                                <div v-else>
                                    <i class="el-icon-upload"></i><p>{{$t('common.fileUpload.uploadClick')}}</p>
                                    <div class="tips">
                                        <p>{{$t('common.fileUpload.typeFileTip1')}}
                                            <span>{{tipsArr}}</span>
                                            {{$t('common.fileUpload.typeFileTip')}}
                                        </p>
                                    </div>
                                </div>
                        </el-upload>
                    </div>
                    <div class="dialog-footer">
                        <el-button type="primary" :disabled="!fileUrl || !(file && file.percentage === 100 )" @click="doBatchUpload">{{$t('common.fileUpload.submitBtn')}}</el-button>
                    </div>
                </div>
        </el-dialog>

    </div>
</template>

<script>
    import { batchUpload,confirmPath } from '@/api/chat';
    import uploadChunk from "@/mixins/uploadChunk";
    export default {
        props:['fileTypeArr','sessionId'],
        mixins: [uploadChunk],
        data(){
            return{
                fileIdList:[],
                fileList:[],
                fileType:'',
                //上传文件弹框
                loading:false,
                dialogVisible:false,
                fileUrl:'',
                //上传文件
                imgConfig:["jpeg", "PNG", "png", "JPG", "jpg",'bmp','webp'],
                audioConfig:['mp3','wav'],
                tipsArr:'',
                tipsObj:{
                    'image/*':['.jpg', '.jpeg', '.png','.webp'],
                    'audio/*':['.wav', '.mp3'],
                    'doc/*':['.txt','.csv','.xlsx','.docx','.html','.pptx','.pdf']
                },
                chunkFileName:'',
                fileInfo:null
            }
        },
        watch:{
            fileTypeArr(val,oldVal) {
                this.setFileType(val)
            }
        },
        created(){
            this.sessionId = this.sessionId || this.$route.query.sessionId
            if(this.fileTypeArr.length){
                this.setFileType(this.fileTypeArr)
            }
        },
        methods:{
            setFileType(fileTypeArr){
                if(fileTypeArr.length){
                    this.tipsArr = ''
                    let tips_arr=[]
                    fileTypeArr.forEach(item=>{
                        tips_arr = tips_arr.concat(this.tipsObj[item])
                    })
                    this.tipsArr = tips_arr.join(', ')
                }
            },
            openDialog(){
                this.dialogVisible = true
            },
            clearFile(){
                this.fileIdList = []
                this.fileList = []
                this.fileType = ''
                this.fileUrl = ''

            },
            handleClose(){
                this.clearFile()
                this.dialogVisible = false
            },
            uploadOnChange(file, fileList) {
                let filename= file.name
                //通过上传的文件名判断文件类型，用于回显
                let fileType = filename.split('.')[filename.split('.').length-1]
                if(["jpeg", "PNG", "png", "JPG", "jpg",'bmp','webp'].includes(fileType)){
                    this.fileType = 'image/*'
                }
                if(["mp3", "wav"].includes(fileType)){
                    this.fileType = 'audio/*'
                }
                if(['txt','csv','xlsx','docx','html','pptx','pdf'].includes(fileType)){
                    this.fileType = 'doc/*'
                }
                this.fileUrl = URL.createObjectURL(file.raw);
                this.fileList = [];
                this.fileList.push(file);
                if(this.fileList.length > 0){
                    this.maxSizeBytes = 0;
                    this.isExpire = true;
                    this.startUpload();
                }
            },
            uploadFile(fileName,oldFileName,fiePath){
                this.fileInfo = {
                    fileName,
                    fileSize:this.fileList[0]['size'],
                    fileUrl:fiePath,
                }
            },
            doBatchUpload(){
                this.$emit('setFileId',this.fileInfo)
                this.$emit('setFile',this.fileList)
                this.handleClose();
            },
            getFileIdList(){
                return this.fileIdList
            },
        }
    }
</script>

<style lang="scss" scoped>
    .upload-dialog {
        .dialog-body{
            padding:0 20px;
            .upload-title{
                text-align: center;
                font-size: 18px;
                margin-bottom: 20px;
            }
            .upload-box{
                height: 190px;
                width: 100% !important;
                background-color: #fff;
                .el-upload-dragger {
                    .el-icon-upload{
                        margin: 46px 0 10px 0 !important;
                        font-size: 32px !important;
                        line-height: 36px!important;
                        color: #384BF7;
                    }
                    .el-upload__text{
                        margin-top: -10px;
                    }
                }
            }

            .echo-img-box{
                background-color: transparent!important;
                .echo-img{
                    img,video{
                        width: auto;
                        height: 80px;
                        margin: 10px auto;
                        border-radius: 4px;
                        background-color: transparent;
                    }
                    audio{
                        width: 300px;
                        height: 54px;
                        margin: 50px auto;
                    }
                }
                .docFile{
                    img{
                        margin:0;
                        width:60px;
                        height: 100px;
                    }
                }
            }
            .tips{
                position: absolute;
                bottom: 16px;
                left: 0;
                right: 0;
                p{
                    color: #9d8d8d!important;
                }
            }
        }
        .dialog-footer{
            text-align: center;
            margin: 30px 0 20px 0;
        }
    }
</style>
