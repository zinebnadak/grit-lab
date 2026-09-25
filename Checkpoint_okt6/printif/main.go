// Input: a string 
// Output: returns either "G" and a new line or "Invalid input" 
// Pseudocode: 
// Edge: If it's an empty string return G followed by a newline \n.

package main

import (
	"fmt"
)

func main() {
	fmt.Print(PrintIf("abcdefz"))
	fmt.Print(PrintIf("abc"))
	fmt.Print(PrintIf(""))
	fmt.Print(PrintIf("14"))
}


func PrintIf(str string) string {
	if len(str) >= 3 || len(str) == 0 { 
		return "G\n"
	} else {
		return "Invalid Input\n"
	}
}