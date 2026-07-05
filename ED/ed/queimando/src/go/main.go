package main

import (
	"bufio"
	//"crypto/dsa"
	"fmt"
	"os"
)

type Pos struct {

	l, c int

}

func burnTrees(grid [][]rune, l, c int) {
	
	stack := NewStack[Pos]()
	stack.Push(Pos{l, c})
	
	for !stack.IsEmpty() {
		bPos := stack.Pop()
		if grid[bPos.l][bPos.c] == '#' {
			
			grid[bPos.l][bPos.c] = 'o'

			if bPos.l + 1 < len(grid) {stack.Push(Pos{bPos.l+1, bPos.c})}
			if bPos.c + 1 < len(grid[0]) {stack.Push(Pos{bPos.l, bPos.c+1})}
			if bPos.l - 1 >= 0 {stack.Push(Pos{bPos.l-1, bPos.c})}
			if bPos.c - 1  >= 0 {stack.Push(Pos{bPos.l, bPos.c-1})}

		}
	}
	// Essa função deve usar uma list como pilha
	// e marcar as árvores na matriz como queimados
	// Uma sugestão de como fazer isso é:
	// - adicionar a primeira posição na pilha
	// - enquanto a pilha não estiver vazia:
	//   - retirar o elemento do topo
	//   - se puder ser queimado, queime e adicione seus vizinhos à pilha

	return
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	line := scanner.Text()
	var nl, nc, lfire, cfire int
	fmt.Sscanf(line, "%d %d %d %d", &nl, &nc, &lfire, &cfire)

	grid := make([][]rune, 0, nl)
	for range nl {
		scanner.Scan()
		line := []rune(scanner.Text())
		grid = append(grid, line)
	}
	burnTrees(grid, lfire, cfire)
	showGrid(grid)
}

func showGrid(mat [][]rune) {
	for _, linha := range mat {
		fmt.Println(string(linha))
	}
}
