package main

import (
	"fmt"
	"strconv"
)

// mostra a lista com o elemento sword destacado
func ToStr(l *DList[int], sword *DNode[int]) string {
	var str string = "[ "

	for i := l.Front(); i != l.End(); i = i.Next() {

		if i.Value == sword.Value {
			str += strconv.Itoa(sword.Value) + "> "
		} else {
			str += strconv.Itoa(i.Value) + " "
		}

	}
	str += "]"

	return str
}

// move para frente na lista circular
func Next(l *DList[int], it *DNode[int]) *DNode[int] {
	if it.next == it.root {
		return it.root.next
	}
	return it.next
}

func main() {
	var qtd, chosen int
	fmt.Scan(&qtd, &chosen)
	l := NewDList[int]()
	for i := 1; i <= qtd; i++ {
		l.PushBack(i)
	}
	sword := l.Front()
	for range chosen - 1 {
		sword = Next(l, sword)
	}
	for range qtd - 1 {
		fmt.Println(ToStr(l, sword))
		l.Erase(Next(l, sword))
		sword = Next(l, sword)
	}
	fmt.Println(ToStr(l, sword))
}
