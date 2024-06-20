package main

import (
	"fmt"
	"hello/models"
	"hello/models/db"
	"hello/utils"
	"net"
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

const (
	NetworkIP   = "192.168.1.0"
	NetworkMask = "24"
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
	userIP := getUserIP(contest.Request)
	contest.ShouldBindJSON(&user)
	error := models.Check(user)
	if error != nil {
		contest.JSON(http.StatusUnauthorized, gin.H{"message": "existing"})
		return
	}
	err := user.Save(userIP)

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
func getUserIP(r *http.Request) string {
	if ip := r.Header.Get("X-Forwarded-For"); ip != "" {
		return ip
	}
	ip, _, _ := net.SplitHostPort(r.RemoteAddr)
	return ip
}

func isSameNetwork(userIP, networkIP, networkMask string) bool {
	_, ipnet, err := net.ParseCIDR(networkIP + "/" + networkMask)
	if err != nil {
		return false
	}

	userIPAddr := net.ParseIP(userIP)
	return ipnet.Contains(userIPAddr)
}

func signin(contest *gin.Context) {
	var user models.Users
	userID := getUserIP(contest.Request)
	if !isSameNetwork(userID, NetworkIP, NetworkMask) {
		contest.JSON(http.StatusOK, gin.H{"message": "User is on the same network"})
	}
	contest.ShouldBindJSON(&user)
	err := user.Validate()

	if err != nil {
		contest.JSON(http.StatusUnauthorized, gin.H{"message": "invalid"})
		return

	}
	token, err := utils.GenerateToken(user.Mail, user.Username)

	if err != nil {
		contest.JSON(http.StatusInternalServerError, gin.H{"message": "ise"}) //ise stands for internal server error
		return

	}
	contest.JSON(http.StatusOK, gin.H{"token": token})
}
