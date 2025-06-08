package utils

import "golang.org/x/crypto/bcrypt"

func HashPassword(plainText string)(string,error){
	bytes, err :=bcrypt.GenerateFromPassword([]byte(plainText),14)
	return string(bytes),err
}


func ComparePassword(plainTextPassword, hashedPassowrd string)bool{
	err:=bcrypt.CompareHashAndPassword([]byte(hashedPassowrd),[]byte(plainTextPassword))
	return err ==nil // no error that means valid return true
}