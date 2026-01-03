package app

import (
	"github.com/argumpamungkas/go-ap-store-app/app/controllers"
	"github.com/gorilla/mux"
)

// init route
func (server *Server) initializeRoutes() {
	server.Router = mux.NewRouter()
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
}
