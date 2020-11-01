package service

import (
	"github.com/danikg/go-todo-rest-api/internal/models"
)

// TodoItemService ...
type TodoItemService struct {
	TodoItemRepo models.ITodoItemRepository
	TodoListRepo models.ITodoListRepository
}

// NewTodoItemService ...
func NewTodoItemService(todoItemRepo models.ITodoItemRepository, todoListRepo models.ITodoListRepository) *TodoItemService {
	return &TodoItemService{
		TodoItemRepo: todoItemRepo,
		TodoListRepo: todoListRepo,
	}
}

// GetAll returns all todo items by todo list id
func (t *TodoItemService) GetAll(listID uint) ([]models.TodoItem, error) {
	todoList, err := t.TodoListRepo.GetSingle(listID)
	if err != nil {
		return []models.TodoItem{}, err
	}
	return t.TodoItemRepo.GetAll(todoList.ID)
}

// GetSingle returns a todo item by id
func (t *TodoItemService) GetSingle(id uint) (models.TodoItem, error) {
	return t.TodoItemRepo.GetSingle(id)
}

// Create creates a new todo item
func (t *TodoItemService) Create(listID uint, todoItem *models.TodoItem) error {
	return t.TodoItemRepo.Create(listID, todoItem)
}

// Update updates the todo item
func (t *TodoItemService) Update(id uint, todoItemData *models.TodoItem) (models.TodoItem, error) {
	return t.TodoItemRepo.Update(id, todoItemData)
}

// Delete removes the todo item
func (t *TodoItemService) Delete(id uint) error {
	return t.TodoItemRepo.Delete(id)
}
