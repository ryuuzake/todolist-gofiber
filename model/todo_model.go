package model

import (
	"github.com/gofrs/uuid"
	"time"
)

type Todo struct {
	Id        uuid.UUID `db:"id" json:"id"`
	UserId    uuid.UUID `db:"user_id" json:"user_id"`
	Title     string    `db:"title" json:"title"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type Todolist struct {
	Id        int
	TodoId    int
	StatusId  int
	Task      string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Status struct {
	Id   int
	Name string
}

type Attachment struct {
	Id         int
	TodolistId int
	Url        string
	Caption    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
