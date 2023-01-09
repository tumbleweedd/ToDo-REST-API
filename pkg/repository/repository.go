package repository

import (
	"github.com/jmoiron/sqlx"
	"github.com/tumbleweedd/todo-app/pkg/model"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GetUser(username, password string) (model.User, error)
}

type TodoList interface {
	Create(userId int, list model.TodoList) (int, error)
	GetAll(id int) ([]model.TodoList, error)
	GetById(userId int, listId int) (model.TodoList, error)
	Delete(userId int, listId int) error
	Update(userId, listId int, input model.UpdateListInput) error
}

type TodoItem interface {
	Create(listId int, item model.TodoItem) (int, error)
	GetAll(userId, listId int) ([]model.TodoItem, error)
	GetById(userId, itemId int) (model.TodoItem, error)
	Update(userId, listId int, input model.UpdateItemInput) error
	Delete(userId int, itemId int) error
}

type Repository struct {
	Authorization
	TodoList
	TodoItem
}

func NewRepository(db *sqlx.DB) *Repository {
	return &Repository{
		Authorization: NewAuthPostgres(db),
		TodoList:      NewTodoListPostgres(db),
		TodoItem:      NewTodoItemPostgres(db),
	}
}
