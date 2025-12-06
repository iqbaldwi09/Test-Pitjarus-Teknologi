package entity

type Store struct {
	StoreID   int    `gorm:"primaryKey" json:"store_id"`
	StoreName string `json:"store_name"`
	AccountID int    `json:"account_id"`
	AreaID    int    `json:"area_id"`
	IsActive  bool   `json:"is_active"`

	Account StoreAccount `gorm:"foreignKey:AccountID"`
	Area    StoreArea    `gorm:"foreignKey:AreaID"`
}

func (Store) TableName() string {
	return "store"
}
