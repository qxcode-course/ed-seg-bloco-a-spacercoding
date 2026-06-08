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
	next  *Node
	prev  *Node
	root  *Node
}

type LList struct {
	root *Node
	size int
}

func NewLList() *LList {
	list := &LList{}
	list.root = &Node{root: nil}
	list.root.next = list.root
	list.root.prev = list.root
	list.root.root = list.root // nó sentinela aponta pra si mesmo
	return list
}

func (l *LList) PushBack(value int) {
	l.insertBefore(l.root, value)
	l.size++
}

func (l *LList) insertBefore(mark *Node, value int) {
	n := &Node{
		Value: value,
		root:  l.root,
	}
	n.prev = mark.prev
	n.next = mark
	mark.prev.next = n
	mark.prev = n
}
func (l *LList) insertAfter(mark *Node, value int) {
	n := &Node{
		Value: value,
		root:  l.root,
	}
	n.prev = mark
	n.next = mark.next
	mark.next.prev = n
	mark.next = n
}

func addsorted(l *LList, value int) {

	for i := l.root.next; i != l.root; i = i.next {

		if value < i.Value {
			l.insertBefore(i, value)
			l.size++
			return
		}

	}
	l.PushBack(value)

}

func equals(l1, l2 *LList) bool {
	if l1.size != l2.size {return false}

	for i, j:= l1.root.next, l2.root.next; i != l1.root; i, j = i.next, j.next {

		if i.Value != j.Value {return false}

	}

	return true
}

func reverse(l *LList) {
	if l.size == 0 {return}
	var cont int
	var auxNode *Node

	for ; cont < l.size / 2; cont++ {

		

	}

}

func str2list(serial string) *LList {
	serial = serial[1 : len(serial)-1]
	ll := NewLList()
	if serial == "" {
		return ll
	}
	for _, p := range strings.Split(serial, ",") {
		value, _ := strconv.Atoi(p)
		ll.PushBack(value)
	}
	return ll
}
func (l *LList)list2str() string {
	var str string = "["

	for i := l.root.next; i != l.root.prev; i = i.next {

		str += strconv.Itoa(i.Value) + ", "

	}
	if l.size != 0 {
	str += strconv.Itoa(l.root.prev.Value)
	}
	str  += "]"

	return str
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

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
		case "compare":
			lla := str2list(args[1])
			llb := str2list(args[2])
			if equals(lla, llb) {
				fmt.Println("iguais")
			} else {
				fmt.Println("diferentes")
			}
		case "addsorted":
			lla := NewLList()
			for i := 1; i < len(args); i++ {
				value, _ := strconv.Atoi(args[i])
				addsorted(lla, value)
			}
			fmt.Println(lla.list2str())
		case "reverse":
			lla := str2list(args[1])
			reverse(lla)
			fmt.Println(lla.list2str())
		case "merge":
			// lla := str2list(args[1])
			// llb := str2list(args[2])
			// merged := merge(lla, llb)
			// fmt.Println(merged)
		case "end":
			return
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
