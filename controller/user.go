package controller

import (
	"authentication-modules/models"
	"authentication-modules/repository"
	"net/http"

	"github.com/gin-gonic/gin"
)

func RegisterUser(c *gin.Context) {
	var reqUser models.ReqUser
	err := c.Bind(&reqUser)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": "binding error, bad request"})
		return
	}
	user, err := repository.CreateUser(reqUser)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "register user successfully", "data": user})
}

func LoginUser(c *gin.Context) {

}
