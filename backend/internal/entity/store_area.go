package entity

type StoreArea struct {
	AreaID   int    `gorm:"primaryKey" json:"area_id"`
	AreaName string `json:"area_name"`
}

func (StoreArea) TableName() string {
	return "store_area"
}
