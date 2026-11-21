package main

import "fmt"

func main() {
	/*
		Ошибка 1: нет закрывающей кавычки
		main.go:6:22: newline in string
		Строка не закрыта и тянется до конца строки файла.
	*/

	// fmt.Println(Viktor")

	fmt.Println("Linux")
	fmt.Println("software developer Go backend")

	/*
		Ошибка 2: нет последней }
		main.go:10:1: syntax error: unexpected EOF, expected }
		Файл закончился, а компилятор ждал закрывающую скобку.
	*/
}
