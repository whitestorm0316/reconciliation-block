<template>
  <div>
    <div class="gva-form-box">
      <el-form :model="formData" ref="elFormRef" label-position="right" :rules="rule" label-width="80px">
        <el-form-item label="交易员工ID:" prop="tradingUserId">
          <el-input v-model.number="formData.tradingUserId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="交易ID:" prop="transactionNumber">
          <el-input v-model="formData.transactionNumber" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="交易时间:" prop="transactionTime">
          <el-date-picker v-model="formData.transactionTime" type="date" placeholder="选择日期" :clearable="true"></el-date-picker>
        </el-form-item>
        <el-form-item label="交易类型:" prop="transactionType">
          <el-input v-model="formData.transactionType" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="交易对方:" prop="counterparty">
          <el-input v-model="formData.counterparty" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="支付方式:" prop="paymentMethod">
          <el-input v-model="formData.paymentMethod" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="收入或者支出:" prop="incomeOrExpenses">
          <el-input v-model="formData.incomeOrExpenses" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="金额:" prop="money">
          <el-input-number v-model="formData.money" :precision="2" :clearable="true"></el-input-number>
        </el-form-item>
        <el-form-item label="是否已对账:" prop="reconciled">
          <el-switch v-model="formData.reconciled" active-color="#13ce66" inactive-color="#ff4949" active-text="是" inactive-text="否" clearable ></el-switch>
        </el-form-item>
        <el-form-item label="账单所属机构:" prop="organization">
          <el-input v-model="formData.organization" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="对账员ID:" prop="reconciler">
          <el-input v-model.number="formData.reconciler" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="审核员ID:" prop="reviewer">
          <el-input v-model.number="formData.reviewer" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="审核进度:" prop="progress">
          <el-input v-model="formData.progress" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="审核理由:" prop="reason">
          <el-input v-model="formData.reason" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="差错原因:" prop="errorReason">
          <el-input v-model="formData.errorReason" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="备注:" prop="note">
          <el-input v-model="formData.note" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="本地账单ID:" prop="localBillId">
          <el-input v-model.number="formData.localBillId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item label="第三方账单ID:" prop="thirdPartyBillId">
          <el-input v-model.number="formData.thirdPartyBillId" :clearable="true" placeholder="请输入" />
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="save">保存</el-button>
          <el-button type="primary" @click="back">返回</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
export default {
  name: 'ErrorBill'
}
</script>

<script setup>
import {
  createErrorBill,
  updateErrorBill,
  findErrorBill
} from '@/api/error_bill'

// 自动获取字典
import { getDictFunc } from '@/utils/format'
import { useRoute, useRouter } from "vue-router"
import { ElMessage } from 'element-plus'
import { ref, reactive } from 'vue'
const route = useRoute()
const router = useRouter()

const type = ref('')
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
            organization: '',
            reconciler: 0,
            reviewer: 0,
            progress: '',
            reason: '',
            errorReason: '',
            note: '',
            localBillId: 0,
            thirdPartyBillId: 0,
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

// 初始化方法
const init = async () => {
 // 建议通过url传参获取目标数据ID 调用 find方法进行查询数据操作 从而决定本页面是create还是update 以下为id作为url参数示例
    if (route.query.id) {
      const res = await findErrorBill({ ID: route.query.id })
      if (res.code === 0) {
        formData.value = res.data.reeb
        type.value = 'update'
      }
    } else {
      type.value = 'create'
    }
}

init()
// 保存按钮
const save = async() => {
      elFormRef.value?.validate( async (valid) => {
         if (!valid) return
            let res
           switch (type.value) {
             case 'create':
               res = await createErrorBill(formData.value)
               break
             case 'update':
               res = await updateErrorBill(formData.value)
               break
             default:
               res = await createErrorBill(formData.value)
               break
           }
           if (res.code === 0) {
             ElMessage({
               type: 'success',
               message: '创建/更改成功'
             })
           }
       })
}

// 返回按钮
const back = () => {
    router.go(-1)
}

</script>

<style>
</style>
