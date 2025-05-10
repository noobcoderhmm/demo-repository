package utils

import (
	"fmt"
	"testing"
)

// Test generated using Keploy
func TestIsEmailValid_ValidAndInvalidEmails_123(t *testing.T) {
	tests := []struct {
		email    string
		expected bool
	}{
		{"test@example.com", true},
		{"user.name+tag+sorting@example.com", true},
		{"user@sub.example.com", true},
		{"invalid-email", false},
		{"@missingusername.com", false},
		{"missingatsign.com", false},
		{"missingdomain@.com", false},
	}

	for _, tt := range tests {
		t.Run(tt.email, func(t *testing.T) {
			result := IsEmailValid(tt.email)
			if result != tt.expected {
				t.Errorf("IsEmailValid(%q) = %v; want %v", tt.email, result, tt.expected)
			}
		})
	}
}

// Test generated using Keploy

func TestCalculateSum_PositiveAndNegativeNumbers_456(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{1, 2, 3},
		{-1, -2, -3},
		{0, 0, 0},
		{100, 200, 300},
		{-100, 50, -50},
	}

	for _, tt := range tests {
		t.Run(fmt.Sprintf("%d+%d", tt.a, tt.b), func(t *testing.T) {
			result := CalculateSum(tt.a, tt.b)
			if result != tt.expected {
				t.Errorf("CalculateSum(%d, %d) = %d; want %d", tt.a, tt.b, result, tt.expected)
			}
		})
	}
}

