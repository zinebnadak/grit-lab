// Input: takes a string
// Output: modified string displaying it repeating each alphabetical character as many times as its alphabetical index
/* Pseodo code: takes a string as input and returns a modified one. I will initialise an empty string, and then loop through the whole string char by char appending to the new string on following conditions: Non-letters are kept exactly as they are, only letters get repeated depending on what position they are in the alphabet (a = 1,, z = 26), and this counts for both lowercase and uppercase letters, which keep their case. I will know the letters position in the alphabet (number we will use) by counting this way:

- IF IT IS LOWERCASE how far the letter is from the rune 'a' then adding 1
- IF IT IS UPPERCASE how far the letter is from the rune ‘A’ then adding 

Last I return the string variable's value at the end

*/


package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.RepeatAlpha("abc"))
	fmt.Println(piscine.RepeatAlpha("Choumi."))
	fmt.Println(piscine.RepeatAlpha(""))
	fmt.Println(piscine.RepeatAlpha("abacadaba 01!"))
}


func RepeatAlpha(s string) string {
}






































































package main

import (
	"fmt"
	"piscine"
)

func main() {
	fmt.Println(piscine.RepeatAlpha("abc"))
	fmt.Println(piscine.RepeatAlpha("Choumi."))
	fmt.Println(piscine.RepeatAlpha(""))
	fmt.Println(piscine.RepeatAlpha("abacadaba 01!"))
}


func RepeatAlpha(s string) string {
}