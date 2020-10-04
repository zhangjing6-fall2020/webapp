package entity

type QuestionCategory struct {
	QuestionID string   `gorm:"primary_key" json:"question_id"`
	Question   Question `gorm:"ForeignKey:ID;AssociationForeignKey:QuestionID"  json:"question"`
	CategoryID string   `gorm:"primary_key" json:"category_id"`
	Category   Question `gorm:"ForeignKey:ID;AssociationForeignKey:CategoryID"  json:"category"`
}