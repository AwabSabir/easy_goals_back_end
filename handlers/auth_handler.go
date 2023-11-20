package handlers

import (
	"database/sql"
	"easygoals/db"
	"easygoals/model"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func LoginUser(c *gin.Context) {
	response := model.BaseModel{
		Status:  true,
		Message: "Login sucessfully",
		Code:    "LOGIN_API",
		Data:    map[string]string{"name": "awab Sabir"},
	}
	c.JSON(http.StatusOK, response)
}

func RegisterUser(c *gin.Context) {
	var userModel = &model.User{}
	err := c.BindJSON(userModel)
	if err != nil {
		log.Println(err.Error())
	}

	database := db.ConnectDb()
	defer func(database *sql.DB) {
		err := database.Close()
		if err != nil {
			log.Panicf(err.Error())
		}
	}(database)

	response := model.BaseModel{
		Status:  true,
		Message: "Register sucessfuly please varify",
		Code:    "REGISTER_API",
	}
	c.JSON(http.StatusOK, response)
}

func GetAllUser(c *gin.Context) {
	databasee := db.ConnectDb()
	defer func(databasee *sql.DB) {
		err := databasee.Close()
		if err != nil {
			log.Panicf(err.Error())
		}
	}(databasee)

	var data, err = databasee.Query("SELECT * FROM `users`")
	if err != nil {
		log.Panicf(err.Error())
	}
	defer data.Close()
	var usrsList []model.User
	for data.Next() {
		var (
			id       int
			fName    string
			lName    string
			email    string
			password string
		)
		if err := data.Scan(&id, &fName, &lName, &email, &password); err != nil {
			log.Panicf(err.Error())
		}
		usrsList = append(usrsList, model.User{
			Name:     fName,
			Email:    email,
			Password: password,
		})
	}
	c.JSON(http.StatusOK, model.BaseModel{
		Status:  true,
		Code:    "GET_ALL_USERS",
		Message: "All user fetch sucessfully",
		Data:    usrsList,
	})

}
