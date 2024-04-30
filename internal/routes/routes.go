package routes

import (
	"github.com/YehorChervonyi/SimpleAPI/internal/handlers"
	"github.com/YehorChervonyi/SimpleAPI/internal/middleware"
	"github.com/julienschmidt/httprouter"
)

func InitRoutes(router *httprouter.Router) {
	//POST
	router.POST("/calculate", middleware.ValidateInput(handlers.CalculateFactorial))

}
