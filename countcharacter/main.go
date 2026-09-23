// Input: a string and a rune (a single charachter) you want to look up in the string 
// Output: returns an int telling us how many times the char appeared in the string
// Edge cases: an empty string returns 0, and a missing character returns 0
// logs: '5'  is the CHARACHTER five, just 5 is the number five so the second test will give 0 bcs they are not the same thing. CountChar(" ", ' ') returns 3 


// Pseudocode: We initialize a counter to 0, then I need to loop through the string charachter by charachter, for each one I check if it is exactly equal to the char we are looking for. If it is we add one to the counter. lastly we return the counter

package main

import (
	"fmt"
)

func main() {
	fmt.Println(CountChar("Hello World", 'l'))
	fmt.Println(CountChar("5  balloons", 5))
	fmt.Println(CountChar("   ", ' '))
	fmt.Println(CountChar("The 7 deadly sins", '7'))
}


func CountChar(str string, c rune) int {
    counter := 0

	for _, charachter := range str {
		if charachter == c {
			counter += 1
		}
	}
	return counter 
}

