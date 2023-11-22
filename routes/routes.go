package routes

import (
	"easygoals/handlers"
	"github.com/gin-gonic/gin"
)

func Routes(r *gin.Engine) {
	r.GET("/", handlers.DefaultHandler)
	r.GET("/v1/Users", handlers.GetAllUser)
	r.NoRoute(handlers.NoRouteFound)
}
