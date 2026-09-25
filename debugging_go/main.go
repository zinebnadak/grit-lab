package main

import "fmt"

const string_1 = "Hello"
var new_string := ""
for _, charachter := range string_1 {
	var repeatNum = (charachter - 'a') + 1
	new_string += charachter * repeat_num 
}



func main() {
	fmt.Println(repeatNum) // 3
}

/*
func main(s string) string {
	new_string := ""

	for _, charachter := range s {

	}

	return new_string 
}
*/

// FILE FOR DEBUGGING 