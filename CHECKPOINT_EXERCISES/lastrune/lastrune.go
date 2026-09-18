package main

import "fmt"

func LastRune(s string) rune {
	r := []rune(s)
	return r[len(r)-1]
}

func main() {
	fmt.Printf("%c\n", LastRune("hello"))
}