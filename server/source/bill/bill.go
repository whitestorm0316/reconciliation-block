package bill

import (
	"context"
	"server/model/bill"
	"server/service/system"

	"gorm.io/gorm"
)

type initBill struct{}

const initOrderApi = system.InitOrderSystem + 1

// auto run
func init() {
	system.RegisterInit(initOrderApi, &initBill{})
}

func (i initBill) InitializerName() string {
	return bill.Bill{}.TableName()
}

func (i *initBill) MigrateTable(ctx context.Context) (context.Context, error) {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return ctx, system.ErrMissingDBContext
	}
	return ctx, db.AutoMigrate(&bill.Bill{})
}

func (i *initBill) TableCreated(ctx context.Context) bool {
	db, ok := ctx.Value("db").(*gorm.DB)
	if !ok {
		return false
	}
	return db.Migrator().HasTable(&bill.Bill{})
}

func (i *initBill) InitializeData(ctx context.Context) (context.Context, error) {
	return ctx, nil
}

func (i *initBill) DataInserted(ctx context.Context) bool {
	return true
}
