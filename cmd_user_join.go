package main

import (
	"github.com//"
)

type UserJoin struct {
	user *User
}

func (client *Client) CmdUserJoin() *User {
	user := &User{}
	return user
}
