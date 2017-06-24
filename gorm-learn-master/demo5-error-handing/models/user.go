package models

import (
	"encoding/json"
	"time"

	"github.com/jinzhu/gorm"
)

// User 用户
type User struct {
	gorm.Model
	Birthday time.Time
	Age      int
	Name     string `gorm:"size:255"` // Default size for string is 255, reset it with this tag
	// Num      int    `gorm:"AUTO_INCREMENT"`

	CreditCard CreditCard // 一对一 (has one - 使用CreditCard 的 UserID 作为外键)
	Emails     []Email    // 一对多 (has many - 使用 Email 的 UserID 作为外键)

	BillingAddress   Address // 一对一 (belongs to - 使用 BillingAddressID 作为外键)
	BillingAddressID int     `gorm:"type:int(11)"`

	ShippingAddress   Address // 一对一 (belongs to - 使用 ShippingAddressID 作为外键)
	ShippingAddressID int

	IgnoreMe  int        `gorm:"-"`                         // 忽略该字段，不是数据库字段
	Languages []Language `gorm:"many2many:user_languages;"` // 多对多, 'user_languages' 为连接表
}

func (u User) String() string {
	str, err := json.Marshal(&u)
	if err != nil {
		return "json error"
	}
	return string(str)
}
