package main

import "fmt"

func qtdIguais(num1, num2, num3 int) int {

    if num1 == num2 && num2 == num3 {return 3
    }else if num1 == num2 || num1 == num3 || num2 == num3 {return 2}

    return 0
}

func main() {
    var num1, num2, num3 int

    fmt.Scan(&num1, &num2, &num3)

    fmt.Println(qtdIguais(num1, num2, num3))
}
