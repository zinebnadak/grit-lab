package main
// Input: a string 
// Output: integer
// An "alphabetic character" means a charachter a-z/A-Z
// Edge cases: an empty string, no letters at all (mixed content (digits, spaces) are just non-matches, not special cases)

// Pseudocode = We initialize a counter starting at 0, and use a for loop , looping charachter by charachter, and add 1 to our counter if we encounter an alpha charachter. An alpha character needs to be compared with the range 'a' to 'z' or the range 'A' to 'Z'. At the end of the loop we just return the counters value.
// Logs: no need to check if the string is empty bcs the counter starts at 0 already. If you dont want to use index in index, charachter set index to _


import (
	"fmt"
)

func main() {
	fmt.Println(CountAlpha("Hello world"))
	fmt.Println(CountAlpha("H e l l o"))
	fmt.Println(CountAlpha("H1e2l3l4o"))
}


func CountAlpha(s string) int {
	counter := 0

	// 
	for _, character := range s {
		if character >= 'a' && character <= 'z' || character >= 'A' && character <= 'Z' {
			counter += 1
		}
	}

	return counter
}



