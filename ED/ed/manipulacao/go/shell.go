package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"slices"
	"math"
)

func getMen(vet []int) []int {
	var sortedVet = make([]int, 0)

	for i := range vet {

		if vet[i] > 0 {

			sortedVet = append(sortedVet, vet[i])

		}

	}

	_ = vet

	return sortedVet
}

func getCalmWomen(vet []int) []int {
	var vetCalmWoman = make([]int, 0)

	for i := range vet {

		if vet[i] > -10 && vet[i] < 0 {

			vetCalmWoman = append(vetCalmWoman, vet[i])

		}

	}

	_ = vet
	return vetCalmWoman
}

func sortVet(vet []int) []int {

	slices.Sort(vet)

	return vet
}

func sortStress(vet []int) []int {

	for i := 0; i < len(vet) - 1; i++ {

		if math.Abs(float64(vet[i])) > math.Abs(float64(vet[i+1])) {

			var varaux int
			varaux = vet[i+1]
			vet[i+1] = vet[i]
			vet[i] = varaux

		} 

	}

	return vet
}

func reverse(vet []int) []int {
	var vetReverse = make([]int, 0)

	for i := len(vet) - 1; i > -1; i-- {

		vetReverse = append(vetReverse, vet[i])

	}

	_ = vet
	return vetReverse
}

func unique(vet []int) []int {
	var vetUnique = make([]int, 0)

	for i := range vet {

		if vet[i] != vetUnique[i] {

			vetUnique = append(vetUnique, vet[i])

		}

	}

	_ = vet
	return sortVet(vetUnique)
}

func repeated(vet []int) []int {
	var vetRepeated = make([]int, 0)

	sortVet(vet)

	for i := range len(vet) - 1 {

		if vet[i] == vet[i+1] {

			vetRepeated = append(vetRepeated, vet[i])

		}

	}

	_ = vet
	return vetRepeated
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	for {
		if !scanner.Scan() {
			break
		}
		fmt.Print("$")
		line := scanner.Text()
		args := strings.Split(line, " ")
		fmt.Println(line)

		switch args[0] {
		case "get_men":
			printVec(getMen(str2vet(args[1])))
		case "get_calm_women":
			printVec(getCalmWomen(str2vet(args[1])))
		case "sort":
			printVec(sortVet(str2vet(args[1])))
		case "sort_stress":
			printVec(sortStress(str2vet(args[1])))
		case "reverse":
			array := str2vet(args[1])
			other := reverse(array)
			printVec(array)
			printVec(other)
		case "unique":
			printVec(unique(str2vet(args[1])))
		case "repeated":
			printVec(repeated(str2vet(args[1])))
		case "end":
			return
		}
	}
}

func printVec(vet []int) {
	fmt.Print("[")
	for i, val := range vet {
		if i > 0 {
			fmt.Print(", ")
		}
		fmt.Print(val)
	}
	fmt.Println("]")
}

func str2vet(s string) []int {
	if s == "[]" {
		return nil
	}
	s = s[1 : len(s)-1]
	parts := strings.Split(s, ",")
	var vet []int
	for _, part := range parts {
		n, _ := strconv.Atoi(part)
		vet = append(vet, n)
	}
	return vet
}
