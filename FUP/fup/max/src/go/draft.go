package main

import "fmt"

func qualOMaior(num1, num2 int) int {

    if num1 > num2 {return num1}

    return num2
}

func main() {
    var num1, num2 int

    fmt.Scan(&num1, &num2)

    fmt.Println(qualOMaior(num1, num2))
}
