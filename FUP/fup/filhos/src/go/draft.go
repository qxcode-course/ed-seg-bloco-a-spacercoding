package main
import "fmt"
func main() {
    var idadeMaisNovo, qtdFilhos int

    fmt.Scan(&idadeMaisNovo, &qtdFilhos)

    for i := idadeMaisNovo; i < idadeMaisNovo + (qtdFilhos * 2); i += 2 {

        fmt.Println(i)

    }
}
