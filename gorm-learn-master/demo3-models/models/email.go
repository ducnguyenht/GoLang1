package models

// Email 邮箱
type Email struct {
	ID         int
	UserID     int    `gorm:"index"`                          // 外键 (belongs to), `index` 为此列创建索引
	Email      string `gorm:"type:varchar(100);unique_index"` // `type` 设置 sql 类型, `unique_index` 为此列创建唯一索引
	Subscribed bool
}
