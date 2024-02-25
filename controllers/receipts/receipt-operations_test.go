package controllers

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/processortest/models"
	"github.com/processortest/utils"
	"github.com/stretchr/testify/assert"
)

func TestProcessReceipts_ValidJSON(t *testing.T) {
	teststr := `{
        "retailer": "Target",
        "purchaseDate": "2022-01-01",
        "purchaseTime": "13:01",
        "items": [
          {
            "shortDescription": "Mountain Dew 12PK",
            "price": "6.49"
          },{
            "shortDescription": "Emils Cheese Pizza",
            "price": "12.25"
          },{
            "shortDescription": "Knorr Creamy Chicken",
            "price": "1.26"
          },{
            "shortDescription": "Doritos Nacho Cheese",
            "price": "3.35"
          },{
            "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
            "price": "12.00"
          }
        ]
      }`
	request, _ := http.NewRequest("POST", "/receipts/process", strings.NewReader(teststr))

	handler := http.HandlerFunc(ProcessReceipts)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, request)

	assert.Equal(t, http.StatusOK, rr.Code)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code : got %v want %v\n", status, http.StatusOK)
	}

	var result map[string]string
	json.Unmarshal(rr.Body.Bytes(), &result)
	assert.NotEmpty(t, result["id"])
}

func TestProcessReceipts_ValidJSON1(t *testing.T) {
	teststr := `{
        "retailer": "Target",
        "purchaseDate": "2022-01-01",
        "purchaseTime": "13:01",
        "items": [
          {
            "shortDescription": "Mountain Dew 12PK",
            "price": "6.49"
          }
        ]
      }`
	request, _ := http.NewRequest("POST", "/receipts/process", strings.NewReader(teststr))

	handler := http.HandlerFunc(ProcessReceipts)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, request)

	assert.Equal(t, http.StatusOK, rr.Code)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code : got %v want %v\n", status, http.StatusOK)
	}

	var result map[string]string
	json.Unmarshal(rr.Body.Bytes(), &result)
	assert.NotEmpty(t, result["id"])
}

func TestGetPoints_ValidReceiptID(t *testing.T) {
	receiptID := utils.GenerateReceiptID()
	receipts[receiptID] = models.Receipt{
		Retailer:     "Example Retailer",
		PurchaseDate: "2022-02-23",
		PurchaseTime: "15:30",
		Items:        []models.Item{{ShortDescription: "Item 1", Price: "10.00"}},
		Total:        "10.00",
	}

	request, _ := http.NewRequest("GET", "/receipts/"+receiptID+"/points", nil)
	response := httptest.NewRecorder()

	GetPoints(response, request)

	assert.Equal(t, http.StatusOK, response.Code)

	var result models.PointsResponse
	json.Unmarshal(response.Body.Bytes(), &result)
	//assert.Equal(t, utils.CalculatePoints(receipts[receiptID]), result.Points)
}

func TestGetPoints_InvalidReceiptID(t *testing.T) {
	invalidReceiptID := "invalidID"
	request, _ := http.NewRequest("GET", "/receipts/"+invalidReceiptID+"/points", nil)
	response := httptest.NewRecorder()

	GetPoints(response, request)

	assert.Equal(t, http.StatusNotFound, response.Code)
}

////////////////////////////////////////////

func TestProcessReceipts_ValidJSON2(t *testing.T) {
	teststr := `{
      "retailer": "Target",
      "purchaseDate": "2022-01-01",
      "purchaseTime": "13:01",
      "items": [
        {
          "shortDescription": "Mountain Dew 12PK",
          "price": "6.49"
        },{
          "shortDescription": "Emils Cheese Pizza",
          "price": "12.25"
        },{
          "shortDescription": "Knorr Creamy Chicken",
          "price": "1.26"
        },{
          "shortDescription": "Doritos Nacho Cheese",
          "price": "3.35"
        },{
          "shortDescription": "   Klarbrunn 12-PK 12 FL OZ  ",
          "price": "12.00"
        }
      ]
    }`
	request, _ := http.NewRequest("POST", "/receipts/process", strings.NewReader(teststr))

	handler := http.HandlerFunc(ProcessReceipts)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, request)

	assert.Equal(t, http.StatusOK, rr.Code)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code : got %v want %v\n", status, http.StatusOK)
	}

	var result map[string]string
	json.Unmarshal(rr.Body.Bytes(), &result)
	assert.NotEmpty(t, result["id"])
}

func TestProcessReceipts_ValidJSON4(t *testing.T) {
	teststr := `{
      "retailer": "Target",
      "purchaseDate": "2022-01-01",
      "purchaseTime": "13:01",
      "items": [
        {
          "shortDescription": "Mountain Dew 12PK",
          "price": "6.49"
        }
      ]
    }`
	request, _ := http.NewRequest("POST", "/receipts/process", strings.NewReader(teststr))

	handler := http.HandlerFunc(ProcessReceipts)
	rr := httptest.NewRecorder()

	handler.ServeHTTP(rr, request)

	assert.Equal(t, http.StatusOK, rr.Code)

	if status := rr.Code; status != http.StatusOK {
		t.Errorf("handler returned wrong status code : got %v want %v\n", status, http.StatusOK)
	}

	var result map[string]string
	json.Unmarshal(rr.Body.Bytes(), &result)
	assert.NotEmpty(t, result["id"])
}
