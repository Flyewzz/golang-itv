package mocks

import "errors"

var (
	FAIL_EXECUTOR_ERROR   error = errors.New("Some error... :(")
	STORE_INCORRECT_INDEX error = errors.New("Incorrect index")
)
