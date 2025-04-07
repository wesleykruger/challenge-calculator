# Challenge Calculator

A simple calculator that adds numbers from a string input.

## Features

- Adds numbers from a string input
- Handles decimal numbers
- Handles empty input
- Handles invalid input
- Handles missing numbers
- Handles whitespace
- Uses decimal arithmetic for precise calculations

## Technical Details

### Decimal Arithmetic

This calculator uses the `github.com/shopspring/decimal` library for all numerical operations.
- Unlike floating-point arithmetic (float32/float64), decimal arithmetic provides exact decimal representation. This eliminates rounding errors that can occur with floating-point calculations.
- The decimal library can also handle much larger numbers than float64 without loss of precision, making it suitable for financial and scientific calculations. It also helps avoid overflows.

### Input Format

The calculator accepts input in the following format:
- Numbers separated by a delimiter.
- Numbers must be positive (unless allowed via the `allowNegatives` argument).
- Numbers greater than or equal to the max allowed value will be omitted from the sum. By default, this is 1000.
- Numbers can be decimal or whole.
- Empty input is allowed.
- Missing numbers are treated as zero.
- Whitespace is allowed around numbers and delimiters.

### Delimiters
The calculator accepts input separated by the following delimters:
- Commas (,)
- Newline characters ('\n'). If the `defaultDelimiter` argument is used, this can be changed.
- Custom Single-Character Delimiters: Supports the use of a single character using the format `//{delimiter}\n{numbers}`
- Custom Multi-Character Delimiters: Supports the use multiple characters as a single delimiter using the format `//[{delimiter1}][{delimiter2}]...\n{numbers}`

### Arguments
The calculator accepts the following arguments on startup:
- logLevel: Determines the application log level
- defaultDelimiter: Allows for an alternate default delmiter in addition to ",". If this argument is omitted, the system will default to the newline character "/n".
- allowNegatives: If set to true, negative numbers will be allowed in calculations.
- maxNumber: Accepts an integer which can be used as the maximum allowed value in a calculation. If omitted, this will default to 1000.

### Logging

The project uses Zerolog for structured logging and accepts a flag at runtime to set the logging level. If no flag is specified, it will default to Info level.

Example:
`go run main.go -log debug`

## Usage

```bash
go run main.go
```

Then enter numbers in the format: `number1,number2`
Invalid or missing values will be treated as 0 for the purpose of calculating values.

## Testing

Run the test suite:
```bash
go test ./...
```

## Dependencies

- github.com/shopspring/decimal
- github.com/stretchr/testify
- github.com/rs/zerolog
