package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"

	"challenge-calculator/calculate"
	"challenge-calculator/logger"
	"challenge-calculator/validate"
)

var logLevel = flag.String("log", "info", "Set the log level (debug, info, error)")

func main() {
	flag.Parse()
	logger.SetLogLevel(logger.LogLevel(*logLevel))

	scanner := bufio.NewScanner(os.Stdin)
	logger.UserMsg("Please enter the numbers to be calculated, separated by a comma:")

	for scanner.Scan() {
		input := scanner.Text()
		unescapedInput := validate.UnescapeNewline(input)

		result, err := calculate.Add(unescapedInput)
		if err != nil {
			logger.UserMsg(fmt.Sprintf("Error calculating result: %v", err))
			os.Exit(1)
		}

		logger.UserMsg(result)
	}

	if err := scanner.Err(); err != nil {
		logger.Error(fmt.Sprintf("Error reading input: %v", err))
		os.Exit(1)
	}
}
