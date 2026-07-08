package main

import (
	"bufio"
	"fmt"
	"os"
)

type Pos struct {
	l, c int
}

func getNeig(p Pos) []Pos {
	return []Pos{{p.l, p.c - 1}, {p.l - 1, p.c}, {p.l, p.c + 1}, {p.l + 1, p.c}}
}

func inside(grid [][]byte, p Pos) bool {
	return !(p.l < 0 || p.l >= len(grid) || p.c < 0 || p.c >= len(grid[0]))
}

func busca(grid [][]byte, pos Pos) {
    if !inside(grid, pos) || grid[pos.l][pos.c] == '0'{
        return
    }

    grid[pos.l][pos.c] = '0'

	varAux := getNeig(pos)

    busca(grid, varAux[0])
    busca(grid, varAux[1])
    busca(grid, varAux[2])
    busca(grid, varAux[3])
}

// Não modifique a assinatura da função numIslands
// Ela é a função que será chamada no LeetCode para resolver o problema
func numIslands(grid [][]byte) int {
	var qtdIlhas int

	for i := range grid {
        for j := range grid[0] {
            if grid[i][j] == '1' {
                qtdIlhas++
                busca(grid, Pos{i, j})
            }
        }
    }

	return qtdIlhas
}

// Não modifique a função main
func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc int
	fmt.Sscanf(line, "%d %d", &nl, &nc)
	grid := make([][]byte, nl)
	for i := 0; i < nl; i++ {
		scanner.Scan()
		grid[i] = []byte(scanner.Text())
	}
	result := numIslands(grid)
	fmt.Println(result)
}