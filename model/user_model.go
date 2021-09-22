package model

import "time"

type User struct {
	Id             int
	FullName       string
	Email          string
	Password       string
	ForgetPassword string
	CreatedAt      time.Time
	UpdatedAt      time.Time
}

// TODO: Add Role Struct
