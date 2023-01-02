package schemas

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Fullname string `json:"fullname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}
