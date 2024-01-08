package models

import (
	"errors"

	"example.com/rest-api/db"
	"example.com/rest-api/util"
)

type User struct {
	ID       int64
	Email    string `binding:"required"`
	Password string `binding:"required"`
}

func (user *User) Save() error {
	query := `
	INSERT INTO users (email,password) 
	VALUES (?,?)
	`

	stmt, err := db.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return err
	}
	result, err := stmt.Exec(user.Email, hashedPassword)
	if err != nil {
		return err
	}
	_, err = result.LastInsertId()
	return err
}

func (u *User) ValidateCredentials() error {
	query := `SELECT id,password FROM users where email=?`
	row := db.DB.QueryRow(query, u.Email)
	var retrievedPassword string
	err := row.Scan(&u.ID, &retrievedPassword)
	if err != nil {
		return errors.New("credentials are not valid! ")
	}
	passwordIsValid := util.CheckPasswordHashed(u.Password, retrievedPassword)
	if !passwordIsValid {
		return errors.New("credentials are not valid! ")
	}
	return nil
}
