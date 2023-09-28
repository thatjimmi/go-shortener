package model

import (
	"os"

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

	dsn := os.Getenv("DATABASE_URL")

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
