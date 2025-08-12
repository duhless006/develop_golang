package main

import (
	"fmt"
	"os"
	"strconv"
)

func addition(a float64, b float64) float64 {
	return a + b
}

func subtraction(a float64, b float64) float64 {
	return a - b
}

func multiplication(a float64, b float64) float64 {
	return a * b
}

func division(a float64, b float64) float64 {
	if b == 0 {
		return 0
	}
	return a / b
}
func main() {
	if len(os.Args) != 4 {
		fmt.Println("Ошибка! нужно ввести оператор и 2 числа!\n")
		return
	}
	operator := os.Args[1]
	a, err1 := strconv.ParseFloat(os.Args[2], 64)
	b, err2 := strconv.ParseFloat(os.Args[3], 64)

	if err1 != nil || err2 != nil {
		fmt.Println("Ошибка! Ввведите 2 аргумента в виде чисел!")
		return
	}

	switch operator {
	case "+":
		resultA := addition(a, b)
		fmt.Printf("%.2f + %.2f = %.2f\n", a, b, resultA)

	case "-":
		resultS := subtraction(a, b)
		fmt.Printf("%.2f - %.2f = %.2f\n", a, b, resultS)

	case "**":
		resultM := multiplication(a, b)
		fmt.Printf("%.2f * %.2f = %.2f\n", a, b, resultM)

	case "/":
		resultD := division(a, b)
		if b == 0 {
			fmt.Println("На 0 делить нельзя!\n")
			break

		} else {
			fmt.Printf("%.2f / %.2f = %.2f\n", a, b, resultD)
		}
	default:
		fmt.Println("Оператор не распознан!\n")

	}
}
