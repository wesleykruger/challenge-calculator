package main

import (
	"testing"

	"challenge-calculator/calculate"
	"challenge-calculator/logger"

	"github.com/stretchr/testify/assert"
)

func TestMainFunctionality(t *testing.T) {
	t.Run("log level setting", func(t *testing.T) {
		*logLevel = "info"
		main()
		assert.Equal(t, logger.LogLevelInfo, logger.LogLevel(*logLevel))

		*logLevel = "debug"
		main()
		assert.Equal(t, logger.LogLevelDebug, logger.LogLevel(*logLevel))

		*logLevel = "error"
		main()
		assert.Equal(t, logger.LogLevelError, logger.LogLevel(*logLevel))
	})

	t.Run("calculation", func(t *testing.T) {
		tests := []struct {
			name        string
			input       string
			expected    string
			expectedErr bool
		}{
			{
				name:        "simple addition",
				input:       "1,2",
				expected:    "1+2 = 3",
				expectedErr: false,
			},
			{
				name:        "single number",
				input:       "5",
				expected:    "5 = 5",
				expectedErr: false,
			},
			{
				name:        "empty input",
				input:       "",
				expected:    "0 = 0",
				expectedErr: false,
			},
			{
				name:        "many numbers",
				input:       "1,2,3",
				expected:    "1+2+3 = 6",
				expectedErr: false,
			},
			{
				name:        "invalid numbers",
				input:       "abc,def",
				expected:    "0+0 = 0",
				expectedErr: false,
			},
		}

		for _, test := range tests {
			t.Run(test.name, func(t *testing.T) {
				sum, err := calculate.Add(test.input)
				if test.expectedErr {
					assert.Error(t, err)
				} else {
					assert.NoError(t, err)
					assert.Equal(t, test.expected, sum)
				}
			})
		}
	})
}
