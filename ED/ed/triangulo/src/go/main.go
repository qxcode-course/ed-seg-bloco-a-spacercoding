package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func toStInt(vet []int) {

	fmt.Print("[ ")
	for i := range vet {

		fmt.Printf("%d ", vet[i])

	}
	fmt.Println("]")

}

func processa(vet []int) {
	var vetSoma = make([]int, 0)

	if len(vet) == 1 {

		toStInt(vet)
		return

	} else {

		for i := range len(vet) - 1 {

			vetSoma = append(vetSoma, vet[i]+vet[i+1])
			i++

		}

		processa(vetSoma)

	}
	toStInt(vet)

	// 1. defina o ponto de parada
	// 2. monte o vetor auxiliar com os resultados das somas
	// 3. chame recursivamente a função processa para o vetor auxiliar
	// 4. imprima o vetor original
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	if !scanner.Scan() {
		return
	}
	line := scanner.Text()
	parts := strings.Fields(line)
	vet := []int{}
	for _, part := range parts {
		if value, err := strconv.Atoi(part); err == nil {
			vet = append(vet, value)
		}
	}
	processa(vet)
}

func Join[T any](v []T, sep string) string {
	if len(v) == 0 {
		return ""
	}
	s := ""
	for i, x := range v {
		if i > 0 {
			s += sep
		}
		s += fmt.Sprintf("%v", x)
	}
	return s
}
