/*
Описание: Напиши программу, которая запрашивает строку чисел, разделенных пробелами (например, "1.5 2.5 0").
Преобразуй строку в слайс []float64, исключая отрицательные числа,
и передай его в sumNonNegative. Повторяй ввод при некорректных данных.
*/
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func sumNonNegative(numbers []float64) (float64, error) {
	if len(numbers) == 0 { // если длинна строки равна 0 возвращаем ошибку
		return 0, fmt.Errorf("empty slice")
	}
	sum := 0.0                       // инициализируем переменную sum для хранения суммы
	hasPositive := false             // флаг для проверки положительных чисел
	for _, number := range numbers { // перебераем числа в слайсе
		if number >= 0 { // если число больше или =0
			sum += number   // суммируем его в переременную
			if number > 0 { // если число больше 0
				hasPositive = true // флаг становится истиной
			}
		}

	}
	if !hasPositive { // если флаг лож возвращаем ошибку
		return 0, fmt.Errorf("no numbers greater than zero found")

	}
	return sum, nil // возвращаем сумму и ошибку если есть
}
func main() {
	reader := bufio.NewReader(os.Stdin) // запрашиваем ввод от пользователя
	numbers := make([]float64, 0)       // создаем дополнительный слайс
	for {
		fmt.Println("Введите число (или 'stop' для завершения): ")
		input, err := reader.ReadString('\n') // ввод пользователя через пробел
		if err != nil {                       // при возникновении ошибки выводим ее и продолжаем цикл
			fmt.Printf("ошибка чтения: %v", err)
			continue
		}

		input = strings.TrimSpace(input) // убираем лишние пробелы

		if strings.TrimSpace(input) == "stop" { //если ввод равен стоп
			if len(numbers) == 0 { // если длинна слайса равна 0 печатаем ошибку
				fmt.Println("Ошибка: no numbers entered")
				return
			}

			sum, err := sumNonNegative(numbers) // инициализируем в нашу переменную функцию
			if err != nil {                     // если есть ошибка печатаем ее
				fmt.Printf("Слайс: %v\nОшибка: %v\n", numbers, err)
			} else { // иначе выводим результат
				fmt.Printf("Слайс: %v\nСумма: %.2f\n", numbers, sum)
			}
			return
		}

		num, err := strconv.ParseFloat(input, 64) // преобразуем нашу пременную к флоату
		if err != nil {                           // при возникновении ошибки выводим ее и продолжаем
			fmt.Printf("Ошибка: invalid number %q\n", input)
			continue
		}

		if num < 0 { //если наше число меньше 0 выводим ошибку
			fmt.Printf("Ошибка: negative number not allowed\n")
			continue
		}
		numbers = append(numbers, num) // добавляем значения в переменную numbers

	}
}
