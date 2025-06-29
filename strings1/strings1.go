package main

import (
	"bufio"
	"fmt"
	"os"
	"unicode/utf8"
)

func main() {
	fmt.Println("Введите текст:")
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	scanner := reader.Text()
	fmt.Print(utf8.RuneCountInString(scanner))
}
