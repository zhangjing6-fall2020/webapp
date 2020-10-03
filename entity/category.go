package entity

type Category struct {
	ID string `gorm:"primary_key" json:"category_id"`
	Category   string `json:"category"`
}