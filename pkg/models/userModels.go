package models

import "gorm.io/gorm"

type Users struct {
	Id        uint64 `gorm:"primaryKey;autoIncrement"`
	FirstName string
	LastName  string
	UserName  string
	Password  string
	gorm.Model
}
