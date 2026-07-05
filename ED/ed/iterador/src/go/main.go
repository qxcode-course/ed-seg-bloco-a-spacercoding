package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type MyList struct {
	data []int
}

type Iterator struct {
	data  []int
	index int
}

type CyclicIterator struct {

	data []int
	index int

}

func NewMyList(values []int) *MyList {
	return &MyList{data: values}
}

func (l *MyList) Iterator() *Iterator {
	return &Iterator{data: l.data, index: -1}
}

func (i *Iterator) HasNext() bool {
	return i.index < len(i.data)-1
}

func (ci *CyclicIterator) Next() int {
	if ci.index == len(ci.data) -1 {
		ci.index = 0
		return ci.data[0]
	}

	ci.index++

	return ci.data[ci.index]
}

func (i *Iterator) Next() int {
	if i.index == len(i.data) {
		panic(fmt.Errorf("No more elements"))
	}
	i.index += 1
	return i.data[i.index]
}

func (list *MyList) ReverseIterator() Iterator {
	var newVet []int = make([]int, len(list.data))

	for i := len(list.data) -1; i > -1; i-- {

		newVet[i] = list.data[(len(list.data) -1) - i]

	}

	return Iterator{data: newVet, index: -1}
}

func (list *MyList) CyclicIterator() CyclicIterator {

	return *&CyclicIterator{data: list.data, index: -1}

}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	mylist := NewMyList([]int{})
	for scanner.Scan() {
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			break
		case "read":
			for i := 1; i < len(args); i++ {
				slice := make([]int, len(args)-1)
				for i, value := range args[1:] {
					slice[i], _ = strconv.Atoi(value)
				}
				mylist = NewMyList(slice)
			}
		case "show":
			fmt.Print("[ ")
			for it := mylist.Iterator(); it.HasNext(); {
				fmt.Printf("%v ", it.Next())
			}
			fmt.Println("]")
		case "reverse":
			fmt.Print("[ ")
			for it := mylist.ReverseIterator(); it.HasNext(); {
				fmt.Printf("%v ", it.Next())
			}
			fmt.Println("]")
		case "cyclic":
			qtd, _ := strconv.Atoi(args[1])
			fmt.Print("[ ")
			it := mylist.CyclicIterator()
			for range qtd {
				fmt.Printf("%v ", it.Next())
			}
			fmt.Println("]")
		 }
	}

}
