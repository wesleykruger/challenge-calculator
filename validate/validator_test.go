package validate

import (
	"fmt"
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestSanitizeInput(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    []decimal.Decimal
		expectedErr error
	}{
		{
			name:        "single number",
			input:       "123",
			expected:    []decimal.Decimal{decimal.NewFromInt(123)},
			expectedErr: nil,
		},
		{
			name:        "two numbers",
			input:       "1,5",
			expected:    []decimal.Decimal{decimal.NewFromInt(1), decimal.NewFromInt(5)},
			expectedErr: nil,
		},
		{
			name:        "empty string",
			input:       "",
			expected:    []decimal.Decimal{decimal.Zero},
			expectedErr: nil,
		},
		{
			name:        "mixed positive and negative",
			input:       "4,-3",
			expected:    []decimal.Decimal{decimal.NewFromInt(4), decimal.NewFromInt(-3)},
			expectedErr: nil,
		},
		{
			name:        "too many numbers",
			input:       "1,2,3",
			expected:    nil,
			expectedErr: fmt.Errorf("invalid input, a maximum of %d numbers are allowed, received %d", maxNumbers, 3),
		},
		{
			name:        "negative number",
			input:       "-123,456",
			expected:    []decimal.Decimal{decimal.NewFromInt(-123), decimal.NewFromInt(456)},
			expectedErr: nil,
		},
		{
			name:        "decimal numbers",
			input:       "123.45,67.89",
			expected:    []decimal.Decimal{decimal.NewFromFloat(123.45), decimal.NewFromFloat(67.89)},
			expectedErr: nil,
		},
		{
			name:        "missing first number",
			input:       ",12",
			expected:    []decimal.Decimal{decimal.Zero, decimal.NewFromInt(12)},
			expectedErr: nil,
		},
		{
			name:        "missing second number",
			input:       "12,",
			expected:    []decimal.Decimal{decimal.NewFromInt(12), decimal.Zero},
			expectedErr: nil,
		},
		{
			name:        "missing both numbers",
			input:       ",",
			expected:    []decimal.Decimal{decimal.Zero, decimal.Zero},
			expectedErr: nil,
		},
		{
			name:        "invalid number in second position",
			input:       "5,tytyt",
			expected:    []decimal.Decimal{decimal.NewFromInt(5), decimal.Zero},
			expectedErr: nil,
		},
		{
			name:        "invalid number in first position",
			input:       "tytyt,5",
			expected:    []decimal.Decimal{decimal.Zero, decimal.NewFromInt(5)},
			expectedErr: nil,
		},
		{
			name:        "whitespace handling",
			input:       " 1 , 2 ",
			expected:    []decimal.Decimal{decimal.NewFromInt(1), decimal.NewFromInt(2)},
			expectedErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := SanitizeInput(test.input)
			if test.expectedErr != nil {
				assert.Equal(t, test.expectedErr, err)
			} else {
				assert.Equal(t, test.expected, result)
				assert.NoError(t, err)
			}
		})
	}
}

func TestSplitInput(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "single number",
			input:    "20",
			expected: []string{"20"},
		},
		{
			name:     "two numbers",
			input:    "1,5",
			expected: []string{"1", "5"},
		},
		{
			name:     "empty string",
			input:    "",
			expected: []string{""},
		},
		{
			name:     "missing first number",
			input:    ",12",
			expected: []string{"", "12"},
		},
		{
			name:     "missing second number",
			input:    "12,",
			expected: []string{"12", ""},
		},
		{
			name:     "missing both numbers",
			input:    ",",
			expected: []string{"", ""},
		},
		{
			name:     "whitespace handling",
			input:    " 1 , 2 ",
			expected: []string{"1", "2"},
		},
		{
			name:     "multiple delimiters",
			input:    "1,2,3",
			expected: []string{"1", "2", "3"},
		},
		{
			name:     "decimal numbers",
			input:    "1.5,2.7",
			expected: []string{"1.5", "2.7"},
		},
		{
			name:     "negative numbers",
			input:    "-1,-2",
			expected: []string{"-1", "-2"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := splitInput(test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestValidateNumberOfValues(t *testing.T) {
	tests := []struct {
		name        string
		input       []string
		expectedErr error
	}{
		{
			name:        "single number",
			input:       []string{"20"},
			expectedErr: nil,
		},
		{
			name:        "empty string",
			input:       []string{""},
			expectedErr: nil,
		},
		{
			name:        "two numbers",
			input:       []string{"1", "5"},
			expectedErr: nil,
		},
		{
			name:        "empty slice",
			input:       []string{},
			expectedErr: nil,
		},
		{
			name:        "too many numbers",
			input:       []string{"1", "2", "3"},
			expectedErr: fmt.Errorf("invalid input, a maximum of %d numbers are allowed, received %d", maxNumbers, 3),
		},
		{
			name:        "nil slice",
			input:       nil,
			expectedErr: nil,
		},
		{
			name:        "missing first number",
			input:       []string{"", "12"},
			expectedErr: nil,
		},
		{
			name:        "missing second number",
			input:       []string{"12", ""},
			expectedErr: nil,
		},
		{
			name:        "missing both numbers",
			input:       []string{"", ""},
			expectedErr: nil,
		},
		{
			name:        "invalid number in second position",
			input:       []string{"5", "tytyt"},
			expectedErr: nil,
		},
		{
			name:        "invalid number in first position",
			input:       []string{"tytyt", "5"},
			expectedErr: nil,
		},
		{
			name:        "decimal numbers",
			input:       []string{"1.5", "2.7"},
			expectedErr: nil,
		},
		{
			name:        "negative numbers",
			input:       []string{"-1", "-2"},
			expectedErr: nil,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := validateNumberOfValues(test.input)
			if test.expectedErr != nil {
				assert.Equal(t, test.expectedErr, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestParseDecimal(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected decimal.Decimal
	}{
		{
			name:     "valid number",
			input:    "123",
			expected: decimal.NewFromInt(123),
		},
		{
			name:     "invalid number",
			input:    "tytyt",
			expected: decimal.Zero,
		},
		{
			name:     "empty string",
			input:    "",
			expected: decimal.Zero,
		},
		{
			name:     "negative number",
			input:    "-123",
			expected: decimal.NewFromInt(-123),
		},
		{
			name:     "decimal number",
			input:    "123.45",
			expected: decimal.NewFromFloat(123.45),
		},
		{
			name:     "invalid decimal numbers",
			input:    "123.45.67",
			expected: decimal.Zero,
		},
		{
			name:     "zero",
			input:    "0",
			expected: decimal.NewFromInt(0),
		},
		{
			name:     "negative decimal number",
			input:    "-123.45",
			expected: decimal.NewFromFloat(-123.45),
		},
		{
			name:     "very small number",
			input:    "0.0123",
			expected: decimal.NewFromFloat(0.0123),
		},
		{
			name:     "very large number",
			input:    "12300000000",
			expected: decimal.NewFromInt(12300000000),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := parseDecimal(test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}
