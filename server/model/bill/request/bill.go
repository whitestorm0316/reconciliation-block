package request

import "server/model/bill"

type CreateBill struct {
	bill.Bill
}

type UpdateBill struct {
	bill.Bill
}

type DeleteBill struct {
	bill.Bill
}
