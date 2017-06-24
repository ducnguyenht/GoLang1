package demo5

import (
	"fmt"
	"time"

	models "github.com/gorm-learn-master/demo5-error-handing/models"
	"github.com/jinzhu/gorm"
)

// Create 创建
func Create(db *gorm.DB) {
	fmt.Println("Create")
	user := models.User{
		Name:            "linhaobin",
		Birthday:        time.Now(),
		BillingAddress:  models.Address{Address1: "Billing Address - Address 1", Address2: "Billing Address - Address 2"},
		ShippingAddress: models.Address{Address1: "Shipping Address - Address 1", Address2: "Shipping Address - Address 2"},
		Emails: []models.Email{
			{Email: "linhaobin28@qq.com"},
			{Email: "864037833@qq.com"},
		},
		Languages: []models.Language{{Name: "ZH"}, {Name: "EN"}},
	}

	db.Create(&user)
}
