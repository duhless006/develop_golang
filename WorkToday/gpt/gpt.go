package main

import (
	"errors"
	"fmt"
)

func average(numbers []float64) (float64, error) {
	sum := 0.0
	if len(numbers) == 0 {
		return 0, errors.New("empty slice")
	}

	for _, number := range numbers {
		if number < 0 {
			return 0, errors.New("negative number")
		}
		sum += number
	}
	return sum / float64(len(numbers)), nil
}

func main() {
	testAverage := []float64{1, 2, 3}
	result, err := average(testAverage)
	if err != nil {
		fmt.Println("Ошибка :", err)

	}
	fmt.Println("Слайс ", testAverage, ":", "Среднее :", result)

	testAverage2 := []float64{}
	result2, err := average(testAverage2)
	if err != nil {
		fmt.Println("Слайс ", testAverage2, ":", "Ошибка:", err)

	} else {
		fmt.Println("Слайс ", testAverage2, ":", "Среднее :", result2)
	}

	testAverage3 := []float64{1, -2, 3}
	result3, err := average(testAverage3)
	if err != nil {
		fmt.Println("Слайс ", testAverage3, ":", "Ошибка:", err)

	} else {
		fmt.Println("Слайс ", testAverage3, ":", "Среднее :", result3)
	}
}
