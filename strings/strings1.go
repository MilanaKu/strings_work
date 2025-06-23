package main

import "fmt"

func main() {
	var text string
	fmt.Println("Введите текст:")
	fmt.Scanln(&text)
	var k int
	for range text {
		k++
	}
	fmt.Println("Количество символов:", k)
}
