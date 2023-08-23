package dbManager

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type manager struct {
	db *gorm.DB
}

var mgr *manager

func init() {
	dsn := "host=localhost user=user password=password dbname=rsdb port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		println("DB connection error: ", err)
	}

	mgr = &manager{db: db}
}

func GetDb() *gorm.DB {
	return mgr.db
}
