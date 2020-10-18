package entity

import (
	"time"
)

type Question struct {
	ID               string     `gorm:"primary_key" json:"question_id"`
	CreatedTimestamp time.Time  `json:"created_timestamp"`
	UpdatedTimestamp time.Time  `json:"updated_timestamp"`
	QuestionText     string     `json:"question_text"`
	UserID           string     `json:"user_id"`
	User             User       `gorm:"ForeignKey:ID;AssociationForeignKey:UserID" json:"-"`
	Categories       []Category `json:"categories"`
	Answers          []Answer   `json:"answers"`
	Attachments      []File     `json:"attachments"`
}
