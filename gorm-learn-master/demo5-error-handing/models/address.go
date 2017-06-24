package models

// Address 地址
type Address struct {
	ID       int
	Address1 string `gorm:"not null;unique"` // 设置此列 不能为空( not nullable ) 和 唯一( unique )
	Address2 string `gorm:"type:varchar(100);unique"`
	Post     string
}
