package handlers

import (
	"github.com/YehorChervonyi/SimpleAPI/internal/factorial"
	"github.com/YehorChervonyi/SimpleAPI/internal/service"
)

type Handlers struct {
	FactorialHandler *FactorialHandler
}

var handler *Handlers

func InitHandlers() {
	factorialRepository := factorial.NewFactorialRepository()
	factorialService := service.NewFactorialService(factorialRepository)
	factorialHandler := NewFactorialHandler(factorialService)

	handler = &Handlers{
		FactorialHandler: factorialHandler,
	}
}

func GetHandlers() *Handlers {
	return handler
}
