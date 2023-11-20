package utils

import (
	"easygoals/model"
	"github.com/gin-gonic/gin"
	"log"
)

func UserValidationRegister(c *gin.Context) {
	var userModel = &model.User{}
	err := c.BindJSON(userModel)
	if err != nil {
		log.Fatal(err.Error())
	}
	var dd = validateUserData(userModel)
	if dd != "" {
		c.JSON(400, model.BaseModel{
			Status:  false,
			Message: dd,
			Code:    "REGISTER_API",
		})
		c.Abort()
	}
	c.Next()
}

func validateUserData(myUSer *model.User) string {

	if myUSer.Name == "" {
		return "Name is Empty"
	}
	if myUSer.Email == "" {
		return "Email not be null"
	}
	if myUSer.Password == "" {
		return "Password not be null"
	}

	return ""
}
