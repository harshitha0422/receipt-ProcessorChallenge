package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/processortest/routes"
)

func main() {

	router := mux.NewRouter()

	routes.RegisterRoutes(router)

	http.Handle("/", router)
	fmt.Println("Starting server....")

	err := http.ListenAndServe(":9011", router)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}

}
