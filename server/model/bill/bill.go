package bill

import (
	"server/global"
)

type Bill struct {
	global.GVA_MODEL
	// 申请人名字
	Applicant string `json:"applicant" gorm:"index;comment:申请人名字"`
	// 申请人ID
	ApplicantId uint `json:"applicantId" gorm:"index;comment:申请人ID"`
	// 审批人名字
	Approver string `json:"approver" gorm:"default:系统用户;comment:审批人名字"`
	// 审批人ID
	ApproverId uint `json:"approverId" gorm:"index;comment:审批人ID"`
	// 申报类型  1表示收入  2表示支出
	Type uint `json:"type" gorm:"index;comment:申报类型"`
	// 管理员
	Admin string `json:"admin" gorm:"index;comment:管理员"`
	// 管理员ID
	AdminId uint `json:"adminId" gorm:"index;comment:管理员ID"`
	// 申请日期
	Time string `json:"time" gorm:"comment:申请日期"`
	// 金额
	Money string `json:"money" gorm:"comment:金额"`
	// 发票URL
	VoucherImgUrl string `json:"voucherImgUrl" gorm:"default:https://qmplusimg.henrongyi.top/gva_header.jpg;comment:发票URL"`
	// 申报理由
	Reason string `json:"reason" gorm:"comment:申报理由"`
	// 审批状态 0表示未审批  1表示审批通过  2表示审批未通过
	ApprovalState uint `json:"approvalState" gorm:"comment:审批状态  "`
	// 确认状态 0表示未确认   1表示确认通过  2表示驳回
	ConfirmState uint `json:"confirmState" gorm:"comment:确认状态"`
	// 审批信息
	AppravalInfo string `json:"appravalInfo" gorm:"comment:审批信息"`
	// 确认信息
	ConfirmInfo string `json:"confirmInfo" gorm:"comment:确认信息"`

	// 发起人
	Initiator string `json:"initiator" gorm:"comment:发起人"`

	// 发起人ID
	InitiatorId uint `json:"initiatorId" gorm:"comment:发起人ID"`

	// 是否归档，归档后不能修改
	Archive bool `json:"archive" gorm:"comment: 是否归档"`
}

func (Bill) TableName() string {
	return "bills"
}
