/*
Rekursion används i BST naturnligt eftersom ett träd är uppbyggt av mindre träd (subträd) 
som i sin tur är uppbyggda av ännu mindre träd.
*/

type Node struct {
	data int
	left *Node
	right *Node
}

/*
BST-regeln (Binary Search Tree property): för varje nod gäller att allt i dess vänstra subträd är mindre än noden
, och allt i dess högra subträd är större eller lika.
Exempel (från videon): rot = 4, vänster subträd innehåller 1,2,3 (alla mindre än 4), höger subträd innehåller 5,6,7 (alla större)
*/