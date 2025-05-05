package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

// Функция перевода десятичного числа в указанную систему счисления
func decToBase(n int, base int) string {
	const table = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	if n < base {
		return string(table[n])
	}
	return decToBase(n/base, base) + string(table[n%base])
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var fromBase, toBase int

	// Ввод начальной системы счисления
	for {
		fmt.Print("Из какой системы счисления надо будет переводить? - ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		n, err := strconv.Atoi(input)
		if err == nil && n >= 2 && n <= 36 {
			fromBase = n
			break
		}
		fmt.Println("Введи целое число от 2 до 36")
	}

	// Ввод конечной системы счисления
	for {
		fmt.Print("В какую систему счисления надо будет переводить? - ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		n, err := strconv.Atoi(input)
		if err == nil && n >= 2 && n <= 36 {
			toBase = n
			break
		}
		fmt.Println("Введи целое число от 2 до 36")
	}

	// Ввод числа и перевод в десятичную систему
	var value int
	for {
		fmt.Print("Какое число надо будет переводить? - ")
		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)
		n, err := strconv.ParseInt(input, fromBase, 64)
		if err == nil {
			value = int(n)
			break
		}
		fmt.Println("Введи любое целое число. Примечание: цифры в числе не должны быть меньше системы счисления, из которой переводишь")
	}

	// Вывод результата
	fmt.Println("Результат:", decToBase(value, toBase))

	// Пауза
	fmt.Println("\nНажмите ENTER для выхода.")
	reader.ReadString('\n')
}
