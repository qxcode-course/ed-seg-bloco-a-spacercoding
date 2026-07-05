package main

import (
	"bufio"
	"fmt"
	"os"
)

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
		fmt.Println(labirinto[i])
	}


}
