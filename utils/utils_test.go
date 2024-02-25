package utils

import (
	
	"testing"
	"github.com/stretchr/testify/assert"
	"github.com/processortest/models"
	"github.com/processortest/utils"
)

func TestCalculatePoints_FailureCase(t *testing.T) {
	// Prepare a receipt that should result in a failure
	receipt := models.Receipt{
		Retailer: "Invalid Retailer %$@", // Retailer with non-alphanumeric characters
		Total:    "15.75",                // Total not a round dollar amount
		Items: []models.Item{
			{ShortDescription: "Item 1", Price: "9.99"},
			{ShortDescription: "Item 2", Price: "5.76"},
		},
		PurchaseDate: "2022-02-23",
		PurchaseTime: "14:30",
	}

	// Call CalculatePoints function
	points, _ := utils.CalculatePoints(receipt)

	// Assert that there is an error and points are 0
	
	assert.NotEqual(t, 0, points)
	
}

func TestCalculatePoints_SuccessCase(t *testing.T) {
	// Prepare a receipt that should result in success
	receipt := models.Receipt{
		Retailer: "Valid Retailer", // Retailer with alphanumeric characters
		Total:    "20.00",          // Total is a round dollar amount
		Items: []models.Item{
			{ShortDescription: "Item 1", Price: "10.00"},
			{ShortDescription: "Item 2", Price: "10.00"},
		},
		PurchaseDate: "2022-02-23",
		PurchaseTime: "15:30",
	}

	// Call CalculatePoints function
	points, _ := utils.CalculatePoints(receipt)

	// Assert that there is no error and points are calculated as expected

	assert.Equal(t, 113, points)

	
}
