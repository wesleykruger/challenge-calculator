package validate

import (
	"fmt"
	"strings"

	"challenge-calculator/logger"

	"github.com/shopspring/decimal"
)

var defaultDelimiters = []string{",", "\n"}
var customDelimiters []string

func ValidateInput(input string) ([]decimal.Decimal, error) {
	logger.Debug(fmt.Sprintf("Starting input validation: %s", input))

	// Reset custom delimiters for each validation
	customDelimiters = nil

	modifiedInput, err := processCustomDelimiters(input)
	if err != nil {
		return nil, err
	}

	sanitizedValues, err := sanitizeInput(modifiedInput)
	if err != nil {
		return nil, err
	}

	negativeNumbers := findNegativeNumbers(sanitizedValues)
	if len(negativeNumbers) > 0 {
		return nil, fmt.Errorf("invalid input: negative numbers found: %s", strings.Join(negativeNumbers, ", "))
	}

	return sanitizedValues, nil
}

func processCustomDelimiters(input string) (string, error) {
	if !strings.HasPrefix(input, "//") {
		return input, nil
	}

	delimiterEnd := strings.Index(input, "\n")
	if delimiterEnd == -1 {
		return input, nil
	}

	// Extract the delimiter definition part (without the //)
	delimiterDef := input[2:delimiterEnd]

	if strings.HasPrefix(delimiterDef, "[") {
		startIdx := 0
		for startIdx < len(delimiterDef) {
			openBracket := strings.IndexRune(delimiterDef[startIdx:], '[')
			if openBracket == -1 {
				break
			}
			openBracket += startIdx

			closeBracket := strings.IndexRune(delimiterDef[openBracket:], ']')
			if closeBracket == -1 {
				return input, fmt.Errorf("invalid delimiter format: missing closing bracket")
			}
			closeBracket += openBracket

			// Extract the delimiter between brackets
			if closeBracket-openBracket > 1 {
				delimiter := delimiterDef[openBracket+1 : closeBracket]
				if delimiter != "" {
					customDelimiters = append(customDelimiters, delimiter)
				}
			}

			startIdx = closeBracket + 1
		}
	} else {
		if len(delimiterDef) != 1 {
			return input, fmt.Errorf("invalid custom delimiter: %q", delimiterDef)
		}
		customDelimiters = append(customDelimiters, delimiterDef)
	}

	// Return the input with delimiter definition removed
	return input[delimiterEnd+1:], nil
}

func sanitizeInput(input string) ([]decimal.Decimal, error) {
	logger.Debug(fmt.Sprintf("Starting input sanitization: %s", input))

	if len(strings.TrimSpace(input)) == 0 {
		logger.Debug("Empty input received, returning [0]")
		return []decimal.Decimal{decimal.Zero}, nil
	}

	splitValues := splitInput(input)

	var sanitizedValues []decimal.Decimal
	for _, number := range splitValues {
		convertedNumber := parseDecimal(number)
		sanitizedValues = append(sanitizedValues, convertedNumber)
	}

	logger.Debug(fmt.Sprintf("Input sanitization completed: %v", sanitizedValues))
	return sanitizedValues, nil
}

func splitInput(input string) []string {
	trimmedInput := strings.TrimSpace(input)
	result := trimmedInput

	allDelimiters := append([]string{}, customDelimiters...)
	allDelimiters = append(allDelimiters, defaultDelimiters...)

	separator := ","
	for _, delimiter := range allDelimiters {
		result = strings.ReplaceAll(result, delimiter, separator)
	}

	parts := strings.Split(result, separator)
	var cleanParts []string

	for _, part := range parts {
		trimmed := strings.TrimSpace(part)
		if trimmed != "" {
			cleanParts = append(cleanParts, trimmed)
		}
	}

	return cleanParts
}

func parseDecimal(val string) decimal.Decimal {
	if val == "" {
		return decimal.Zero
	}

	number, err := decimal.NewFromString(val)
	if err != nil {
		logger.Debug(fmt.Sprintf("Invalid number format '%s', converting to 0", val))
		return decimal.Zero
	}
	return number
}

func UnescapeNewline(input string) string {
	return strings.ReplaceAll(input, "\\n", "\n")
}

func findNegativeNumbers(numbers []decimal.Decimal) []string {
	var negativeNumbers []string
	for _, number := range numbers {
		if number.Sign() == -1 {
			negativeNumbers = append(negativeNumbers, number.String())
		}
	}
	return negativeNumbers
}
