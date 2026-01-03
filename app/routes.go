package app

import "github.com/argumpamungkas/go-ap-store-app/app/controllers"

// init route
func (server *Server) initializeRoutes() {
	server.Router.HandleFunc("/", controllers.Home).Methods("GET")
}
