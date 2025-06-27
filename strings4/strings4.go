package main

import (
	"bufio"
	"fmt"
	"os"
)

func check(form string) {
	var open_bracket int
	var close_bracket int
	for _, k := range form {
		if k == '(' {
			open_bracket++
		}
		if k == ')' {
			close_bracket++
		}
	}
	if open_bracket == close_bracket {
		fmt.Printf("Скобки расставлены верно, %d открывающиеся, %d закрывающиеся\n", open_bracket, close_bracket)
	} else {
		fmt.Printf("Скобки расставлены неправильно, %d открывающиеся, %d закрывающиеся\n", open_bracket, close_bracket)
	}
}
func main() {
	fmt.Println("Введите формулу:")
	reader := bufio.NewReader(os.Stdin)
	formula, _ := reader.ReadString('\n')
	check(formula)

}
