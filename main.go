package main

import (
	"bufio"
	"fmt"
	"os"
)

func reverseString(s string) string {
	r := []rune(s)
	l := len(r)
	for i := 0; i < l/2; i++ {
		j := l - i - 1
		r[i], r[j] = r[j], r[i]
	}
	return string(r)
}

func main() {
	fmt.Println("Введите строку: ")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	s := scanner.Text()
	fmt.Println(reverseString(s))
}
