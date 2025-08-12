package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)
func sumNonNegative(input []float64) (float64, error) {
	if len(input) == 0 {
		return 0, fmt.Errorf("error input")
	}
	for _, inp := range input {
		if inp < 0 {
			return 0, fmt.Errorf("Ошибка: negative number not allowed")
		} else {
			sum := 0.0
			for _, number := reader numbers {
				if number >= 0 {
					sum += number
				}
			}
			return sum, nil
		}
	
	}
	return 0, nil
}
func main() {
	for i := 0; i != "stop"; i++ {
		fmt.Printin("Введите число (или 'stop' для завершения): ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			og.Fatal(err)
		}
		if input == "stop" {
			fmt.Printf("Слайс: %.1f\n Сумма: %.1f\n")
		}
		input = strings.TrimSpace(input)
		output, err := strconv.ParseFloat(input, 64)
		if err != nil {
			log.Fatal(err)
		}
	
}
