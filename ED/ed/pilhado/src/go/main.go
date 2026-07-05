package main

import (
	"bufio"
	"fmt"
	"os"
)

func Resolver(labirinto [][]string) [][]string {
	caminho := NewStack[string]()
	limites := NewStack[string]()

	for i := range labirinto {
		for j := range labirinto[i] {
			if labirinto[i][j] == "I" {caminho.Push("I")}
		}
	}

	return labirinto
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var nLinha, nColuna int

	fmt.Scan(&nLinha, &nColuna)

	var labirinto [][]string = make([][]string, nLinha)

	for i := range nLinha {
		scanner.Scan()
		x := ""
		x = scanner.Text()
		labirinto[i] = append(labirinto[i], x)
		fmt.Println(Resolver(labirinto))
	}

}
