package main

import (
	"fmt"
)

func mdc(a, b int) int {

	if a == 0 {return b
	}else if b == 0 {return a}

	varAux := a % b

	return mdc(b, varAux)
}

func main() {
	var a, b int
	fmt.Scan(&a, &b)
	fmt.Println(mdc(a, b))
}
