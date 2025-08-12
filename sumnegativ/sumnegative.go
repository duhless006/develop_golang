package main

import (
	"fmt"
)

func sumNonNegative(numbers []float64) (float64, error) {
	if len(numbers) == 0 {
		return 0, fmt.Errorf("empty slice")
	}
	sum := 0.0
	hasPositive := false
	for _, number := range numbers {
		if number >= 0 {
			sum += number
			if number > 0 {
				hasPositive = true
			}
		}

	}
	if !hasPositive {
		return 0, fmt.Errorf("no positive numbers found")
	}
	return sum, nil
}
func main() {
	tests := [][]float64{
		{1.5, -2, 3, 0},
		{},
		{-1, -2, 0},
		{0, 0},
	}
	for i, test := range tests {
		result, err := sumNonNegative(test)
		if err != nil {
			fmt.Printf("Тест %d, Слайс %v: Ошибка: %v\n", i+1, test, err)
		} else {
			fmt.Printf("Тест %d, Слайс %v: Сумма: %.1f\n", i+1, test, result)
		}
	}

}
