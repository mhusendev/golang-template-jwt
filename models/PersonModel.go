package models

import (
	"time"

	"gorm.io/gorm"
)

type Person struct {
	gorm.Model
	ID        uint
	Name      string
	Address   string
	Phone     string
	Username  string `gorm:"unique;not null"`
	Password  string `gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

func (Person) TableName() string {
	return "persons"
}
