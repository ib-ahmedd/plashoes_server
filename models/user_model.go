package models

import (
	"plashoes-server/db"
	"plashoes-server/utils"
)

func (user User) Save() (User,error) {
	query := "INSERT INTO users(user_name, email, password, mobile_no, date_of_birth, gender, country, postal_code, address, country_code) VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := db.DB.Prepare(query)

	if err != nil {
		return User{},err
	}

	defer stmt.Close()

	hashedPassword, err := utils.HashPassword(user.Password)

	if err != nil {
		return User{},err
	}

	result, err := stmt.Exec(user.User_name, user.Email, hashedPassword, user.Mobile_no, user.Date_of_birth, user.Gender, user.Country, user.Postal_code, user.Address, user.Country)

	if err != nil {
		return User{},err
	}

	userID, err := result.LastInsertId()

	user.ID = userID

	return user,err
}

func CheckUserExists(email string) (bool, error) {
	var count int
	query := "SELECT COUNT(1)FROM users WHERE email = ? "

	err := db.DB.QueryRow(query, email).Scan(&count)

	if err != nil {
		return true, err
	}

	if count > 0 {
		return true, nil
	}else{
		return false,nil
	}
}