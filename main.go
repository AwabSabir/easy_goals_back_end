package main

import (
	"easygoals/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.AuthRouter(router)
	routes.Routes(router)
	err := router.Run("127.0.0.1:8080")

	if err != nil {
		return
	}
}
