package main

import "fmt"

func calcular(num1, num2 int, operador string) int {

    if operador == "+" {return num1 + num2
    }else if operador == "-" {return num1 - num2
    }else if operador == "*" {return num1 * num2}

    return num1 / num2
}

func main() {
    var num1, num2 int
    var operador string

    fmt.Scan(&num1, &num2, &operador)

    fmt.Println(calcular(num1, num2, operador))
}
