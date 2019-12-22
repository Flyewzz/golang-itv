package features

import (
	"github.com/Flyewzz/golang-itv/workers/models"
)

// !IMPORTANT!
// To use this function to work only with bufferized channels
func IsClosed(ch chan models.Job) bool {
	_, ok := <-ch
	return (ok == false)
}
