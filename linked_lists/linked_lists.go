// kör med go run linked_lists/linked_lists.go
package main
import "fmt"

type Node struct {
	data int 
	next *Node 
}

// function for inserting 
func add_to_end (head *Node, value int) *Node {
	new_node := &Node{data: value, next: nil}

	current := head 
	for current.next != nil { // OBS: här VILL vi kolla .next, till skillnad från print-loopen!
	current = current.next
	}

	current.next = new_node
	return head // returnera hela listan igen med start fån head
}


func main () {

	node_3 := &Node{data: 15,next: nil}
	node_2 := &Node{data: 10,next: node_3}
	node_1 := &Node{data: 5,next: node_2} 

	head := node_1
	current_node := head

	head = add_to_end(head, 20) // enkelt anrop, ingen loop behövs här

	// traversera och skriv ut hela listan för att verifiera
	for current_node != nil {
		fmt.Println(current_node.data)
		current_node = current_node.next
	}
}
