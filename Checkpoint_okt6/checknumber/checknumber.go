package main

import "fmt"

func CheckNumber(arg string) bool {
	for _, character := range arg { // both the byte-index and the rune (character) at that position
		if character >= '0' && character <= '9' { // '0' is not the string "0", we need the integer value so use '
			return true
		} else {
			continue
		}
	}
	return false
}


func main() {
	fmt.Println(CheckNumber("Hello"))
	fmt.Println(CheckNumber("Hello1"))
}