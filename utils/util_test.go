package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test generated using Keploy
func TestCalculateSum_PositiveNegativeZero_456(t *testing.T) {
	// Arrange
	testCases := []struct {
		a, b     int
		expected int
	}{
		{a: 5, b: 3, expected: 8},
		{a: -5, b: -3, expected: -8},
		{a: 5, b: -3, expected: 2},
		{a: 0, b: 0, expected: 0},
	}

	// Act & Assert
	for _, tc := range testCases {
		result := CalculateSum(tc.a, tc.b)
		assert.Equal(t, tc.expected, result, "Expected %d + %d = %d, but got %d", tc.a, tc.b, tc.expected, result)
	}
}

// Test generated using Keploy

func TestIsEmailValid_ValidInvalidEmails_123(t *testing.T) {
	// Arrange
	testCases := []struct {
		email    string
		expected bool
	}{
		{email: "test@example.com", expected: true},
		{email: "user.name+tag+sorting@example.com", expected: true},
		{email: "user@sub.domain.com", expected: true},
		{email: "invalid-email", expected: false},
		{email: "missing@domain", expected: false},
		{email: "@missingusername.com", expected: false},
		{email: "missingatsign.com", expected: false},
		{email: "user@.com", expected: false},
	}

	// Act & Assert
	for _, tc := range testCases {
		result := IsEmailValid(tc.email)
		assert.Equal(t, tc.expected, result, "Expected IsEmailValid(%s) to return %v, but got %v", tc.email, tc.expected, result)
	}
}

// func TestCalculateSum_PositiveNegativeZero_457(t *testing.T) {
// 	// Arrange
// 	testCases := []struct {
// 		a, b     int
// 		expected int
// 	}{
// 		{a: 5, b: 3, expected: 8},
// 		{a: -5, b: -3, expected: -8},
// 		{a: 5, b: -3, expected: 2},
// 		{a: 0, b: 0, expected: 0},
// 	}
// 	testCases = nil
// 	testCases[0].a = 5
// 	// Act & Assert
// 	for _, tc := range testCases {
// 		result := CalculateSum(tc.a, tc.b)
// 		assert.Equal(t, tc.expected, result, "Expected %d + %d = %d, but got %d", tc.a, tc.b, tc.expected, result)
// 	}
// }

