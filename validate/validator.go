package validate

import (
	"errors"
	"fmt"
	"regexp"
	"slices"
	"strings"
	"unicode/utf8"

	"challenge-calculator/logger"

	"github.com/shopspring/decimal"
)

var inputDelimiters = []rune{',', '\n'}

func ValidateInput(input string) ([]decimal.Decimal, error) {
	logger.Debug(fmt.Sprintf("Starting input validation: %s", input))

	err := handleCustomDelimiter(input)
	if err != nil {
		return nil, err
	}

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

	splitValues := splitInput(input)

	var sanitizedValues []decimal.Decimal
	for _, number := range splitValues {
		convertedNumber := parseDecimal(number)
		sanitizedValues = append(sanitizedValues, convertedNumber)
	}

	logger.Debug(fmt.Sprintf("Input sanitization completed: %v", sanitizedValues))
	return sanitizedValues, nil
}

func handleCustomDelimiter(input string) error {
	hasCustomDelimiter, customDelimiters, err := checkForCustomDelimiters(input)
	if err != nil {
		return err
	}

	if hasCustomDelimiter {
		inputDelimiters = append(inputDelimiters, customDelimiters...)
	}

	return nil
}

func checkForCustomDelimiters(input string) (bool, []rune, error) {
	// 1. Match multiple delimiters of any length: //[<d1>][<d2>]...\n
	multi := regexp.MustCompile(`^//(\[.*\])\n`)
	if multi.MatchString(input) {
		brackets := regexp.MustCompile(`\[(.+?)\]`)
		matches := brackets.FindAllStringSubmatch(input, -1)

		if len(matches) == 0 {
			return false, nil, nil
		}

		delimiters := make([]rune, 0, len(matches))
		for _, match := range matches {
			if match[1] == "" {
				return false, nil, errors.New("custom delimiter cannot be empty")
			}
			delimiters = append(delimiters, []rune(match[1])...)
		}
		return true, []rune(delimiters), nil
	}

	// 2. Match single bracketed format: //[<delimiter>]\n
	bracketed := regexp.MustCompile(`^//\[(.+)\]\n`)
	if match := bracketed.FindStringSubmatch(input); len(match) == 2 {
		if match[1] == "" {
			return false, nil, errors.New("custom delimiter cannot be empty")
		}
		return true, []rune(match[1]), nil
	}

	// 3. Match single-char format: //{delimiter}\n
	re := regexp.MustCompile(`^//(.+)\n`)
	match := re.FindStringSubmatch(input)

	if len(match) <= 1 {
		return false, nil, nil
	}

	delimiterStr := match[1]
	if utf8.RuneCountInString(delimiterStr) != 1 {
		return false, nil, fmt.Errorf("invalid custom delimiter %q", delimiterStr)
	}

	return true, []rune(delimiterStr), nil
}

func splitInput(input string) []string {
	trimmedInput := strings.TrimSpace(input)
	parts := strings.FieldsFunc(trimmedInput, func(r rune) bool {
		return slices.Contains(inputDelimiters, r)
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
