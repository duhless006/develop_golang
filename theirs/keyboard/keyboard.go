package keyboard

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func GetFloat() (float64, error) {
	reader := bufio.NewReader(os.Stdin)   // буферизированный считыватель
	input, err := reader.ReadString('\n') // получаем данные от пользователя и завершаем чтение после нажатия enter
	if err != nil {                       //если ошибка не равна nil // сохраняем ее в переменную err
		return 0, err
	}

	input = strings.TrimSpace((input))           // удаляем пробельные символы, пропуски
	number, err := strconv.ParseFloat(input, 64) // преобразуем из строки в формат float64
	if err != nil {                              // проверка если не равно nil
		return 0, err
	}
	return number, err
}
