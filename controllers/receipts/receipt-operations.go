package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"sync"

	"github.com/gorilla/mux"
	"github.com/processortest/models"
	"github.com/processortest/utils"
)

var (
	receipts      = make(map[string]models.Receipt)
	receiptsMutex sync.Mutex
)

// ProcessReceipts handles the processing of receipts.
func ProcessReceipts(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Inside process receipts ....")
	var receipt models.Receipt
	err := json.NewDecoder(r.Body).Decode(&receipt)
	if err != nil {
		http.Error(w, "Invalid JSON payload", http.StatusBadRequest)
		return
	}

	// Lock the mutex to ensure safe access to the receipts map
	receiptsMutex.Lock()
	defer receiptsMutex.Unlock()

	receiptID := utils.GenerateReceiptID()

	receipts[receiptID] = receipt

	json.NewEncoder(w).Encode(map[string]string{"id": receiptID})
}

// GetPoints handles the retrieval of points for a given receipt ID.
func GetPoints(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	receiptID := params["id"]
	fmt.Println(receiptID)

	// Lock the mutex to ensure safe access to the receipts map
	receiptsMutex.Lock()
	defer receiptsMutex.Unlock()

	receipt, exists := receipts[receiptID]
	if !exists {
		http.Error(w, "Receipt not found", http.StatusNotFound)
		return
	}

	points, err := utils.CalculatePoints(receipt)
	if err != nil {
		http.Error(w, "Error calculating points", http.StatusInternalServerError)
		return
	}

	json.NewEncoder(w).Encode(models.PointsResponse{Points: points})
}
