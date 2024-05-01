package main

import (
	"fmt"
	"github.com/YehorChervonyi/SimpleAPI/config"
	"github.com/YehorChervonyi/SimpleAPI/internal/routes"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
)

func main() {
	cfg := config.LoadConfig()
	port := strconv.Itoa(cfg.Server.Port)
	router := httprouter.New()
	routes.InitRoutes(router)
	fmt.Printf("Satarting server on post %s...", port)
	http.ListenAndServe(":"+port, router)
}
