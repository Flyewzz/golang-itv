package interfaces

import (
	"net/http"

	"github.com/Flyewzz/golang-itv/models"
)

type Executor interface {
	Execute(client *http.Client, method, url string, id int) (*models.Response, error)
}
