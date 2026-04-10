package main

import "fmt"

func verificacaoPositivo(num int) string {

    if num >= 0 {return "SIM"}

    return "NAO"
}

func main() {
    var num int

    fmt.Scan(&num)

    fmt.Println(verificacaoPositivo(num))
}
