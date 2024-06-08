package utils

import (
	"errors"
	"hello/models/db"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type UserClaims struct {
	Mail     string
	Usertype string
}

const secretkey = "supersecret"

func GenerateToken(mail string, username string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"mail":     mail,
		"username": username,
		"exp":      time.Now().Add(time.Hour * 2).Unix(),
	})
	return token.SignedString([]byte(secretkey))
}

func Validate(token string) (UserClaims, error) {
	parsedtoken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretkey), nil
	})
	if err != nil {
		return UserClaims{}, errors.New("could not parse")
	}
	if !parsedtoken.Valid {
		return UserClaims{}, errors.New("invalid token")
	}

	claims, ok := parsedtoken.Claims.(jwt.MapClaims)
	if !ok {
		return UserClaims{}, errors.New("invalid claims")
	}

	email, ok := claims["mail"].(string)
	if !ok {
		return UserClaims{}, errors.New("invalid email claim")
	}
	query := "SELECT usertype FROM users WHERE mail = ?"
	row := db.DB.QueryRow(query, email)
	var usertype string
	if err := row.Scan(&usertype); err != nil {
		return UserClaims{}, errors.New("credentials invalid")
	}

	return UserClaims{Mail: email, Usertype: usertype}, nil
}
