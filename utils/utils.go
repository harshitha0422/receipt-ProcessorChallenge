package utils

import (
	
	
	"fmt"
    "sync"
    "time"
	"math"
	"strings"
	"strconv" // Add import for strconv package
	"github.com/processortest/models"
)

var receiptsMutex sync.Mutex

// CalculatePoints calculates the points for a given receipt based on the specified rules.
func CalculatePoints(receipt models.Receipt) (int,error) {
    points := 0

	// Rule: One point for every alphanumeric character in the retailer name.
	points += len(strings.ReplaceAll(receipt.Retailer, " ", ""))

	// Rule: 50 points if the total is a round dollar amount with no cents.
	if isRoundDollarAmount(receipt.Total) {
		points += 50
	}

	// Rule: 25 points if the total is a multiple of 0.25.
	if isMultipleOfQuarter(receipt.Total) {
		points += 25
	}

	// Rule: 5 points for every two items on the receipt.
	points += (len(receipt.Items) / 2) * 5

	// Rule: If the trimmed length of the item description is a multiple of 3,
	// multiply the price by 0.2 and round up to the nearest integer.
	for _, item := range receipt.Items {
		trimmedLength := len(strings.TrimSpace(item.ShortDescription))
		if trimmedLength%3 == 0 {
			price, err := strconv.ParseFloat(item.Price, 64)
            if err != nil {
                // Handle parsing error
				return 0, fmt.Errorf("failed to parse price: %w", err)
                // continue
            }
			points += int(math.Ceil(price * 0.2))
		}
	}

	// Rule: 6 points if the day in the purchase date is odd.
	purchaseDate, err := time.Parse("2006-01-02", receipt.PurchaseDate)
	if err == nil && purchaseDate.Day()%2 == 1 {
		points += 6
	}

	// Rule: 10 points if the time of purchase is after 2:00pm and before 4:00pm.
	purchaseTime, err := time.Parse("15:04", receipt.PurchaseTime)
	if err == nil && purchaseTime.After(time.Date(0, 1, 1, 14, 0, 0, 0, time.UTC)) &&
		purchaseTime.Before(time.Date(0, 1, 1, 16, 0, 0, 0, time.UTC)) {
		points += 10
	}

	return points, nil
}

// Helper function to check if the total is a round dollar amount with no cents.
func isRoundDollarAmount(total string) bool {
	// Assuming total is in the format "xx.xx"
	return strings.HasSuffix(total, ".00")
}

// Helper function to check if the total is a multiple of 0.25.
func isMultipleOfQuarter(total string) bool {
	// Assuming total is in the format "xx.xx"
	value := parseTotal(total)
	return math.Mod(value, 0.25) == 0
}

// Helper function to parse the total value as a float64.
func parseTotal(total string) float64 {
	value, err := strconv.ParseFloat(total, 64)
	if err != nil {
		return 0.0
	}
	return value
}

// GenerateReceiptID generates a unique ID for a receipt.
func GenerateReceiptID() string {
    return fmt.Sprintf("%x", time.Now().UnixNano())
}
