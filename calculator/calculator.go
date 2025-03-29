package calculator

import (
	"math"
)

func Calculator(number1 float32, operator string, numbers ...float32) float32 {
	var number2 float32
	if len(numbers) > 0 {
		number2 = numbers[0]
	}

	switch operator {
	case "+":
		return number1 + number2
	case "-":
		return number1 - number2
	case "*":
		return number1 * number2
	case "/":
		if number2 == 0 {
			panic("Cannot divide by zero")
		}
		return number1 / number2
	case "%":
		return float32(int(number1) % int(number2))
	case "^":
		return float32(math.Pow(float64(number1), float64(number2)))
	case "sqrt":
		if number1 < 0 {
			panic("Undefined")
		}
		return float32(math.Sqrt(float64(number1)))
	case "OR":
		var answer int = int(number1) | int(number2)
		return float32(answer)
	case "AND":
		var answer int = int(number1) & int(number2)
		return float32(answer)
	case "XOR":
		var answer int = int(number1) ^ int(number2)
		return float32(answer)
	default:
		panic("Invalid operator: " + operator)
	}
}
