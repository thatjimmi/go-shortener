package model

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var db *gorm.DB

type Lenk struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Redirect string `json:"redirect" gorm:"not null"`
	Shorten  string `json:"shorten" gorm:"unique;not null"`
	Clicked  uint   `json:"clicked"`
	Random   bool   `json:"random"`
}

func Setup() {

	dsn := "host=127.0.0.1 user=admin password=test dbname=admin port=5987 sslmode=disable"

	var err error
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	err = db.AutoMigrate(&Lenk{})
	if err != nil {
		panic("failed to migrate database")
	}
}
