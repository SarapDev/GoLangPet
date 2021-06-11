package account

import "context"

type User struct {
	ID 			string `json:"id,omitempty"`
	Email 		string `json:"email"`
	Password 	string `json:"password"`
}

type Repository interface {
	CreateUser(context context.Context, user User) error
	GetUser(context context.Context, id string) (string, error)
}