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
		first_letter := strings.ToUpper(string(k[0]))
		rest_letters := strings.ToLower(string(k[1:]))
		split_s[i] = first_letter + rest_letters
	}
	return strings.Join(split_s, " ")
}
func min() {
	fmt.Println("Введите текст:")
	reader := bufio.NewScanner(os.Stdin)
	text := reader.Text()
	changed_text := capitalizeWords(text)
	fmt.Println(changed_text)

}
