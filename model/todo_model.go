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
	Id        uuid.UUID `db:"id" json:"id"`
	TodoId    uuid.UUID `db:"todo_id" json:"todo_id"`
	StatusId  uuid.UUID `db:"status_id" json:"status_id"`
	Task      string    `db:"task" json:"task"`
	CreatedAt time.Time `db:"created_at" json:"created_at"`
	UpdatedAt time.Time `db:"updated_at" json:"updated_at"`
}

type Status struct {
	Id   uuid.UUID `db:"id" json:"id"`
	Name string    `db:"name" json:"name"`
}

type Attachment struct {
	Id         int
	TodolistId int
	Url        string
	Caption    string
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
