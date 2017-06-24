package main

import (
	"time"

	"github.com/jinzhu/gorm"
)

// Customer 顾客
type Customer struct {
	gorm.Model
	Birthday time.Time
	Age      int
	Name     string `gorm:"size:255"` // Default size for string is 255, reset it with this tag
}
