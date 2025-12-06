package entity

type Brand struct {
	BrandID   int    `gorm:"primaryKey" json:"brand_id"`
	BrandName string `json:"brand_name"`
}

func (Brand) TableName() string {
	return "product_brand"
}
