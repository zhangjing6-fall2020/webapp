package entity

type AnswerFile struct {
	ID       string `gorm:"primary_key" json:"file_id"`
	File     File   `gorm:"ForeignKey:ID;AssociationForeignKey:ID"  json:"-"`
	AnswerID string `json:"answer_id"`
	Answer   Answer `gorm:"ForeignKey:ID;AssociationForeignKey:AnswerID"  json:"-"`
}
