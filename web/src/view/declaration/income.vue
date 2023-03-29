<template>
  <div>
    <el-form :model="form"  label-width="100px">
  <el-form-item label="申请人" prop="applicant" >
    <el-col :span="6">
      <el-input v-model="form.applicant"></el-input>
    </el-col>
  </el-form-item>
  <el-form-item label="申请人ID" prop="applicantId" >
  <el-col :span="6">
    <el-input v-model="form.applicantId" type="number"></el-input>
  </el-col>
</el-form-item>
  <el-form-item label="审批人ID" prop="approverId">
    <el-select v-model="form.approverId" placeholder="请选择活动区域" @change="getApprover()" >
      <el-option v-for="user in declarationUser" :label="`${user.ID}`" :value="`${user.ID}`"></el-option>
    </el-select>
  </el-form-item>
  <el-form-item label="审批人姓名" prop="approverId"> {{ form.approver }}
  </el-form-item>
  <el-form-item label="申报时间" required>
    <el-col :span="11">
      <el-form-item prop="time">
        <el-date-picker type="date" placeholder="选择日期" v-model="form.time" style="width: 100%;"></el-date-picker>
      </el-form-item>
    </el-col>
  </el-form-item>
  <el-form-item label="金额" prop="money">
    <el-col :span="4">
      <el-input v-model="form.money" type="number"></el-input>
  </el-col>
</el-form-item>
<el-form-item label="发票" >
  <el-upload
  class="avatar-uploader"
  :action="`${path}/fileUploadAndDownload/upload`"
  :headers="{ 'x-token': userStore.token }"
  :show-file-list="false"
  :multiple="false"
  :on-success="handleImageSuccess"
  :before-upload="beforeImageUpload"
  >
  <!-- <CustomPic  v-if="imageUrl" pic-type="file" :pic-src="imageUrl" /> -->
  <img v-if="imageUrl" :src="`${path}/${imageUrl}`" class="avatar">
  <el-icon v-else class="avatar-uploader-icon"><Plus /></el-icon>
  </el-upload>
</el-form-item>
  <el-form-item label="申报理由" prop="reason">
    <el-input type="textarea" v-model="form.reason"></el-input>
  </el-form-item>
  <el-form-item>
    <el-button type="primary" @click="onSubmit()">{{ buttonLable }}</el-button>
    <el-button @click="onSubmit()">重置</el-button>
  </el-form-item>
</el-form>
  </div>
</template>
<style scoped>
.avatar-uploader .avatar {
width: 250px;
display: block;
}
</style>

<style>
.avatar-uploader .el-upload {
border: 1px dashed var(--el-border-color);
border-radius: 6px;
cursor: pointer;
position: relative;
overflow: hidden;
transition: var(--el-transition-duration-fast);
}

.avatar-uploader .el-upload:hover {
border-color: var(--el-color-primary);
}

.el-icon.avatar-uploader-icon {
font-size: 28px;
color: #8c939d;
width: 178px;
height: 178px;
text-align: center;
}
avatar-uploader .el-upload {
border: 1px dashed var(--el-border-color);
border-radius: 6px;
cursor: pointer;
position: relative;
overflow: hidden;
transition: var(--el-transition-duration-fast);
}

.avatar-uploader .el-upload:hover {
border-color: var(--el-color-primary);
}

.el-icon.avatar-uploader-icon {
font-size: 28px;
color: #8c939d;
width: 178px;
height: 178px;
text-align: center;
}
</style>


<script  setup>
import { ref } from 'vue'
import { ElMessage } from 'element-plus'
import { reactive } from 'vue'
import { useUserStore } from '@/pinia/modules/user'
import {Plus} from '@element-plus/icons-vue'
import {getUsersByAuthorityId} from '@/api/user'
import {createBill,updateBill,getBillList,deleteBill} from '@/api/bill'
import { fromPairs } from 'lodash'
import { useRoute } from 'vue-router'
import {  watch } from 'vue'

// do not use same name with ref
const imageUrl = ref('')

const userStore = useUserStore()
const router = useRoute()
const path = ref(import.meta.env.VITE_BASE_API)
console.log("userinfo##",userStore.userInfo.ID)

const buttonLable = ref('立即创建')
const declarationUser = ref([])
const tempDeclarationUser = ref([])

const form = ref({
  applicant: '',
  approverId: 0,
  time: '',
  voucherImgUrl: '',
  reason: '',
  money: 0,
  applicantId: 0,
  approver: '',

})

const init  = ()=> {
  console.log(router.query)
  {
    form.value.applicant = router.query.applicant
    form.value.applicantId = router.query.applicantId
    form.value.time = router.query.time
    form.value.approver = router.query.approver
    form.value.money = router.query.money
    form.value.voucherImgUrl = router.query.voucherImgUrl
    imageUrl.value = router.query.voucherImgUrl
    form.value.reason = router.query.reason
    form.value.type = router.query.type
    console.log(router.query.ID)
    if(typeof(router.query.ID) != "undefined" ){
      buttonLable.value = '立即修改'
    }
  }
}
init()
const onSubmit = async() => {
  {
    form.value.applicantId = Number(form.value.applicantId)
    form.value.voucherImgUrl = imageUrl.value
    form.value.type = 1
    form.value.money = form.value.money.toString()
    form.value.approverId = Number(form.value.approverId)
    form.value.initiator = userStore.userInfo.userName
    form.value.initiatorId = Number(userStore.userInfo.ID)
  }
  if (buttonLable.value == '立即修改'){
    const data = await updateBill(form.value)
    console.log(data)
  }else{
    const data = await createBill(form.value)
    console.log(data)
   }
    
} 
const beforeImageUpload = (file) => {
}

const handleImageSuccess = (res) => {
const { data } = res
console.log(data)
if (data.file) {
  imageUrl.value = data.file.url
  console.log("###",data.file.url)
}
}

watch(form.approverId,(newVal,oldVal)=>{
  console.log(newVal,oldVal)
})

const getDeclationUsers = async()=>{
const data = await getUsersByAuthorityId({authority_id: "2"})
if (data.code === 0){
  declarationUser.value = data.data
}
console.log(data)
}
getDeclationUsers()

const getApprover = () =>{
    for (const index in declarationUser.value){
      console.log(declarationUser.value[index].ID)
      console.log(form.value.approverId)
        if (declarationUser.value[index].ID == form.value.approverId){
            form.value.approver = declarationUser.value[index].userName
            console.log(form.value.approver)
            return 
        }
        
    }
}

</script>

