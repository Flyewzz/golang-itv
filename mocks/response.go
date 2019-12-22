package mocks

import (
	"github.com/Flyewzz/golang-itv/models"
)

func GetMockStandardSuccResponse() *models.Response {
	return &models.Response{
		Status:        "200 OK",
		Headers:       "Header1: Header1",
		Body:          "body",
		ContentLength: 4,
	}
}
