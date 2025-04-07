package validate

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestValidateInput(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    []decimal.Decimal
		expectedErr string
	}{
		{
			name:        "valid input with two numbers",
			input:       "1,2",
			expected:    []decimal.Decimal{decimal.NewFromInt(1), decimal.NewFromInt(2)},
			expectedErr: "",
		},
		{
			name:        "newline delimiter",
			input:       "1\n2",
			expected:    []decimal.Decimal{decimal.NewFromInt(1), decimal.NewFromInt(2)},
			expectedErr: "",
		},
		{
			name:        "mixed delimiters",
			input:       "1\n2,3",
			expected:    []decimal.Decimal{decimal.NewFromInt(1), decimal.NewFromInt(2), decimal.NewFromInt(3)},
			expectedErr: "",
		},
		{
			name:        "newline with whitespace",
			input:       " 1 \n 2 ",
			expected:    []decimal.Decimal{decimal.NewFromInt(1), decimal.NewFromInt(2)},
			expectedErr: "",
		},
		{
			name:        "multiple newlines",
			input:       "1\n\n2",
			expected:    []decimal.Decimal{decimal.NewFromInt(1), decimal.NewFromInt(0), decimal.NewFromInt(2)},
			expectedErr: "",
		},
		{
			name:        "newline with empty lines",
			input:       "1\n\n2\n",
			expected:    []decimal.Decimal{decimal.NewFromInt(1), decimal.NewFromInt(0), decimal.NewFromInt(2)},
			expectedErr: "",
		},
		{
			name:        "single number",
			input:       "123",
			expected:    []decimal.Decimal{decimal.NewFromInt(123)},
			expectedErr: "",
		},
		{
			name:        "two numbers",
			input:       "1,5",
			expected:    []decimal.Decimal{decimal.NewFromInt(1), decimal.NewFromInt(5)},
			expectedErr: "",
		},
		{
			name:        "empty string",
			input:       "",
			expected:    []decimal.Decimal{decimal.Zero},
			expectedErr: "",
		},
		{
			name:        "mixed positive and negative",
			input:       "4,-3",
			expected:    nil,
			expectedErr: "invalid input: negative numbers found: -3",
		},
		{
			name:        "many numbers",
			input:       "1,2,3",
			expected:    []decimal.Decimal{decimal.NewFromInt(1), decimal.NewFromInt(2), decimal.NewFromInt(3)},
			expectedErr: "",
		},
		{
			name:        "negative number",
			input:       "-123,456",
			expected:    nil,
			expectedErr: "invalid input: negative numbers found: -123",
		},
		{
			name:        "decimal numbers",
			input:       "123.45,67.89",
			expected:    []decimal.Decimal{decimal.NewFromFloat(123.45), decimal.NewFromFloat(67.89)},
			expectedErr: "",
		},
		{
			name:        "missing first number",
			input:       ",12",
			expected:    []decimal.Decimal{decimal.NewFromInt(0), decimal.NewFromInt(12)},
			expectedErr: "",
		},
		{
			name:        "missing second number",
			input:       "12,",
			expected:    []decimal.Decimal{decimal.NewFromInt(12), decimal.NewFromInt(0)},
			expectedErr: "",
		},
		{
			name:        "missing both numbers",
			input:       ",",
			expected:    []decimal.Decimal{decimal.NewFromInt(0), decimal.NewFromInt(0)},
			expectedErr: "",
		},
		{
			name:        "invalid number in second position",
			input:       "5,tytyt",
			expected:    []decimal.Decimal{decimal.NewFromInt(5), decimal.Zero},
			expectedErr: "",
		},
		{
			name:        "invalid number in first position",
			input:       "tytyt,5",
			expected:    []decimal.Decimal{decimal.Zero, decimal.NewFromInt(5)},
			expectedErr: "",
		},
		{
			name:        "whitespace handling",
			input:       " 1 , 2 ",
			expected:    []decimal.Decimal{decimal.NewFromInt(1), decimal.NewFromInt(2)},
			expectedErr: "",
		},
		{
			name:        "custom multiple delimiters",
			input:       "//[*][!!][r9r]\n11r9r22*hh*33!!44",
			expected:    []decimal.Decimal{decimal.NewFromInt(11), decimal.NewFromInt(22), decimal.Zero, decimal.NewFromInt(33), decimal.NewFromInt(44)},
			expectedErr: "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := ValidateInput(test.input)
			if test.expectedErr != "" {
				assert.EqualError(t, err, test.expectedErr)
			} else {
				assert.Equal(t, test.expected, result)
				assert.NoError(t, err)
			}
		})
	}
}

func TestSanitizeInput(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    []decimal.Decimal
		expectedErr error
	}{
		{
			name:        "valid input with two numbers",
			input:       "1,2",
			expected:    []decimal.Decimal{decimal.NewFromInt(1), decimal.NewFromInt(2)},
			expectedErr: nil,
		},
		{
			name:        "newline delimiter",
			input:       "1\n2",
			expected:    []decimal.Decimal{decimal.NewFromInt(1), decimal.NewFromInt(2)},
			expectedErr: nil,
		},
		{
			name:        "mixed delimiters",
			input:       "1\n2,3",
			expected:    []decimal.Decimal{decimal.NewFromInt(1), decimal.NewFromInt(2), decimal.NewFromInt(3)},
			expectedErr: nil,
		},
		{
			name:        "newline with whitespace",
			input:       " 1 \n 2 ",
			expected:    []decimal.Decimal{decimal.NewFromInt(1), decimal.NewFromInt(2)},
			expectedErr: nil,
		},
		{
			name:        "multiple newlines",
			input:       "1\n\n2",
			expected:    []decimal.Decimal{decimal.NewFromInt(1), decimal.NewFromInt(0), decimal.NewFromInt(2)},
			expectedErr: nil,
		},
		{
			name:        "newline with empty lines",
			input:       "1\n\n2\n",
			expected:    []decimal.Decimal{decimal.NewFromInt(1), decimal.NewFromInt(0), decimal.NewFromInt(2)},
			expectedErr: nil,
		},
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
			name:        "many numbers",
			input:       "1,2,3",
			expected:    []decimal.Decimal{decimal.NewFromInt(1), decimal.NewFromInt(2), decimal.NewFromInt(3)},
			expectedErr: nil,
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
			expected:    []decimal.Decimal{decimal.NewFromInt(0), decimal.NewFromInt(12)},
			expectedErr: nil,
		},
		{
			name:        "missing second number",
			input:       "12,",
			expected:    []decimal.Decimal{decimal.NewFromInt(12), decimal.NewFromInt(0)},
			expectedErr: nil,
		},
		{
			name:        "missing both numbers",
			input:       ",",
			expected:    []decimal.Decimal{decimal.NewFromInt(0), decimal.NewFromInt(0)},
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
			result, err := sanitizeInput(test.input)
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
		delims   []string
		expected []string
	}{
		{
			name:     "single number",
			input:    "20",
			delims:   nil,
			expected: []string{"20"},
		},
		{
			name:     "two numbers with default delimiter",
			input:    "1,5",
			delims:   nil,
			expected: []string{"1", "5"},
		},
		{
			name:     "empty string",
			input:    "",
			delims:   nil,
			expected: []string{"0"},
		},
		{
			name:     "missing first number",
			input:    ",12",
			delims:   nil,
			expected: []string{"0", "12"},
		},
		{
			name:     "missing second number",
			input:    "12,",
			delims:   nil,
			expected: []string{"12", "0"},
		},
		{
			name:     "missing both numbers",
			input:    ",",
			delims:   nil,
			expected: []string{"0", "0"},
		},
		{
			name:     "whitespace handling",
			input:    " 1 , 2 ",
			delims:   nil,
			expected: []string{"1", "2"},
		},
		{
			name:     "multiple delimiters",
			input:    "1,2,3",
			delims:   nil,
			expected: []string{"1", "2", "3"},
		},
		{
			name:     "decimal numbers",
			input:    "1.5,2.7",
			delims:   nil,
			expected: []string{"1.5", "2.7"},
		},
		{
			name:     "negative numbers",
			input:    "-1,-2",
			delims:   nil,
			expected: []string{"-1", "-2"},
		},
		{
			name:     "custom delimiter",
			input:    "1*2*3",
			delims:   []string{"*"},
			expected: []string{"1", "2", "3"},
		},
		{
			name:     "multiple custom delimiters",
			input:    "11r9r22*hh*33!!44",
			delims:   []string{"*", "!!", "r9r"},
			expected: []string{"11", "22", "hh", "33", "44"},
		},
		{
			name:     "mixed default and custom delimiters",
			input:    "1,2*3\n4",
			delims:   []string{"*"},
			expected: []string{"1", "2", "3", "4"},
		},
		{
			name:     "consecutive delimiters",
			input:    "1,,2",
			delims:   nil,
			expected: []string{"1", "0", "2"},
		},
		{
			name:     "delimiter at start and end",
			input:    ",1,2,",
			delims:   nil,
			expected: []string{"0", "1", "2", "0"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Set custom delimiters for this test
			customDelimiters = test.delims
			result := splitInput(test.input)
			assert.Equal(t, test.expected, result)
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
		{
			name:     "leading zeros",
			input:    "00123",
			expected: decimal.NewFromInt(123),
		},
		{
			name:  "trailing zeros",
			input: "123.4500",
			expected: func() decimal.Decimal {
				d, _ := decimal.NewFromString("123.4500")
				return d
			}(),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := parseDecimal(test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestUnescapeNewline(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "single escaped newline",
			input:    "1\\n2",
			expected: "1\n2",
		},
		{
			name:     "multiple escaped newlines",
			input:    "1\\n2\\n3",
			expected: "1\n2\n3",
		},
		{
			name:     "no escaped newlines",
			input:    "1,2,3",
			expected: "1,2,3",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "mixed delimiters",
			input:    "1\\n2,3",
			expected: "1\n2,3",
		},
		{
			name:     "escaped newline at start",
			input:    "\\n1,2",
			expected: "\n1,2",
		},
		{
			name:     "escaped newline at end",
			input:    "1,2\\n",
			expected: "1,2\n",
		},
		{
			name:     "consecutive escaped newlines",
			input:    "1\\n\\n2",
			expected: "1\n\n2",
		},
		{
			name:     "escaped newline with whitespace",
			input:    " 1 \\n 2 ",
			expected: " 1 \n 2 ",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := UnescapeNewline(test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestFindNegativeNumbers(t *testing.T) {
	tests := []struct {
		name     string
		input    []decimal.Decimal
		expected []string
	}{
		{
			name:     "no negative numbers",
			input:    []decimal.Decimal{decimal.NewFromInt(1), decimal.NewFromInt(2), decimal.NewFromInt(3)},
			expected: nil,
		},
		{
			name:     "single negative number",
			input:    []decimal.Decimal{decimal.NewFromInt(1), decimal.NewFromInt(-2), decimal.NewFromInt(3)},
			expected: []string{"-2"},
		},
		{
			name:     "multiple negative numbers",
			input:    []decimal.Decimal{decimal.NewFromInt(-1), decimal.NewFromInt(-2), decimal.NewFromInt(-3)},
			expected: []string{"-1", "-2", "-3"},
		},
		{
			name:     "negative decimal numbers",
			input:    []decimal.Decimal{decimal.NewFromFloat(-1.5), decimal.NewFromFloat(-2.5)},
			expected: []string{"-1.5", "-2.5"},
		},
		{
			name:     "mixed positive and negative",
			input:    []decimal.Decimal{decimal.NewFromInt(1), decimal.NewFromInt(-2), decimal.NewFromInt(3), decimal.NewFromInt(-4)},
			expected: []string{"-2", "-4"},
		},
		{
			name:     "empty slice",
			input:    []decimal.Decimal{},
			expected: nil,
		},
		{
			name:     "zero values",
			input:    []decimal.Decimal{decimal.Zero, decimal.Zero, decimal.Zero},
			expected: nil,
		},
		{
			name:     "very small negative numbers",
			input:    []decimal.Decimal{decimal.NewFromFloat(-0.0001), decimal.NewFromFloat(-0.0002)},
			expected: []string{"-0.0001", "-0.0002"},
		},
		{
			name:     "very large negative numbers",
			input:    []decimal.Decimal{decimal.NewFromInt(-1000000), decimal.NewFromInt(-2000000)},
			expected: []string{"-1000000", "-2000000"},
		},
		{
			name:     "mixed decimal and integer negatives",
			input:    []decimal.Decimal{decimal.NewFromInt(-1), decimal.NewFromFloat(-2.5), decimal.NewFromInt(-3)},
			expected: []string{"-1", "-2.5", "-3"},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := findNegativeNumbers(test.input)
			assert.Equal(t, test.expected, result)
		})
	}
}

func TestProcessCustomDelimiters(t *testing.T) {
	tests := []struct {
		name           string
		input          string
		expectedInput  string
		expectedDelims []string
		expectedErr    string
	}{
		{
			name:           "no custom delimiter",
			input:          "1,2",
			expectedInput:  "1,2",
			expectedDelims: nil,
			expectedErr:    "",
		},
		{
			name:           "single character delimiter",
			input:          "//*\n1*2",
			expectedInput:  "1*2",
			expectedDelims: []string{"*"},
			expectedErr:    "",
		},
		{
			name:           "multiple bracketed delimiters",
			input:          "//[*][!!][r9r]\n11r9r22*hh*33!!44",
			expectedInput:  "11r9r22*hh*33!!44",
			expectedDelims: []string{"*", "!!", "r9r"},
			expectedErr:    "",
		},
		{
			name:           "single bracketed delimiter",
			input:          "//[***]\n1***2",
			expectedInput:  "1***2",
			expectedDelims: []string{"***"},
			expectedErr:    "",
		},
		{
			name:           "empty brackets",
			input:          "//[]\n1,2",
			expectedInput:  "1,2",
			expectedDelims: nil,
			expectedErr:    "",
		},
		{
			name:           "missing closing bracket",
			input:          "//[***\n1***2",
			expectedInput:  "1***2",
			expectedDelims: nil,
			expectedErr:    "invalid delimiter format: missing closing bracket",
		},
		{
			name:           "invalid simple delimiter",
			input:          "//**\n1*2",
			expectedInput:  "1*2",
			expectedDelims: nil,
			expectedErr:    "invalid custom delimiter: \"**\"",
		},
		{
			name:           "no newline after delimiter",
			input:          "//*1*2",
			expectedInput:  "//*1*2",
			expectedDelims: nil,
			expectedErr:    "",
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			// Reset customDelimiters before each test
			customDelimiters = nil

			result, err := processCustomDelimiters(test.input)
			if test.expectedErr != "" {
				assert.EqualError(t, err, test.expectedErr)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expectedInput, result)
				assert.Equal(t, test.expectedDelims, customDelimiters)
			}
		})
	}
}
