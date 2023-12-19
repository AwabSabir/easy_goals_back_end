package main

import (
	"easygoals/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routes.AuthRouter(router)
	routes.Routes(router)
	const mobileIp = "192.168.10.11:8080"
	const localHost = "127.0.0.1:8080"
	err := router.Run(localHost)

	if err != nil {
		return
	}
}
