package main

import "fmt"

type Node struct {
	data     int
	next     *Node
	previous *Node
}
type List struct{ head *Node }

// inicializando uma nova lista
func NewList() *List {
	//criação de uma nova Lista
	var newList *List = &List{}

	//nascimento da cabeça
	newList.head = &Node{}

	//a cabeça vai ser o ponto fixo, então ela precisa ser o primeiro e o último elemento
	//preenchendo a cabeça
	newList.head.next = newList.head
	newList.head.previous = newList.head

	return newList
}

func insert(node *Node, value int) {
	var newNode *Node = &Node{data: value, next: nil, previous: nil}

	//endereçamento do elemento a ser inserido
	newNode.next = node
	newNode.previous = node.previous

	//atualização dos elementos da lisa
	node.previous.next = newNode //o próximo do anterior agora é o novo node
	//A ORDEM IMPORTA
	node.previous = newNode
}
func (list *List) PushFront(value int) { insert(list.Front(), value) }
func (list *List) PushBack(value int)  { insert(list.End(), value) }

func remove(node *Node) {
	//o próximo se liga ao anterior
	node.next.previous = node.previous

	//o anterior se liga ao próximo
	node.previous.next = node.next

	//uma vez isolado, esqueça todos os seus vínculos
	node.next = nil
	node.previous = nil
}
func (list *List) PopFront() { remove(list.Front()) }
func (list *List) PopBack()  { remove(list.Back()) }

// abstração do primeiro elemento
// abstração do último elemento
// abstração do fim da lista
func (list *List) Front() *Node { return list.head.next }
func (list *List) Back() *Node  { return list.head.previous }
func (list *List) End() *Node   { return list.head }

func println(list *List) {

	fmt.Print("[ ")
	for i := list.Front(); i != list.End(); i = i.next {

		fmt.Print(i.data, " ")

	}
	fmt.Println("]")

}

func main() {
	var myList List = *NewList()

	//primeira inserção é após a cabeça, já que a lista só tem ela
	myList.PushBack(1)

	//inserindo no fim
	myList.PushBack(2)

	//inserindo no fim
	myList.PushBack(3)

	//inserindo no começo
	myList.PushFront(-1)

	println(&myList)
}
