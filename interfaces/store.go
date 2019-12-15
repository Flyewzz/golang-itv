package interfaces

import (
	"github.com/Flyewzz/golang-itv/models"
)

type Store interface {
	Add(request *models.Request) int
	GetAll() []models.Request
	GetByPage(page int) ([]models.Request, error)
	GetById(id int) (*models.Request, error)
	RemoveById(id int) error
	RemoveAll()
}
