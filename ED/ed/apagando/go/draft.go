package main

import (
	"fmt"
)

func atribuirInt(endereco *int) {

	fmt.Scan(endereco)

}

type Node struct {

	data int
	next *Node

}

func novoNode(nodeValue int) *Node {

	return &Node{data: nodeValue, next: nil}

}

type linkedList struct {

	head *Node

}

func instancearLL() *linkedList {

	return &linkedList{head: nil}

}


func main() {
	var qtd_pessoas int

	atribuirInt(&qtd_pessoas)

	var fila linkedList
	fila = *instancearLL()

	for range qtd_pessoas {
		var x int
		atribuirInt(&x)

		if fila.head == nil {

			fila.head = novoNode(x)

		} else {

			fila.head.next = novoNode(x)


		}


	}
	fmt.Println(fila.head.data)

}
