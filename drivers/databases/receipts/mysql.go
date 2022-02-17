package receipts

import (
	"go-drop-logistik/business/receipts"

	"gorm.io/gorm"
)

type mysqlReceiptRepository struct {
	Conn *gorm.DB
}

func NewMySQLReceiptRepository(conn *gorm.DB) receipts.Repository {
	return &mysqlReceiptRepository{
		Conn: conn,
	}
}
