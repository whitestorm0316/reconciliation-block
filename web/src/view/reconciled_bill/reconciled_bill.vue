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
        <el-form-item label="账单所属机构">
         <el-input v-model="searchInfo.organization" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="对账员ID">
            
             <el-input v-model.number="searchInfo.reconciler" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="审核员ID">
            
             <el-input v-model.number="searchInfo.reviewer" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="备注">
         <el-input v-model="searchInfo.note" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="本地账单ID">
            
             <el-input v-model.number="searchInfo.localBillId" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item label="第三方账单ID">
            
             <el-input v-model.number="searchInfo.thirdPartyBillId" placeholder="搜索条件" />

        </el-form-item>
            <el-form-item label="是否上传至区块链" prop="blocked">
            <el-select v-model="searchInfo.blocked" clearable placeholder="请选择">
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
        <el-form-item label="合约名称">
         <el-input v-model="searchInfo.contractName" placeholder="搜索条件" />

        </el-form-item>
        <el-form-item>
          <el-button type="primary" icon="search" @click="onSubmit">查询</el-button>
          <el-button icon="refresh" @click="onReset">重置</el-button>
        </el-form-item>
      </el-form>
    </div>
    <div class="gva-table-box">
        <div class="gva-btn-list">
            <el-button type="primary" icon="plus" @click="openDialog">新增</el-button>
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
        <el-table-column align="left" label="账单所属机构" prop="organization" width="120" />
        <el-table-column align="left" label="对账员ID" prop="reconciler" width="120" />
        <el-table-column align="left" label="审核员ID" prop="reviewer" width="120" />
        <el-table-column align="left" label="备注" prop="note" width="120" />
        <el-table-column align="left" label="本地账单ID" prop="localBillId" width="120" />
        <el-table-column align="left" label="第三方账单ID" prop="thirdPartyBillId" width="120" />
        <el-table-column align="left" label="是否上传至区块链" prop="blocked" width="120">
            <template #default="scope">{{ formatBoolean(scope.row.blocked) }}</template>
        </el-table-column>
        <el-table-column align="left" label="合约名称" prop="contractName" width="120" />
        <el-table-column align="left" label="按钮组">
            <template #default="scope">
            <el-button type="primary" link icon="edit" class="table-button" @click="updateReconciledBillFunc(scope.row)">变更</el-button>
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
        <el-form-item label="账单所属机构:"  prop="organization" >
          <el-input v-model="formData.organization" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="对账员ID:"  prop="reconciler" >
          <el-input v-model.number="formData.reconciler" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="审核员ID:"  prop="reviewer" >
          <el-input v-model.number="formData.reviewer" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="备注:"  prop="note" >
          <el-input v-model="formData.note" :clearable="true"  placeholder="请输入" />
        </el-form-item>
        <el-form-item label="本地账单ID:"  prop="localBillId" >
          <el-input v-model.number="formData.localBillId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="第三方账单ID:"  prop="thirdPartyBillId" >
          <el-input v-model.number="formData.thirdPartyBillId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="是否上传至区块链:"  prop="blocked" >
          <el-switch v-model="formData.blocked" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="合约名称:"  prop="contractName" >
          <el-input v-model="formData.contractName" :clearable="true"  placeholder="请输入" />
        </el-form-item>
      </el-form>
      <template #footer>
        <div class="dialog-footer">
          <el-button @click="closeDialog">取 消</el-button>
          <el-button type="primary" @click="enterDialog">确 定</el-button>
        </div>
      </template>
    </el-dialog>
  </div>
</template>

<script>
export default {
  name: 'ReconciledBill'
}
</script>

<script setup>
import {
  createReconciledBill,
  deleteReconciledBill,
  deleteReconciledBillByIds,
  updateReconciledBill,
  findReconciledBill,
  getReconciledBillList
} from '@/api/reconciled_bill'

// 全量引入格式化工具 请按需保留
import { getDictFunc, formatDate, formatBoolean, filterDict } from '@/utils/format'
import { ElMessage, ElMessageBox } from 'element-plus'
import { ref, reactive } from 'vue'

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
        organization: '',
        reconciler: 0,
        reviewer: 0,
        note: '',
        localBillId: 0,
        thirdPartyBillId: 0,
        blocked: false,
        contractName: '',
        })

// 验证规则
const rule = reactive({
               transactionNumber : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
               money : [{
                   required: true,
                   message: '',
                   trigger: ['input','blur'],
               }],
})

const elFormRef = ref()


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
  if (searchInfo.value.blocked === ""){
      searchInfo.value.blocked=null
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
  const table = await getReconciledBillList({ page: page.value, pageSize: pageSize.value, ...searchInfo.value })
  if (table.code === 0) {
    tableData.value = table.data.list
    total.value = table.data.total
    page.value = table.data.page
    pageSize.value = table.data.pageSize
  }
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
            deleteReconciledBillFunc(row)
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
      const res = await deleteReconciledBillByIds({ ids })
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
const updateReconciledBillFunc = async(row) => {
    const res = await findReconciledBill({ ID: row.ID })
    type.value = 'update'
    if (res.code === 0) {
        formData.value = res.data.rerb
        dialogFormVisible.value = true
    }
}


// 删除行
const deleteReconciledBillFunc = async (row) => {
    const res = await deleteReconciledBill({ ID: row.ID })
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
        organization: '',
        reconciler: 0,
        reviewer: 0,
        note: '',
        localBillId: 0,
        thirdPartyBillId: 0,
        blocked: false,
        contractName: '',
        }
}
// 弹窗确定
const enterDialog = async () => {
     elFormRef.value?.validate( async (valid) => {
             if (!valid) return
              let res
              switch (type.value) {
                case 'create':
                  res = await createReconciledBill(formData.value)
                  break
                case 'update':
                  res = await updateReconciledBill(formData.value)
                  break
                default:
                  res = await createReconciledBill(formData.value)
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
