package main

import (
	"backend/dbconfig"
	"backend/logging"
	"backend/routehandler"
	"fmt"
	"net/http"

	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
)

// @title POC Dashboard
// @version 1.0
// @description POC Dashboard
// @host localhost:8080
func main() {
	logging.Init()
	router := mux.NewRouter()

	db, err := dbconfig.Connect()
	if err != nil {
		logging.Logger.WithError(err).Fatal("Failed to connect to the database")
	}
	routehandler.HandlePocRoutes(router, db)

	corsMiddleware := handlers.CORS(
		handlers.AllowedOrigins([]string{"*", "http://localhost:8080", "http://localhost:3000"}),
		handlers.AllowedMethods([]string{"GET", "POST", "PUT", "DELETE", "OPTIONS"}),
		handlers.AllowedHeaders([]string{"Content-Type"}),
	)

	routerWithCORS := corsMiddleware(router)

	addr := fmt.Sprintf(":%s", dbconfig.Port)
	logging.Logger.Infof("Server is listening on %s", addr)

	logging.Logger.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", dbconfig.Port), routerWithCORS))
}
