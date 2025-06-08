package services

import (
	"errors"

	"bappa.com/rest/db"
	"bappa.com/rest/models"
	"bappa.com/rest/utils"
)

func SaveUser (user *models.User) error{
	query :=`
	INSERT INTO users(email,password)
	VALUES (?,?)
	`
	stmt,err :=db.DB.Prepare(query)
	if err != nil {return err}

	defer stmt.Close()

	hashedPassowrd,err :=utils.HashPassword(user.Password)
	if err != nil {return err}

	// result,err:=stmt.Exec(user.Email,user.Password)
	result,err:=stmt.Exec(user.Email,hashedPassowrd)
	if err != nil {return err}

	userId,err :=result.LastInsertId()
	user.ID =userId

	return err

}


func ValidateCredentials(user *models.User)error{
	query :=`SELECT password FROM users WHERE email=?`

	row:=db.DB.QueryRow(query,user.Email)

	var retrivedPassword string
	err:=row.Scan(&retrivedPassword)
	if err != nil {return errors.New("Credentials are not valid")}

	//compare password
	passwordIsValid :=utils.ComparePassword(user.Password,retrivedPassword)

	if !passwordIsValid {
		return errors.New("Credentials are not valid")
	}

	return nil; //valid


}