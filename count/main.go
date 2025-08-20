package main

import (
	"fmt"
	"sort"
)

// функция считает кол-во голосов и выводит их. в алфавитном порядке
func main() {
	grades := map[string]float64{"Alma": 74.2, "Rohit": 86.5, "Carl": 59.7}
	var names []string
	for name := range grades { // перебираем ключи в мапе grades
		names = append(names, name) // добавляем ключи в слайс names
	}
	sort.Strings(names)          // сортируем слайс names в алфавитном порядке
	for _, name := range names { // перебираем отсортированный слайс names
		fmt.Printf("%s has a grade of %0.1f%%\n", name, grades[name])
	}
}

/* func main() {
lines, err := datafile.GetStrings("votes.txt")
if err != nil {
log.Fatal(err)
}
counts := make(map[string]int)
for _, line := range lines {
counts[line]++
}
for name, count := range counts {
fmt.Printf("Votes for %s: %d\n", name, count)
}
}
*/

/* var names []string           // слайс для хранения имен
var counts []int             // слайс для хранения количества голосов
for _, line := range lines { // перебираем каждую строку из файла
	matched := false             // флаг для проверки, найдено ли имя
	for i, name := range names { // перебираем имена в слайсе names
		if name == line { // если имя совпадает с текущей строкой
			counts[i]++    // увеличиваем счетчик голосов для этого имени
			matched = true // устанавливаем флаг в true
		}
	}
	//	если имя не найдено в слайсе names
	if matched == false {
		names = append(names, line) // добавляем имя в слайс names
		counts = append(counts, 1)  // добавляем счетчик голосов для нового имени
	}
}
for i, name := range names { // перебираем имена в слайсе names
	fmt.Printf("%s: %d\n", name, counts[i]) // выводим имя и количество голосов
}
*/
