package models

import (
	"errors"
	"hello/models/db"
	"hello/utils"
)

type Users struct {
	Username string
	Mail     string
	Password string
}

func (u *Users) Save() error {
	query := "INSERT INTO users(username,mail,password) VALUES (?,?,?)"
	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	hashed, erro := utils.HashPassword(u.Password)
	if erro != nil {
		return erro
	}
	_, err = stmt.Exec(u.Username, u.Mail, hashed)
	if err != nil {
		return err
	}
	return nil
}

func (u *Users) Validate() error {
	query := "SELECT password FROM users WHERE mail = ?"
	row := db.DB.QueryRow(query, u.Mail)
	var retrieved string
	if err := row.Scan(&retrieved); err != nil {
		return errors.New("Credentials invalid")
	}

	passwordValid := utils.Check(u.Password, retrieved)

	if !passwordValid {
		return errors.New("Credentials invalid")
	}

	return nil
}

func Check(user Users) error {
	query := "SELECT password FROM users WHERE mail = ?"
	row := db.DB.QueryRow(query, user.Mail)
	var retrieved string
	if err := row.Scan(&retrieved); err != nil {
		return nil
	}
	return errors.New("existing account")
}
