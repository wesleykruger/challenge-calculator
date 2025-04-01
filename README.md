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
- Two numbers separated by a delimiter
- Numbers must be positive
- Numbers greater than or equal to 1000 will be omitted from the sum
- Numbers can be decimal or whole
- Empty input is allowed
- Missing numbers are treated as zero
- Whitespace is allowed around numbers and delimiters

### Delimiters
The Calculator accepts input separated by the following delimters:
- Commas (,)
- Newline characters ('\n')
- Custom Delimiters: Supports the use of a single character using the format `//{delimiter}\n{numbers}`

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
