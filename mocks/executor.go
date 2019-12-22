package mocks

import (
	"net/http"

	"github.com/Flyewzz/golang-itv/models"
)

type SuccessExecutor struct{}

func NewMockSuccessExecutor() *SuccessExecutor { return &SuccessExecutor{} }

func (ex *SuccessExecutor) Execute(client *http.Client, task *models.Task) (*models.Response, error) {
	return GetMockStandardSuccResponse(), nil
}

type FailExecutor struct{}

func NewMockFailExecutor() *FailExecutor { return &FailExecutor{} }

func (ex *FailExecutor) Execute(client *http.Client, task *models.Task) (*models.Response, error) {
	return nil, FAIL_EXECUTOR_ERROR
}
