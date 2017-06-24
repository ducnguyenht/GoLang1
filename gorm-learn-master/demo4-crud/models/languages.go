package models

// Language 语言
type Language struct {
	ID   int
	Name string `gorm:"index:idx_name_code"` // 为此列创建索引并命名为``，如果找到其他字段定义为相同名称，将创建组合索引
	Code string `gorm:"index:idx_name_code"` // 使用 `unique_index` 也可以
}
