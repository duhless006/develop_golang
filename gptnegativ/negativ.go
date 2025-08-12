// Задача 1: Фильтрация отрицательных чисел, Тема: Слайсы, обработка ошибок
package main

import (
	"fmt"
)

func filterPositive(numbers []float64) ([]float64, error) {
	if len(numbers) == 0 {
		return nil, fmt.Errorf("empty slice : %v", numbers) // достаточно будет : fmt.Errorf("empty slice").
	}
	var negatives []float64 // лучше использовать майк negatives := make([]float64, 0)
	for _, number := range numbers {
		if number >= 0 {
			negatives = append(negatives, number)
		}
	}
	return negatives, nil
}

func main() {
	negativ := []float64{1.5, -2.0, 3.0, 0.0}
	result, err := filterPositive(negativ)
	if err != nil {
		fmt.Println("error", err)
		return
	}
	fmt.Println("Слайс ", negativ, ":", "Отфильтрованный: ", result) //  ранний выход из-за ретерна, лучший способ сделать как во втором тесте

	negativ2 := []float64{}
	result2, err := filterPositive(negativ2)
	if err != nil {
		fmt.Println("Слайс ", negativ2, "Ошибка: empty slice") // тут жесткий вывод, лучше юзать fmt.Println("Ошибка:", err).

	} else {
		fmt.Println("Слайс ", negativ2, ":", "Отфильтрованный: ", result2)
	}
	negativ3 := []float64{1.0}
	result3, err := filterPositive(negativ3)
	if err != nil {
		fmt.Println("Слайс ", negativ3, ":", "Ошибка: ", err)
		return

	} else {
		fmt.Println("Слайс ", negativ3, ":", "Отфильтрованный: ", result3)
	}
	// Отсутствие проверки на случай, когда все числа отрицательные:
	// В filterPositive нет проверки на случай, когда все числа отрицательные
	// (например, [-1, -2]). Это не было в изначальной задаче,
	// но добавление такой проверки сделает функцию более надежной.
}

/*
//func main() {
	tests := [][]float64{
		{1.5, -2.0, 3.0, 0.0}, // Нормальный слайс
		{},                     // Пустой слайс
		{1.0},                 // Слайс с одним элементом
		{-1.0, -2.0},         // Все отрицательные
	}

	for i, test := range tests {
		result, err := filterPositive(test)
		if err != nil {
			fmt.Printf("Тест %d, Слайс %v: Ошибка: %v\n", i+1, test, err)
		} else {
			fmt.Printf("Тест %d, Слайс %v: Отфильтрованный: %v\n", i+1, test, result)
		}
	}
}
*/
