package entity

import (
	"time"
)

type File struct {
	ID            string    `gorm:"primary_key" json:"file_id"`
	FileName      string    `json:"file_name"`
	S3ObjectName  string    `json:"s3_object_name"`
	CreatedDate   time.Time `json:"created_date"`
	IsQuestion    bool      `json:"is_question"`    //true: question; false; answer
	AcceptRanges  string    `json:"accept_ranges"`  //metadata from s3
	ContentLength int64     `json:"content_length"` //metadata from s3
	ContentType   string    `json:"content_type"`   //metadata from s3
	ETag          string    `json:"e_tag"`          //metadata from s3
	LastModified  time.Time `json:"last_modified"`  //metadata from s3
}
