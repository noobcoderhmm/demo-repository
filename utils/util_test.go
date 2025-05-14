package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCalculateSum_PositiveNegativeZeroValues_456 tests the CalculateSum function with positive, negative, and zero values.
func TestIsEmailValid_ValidEmail_123(t *testing.T) {
	// Arrange
	validEmail := "test@example.com"

	// Act
	isValid := IsEmailValid(validEmail)

	// Assert
	assert.True(t, isValid, "Expected email to be valid")
}

func TestCalculateSum_PositiveNegativeZeroValues_456(t *testing.T) {
	// Arrange
	testCases := []struct {
		a, b     int
		expected int
	}{
		{a: 5, b: 3, expected: 8},
		{a: -5, b: -3, expected: -8},
		{a: 0, b: 0, expected: 0},
		{a: -5, b: 5, expected: 0},
		{a: 10, b: -3, expected: 7},
	}

	// Act & Assert
	for _, tc := range testCases {
		result := CalculateSum(tc.a, tc.b)
		assert.Equal(t, tc.expected, result, "Expected %d + %d = %d, but got %d", tc.a, tc.b, tc.expected, result)
	}
}

// TestIsEmailValid_ValidEmail_123 tests IsEmailValid with a valid email.

