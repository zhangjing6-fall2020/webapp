package entity

import (
	"time"
)

type Question struct {
	ID               string    `json:"question_id"`
	CreatedTimestamp time.Time `json:"question_created_timestamp"`
	UpdatedTimestamp time.Time `json:"question_updated_timestamp"`
	QuestionText             string    `json:"question_text"`
	UserID                   string    `json:"user_id"`
	User                     User `gorm:"ForeignKey:ID;AssociationForeignKey:UserID"`
}
