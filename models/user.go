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

type ReqUser struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

func (req *ReqUser) ToUser() User {
	return User{
		FullName: req.FullName,
		Email:    req.Email,
		Password: req.Password,
	}
}
