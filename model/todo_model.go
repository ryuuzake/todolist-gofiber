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
	Id         uuid.UUID `db:"id" json:"id"`
	TodolistId uuid.UUID `db:"todolist_id" json:"todolist_id"`
	Url        string    `db:"url" json:"url"`
	Caption    string    `db:"caption" json:"caption"`
	CreatedAt  time.Time `db:"created_at" json:"created_at"`
	UpdatedAt  time.Time `db:"updated_at" json:"updated_at"`
}
