package bonus

//Implemente uma função que inverta
//uma [lista encadeada](https://www.ime.usp.br/~pf/algoritmos/aulas/lista.html#:~:text=Uma%20lista%20encadeada%20%C3%A9%20uma,segunda%2C%20e%20assim%20por%20diante),
//ou seja, altere a ordem dos elementos da lista de forma reversa.

type Node struct {
	Value int
	Next  *Node
}

type LinkedList struct {
	Head *Node
}

func InvertLinkedList(list *LinkedList) {
	var (
		currentNode *Node
		nextNode    *Node
	)

	currentNode = list.Head
	nextNode = list.Head.Next
	currentNode.Next = nil

	for nextNode != nil {
		temp := nextNode.Next
		nextNode.Next = currentNode
		currentNode = nextNode
		nextNode = temp
	}

	list.Head = currentNode
}
