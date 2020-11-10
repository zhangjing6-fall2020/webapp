package entity

type Category struct {
	ID       string `gorm:"primary_key" json:"category_id"`
	Category string `gorm:"unique;not null" json:"category"`
}
