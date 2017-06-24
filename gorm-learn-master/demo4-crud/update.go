package main

import (
	"fmt"

	"github.com/gorm-learn-master/demo4-crud/models"
	"github.com/jinzhu/gorm"
)

// Update 更新
func Update(db *gorm.DB) {
	fmt.Println("Update")

	// Save
	fmt.Println("---Save")
	var user models.User
	db.First(&user)

	user.Name = "linhaobin 2"
	user.Age = 25

	db.Save(&user)
	fmt.Println(&user)

	// Update
	fmt.Println("---update")
	db.Model(&user).Update("Name", "linhaobin 3")
	fmt.Println(&user)
	db.Model(&user).Updates(&models.User{Name: "linhaobin 4", Age: 26})
	fmt.Println(&user)

	// Batch Updates
	fmt.Println("---batch updates")
	db.Model(&models.User{}).Where("name like ?", "%lin%").Updates(&models.User{Age: 30})
	fmt.Println(&user)

}
