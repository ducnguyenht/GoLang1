package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/gorm-learn-master/demo4-crud/models"
	"github.com/jinzhu/gorm"
)

func main() {
	db, err := gorm.Open("mysql", "sa:123456@tcp(localhost:3306)/ductestmysql?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("ket noi that bai")
		fmt.Println(err)
	} else {
		fmt.Println("ket noi thanh cong")
	}

	db.LogMode(true)

	// 创建表
	autoMigrate(db)

	// 增删改查
	Create(db)
	Query(db)
	Update(db)
	Delete(db)

	defer func() {
		fmt.Scanf("%s")
		db.Close()
		fmt.Println("关闭连接")
	}()
}

// 创建表
func autoMigrate(db *gorm.DB) {
	db.DropTableIfExists(
		&models.Address{},
		&models.CreditCard{},
		&models.Email{},
		&models.Language{},
		&models.User{},
	)

	db.AutoMigrate(
		&models.Address{},
		&models.CreditCard{},
		&models.Email{},
		&models.Language{},
		&models.User{},
	)
}
