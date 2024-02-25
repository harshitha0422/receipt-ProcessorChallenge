package main

import (
	"github.com/processortest/routes"
	"fmt"
	"net/http"
	"github.com/gorilla/mux"
)

func main() {

	router := mux.NewRouter()
     
	
    fmt.Println("Starting server1....")
	routes.RegisterRoutes(router)

	http.Handle("/", router)
	fmt.Println("Starting server....")
	
	err := http.ListenAndServe(":9011", router)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
	

}
