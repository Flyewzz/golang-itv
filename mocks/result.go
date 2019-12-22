package mocks

import "github.com/Flyewzz/golang-itv/models"

func GetMockStandardSuccResult() *models.Result {
	return &models.Result{
		Response: GetMockStandardSuccResponse(),
		Error:    nil,
	}
}

func GetMockStandardFailResult() *models.Result {
	return &models.Result{
		Response: nil,
		Error:    FAIL_EXECUTOR_ERROR,
	}
}
