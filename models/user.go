package models

import "github.com/jinzhu/gorm"

type Shouter struct {
	gorm.Model
	Login     		 Login `gorm:"polymorphic:Owner;"`
	FirstName      string    `gorm:"size:255"` // Default size for string is 255, reset it with this tag
	LastName       string    `gorm:"size:255"` // Default size for string is 255, reset it with this tag
}
