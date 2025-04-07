package calculate

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name            string
		input           string
		expectedFormula string
		expectedErr     bool
	}{
		{
			name:            "newline delimiter",
			input:           "1\n2",
			expectedFormula: "1+2 = 3",
			expectedErr:     false,
		},
		{
			name:            "mixed delimiters",
			input:           "1\n2,3",
			expectedFormula: "1+2+3 = 6",
			expectedErr:     false,
		},
		{
			name:            "newline with whitespace",
			input:           " 1 \n 2 ",
			expectedFormula: "1+2 = 3",
			expectedErr:     false,
		},
		{
			name:            "multiple newlines",
			input:           "1\n\n2",
			expectedFormula: "1+0+2 = 3",
			expectedErr:     false,
		},
		{
			name:            "newline with empty lines",
			input:           "1\n\n2\n",
			expectedFormula: "1+0+2 = 3",
			expectedErr:     false,
		},
		{
			name:            "simple addition",
			input:           "1,2",
			expectedFormula: "1+2 = 3",
			expectedErr:     false,
		},
		{
			name:            "single number",
			input:           "5",
			expectedFormula: "5 = 5",
			expectedErr:     false,
		},
		{
			name:            "empty input",
			input:           "",
			expectedFormula: "0 = 0",
			expectedErr:     false,
		},
		{
			name:            "many numbers",
			input:           "1,2,3",
			expectedFormula: "1+2+3 = 6",
			expectedErr:     false,
		},
		{
			name:            "invalid number format",
			input:           "abc,def",
			expectedFormula: "0+0 = 0",
			expectedErr:     false,
		},
		{
			name:            "decimal addition",
			input:           "1.5,2.5",
			expectedFormula: "1.5+2.5 = 4",
			expectedErr:     false,
		},
		{
			name:            "negative numbers",
			input:           "-1,-2",
			expectedFormula: "-1+-2 = -3",
			expectedErr:     true,
		},
		{
			name:            "mixed positive and negative",
			input:           "5,-3",
			expectedFormula: "5+-3 = 2",
			expectedErr:     true,
		},
		{
			name:            "negative number with newline deliminator",
			input:           "-1\n-2",
			expectedFormula: "-1+-2 = -3",
			expectedErr:     true,
		},
		{
			name:            "large numbers",
			input:           "1000000,2000000",
			expectedFormula: "0+0 = 0",
			expectedErr:     false,
		},
		{
			name:            "small decimal numbers",
			input:           "0.1,0.2",
			expectedFormula: "0.1+0.2 = 0.3",
			expectedErr:     false,
		},
		{
			name:            "missing first number",
			input:           ",5",
			expectedFormula: "0+5 = 5",
			expectedErr:     false,
		},
		{
			name:            "missing second number",
			input:           "5,",
			expectedFormula: "5+0 = 5",
			expectedErr:     false,
		},
		{
			name:            "missing both numbers",
			input:           ",",
			expectedFormula: "0+0 = 0",
			expectedErr:     false,
		},
		{
			name:            "many numbers",
			input:           "1,2,3,4,5,6,7,8,9,10",
			expectedFormula: "1+2+3+4+5+6+7+8+9+10 = 55",
			expectedErr:     false,
		},
		{
			name:            "many numbers with some invalid",
			input:           "1,2,3,4,5,6,7,8,9,10,abc",
			expectedFormula: "1+2+3+4+5+6+7+8+9+10+0 = 55",
			expectedErr:     false,
		},
		{
			name:            "mixed large numbers",
			input:           "1000000,2000000,3\n4",
			expectedFormula: "0+0+3+4 = 7",
			expectedErr:     false,
		},
		{
			name:            "custom multiple delimiters",
			input:           "//[*][!!][r9r]\n11r9r22*hh*33!!44",
			expectedFormula: "11+22+0+33+44 = 110",
			expectedErr:     false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := Add(test.input)
			if test.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expectedFormula, result)
			}
		})
	}
}

func TestNumberExceedsMaxValue(t *testing.T) {
	tests := []struct {
		name     string
		input    decimal.Decimal
		expected bool
	}{
		{
			name:     "number below max",
			input:    decimal.NewFromInt(999),
			expected: false,
		},
		{
			name:     "number at max",
			input:    decimal.NewFromInt(1000),
			expected: false,
		},
		{
			name:     "number above max",
			input:    decimal.NewFromInt(1001),
			expected: true,
		},
		{
			name:     "zero",
			input:    decimal.Zero,
			expected: false,
		},
		{
			name:     "negative number",
			input:    decimal.NewFromInt(-1000),
			expected: false,
		},
		{
			name:     "negative number below abs max",
			input:    decimal.NewFromInt(-1001),
			expected: false,
		},
		{
			name:     "decimal below max",
			input:    decimal.NewFromFloat(999.99),
			expected: false,
		},
		{
			name:     "decimal at max",
			input:    decimal.NewFromFloat(1000.00),
			expected: false,
		},
		{
			name:     "decimal above max",
			input:    decimal.NewFromFloat(1000.01),
			expected: true,
		},
		{
			name:     "very large number",
			input:    decimal.NewFromInt(1000000),
			expected: true,
		},
		{
			name:     "very small number",
			input:    decimal.NewFromFloat(0.0001),
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := numberExceedsMaxValue(test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}
