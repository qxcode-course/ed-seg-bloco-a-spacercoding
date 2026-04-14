package main
import "fmt"
func main() {
    var capacidade, movimentacao, qtdAtual int

    fmt.Scan(&capacidade)

    for {
        fmt.Scan(&movimentacao)

        qtdAtual += movimentacao

        if qtdAtual == 0 {

            fmt.Println("vazio")

        }else if qtdAtual < capacidade {

            fmt.Println("ainda cabe")

        }else if qtdAtual >= capacidade && qtdAtual < capacidade * 2 {

            fmt.Println("lotado")

        }else if qtdAtual >= capacidade * 2 {

            fmt.Println("hora de partir")
            break

        }

    }
}
