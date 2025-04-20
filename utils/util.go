package utils

import (
	"regexp"
)

// IsEmailValid checks if the email is valid using a basic regex
func IsEmailValid(email string) bool {
	
	// Very basic regex for illustration
	re := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$`)
	return re.MatchString(email)
}

// CalculateSum returns the sum of two integers
func CalculateSum(a, b int) int {
	return a + b
}
