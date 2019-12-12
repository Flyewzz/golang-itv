package handlers

import (
	. "github.com/Flyewzz/golang-itv/interfaces"
)

type UserHandler struct {
	Executor        Executor
	StoreController Store
}

func NewUserHandler(ex Executor, sc Store) *UserHandler {
	return &UserHandler{
		Executor:        ex,
		StoreController: sc,
	}
}
