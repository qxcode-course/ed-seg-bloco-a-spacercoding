package main

import "fmt"

func calcFilhoMaisVelho(idadeMonica, idadeFilho1, idadeFilho2 int) int {
    var idadeFilho3 int = idadeMonica - (idadeFilho1 + idadeFilho2)
    
    if idadeFilho1 > idadeFilho2 && idadeFilho1 > idadeFilho3 {return idadeFilho1
    }else if idadeFilho2 > idadeFilho1 && idadeFilho2 > idadeFilho3 {return idadeFilho2
    }else if idadeFilho3 > idadeFilho2 && idadeFilho3 > idadeFilho1 {return idadeFilho3}

    return 0
}

func main() {
    var idadeMonica, idadeFilho1, idadeFilho2 int

    fmt.Scan(&idadeMonica, &idadeFilho1, &idadeFilho2)

    fmt.Println(calcFilhoMaisVelho(idadeMonica, idadeFilho1, idadeFilho2))
}
