package routes

import (
	"github.com/YehorChervonyi/SimpleAPI/internal/handlers"
	"github.com/YehorChervonyi/SimpleAPI/internal/middleware"
	"github.com/julienschmidt/httprouter"
)

func InitRoutes(router *httprouter.Router) {
	handlers.InitHandlers()
	allHandlers := handlers.GetHandlers()
	//POST
	router.POST("/calculate", middleware.ValidateInput(allHandlers.FactorialHandler.CalculateFactorial))

}
