package datafile

import (
	"bufio"
	"os"
)

func GetStrings(fileName string) ([]string, error) {
	var lines []string             // переменная для хранения строк
	file, err := os.Open(fileName) // открываем файл
	if err != nil {
		return nil, err
	}
	scanner := bufio.NewScanner(file) // создаем новый сканер для чтения файла
	for scanner.Scan() {              // читаем файл построчно
		line := scanner.Text()      // получаем текущую строку
		lines = append(lines, line) // добавляем строку в срез
	}
	err = file.Close() // закрываем файл после чтения
	if err != nil {
		return nil, err
	}
	if scanner.Err() != nil { // проверяем на ошибки при чтении
		return nil, scanner.Err() // если есть ошибка, возвращаем ее
	}
	return lines, nil
}
