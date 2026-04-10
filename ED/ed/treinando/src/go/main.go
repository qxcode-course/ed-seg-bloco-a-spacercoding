package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func tostr(vet []int) string {
	var vetAux = make([]int, 0)
	var indexString string

	if len(vet) == 0 {

		return "[]"

	}

	if len(vet) == 1 {

		indexString = "[" + strconv.Itoa(vet[0])
		return indexString + "]"

	}

	vetAux = append(vetAux, vet[1:]...)
	indexString = strconv.Itoa(vet[0])
	tostr(vetAux)

	return indexString + ", "
}

func tostrrev(vet []int) string {
	var vetAux = make([]int, 0)
	var indexString string

	if len(vet) == 0 {

		return "[]"

	}

	if len(vet) == 1 {

		indexString = "[" + strconv.Itoa(vet[0])
		return indexString + "]"

	}

	vetAux = append(vetAux, vet[:len(vet)-1]...)
	indexString = strconv.Itoa(vet[len(vet)-1])
	tostrrev(vetAux)

	return ", " + indexString
}

// reverse: inverte os elementos do slice
func reverse(vet []int) {
	if len(vet) == 0 || len(vet) == 1 {
		return
	}

	reverse(vet)
}

// sum: soma dos elementos do slice
func sum(vet []int) int {
	if len(vet) == 0 {
		return 0
	}
	if len(vet) == 1 {
		return vet[0]
	}

	vet[1] += vet[0]
	vet = append(vet, vet[1:]...)

	return sum(vet)
}

// mult: produto dos elementos do slice
func mult(vet []int) int {
	if len(vet) == 0 {
		return 0
	}
	if len(vet) == 1 {
		return vet[0]
	}

	vet[1] = vet[1] * vet[0]
	vet = append(vet, vet[1:]...)

	return mult(vet)
}

// min: retorna o índice e valor do menor valor
// crie uma função recursiva interna do modelo
// var rec func(v []int) (int, int)
// para fazer uma recursão que retorna valor e índice
func min(vet []int) int {
	if len(vet) == 0 {
		return 0
	}
	if len(vet) == 1 {
		return vet[0]
	}

	if vet[1] <= vet[0] {

		vet = append(vet, vet[1:]...)

	} else {

		vet[1] = vet[0]

	}

	return min(vet)
}

func main() {
	var vet []int
	scanner := bufio.NewScanner(os.Stdin)
	for {
		if !scanner.Scan() {
			break
		}
		line := scanner.Text()
		args := strings.Fields(line)
		fmt.Println("$" + line)

		switch args[0] {
		case "end":
			return
		case "read":
			vet = nil
			for _, arg := range args[1:] {
				if val, err := strconv.Atoi(arg); err == nil {
					vet = append(vet, val)
				}
			}
		case "tostr":
			fmt.Println(tostr(vet))
		case "torev":
			fmt.Println(tostrrev(vet))
		case "reverse":
			reverse(vet)
		case "sum":
			fmt.Println(sum(vet))
		case "mult":
			fmt.Println(mult(vet))
		case "min":
			fmt.Println(min(vet))
		default:
			fmt.Println("fail: comando invalido")
		}
	}
}
