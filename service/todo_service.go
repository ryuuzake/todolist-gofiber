package service

import "github.com/ryuuzake/todolist-gofiber/model"

type TodoService interface {
	GetAllTodo() ([]model.Todo, error)
	GetTodoById(id int) (model.Todo, error)
	CreateTodo(todo model.Todo) error
	UpdateTodo(id int, todo model.Todo) error
	DeleteTodo(id int) error
	CreateTodolistWithTodoId(id int, todolist model.Todolist) error
	UpdateTodolistWithTodoId(id int, todolist model.Todolist) error
	CreateAttachmentWithTodolistId(id int, attachment model.Attachment) error
	UpdateAttachmentWithId(id int, attachment model.Attachment) error
}
