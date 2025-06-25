package models

import (
	"plashoes-server/db"
)

func (OTPRequest OTP) Save() error {

	var query string

	emailInDatabase,err := checkEmailAdded(OTPRequest.Email)

	if err != nil {
		return err
	}

	if emailInDatabase {
		query = "UPDATE OTPs SET code = ? WHERE email = ?"
	}else{
		query = "INSERT INTO OTPs (code, email) VALUES (?,?)"
	}

	_, err = db.DB.Exec(query, OTPRequest.Code, OTPRequest.Email, )

	return err
}
func (OTPRequest OTP) OTPCorrect() (bool, error) {
	query := "SELECT COUNT(1) FROM OTPs WHERE email = ? and code = ?"

	var count int

	err := db.DB.QueryRow(query, OTPRequest.Email, OTPRequest.Code).Scan(&count)

	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}else{
		return false,nil
	}
}

func DeleteOTP(email string)error{
	query := "DELETE FROM OTPs WHERE email = ?"
	
	_,err := db.DB.Exec(query, email)

	return err
}

func checkEmailAdded(email string) (bool, error){
	query := "SELECT COUNT(1) FROM OTPs WHERE email = ?"

	var count int

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