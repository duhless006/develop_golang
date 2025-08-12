// Улучшенная фильтрация с дополнительной проверкой
package main

import (
	"fmt"
)

func filterPositive(numbers []float64) ([]float64, error) {
	if len(numbers) == 0 {
		return nil, fmt.Errorf("empty slice")
	}
	positives := make([]float64, 0)
	hasPositive := false
	for _, number := range numbers {
		if number >= 0 {
			positives = append(positives, number)
			if number > 0 {
				hasPositive = true
			}
		}

	}
	if !hasPositive {
		return nil, fmt.Errorf("no positive numbers found")
	}
	return positives, nil
}
func main() {
	tests := [][]float64{
		{1.5, -2, 3, 0},
		{},
		{1},
		{-1, -2, 0},
		{0, 0},
	}
	for i, test := range tests {
		result, err := filterPositive(test)
		if err != nil {
			fmt.Printf("Тест %d, Слайс %v: Ошибка: %v\n", i+1, test, err)
		} else {
			fmt.Printf("Тест %d, Слайс %v: Отфильтрованный: %v\n", i+1, test, result) // Замечание 4: Отличное форматирование, но ранний возврат исправлен
		}
	}
}
