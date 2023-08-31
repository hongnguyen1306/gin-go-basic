package controllers

import (
	"app/db"
	"app/models"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-pg/pg/v10"
	"github.com/go-pg/pg/v10/orm"
	guuid "github.com/google/uuid"
)

func CreateUserTable(db *pg.DB) error {
	model := []interface{}{
		(*models.User)(nil),
	}
	err := db.Model(model).CreateTable(&orm.CreateTableOptions{
		Temp:        false,
		IfNotExists: true,
	})

	if err != nil {
		log.Printf("Error while creating user table, Reason: %v\n", err)
		return err
	}
	log.Printf("User table created")
	return nil
}

func CreateUser(c *gin.Context) {
	dbConnect := db.ConnectDatabase()
	var user models.User
	c.BindJSON(&user)
	fullName := user.FullName
	employeeCode := user.EmployeeCode
	email := user.Email
	role := user.Role
	id := guuid.New().String()

	_, insertError := dbConnect.Model(&models.User{
		Id:           id,
		FullName:     fullName,
		EmployeeCode: employeeCode,
		Email:        email,
		Role:         role,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}).Insert()
	if insertError != nil {
		log.Printf("Error while inserting new user into db, Reason: %v\n", insertError)
		c.JSON(http.StatusInternalServerError, gin.H{
			"status":  http.StatusInternalServerError,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusCreated, gin.H{
		"status":  http.StatusCreated,
		"message": "User created Successfully",
	})
}

func GetSingleUser(c *gin.Context) {
	dbConnect := db.ConnectDatabase()
	userId := c.Param("userId")
	user := &models.User{Id: userId}
	err := dbConnect.Model(user).WherePK().Select()

	if err != nil {
		log.Printf("Error while getting a single user, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "User not found",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "Single User",
		"data":    user,
	})
}

func GetAllUsers(c *gin.Context) {
	dbConnect := db.ConnectDatabase()
	var user []models.User
	err := dbConnect.Model(user).WherePK().Select()

	if err != nil {
		log.Printf("Error while getting all users, Reason: %v\n", err)
		c.JSON(http.StatusNotFound, gin.H{
			"status":  http.StatusNotFound,
			"message": "Something went wrong",
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status":  http.StatusOK,
		"message": "All Users",
		"data":    user,
	})
}
