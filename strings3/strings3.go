package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func capitalizeWords(s string) string {
	split_s := strings.Split(s, " ")
	for i, k := range split_s {
		runes := []rune(k)
		first_letter := strings.ToUpper(string(runes[0]))
		rest_letters := strings.ToLower(string(runes[1:]))
		split_s[i] = first_letter + rest_letters
	}
	return strings.Join(split_s, " ")
}
func main() {
	fmt.Println("Введите текст:")
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	text := reader.Text()
	changed_text := capitalizeWords(text)
	fmt.Println(changed_text)

}

