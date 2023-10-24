package models

type User struct {
	// gorm.Model
	ID              int64
	Firstname       string
	Lastname        string
	Password        string
	Tel             string
	Lineid          string
	Email           string
	Token           string
	TokenExpireTime string
}
