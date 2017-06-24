package models

import "github.com/jinzhu/gorm"

// CreditCard 信用卡
type CreditCard struct {
	gorm.Model
	UserID uint
	Number string
}
