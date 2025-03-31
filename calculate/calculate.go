package calculate

import (
	"fmt"

	"challenge-calculator/logger"
	"challenge-calculator/validate"

	"github.com/shopspring/decimal"
)

func Add(input string) (decimal.Decimal, error) {
	logger.Debug(fmt.Sprintf("Starting addition calculation for input: %s", input))
	sum := decimal.Zero

	numbers, err := validate.ValidateInput(input)
	if err != nil {
		logger.Error(fmt.Sprintf("Error validating input: %v", err))
		return decimal.Zero, err
	}

	for _, num := range numbers {
		sum = sum.Add(num)
	}

	formattedResult := sum.String()
	logger.Debug(fmt.Sprintf("Calculation completed: %v = %v", numbers, formattedResult))
	return sum, nil
}
