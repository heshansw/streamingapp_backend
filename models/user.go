package models

type User struct {
	ID       int    `int:"id"`
	Username string `string:"username"`
	Password string `string:"password"`
}
