package interfaces

import (
	"github.com/Flyewzz/golang-itv/models"
)

type Store interface {
	Add(task *models.Task) int
	GetAll() []models.Task
	GetTasksByPage(page int) ([]models.Task, error)
	GetById(id int) (*models.Task, error)
	RemoveById(id int) error
	RemoveAll()
}
