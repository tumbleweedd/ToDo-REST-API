package service

import (
	"github.com/tumbleweedd/todo-app/pkg/model"
	"github.com/tumbleweedd/todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user model.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}
type TodoList interface {
	Create(userId int, list model.TodoList) (int, error)
	GetAll(userId int) ([]model.TodoList, error)
	GetById(userId int, listId int) (model.TodoList, error)
	Delete(userId int, listId int) error
	Update(userId, listId int, input model.UpdateListInput) error
}
type TodoItem interface {
	Create(userId int, listId int, item model.TodoItem) (int, error)
	GetAll(userId, listId int) ([]model.TodoItem, error)
	GetById(userId int, itemId int) (model.TodoItem, error)
	Update(userId, listId int, input model.UpdateItemInput) error
	Delete(userId int, itemId int) error
}

type Service struct {
	Authorization
	TodoList
	TodoItem
}

func NewService(repository *repository.Repository) *Service {
	return &Service{
		Authorization: NewAuthService(repository.Authorization),
		TodoList:      NewTodoListService(repository.TodoList),
		TodoItem:      NewTodoItemService(repository.TodoItem, repository.TodoList),
	}
}
