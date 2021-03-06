package webservices

import (
	"github.com/danikg/go-todo-rest-api/models"
	repos "github.com/danikg/go-todo-rest-api/repositories"
)

// TodoListService ...
type TodoListService struct {
	UserRepo     repos.IUserRepository
	TodoListRepo repos.ITodoListRepository
}

// NewTodoListService ...
func NewTodoListService(userRepo repos.IUserRepository, todoListRepo repos.ITodoListRepository) *TodoListService {
	return &TodoListService{
		UserRepo:     userRepo,
		TodoListRepo: todoListRepo,
	}
}

// GetAll returns all todo lists by user id
func (t *TodoListService) GetAll(userID uint) ([]models.TodoList, error) {
	user, err := t.UserRepo.GetSingle(userID)
	if err != nil {
		return []models.TodoList{}, err
	}

	return t.TodoListRepo.GetAll(user.ID)
}

// GetSingle returns a todo list by id
func (t *TodoListService) GetSingle(id uint) (models.TodoList, error) {
	return t.TodoListRepo.GetSingle(id)
}

// Create creates a new todo list
func (t *TodoListService) Create(userID uint, todoList *models.TodoList) error {
	return t.TodoListRepo.Create(userID, todoList)
}

// Update updates the todo list
func (t *TodoListService) Update(id uint, todoListData *models.TodoList) (models.TodoList, error) {
	return t.TodoListRepo.Update(id, todoListData)
}

// Delete removes the todo list
func (t *TodoListService) Delete(id uint) error {
	return t.TodoListRepo.Delete(id)
}
