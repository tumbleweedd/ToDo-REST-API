package service

import (
	"github.com/tumbleweedd/todo-app"
	"github.com/tumbleweedd/todo-app/pkg/repository"
)

type Authorization interface {
	CreateUser(user todo.User) (int, error)
	GenerateToken(username, password string) (string, error)
	ParseToken(token string) (int, error)
}
type TodoList interface {
	Create(userId int, list todo.TodoList) (int, error)
	GetAll(userId int) ([]todo.TodoList, error)
	GetById(userId int, listId int) (todo.TodoList, error)
	Delete(userId int, listId int) error
	Update(userId, listId int, input todo.UpdateListInput) error
}
type TodoItem interface {
	Create(userId int, listId int, item todo.TodoItem) (int, error)
	GetAll(userId, listId int) ([]todo.TodoItem, error)
	GetById(userId int, itemId int) (todo.TodoItem, error)
	Update(userId, listId int, input todo.UpdateItemInput) error
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
