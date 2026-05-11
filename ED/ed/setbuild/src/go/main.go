package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type VetOrdenado struct {
	data           []int
	size, capacity int
}

func NewSet(value int) VetOrdenado {
	var v VetOrdenado

	v.capacity = value
	v.data = make([]int, v.capacity)
	v.size = 0

	return v
}
func (v *VetOrdenado) reserve(value int) {
	var vetAux []int = make([]int, value)

	v.capacity = value
	for i := range v.size {

		vetAux[i] = v.data[i]

	}
	v.data = vetAux

}
func (v *VetOrdenado) binarySearch(value int) int {

	for i := range v.size {

		if v.data[i] == value {
			return i
		}

	}

	return -1
}

func (v *VetOrdenado) Insert(value int) {
	if !v.Contains(value) {

		v.reserve(v.capacity + 1)
		v.data[v.size] = value

	}

	var index int

	for i := range v.size {

		if v.data[i] > value {

			index = i
			break

		}

	}

	v.reserve(v.capacity + 1)
	for i := v.size; i > index; i-- {

		v.data[i] = v.data[i-1]

	}
	v.data[index] = value

}
func (v *VetOrdenado) Contains(value int) bool {

	for i := range v.size {

		if v.data[i] == value {
			return true
		}

	}

	return false
}
func (v *VetOrdenado) Erase(value int) bool {
	if !v.Contains(value) {
		return false
	}

	var index int

	for i := range v.size {

		if v.data[i] > value {

			index = i
			break

		}

	}

	for i := index; i < v.size; i++ {

		v.data[i] = v.data[i+1]

	}
	v.reserve(v.capacity - 1)
	v.size--

	return true
}
func (v *VetOrdenado) String() string {
	if v.size == 0 {
		return "[]"
	}
	var s string = "["

	for i := range v.size - 1 {

		s += strconv.Itoa(v.data[i]) + ","

	}
	s += strconv.Itoa(v.data[v.size-1]) + "]"

	return s
}

func main() {
	var line, cmd string
	scanner := bufio.NewScanner(os.Stdin)

	v := NewSet(0)
	for scanner.Scan() {
		fmt.Print("$")
		line = scanner.Text()
		fmt.Println(line)
		parts := strings.Fields(line)
		if len(parts) == 0 {
			continue
		}
		cmd = parts[0]

		switch cmd {
		case "end":
			return
		case "init":
			value, _ := strconv.Atoi(parts[1])
			v = NewSet(value)
		case "insert":
			for _, part := range parts[1:] {
				value, _ := strconv.Atoi(part)
				v.Insert(value)
			}
		case "show":
			fmt.Println(v.String())
		case "erase":
			value, _ := strconv.Atoi(parts[1])
			v.Erase(value)
		case "contains":
			value, _ := strconv.Atoi(parts[1])
			if !v.Erase(value) {
				fmt.Println("value not found")
			}
		case "clear":
			for i := v.size; i > 0; i-- {
				v.Erase(v.data[i])
			}
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
