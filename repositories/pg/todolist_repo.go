package pg

import (
	"github.com/danikg/go-todo-rest-api/models"
	"gorm.io/gorm"
)

// TodoListRepository ...
type TodoListRepository struct {
	Conn *gorm.DB
}

// NewTodoListRepository ...
func NewTodoListRepository(conn *gorm.DB) *TodoListRepository {
	return &TodoListRepository{Conn: conn}
}

// GetAll returns all todo lists by user id
func (t *TodoListRepository) GetAll(userID uint) ([]models.TodoList, error) {
	todoLists := []models.TodoList{}
	err := t.Conn.Find(&todoLists, "user_id = ?", userID).Error
	return todoLists, err
}

// GetSingle returns a todo list by id
func (t *TodoListRepository) GetSingle(id uint) (models.TodoList, error) {
	todoList := models.TodoList{}
	err := t.Conn.First(&todoList, id).Error
	return todoList, err
}

// Create creates a new todo list
func (t *TodoListRepository) Create(userID uint, todoList *models.TodoList) error {
	todoList.UserID = userID
	return t.Conn.Create(todoList).Error
}

// Update updates the todo list
func (t *TodoListRepository) Update(id uint, todoListData *models.TodoList) (models.TodoList, error) {
	todoList, err := t.GetSingle(id)
	if err != nil {
		return todoList, err
	}

	err = t.Conn.Model(&todoList).Update("Name", todoListData.Name).Error
	return todoList, err
}

// Delete removes the todo list
func (t *TodoListRepository) Delete(id uint) error {
	todoList, err := t.GetSingle(id)
	if err != nil {
		return err
	}
	return t.Conn.Unscoped().Delete(&todoList).Error
}
