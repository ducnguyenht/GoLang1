package demo5

import (
	"fmt"

	models "github.com/gorm-learn-master/demo5-error-handing/models"
	"github.com/jinzhu/gorm"
)

// ErrorHanding 错误处理
func ErrorHanding(db *gorm.DB) {
	fmt.Println("Delete")

	var users []models.User
	// error
	if err := db.Where("sasdf", 1).Find(&users).Error; err != nil {
		fmt.Println(err)
		fmt.Println("错误")
	}
}
