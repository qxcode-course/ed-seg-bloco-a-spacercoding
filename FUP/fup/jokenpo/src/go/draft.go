package main

import "fmt"

func resultado(jogada1, jogada2 string) string {

    if (jogada1 == "R" && jogada2 == "S") || (jogada1 == "S" && jogada2 == "P") || (jogada1 == "P" && jogada2 == "R") {return "jog1"
    }else if (jogada2 == "R" && jogada1 == "S") || (jogada2 == "S" && jogada1 == "P") || (jogada2 == "P" && jogada1 == "R") {return "jog2"}

    return "empate"
}

func main() {
    var jogada1, jogada2 string

    fmt.Scan(&jogada1, &jogada2)

    fmt.Println(resultado(jogada1, jogada2))
}
