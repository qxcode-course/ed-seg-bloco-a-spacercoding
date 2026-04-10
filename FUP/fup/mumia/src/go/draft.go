package main

import (
    "fmt"
)

func status(nome string, idade int) string {

    if idade < 12 {return nome + " eh crianca"
    }else if idade < 18 {return nome + " eh jovem"
    }else if idade < 65 {return nome + " eh adulto"
    }else if idade < 1000 {return nome + " eh idoso"}

    return nome + " eh mumia" 
}

func main() {
    var nome string
    var idade int

    fmt.Scan(&nome, &idade)

    fmt.Println(status(nome, idade))
}
