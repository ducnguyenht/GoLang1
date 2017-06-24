package main

import (
	"fmt"

	"github.com/gorm-learn-master/demo4-crud/models"
	"github.com/jinzhu/gorm"
)

// Query 查询
func Query(db *gorm.DB) {
	fmt.Println("Query")

	fmt.Println("first")
	var user models.User
	db.First(&user)
	fmt.Println(&user)

	fmt.Println("---id = 9999")
	db.First(&user, 9999) // 数据库不存在
	fmt.Println(&user)

	// where
	fmt.Println("---name = linhaobin")
	var users []models.User
	db.Where("name = ?", "linhaobin").Find(&users)
	fmt.Println(&users)

	// struct
	fmt.Println("---struct: name = linhaobin")
	db.Where(&models.User{Name: "linhaobin"}).Find(&users)
	fmt.Println(&users)

	// preload
	fmt.Println("---preload Email: name = linhaobin")
	db.Where(&models.User{Name: "linhaobin"}).Preload("CreditCard").Preload("Emails").Preload("BillingAddress").Preload("ShippingAddress").Preload("Languages").Find(&users)
	fmt.Println(&users)
}
