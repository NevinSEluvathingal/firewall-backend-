package main

import (
	"fmt"
	"hello/models"
	"hello/models/db"
	"hello/utils"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	db.InitDB()
	server := gin.Default()

	server.Use(cors.New(cors.Config{ //used to remove cors interuptions
		AllowOrigins:     []string{"http://localhost:5173"},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	server.Use(cors.Default())
	server.POST("/signup", signup)
	server.POST("/signin", signin)
	server.GET("/getuser", userinfo)
	server.GET("/barinfo", Barinfo)
	server.Run("0.0.0.0:4000")
}

func userinfo(context *gin.Context) {
	var info utils.UserClaims
	token := context.Request.Header.Get("Authorization")
	fmt.Printf("token:%s\n", token)
	if token == "" {
		fmt.Println("blank aane tto")
		context.JSON(http.StatusOK, gin.H{"token": "blank text aane"})
		context.JSON(http.StatusUnauthorized, gin.H{"message": "not authorized"})
		return
	}
	info, err := utils.Validate(token)
	if err != nil {
		fmt.Println("token is wrong")
		context.JSON(http.StatusUnauthorized, gin.H{"message": "invalid token"})
		return
	}
	fmt.Printf("\n%s\n", info)
	context.JSON(http.StatusOK, gin.H{"userinfo": info})
}
func signup(contest *gin.Context) {
	var user models.Users
	contest.ShouldBindJSON(&user)
	err := user.Save()

	if err != nil {
		contest.JSON(http.StatusUnauthorized, gin.H{"message": "could not save"})
		return

	}

	contest.JSON(http.StatusOK, gin.H{"message": "event created....i mean entered to database"})
}
func Barinfo(context *gin.Context) {
	var BarDetails models.Bar
	email := context.Query("email")
	if email == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "email is empty"})
	}
	BarDetails, err := models.Bardetails(email)
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"message": "bar details error"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"Bar": BarDetails})
}

func signin(contest *gin.Context) {
	var user models.Users
	contest.ShouldBindJSON(&user)
	err := user.Validate()

	if err != nil {
		contest.JSON(http.StatusUnauthorized, gin.H{"message": "invalid credentials"})
		return

	}
	token, err := utils.GenerateToken(user.Mail, user.Username)

	if err != nil {
		contest.JSON(http.StatusInternalServerError, gin.H{"message": "could not"})
		return

	}
	contest.JSON(http.StatusOK, gin.H{"token": token})
}
