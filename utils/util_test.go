package utils

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Test generated using Keploy
func TestIsEmailValid_RegexBehavior_512(t *testing.T) {
	testCases := []struct {
		name     string
		email    string
		expected bool // Expected result based *strictly* on the regex: ^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,}$
	}{
		{"Valid Simple Email", "test@example.com", true},
		{"Valid Email with Subdomain", "test@sub.example.com", true},
		{"Valid Email with Plus", "test+alias@example.com", true},
		{"Valid Email with Dot", "test.name@example.com", true},
		{"Valid Email with Hyphen", "test-name@example.com", true},
		{"Valid Email with Digits", "test123name@example123.com", true},
		{"Valid Email with Underscore", "test_name@example.com", true},
		{"Valid Email with Percent", "test%name@example.com", true}, // Allowed by regex
		{"Valid TLD length 2", "test@example.co", true},
		{"Valid TLD length 3", "test@example.org", true},
		{"Valid Domain Starts with Dot", "test@.example.com", true},            // Allowed by regex [a-z0-9.\-]
		{"Valid Domain Ends with Hyphen", "test@example-.com", true},           // Allowed by regex [a-z0-9.\-]
		{"Valid Domain Consecutive Dots", "test@example..com", true},           // Allowed by regex [a-z0-9.\-]
		{"Valid Domain Consecutive Hyphens", "test@example--hyphen.com", true}, // Allowed by regex [a-z0-9.\-]

		{"Invalid No At Symbol", "testexample.com", false},
		{"Invalid No Domain", "test@", false},
		{"Invalid No TLD", "test@example", false},
		{"Invalid TLD Too Short", "test@example.c", false}, // Regex requires {2,}
		{"Invalid Characters in Local", "test<script>@example.com", false},
		{"Invalid Characters in Domain", "test@exa<m>ple.com", false},
		{"Invalid Space in Local", "test name@example.com", false},
		{"Invalid Space in Domain", "test@example site.com", false},
		{"Invalid Starts with Hyphen Local", "-test@example.com", true}, // Allowed by regex ^[a-z0-9._%+\-]+
		{"Invalid Starts with Plus Local", "+test@example.com", true},   // Allowed by regex ^[a-z0-9._%+\-]+
		{"Invalid Starts with Dot Local", ".test@example.com", true},    // Allowed by regex ^[a-z0-9._%+\-]+
		{"Empty String", "", false},
		{"Only At Symbol", "@", false},
		{"At Symbol No Local Part", "@example.com", false}, // Regex requires non-empty local part ^[a-z0-9._%+\-]+
		{"At Symbol No Domain Part", "test@", false},       // Regex requires non-empty domain part [a-z0-9.\-]+\.[a-z]{2,}$
		{"Invalid Upper Case", "Test@Example.com", false},  // Regex uses [a-z] only
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := IsEmailValid(tc.email)
			assert.Equal(t, tc.expected, actual, fmt.Sprintf("Email: %s", tc.email))
		})
	}
}

func TestCalculateSum_PositiveNegativeZeroValues_002(t *testing.T) {
	// Arrange
	testCases := []struct {
		a, b     int
		expected int
	}{
		{a: 5, b: 3, expected: 8},
		{a: -5, b: -3, expected: -8},
		{a: 0, b: 0, expected: 0},
		{a: 5, b: -3, expected: 2},
		{a: -5, b: 3, expected: -2},
	}

	// Act & Assert
	for _, tc := range testCases {
		t.Run(fmt.Sprintf("Sum_%d_%d", tc.a, tc.b), func(t *testing.T) {
			result := CalculateSum(tc.a, tc.b)
			assert.Equal(t, tc.expected, result, "Expected sum to match")
		})
	}
}

// Test generated using Keploy

