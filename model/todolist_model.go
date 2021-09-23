package model

import "time"

type Todolist struct {
	Id        int
	Task      string
	CreatedAt time.Time
	UpdatedAt time.Time
}
