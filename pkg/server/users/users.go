package users

import "social/pkg/server"

type User struct {
	server *server.Server
}

func InitUsers(s server.Server) *User {
	return &User{
		server: &s,
	}
}
