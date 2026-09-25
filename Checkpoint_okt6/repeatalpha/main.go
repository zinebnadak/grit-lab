// Input: takes a string
// Output: modified string displaying it repeating each alphabetical character as many times as its alphabetical index
/* Pseodo code: takes a string as input and returns a modified one. 
I will initialise an empty string
, and then loop through the whole string char by char appending to the new string on following conditions: 
Non-letters are kept exactly as they are, 
only letters get repeated depending on what position they are in the alphabet (a = 1,, z = 26), 
and this counts for both lowercase and uppercase letters, which keep their case. 
I will know the letters position in the alphabet (number we will use) by counting this way:

- IF IT IS LOWERCASE how far the letter is from the rune 'a' then adding 1
- IF IT IS UPPERCASE how far the letter is from the rune ‘A’ then adding 

Last I return the string variable's value at the end
*/

/* logs: 
for _, charachter := range s
In python "a" * 3 gives "aaa", but not in go!
Write else block like this: } else { or } else if  {
*/


package main

import (
	"fmt"
)

func main() {
	fmt.Println(RepeatAlpha("abc"))
	fmt.Println(RepeatAlpha("Choumi."))
	fmt.Println(RepeatAlpha(""))
	fmt.Println(RepeatAlpha("abacadaba 01!"))
}


func RepeatAlpha(s string) string {
	new_string := ""

	for _, charachter := range s {

		if charachter >= 'a' && charachter <= 'z' {
			repeat_num_l := (charachter - 'a') + 1 // find lowercase letters position in alphabet, and repeat corresponding times 
			for index := 0; index < int(repeat_num_l); index++ {
				new_string += string(charachter)
			}
		} else if charachter >= 'A' && charachter <= 'Z' {
			repeat_num_u := (charachter - 'A') + 1 // find uppercase letters position in alphabet, and repeat corresponding times
			for index := 0; index < int(repeat_num_u); index++ {
				new_string += string(charachter)
			}
		} else {
			new_string += string(charachter)
		}
	}
	return new_string 
}
