package calculate

import (
	"fmt"
	"strings"

	"challenge-calculator/logger"
	"challenge-calculator/validate"

	"github.com/shopspring/decimal"
)

const maxValidNumber = 1000

func Add(input string) (string, error) {
	logger.Debug(fmt.Sprintf("Starting addition calculation for input: %s", input))
	sum := decimal.Zero
	var formulaParts []string

	numbers, err := validate.ValidateInput(input)
	if err != nil {
		logger.Error(fmt.Sprintf("Error validating input: %v", err))
		return "", err
	}

	for _, num := range numbers {
		if !numberExceedsMaxValue(num) {
			logger.Debug(fmt.Sprintf("Adding number %s to sum", num.String()))
			sum = sum.Add(num)
			formulaParts = append(formulaParts, num.String())
		} else {
			logger.Debug(fmt.Sprintf("Number %s is too large, omitting from sum", num.String()))
			formulaParts = append(formulaParts, "0")
		}
	}

	formula := strings.Join(formulaParts, "+")
	if len(formulaParts) > 0 {
		formula += " = " + sum.String()
	} else {
		formula = "0 = 0"
	}

	logger.Debug(fmt.Sprintf("Calculation completed: %s", formula))
	return formula, nil
}

func numberExceedsMaxValue(number decimal.Decimal) bool {
	return number.GreaterThan(decimal.NewFromInt(maxValidNumber))
}
