package main

import "fmt"

func estimarEstado(idade int) (string) {

    var estado string

    if idade <= 12 {estado = "crianca"
} else if idade <= 18 {estado = "jovem"
} else if idade < 65 {estado = "adulto"
} else if idade < 1000 {estado = "idoso"
} else {estado = "mumia"}

    return estado

}

func main() {

    var nome string
    var idade int

    fmt.Scan(&nome, &idade)

    fmt.Printf("%s eh %s\n", nome, estimarEstado(idade))

}
