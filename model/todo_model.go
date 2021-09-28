package model

import "time"

type Todo struct {
	Id        int
	UserId    int
	Title     string
	CreatedAt time.Time
	UpdatedAt time.Time
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
