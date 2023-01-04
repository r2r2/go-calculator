package calculator

import "fmt"

func SimpleCalculator(a, b float64, operator string) (float64, error) {
	switch operator {
	case "+":
		return add(a, b), nil
	case "-":
		return subtract(a, b), nil
	case "/":
		return divide(a, b)
	case "*":
		return multiply(a, b), nil
	}
	return 0, fmt.Errorf("unknown operator: %s", operator)
}

func add(a, b float64) float64 {
	return a + b
}

func subtract(a, b float64) float64 {
	return a - b
}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, fmt.Errorf("error: division by zero")
	}
	return a / b, nil
}

func multiply(a, b float64) float64 {
	return a * b
}
