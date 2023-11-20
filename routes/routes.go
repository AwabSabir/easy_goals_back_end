package routes

import (
	"easygoals/handlers"
	"easygoals/utils"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.GET("/", handlers.DefaultHandler)
	r.GET("/v1/Users", handlers.GetAllUser)
	r.GET("/v1/login", handlers.LoginUser)
	r.Use(utils.UserValidationRegister)
	r.POST("/v1/register", handlers.RegisterUser)

	r.NoRoute(handlers.NoRouteFound)
}
