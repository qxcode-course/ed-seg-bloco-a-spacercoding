package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {

	Value int
	next *Node
	previous *Node

}
type LList struct {

	head *Node

}

func (ll *LList) Front() *Node { return ll.head.next }
func (ll *LList) Back() *Node { return ll.head.previous }
func (ll *LList) End() *Node { return ll.head }

func (node *Node) Next() *Node { return node.next }
func (node *Node) Previous() *Node { return node.previous }

func NewLList() *LList {
	var newLList *LList

	newLList = &LList{}
	newLList.head = &Node{}

	newLList.head.next = newLList.head
	newLList.head.previous = newLList.head

	return newLList
}

func (ll *LList) Size() int {
	var cont int

	for i := ll.Front(); i != ll.End(); i = i.Next() {

		cont++

	}

	return cont
}
func (ll *LList) Search(value int) *Node {

	for i := ll.Front(); i != ll.End(); i = i.next {if i.Value == value {return i}}

	return nil
}

func insert(node *Node, value int) {
	var newNode *Node = &Node{Value: value, next: nil, previous: nil}

	newNode.Value = value
	newNode.next = node
	newNode.previous = node.previous

	node.previous.next = newNode
	node.previous = newNode
}
func (ll *LList) PushFront(value int) { insert(ll.Front(), value) }
func (ll *LList) PushBack(value int) { insert(ll.End(), value) }
func (ll *LList) Insert(node *Node, value int) { insert(node, value) }

func remove(node *Node) {
	
	node.next.previous = node.previous
	node.previous.next = node.next
	
	node.next = nil
	node.previous = nil
	
}
func (ll *LList) PopFront() { if ll.Size() != 0 {remove(ll.Front())} }
func (ll *LList) PopBack() { if ll.Size() != 0 {remove(ll.Back())} }
func (ll *LList) Clear() { for range ll.Size() {remove(ll.Front())} }
func (ll *LList) Remove(node *Node) { remove(node) }



func (ll *LList) String() string {
	var str string = "["

	if ll.Size() != 0 {

		for i := ll.Front(); i != ll.Back(); i = i.Next() {str += strconv.Itoa(i.Value) + ", "}
		str += strconv.Itoa(ll.Back().Value)

	}
	str += "]"

	return str
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := NewLList()

	for {
		fmt.Print("$")
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		fmt.Println(line)
		args := strings.Fields(line)

		if len(args) == 0 {
			continue
		}

		cmd := args[0]

		switch cmd {
		case "show":
			fmt.Println(ll.String())
		case "size":
			fmt.Println(ll.Size())
		case "push_back":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushBack(num)
			}
		case "push_front":
			for _, v := range args[1:] {
				num, _ := strconv.Atoi(v)
				ll.PushFront(num)
			}
		case "pop_back":
			ll.PopBack()
		case "pop_front":
			ll.PopFront()
		case "clear":
			ll.Clear()
		case "walk":
			fmt.Print("[ ")
			for node := ll.Front(); node != ll.End(); node = node.Next() {
				fmt.Printf("%v ", node.Value)
			}
			fmt.Print("]\n[ ")
			for node := ll.Back(); node != ll.End(); node = node.Previous() {
				fmt.Printf("%v ", node.Value)
			}
			fmt.Println("]")
		case "replace":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				node.Value = newvalue
			} else {
				fmt.Println("fail: not found")
			}
		case "insert":
			oldvalue, _ := strconv.Atoi(args[1])
			newvalue, _ := strconv.Atoi(args[2])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Insert(node, newvalue)
			} else {
				fmt.Println("fail: not found")
			}
		case "remove":
			oldvalue, _ := strconv.Atoi(args[1])
			node := ll.Search(oldvalue)
			if node != nil {
				ll.Remove(node)
			} else {
				fmt.Println("fail: not found")
			}
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
