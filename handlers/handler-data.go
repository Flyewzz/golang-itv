package handlers

import (
	. "github.com/Flyewzz/golang-itv/interfaces"
)

type HandlerData struct {
	Executor        Executor
	StoreController Store
	Dispatcher      Dispatcher
}

func NewHandlerData(ex Executor, sc Store, d Dispatcher) *HandlerData {
	return &HandlerData{
		Executor:        ex,
		StoreController: sc,
		Dispatcher:      d,
	}
}
