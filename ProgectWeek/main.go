package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"unicode/utf8"
)

func countsString(fileName string) (map[string]int, error) {

	file, err := os.Open(fileName)
	if err != nil {
		return nil, fmt.Errorf("Ошибка открытия файла")

	}
	defer file.Close()
	lines := map[string]int{}
	totalLine := 0
	totalWord := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
		//могу засунуть в мейн
		lenLine := len(line)
		count := utf8.RuneCountInString(line)
		totalLine++
	}
	for scanner.Scan() {
		line := scanner.Text()
		fileds := strings.Fields(line)
		totalWord++

	}
	if scanner.Err() != nil {
		return nil, scanner.Err()
	}

}

func main() {
	countsWords, err := countsString("data.txt")
	if err != nil {

	}
}
