package main

import (
	"time"
)

type User struct {
	Id       int64
	UserType int
	LoggedIn bool
	LoggedAt time.Time
}

func NewUser() *User {
	return &User{}
}

func Join() *User {
	return &User{}
}

func Login(user *User) *User {
	user.LoggedAt = time.Now()
	user.LoggedIn = true
	return user
}

func Logout(user *User) *User {
	user.LoggedIn = false
	return user
}
