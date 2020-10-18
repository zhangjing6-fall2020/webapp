package entity

import "time"

//used in GetObjectMetaData, not in database
type Metadata struct {
	AcceptRanges  *string
	ContentLength *int64
	ContentType   *string
	ETag          *string
	LastModified  *time.Time
}
