package validate

import (
	"fmt"
	"strings"

	"challenge-calculator/logger"

	"github.com/shopspring/decimal"
)

const (
	inputDelimiter = ","
	maxNumbers     = 2
)

func SanitizeInput(input string) ([]decimal.Decimal, error) {
	logger.Debug(fmt.Sprintf("Starting input sanitization: %s", input))

	if len(strings.TrimSpace(input)) == 0 {
		logger.Debug("Empty input received, returning [0]")
		return []decimal.Decimal{decimal.Zero}, nil
	}

	splitValues := splitInput(input)

	if err := validateNumberOfValues(splitValues); err != nil {
		return nil, err
	}

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
	parts := strings.Split(trimmedInput, inputDelimiter)

	// Trim whitespace from each part
	for i, part := range parts {
		parts[i] = strings.TrimSpace(part)
	}

	return parts
}

func validateNumberOfValues(values []string) error {
	if len(values) > maxNumbers {
		return fmt.Errorf("invalid input, a maximum of %d numbers are allowed, received %d", maxNumbers, len(values))
	}
	return nil
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
