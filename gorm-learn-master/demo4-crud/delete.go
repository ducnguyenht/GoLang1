package main

import (
	"fmt"

	"github.com/gorm-learn-master/demo4-crud/models"
	"github.com/jinzhu/gorm"
)

// Delete 更新
func Delete(db *gorm.DB) {
	fmt.Println("Delete")

	// Delete
	fmt.Println("---Delete")
	var user models.User
	db.First(&user)

	db.Delete(&user)
	fmt.Println(&user)

	// Batch Delete
	db.Where("name like ?", "%lin%").Delete(&models.User{})

	// Soft Delete and Unscoped
	fmt.Println("---Soft Delete and Unscoped")
	var users []models.User
	db.Find(&users)
	fmt.Println(&users)
	db.Unscoped().Find(&users)
	fmt.Println(&users)
}
