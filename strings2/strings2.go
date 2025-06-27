package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Введите текст:")
	reader := bufio.NewScanner(os.Stdin)
	reader.Scan()
	text := reader.Text()
	var a, b, c, d, e, f, g, h, i, j int
	for _, k := range text {
		if k == 'а' || k == 'А' {
			a++
		}
		if k == 'е' || k == 'Е' {
			b++
		}
		if k == 'ё' || k == 'Ё' {
			c++
		}
		if k == 'и' || k == 'И' {
			d++
		}
		if k == 'о' || k == 'О' {
			e++
		}
		if k == 'у' || k == 'У' {
			f++
		}
		if k == 'ы' {
			g++
		}
		if k == 'э' || k == 'Э' {
			h++
		}
		if k == 'ю' || k == 'Ю' {
			i++
		}
		if k == 'я' || k == 'Я' {
			j++
		}
	}
	fmt.Printf("а-%d\nе-%d\nё-%d\nи-%d\nо-%d\nу-%d\nы-%d\nэ-%d\nю-%d\nя-%d\n", a, b, c, d, e, f, g, h, i, j)

}
