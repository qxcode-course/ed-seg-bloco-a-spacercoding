package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func ondeEraParaEstar(slice []int, value int) int {

	for i := range slice {

		if slice[i] > value {
			return i
		}

	}

	return len(slice)

}

func BetterSearch(slice []int, value int) (bool, int) {
	var inicio, meio, fim int = 0, (len(slice) - 1) / 2, len(slice) - 1

	for inicio <= fim {
		meio = (inicio + fim) / 2

		if value == slice[meio] {
			return true, meio
		}

		if value < slice[meio] {

			fim = meio - 1

		} else {

			inicio = meio + 1

		}

	}

	return false, ondeEraParaEstar(slice, value)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	parts := strings.Split(scanner.Text(), " ")
	slice := []int{}
	for _, elem := range parts[1 : len(parts)-1] {
		value, _ := strconv.Atoi(elem)
		slice = append(slice, value)
	}
	scanner.Scan()
	value, _ := strconv.Atoi(scanner.Text())
	found, result := BetterSearch(slice, value)
	if found {
		fmt.Println("V", result)
	} else {
		fmt.Println("F", result)
	}
}
