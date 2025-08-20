задача 1
package main

import (
	"fmt"
)

func main() {
	var nuts = [3]int{10, 0, 8}
	sumNuts := 0
	for _, nut := range nuts {
		if nut == 0 {
			fmt.Println("В массиве по одной из позиций закончились орехи")
		}
		sumNuts += nut
	}
	fmt.Println("Общее количесво орехов на складе: ", sumNuts)

}

задача 2
package main

import (
	"fmt"
)

func main() {
	nuts := []int{}
	nuts = append(nuts, 10)
	nuts = append(nuts, 5)
	nuts = append(nuts, 8)
	fmt.Println("длинна: ", len(nuts))
	fmt.Println("емкость: ", cap(nuts))
	nuts = nuts[:len(nuts)-1]
	fmt.Println("длинна после удаления последнего элемента: ", len(nuts))
	fmt.Println("емкость после удаления последнего элемента: ", cap(nuts))

}

задача 3
package main

import (
	"fmt"
)

func main() {
	nuts := make(map[string]int)
	nuts["арахис"] = 10
	nuts["фундук"] = 8
	nuts["грецкий"] = 5
	value, ok := nuts["кешью"]
	if ok {
		fmt.Printf("кешью %d\n", value)
	} else {
		fmt.Println("Кешью не найден")
	}
	nuts["кешью"] = 3
	delete(nuts, "грецкий")
	for nut, value := range nuts {
		fmt.Printf("Добавляем %s: %d\n", nut, value)
	}
}

задача 4

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func sumNuts(fileName string) (int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return 0, fmt.Errorf("ошибка открытия файла: %v", err)
	}
	defer file.Close()
	count := 0
	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		if count > 2 {
			return 0, fmt.Errorf("Слишком много строк в файле %v")
		}
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fileds) != 2 {
			return 0, fmt.Errorf("неверный фрмат строки")
		}
		total, err := strconv.Atoi(fileds[1])
		if err != nil {
			return 0, fmt.Errorf("Ошибка преобразования числа!")
		}
		sum += total
		count++
	}
	
	if err != nil {
		return 0, fmt.Errorf("ошибка закрытия файла!")
	}
	if scanner.Err() != nil {
		return 0, scanner.Err()
	}

	return sum, nil
}
func main() {
	nuts, err := sumNuts("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Всего орехов на складе: %d\n", nuts)
}

задача 5

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func addNuts(fileName string) ([]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("error open file %v")
	}
	defer file.Close()
	nutCount := []int{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if len(fileds) != 2 {
			return nil, fmt.Errorf("Ожидаемое максимальное количество строк в файле не более 3-х! %v")
		}
		total, err := strconv.Atoi(fileds[1])

		if err != nil {
			return nil, fmt.Errorf("Ошибка преобразования числа! %v")
		}
		nutCount = append(nutCount, total)

	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("ошибка чтения файла: %v", err)
	}
	sort.Ints(nutCount); sort.Sort(sort.Reverse(sort.IntSlice(nutCount)))
	nut := 3
	nutCount = append(nutCount, nut)
	return nutCount, nil
}
func main() {
	nuts, err := addNuts("input.txt")
	if err != nil {
		fmt.Println("Ошибка:", err)
        return
	}

	fmt.Println(nuts)

}

задача 6

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
)

func countNuts(fileName string) (map[string]int, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("error open file %v", err)
	}
	defer file.Close()
	nutCount := make(map[string]int)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines := scanner.Text()
		fields := strings.Fields(line)
        if len(fields) < 1 {
            continue
        }
        nutCount[fields[0]]++
	}
	if err := scanner.Err(); err != nil {
        return nil, fmt.Errorf("ошибка чтения файла: %v", err)
    }
	return nutCount, nil
}
func main() {
	nuts, err := countNuts("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	keys := make([]string, 0, len(nuts))
	for nut := range nuts {
		keys = append(keys, nut)
	}
	sort.Strings(keys)

	for _, nut := range keys {
		count := nuts[nut]
		fmt.Printf("%s: %d\n", nut, count)
	}

}

задача 7

package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)

func sliceMapMass(fileName string) ([]string, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("error open file %v", err)
	}
	defer file.Close()

	var nutCount []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := strings.ToLower(scanner.Text())
		fields := strings.Fields(line)
		nutCount = append(nutCount, fields...)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error scanning file: %v", err)
	}
	return nutCount, nil

}

func readRunes(fileName string) ([]rune, error) {
	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("error opening file: %v", err)
	}
	defer file.Close()

	var runes []rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		runes = append(runes, []rune(line)...)
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error scanning file: %v", err)
	}

	return runes, nil
}

type WordCount struct {
	Word  string
	Count int
}

func main() {
	totals, err := sliceMapMass("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	wordFreq := make(map[string]int)
	for _, total := range totals {
		wordFreq[total]++
	}
	var wordCounts []WordCount
	for word, count := range wordFreq {
		wordCounts = append(wordCounts, WordCount{Word: word, Count: count})
	}
	sort.Slice(wordCounts, func(i, j int) bool {
		return wordCounts[i].Count > wordCounts[j].Count
	})

	var topThree [3]string
	for i := 0; i < len(wordCounts) && i < 3; i++ {
		topThree[i] = wordCounts[i].Word
	}

	fmt.Println("Топ-3 самых частых слов:")
	for i, word := range topThree {
		if word != "" { // Пропускаем пустые элементы, если слов меньше 3
			fmt.Printf("%d. %s (встречается %d раз)\n", i+1, word, wordFreq[word])
		}
	}
	runes, err := readRunes("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	runeFreq := make(map[rune]int)
	for _, r := range runes {
		runeFreq[r]++
	}

	// Преобразование map в срез для сортировки
	type RuneCount struct {
		Rune  rune
		Count int
	}
	var runeCounts []RuneCount
	for r, count := range runeFreq {
		runeCounts = append(runeCounts, RuneCount{Rune: r, Count: count})
	}

	// Сортировка по убыванию частоты
	sort.Slice(runeCounts, func(i, j int) bool {
		return runeCounts[i].Count > runeCounts[j].Count
	})

	// Вывод топ-3 символов
	fmt.Println("\nТоп-3 самых частых символов:")
	for i := 0; i < len(runeCounts) && i < 3; i++ {
		fmt.Printf("%d. %q (встречается %d раз)\n", i+1, runeCounts[i].Rune, runeCounts[i].Count)
	}
}
