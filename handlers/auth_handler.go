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
)

func LoginUser(c *gin.Context) {
	var request = &model.LoginModel{}
	err := c.ShouldBindJSON(request)
	if err != nil {
		c.JSON(400, model.BaseModel{
			Status:  false,
			Message: err.Error(),
			Code:    "LOGIN_API",
		})
		log.Panic(err.Error())
		return
	}
	log.Print(request)
	myDb := db.ConnectDb()
	defer myDb.Close()

	query := "SELECT id,fName, email, created_at FROM `users` WHERE email=? AND password = ?"
	var (
		id        int
		fName     string
		email     string
		createdAt string
	)
	err = myDb.QueryRow(query, request.Email, request.Password).Scan(&id, &fName, &email, &createdAt)
	if err != nil {
		c.JSON(400, model.BaseModel{
			Status:  false,
			Message: "User not found",
			Code:    "LOGIN_API",
		})
		log.Panic(err.Error())
		return
	}
	c.JSON(http.StatusOK, model.BaseModel{
		Status:  true,
		Message: "Login Sucessfuly",
		Code:    "LOGIN API",
		Data: gin.H{
			"id":        id,
			"Name":      fName,
			"email":     email,
			"createdAt": createdAt,
		},
	})
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
		userFounded, _ := findUser(userModel.Email)
		if userFounded {
			response := model.BaseModel{
				Status:  false,
				Message: "User Already Registerd",
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
					Data: map[string]any{
						"name":      user.Name,
						"email":     user.Email,
						"createdAt": user.CreatedAt,
					},
				}
				c.JSON(http.StatusOK, response)
			} else {
				response := model.BaseModel{
					Status:  false,
					Message: "some thing went wrrong",
					Code:    "REGISTER_API",
				}
				c.JSON(400, response)
			}
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
			createdAt string
		)
		err = database.QueryRow("SELECT  id,fName, email, created_at FROM `users` WHERE `id` = ?", lastInsertId).Scan(&id, &fName, &email, &createdAt)
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

	var data, err = databasee.Query("SELECT id,fName, email, created_at From users; ")
	if err != nil {
		log.Panicf(err.Error())
	}
	defer data.Close()
	var usrsList []model.User
	for data.Next() {
		var (
			id    int
			fName string
			email string
		)
		if err := data.Scan(&id, &fName, &email); err != nil {
			log.Panicf(err.Error())
		}
		usrsList = append(usrsList, model.User{
			Name:  fName,
			Email: email,
		})
	}
	c.JSON(http.StatusOK, model.BaseModel{
		Status:  true,
		Code:    "GET_ALL_USERS",
		Message: "All user fetch sucessfully",
		Data:    usrsList,
	})

}

func findUser(email string) (bool, model.User) {
	myDb := db.ConnectDb()
	defer myDb.Close()
	query := "SELECT fName, email, created_at FROM users WHERE email=?"
	var (
		fName     string
		userEmail string
		createdAt *string
	)
	data, err := myDb.Query(query, email)
	if err != nil {
		log.Fatal(err.Error())
	}
	if data.Next() {
		data.Scan(&fName, &userEmail, &createdAt)
	} else {
		return false, model.User{}
	}
	if userEmail != "" {
		return true, model.User{
			Email:     userEmail,
			Name:      fName,
			CreatedAt: createdAt,
		}
	}
	return false, model.User{}
}

func LoginUserDatabase(emai, password string) {
	database := db.ConnectDb()
	defer database.Close()

}
