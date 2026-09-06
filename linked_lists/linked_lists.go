// kör med go run linked_lists/linked_lists.go
package main
import "fmt"

// nil i Go, None i Python
// define new type with "type"
// Node is the name of the type 
// struct to define the structure of the type. Inside the attribute names you want the type to have and the Go types they are 
// no __init__ instructor needed
// this is just a declaration and does not need to be inside func main()

type Node struct {
	data int 
	next *Node 
}


func main () {

	node_3 := &Node{data: 15,next: nil}
	node_2 := &Node{data: 10,next: node_3}
	node_1 := &Node{data: 5,next: node_2} 

	head := node_1
	current_node := head

	// for-loop in GO: 
	for current_node != nil { // current_node istället för current_node.next pga off-by-one buggen
		fmt.Println(current_node.data)
		current_node = current_node.next
	}
}
