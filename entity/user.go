package entity

import (
	"time"
)

type User struct {
	ID             string    `gorm:"primary_key" json:"id"`
	FirstName      string    `json:"first_name"`
	LastName       string    `json:"last_name"`
	Password       string    `json:"password"`
	Username       *string   `gorm:"unique;not null" json:"username"`
	AccountCreated time.Time `json:"account_created"`
	AccountUpdated time.Time `json:"account_updated"`
}
