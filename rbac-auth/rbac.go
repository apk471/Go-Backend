package main

type Role string

const (
	RoleAdmin Role = "admin"
	RoleUser  Role = "user"
)

var userRoles = map[string]Role{
	"admin": "admin",
	"user":  "user",
}

func getRole(username string) Role {
	return userRoles[username]
}
