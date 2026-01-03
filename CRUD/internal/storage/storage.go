package storage

import "github.com/apk471/go-crud-api/internal/types"

type Storage interface{
	CreateUser(name string, email string, age int) (int64, error)
	GetUserById(id int64) (types.User, error)
	GetUser() ([]types.User, error)
}