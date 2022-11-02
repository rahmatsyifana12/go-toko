package app

import (
	"go-toko/app/controllers"
)

func (server* Server) InitializeRoutes() {
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
}