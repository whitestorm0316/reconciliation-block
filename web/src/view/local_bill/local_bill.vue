<template>
  <div>
    <div class="gva-search-box">
      <el-form :inline="true" :model="searchInfo" class="demo-form-inline" @keyup.enter="onSubmit">
      <el-form-item label="创建时间">
      <el-date-picker v-model="searchInfo.startCreatedAt" type="datetime" placeholder="开始时间"></el-date-picker>
       —
      <el-date-picker v-model="searchInfo.endCreatedAt" type="datetime" placeholder="结束时间"></el-date-picker>
      </el-form-item>
        <el-form-item label="交易员工ID">
            
             <el-input v-model.number="searchInfo.tradingUserId" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="交易ID">
         <el-input v-model="searchInfo.transactionNumber" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="交易时间">
            
            <el-date-picker v-model="searchInfo.startTransactionTime" type="datetime" placeholder="搜索条件（起）"></el-date-picker>
            —
            <el-date-picker v-model="searchInfo.endTransactionTime" type="datetime" placeholder="搜索条件（止）"></el-date-picker>

        </el-form-item>
        <el-form-item label="交易类型">
         <el-input v-model="searchInfo.transactionType" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="交易对方">
         <el-input v-model="searchInfo.counterparty" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="支付方式">
         <el-input v-model="searchInfo.paymentMethod" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="收入或者支出">
         <el-input v-model="searchInfo.incomeOrExpenses" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="金额">
            
            <el-input v-model.number="searchInfo.startMoney" placeholder="搜索条件（起）" />
            —
            <el-input v-model.number="searchInfo.endMoney" placeholder="搜索条件（止）" />

        </el-form-item>
            <el-form-item label="是否已对账" prop="reconciled">
            <el-select v-model="searchInfo.reconciled" clearable placeholder="请选择">
                <el-option
                    key="true"
                    label="是"
                    value="true">
                </el-option>
                <el-option
                    key="false"
                    label="否"
                    value="false">
                </el-option>
            </el-select>
            </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-upload 
            class="upload-demo"
            action=""
            :on-change="handleChange"
            :on-exceed="handleExceed"
            :on-remove="handleRemove"
            :file-list="fileListUpload"
            accept=".csv"
            :auto-upload="false">
              <el-button  type="primary" icon="upload">CSV上传</el-button>
            </el-upload>
            <el-button  style="margin-left: 10px;" type="primary" icon="download" @click="exportcsv">导出csv</el-button>
            <el-button  style="margin-left: 10px;" type="primary" icon="plus" @click="openDialog">新增</el-button>
            <el-button  style="margin-left: 10px;" type="primary" icon="check" @click="openReconciliationDialog">对账</el-button>
            <el-button type="primary" icon="right" @click="openErrorDealDialog">差错处理</el-button>
          
            <el-popover v-model:visible="deleteVisible" placement="top" width="160">
            <p>确定要删除吗？</p>
            <div style="text-align: right; margin-top: 8px;">
                <el-button type="primary" link @click="deleteVisible = false">取消</el-button>
                <el-button type="primary" @click="onDelete">确定</el-button>
            </div>
            <template #reference>
                <el-button icon="delete" style="margin-left: 10px;" :disabled="!multipleSelection.length" @click="deleteVisible = true">删除</el-button>
            </template>
            </el-popover>
        </div>
        <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="tableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="交易ID" prop="transactionNumber" width="120" />
        <el-table-column align="left" label="交易员工ID" prop="tradingUserId" width="120" />
         <el-table-column align="left" label="交易时间" width="180">
            <template #default="scope">{{ formatDate(scope.row.transactionTime) }}</template>
         </el-table-column>
        <el-table-column align="left" label="交易类型" prop="transactionType" width="120" />
        <el-table-column align="left" label="交易对方" prop="counterparty" width="120" />
        <el-table-column align="left" label="支付方式" prop="paymentMethod" width="120" />
        <el-table-column align="left" label="收入或者支出" prop="incomeOrExpenses" width="120" />
        <el-table-column align="left" label="金额" prop="money" width="120" />
        <el-table-column align="left" label="是否已对账" prop="reconciled" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.reconciled) }}</template>
        </el-table-column>
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="updateLocalBillFunc(scope.row)">变更</el-button>
            <el-button type="primary" link icon="delete" @click="deleteRow(scope.row)">删除</el-button>
            </template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            @current-change="handleCurrentChange"
            @size-change="handleSizeChange"
            />
        </div>
    </div>
    <el-dialog v-model="dialogFormVisible" :before-close="closeDialog" title="弹窗操作">
      <el-form :model="formData" label-position="right" ref="elFormRef" :rules="rule" label-width="80px">
        <el-form-item label="交易员工ID:"  prop="tradingUserId" >
          <el-input v-model.number="formData.tradingUserId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="交易ID:"  prop="transactionNumber" >
          <el-input v-model="formData.transactionNumber" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="交易时间:"  prop="transactionTime" >
          <el-date-picker v-model="formData.transactionTime" type="date" style="width:100%" placeholder="选择日期" :clearable="true"  />
        </el-form-item>
        <el-form-item label="交易类型:"  prop="transactionType" >
          <el-input v-model="formData.transactionType" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="交易对方:"  prop="counterparty" >
          <el-input v-model="formData.counterparty" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="支付方式:"  prop="paymentMethod" >
          <el-input v-model="formData.paymentMethod" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="收入或者支出:"  prop="incomeOrExpenses" >
          <el-input v-model="formData.incomeOrExpenses" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="金额:"  prop="money" >
          <el-input-number v-model="formData.money"  style="width:100%" :precision="2" :clearable="true"  />
        </el-form-item>
        <el-form-item label="是否已对账:"  prop="reconciled" >
          <el-switch v-model="formData.reconciled" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
    <el-dialog v-model="errorDealFormVisible" :before-close="closeErrorDealDialog" title="弹窗操作">
      <el-form :model="errorBillFormData" label-position="right" ref="elFormRef"  label-width="80px">
        <el-form-item label="差错原因"  prop="errorReason" >
          <el-input v-model="errorReason" :clearable="true"  placeholder="请输入" />
        </el-form-item >
        <el-form-item label="审批员">
          <el-select  v-model="reviewer" placeholder="请选择审批的管理员" >
            <el-option el-option v-for="user in selectionReviewer" :label="`${user.userName}`" :value="`${user.ID}`"></el-option>
          </el-select>
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeErrorDealDialog">取 消</el-button>
          <el-button type="primary" @click="enterErrorDealDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>

    <el-dialog v-model="csvDialogVisible" :before-close="closeCsvDialog" title="弹窗操作">
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="csvTableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="交易ID" prop="transactionNumber" width="120" />
        <el-table-column align="left" label="交易员工ID" prop="tradingUserId" width="120" />
         <el-table-column align="left" label="交易时间" width="180">
            <template #default="scope">{{ formatDate(scope.row.transactionTime) }}</template>
         </el-table-column>
        <el-table-column align="left" label="交易类型" prop="transactionType" width="120" />
        <el-table-column align="left" label="交易对方" prop="counterparty" width="120" />
        <el-table-column align="left" label="支付方式" prop="paymentMethod" width="120" />
        <el-table-column align="left" label="收入或者支出" prop="incomeOrExpenses" width="120" />
        <el-table-column align="left" label="金额" prop="money" width="120" />
        <el-table-column align="left" label="是否已对账" prop="reconciled" width="120"> 
            <template #default="scope">{{ formatBoolean(scope.row.reconciled) }}</template>
        </el-table-column>
        </el-table>
        <div class="gva-pagination">
            <el-pagination
            layout="total, sizes, prev, pager, next, jumper"
            :current-page="page"
            :page-size="pageSize"
            :page-sizes="[10, 30, 50, 100]"
            :total="total"
            />
        </div>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeCsvDialog">取 消</el-button>
          <el-button type="primary" @click="enterCsvDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
    <el-dialog v-model="reconciliationDialogVisible" :before-close="closeReconciliationDialog" title="弹窗操作">
      <el-table
        ref="multipleTable"
        style="width: 100%"
        tooltip-effect="dark"
        :data="reconciledTableData"
        row-key="ID"
        @selection-change="handleSelectionChange"
        >
        <el-table-column type="selection" width="55" />
        <el-table-column align="left" label="是否匹配成功" prop="isMatch" width="120"> 
            <template #default="scope">{{ formatBoolean(scope.row.isMatch) }}</template>
        </el-table-column>
        <el-table-column align="left" label="日期" width="180">
            <template #default="scope">{{ formatDate(scope.row.CreatedAt) }}</template>
        </el-table-column>
        <el-table-column align="left" label="交易员工ID" prop="tradingUserId" width="120" />
        <el-table-column align="left" label="交易ID" prop="transactionNumber" width="120" />
         <el-table-column align="left" label="交易时间" width="180">
            <template #default="scope">{{ formatDate(scope.row.transactionTime) }}</template>
         </el-table-column>
        <el-table-column align="left" label="交易类型" prop="transactionType" width="120" />
        <el-table-column align="left" label="交易对方" prop="counterparty" width="120" />
        <el-table-column align="left" label="支付方式" prop="paymentMethod" width="120" />
        <el-table-column align="left" label="收入或者支出" prop="incomeOrExpenses" width="120" />
        <el-table-column align="left" label="金额" prop="money" width="120" />

        <el-table-column align="left" label="第三方交易ID" prop="thirdTransactionNumber" width="120" />
         <el-table-column align="left" label="第三方交易时间" width="180">
            <template #default="scope">{{ formatDate(scope.row.thirdTransactionTime) }}</template>
         </el-table-column>
        <el-table-column align="left" label="第三方交易类型" prop="thirdTransactionType" width="120" />
        <el-table-column align="left" label="第三方交易对方" prop="thirdCounterparty" width="120" />
        <el-table-column align="left" label="第三方支付方式" prop="thirdPaymentMethod" width="120" />
        <el-table-column align="left" label="第三方收入或者支出" prop="thirdIncomeOrExpenses" width="120" />
        <el-table-column align="left" label="第三方金额" prop="thirdMoney" width="120" />
        <el-table-column align="left" label="第三方账单所属机构" prop="thirdOrganization" width="120" />

        <el-table-column align="left" label="本地账单ID" prop="ID" width="120" />
        <el-table-column align="left" label="第三方账单ID" prop="thirdID" width="120" />
        </el-table>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeReconciliationDialog">取 消</el-button>
          <el-button  style="margin-left: 10px;" type="primary" icon="check" @click="enterReconciliationDialog">对账</el-button>
          <el-button type="primary" icon="right" @click="openErrorDealDialog">差错处理</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'LocalBill'
}
</script>

<script setup>
import {
  createLocalBill,
  deleteLocalBill,
  deleteLocalBillByIds,
  updateLocalBill,
  findLocalBill,
  getLocalBillList
} from '@/api/local_bill'

import {
  createErrorBill
} from '@/api/error_bill'

import {
  getThirdPartyBillList,
  updateThirdPartyBill
} from '@/api/third_party_bill'

import {
  createReconciledBill
} from '@/api/reconciled_bill'


// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'
import {getUsersByAuthorityId} from '@/api/user'
import Papa from 'papaparse';
import { saveAs } from "file-saver";
import { useUserStore } from '@/pinia/modules/user'

const userStore = useUserStore()

// 自动化生成的字典（可能为空）以及字段
const formData = ref({
        tradingUserId: 0,
        transactionNumber: '',
        transactionTime: new Date(),
        transactionType: '',
        counterparty: '',
        paymentMethod: '',
        incomeOrExpenses: '',
        money: 0,
        reconciled: false,
        })

// 验证规则
const rule = reactive({
               money : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()


//  ============对账部分==============
const reconciliationDialogVisible = ref(false)
const reconciledTableData = ref([])


// 打开对账处理弹窗
const  openReconciliationDialog = async()=>{
  reconciliationDialogVisible.value = true
  multipleSelection.value &&
    multipleSelection.value.map(item => {
      const searchInfo = {}
      searchInfo.transactionNumber = item.transactionNumber
      searchInfo.reconciled = false
      getThirdPartyBillList({ page: page.value, pageSize: pageSize.value, ...searchInfo }).then(
        res =>{
          if (res.data.list.length == 0){
            item.isMatch = false
            item.thirdTransactionTime = null
          }else{
            item.isMatch = true
            item.thirdTransactionNumber = res.data.list[0].transactionNumber
            item.thirdTransactionTime = res.data.list[0].transactionTime
            item.thirdTransactionType = res.data.list[0].transactionType
            item.thirdCounterparty =  res.data.list[0].counterparty
            item.thirdPaymentMethod =  res.data.list[0].paymentMethod
            item.thirdMoney = res.data.list[0].money
            item.thirdOrganization = res.data.list[0].organization
            item.thirdCreatedAt = res.data.list[0].CreatedAt
            item.third = res.data.list[0]
          }
          
          reconciledTableData.value.push(item)
          console.log(res.data.list)
          
        }
      )
    })
    multipleSelection.value = []

  
}
// 关闭对账弹窗
const closeReconciliationDialog = ()=>{
  reconciliationDialogVisible.value = false
  reconciledTableData.value = []
  console.log(reviewer.value)
}

const enterReconciliationDialog = async()=>{
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要对账的数据'
    })
    return
  }
  multipleSelection.value &&
    multipleSelection.value.map(item => {
      console.log('item:',item)
      item.organization = item.thirdOrganization
      item.localBillId = item.ID
      item.thirdPartyBillId = item.thirdID
      item.reconciler = userStore.userInfo.ID
      
      
      createReconciledBill(item).then(res =>{
        
        if (res.code === 0){
          ElMessage({
          type: 'success',
          message: 'ID:'+item.ID+'对账处理完成'
        })
        item.reconciled = true
        updateLocalBill(item)
        item.third.reconciled = true
        updateThirdPartyBill( item.third)
        }
      })
      
    })

}


//  ============差错处理部分 ==========
const errorBillFormData = ref({
        tradingUserId: 0,
        transactionNumber: '',
        transactionTime: new Date(),
        transactionType: '',
        counterparty: '',
        paymentMethod: '',
        incomeOrExpenses: '',
        money: 0,
        reconciled: false,
        organization: '',
        reconciler: 0,
        reviewer: 0,
        reason: '',
        errorReason: '',
        note: '',
        localBillId: 0,
        thirdPartyBillId: 0,
        })

const errorReason = ref('')
const reviewer = ref(0)
const selectionReviewer = ref([])


const getReviewers = async()=>{
  const data = await getUsersByAuthorityId({authority_id: "3"})
  if (data.code === 0){
    selectionReviewer.value = data.data
  }
  console.log(data)
}


// 打开差错处理弹窗
const openErrorDealDialog = ()=>{
  errorDealFormVisible.value = true
  getReviewers()
}
// 关闭差错弹窗
const closeErrorDealDialog = ()=>{
  errorDealFormVisible.value = false
  console.log(reviewer.value)
}

const enterErrorDealDialog = async()=>{
  const ids = []
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要差错处理的数据'
    })
    return
  }
  multipleSelection.value &&
    multipleSelection.value.map(item => {
      ids.push(item.ID)
      item.errorReason = errorReason.value
      item.reviewer = Number(reviewer.value)
      item.localBillId = Number(item.ID)
      console.log('item:',item)
      createErrorBill(item).then(res =>{
        if (res.code === 0){
            ElMessage({
            type: 'success',
            message: 'ID:'+item.ID+'差错处理完成'
          })
        }else{
          ElMessage({
            type: 'error',
            message: 'ID:'+item.ID+'差错处理失败'
        })
      }

      item.reconciled = true 
      updateLocalBill(item).then(res =>{
        if(res.code === 0){
          console.log(修改本地账单成功)
        }
      })
      item.third.reconciled = true
      updateThirdPartyBill(item.third).then(res => {
        if (res.code === 0){
          console.log(修改第三方账单成功)
        }
      })
    })
            
    })

  for (let i = 0; i < reconciledTableData.value.length; i++) {
    const id =  reconciledTableData.value[i].ID;
    for(let j = 0;j<ids.length;j++){
      if(id == ids[j]){
        reconciledTableData.value.splice(i,1)
      }
    }
  }

  
  errorDealFormVisible.value = true
  getTableData()
}

//  ============CSV导入部分 ==========
const csvTableData = ref([])
const csvDialogVisible = ref(false)

// 打开csv数据弹窗
const openCsvDialog = ()=>{
  csvDialogVisible.value = true
  getReviewers()
}
// 关闭csv数据弹窗
const closeCsvDialog = ()=>{
  csvDialogVisible.value = false
  csvTableData = ref([])
  console.log(reviewer.value)
}

const enterCsvDialog = async()=>{
  if (multipleSelection.value.length === 0) {
    ElMessage({
      type: 'warning',
      message: '请选择要添加的数据'
    })
    return
  }
  multipleSelection.value &&
    multipleSelection.value.map(item => {
      const res = createLocalBill(item)
    })
    csvDialogVisible.value = false
    getTableData()
}
const handleChange = (file,fileList)=>{
    const fileTemp = file.raw
    if (fileTemp) {
        if ((fileTemp.type == 'text/csv') || (fileTemp.type == 'application/vnd.ms-excel')) {
        importcsv(file.raw)
      } else { 
        console.log(fileTemp.type)
        ElMessage({
          type: 'warning',
          message: '附件格式错误，请删除后重新上传！'
        })
      }
    } else {
      ElMessage({
        type: 'warning',
            message: '请上传附件！'
          })
      }
}

const importcsv =(obj)=>{
  Papa.parse(obj, {
    complete (results) {
      console.log(results)//这个是csv文件的数据
    
      //遍历csv文件中的数据，存放到data中 方法不唯一，可自己更改
      for (let i = 1; i < results.data.length; i++) {
        let obj = {}
        obj.transactionNumber = results.data[i][1]
        obj.tradingUserId = Number(results.data[i][0])
        obj.transactionTime = new Date(results.data[i][2])
        obj.transactionType = results.data[i][3]
        obj.counterparty = results.data[i][4]
        obj.paymentMethod = results.data[i][5]
        obj.incomeOrExpenses = results.data[i][6]
        obj.money = Number(results.data[i][7])
        obj.reconciled = getBoolean( results.data[i][8])
        csvTableData.value.push(obj)
      }
      csvTableData.value.splice(csvTableData.value.length-1,csvTableData.value.length)
      openCsvDialog()
    }
  })
}

const getBoolean = (bstr)=>{
  if (bstr == 'TRUE'){
    return true
  }else {
    return false
  }
}

const exportcsv =()=>{
  const csv = Papa.unparse(multipleSelection.value)
  const blob = new Blob([csv], { type: 'text/csv;charset=utf-8' })
  saveAs(blob, 'data.csv')
}


// =========== 表格控制部分 ===========
const page = ref(1)
const total = ref(0)
const pageSize = ref(10)
const tableData = ref([])
const searchInfo = ref({})

// 重置
const onReset = () => {
  searchInfo.value = {}
  getTableData()
}

// 搜索
const onSubmit = () => {
  page.value = 1
  pageSize.value = 10
  if (searchInfo.value.reconciled === ""){
      searchInfo.value.reconciled=null
  }
  getTableData()
}

// 分页
const handleSizeChange = (val) => {
  pageSize.value = val
  getTableData()
}

// 修改页面容量
const handleCurrentChange = (val) => {
  page.value = val
  getTableData()
}

// 查询
const getTableData = async() => {
  const table = await getLocalBillList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
  console.log(table)
}

getTableData()

// ============== 表格控制部分结束 ===============

// 获取需要的字典 可能为空 按需保留
const setOptions = async () =>{
}

// 获取需要的字典 可能为空 按需保留
setOptions()


// 多选数据
const multipleSelection = ref([])
// 多选
const handleSelectionChange = (val) => {
    multipleSelection.value = val
}

// 删除行
const deleteRow = (row) => {
    ElMessageBox.confirm('确定要删除吗?', '提示', {
        confirmButtonText: '确定',
        cancelButtonText: '取消',
        type: 'warning'
    }).then(() => {
            deleteLocalBillFunc(row)
        })
    }


// 批量删除控制标记
const deleteVisible = ref(false)

// 多选删除
const onDelete = async() => {
      const ids = []
      if (multipleSelection.value.length === 0) {
        ElMessage({
          type: 'warning',
          message: '请选择要删除的数据'
        })
        return
      }
      multipleSelection.value &&
        multipleSelection.value.map(item => {
          ids.push(item.ID)
        })
      const res = await deleteLocalBillByIds({ ids })
      if (res.code === 0) {
        ElMessage({
          type: 'success',
          message: '删除成功'
        })
        if (tableData.value.length === ids.length && page.value > 1) {
          page.value--
        }
        deleteVisible.value = false
        getTableData()
      }
    }



// 行为控制标记（弹窗内部需要增还是改）
const type = ref('')

// 更新行
const updateLocalBillFunc = async(row) => {
    const res = await findLocalBill({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.relb
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteLocalBillFunc = async (row) => {
    const res = await deleteLocalBill({ ID: row.ID })
    if (res.code === 0) {
        ElMessage({
                type: 'success',
                message: '删除成功'
            })
            if (tableData.value.length === 1 && page.value > 1) {
            page.value--
        }
        getTableData()
    }
}

// 弹窗控制标记
const dialogFormVisible = ref(false)
const errorDealFormVisible = ref(false)

// 打开弹窗
const openDialog = () => {
    type.value = 'create'
    dialogFormVisible.value = true
}



// 关闭弹窗
const closeDialog = () => {
    dialogFormVisible.value = false
    formData.value = {
        tradingUserId: 0,
        transactionNumber: '',
        transactionTime: new Date(),
        transactionType: '',
        counterparty: '',
        paymentMethod: '',
        incomeOrExpenses: '',
        money: 0,
        reconciled: false,
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createLocalBill(formData.value)
                  break
                case 'update':
                  res = await updateLocalBill(formData.value)
                  break
                default:
                  res = await createLocalBill(formData.value)
                  break
              }
              if (res.code === 0) {
                ElMessage({
                  type: 'success',
                  message: '创建/更改成功'
                })
                closeDialog()
                getTableData()
              }
      })
}
</script>

<style>
</style>
