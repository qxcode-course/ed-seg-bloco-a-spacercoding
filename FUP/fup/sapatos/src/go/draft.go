package main
import "fmt"
func main() {
    var num1, num2, soma int

    fmt.Scan(&num1, &num2)

    if num1 > num2 {fmt.Println("invalido")
    } else {

        for i := num1; i <= num2; i++ {

            if i % 6 == 0 {soma += i}

        }

        fmt.Println(soma)

    }
}
