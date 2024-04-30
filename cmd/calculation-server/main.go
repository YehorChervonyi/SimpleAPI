package main

import (
	"fmt"
	"github.com/YehorChervonyi/SimpleAPI/internal/routes"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func main() {
	router := httprouter.New()
	routes.InitRoutes(router)
	fmt.Println("Satarting server on post 8989...")
	http.ListenAndServe(":8989", router)
}
