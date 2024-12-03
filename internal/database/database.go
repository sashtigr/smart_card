package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"log"
	"time"
)

var DB *gorm.DB

type Event struct {
	ID        uint      `gorm:"primaryKey"`
	CardUID   string    `gorm:"not null"`
	Success   bool      `gorm:"not null"`
	Timestamp time.Time `gorm:"autoCreateTime"`
}

func InitDB() {
	var err error
	DB, err = gorm.Open(sqlite.Open("access_control.db"), &gorm.Config{})
	if err != nil {
		log.Fatalf("Ошибка подключения к БД: %v", err)
	}

	DB.AutoMigrate(&User{}, &Event{})

	DB.Create(&User{Name: "Test User", CardUID: "12345678", Access: true})
}
