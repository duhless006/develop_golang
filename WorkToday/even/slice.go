/*
Описание: Напиши функцию countEven(numbers []int) int,
которая возвращает количество чётных чисел в слайсе целых чисел.
Если слайс пустой, верни 0.
Протестируй с []int{1, 2, 3, 4, 5}, []int{2, 4, 6}, []int{}.
*/
package main

import (
	"fmt"
)

func countEven(numbers []int) int {
	if len(numbers) == 0 {
		return 0
	}
	count := 0
	for _, number := range numbers {
		if number%2 == 0 {
			count += 1
		}
	}
	return count
}
func main() {
	digits := []int{1, 2, 3, 4, 5}
	evenCount := countEven(digits)
	fmt.Println("Тест 1, Слайс: ", digits, "Четных чисел: ", evenCount)

	digits2 := []int{2, 4, 6}
	evenCount2 := countEven(digits2)
	fmt.Println("Тест 2, Слайс: ", digits2, "Четных чисел: ", evenCount2)

	digits3 := []int{}
	evenCount3 := countEven(digits3)
	fmt.Println("Тест 3, Слайс: ", digits3, "Четных чисел: ", evenCount3)
}

/*func main() {
	lists := [][]int{
		{1, 2, 3, 4, 5},
		{2, 4, 6},
		{},
	}
	for i, list := range lists {
		evenCount := countEven(list)
		if evenCount == 0 {
			fmt.Println("Тест", i+1, ", Слайс: ", list, "Четных чисел: ", evenCount)
			continue
		} else {
			fmt.Println("Тест", i+1, ", Слайс: ", list, "Четных чисел: ", evenCount)
		}
	}

}
*/
