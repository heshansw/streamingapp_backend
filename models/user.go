package models

type Users struct {
	ID       int    `json:"user_id" gorm:"primary_key"`
	Username string `json:"username"`
	Password string `json:"password"`
}
