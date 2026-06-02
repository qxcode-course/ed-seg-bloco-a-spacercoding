package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Node struct {
	data     int
	next     *Node
	previous *Node
}
type LList struct {
	head *Node
}

func (ll *LList) Front() *Node { return ll.head.next }
func (ll *LList) Back() *Node  { return ll.head.previous }
func (ll *LList) End() *Node   { return ll.head }

func NewLList() *LList {
	var newLList *LList

	newLList = &LList{}
	newLList.head = &Node{}

	newLList.head.next = newLList.head
	newLList.head.previous = newLList.head

	return newLList
}

func insert(value int, node *Node) {
	var newNode *Node = &Node{data: value, next: nil, previous: nil}

	newNode.next = node
	newNode.previous = node.previous

	node.previous.next = newNode
	node.previous = newNode
}
func (ll *LList) PushFront(value int) { insert(value, ll.Front()) }
func (ll *LList) PushBack(value int)  { insert(value, ll.End()) }

func erase(node *Node) {

	node.next.previous = node.previous
	node.previous.next = node.next

	node.next = nil
	node.previous = nil

}
func (ll *LList) PopFront() { if ll.Size() != 0 {erase(ll.Front())} }
func (ll *LList) PopBack()  { if ll.Size() != 0 {erase(ll.Back())} }
func (ll *LList) Clear() {

	for range ll.Size() {

		erase(ll.Front())

	}

}

func (ll *LList) Size() int {
	var count int

	for i := ll.Front(); i != ll.End(); i = i.next { count++ }

	return count
}

func (ll *LList) String() string {
	var str string = "["

	if ll.Size() != 0 {
	for i := ll.Front(); i != ll.Back(); i = i.next { str += strconv.Itoa(i.data) + ", " }
	
	str += strconv.Itoa(ll.Back().data)
	}
	str += "]"

	return str
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	ll := *NewLList()

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
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
