package entity

import (
	"time"
)

type File struct {
	ID           string    `gorm:"primary_key" json:"file_id"`
	FileName     string    `json:"file_name"`
	S3ObjectName string    `json:"s3_object_name"`
	CreatedDate  time.Time `json:"created_date"`
	IsQuestion   bool      `json:"is_question"` //true: question; false; answer
}
