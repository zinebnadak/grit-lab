package main

import (
	"fmt"
	"myproject/mypackage"
)

func main() {
	message := mypackage.Greet()
	fmt.Println(message)
}


