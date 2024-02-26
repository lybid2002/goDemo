package models

type User struct {
	Id       int
	Username string
	Password string
}

type UserInput struct {
	Username string `json:"name" binding:"required,min=4"`
	Email    string `json:"email" binding:"required,email,uniqueEmail"`
	Password string `json:"password" binding:"required,min=8"`
}

func CreateUser(user *User) error {
	return nil
}

func AuthenticateUser(user *User) (*User, error) {
	return &User{}, nil
}
