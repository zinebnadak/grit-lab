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
	next *Node // explicit pekare till en annan Node av typ pekare, precis som typen som görs när vi skapar en Node
}


func main () {
	// creating an instans of a Node this way

	node_3 := &Node{data: 15,next: nil}
	node_2 := &Node{data: 10,next: node_3}
	node_1 := &Node{data: 5,next: node_2} 

	// &Node{...} skapar en Node och & tar adressen till den, så n1 blir av typen *Node (en pekare), inte Node direkt.
	// Analogi: tänk dig ett hus (Node-structen med dess data) och dess gatuadress (pekaren *Node). 
	// Node{...} bygger huset. &Node{...} ger dig lappen med adressen till huset , inte huset självt.

	// &-operatorn: tar ett hus (en struct) och ger dig adressen till det (struct -> pekare)
	// *-operatorn (dereferensering): gör tvärtom ,tar en adress (pekare) och ger dig huset som ligger där (pekare -> struct)

	fmt.Println(node_1.data)
	
	fmt.Println(node_2.data) // en adress skapad med &

	fmt.Println(node_3.data)
	
}
