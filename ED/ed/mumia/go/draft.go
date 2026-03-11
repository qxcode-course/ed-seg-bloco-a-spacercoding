package main

import "fmt"

func ageSort(nome string, idade int) string {

	if idade < 12 {
		return nome + " eh crianca\n"
	} else if idade < 18 {
		return nome + " eh jovem\n"
	} else if idade < 65 {
		return nome + " eh adulto\n"
	} else if idade < 1000 {
		return nome + " eh idoso\n"
	} else {
		return nome + " eh mumia\n"
	}

}

func main() {

	var nome string
	var idade int

	fmt.Scan(&nome, &idade)

	fmt.Print(ageSort(nome, idade))

}
