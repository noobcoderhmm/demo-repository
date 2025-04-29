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
		{a: 0, b: 0, expected: 0},
		{a: 10, b: -5, expected: 5},
		{a: -10, b: 5, expected: -5},
	}

	// Act & Assert
	for _, tc := range testCases {
		result := CalculateSum(tc.a, tc.b)
		assert.Equal(t, tc.expected, result, "Expected %d + %d = %d, got %d", tc.a, tc.b, tc.expected, result)
	}
}

// Test generated using Keploy

func TestIsEmailValid_ValidEmails_789(t *testing.T) {
	// Arrange
	validEmails := []string{
		"test@example.com",
		"user.name+tag+sorting@example.com",
		"x@example.com",
		"example-indeed@strange-example.com",
	}

	// Act & Assert
	for _, email := range validEmails {
		result := IsEmailValid(email)
		assert.True(t, result, "Expected email %s to be valid, but got invalid", email)
	}
}

