package routes

import (
	"easygoals/handlers"
	"github.com/gin-gonic/gin"
)

func AuthRouter(r *gin.Engine) {
	r.POST("/v1/login", handlers.LoginUser)
	//r.Use(utils.UserValidationRegister)
	r.POST("/v1/register", handlers.RegisterUser)

}
