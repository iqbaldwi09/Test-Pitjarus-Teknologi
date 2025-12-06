package entity

type StoreAccount struct {
	AccountID   int    `gorm:"primaryKey" json:"account_id"`
	AccountName string `json:"account_name"`
}

func (StoreAccount) TableName() string {
	return "store_account"
}
