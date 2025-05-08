package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test generated using Keploy
func TestIsEmailValid_ValidAndInvalidEmails_001(t *testing.T) {
	// Arrange
	validEmail := "test@example.com"
	invalidEmail := "invalid-email"

	// Act
	isValid := IsEmailValid(validEmail)
	isInvalid := IsEmailValid(invalidEmail)

	// Assert
	assert.True(t, isValid, "Expected valid email to return true")
	assert.False(t, isInvalid, "Expected invalid email to return false")
}

// Test generated using Keploy

func TestCalculateSum_PositiveNegativeZeroValues_002(t *testing.T) {
	// Arrange
	a, b := 5, 10
	c, d := -3, -7
	e, f := 0, 0

	// Act
	sum1 := CalculateSum(a, b)
	sum2 := CalculateSum(c, d)
	sum3 := CalculateSum(e, f)

	// Assert
	assert.Equal(t, 15, sum1, "Expected sum of 5 and 10 to be 15")
	assert.Equal(t, -10, sum2, "Expected sum of -3 and -7 to be -10")
	assert.Equal(t, 0, sum3, "Expected sum of 0 and 0 to be 0")
}

