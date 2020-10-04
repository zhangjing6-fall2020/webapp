package entity

import (
	"time"
)

type Answer struct {
	ID               string    `gorm:"primary_key" json:"answer_id"`
	CreatedTimestamp time.Time `json:"created_timestamp"`
	UpdatedTimestamp time.Time `json:"updated_timestamp"`
	AnswerText       string    `json:"answer_text"`
	UserID           string    `json:"user_id"`
	User             User      `gorm:"ForeignKey:ID;AssociationForeignKey:UserID" json:"-"`
	QuestionID       string    `json:"question_id"`
	Question         Question  `gorm:"ForeignKey:ID;AssociationForeignKey:QuestionID" json:"-"`
}