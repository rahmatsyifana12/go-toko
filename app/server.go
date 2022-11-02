package app

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

type Server struct {
	DB *gorm.DB
	Router *mux.Router
}

// Initialize is used to initiate the server
func (server* Server) Initialize() {
	server.Router = mux.NewRouter()
	server.InitializeRoutes()
}

// Run is used to run the server
func (server* Server) Run(addr string) {
	fmt.Printf("Listening to localhost%s", addr)
	log.Fatal(http.ListenAndServe(addr, server.Router))
}

func Run() {
	server := Server{}

	server.Initialize()
	server.Run(":5000")
}