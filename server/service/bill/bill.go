package bill

import (
	"fmt"
	"server/global"
	"server/model/bill"
	"server/model/common/request"
	"strconv"
)

type BillService struct {
}

func (b *BillService) CreateBill(bl bill.Bill) error {
	if err := global.GVA_DB.Create(&bl).Error; err != nil {
		return err
	}
	return nil
}
func (b *BillService) GetBillList(info request.PageInfo) (interface{}, int64, error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	fmt.Println(info)
	db := global.GVA_DB.Model(&bill.Bill{})
	var billList []bill.Bill
	var count int64
	if err := db.Model(&bill.Bill{}).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	err := db.Limit(limit).Offset(offset).Find(&billList).Error
	if err != nil {
		return nil, 0, err
	}

	fmt.Println(count)

	return billList, count, err
}

// 通过名字查找
func (b *BillService) GetBillListByName(info request.PageInfo) (interface{}, int64, error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&bill.Bill{})
	var billList []bill.Bill
	var count int64
	selectName := "%" + info.Keyword + "%"
	err := db.Where("applicant like ? or  approver like ? or admin like ?", selectName, selectName, selectName).Limit(limit).Offset(offset).Find(&billList).Limit(-1).Offset(-1).Count(&count).Error
	if err != nil {
		return nil, 0, err
	}

	return billList, count, nil
}

func (b *BillService) UpdateBill(bl bill.Bill) error {
	db := global.GVA_DB.Model(&bill.Bill{})
	var c int64
	fmt.Printf("%+v\n", bl)
	if err := db.Where("id=?", bl.ID).Count(&c).Error; err != nil {
		return err
	}
	fmt.Println(c)
	if c == 1 {
		if err := db.Save(&bl).Error; err != nil {
			return err
		}
	}
	return nil
}

func (b *BillService) DeleteBill(bl bill.Bill) error {
	db := global.GVA_DB.Model(&bill.Bill{})
	if err := db.Delete("id=?", bl.ID).Error; err != nil {
		return err
	}
	return nil
}

func (b *BillService) GetBillByApproverId(info request.PageInfo) (interface{}, int64, error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&bill.Bill{})
	var billList []bill.Bill
	var count int64
	id, err := strconv.Atoi(info.Keyword)
	if err != nil {
		return nil, 0, err
	}
	if err := db.Where("approver_id=?", id).Limit(limit).Offset(offset).Find(&billList).Limit(-1).Offset(-1).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return billList, count, nil
}
func (b *BillService) GetBillByInitiatorId(info request.PageInfo) (interface{}, int64, error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&bill.Bill{})
	var billList []bill.Bill
	var count int64
	id, err := strconv.Atoi(info.Keyword)
	if err != nil {
		return nil, 0, err
	}
	if err := db.Where("initiator_id=?", id).Limit(limit).Offset(offset).Find(&billList).Limit(-1).Offset(-1).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return billList, count, nil
}

func (b *BillService) GetBillByApplicantId(info request.PageInfo) (interface{}, int64, error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&bill.Bill{})
	var billList []bill.Bill
	var count int64
	id, err := strconv.Atoi(info.Keyword)
	if err != nil {
		return nil, 0, err
	}
	if err := db.Where("applicant_id=?", id).Limit(limit).Offset(offset).Find(&billList).Limit(-1).Offset(-1).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return billList, count, nil
}
func (b *BillService) GetBillByAdminId(info request.PageInfo) (interface{}, int64, error) {
	limit := info.PageSize
	offset := info.PageSize * (info.Page - 1)
	db := global.GVA_DB.Model(&bill.Bill{})
	var billList []bill.Bill
	var count int64
	id, err := strconv.Atoi(info.Keyword)
	if err != nil {
		return nil, 0, err
	}
	if err := db.Where("admin_id=?", id).Limit(limit).Offset(offset).Find(&billList).Limit(-1).Offset(-1).Count(&count).Error; err != nil {
		return nil, 0, err
	}
	return billList, count, nil
}
