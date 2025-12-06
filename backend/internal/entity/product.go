package entity

type Product struct {
	ProductID   int    `gorm:"primaryKey" json:"product_id"`
	ProductName string `json:"product_name"`
	BrandID     int    `json:"brand_id"`

	Brand Brand `gorm:"foreignKey:BrandID"`
}

func (Product) TableName() string {
	return "product"
}
