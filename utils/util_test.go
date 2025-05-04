package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test generated using Keploy
func TestIsEmailValid_ValidEmail_123(t *testing.T) {
	// Arrange
	validEmail := "test@example.com"

	// Act
	result := IsEmailValid(validEmail)

	// Assert
	assert.True(t, result, "Expected valid email to return true")
}

// Test generated using Keploy

func TestCalculateSum_PositiveIntegers_101(t *testing.T) {
	// Arrange
	a, b := 5, 10

	// Act
	result := CalculateSum(a, b)

	// Assert
	assert.Equal(t, 15, result, "Expected sum of 5 and 10 to be 15")
}

