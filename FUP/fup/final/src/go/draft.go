package main

import "fmt"

func calcularMedia(nota1, nota2, nota_final int) string {
    var media int = (nota1 + nota2) / 2

    if media >= 7 {return "aprovado"
    }else if media < 4 {return "reprovado"} 

    if (media + nota_final) / 2 >= 5 {return "aprovado na final"}

    return "reprovado na final"
}

func main() {
    var nota1, nota2, nota_final int

    fmt.Scan(&nota1, &nota2, &nota_final)

    fmt.Println(calcularMedia(nota1, nota2, nota_final))
}
