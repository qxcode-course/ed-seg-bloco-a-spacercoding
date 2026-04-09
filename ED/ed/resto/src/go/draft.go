package main

import "fmt"

func divPor2(num int) {
	var resultado, resto int
    
    if num == 0 {
    
    return
    
    }

	resultado = num / 2
	resto = num % 2
	divPor2(resultado)

	fmt.Println(resultado, resto)
}

func main() {
	var num int

	fmt.Scan(&num)

	divPor2(num)
}
