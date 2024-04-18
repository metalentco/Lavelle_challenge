package main

import (
	"fmt"
	"regexp"
)

func isValidCreditCard(number string) bool {
	// Regular expression for validating the credit card format
	creditCardRegex := regexp.MustCompile(`^(4|5|6)\d{3}(-?\d{4}){3}$`)

	// Check the format of the card number
	if !creditCardRegex.MatchString(number) {
		return false
	}

	// Remove hyphens to check for consecutive repeated digits
	cleanNumber := regexp.MustCompile("-").ReplaceAllString(number, "")

	// Check for more than 3 consecutive repeated digits
	for i := 0; i < len(cleanNumber)-3; i++ {
		if cleanNumber[i] == cleanNumber[i+1] && cleanNumber[i+1] == cleanNumber[i+2] && cleanNumber[i+2] == cleanNumber[i+3] {
			return false
		}
	}

	return true
}

func main() {
	// Sample input
	cardNumbers := []string{
		"4123456789123456",
		"5123-4567-8912-3456",
		"61234-567-8912-3456",
		"4123356789123456",
		"5133-3367-8912-3456",
		"5123-3567-8912-3446",
	}

	// Validate each card number
	for _, number := range cardNumbers {
		if isValidCreditCard(number) {
			fmt.Println("Valid")
		} else {
			fmt.Println("Invalid")
		}
	}
}
