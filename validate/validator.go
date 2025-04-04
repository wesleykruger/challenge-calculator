package validate

import (
	"fmt"
	"slices"
	"strings"

	"challenge-calculator/logger"

	"github.com/shopspring/decimal"
)

var inputDelimiter = []rune{',', '\n'}

func ValidateInput(input string) ([]decimal.Decimal, error) {
	logger.Debug(fmt.Sprintf("Starting input validation: %s", input))

	sanitizedValues, err := sanitizeInput(input)
	if err != nil {
		return nil, err
	}

	negativeNumbers := findNegativeNumbers(sanitizedValues)
	if len(negativeNumbers) > 0 {
		return nil, fmt.Errorf("invalid input: negative numbers found: %s", strings.Join(negativeNumbers, ", "))
	}

	return sanitizedValues, nil
}

func sanitizeInput(input string) ([]decimal.Decimal, error) {
	logger.Debug(fmt.Sprintf("Starting input sanitization: %s", input))

	if len(strings.TrimSpace(input)) == 0 {
		logger.Debug("Empty input received, returning [0]")
		return []decimal.Decimal{decimal.Zero}, nil
	}

	unescapedInput := unescapeNewline(input)

	splitValues := splitInput(unescapedInput)

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
	parts := strings.FieldsFunc(trimmedInput, func(r rune) bool {
		return slices.Contains(inputDelimiter, r)
	})

	// Trim whitespace from each part
	for i, part := range parts {
		parts[i] = strings.TrimSpace(part)
	}

	return parts
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

func unescapeNewline(input string) string {
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
