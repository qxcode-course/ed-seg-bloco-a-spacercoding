package main

import "fmt"

func verificacaoMultiplo7(num int) string {

    if num % 7 == 0 {return "SIM"}

        return "NAO"
}

func main() {
    var num int

    fmt.Scan(&num)

    fmt.Println(verificacaoMultiplo7(num))
}
