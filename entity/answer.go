package entity

import (
	"time"
)

type Answer struct {
	ID               string    `json:"answer_id"`
	CreatedTimestamp time.Time `json:"answer_created_timestamp"`
	UpdatedTimestamp time.Time `json:"answer_updated_timestamp"`
	AnswerText             string    `json:"answer_text"`
	UserID                 string    `json:"user_id"`
	User                   User `gorm:"ForeignKey:ID;AssociationForeignKey:UserID"`
	QuestionID             string `json:"question_id"`
	Question               Question `gorm:"ForeignKey:ID;AssociationForeignKey:QuestionID"`
}