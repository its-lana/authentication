package repository

import (
	"authentication-modules/config"
	"authentication-modules/dto"
	"authentication-modules/models"
)

func CreateUser(req dto.ReqUser) (*models.User, error) {
	newUser := req.ToUser()
	if err := config.DB.Debug().Create(&newUser).Error; err != nil {
		return nil, err
	}
	return &newUser, nil
}

func RetrieveUserByID(id int) (*models.User, error) {
	var user models.User
	if err := config.DB.Debug().First(&user, id).Error; err != nil {
		return nil, err
	}
	return &user, nil
}
