package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {
	var brank string
	fmt.Println(brank)

	fmt.Println("foo\nbar")
	fmt.Println(`hoge\nfuga`)
	fmt.Println(`
		hoge

    fuga`)

	var pi, alpha, omega rune = 960, 940, 969
	var bang byte = 33

	fmt.Printf("%v %v %v %v\n", pi, alpha, omega, bang)
	fmt.Printf("%c %c %c %c\n", pi, alpha, omega, bang)

	c := 'y'
	c = c + 3
	if c > 'z' {
		c = c - 26
	}
	fmt.Printf("%c\n\n", c)

	message := "uv vagreangvbany fcnpr fgngvba"
	for i := 0; i < len(message); i++ {
		c := message[i]
		if c >= 'a' && c <= 'z' {
			c = c + 13
			if c > 'z' {
				c = c - 26
			}
		}
		fmt.Printf("%c", c)
	}
	fmt.Print("\n")

	question := "ほげふが?"
	fmt.Println(len(question), "bytes")
	fmt.Println(utf8.RuneCountInString(question), "runes")
	ch, size := utf8.DecodeRuneInString(question)
	fmt.Printf("First rune %c %v bytes\n", ch, size)

	for i, c := range question {
		fmt.Printf("%v %c\n", i, c)
	}
}
