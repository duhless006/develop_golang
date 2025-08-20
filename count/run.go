package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	nuts := make(map[string]int)
	nuts["арахис"] = 10
	nuts["кешью"] = 5
	nuts["грецкий"] = 8
	//fmt.Println(nuts)
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	checkNuts := strings.TrimSpace(scanner.Text())
	count, ok := nuts[checkNuts]
	if ok {
		fmt.Printf("%s: %d\n", checkNuts, count)

	} else {
		fmt.Println("таких орехов не найдено\n")
	}

}
