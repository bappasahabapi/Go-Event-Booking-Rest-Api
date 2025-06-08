package services

import (
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