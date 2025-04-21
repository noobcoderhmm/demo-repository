package utils

import (
	"fmt"
	"testing"
)

// Test generated using Keploy
func TestIsEmailValid_ValidEmails_123(t *testing.T) {
	validEmails := []string{
		"test@example.com",
		"user.name+tag+sorting@example.com",
		"x@x.au",
		"example-indeed@strange-example.com",
	}

	for _, email := range validEmails {
		t.Run(email, func(t *testing.T) {
			result := IsEmailValid(email)
			if !result {
				t.Errorf("Expected email %s to be valid, but got invalid", email)
			}
		})
	}
}

// Test generated using Keploy

func TestCalculateSum_BasicCases_789(t *testing.T) {
	testCases := []struct {
		a, b     int
		expected int
	}{
		{1, 2, 3},
		{0, 0, 0},
		{-1, -1, -2},
		{100, 200, 300},
		{-50, 50, 0},
	}

	for _, tc := range testCases {
		t.Run(fmt.Sprintf("%d+%d", tc.a, tc.b), func(t *testing.T) {
			result := CalculateSum(tc.a, tc.b)
			if result != tc.expected {
				t.Errorf("Expected sum of %d and %d to be %d, but got %d", tc.a, tc.b, tc.expected, result)
			}
		})
	}
}

