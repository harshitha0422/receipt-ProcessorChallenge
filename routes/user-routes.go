package routes

import (
	"github.com/gorilla/mux"
	controllers "github.com/processortest/controllers/receipts"
)

var RegisterRoutes = func(router *mux.Router) {
	router.HandleFunc("/receipts/process", controllers.ProcessReceipts).Methods("POST")
	router.HandleFunc("/receipts/{id}/points", controllers.GetPoints).Methods("GET")
}
