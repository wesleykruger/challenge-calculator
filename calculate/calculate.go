package calculate

import (
	"fmt"

	"challenge-calculator/logger"
	"challenge-calculator/validate"

	"github.com/shopspring/decimal"
)

const maxValidNumber = 1000

func Add(input string) (decimal.Decimal, error) {
	logger.Debug(fmt.Sprintf("Starting addition calculation for input: %s", input))
	sum := decimal.Zero

	numbers, err := validate.ValidateInput(input)
	if err != nil {
		logger.Error(fmt.Sprintf("Error validating input: %v", err))
		return decimal.Zero, err
	}

	for _, num := range numbers {
		if !numberExceedsMaxValue(num) {
			logger.Debug(fmt.Sprintf("Number %s is too large, omitting from sum", num.String()))
			sum = sum.Add(num)
		}
	}

	formattedResult := sum.String()
	logger.Debug(fmt.Sprintf("Calculation completed: %v = %v", numbers, formattedResult))
	return sum, nil
}

func numberExceedsMaxValue(number decimal.Decimal) bool {
	return number.GreaterThan(decimal.NewFromInt(maxValidNumber))
}
