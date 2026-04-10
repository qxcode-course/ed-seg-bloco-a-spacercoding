package main

import "fmt"

func arredondar(operacao string, num float64) int {

    if operacao == "c" { return int(num) + 1
    } else if operacao == "r" {
        if num - float64(int(num)) >= 0.5 {return int(num) + 1
        }else {return int(num)}
    }

    return int(num)
}

func main() {
    var operacao string
    var num float64

    fmt.Scan(&operacao, &num)

    fmt.Println(arredondar(operacao, num))
}
