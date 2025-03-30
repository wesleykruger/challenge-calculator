package calculate

import (
	"testing"

	"github.com/shopspring/decimal"
	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	tests := []struct {
		name        string
		input       string
		expected    decimal.Decimal
		expectedErr bool
	}{
		{
			name:        "simple addition",
			input:       "1,2",
			expected:    decimal.NewFromInt(3),
			expectedErr: false,
		},
		{
			name:        "single number",
			input:       "5",
			expected:    decimal.NewFromInt(5),
			expectedErr: false,
		},
		{
			name:        "empty input",
			input:       "",
			expected:    decimal.Zero,
			expectedErr: false,
		},
		{
			name:        "too many numbers",
			input:       "1,2,3",
			expected:    decimal.Zero,
			expectedErr: true,
		},
		{
			name:        "invalid number format",
			input:       "abc,def",
			expected:    decimal.Zero,
			expectedErr: false,
		},
		{
			name:  "decimal addition",
			input: "1.5,2.5",
			expected: func() decimal.Decimal {
				d, _ := decimal.NewFromString("1.5")
				d2, _ := decimal.NewFromString("2.5")
				return d.Add(d2)
			}(),
			expectedErr: false,
		},
		{
			name:        "negative numbers",
			input:       "-1,-2",
			expected:    decimal.NewFromInt(-3),
			expectedErr: false,
		},
		{
			name:        "mixed positive and negative",
			input:       "5,-3",
			expected:    decimal.NewFromInt(2),
			expectedErr: false,
		},
		{
			name:        "large numbers",
			input:       "1000000,2000000",
			expected:    decimal.NewFromInt(3000000),
			expectedErr: false,
		},
		{
			name:  "small decimal numbers",
			input: "0.1,0.2",
			expected: func() decimal.Decimal {
				d, _ := decimal.NewFromString("0.1")
				d2, _ := decimal.NewFromString("0.2")
				return d.Add(d2)
			}(),
			expectedErr: false,
		},
		{
			name:        "missing first number",
			input:       ",5",
			expected:    decimal.NewFromInt(5),
			expectedErr: false,
		},
		{
			name:        "missing second number",
			input:       "5,",
			expected:    decimal.NewFromInt(5),
			expectedErr: false,
		},
		{
			name:        "missing both numbers",
			input:       ",",
			expected:    decimal.Zero,
			expectedErr: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result, err := Add(test.input)
			if test.expectedErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
				assert.Equal(t, test.expected, result)
			}
		})
	}
}
