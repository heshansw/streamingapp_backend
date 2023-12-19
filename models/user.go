package models

type Users struct {
	UserId   int    `json:"user_id" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserInput struct {
	Username string `json:"username" validate:"required"`
	Password string `json:"password" validate:"required"`
}
