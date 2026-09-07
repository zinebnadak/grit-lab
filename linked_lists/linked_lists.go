package main
import "fmt"

// creating a type with two attributes: data and next

type Node struct {
	data int
	next *Node
}

// function for insertting at the end

func add_to_end (head *Node, value int) *Node {
	new_node := &Node{data: value, next: nil} // we set next to nil bcs it will be the last oen

	current := head
	for current.next != nil { // gå igenom hela listan tills vi med nil identify when we are at the end 
		current = current.next
	}

	current.next = new_node // efter for loopen klar blir sista nodesa next det nya värdet
	return head // return the new head so that we can assign the old head with the new one only by calling the function
}

// Creating the linked list: creating three instanses of the type: Node. Remember the referenced in next needs to be defined before its used, bcs Go reads from top to bottom. 
func main () {

	node_3 := &Node{data: 3, next: nil}
	node_2 := &Node{data: 2, next: node_3}
	node_1 := &Node{data: 1, next: node_2}

	/*
	fmt.Println(node_3) //skriver ut pekaren med adressen 
	fmt.Println(node_2.data) //skriver ut datan
	fmt.Println(node_1.next) // skriver ut nästa nodens pekare
	*/

	head := node_1
	head = add_to_end(head,4) //updating the head/linked list (with a new added element)

	current_node := head

	for current_node != nil {
		fmt.Println(current_node.data)
		current_node = current_node.next
	}
}


