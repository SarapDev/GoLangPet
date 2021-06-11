package account

import "context"

type Service interface {
	CreateUser(context context.Context, email string, password string) (string, error)
	GetUser(context context.Context, id string) (string, error)
}
