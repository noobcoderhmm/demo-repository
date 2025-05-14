package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCalculateSum_PositiveIntegers_789 tests the CalculateSum function with positive integers.
func TestIsEmailValid_ValidEmails_123(t *testing.T) {
	validEmails := []string{
		"test@example.com",
		"user.name+tag+sorting@example.com",
		"x@example.com",
		"example-indeed@strange-example.com",
	}

	for _, email := range validEmails {
		result := IsEmailValid(email)
		assert.True(t, result, "Expected email %s to be valid", email)
	}
}

func TestCalculateSum_PositiveIntegers_789(t *testing.T) {
	result := CalculateSum(3, 5)
	assert.Equal(t, 8, result, "Expected sum of 3 and 5 to be 8")
}

// TestIsEmailValid_ValidEmails_123 tests the IsEmailValid function with valid email addresses.

