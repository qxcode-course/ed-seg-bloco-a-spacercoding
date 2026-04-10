package main

import "fmt"

func valorAbsoluto(num1, num2 int) int {

    if num1 > num2 {return num1 - num2}

    return num2 - num1
}

func main() {
    var num1, num2 int

    fmt.Scan(&num1, &num2)

    fmt.Println(valorAbsoluto(num1, num2))
}
