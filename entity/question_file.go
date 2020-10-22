package entity

type QuestionFile struct {
	ID         string   `gorm:"primary_key" json:"file_id"`
	File       File     `gorm:"ForeignKey:ID;AssociationForeignKey:ID"  json:"-"`
	QuestionID string   `json:"question_id"`
	Question   Question `gorm:"ForeignKey:ID;AssociationForeignKey:QuestionID"  json:"-"`
}
