package handlers

import (
	"database/sql"
	"easygoals/db"
	"easygoals/model"
	"easygoals/utils"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"time"
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
	log.Println("Register call")
	var userModel = &model.User{}
	err := c.BindJSON(userModel)
	if err != nil {
		log.Println(err.Error())
	}
	validate := utils.ValidateUserData(userModel)
	if validate != "" {
		response := model.BaseModel{
			Status:  false,
			Message: validate,
			Code:    "REGISTER_API",
		}
		c.JSON(400, response)
	} else {
		added, user := userInsertIntoDb(userModel)
		if added {
			log.Println(userModel)
			response := model.BaseModel{
				Status:  true,
				Message: "Register sucessfuly please varify",
				Code:    "REGISTER_API",
				Data:    user,
			}
			c.JSON(http.StatusOK, response)
		} else {
			response := model.BaseModel{
				Status:  false,
				Message: validate,
				Code:    "REGISTER_API",
			}
			c.JSON(400, response)
		}
	}

}

func userInsertIntoDb(userModel *model.User) (bool, model.User) {
	var isDataAdded = false

	database := db.ConnectDb()
	defer func(database *sql.DB) {
		err := database.Close()
		if err != nil {
			log.Panicf(err.Error())
		}
	}(database)
	query := "INSERT INTO `users`(`fName`, `email`, `password`) VALUES (?,?,?)"
	insert, err := database.Prepare(query)
	if err != nil {
		log.Panic(err.Error())
	}

	res, err := insert.Exec(userModel.Name, userModel.Email, userModel.Password)
	if err != nil {
		log.Fatal(err.Error())
	}
	lastInsertId, err := res.LastInsertId()
	var registerdUSer = model.User{}
	if lastInsertId != 0 {
		isDataAdded = true
		var (
			id        int
			fName     string
			email     string
			password  string
			createdAt string
		)
		err = database.QueryRow("SELECT * FROM `users` WHERE `id` = ?", lastInsertId).Scan(&id, &fName, &email, &password, &createdAt)
		if err != nil {
			log.Panic(err.Error())
		}
		registerdUSer = model.User{
			Name:      fName,
			Email:     email,
			CreatedAt: &createdAt,
		}
	}
	if err != nil {
		isDataAdded = false
		fmt.Println(err.Error())
	}
	defer insert.Close()
	return isDataAdded, registerdUSer
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
			id        int
			fName     string
			email     string
			password  string
			createdAt time.Time
		)
		if err := data.Scan(&id, &fName, &email, &password, &createdAt); err != nil {
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
