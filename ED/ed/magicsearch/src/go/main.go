package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func searchLastOccurance(value int, slice []int) int {

	for i := len(slice) - 1; i >= 0; i-- {

		//fmt.Println("procurar")
		if slice[i] == value {return i}

	}

	return -1
}

func allocateIn(slice []int, value int) int {

	for i := range slice {

		if value < slice[i] {

			//fmt.Println("alocar")
			return i

		}

	}

	return len(slice)
}

func MagicSearch(slice []int, value int) int {
	_, _ = slice, value

	if len(slice) == 0 {return 0}

	var inicio, meio, fim int
	inicio = 0
	meio = (len(slice) -1 )/2
	fim = len(slice) -1

	for inicio <= fim{
		meio = (inicio + fim) /2

		if slice[meio] == value {

			return searchLastOccurance(value, slice)
		
		}

		if slice[meio] > value {

			fim = meio - 1

		} else {

			inicio = meio + 1
		}

	}

	return allocateIn(slice, value)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Fields(scanner.Text())
	slice := make([]int, 0, 1)
	for _, elem := range parts[1 : len(parts)-1] {
		value, _ := strconv.Atoi(elem)
		slice = append(slice, value)
	}

	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	result := MagicSearch(slice, value)
	fmt.Println(result)
}
