package models

type User struct {
	ID       int
	Username string `json:"username"`
	Password string
}
