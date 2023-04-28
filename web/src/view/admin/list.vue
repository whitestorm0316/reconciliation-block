<template>
  <el-form ref="searchForm" :inline="true" :model="search">
          <el-form-item label="">
            <el-input v-model="search.keyword" class="keyword" placeholder="请输入文件名或备注" />
          </el-form-item>

          <el-form-item>
            <el-button type="primary" icon="search" @click="getTableData">查询</el-button>
          </el-form-item>
        </el-form>

  <el-table :data="tableData">
        <el-table-column align="left" label="预览" width="100">
          <template #default="scope">
            <CustomPic pic-type="file" :pic-src="scope.row.voucherImgUrl" />
          </template>
        </el-table-column>
        <el-table-column align="left" label="日期" prop="time" width="180">
          <template #default="scope">
            <div>{{ formatDate(scope.row.time) }}</div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="申请人名字" prop="applicant" width="180">
          <template #default="scope">
            <div>{{scope.row.applicant}}</div>
          </template>
        </el-table-column>
        <el-table-column align="left" label="申请人ID" prop="applicantId" min-width="100" />
        <el-table-column align="left" label="审批人" prop="approver" min-width="100" />
        <el-table-column align="left" label="类型" prop="type" min-width="100">
            <template #default="scope">
            <el-tag
              :type="scope.row.type === 1 ? 'success' : 'danger'"
              disable-transitions
            >{{ formatType(scope.row.type) }}
            </el-tag>
          </template>
     
        </el-table-column>
        <el-table-column align="left" label="审批员审核状态" prop="approvalState" min-width="100">
            <template #default="scope">
           {{ formatState(scope.row.approvalState) }}
          </template>
        </el-table-column>
        <el-table-column align="left" label="管理员审核状态" prop="approvalState" min-width="100">
            <template #default="scope">
           {{ formatState(scope.row.ConfirmState) }}
          </template>
        </el-table-column>


        <el-table-column align="left" label="是否归档" prop="acchive" min-width="100" >
          <template #default="scope">
           {{ formatArchive(scope.row.archive)}}
          </template>
        </el-table-column>

        <el-table-column align="left" label="原因" prop="reason" min-width="100" />
        <el-table-column align="left" label="操作" width="160">
          <template #default="scope">
            <el-button icon="InfoFilled" type="primary" link @click="updateBill(scope.row)">详情</el-button>
            <el-button icon="delete" type="primary" link @click="deleteBillFunc(scope.row)">删除</el-button>
          </template>
        </el-table-column>
      </el-table>
      <el-pagination
          background
          v-model:current-page="page"
          @current-change="getTableData"
          layout="prev, pager, next"
          :total="total">
      </el-pagination>
</template>

<script setup>
import { getFileList, deleteFile, editFileName } from '@/api/fileUploadAndDownload'
import { downloadImage } from '@/utils/downloadImg'
import CustomPic from '@/components/customPic/index.vue'
import UploadImage from '@/components/upload/image.vue'
import UploadCommon from '@/components/upload/common.vue'
import { formatDate } from '@/utils/format'
import WarningBar from '@/components/warningBar/warningBar.vue'
import { useUserStore } from '@/pinia/modules/user'
import {getBillList,deleteBill} from '@/api/bill'
import { useRouter } from "vue-router"



import { ref } from 'vue'
import { ElMessage, ElMessageBox } from 'element-plus'


const userStore = useUserStore()

const router = useRouter()

const path = ref(import.meta.env.VITE_BASE_API)

const imageUrl = ref('')
const imageCommon = ref('')

const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const search = ref({})
const tableData = ref([])

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  console.log(page)
  tableData.value = []

  const table = await getBillList({ page: page.value, pageSize: pageSize.value, keyword: userStore.userInfo.ID.toString(),keyType: "adminId" })
  if (table.code === 0) {
    let tempTotal = 0
    for (const index in table.data.list){
      const bill = table.data.list[index]
      console.log("billl",table.data.list[index])
      if(bill.confirmState != 0){
        tableData.value.push(bill)
        tempTotal++
      }
    }
    total.value = tempTotal
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
  console.log(table)
}
getTableData()

const deleteBillFunc = async(row) => {
  ElMessageBox.confirm('此操作将永久删除该记录, 是否继续?', '提示', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    type: 'warning',
  })
    .then(async() => {
      const res = await deleteBill(row)
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功!',
        })
        if (tableData.value.length === 1 && page.value > 1) {
          page.value--
        }
        getTableData()
      }
    })
    .catch(() => {
      ElMessage({
        type: 'info',
        message: '已取消删除',
      })
      console.log(page)
    })
}



const formatType = (type) =>{
  if (type == 1){
    return "收入"
  }
  if (type == 2){
    return "支出"
  }
  return "未知"
}

const DeleteBill = async(row)=>{
  const data = await deleteBill(row)
  if (data.code === 0){
    console.log('删除成功')
  }
  console.log(data)
  getTableData()
}


const formatState = (state) =>{
  if (state == 1){
    return "通过"
  }
  if (state == 2){
    return "驳回"
  }
  return "未审核"
}

const formatArchive= (archive) =>{
  if (archive){
    return "归档"
  }
  return "未归档"
}
/**
 * 编辑文件名或者备注
 * @param row
 * @returns {Promise<void>}
 */
const editFileNameFunc = async(row) => {
  ElMessageBox.prompt('请输入文件名或者备注', '编辑', {
    confirmButtonText: '确定',
    cancelButtonText: '取消',
    inputPattern: /\S/,
    inputErrorMessage: '不能为空',
    inputValue: row.name
  }).then(async({ value }) => {
    row.name = value
    // console.log(row)
    const res = await editFileName(row)
    if (res.code === 0) {
      ElMessage({
        type: 'success',
        message: '编辑成功!',
      })
      getTableData()
    }
  }).catch(() => {
    ElMessage({
      type: 'info',
      message: '取消修改'
    })
  })
}

</script>

<script>

export default {
  name: 'Upload',
}
</script>
<style scoped>
.name {
  cursor: pointer;
}

.upload-btn + .upload-btn {
  margin-left: 12px;
}
</style>

