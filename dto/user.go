package dto

import "authentication-modules/models"

type ReqUser struct {
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"-"`
}

func (req *ReqUser) ToUser() models.User {
	return models.User{
		FullName: req.FullName,
		Email:    req.Email,
		Password: req.Password,
	}
}
