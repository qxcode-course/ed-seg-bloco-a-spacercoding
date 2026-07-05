package main
import "fmt"
func main() {
    var total int
    var equacao string

    fmt.Scan(&equacao)

    for i := range equacao {
        if i != len(equacao) -1 && (equacao[i] == '(' && equacao[i+1] == ']') {continue}
        if i == len(equacao) -1 && (equacao[i] == '(' || equacao[i] == '[') {continue}

        if equacao[i] == '(' {total++
        } else if equacao[i] == ')' {total--
        } else if equacao[i] == '[' {total += 2
        } else {total -= 2}

    }
    if total == 0 {fmt.Println("balanceado")
    } else {fmt.Println("nao balanceado")}
}