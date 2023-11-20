package handlers

import (
	"easygoals/model"
	"github.com/gin-gonic/gin"
	"net/http"
)

func DefaultHandler(c *gin.Context) {
	c.String(http.StatusOK, "Server is Running")
}

func NoRouteFound(c *gin.Context) {
	var baseModel = model.BaseModel{
		Status:  true,
		Message: "No Route Found",
	}

	c.JSON(http.StatusNotFound, baseModel)
}
