package utils

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

// TestCalculateSum_PositiveNegativeZero_002 tests the CalculateSum function with positive, negative, and zero values.
func TestCalculateSum_PositiveNegativeZero_002(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{3, 5, 8},
		{-3, -5, -8},
		{0, 0, 0},
		{10, -5, 5},
		{-10, 5, -5},
	}

	for _, tt := range tests {
		result := CalculateSum(tt.a, tt.b)
		assert.Equal(t, tt.expected, result, "Expected %d + %d = %d, got %d", tt.a, tt.b, tt.expected, result)
	}
}

// TestCalculateSub_PositiveNegativeZero_003 tests the CalculateSub function with positive, negative, and zero values.

func TestCalculateSub_PositiveNegativeZero_003(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{5, 3, 2},
		{-5, -3, -2},
		{0, 0, 0},
		{10, -5, 15},
		{-10, 5, -15},
	}

	for _, tt := range tests {
		result := CalculateSub(tt.a, tt.b)
		assert.Equal(t, tt.expected, result, "Expected %d - %d = %d, got %d", tt.a, tt.b, tt.expected, result)
	}
}

// TestCalculatemul_PositiveNegativeZero_004 tests the Calculatemul function with positive, negative, and zero values.

func TestCalculatemul_PositiveNegativeZero_004(t *testing.T) {
	tests := []struct {
		a, b     int
		expected int
	}{
		{3, 5, 15},
		{-3, -5, 15},
		{0, 5, 0},
		{10, -5, -50},
		{-10, 5, -50},
	}

	for _, tt := range tests {
		result := Calculatemul(tt.a, tt.b)
		assert.Equal(t, tt.expected, result, "Expected %d * %d = %d, got %d", tt.a, tt.b, tt.expected, result)
	}
}

