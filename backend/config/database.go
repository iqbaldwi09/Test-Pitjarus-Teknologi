package config

import (
	"log"

	"go-backend/internal/entity"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func ConnectDB() *gorm.DB {
	dsn := "root:@tcp(127.0.0.1:3306)/dev_test114?charset=utf8mb4&parseTime=True&loc=Local"

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Database connection failed:", err)
	}

	db.AutoMigrate(
		&entity.Store{},
		&entity.StoreAccount{},
		&entity.StoreArea{},
		&entity.Product{},
		&entity.Brand{},
		&entity.ReportProduct{},
	)

	return db
}
