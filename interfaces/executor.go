package interfaces

import (
	"net/http"

	"github.com/Flyewzz/golang-itv/models"
)

type Executor interface {
	Execute(client *http.Client, task *models.Task) (*models.Response, error)
}
