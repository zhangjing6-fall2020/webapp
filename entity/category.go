package entity

type Category struct {
	CategoryID string `json:"category_id"`
	Category   string `json:"category"`
	QuestionID string `json:"question_id"`
	Question   Question `gorm:"ForeignKey:ID;AssociationForeignKey:QuestionID"`
}