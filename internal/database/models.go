package database

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Name    string
	CardUID string `gorm:"unique"`
	Access  bool
}
