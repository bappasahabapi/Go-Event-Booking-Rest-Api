package services

import (
	"bappa.com/rest/db"
	"bappa.com/rest/models"
)

func SaveUser (user *models.User) error{
	query :=`
	INSERT INTO users(email,password)
	VALUES (?,?)
	`
	stmt,err :=db.DB.Prepare(query)
	if err != nil {return err}

	defer stmt.Close()

	result,err:=stmt.Exec(user.Email,user.Password)
	if err != nil {return err}

	userId,err :=result.LastInsertId()
	user.ID =userId

	return err

}