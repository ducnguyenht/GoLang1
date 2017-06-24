package main

import (
	"fmt"

	"github.com/gorm-learn-master/demo4-crud/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

/*"github.com/linhaobin/gorm-learn/demo3-models/models"*/
func main() {
	db, err := gorm.Open("mysql", "sa:123456@tcp(localhost:3306)/ductestmysql?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("ket noi that bai")
		fmt.Println(err)
	} else {
		fmt.Println("ket noi thanh cong")
	}
	db.AutoMigrate(&models.Address{}, &models.CreditCard{}, &models.Email{}, &models.Language{}, &models.User{})

	defer func() {
		fmt.Scanf("%s")
		db.Close()
		fmt.Println("关闭连接")
	}()
}
