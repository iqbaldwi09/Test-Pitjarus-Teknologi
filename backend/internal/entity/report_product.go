package entity

import "time"

type ReportProduct struct {
	ReportID   int       `gorm:"primaryKey" json:"report_id"`
	StoreID    int       `json:"store_id"`
	ProductID  int       `json:"product_id"`
	Compliance int       `json:"compliance"`
	Tanggal    time.Time `json:"tanggal"`

	Store   Store   `gorm:"foreignKey:StoreID"`
	Product Product `gorm:"foreignKey:ProductID"`
}

func (ReportProduct) TableName() string {
	return "report_product"
}
