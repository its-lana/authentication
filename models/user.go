package models

import "time"

type User struct {
	ID        int       `gorm:"primaryKey;autoIncrement" json:"id"`
	FullName  string    `gorm:"size:64" json:"full_name"`
	Email     string    `gorm:"size:64" json:"email"`
	Password  string    `gorm:"size:128" json:"-"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
