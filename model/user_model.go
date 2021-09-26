package model

import "time"

type User struct {
	Id             int
	FullName       string
	Email          string
	Password       string
	ForgetPassword string
	RoleId         int
	Role           Role
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

type Role struct {
	Id          int
	Name        string
	Description string
}
