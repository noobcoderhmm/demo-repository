package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestIsEmailValid_ValidAndInvalidEmails_123 tests the IsEmailValid function with valid and invalid email inputs.
func TestIsEmailValid_ValidAndInvalidEmails_123(t *testing.T) {
	// Arrange
	validEmail := "test@example.com"
	invalidEmail := "invalid-email"

	// Act
	validResult := IsEmailValid(validEmail)
	invalidResult := IsEmailValid(invalidEmail)

	// Assert
	assert.True(t, validResult, "Expected valid email to return true")
	assert.False(t, invalidResult, "Expected invalid email to return false")
}

// TestCalculateSum_PositiveNegativeZero_456 tests the CalculateSum function with positive, negative, and zero values.

func TestCalculateSum_PositiveNegativeZero_456(t *testing.T) {
	// Arrange
	a, b := 5, 3
	c, d := -5, -3
	e, f := 0, 0

	// Act
	positiveSum := CalculateSum(a, b)
	negativeSum := CalculateSum(c, d)
	zeroSum := CalculateSum(e, f)

	// Assert
	assert.Equal(t, 8, positiveSum, "Expected sum of 5 and 3 to be 8")
	assert.Equal(t, -8, negativeSum, "Expected sum of -5 and -3 to be -8")
	assert.Equal(t, 0, zeroSum, "Expected sum of 0 and 0 to be 0")
}

