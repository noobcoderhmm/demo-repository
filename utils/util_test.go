package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestIsEmailValid_ValidEmail_123 tests the IsEmailValid function with a valid email input.
func TestIsEmailValid_ValidEmail_123(t *testing.T) {
	// Arrange
	email := "test@example.com"

	// Act
	result := IsEmailValid(email)

	// Assert
	assert.True(t, result, "Expected the email to be valid")
}

// TestCalculateSum_PositiveIntegers_789 tests the CalculateSum function with positive integers.

func TestCalculateSum_PositiveIntegers_789(t *testing.T) {
	// Arrange
	a, b := 5, 10

	// Act
	result := CalculateSum(a, b)

	// Assert
	assert.Equal(t, 15, result, "Expected the sum of 5 and 10 to be 15")
}

