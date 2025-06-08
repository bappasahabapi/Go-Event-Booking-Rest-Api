// package models

// type User struct {
// 	ID       int64  `json:"id"`
// 	Email    string `json:"email" binding:"required"`
// 	Password string `json:"password" binding:"required"`
// }


package models

type User struct {
	ID       int64    `json:"id"`       // optional for POST
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}
