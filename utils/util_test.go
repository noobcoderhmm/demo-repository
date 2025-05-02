package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test generated using Keploy
func TestIsEmailValid_ValidEmails_123(t *testing.T) {
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
		assert.True(t, result, "Expected true for valid email: %s", email)
	}
}

func TestCalculateSum_PositiveIntegers_789(t *testing.T) {
	// Arrange
	a, b := 5, 10

	// Act
	result := CalculateSum(a, b)

	// Assert
	assert.Equal(t, 15, result, "Expected sum of %d and %d to be %d", a, b, 15)
}

// Test generated using Keploy

