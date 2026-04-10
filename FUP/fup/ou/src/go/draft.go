package main

import "fmt"

func verificacao3Ou5(num int) string {

    if num == 3 || num == 5 {return "SIM"}

    return "NAO"
}

func main() {
    var num int

    fmt.Scan(&num)

    fmt.Println(verificacao3Ou5(num))
}
